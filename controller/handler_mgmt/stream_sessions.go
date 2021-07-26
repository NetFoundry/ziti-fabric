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
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/fabric/controller/handler_common"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/events"
	"github.com/openziti/fabric/pb/mgmt_pb"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/identity/identity"
)

type streamSessionsHandler struct {
	network        *network.Network
	streamHandlers []network.SessionEventHandler
}

func newStreamSessionsHandler(network *network.Network) *streamSessionsHandler {
	return &streamSessionsHandler{network: network}
}

func (*streamSessionsHandler) ContentType() int32 {
	return int32(mgmt_pb.ContentType_StreamSessionsRequestType)
}

func (handler *streamSessionsHandler) HandleReceive(msg *channel2.Message, ch channel2.Channel) {
	request := &mgmt_pb.StreamSessionsRequest{}
	if err := proto.Unmarshal(msg.Body, request); err != nil {
		handler_common.SendFailure(msg, ch, err.Error())
		return
	}

	sessionsStreamHandler := &SessionsStreamHandler{ch: ch}
	handler.streamHandlers = append(handler.streamHandlers, sessionsStreamHandler)
	events.AddSessionEventHandler(sessionsStreamHandler)
}

func (handler *streamSessionsHandler) HandleClose(ch channel2.Channel) {
	for _, listener := range handler.streamHandlers {
		events.RemoveSessionEventHandler(listener)
	}
}

type SessionsStreamHandler struct {
	ch channel2.Channel
}

func (handler *SessionsStreamHandler) SessionCreated(sessionId *identity.TokenId, clientId *identity.TokenId, serviceId string, path *network.Path) {
	event := &mgmt_pb.StreamSessionsEvent{
		EventType: mgmt_pb.StreamSessionEventType_SessionCreated,
		SessionId: sessionId.Token,
		ClientId:  clientId.Token,
		ServiceId: serviceId,
		Path:      NewPath(path),
	}
	handler.sendEvent(event)
}

func (handler *SessionsStreamHandler) SessionDeleted(sessionId *identity.TokenId, clientId *identity.TokenId) {
	event := &mgmt_pb.StreamSessionsEvent{
		EventType: mgmt_pb.StreamSessionEventType_SessionDeleted,
		SessionId: sessionId.Token,
		ClientId:  clientId.Token,
	}
	handler.sendEvent(event)
}

func (handler *SessionsStreamHandler) PathUpdated(sessionId *identity.TokenId, path *network.Path) {
	event := &mgmt_pb.StreamSessionsEvent{
		EventType: mgmt_pb.StreamSessionEventType_PathUpdated,
		SessionId: sessionId.Token,
		Path:      NewPath(path),
	}
	handler.sendEvent(event)
}

func (handler *SessionsStreamHandler) sendEvent(event *mgmt_pb.StreamSessionsEvent) {
	body, err := proto.Marshal(event)
	if err != nil {
		pfxlog.Logger().Errorf("unexpected error serializing StreamSessionsEvent (%s)", err)
		return
	}

	responseMsg := channel2.NewMessage(int32(mgmt_pb.ContentType_StreamSessionsEventType), body)
	if err := handler.ch.Send(responseMsg); err != nil {
		pfxlog.Logger().Errorf("unexpected error sending StreamMetricsEvent (%s)", err)
		handler.close()
	}
}

func (handler *SessionsStreamHandler) close() {
	if err := handler.ch.Close(); err != nil {
		pfxlog.Logger().WithError(err).Error("unexpected error closing mgmt channel")
	}
	events.RemoveSessionEventHandler(handler)
}
