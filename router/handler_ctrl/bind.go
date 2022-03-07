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

package handler_ctrl

import (
	"github.com/openziti/channel"
	"github.com/openziti/channel/latency"
	"github.com/openziti/fabric/controller/xctrl"
	"github.com/openziti/fabric/router/forwarder"
	"github.com/openziti/fabric/router/xgress"
	"github.com/openziti/fabric/router/xlink"
	"github.com/openziti/fabric/trace"
	"github.com/openziti/foundation/identity/identity"
)

type bindHandler struct {
	id                 *identity.TokenId
	dialerCfg          map[string]xgress.OptionsData
	xlinkDialers       []xlink.Dialer
	ctrl               xgress.CtrlChannel
	forwarder          *forwarder.Forwarder
	xctrls             []xctrl.Xctrl
	closeNotify        chan struct{}
	ctrlAddressChanger CtrlAddressChanger
	traceHandler       *channel.TraceHandler
	linkRegistry       xlink.Registry
}

func NewBindHandler(id *identity.TokenId,
	dialerCfg map[string]xgress.OptionsData,
	xlinkDialers []xlink.Dialer,
	ctrl xgress.CtrlChannel,
	forwarder *forwarder.Forwarder,
	xctrls []xctrl.Xctrl,
	ctrlAddressChanger CtrlAddressChanger,
	traceHandler *channel.TraceHandler,
	linkRegistry xlink.Registry,
	closeNotify chan struct{}) channel.BindHandler {
	return &bindHandler{
		id:                 id,
		dialerCfg:          dialerCfg,
		xlinkDialers:       xlinkDialers,
		ctrl:               ctrl,
		forwarder:          forwarder,
		xctrls:             xctrls,
		closeNotify:        closeNotify,
		ctrlAddressChanger: ctrlAddressChanger,
		traceHandler:       traceHandler,
		linkRegistry:       linkRegistry,
	}
}

func (self *bindHandler) BindChannel(binding channel.Binding) error {
	linkDialerPool := &handlerPool{
		options:     self.forwarder.Options.LinkDial,
		closeNotify: self.closeNotify,
	}
	binding.AddTypedReceiveHandler(newDialHandler(self.id, self.ctrl, self.xlinkDialers, linkDialerPool, self.linkRegistry))
	binding.AddTypedReceiveHandler(newRouteHandler(self.id, self.ctrl, self.dialerCfg, self.forwarder, self.closeNotify))
	binding.AddTypedReceiveHandler(newValidateTerminatorsHandler(self.ctrl, self.dialerCfg))
	binding.AddTypedReceiveHandler(newUnrouteHandler(self.forwarder))
	binding.AddTypedReceiveHandler(newTraceHandler(self.id, self.forwarder.TraceController()))
	binding.AddTypedReceiveHandler(newInspectHandler(self.id, self.linkRegistry))
	binding.AddTypedReceiveHandler(newSettingsHandler(self.ctrlAddressChanger))
	binding.AddTypedReceiveHandler(newFaultHandler(self.linkRegistry))

	binding.AddPeekHandler(trace.NewChannelPeekHandler(self.id.Token, binding.GetChannel(), self.forwarder.TraceController(), trace.NewChannelSink(binding.GetChannel())))
	latency.AddLatencyProbeResponder(binding)

	if self.traceHandler != nil {
		binding.AddPeekHandler(self.traceHandler)
	}

	for _, x := range self.xctrls {
		if err := binding.Bind(x); err != nil {
			return err
		}
	}

	return nil
}
