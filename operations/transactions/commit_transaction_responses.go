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

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// CommitTransactionOKCode is the HTTP code returned for type CommitTransactionOK
const CommitTransactionOKCode int = 200

/*CommitTransactionOK Transaction succesfully commited

swagger:response commitTransactionOK
*/
type CommitTransactionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Transaction `json:"body,omitempty"`
}

// NewCommitTransactionOK creates CommitTransactionOK with default headers values
func NewCommitTransactionOK() *CommitTransactionOK {

	return &CommitTransactionOK{}
}

// WithPayload adds the payload to the commit transaction o k response
func (o *CommitTransactionOK) WithPayload(payload *models.Transaction) *CommitTransactionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit transaction o k response
func (o *CommitTransactionOK) SetPayload(payload *models.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitTransactionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CommitTransactionAcceptedCode is the HTTP code returned for type CommitTransactionAccepted
const CommitTransactionAcceptedCode int = 202

/*CommitTransactionAccepted Configuration change accepted and reload requested

swagger:response commitTransactionAccepted
*/
type CommitTransactionAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Transaction `json:"body,omitempty"`
}

// NewCommitTransactionAccepted creates CommitTransactionAccepted with default headers values
func NewCommitTransactionAccepted() *CommitTransactionAccepted {

	return &CommitTransactionAccepted{}
}

// WithReloadID adds the reloadId to the commit transaction accepted response
func (o *CommitTransactionAccepted) WithReloadID(reloadID string) *CommitTransactionAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the commit transaction accepted response
func (o *CommitTransactionAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the commit transaction accepted response
func (o *CommitTransactionAccepted) WithPayload(payload *models.Transaction) *CommitTransactionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit transaction accepted response
func (o *CommitTransactionAccepted) SetPayload(payload *models.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitTransactionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CommitTransactionBadRequestCode is the HTTP code returned for type CommitTransactionBadRequest
const CommitTransactionBadRequestCode int = 400

/*CommitTransactionBadRequest Bad request

swagger:response commitTransactionBadRequest
*/
type CommitTransactionBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCommitTransactionBadRequest creates CommitTransactionBadRequest with default headers values
func NewCommitTransactionBadRequest() *CommitTransactionBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CommitTransactionBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the commit transaction bad request response
func (o *CommitTransactionBadRequest) WithConfigurationVersion(configurationVersion int64) *CommitTransactionBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the commit transaction bad request response
func (o *CommitTransactionBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the commit transaction bad request response
func (o *CommitTransactionBadRequest) WithPayload(payload *models.Error) *CommitTransactionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit transaction bad request response
func (o *CommitTransactionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitTransactionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CommitTransactionNotFoundCode is the HTTP code returned for type CommitTransactionNotFound
const CommitTransactionNotFoundCode int = 404

/*CommitTransactionNotFound The specified resource was not found

swagger:response commitTransactionNotFound
*/
type CommitTransactionNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCommitTransactionNotFound creates CommitTransactionNotFound with default headers values
func NewCommitTransactionNotFound() *CommitTransactionNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CommitTransactionNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the commit transaction not found response
func (o *CommitTransactionNotFound) WithConfigurationVersion(configurationVersion int64) *CommitTransactionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the commit transaction not found response
func (o *CommitTransactionNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the commit transaction not found response
func (o *CommitTransactionNotFound) WithPayload(payload *models.Error) *CommitTransactionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit transaction not found response
func (o *CommitTransactionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CommitTransactionDefault General Error

swagger:response commitTransactionDefault
*/
type CommitTransactionDefault struct {
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

// NewCommitTransactionDefault creates CommitTransactionDefault with default headers values
func NewCommitTransactionDefault(code int) *CommitTransactionDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CommitTransactionDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the commit transaction default response
func (o *CommitTransactionDefault) WithStatusCode(code int) *CommitTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the commit transaction default response
func (o *CommitTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the commit transaction default response
func (o *CommitTransactionDefault) WithConfigurationVersion(configurationVersion int64) *CommitTransactionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the commit transaction default response
func (o *CommitTransactionDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the commit transaction default response
func (o *CommitTransactionDefault) WithPayload(payload *models.Error) *CommitTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit transaction default response
func (o *CommitTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
