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

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetStorageEndpointsHandlerFunc turns a function with the right signature into a get storage endpoints handler
type GetStorageEndpointsHandlerFunc func(GetStorageEndpointsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStorageEndpointsHandlerFunc) Handle(params GetStorageEndpointsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetStorageEndpointsHandler interface for that can handle valid get storage endpoints params
type GetStorageEndpointsHandler interface {
	Handle(GetStorageEndpointsParams, interface{}) middleware.Responder
}

// NewGetStorageEndpoints creates a new http.Handler for the get storage endpoints operation
func NewGetStorageEndpoints(ctx *middleware.Context, handler GetStorageEndpointsHandler) *GetStorageEndpoints {
	return &GetStorageEndpoints{Context: ctx, Handler: handler}
}

/*GetStorageEndpoints swagger:route GET /services/haproxy/storage Discovery getStorageEndpoints

Return list of HAProxy storage endpoints

Returns a list of endpoints that use HAProxy storage for persistency, e.g. maps, ssl certificates...

*/
type GetStorageEndpoints struct {
	Context *middleware.Context
	Handler GetStorageEndpointsHandler
}

func (o *GetStorageEndpoints) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStorageEndpointsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
