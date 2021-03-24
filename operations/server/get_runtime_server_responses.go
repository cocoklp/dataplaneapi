// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// GetRuntimeServerOKCode is the HTTP code returned for type GetRuntimeServerOK
const GetRuntimeServerOKCode int = 200

/*GetRuntimeServerOK Successful operation

swagger:response getRuntimeServerOK
*/
type GetRuntimeServerOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeServer `json:"body,omitempty"`
}

// NewGetRuntimeServerOK creates GetRuntimeServerOK with default headers values
func NewGetRuntimeServerOK() *GetRuntimeServerOK {

	return &GetRuntimeServerOK{}
}

// WithPayload adds the payload to the get runtime server o k response
func (o *GetRuntimeServerOK) WithPayload(payload *models.RuntimeServer) *GetRuntimeServerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime server o k response
func (o *GetRuntimeServerOK) SetPayload(payload *models.RuntimeServer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeServerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRuntimeServerNotFoundCode is the HTTP code returned for type GetRuntimeServerNotFound
const GetRuntimeServerNotFoundCode int = 404

/*GetRuntimeServerNotFound The specified resource was not found

swagger:response getRuntimeServerNotFound
*/
type GetRuntimeServerNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRuntimeServerNotFound creates GetRuntimeServerNotFound with default headers values
func NewGetRuntimeServerNotFound() *GetRuntimeServerNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetRuntimeServerNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get runtime server not found response
func (o *GetRuntimeServerNotFound) WithConfigurationVersion(configurationVersion int64) *GetRuntimeServerNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get runtime server not found response
func (o *GetRuntimeServerNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get runtime server not found response
func (o *GetRuntimeServerNotFound) WithPayload(payload *models.Error) *GetRuntimeServerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime server not found response
func (o *GetRuntimeServerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeServerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRuntimeServerDefault General Error

swagger:response getRuntimeServerDefault
*/
type GetRuntimeServerDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRuntimeServerDefault creates GetRuntimeServerDefault with default headers values
func NewGetRuntimeServerDefault(code int) *GetRuntimeServerDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetRuntimeServerDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get runtime server default response
func (o *GetRuntimeServerDefault) WithStatusCode(code int) *GetRuntimeServerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get runtime server default response
func (o *GetRuntimeServerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get runtime server default response
func (o *GetRuntimeServerDefault) WithConfigurationVersion(configurationVersion int64) *GetRuntimeServerDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get runtime server default response
func (o *GetRuntimeServerDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get runtime server default response
func (o *GetRuntimeServerDefault) WithPayload(payload *models.Error) *GetRuntimeServerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime server default response
func (o *GetRuntimeServerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeServerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
