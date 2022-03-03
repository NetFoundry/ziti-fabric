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

package xgress

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel"
)

type Decoder struct{}

const DECODER = "data"

func (d Decoder) Decode(msg *channel.Message) ([]byte, bool) {
	switch msg.ContentType {
	case int32(ContentTypePayloadType):
		if payload, err := UnmarshallPayload(msg); err == nil {
			return DecodePayload(payload)
		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
		}

	case int32(ContentTypeAcknowledgementType):
		if ack, err := UnmarshallAcknowledgement(msg); err == nil {
			meta := channel.NewTraceMessageDecode(DECODER, "Acknowledgement")
			meta["circuitId"] = ack.CircuitId
			meta["sequence"] = fmt.Sprintf("len(%d)", len(ack.Sequence))
			switch ack.GetOriginator() {
			case Initiator:
				meta["originator"] = "i"
			case Terminator:
				meta["originator"] = "e"
			}

			data, err := meta.MarshalTraceMessageDecode()
			if err != nil {
				return nil, true
			}

			return data, true

		} else {
			pfxlog.Logger().Errorf("unexpected error (%s)", err)
		}
	}

	return nil, false
}

func DecodePayload(payload *Payload) ([]byte, bool) {
	meta := channel.NewTraceMessageDecode(DECODER, "Payload")
	meta["circuitId"] = payload.CircuitId
	meta["sequence"] = payload.Sequence
	switch payload.GetOriginator() {
	case Initiator:
		meta["originator"] = "i"
	case Terminator:
		meta["originator"] = "e"
	}
	if payload.Flags != 0 {
		meta["flags"] = payload.Flags
	}
	meta["length"] = len(payload.Data)

	data, err := meta.MarshalTraceMessageDecode()
	if err != nil {
		return nil, true
	}

	return data, true
}
