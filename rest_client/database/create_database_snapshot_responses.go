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

// CreateDatabaseSnapshotReader is a Reader for the CreateDatabaseSnapshot structure.
type CreateDatabaseSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDatabaseSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDatabaseSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateDatabaseSnapshotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDatabaseSnapshotTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDatabaseSnapshotOK creates a CreateDatabaseSnapshotOK with default headers values
func NewCreateDatabaseSnapshotOK() *CreateDatabaseSnapshotOK {
	return &CreateDatabaseSnapshotOK{}
}

/*
CreateDatabaseSnapshotOK describes a response with status code 200, with default header values.

Base empty response
*/
type CreateDatabaseSnapshotOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this create database snapshot o k response has a 2xx status code
func (o *CreateDatabaseSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create database snapshot o k response has a 3xx status code
func (o *CreateDatabaseSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create database snapshot o k response has a 4xx status code
func (o *CreateDatabaseSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create database snapshot o k response has a 5xx status code
func (o *CreateDatabaseSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create database snapshot o k response a status code equal to that given
func (o *CreateDatabaseSnapshotOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateDatabaseSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /database][%d] createDatabaseSnapshotOK  %+v", 200, o.Payload)
}

func (o *CreateDatabaseSnapshotOK) String() string {
	return fmt.Sprintf("[POST /database][%d] createDatabaseSnapshotOK  %+v", 200, o.Payload)
}

func (o *CreateDatabaseSnapshotOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *CreateDatabaseSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatabaseSnapshotUnauthorized creates a CreateDatabaseSnapshotUnauthorized with default headers values
func NewCreateDatabaseSnapshotUnauthorized() *CreateDatabaseSnapshotUnauthorized {
	return &CreateDatabaseSnapshotUnauthorized{}
}

/*
CreateDatabaseSnapshotUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type CreateDatabaseSnapshotUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this create database snapshot unauthorized response has a 2xx status code
func (o *CreateDatabaseSnapshotUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create database snapshot unauthorized response has a 3xx status code
func (o *CreateDatabaseSnapshotUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create database snapshot unauthorized response has a 4xx status code
func (o *CreateDatabaseSnapshotUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create database snapshot unauthorized response has a 5xx status code
func (o *CreateDatabaseSnapshotUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create database snapshot unauthorized response a status code equal to that given
func (o *CreateDatabaseSnapshotUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateDatabaseSnapshotUnauthorized) Error() string {
	return fmt.Sprintf("[POST /database][%d] createDatabaseSnapshotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDatabaseSnapshotUnauthorized) String() string {
	return fmt.Sprintf("[POST /database][%d] createDatabaseSnapshotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDatabaseSnapshotUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateDatabaseSnapshotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatabaseSnapshotTooManyRequests creates a CreateDatabaseSnapshotTooManyRequests with default headers values
func NewCreateDatabaseSnapshotTooManyRequests() *CreateDatabaseSnapshotTooManyRequests {
	return &CreateDatabaseSnapshotTooManyRequests{}
}

/*
CreateDatabaseSnapshotTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type CreateDatabaseSnapshotTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this create database snapshot too many requests response has a 2xx status code
func (o *CreateDatabaseSnapshotTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create database snapshot too many requests response has a 3xx status code
func (o *CreateDatabaseSnapshotTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create database snapshot too many requests response has a 4xx status code
func (o *CreateDatabaseSnapshotTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create database snapshot too many requests response has a 5xx status code
func (o *CreateDatabaseSnapshotTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create database snapshot too many requests response a status code equal to that given
func (o *CreateDatabaseSnapshotTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateDatabaseSnapshotTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /database][%d] createDatabaseSnapshotTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateDatabaseSnapshotTooManyRequests) String() string {
	return fmt.Sprintf("[POST /database][%d] createDatabaseSnapshotTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateDatabaseSnapshotTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateDatabaseSnapshotTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
