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

package handler_link

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel"
	"github.com/openziti/channel/protobufs"
	"github.com/openziti/fabric/pb/ctrl_pb"
	"github.com/openziti/fabric/router/forwarder"
	"github.com/openziti/fabric/router/xgress"
	"github.com/openziti/fabric/router/xlink"
	"github.com/openziti/foundation/util/concurrenz"
)

type closeHandler struct {
	link          xlink.Xlink
	ctrl          xgress.CtrlChannel
	forwarder     *forwarder.Forwarder
	closeNotify   chan struct{}
	closed        concurrenz.AtomicBoolean
	xlinkRegistry xlink.Registry
}

func newCloseHandler(link xlink.Xlink, ctrl xgress.CtrlChannel, forwarder *forwarder.Forwarder, closeNotify chan struct{}, registry xlink.Registry) *closeHandler {
	return &closeHandler{
		link:          link,
		ctrl:          ctrl,
		forwarder:     forwarder,
		closeNotify:   closeNotify,
		xlinkRegistry: registry,
	}
}

func (self *closeHandler) HandleClose(ch channel.Channel) {
	if self.closed.CompareAndSwap(false, true) {
		log := pfxlog.ContextLogger(ch.Label()).
			WithField("linkId", self.link.Id().Token).
			WithField("routerId", self.link.DestinationId())

		// ensure that both parts of a split link are closed, if one side closes
		go func() {
			_ = self.link.Close()
			// Close can be called from the link registry, so we can't call back into it from the same go-routine
			self.xlinkRegistry.LinkClosed(self.link)
		}()

		log.Info("link closed")

		self.link.HandleCloseNotification(func() {
			fault := &ctrl_pb.Fault{Subject: ctrl_pb.FaultSubject_LinkFault, Id: self.link.Id().Token}
			if err := protobufs.MarshalTyped(fault).Send(self.ctrl.Channel()); err == nil {
				log.Debug("transmitted link fault")
			} else {
				log.WithError(err).Error("unexpected error transmitting link fault")
			}
		})

		self.forwarder.UnregisterLink(self.link)
		close(self.closeNotify)
	}
}
