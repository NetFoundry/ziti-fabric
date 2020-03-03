/*
	Copyright 2019 NetFoundry, Inc.

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
	"github.com/golang/protobuf/proto"
	"github.com/netfoundry/ziti-fabric/controller/network"
	"github.com/netfoundry/ziti-fabric/pb/mgmt_pb"
	"github.com/netfoundry/ziti-foundation/channel2"
)

type removeEndpointHandler struct {
	network *network.Network
}

func newRemoveEndpointHandler(network *network.Network) *removeEndpointHandler {
	return &removeEndpointHandler{network: network}
}

func (h *removeEndpointHandler) ContentType() int32 {
	return int32(mgmt_pb.ContentType_RemoveEndpointRequestType)
}

func (h *removeEndpointHandler) HandleReceive(msg *channel2.Message, ch channel2.Channel) {
	rs := &mgmt_pb.RemoveEndpointRequest{}
	if err := proto.Unmarshal(msg.Body, rs); err != nil {
		sendFailure(msg, ch, err.Error())
		return
	}
	if endpoint, found := h.network.GetEndpoint(rs.EndpointId); found {
		if err := h.network.RemoveEndpoint(endpoint.Id); err != nil {
			sendFailure(msg, ch, err.Error())
		} else {
			sendSuccess(msg, ch, "")
		}
	} else {
		sendFailure(msg, ch, "no such endpoint")
	}
}
