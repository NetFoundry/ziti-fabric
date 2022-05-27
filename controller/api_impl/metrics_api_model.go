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

package api_impl

import (
	"encoding/json"
	"fmt"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/events"
	"github.com/openziti/foundation/metrics/metrics_pb"
	"github.com/pkg/errors"
	"strings"
)

const EntityNameMetrics = "metrics"

func MapInspectResultValueToMetricsModel(inspectResultValue *network.InspectResultValue, format string) (any, error) {
	var result any

	msg := &metrics_pb.MetricsMessage{}
	if err := json.Unmarshal([]byte(inspectResultValue.Value), msg); err == nil {

		var metricEvents []events.MetricsEvent

		adapter := events.NewFilteredMetricsAdapter(nil, nil, events.MetricsHandlerF(func(event *events.MetricsEvent) {
			metricEvents = append(metricEvents, *event)
		}))

		adapter.AcceptMetrics(msg)

		switch format {
		case "json":
			result = metricEvents
		case "prometheus":
			var promMsgs []string

			for _, msg := range metricEvents {
				event := (events.PrometheusMetricsEvent)(msg)
				o, err := event.Marshal()

				if err == nil {
					promMsgs = append(promMsgs, string(o))
				} else {
					promMsgs = append(promMsgs, fmt.Sprint(err))
				}
			}
			result = promMsgs
		default:
			return nil, errors.New(fmt.Sprintf("Unsupported metrics format %s requested", format))
		}
	} else {
		return nil, err
	}
	return result, nil
}

func MapInspectResultToMetricsResult(inspectResult *network.InspectResult, format string) (*string, error) {

	var emit string

	var r []any

	for _, val := range inspectResult.Results {
		s, _ := MapInspectResultValueToMetricsModel(val, format)

		r = append(r, s)
	}

	switch format {
	case "json":
		var js []any
		for _, mg := range r {
			for _, m := range mg.([]events.MetricsEvent) {
				js = append(js, m)
			}
		}
		s, err := json.Marshal(js)
		if err != nil {
			return nil, err
		}
		emit = string(s)
	case "prometheus":
		//emit = strings.Join([]string(r), "")
		var prom string

		for _, m := range r {
			prom += strings.Join(m.([]string), "")
		}
		emit = prom
	}

	return &emit, nil
}
