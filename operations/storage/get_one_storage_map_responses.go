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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// GetOneStorageMapOKCode is the HTTP code returned for type GetOneStorageMapOK
const GetOneStorageMapOKCode int = 200

/*GetOneStorageMapOK Successful operation

swagger:response getOneStorageMapOK
*/
type GetOneStorageMapOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetOneStorageMapOK creates GetOneStorageMapOK with default headers values
func NewGetOneStorageMapOK() *GetOneStorageMapOK {

	return &GetOneStorageMapOK{}
}

// WithPayload adds the payload to the get one storage map o k response
func (o *GetOneStorageMapOK) WithPayload(payload io.ReadCloser) *GetOneStorageMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get one storage map o k response
func (o *GetOneStorageMapOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOneStorageMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetOneStorageMapNotFoundCode is the HTTP code returned for type GetOneStorageMapNotFound
const GetOneStorageMapNotFoundCode int = 404

/*GetOneStorageMapNotFound The specified resource was not found

swagger:response getOneStorageMapNotFound
*/
type GetOneStorageMapNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOneStorageMapNotFound creates GetOneStorageMapNotFound with default headers values
func NewGetOneStorageMapNotFound() *GetOneStorageMapNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetOneStorageMapNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get one storage map not found response
func (o *GetOneStorageMapNotFound) WithConfigurationVersion(configurationVersion int64) *GetOneStorageMapNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get one storage map not found response
func (o *GetOneStorageMapNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get one storage map not found response
func (o *GetOneStorageMapNotFound) WithPayload(payload *models.Error) *GetOneStorageMapNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get one storage map not found response
func (o *GetOneStorageMapNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOneStorageMapNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetOneStorageMapDefault General Error

swagger:response getOneStorageMapDefault
*/
type GetOneStorageMapDefault struct {
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

// NewGetOneStorageMapDefault creates GetOneStorageMapDefault with default headers values
func NewGetOneStorageMapDefault(code int) *GetOneStorageMapDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetOneStorageMapDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get one storage map default response
func (o *GetOneStorageMapDefault) WithStatusCode(code int) *GetOneStorageMapDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get one storage map default response
func (o *GetOneStorageMapDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get one storage map default response
func (o *GetOneStorageMapDefault) WithConfigurationVersion(configurationVersion int64) *GetOneStorageMapDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get one storage map default response
func (o *GetOneStorageMapDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get one storage map default response
func (o *GetOneStorageMapDefault) WithPayload(payload *models.Error) *GetOneStorageMapDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get one storage map default response
func (o *GetOneStorageMapDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOneStorageMapDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
