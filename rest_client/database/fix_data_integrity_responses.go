// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/fabric/rest_model"
)

// FixDataIntegrityReader is a Reader for the FixDataIntegrity structure.
type FixDataIntegrityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FixDataIntegrityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewFixDataIntegrityAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFixDataIntegrityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewFixDataIntegrityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFixDataIntegrityAccepted creates a FixDataIntegrityAccepted with default headers values
func NewFixDataIntegrityAccepted() *FixDataIntegrityAccepted {
	return &FixDataIntegrityAccepted{}
}

/*
FixDataIntegrityAccepted describes a response with status code 202, with default header values.

Base empty response
*/
type FixDataIntegrityAccepted struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this fix data integrity accepted response has a 2xx status code
func (o *FixDataIntegrityAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fix data integrity accepted response has a 3xx status code
func (o *FixDataIntegrityAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fix data integrity accepted response has a 4xx status code
func (o *FixDataIntegrityAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this fix data integrity accepted response has a 5xx status code
func (o *FixDataIntegrityAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this fix data integrity accepted response a status code equal to that given
func (o *FixDataIntegrityAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the fix data integrity accepted response
func (o *FixDataIntegrityAccepted) Code() int {
	return 202
}

func (o *FixDataIntegrityAccepted) Error() string {
	return fmt.Sprintf("[POST /database/fix-data-integrity][%d] fixDataIntegrityAccepted  %+v", 202, o.Payload)
}

func (o *FixDataIntegrityAccepted) String() string {
	return fmt.Sprintf("[POST /database/fix-data-integrity][%d] fixDataIntegrityAccepted  %+v", 202, o.Payload)
}

func (o *FixDataIntegrityAccepted) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *FixDataIntegrityAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFixDataIntegrityUnauthorized creates a FixDataIntegrityUnauthorized with default headers values
func NewFixDataIntegrityUnauthorized() *FixDataIntegrityUnauthorized {
	return &FixDataIntegrityUnauthorized{}
}

/*
FixDataIntegrityUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type FixDataIntegrityUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this fix data integrity unauthorized response has a 2xx status code
func (o *FixDataIntegrityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fix data integrity unauthorized response has a 3xx status code
func (o *FixDataIntegrityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fix data integrity unauthorized response has a 4xx status code
func (o *FixDataIntegrityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this fix data integrity unauthorized response has a 5xx status code
func (o *FixDataIntegrityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this fix data integrity unauthorized response a status code equal to that given
func (o *FixDataIntegrityUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the fix data integrity unauthorized response
func (o *FixDataIntegrityUnauthorized) Code() int {
	return 401
}

func (o *FixDataIntegrityUnauthorized) Error() string {
	return fmt.Sprintf("[POST /database/fix-data-integrity][%d] fixDataIntegrityUnauthorized  %+v", 401, o.Payload)
}

func (o *FixDataIntegrityUnauthorized) String() string {
	return fmt.Sprintf("[POST /database/fix-data-integrity][%d] fixDataIntegrityUnauthorized  %+v", 401, o.Payload)
}

func (o *FixDataIntegrityUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *FixDataIntegrityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFixDataIntegrityTooManyRequests creates a FixDataIntegrityTooManyRequests with default headers values
func NewFixDataIntegrityTooManyRequests() *FixDataIntegrityTooManyRequests {
	return &FixDataIntegrityTooManyRequests{}
}

/*
FixDataIntegrityTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type FixDataIntegrityTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this fix data integrity too many requests response has a 2xx status code
func (o *FixDataIntegrityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fix data integrity too many requests response has a 3xx status code
func (o *FixDataIntegrityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fix data integrity too many requests response has a 4xx status code
func (o *FixDataIntegrityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this fix data integrity too many requests response has a 5xx status code
func (o *FixDataIntegrityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this fix data integrity too many requests response a status code equal to that given
func (o *FixDataIntegrityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the fix data integrity too many requests response
func (o *FixDataIntegrityTooManyRequests) Code() int {
	return 429
}

func (o *FixDataIntegrityTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /database/fix-data-integrity][%d] fixDataIntegrityTooManyRequests  %+v", 429, o.Payload)
}

func (o *FixDataIntegrityTooManyRequests) String() string {
	return fmt.Sprintf("[POST /database/fix-data-integrity][%d] fixDataIntegrityTooManyRequests  %+v", 429, o.Payload)
}

func (o *FixDataIntegrityTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *FixDataIntegrityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
