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
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/fabric/controller/rest_model"
)

// UpdateRouterOKCode is the HTTP code returned for type UpdateRouterOK
const UpdateRouterOKCode int = 200

/*
UpdateRouterOK The update request was successful and the resource has been altered

swagger:response updateRouterOK
*/
type UpdateRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewUpdateRouterOK creates UpdateRouterOK with default headers values
func NewUpdateRouterOK() *UpdateRouterOK {

	return &UpdateRouterOK{}
}

// WithPayload adds the payload to the update router o k response
func (o *UpdateRouterOK) WithPayload(payload *rest_model.Empty) *UpdateRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update router o k response
func (o *UpdateRouterOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRouterBadRequestCode is the HTTP code returned for type UpdateRouterBadRequest
const UpdateRouterBadRequestCode int = 400

/*
UpdateRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response updateRouterBadRequest
*/
type UpdateRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateRouterBadRequest creates UpdateRouterBadRequest with default headers values
func NewUpdateRouterBadRequest() *UpdateRouterBadRequest {

	return &UpdateRouterBadRequest{}
}

// WithPayload adds the payload to the update router bad request response
func (o *UpdateRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update router bad request response
func (o *UpdateRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRouterUnauthorizedCode is the HTTP code returned for type UpdateRouterUnauthorized
const UpdateRouterUnauthorizedCode int = 401

/*
UpdateRouterUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response updateRouterUnauthorized
*/
type UpdateRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateRouterUnauthorized creates UpdateRouterUnauthorized with default headers values
func NewUpdateRouterUnauthorized() *UpdateRouterUnauthorized {

	return &UpdateRouterUnauthorized{}
}

// WithPayload adds the payload to the update router unauthorized response
func (o *UpdateRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update router unauthorized response
func (o *UpdateRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRouterNotFoundCode is the HTTP code returned for type UpdateRouterNotFound
const UpdateRouterNotFoundCode int = 404

/*
UpdateRouterNotFound The requested resource does not exist

swagger:response updateRouterNotFound
*/
type UpdateRouterNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateRouterNotFound creates UpdateRouterNotFound with default headers values
func NewUpdateRouterNotFound() *UpdateRouterNotFound {

	return &UpdateRouterNotFound{}
}

// WithPayload adds the payload to the update router not found response
func (o *UpdateRouterNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateRouterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update router not found response
func (o *UpdateRouterNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRouterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}