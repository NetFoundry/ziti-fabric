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
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/fabric/controller/xt"
	"github.com/openziti/fabric/ctrl_msg"
	"github.com/openziti/fabric/pb/ctrl_pb"
	"github.com/openziti/fabric/router/forwarder"
	"github.com/openziti/fabric/router/handler_xgress"
	"github.com/openziti/fabric/router/xgress"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/identity/identity"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type routeHandler struct {
	id        *identity.TokenId
	ctrl      xgress.CtrlChannel
	dialerCfg map[string]xgress.OptionsData
	forwarder *forwarder.Forwarder
	pool      handlerPool
}

func newRouteHandler(id *identity.TokenId, ctrl xgress.CtrlChannel, dialerCfg map[string]xgress.OptionsData, forwarder *forwarder.Forwarder, closeNotify chan struct{}) *routeHandler {
	handler := &routeHandler{
		id:        id,
		ctrl:      ctrl,
		dialerCfg: dialerCfg,
		forwarder: forwarder,
		pool: handlerPool{
			options:     forwarder.Options.XgressDial,
			closeNotify: closeNotify,
		},
	}

	handler.pool.Start()

	return handler
}

func (rh *routeHandler) ContentType() int32 {
	return int32(ctrl_pb.ContentType_RouteType)
}

func (rh *routeHandler) HandleReceive(msg *channel2.Message, ch channel2.Channel) {
	log := pfxlog.ContextLogger(ch.Label())

	route := &ctrl_pb.Route{}
	if err := proto.Unmarshal(msg.Body, route); err == nil {
		logrus.Debugf("attempt [#%d] for [s/%s]", route.Attempt, route.SessionId)

		if route.Egress != nil {
			if rh.forwarder.HasDestination(xgress.Address(route.Egress.Address)) {
				pfxlog.Logger().Warnf("destination exists for [%s]", route.Egress.Address)
				rh.success(msg, int(route.Attempt), ch, route, nil)
				return
			} else {
				rh.connectEgress(msg, int(route.Attempt), ch, route)
				return
			}
		} else {
			rh.success(msg, int(route.Attempt), ch, route, nil)
		}
	} else {
		log.Errorf("error unmarshaling (%s)", err)
	}
}

func (rh *routeHandler) success(msg *channel2.Message, attempt int, ch channel2.Channel, route *ctrl_pb.Route, peerData xt.PeerData) {
	rh.forwarder.Route(route)

	log := pfxlog.ContextLogger(ch.Label())
	response := ctrl_msg.NewRouteResultSuccessMsg(route.SessionId, attempt)
	for k, v := range peerData {
		response.Headers[int32(k)] = v
	}

	response.ReplyTo(msg)

	if err := rh.ctrl.Channel().Send(response); err == nil {
		log.Debugf("handled route for [s/%s]", route.SessionId)
	} else {
		log.Errorf("send response failed for [s/%s] (%s)", route.SessionId, err)
	}
}

func (rh *routeHandler) fail(msg *channel2.Message, attempt int, ch channel2.Channel, route *ctrl_pb.Route, err error) {
	log := pfxlog.ContextLogger(ch.Label()).
		WithField("sessionId", "s/"+route.SessionId).
		WithField("binding", route.Egress.Binding).
		WithField("address", route.Egress.Destination).
		WithField("attempt", route.Attempt)

	log.WithError(err).Errorf("failed to connect egress")

	response := ctrl_msg.NewRouteResultFailedMessage(route.SessionId, attempt, err.Error())
	response.ReplyTo(msg)
	if err := rh.ctrl.Channel().Send(response); err != nil {
		log.Errorf("send failure response failed for [s/%s] (%s)", route.SessionId, err)
	}
}

func (rh *routeHandler) connectEgress(msg *channel2.Message, attempt int, ch channel2.Channel, route *ctrl_pb.Route) {
	rh.pool.Queue(func() {
		log := pfxlog.ContextLogger(ch.Label()).
			WithField("sessionId", "s/"+route.SessionId).
			WithField("binding", route.Egress.Binding).
			WithField("address", route.Egress.Destination).
			WithField("attempt", route.Attempt)
		log.Debug("route request received")
		if factory, err := xgress.GlobalRegistry().Factory(route.Egress.Binding); err == nil {
			if dialer, err := factory.CreateDialer(rh.dialerCfg[route.Egress.Binding]); err == nil {
				sessionId := &identity.TokenId{Token: route.SessionId, Data: route.Egress.PeerData}

				bindHandler := handler_xgress.NewBindHandler(
					handler_xgress.NewReceiveHandler(rh.forwarder),
					handler_xgress.NewCloseHandler(rh.ctrl, rh.forwarder),
					rh.forwarder)

				if rh.forwarder.Options.XgressDialDwellTime > 0 {
					log.Infof("dwelling [%s] on dial", rh.forwarder.Options.XgressDialDwellTime)
					time.Sleep(rh.forwarder.Options.XgressDialDwellTime)
				}

				if peerData, err := dialer.Dial(route.Egress.Destination, sessionId, xgress.Address(route.Egress.Address), bindHandler); err == nil {
					rh.success(msg, attempt, ch, route, peerData)
				} else {
					rh.fail(msg, attempt, ch, route, errors.Wrapf(err, "error creating route for [s/%s]", route.SessionId))
				}
			} else {
				rh.fail(msg, attempt, ch, route, errors.Wrapf(err, "unable to create dialer for [s/%s]", route.SessionId))
			}
		} else {
			rh.fail(msg, attempt, ch, route, errors.Wrapf(err, "error creating route for [s/%s]", route.SessionId))
		}
	})
}
