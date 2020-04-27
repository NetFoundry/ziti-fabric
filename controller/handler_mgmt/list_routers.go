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

package handler_mgmt

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-fabric/controller/handler_common"
	"github.com/netfoundry/ziti-fabric/controller/model"
	"github.com/netfoundry/ziti-fabric/controller/network"
	"github.com/netfoundry/ziti-fabric/pb/mgmt_pb"
	"github.com/netfoundry/ziti-foundation/channel2"
	"reflect"
)

type listRoutersHandler struct {
	network *network.Network
}

func newListRoutersHandler(network *network.Network) *listRoutersHandler {
	return &listRoutersHandler{network: network}
}

func (h *listRoutersHandler) ContentType() int32 {
	return int32(mgmt_pb.ContentType_ListRoutersRequestType)
}

func (h *listRoutersHandler) HandleReceive(msg *channel2.Message, ch channel2.Channel) {
	log := pfxlog.ContextLogger(ch.Label())

	list := &mgmt_pb.ListRoutersRequest{}
	if err := proto.Unmarshal(msg.Body, list); err == nil {
		response := &mgmt_pb.ListRoutersResponse{Routers: make([]*mgmt_pb.Router, 0)}
		if result, err := h.network.Routers.BaseList(list.Query); err == nil {
			log.Infof("got [%d] routers", len(result.Entities))
			for _, entity := range result.Entities {
				router, ok := entity.(*model.Router)
				if !ok {
					errorMsg := fmt.Sprintf("unexpected result in router list of type: %v", reflect.TypeOf(entity))
					handler_common.SendFailure(msg, ch, errorMsg)
					return
				}

				responseR := &mgmt_pb.Router{
					Id:          router.Id,
					Fingerprint: router.Fingerprint,
				}

				if connR := h.network.GetConnectedRouter(router.Id); connR != nil {
					responseR.Connected = true
					responseR.ListenerAddress = connR.AdvertisedListener
				}

				response.Routers = append(response.Routers, responseR)
			}

			if body, err := proto.Marshal(response); err == nil {
				responseMsg := channel2.NewMessage(int32(mgmt_pb.ContentType_ListRoutersResponseType), body)
				responseMsg.ReplyTo(msg)
				if err := ch.Send(responseMsg); err != nil {
					pfxlog.ContextLogger(ch.Label()).Errorf("unexpected error sending response (%s)", err)
				}

			} else {
				handler_common.SendFailure(msg, ch, err.Error())
			}

		} else {
			handler_common.SendFailure(msg, ch, err.Error())
		}
	}
}
