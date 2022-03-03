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
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/pb/ctrl_pb"
	"github.com/openziti/channel"
	"strings"
)

type faultHandler struct {
	r       *network.Router
	network *network.Network
}

func newFaultHandler(r *network.Router, network *network.Network) *faultHandler {
	return &faultHandler{r: r, network: network}
}

func (h *faultHandler) ContentType() int32 {
	return int32(ctrl_pb.ContentType_FaultType)
}

func (h *faultHandler) HandleReceive(msg *channel.Message, ch channel.Channel) {
	log := pfxlog.ContextLogger(ch.Label())

	fault := &ctrl_pb.Fault{}
	if err := proto.Unmarshal(msg.Body, fault); err != nil {
		log.WithError(err).Error("failed to unmarshal fault message")
		return
	}

	go h.handleFault(msg, ch, fault)
}

func (h *faultHandler) handleFault(msg *channel.Message, ch channel.Channel, fault *ctrl_pb.Fault) {
	log := pfxlog.ContextLogger(ch.Label())

	switch fault.Subject {
	case ctrl_pb.FaultSubject_LinkFault:
		linkId := fault.Id
		if err := h.network.LinkConnected(linkId, false); err == nil {
			if link, found := h.network.GetLink(linkId); found {
				h.network.LinkChanged(link)
			}
			log.Infof("link fault [l/%s]", linkId)
		}

	case ctrl_pb.FaultSubject_IngressFault:
		if err := h.network.RemoveCircuit(fault.Id, false); err != nil {
			invalidCircuitErr := network.InvalidCircuitError{}
			if errors.As(err, &invalidCircuitErr) {
				log.Debugf("error handling ingress fault (%s)", err)
			} else {
				log.Errorf("error handling ingress fault (%s)", err)
			}
		} else {
			log.Debugf("handled ingress fault for (%s)", fault.Id)
		}

	case ctrl_pb.FaultSubject_EgressFault:
		if err := h.network.RemoveCircuit(fault.Id, false); err != nil {
			invalidCircuitErr := network.InvalidCircuitError{}
			if errors.As(err, &invalidCircuitErr) {
				log.Debugf("error handling egress fault (%s)", err)
			} else {
				log.Errorf("error handling egress fault (%s)", err)
			}
		} else {
			log.Debugf("handled egress fault for (%s)", fault.Id)
		}

	case ctrl_pb.FaultSubject_ForwardFault:
		circuitIds := strings.Split(fault.Id, " ")
		h.network.ReportForwardingFaults(&network.ForwardingFaultReport{R: h.r, CircuitIds: circuitIds})

	default:
		log.Errorf("unexpected subject (%s)", fault.Subject.String())
	}
}
