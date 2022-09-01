// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewOverviewParams creates a new OverviewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOverviewParams() *OverviewParams {
	return &OverviewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOverviewParamsWithTimeout creates a new OverviewParams object
// with the ability to set a timeout on a request.
func NewOverviewParamsWithTimeout(timeout time.Duration) *OverviewParams {
	return &OverviewParams{
		timeout: timeout,
	}
}

// NewOverviewParamsWithContext creates a new OverviewParams object
// with the ability to set a context for a request.
func NewOverviewParamsWithContext(ctx context.Context) *OverviewParams {
	return &OverviewParams{
		Context: ctx,
	}
}

// NewOverviewParamsWithHTTPClient creates a new OverviewParams object
// with the ability to set a custom HTTPClient for a request.
func NewOverviewParamsWithHTTPClient(client *http.Client) *OverviewParams {
	return &OverviewParams{
		HTTPClient: client,
	}
}

/*
OverviewParams contains all the parameters to send to the API endpoint

	for the overview operation.

	Typically these are written to a http.Request.
*/
type OverviewParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverviewParams) WithDefaults() *OverviewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverviewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the overview params
func (o *OverviewParams) WithTimeout(timeout time.Duration) *OverviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the overview params
func (o *OverviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the overview params
func (o *OverviewParams) WithContext(ctx context.Context) *OverviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the overview params
func (o *OverviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the overview params
func (o *OverviewParams) WithHTTPClient(client *http.Client) *OverviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the overview params
func (o *OverviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OverviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
