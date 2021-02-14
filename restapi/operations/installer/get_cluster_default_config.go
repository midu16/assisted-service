// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClusterDefaultConfigHandlerFunc turns a function with the right signature into a get cluster default config handler
type GetClusterDefaultConfigHandlerFunc func(GetClusterDefaultConfigParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClusterDefaultConfigHandlerFunc) Handle(params GetClusterDefaultConfigParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetClusterDefaultConfigHandler interface for that can handle valid get cluster default config params
type GetClusterDefaultConfigHandler interface {
	Handle(GetClusterDefaultConfigParams, interface{}) middleware.Responder
}

// NewGetClusterDefaultConfig creates a new http.Handler for the get cluster default config operation
func NewGetClusterDefaultConfig(ctx *middleware.Context, handler GetClusterDefaultConfigHandler) *GetClusterDefaultConfig {
	return &GetClusterDefaultConfig{Context: ctx, Handler: handler}
}

/*GetClusterDefaultConfig swagger:route GET /clusters/default-config installer getClusterDefaultConfig

Get the default values for various cluster properties.

*/
type GetClusterDefaultConfig struct {
	Context *middleware.Context
	Handler GetClusterDefaultConfigHandler
}

func (o *GetClusterDefaultConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetClusterDefaultConfigParams()

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
