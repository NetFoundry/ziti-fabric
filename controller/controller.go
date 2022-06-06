/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel"
	"github.com/openziti/fabric/controller/api_impl"
	"github.com/openziti/fabric/controller/command"
	"github.com/openziti/fabric/controller/handler_ctrl"
	"github.com/openziti/fabric/controller/handler_mgmt"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/controller/raft"
	"github.com/openziti/fabric/controller/xctrl"
	"github.com/openziti/fabric/controller/xmgmt"
	"github.com/openziti/fabric/controller/xt"
	"github.com/openziti/fabric/controller/xt_random"
	"github.com/openziti/fabric/controller/xt_smartrouting"
	"github.com/openziti/fabric/controller/xt_weighted"
	"github.com/openziti/fabric/events"
	"github.com/openziti/fabric/health"
	fabricMetrics "github.com/openziti/fabric/metrics"
	"github.com/openziti/fabric/pb/ctrl_pb"
	"github.com/openziti/fabric/profiler"
	"github.com/openziti/foundation/common"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/metrics"
	"github.com/openziti/foundation/util/concurrenz"
	"github.com/openziti/storage/boltz"
	"github.com/openziti/xweb/v2"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	config             *Config
	network            *network.Network
	raftController     *raft.Controller
	ctrlConnectHandler *handler_ctrl.ConnectHandler
	mgmtConnectHandler *handler_mgmt.ConnectHandler
	xctrls             []xctrl.Xctrl
	xmgmts             []xmgmt.Xmgmt

	xwebFactoryRegistry xweb.Registry
	xweb                xweb.Instance

	ctrlListener channel.UnderlayListener
	mgmtListener channel.UnderlayListener

	shutdownC         chan struct{}
	isShutdown        concurrenz.AtomicBoolean
	agentHandlers     map[byte]func(conn *bufio.ReadWriter) error
	agentBindHandlers []channel.BindHandler
	metricsRegistry   metrics.Registry
	versionProvider   common.VersionProvider
}

func (c *Controller) GetId() *identity.TokenId {
	return c.config.Id
}

func (c *Controller) GetMetricsRegistry() metrics.Registry {
	return c.metricsRegistry
}

func (c *Controller) GetOptions() *network.Options {
	return c.config.Network
}

func (c *Controller) GetCommandDispatcher() command.Dispatcher {
	if c.raftController == nil {
		return nil
	}
	return c.raftController
}

func (c *Controller) GetDb() boltz.Db {
	return c.config.Db
}

func (c *Controller) GetMetricsConfig() *metrics.Config {
	return c.config.Metrics
}

func (c *Controller) GetVersionProvider() common.VersionProvider {
	return c.versionProvider
}

func (c *Controller) GetCloseNotify() <-chan struct{} {
	return c.shutdownC
}

func NewController(cfg *Config, versionProvider common.VersionProvider) (*Controller, error) {
	metricRegistry := metrics.NewRegistry(cfg.Id.Token, nil)

	c := &Controller{
		config:              cfg,
		shutdownC:           make(chan struct{}),
		xwebFactoryRegistry: xweb.NewRegistryMap(),
		metricsRegistry:     metricRegistry,
		versionProvider:     versionProvider,
	}

	if cfg.Raft != nil {
		raftController, err := raft.NewController(cfg.Id, cfg.Raft, metricRegistry)
		if err != nil {
			fmt.Printf("error starting raft %+v\n", err)
			panic(err)
		}

		c.raftController = raftController
		cfg.Db = raftController.GetDb()
	}

	c.registerXts()
	c.loadEventHandlers()

	if n, err := network.NewNetwork(c); err == nil {
		c.network = n
	} else {
		return nil, err
	}

	events.InitTerminatorEventRouter(c.network)
	events.InitRouterEventRouter(c.network)

	if cfg.Ctrl.Options.NewListener != nil {
		c.network.AddRouterPresenceHandler(&OnConnectSettingsHandler{
			config: cfg,
			settings: map[int32][]byte{
				int32(ctrl_pb.SettingTypes_NewCtrlAddress): []byte((*cfg.Ctrl.Options.NewListener).String()),
			},
		})
	}

	if err := c.showOptions(); err != nil {
		return nil, err
	}

	c.initWeb()

	return c, nil
}

func (c *Controller) initWeb() {
	healthChecker, err := c.initializeHealthChecks()
	if err != nil {
		logrus.WithError(err).Fatalf("failed to create health checker")
	}

	c.xweb = xweb.NewDefaultInstance(c.xwebFactoryRegistry, c.config.Id)

	if err := c.xweb.GetRegistry().Add(health.NewHealthCheckApiFactory(healthChecker)); err != nil {
		logrus.WithError(err).Fatalf("failed to create health checks api factory")
	}

	if err := c.xweb.GetRegistry().Add(api_impl.NewManagementApiFactory(c.config.Id, c.network, c.xmgmts)); err != nil {
		logrus.WithError(err).Fatalf("failed to create management api factory")
	}

	if err := c.RegisterXWebHandlerFactory(api_impl.NewMetricsApiFactory(c.config.Id, c.network, c.xmgmts)); err != nil {
		logrus.WithError(err).Fatalf("failed to create metrics api factory")
	}

}

func (c *Controller) Run() error {
	c.startProfiling()

	if err := c.registerComponents(); err != nil {
		return fmt.Errorf("error registering component: %s", err)
	}
	versionInfo := c.network.VersionProvider.AsVersionInfo()
	versionHeader, err := c.network.VersionProvider.EncoderDecoder().Encode(versionInfo)

	if err != nil {
		pfxlog.Logger().Panicf("could not prepare version headers: %v", err)
	}
	headers := map[int32][]byte{
		channel.HelloVersionHeader: versionHeader,
	}

	/**
	 * ctrl listener/accepter.
	 */
	ctrlChannelListenerConfig := channel.ListenerConfig{
		ConnectOptions:   c.config.Ctrl.Options.ConnectOptions,
		PoolConfigurator: fabricMetrics.GoroutinesPoolMetricsConfigF(c.network.GetMetricsRegistry(), "pool.listener.ctrl"),
		Headers:          headers,
	}
	ctrlListener := channel.NewClassicListener(c.config.Id, c.config.Ctrl.Listener, ctrlChannelListenerConfig)
	c.ctrlListener = ctrlListener
	if err := c.ctrlListener.Listen(c.ctrlConnectHandler); err != nil {
		panic(err)
	}

	ctrlAccepter := handler_ctrl.NewCtrlAccepter(c.network, c.xctrls, c.ctrlListener, c.config.Ctrl.Options.Options, c.config.Trace.Handler)
	go ctrlAccepter.Run()
	/* */

	/**
	 * mgmt listener/accepter.
	 */
	mgmtChannelListenerConfig := channel.ListenerConfig{
		ConnectOptions:   c.config.Mgmt.Options.ConnectOptions,
		PoolConfigurator: fabricMetrics.GoroutinesPoolMetricsConfigF(c.network.GetMetricsRegistry(), "pool.listener.mgmt"),
		Headers:          headers,
	}
	mgmtListener := channel.NewClassicListener(c.config.Id, c.config.Mgmt.Listener, mgmtChannelListenerConfig)
	c.mgmtListener = mgmtListener
	if err := c.mgmtListener.Listen(c.mgmtConnectHandler); err != nil {
		panic(err)
	}
	mgmtAccepter := handler_mgmt.NewMgmtAccepter(c.network, c.xmgmts, c.mgmtListener, c.config.Mgmt.Options)
	go mgmtAccepter.Run()

	if err := c.config.Configure(c.xweb); err != nil {
		panic(err)
	}

	go c.xweb.Run()

	// event handlers
	if err := events.WireEventHandlers(c.network.InitServiceCounterDispatch); err != nil {
		panic(err)
	}

	events.InitNetwork(c.network)

	c.network.Run()

	return nil
}

func (c *Controller) GetCloseNotifyChannel() <-chan struct{} {
	return c.shutdownC
}

func (c *Controller) Shutdown() {
	if c.isShutdown.CompareAndSwap(false, true) {
		close(c.shutdownC)

		if c.ctrlListener != nil {
			if err := c.ctrlListener.Close(); err != nil {
				pfxlog.Logger().WithError(err).Error("failed to close ctrl channel listener")
			}
		}

		if c.mgmtListener != nil {
			if err := c.mgmtListener.Close(); err != nil {
				pfxlog.Logger().WithError(err).Error("failed to close mgmt channel listener")
			}
		}

		if c.config.Db != nil {
			if err := c.config.Db.Close(); err != nil {
				pfxlog.Logger().WithError(err).Error("failed to close db")
			}
		}

		go c.xweb.Shutdown()
	}
}

func (c *Controller) showOptions() error {
	if ctrl, err := json.MarshalIndent(c.config.Ctrl.Options, "", "  "); err == nil {
		pfxlog.Logger().Infof("ctrl = %s", string(ctrl))
	} else {
		return err
	}
	if mgmt, err := json.MarshalIndent(c.config.Mgmt.Options, "", "  "); err == nil {
		pfxlog.Logger().Infof("mgmt = %s", string(mgmt))
	} else {
		return err
	}
	return nil
}

func (c *Controller) startProfiling() {
	if c.config.Profile.Memory.Path != "" {
		go profiler.NewMemoryWithShutdown(c.config.Profile.Memory.Path, c.config.Profile.Memory.Interval, c.shutdownC).Run()
	}
	if c.config.Profile.CPU.Path != "" {
		if cpu, err := profiler.NewCPUWithShutdown(c.config.Profile.CPU.Path, c.shutdownC); err == nil {
			go cpu.Run()
		} else {
			logrus.Errorf("unexpected error launching cpu profiling (%v)", err)
		}
	}
}

func (c *Controller) loadEventHandlers() {
	if e, ok := c.config.src["events"]; ok {
		if em, ok := e.(map[interface{}]interface{}); ok {
			for k, v := range em {
				if handlerConfig, ok := v.(map[interface{}]interface{}); ok {
					events.RegisterEventHandler(k, handlerConfig)
				}
			}
		}
	}
}

func (c *Controller) registerXts() {
	xt.GlobalRegistry().RegisterFactory(xt_smartrouting.NewFactory())
	xt.GlobalRegistry().RegisterFactory(xt_random.NewFactory())
	xt.GlobalRegistry().RegisterFactory(xt_weighted.NewFactory())
}

func (c *Controller) registerComponents() error {
	c.ctrlConnectHandler = handler_ctrl.NewConnectHandler(c.config.Id, c.network)
	c.mgmtConnectHandler = handler_mgmt.NewConnectHandler(c.config.Id, c.network)

	return nil
}

func (c *Controller) RegisterXctrl(x xctrl.Xctrl) error {
	if err := c.config.Configure(x); err != nil {
		return err
	}
	if x.Enabled() {
		c.xctrls = append(c.xctrls, x)
		if c.config.Trace.Handler != nil {
			for _, decoder := range x.GetTraceDecoders() {
				c.config.Trace.Handler.AddDecoder(decoder)
			}
		}
	}
	return nil
}

func (c *Controller) RegisterXmgmt(x xmgmt.Xmgmt) error {
	if err := c.config.Configure(x); err != nil {
		return err
	}
	if x.Enabled() {
		c.xmgmts = append(c.xmgmts, x)
	}
	return nil
}

func (c *Controller) GetXWebInstance() xweb.Instance {
	return c.xweb
}

func (c *Controller) GetNetwork() *network.Network {
	return c.network
}

func (c *Controller) Identity() identity.Identity {
	return c.config.Id
}
