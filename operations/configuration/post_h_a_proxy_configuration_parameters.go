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

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPostHAProxyConfigurationParams creates a new PostHAProxyConfigurationParams object
// with the default values initialized.
func NewPostHAProxyConfigurationParams() PostHAProxyConfigurationParams {

	var (
		// initialize parameters with default values

		forceReloadDefault  = bool(false)
		onlyValidateDefault = bool(false)
		skipReloadDefault   = bool(false)
		skipVersionDefault  = bool(false)
	)

	return PostHAProxyConfigurationParams{
		ForceReload: &forceReloadDefault,

		OnlyValidate: &onlyValidateDefault,

		SkipReload: &skipReloadDefault,

		SkipVersion: &skipVersionDefault,
	}
}

// PostHAProxyConfigurationParams contains all the bound params for the post h a proxy configuration operation
// typically these are obtained from a http.Request
//
// swagger:parameters postHAProxyConfiguration
type PostHAProxyConfigurationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*List of Runtime API commands with parameters separated by ';'
	  In: header
	*/
	XRuntimeActions *string
	/*
	  Required: true
	  In: body
	*/
	Data string
	/*If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	  In: query
	  Default: false
	*/
	ForceReload *bool
	/*If set, only validates configuration, without applying it
	  In: query
	  Default: false
	*/
	OnlyValidate *bool
	/*If set, no reload will be initiated and runtime actions from X-Runtime-Actions will be applied
	  In: query
	  Default: false
	*/
	SkipReload *bool
	/*If set, no version check will be done and the pushed config will be enforced
	  In: query
	  Default: false
	*/
	SkipVersion *bool
	/*Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	  In: query
	*/
	Version *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostHAProxyConfigurationParams() beforehand.
func (o *PostHAProxyConfigurationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXRuntimeActions(r.Header[http.CanonicalHeaderKey("X-Runtime-Actions")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body"))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// no validation required on inline body
			o.Data = body
		}
	} else {
		res = append(res, errors.Required("data", "body"))
	}
	qForceReload, qhkForceReload, _ := qs.GetOK("force_reload")
	if err := o.bindForceReload(qForceReload, qhkForceReload, route.Formats); err != nil {
		res = append(res, err)
	}

	qOnlyValidate, qhkOnlyValidate, _ := qs.GetOK("only_validate")
	if err := o.bindOnlyValidate(qOnlyValidate, qhkOnlyValidate, route.Formats); err != nil {
		res = append(res, err)
	}

	qSkipReload, qhkSkipReload, _ := qs.GetOK("skip_reload")
	if err := o.bindSkipReload(qSkipReload, qhkSkipReload, route.Formats); err != nil {
		res = append(res, err)
	}

	qSkipVersion, qhkSkipVersion, _ := qs.GetOK("skip_version")
	if err := o.bindSkipVersion(qSkipVersion, qhkSkipVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRuntimeActions binds and validates parameter XRuntimeActions from header.
func (o *PostHAProxyConfigurationParams) bindXRuntimeActions(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XRuntimeActions = &raw

	return nil
}

// bindForceReload binds and validates parameter ForceReload from query.
func (o *PostHAProxyConfigurationParams) bindForceReload(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPostHAProxyConfigurationParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("force_reload", "query", "bool", raw)
	}
	o.ForceReload = &value

	return nil
}

// bindOnlyValidate binds and validates parameter OnlyValidate from query.
func (o *PostHAProxyConfigurationParams) bindOnlyValidate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPostHAProxyConfigurationParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("only_validate", "query", "bool", raw)
	}
	o.OnlyValidate = &value

	return nil
}

// bindSkipReload binds and validates parameter SkipReload from query.
func (o *PostHAProxyConfigurationParams) bindSkipReload(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPostHAProxyConfigurationParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("skip_reload", "query", "bool", raw)
	}
	o.SkipReload = &value

	return nil
}

// bindSkipVersion binds and validates parameter SkipVersion from query.
func (o *PostHAProxyConfigurationParams) bindSkipVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPostHAProxyConfigurationParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("skip_version", "query", "bool", raw)
	}
	o.SkipVersion = &value

	return nil
}

// bindVersion binds and validates parameter Version from query.
func (o *PostHAProxyConfigurationParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("version", "query", "int64", raw)
	}
	o.Version = &value

	return nil
}
