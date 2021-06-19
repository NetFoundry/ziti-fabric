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
	"bytes"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/fabric/controller/db"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/pb/ctrl_pb"
	"github.com/openziti/fabric/pb/mgmt_pb"
	"github.com/openziti/fabric/router/xgress"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/config"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/metrics"
	"github.com/openziti/foundation/transport"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Config struct {
	Id      *identity.TokenId
	Network *network.Options
	Db      *db.Db
	Trace   struct {
		Handler *channel2.TraceHandler
	}
	Profile struct {
		Memory struct {
			Path     string
			Interval time.Duration
		}
		CPU struct {
			Path string
		}
	}
	Ctrl struct {
		Listener transport.Address
		Options  *channel2.Options
	}
	Mgmt struct {
		Listener transport.Address
		Options  *channel2.Options
	}
	Metrics      *metrics.Config
	HealthChecks struct {
		BoltCheck struct {
			Interval     time.Duration
			Timeout      time.Duration
			InitialDelay time.Duration
		}
	}
	src map[interface{}]interface{}
}

func (config *Config) Configure(sub config.Subconfig) error {
	return sub.LoadConfig(config.src)
}

func LoadConfig(path string) (*Config, error) {
	cfgBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfgmap := make(map[interface{}]interface{})
	if err = yaml.NewDecoder(bytes.NewReader(cfgBytes)).Decode(&cfgmap); err != nil {
		return nil, err
	}
	config.InjectEnv(cfgmap)
	if value, found := cfgmap["v"]; found {
		if value.(int) != 3 {
			panic("config version mismatch: see docs for information on config updates")
		}
	} else {
		panic("no config version: see docs for information on config")
	}

	identityConfig := identity.IdentityConfig{}
	if value, found := cfgmap["identity"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["key"]; found {
				identityConfig.Key = value.(string)
			}
			if value, found := submap["cert"]; found {
				identityConfig.Cert = value.(string)
			}
			if value, found := submap["server_cert"]; found {
				identityConfig.ServerCert = value.(string)
			}
			if value, found := submap["server_key"]; found {
				identityConfig.ServerKey = value.(string)
			}
			if value, found := submap["ca"]; found {
				identityConfig.CA = value.(string)
			}
		}
	}

	cfg := &Config{
		Network: network.DefaultOptions(),
		src:     cfgmap,
	}

	if id, err := identity.LoadIdentity(identityConfig); err != nil {
		return nil, fmt.Errorf("unable to load identity (%s)", err)
	} else {
		cfg.Id = identity.NewIdentity(id)
	}

	if value, found := cfgmap["network"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if options, err := network.LoadOptions(submap); err == nil {
				cfg.Network = options
			} else {
				return nil, fmt.Errorf("invalid 'network' stanza (%s)", err)
			}
		} else {
			pfxlog.Logger().Warn("invalid or empty 'network' stanza")
		}
	}

	dbTrace := false
	if value, found := cfgmap["dbTrace"]; found {
		dbTrace = value.(bool)
	}

	if value, found := cfgmap["db"]; found {
		str, err := db.Open(value.(string), dbTrace)
		if err != nil {
			return nil, err
		}
		cfg.Db = str
	} else {
		panic("config must provide [db]")
	}

	if value, found := cfgmap["trace"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["path"]; found {
				handler, err := channel2.NewTraceHandler(value.(string), cfg.Id)
				if err != nil {
					return nil, err
				}
				handler.AddDecoder(&channel2.Decoder{})
				handler.AddDecoder(&ctrl_pb.Decoder{})
				handler.AddDecoder(&xgress.Decoder{})
				handler.AddDecoder(&mgmt_pb.Decoder{})
				cfg.Trace.Handler = handler
			}
		}
	}

	if value, found := cfgmap["profile"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["memory"]; found {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					if value, found := submap["path"]; found {
						cfg.Profile.Memory.Path = value.(string)
					}
					if value, found := submap["intervalMs"]; found {
						cfg.Profile.Memory.Interval = time.Duration(value.(int)) * time.Millisecond
					} else {
						cfg.Profile.Memory.Interval = 15 * time.Second
					}
				}
			}
			if value, found := submap["cpu"]; found {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					if value, found := submap["path"]; found {
						cfg.Profile.CPU.Path = value.(string)
					}
				}
			}
		}
	}

	if value, found := cfgmap["ctrl"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["listener"]; found {
				listener, err := transport.ParseAddress(value.(string))
				if err != nil {
					return nil, err
				}
				cfg.Ctrl.Listener = listener
			} else {
				panic("config must provide [ctrl/listener]")
			}

			cfg.Ctrl.Options = channel2.DefaultOptions()
			if value, found := submap["options"]; found {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					cfg.Ctrl.Options = channel2.LoadOptions(submap)
					if err := cfg.Ctrl.Options.Validate(); err != nil {
						return nil, fmt.Errorf("error loading channel options for [ctrl/options] (%v)", err)
					}
				}
			}

			if cfg.Trace.Handler != nil {
				cfg.Ctrl.Options.PeekHandlers = append(cfg.Ctrl.Options.PeekHandlers, cfg.Trace.Handler)
			}
		} else {
			panic("config [ctrl] section in unexpected format")
		}
	} else {
		panic("config must provide [ctrl]")
	}

	if value, found := cfgmap["mgmt"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["listener"]; found {
				listener, err := transport.ParseAddress(value.(string))
				if err != nil {
					return nil, err
				}
				cfg.Mgmt.Listener = listener
			} else {
				panic("config must provide [mgmt/listener]")
			}

			cfg.Mgmt.Options = channel2.DefaultOptions()
			if value, found := submap["options"]; found {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					cfg.Mgmt.Options = channel2.LoadOptions(submap)
					if err := cfg.Mgmt.Options.Validate(); err != nil {
						return nil, fmt.Errorf("error loading channel options for [mgmt/options] (%v)", err)
					}
				}
			}
		} else {
			panic("config [mgmt] section in unexpected format")
		}
	} else {
		panic("config must provide [mgmt]")
	}

	if value, found := cfgmap["metrics"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if metricsCfg, err := metrics.LoadConfig(submap); err == nil {
				cfg.Metrics = metricsCfg
			} else {
				return nil, fmt.Errorf("error loading metrics config (%s)", err)
			}
		} else {
			pfxlog.Logger().Warn("invalid or empty [metrics] stanza")
		}
	}

	cfg.HealthChecks.BoltCheck.Interval = 30 * time.Second
	cfg.HealthChecks.BoltCheck.Timeout = 20 * time.Second
	cfg.HealthChecks.BoltCheck.InitialDelay = 30 * time.Second

	configMap := config.NewConfigMap(config.ToStringIntfMap(cfgmap))

	if healthCheckConfig := configMap.Child("healthChecks"); healthCheckConfig != nil {
		if boltCheckConfig := configMap.Child("bolt"); boltCheckConfig != nil {
			if val, found := boltCheckConfig.GetDuration("interval"); found {
				cfg.HealthChecks.BoltCheck.Interval = val
			}
			if val, found := boltCheckConfig.GetDuration("timeout"); found {
				cfg.HealthChecks.BoltCheck.Timeout = val
			}
			if val, found := boltCheckConfig.GetDuration("initialDelay"); found {
				cfg.HealthChecks.BoltCheck.InitialDelay = val
			}
		}
	}

	if configMap.HasError() {
		return nil, err
	}

	configMap = config.NewConfigMap(config.ToStringIntfMap(cfgmap))

	cfg.HealthChecks.BoltCheck.Interval = configMap.Duration("healthChecks.boltCheck.interval", 30*time.Second)
	cfg.HealthChecks.BoltCheck.Timeout = configMap.Duration("healthChecks.boltCheck.timeout", 20*time.Second)
	cfg.HealthChecks.BoltCheck.InitialDelay = configMap.Duration("healthChecks.boltCheck.timeout", 30*time.Second)

	if configMap.HasError() {
		return nil, err
	}

	return cfg, nil
}
