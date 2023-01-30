// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// NewDisconnectParams creates a new DisconnectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisconnectParams() *DisconnectParams {
	return &DisconnectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisconnectParamsWithTimeout creates a new DisconnectParams object
// with the ability to set a timeout on a request.
func NewDisconnectParamsWithTimeout(timeout time.Duration) *DisconnectParams {
	return &DisconnectParams{
		timeout: timeout,
	}
}

// NewDisconnectParamsWithContext creates a new DisconnectParams object
// with the ability to set a context for a request.
func NewDisconnectParamsWithContext(ctx context.Context) *DisconnectParams {
	return &DisconnectParams{
		Context: ctx,
	}
}

// NewDisconnectParamsWithHTTPClient creates a new DisconnectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisconnectParamsWithHTTPClient(client *http.Client) *DisconnectParams {
	return &DisconnectParams{
		HTTPClient: client,
	}
}

/*
DisconnectParams contains all the parameters to send to the API endpoint

	for the disconnect operation.

	Typically these are written to a http.Request.
*/
type DisconnectParams struct {

	// Body.
	Body DisconnectBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disconnect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisconnectParams) WithDefaults() *DisconnectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disconnect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisconnectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disconnect params
func (o *DisconnectParams) WithTimeout(timeout time.Duration) *DisconnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disconnect params
func (o *DisconnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disconnect params
func (o *DisconnectParams) WithContext(ctx context.Context) *DisconnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disconnect params
func (o *DisconnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disconnect params
func (o *DisconnectParams) WithHTTPClient(client *http.Client) *DisconnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disconnect params
func (o *DisconnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the disconnect params
func (o *DisconnectParams) WithBody(body DisconnectBody) *DisconnectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the disconnect params
func (o *DisconnectParams) SetBody(body DisconnectBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DisconnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
