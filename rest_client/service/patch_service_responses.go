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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/fabric/rest_model"
)

// PatchServiceReader is a Reader for the PatchService structure.
type PatchServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchServiceOK creates a PatchServiceOK with default headers values
func NewPatchServiceOK() *PatchServiceOK {
	return &PatchServiceOK{}
}

/*
PatchServiceOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchServiceOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this patch service o k response has a 2xx status code
func (o *PatchServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch service o k response has a 3xx status code
func (o *PatchServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch service o k response has a 4xx status code
func (o *PatchServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch service o k response has a 5xx status code
func (o *PatchServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch service o k response a status code equal to that given
func (o *PatchServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch service o k response
func (o *PatchServiceOK) Code() int {
	return 200
}

func (o *PatchServiceOK) Error() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceOK  %+v", 200, o.Payload)
}

func (o *PatchServiceOK) String() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceOK  %+v", 200, o.Payload)
}

func (o *PatchServiceOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceBadRequest creates a PatchServiceBadRequest with default headers values
func NewPatchServiceBadRequest() *PatchServiceBadRequest {
	return &PatchServiceBadRequest{}
}

/*
PatchServiceBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchServiceBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch service bad request response has a 2xx status code
func (o *PatchServiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch service bad request response has a 3xx status code
func (o *PatchServiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch service bad request response has a 4xx status code
func (o *PatchServiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch service bad request response has a 5xx status code
func (o *PatchServiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch service bad request response a status code equal to that given
func (o *PatchServiceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch service bad request response
func (o *PatchServiceBadRequest) Code() int {
	return 400
}

func (o *PatchServiceBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceBadRequest  %+v", 400, o.Payload)
}

func (o *PatchServiceBadRequest) String() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceBadRequest  %+v", 400, o.Payload)
}

func (o *PatchServiceBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceUnauthorized creates a PatchServiceUnauthorized with default headers values
func NewPatchServiceUnauthorized() *PatchServiceUnauthorized {
	return &PatchServiceUnauthorized{}
}

/*
PatchServiceUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type PatchServiceUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch service unauthorized response has a 2xx status code
func (o *PatchServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch service unauthorized response has a 3xx status code
func (o *PatchServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch service unauthorized response has a 4xx status code
func (o *PatchServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch service unauthorized response has a 5xx status code
func (o *PatchServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch service unauthorized response a status code equal to that given
func (o *PatchServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch service unauthorized response
func (o *PatchServiceUnauthorized) Code() int {
	return 401
}

func (o *PatchServiceUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchServiceUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchServiceUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceNotFound creates a PatchServiceNotFound with default headers values
func NewPatchServiceNotFound() *PatchServiceNotFound {
	return &PatchServiceNotFound{}
}

/*
PatchServiceNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchServiceNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch service not found response has a 2xx status code
func (o *PatchServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch service not found response has a 3xx status code
func (o *PatchServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch service not found response has a 4xx status code
func (o *PatchServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch service not found response has a 5xx status code
func (o *PatchServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch service not found response a status code equal to that given
func (o *PatchServiceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch service not found response
func (o *PatchServiceNotFound) Code() int {
	return 404
}

func (o *PatchServiceNotFound) Error() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceNotFound  %+v", 404, o.Payload)
}

func (o *PatchServiceNotFound) String() string {
	return fmt.Sprintf("[PATCH /services/{id}][%d] patchServiceNotFound  %+v", 404, o.Payload)
}

func (o *PatchServiceNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
