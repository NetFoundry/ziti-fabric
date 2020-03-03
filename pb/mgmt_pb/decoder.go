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

package mgmt_pb

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-foundation/channel2"
)

type Decoder struct{}

const DECODER = "mgmt"

func (d Decoder) Decode(msg *channel2.Message) ([]byte, bool) {
	switch msg.ContentType {
	case int32(ContentType_ListServicesRequestType):
		data, err := channel2.NewTraceMessageDecode(DECODER, "List Services Request").MarshalTraceMessageDecode()
		if err != nil {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

		return data, true

	case int32(ContentType_ListServicesResponseType):
		listServices := &ListServicesResponse{}
		if err := proto.Unmarshal(msg.Body, listServices); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "List Services Response")
			meta["services"] = len(listServices.Services)

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_CreateServiceRequestType):
		createService := &CreateServiceRequest{}
		if err := proto.Unmarshal(msg.Body, createService); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Create Service Request")
			meta["service"] = serviceToString(createService.Service)

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_RemoveServiceRequestType):
		removeService := &RemoveServiceRequest{}
		if err := proto.Unmarshal(msg.Body, removeService); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Remove Service Request")
			meta["serviceId"] = removeService.ServiceId

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_GetServiceRequestType):
		getService := &GetServiceRequest{}
		if err := proto.Unmarshal(msg.Body, getService); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Get Service Request")
			meta["serviceId"] = getService.ServiceId

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_GetServiceResponseType):
		getService := &GetServiceResponse{}
		if err := proto.Unmarshal(msg.Body, getService); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Get Service Response")
			meta["service"] = serviceToString(getService.Service)

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_ListRoutersRequestType):
		data, err := channel2.NewTraceMessageDecode(DECODER, "List Routers Request").MarshalTraceMessageDecode()
		if err != nil {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

		return data, true

	case int32(ContentType_ListRoutersResponseType):
		listRouters := &ListRoutersResponse{}
		if err := proto.Unmarshal(msg.Body, listRouters); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "List Routers Response")
			meta["routers"] = len(listRouters.Routers)

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_CreateRouterRequestType):
		createRouter := &CreateRouterRequest{}
		if err := proto.Unmarshal(msg.Body, createRouter); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Create Router Request")
			meta["router"] = routerToString(createRouter.Router)

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_RemoveRouterRequestType):
		removeRouter := &RemoveRouterRequest{}
		if err := proto.Unmarshal(msg.Body, removeRouter); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Remove Router Request")
			meta["routerId"] = removeRouter.RouterId

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_ListLinksRequestType):
		data, err := channel2.NewTraceMessageDecode(DECODER, "List Links Request").MarshalTraceMessageDecode()
		if err != nil {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

		return data, true

	case int32(ContentType_ListLinksResponseType):
		listLinks := &ListLinksResponse{}
		if err := proto.Unmarshal(msg.Body, listLinks); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "List Links Response")
			meta["links"] = len(listLinks.Links)

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_SetLinkCostRequestType):
		setLinkCost := &SetLinkCostRequest{}
		if err := proto.Unmarshal(msg.Body, setLinkCost); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Set Link Cost Request")
			meta["linkId"] = setLinkCost.LinkId
			meta["cost"] = setLinkCost.Cost

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_SetLinkDownRequestType):
		setLinkDown := &SetLinkDownRequest{}
		if err := proto.Unmarshal(msg.Body, setLinkDown); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "Set Link Down Request")
			meta["linkId"] = setLinkDown.LinkId
			meta["down"] = setLinkDown.Down

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

	case int32(ContentType_ListSessionsRequestType):
		data, err := channel2.NewTraceMessageDecode(DECODER, "List Sessions Request").MarshalTraceMessageDecode()
		if err != nil {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}

		return data, true

	case int32(ContentType_ListSessionsResponseType):
		listSessions := &ListSessionsResponse{}
		if err := proto.Unmarshal(msg.Body, listSessions); err == nil {
			meta := channel2.NewTraceMessageDecode(DECODER, "List Sessions Response")
			meta["sessions"] = len(listSessions.Sessions)

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				pfxlog.Logger().Errorf("unexpected error (%s)", err)
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
			return nil, true
		}
	}

	return nil, false
}

func serviceToString(service *Service) string {
	return fmt.Sprintf("{id=[%s]}", service.Id)
}

func routerToString(router *Router) string {
	return fmt.Sprintf("{id=[%s] fingerprint=[%s] listener=[%s] connected=[%t]}", router.Id, router.Fingerprint, router.ListenerAddress, router.Connected)
}

func (circuit *Circuit) CalculateDisplayPath() string {
	if circuit == nil {
		return ""
	}
	out := ""
	for i := 0; i < len(circuit.Path); i++ {
		if i < len(circuit.Links) {
			out += fmt.Sprintf("[r/%s]->{l/%s}->", circuit.Path[i], circuit.Links[i])
		} else {
			out += fmt.Sprintf("[r/%s]\n", circuit.Path[i])
		}
	}
	return out
}
