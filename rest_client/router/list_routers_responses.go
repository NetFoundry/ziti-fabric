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

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/fabric/rest_model"
)

// ListRoutersReader is a Reader for the ListRouters structure.
type ListRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRoutersOK creates a ListRoutersOK with default headers values
func NewListRoutersOK() *ListRoutersOK {
	return &ListRoutersOK{}
}

/*
ListRoutersOK describes a response with status code 200, with default header values.

A list of routers
*/
type ListRoutersOK struct {
	Payload *rest_model.ListRoutersEnvelope
}

// IsSuccess returns true when this list routers o k response has a 2xx status code
func (o *ListRoutersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list routers o k response has a 3xx status code
func (o *ListRoutersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list routers o k response has a 4xx status code
func (o *ListRoutersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list routers o k response has a 5xx status code
func (o *ListRoutersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list routers o k response a status code equal to that given
func (o *ListRoutersOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListRoutersOK) Error() string {
	return fmt.Sprintf("[GET /routers][%d] listRoutersOK  %+v", 200, o.Payload)
}

func (o *ListRoutersOK) String() string {
	return fmt.Sprintf("[GET /routers][%d] listRoutersOK  %+v", 200, o.Payload)
}

func (o *ListRoutersOK) GetPayload() *rest_model.ListRoutersEnvelope {
	return o.Payload
}

func (o *ListRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListRoutersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRoutersUnauthorized creates a ListRoutersUnauthorized with default headers values
func NewListRoutersUnauthorized() *ListRoutersUnauthorized {
	return &ListRoutersUnauthorized{}
}

/*
ListRoutersUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListRoutersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this list routers unauthorized response has a 2xx status code
func (o *ListRoutersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list routers unauthorized response has a 3xx status code
func (o *ListRoutersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list routers unauthorized response has a 4xx status code
func (o *ListRoutersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list routers unauthorized response has a 5xx status code
func (o *ListRoutersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list routers unauthorized response a status code equal to that given
func (o *ListRoutersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListRoutersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /routers][%d] listRoutersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRoutersUnauthorized) String() string {
	return fmt.Sprintf("[GET /routers][%d] listRoutersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRoutersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
