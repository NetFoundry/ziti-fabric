/*
	Copyright 2020 NetFoundry, Inc.

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
	"github.com/netfoundry/ziti-fabric/controller/network"
	"github.com/netfoundry/ziti-fabric/trace"
	"github.com/netfoundry/ziti-fabric/xmgmt"
	"github.com/netfoundry/ziti-foundation/channel2"
)

type BindHandler struct {
	network *network.Network
	xmgmts  []xmgmt.Xmgmt
}

func NewBindHandler(network *network.Network, xmgmts []xmgmt.Xmgmt) *BindHandler {
	return &BindHandler{network: network, xmgmts: xmgmts}
}

func (bindHandler *BindHandler) BindChannel(ch channel2.Channel) error {
	network := bindHandler.network
	ch.AddReceiveHandler(newCreateRouterHandler(network))
	ch.AddReceiveHandler(newCreateServiceHandler(network))
	ch.AddReceiveHandler(newGetServiceHandler(network))
	ch.AddReceiveHandler(newInspectHandler(network))
	ch.AddReceiveHandler(newListLinksHandler(network))
	ch.AddReceiveHandler(newListRoutersHandler(network))
	ch.AddReceiveHandler(newListServicesHandler(network))
	ch.AddReceiveHandler(newListSessionsHandler(network))
	ch.AddReceiveHandler(newRemoveRouterHandler(network))
	ch.AddReceiveHandler(newRemoveServiceHandler(network))
	ch.AddReceiveHandler(newRemoveSessionHandler(network))
	ch.AddReceiveHandler(newSetLinkCostHandler(network))
	ch.AddReceiveHandler(newSetLinkDownHandler(network))

	ch.AddReceiveHandler(newCreateEndpointHandler(network))
	ch.AddReceiveHandler(newRemoveEndpointHandler(network))
	ch.AddReceiveHandler(newGetEndpointHandler(network))
	ch.AddReceiveHandler(newListEndpointsHandler(network))

	streamMetricHandler := newStreamMetricsHandler(network)
	ch.AddReceiveHandler(streamMetricHandler)
	ch.AddCloseHandler(streamMetricHandler)

	streamSessionsHandler := newStreamSessionsHandler(network)
	ch.AddReceiveHandler(streamSessionsHandler)
	ch.AddCloseHandler(streamSessionsHandler)

	streamTracesHandler := newStreamTracesHandler(network)
	ch.AddReceiveHandler(streamTracesHandler)
	ch.AddCloseHandler(streamTracesHandler)

	ch.AddReceiveHandler(newTogglePipeTracesHandler(network))
	ch.AddPeekHandler(trace.NewChannelPeekHandler(network.GetAppId(), ch, network.GetTraceController(), network.GetTraceEventController()))

	xmgmtDone := make(chan struct{})
	for _, x := range bindHandler.xmgmts {
		if err := ch.Bind(x); err != nil {
			return err
		}
		if err := x.Run(ch, xmgmtDone); err != nil {
			return err
		}
	}
	if len(bindHandler.xmgmts) > 0 {
		ch.AddCloseHandler(newXmgmtCloseHandler(xmgmtDone))
	}

	return nil
}
