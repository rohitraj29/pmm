// Code generated by go-swagger; DO NOT EDIT.

package components

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

// NewGetPXCComponentsParams creates a new GetPXCComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPXCComponentsParams() *GetPXCComponentsParams {
	return &GetPXCComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPXCComponentsParamsWithTimeout creates a new GetPXCComponentsParams object
// with the ability to set a timeout on a request.
func NewGetPXCComponentsParamsWithTimeout(timeout time.Duration) *GetPXCComponentsParams {
	return &GetPXCComponentsParams{
		timeout: timeout,
	}
}

// NewGetPXCComponentsParamsWithContext creates a new GetPXCComponentsParams object
// with the ability to set a context for a request.
func NewGetPXCComponentsParamsWithContext(ctx context.Context) *GetPXCComponentsParams {
	return &GetPXCComponentsParams{
		Context: ctx,
	}
}

// NewGetPXCComponentsParamsWithHTTPClient creates a new GetPXCComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPXCComponentsParamsWithHTTPClient(client *http.Client) *GetPXCComponentsParams {
	return &GetPXCComponentsParams{
		HTTPClient: client,
	}
}

/*
GetPXCComponentsParams contains all the parameters to send to the API endpoint

	for the get PXC components operation.

	Typically these are written to a http.Request.
*/
type GetPXCComponentsParams struct {

	// Body.
	Body GetPXCComponentsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get PXC components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPXCComponentsParams) WithDefaults() *GetPXCComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get PXC components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPXCComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get PXC components params
func (o *GetPXCComponentsParams) WithTimeout(timeout time.Duration) *GetPXCComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get PXC components params
func (o *GetPXCComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get PXC components params
func (o *GetPXCComponentsParams) WithContext(ctx context.Context) *GetPXCComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get PXC components params
func (o *GetPXCComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get PXC components params
func (o *GetPXCComponentsParams) WithHTTPClient(client *http.Client) *GetPXCComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get PXC components params
func (o *GetPXCComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get PXC components params
func (o *GetPXCComponentsParams) WithBody(body GetPXCComponentsBody) *GetPXCComponentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get PXC components params
func (o *GetPXCComponentsParams) SetBody(body GetPXCComponentsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetPXCComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
