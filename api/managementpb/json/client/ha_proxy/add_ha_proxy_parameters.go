// Code generated by go-swagger; DO NOT EDIT.

package ha_proxy

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

// NewAddHAProxyParams creates a new AddHAProxyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddHAProxyParams() *AddHAProxyParams {
	return &AddHAProxyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddHAProxyParamsWithTimeout creates a new AddHAProxyParams object
// with the ability to set a timeout on a request.
func NewAddHAProxyParamsWithTimeout(timeout time.Duration) *AddHAProxyParams {
	return &AddHAProxyParams{
		timeout: timeout,
	}
}

// NewAddHAProxyParamsWithContext creates a new AddHAProxyParams object
// with the ability to set a context for a request.
func NewAddHAProxyParamsWithContext(ctx context.Context) *AddHAProxyParams {
	return &AddHAProxyParams{
		Context: ctx,
	}
}

// NewAddHAProxyParamsWithHTTPClient creates a new AddHAProxyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddHAProxyParamsWithHTTPClient(client *http.Client) *AddHAProxyParams {
	return &AddHAProxyParams{
		HTTPClient: client,
	}
}

/*
AddHAProxyParams contains all the parameters to send to the API endpoint

	for the add HA proxy operation.

	Typically these are written to a http.Request.
*/
type AddHAProxyParams struct {

	// Body.
	Body AddHAProxyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add HA proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddHAProxyParams) WithDefaults() *AddHAProxyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add HA proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddHAProxyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add HA proxy params
func (o *AddHAProxyParams) WithTimeout(timeout time.Duration) *AddHAProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add HA proxy params
func (o *AddHAProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add HA proxy params
func (o *AddHAProxyParams) WithContext(ctx context.Context) *AddHAProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add HA proxy params
func (o *AddHAProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add HA proxy params
func (o *AddHAProxyParams) WithHTTPClient(client *http.Client) *AddHAProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add HA proxy params
func (o *AddHAProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add HA proxy params
func (o *AddHAProxyParams) WithBody(body AddHAProxyBody) *AddHAProxyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add HA proxy params
func (o *AddHAProxyParams) SetBody(body AddHAProxyBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddHAProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
