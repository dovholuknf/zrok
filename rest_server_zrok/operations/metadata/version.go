// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VersionHandlerFunc turns a function with the right signature into a version handler
type VersionHandlerFunc func(VersionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VersionHandlerFunc) Handle(params VersionParams) middleware.Responder {
	return fn(params)
}

// VersionHandler interface for that can handle valid version params
type VersionHandler interface {
	Handle(VersionParams) middleware.Responder
}

// NewVersion creates a new http.Handler for the version operation
func NewVersion(ctx *middleware.Context, handler VersionHandler) *Version {
	return &Version{Context: ctx, Handler: handler}
}

/*
	Version swagger:route GET /version metadata version

Version version API
*/
type Version struct {
	Context *middleware.Context
	Handler VersionHandler
}

func (o *Version) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVersionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
