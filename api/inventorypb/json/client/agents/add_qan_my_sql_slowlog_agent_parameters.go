// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewAddQANMySQLSlowlogAgentParams creates a new AddQANMySQLSlowlogAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddQANMySQLSlowlogAgentParams() *AddQANMySQLSlowlogAgentParams {
	return &AddQANMySQLSlowlogAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddQANMySQLSlowlogAgentParamsWithTimeout creates a new AddQANMySQLSlowlogAgentParams object
// with the ability to set a timeout on a request.
func NewAddQANMySQLSlowlogAgentParamsWithTimeout(timeout time.Duration) *AddQANMySQLSlowlogAgentParams {
	return &AddQANMySQLSlowlogAgentParams{
		timeout: timeout,
	}
}

// NewAddQANMySQLSlowlogAgentParamsWithContext creates a new AddQANMySQLSlowlogAgentParams object
// with the ability to set a context for a request.
func NewAddQANMySQLSlowlogAgentParamsWithContext(ctx context.Context) *AddQANMySQLSlowlogAgentParams {
	return &AddQANMySQLSlowlogAgentParams{
		Context: ctx,
	}
}

// NewAddQANMySQLSlowlogAgentParamsWithHTTPClient creates a new AddQANMySQLSlowlogAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddQANMySQLSlowlogAgentParamsWithHTTPClient(client *http.Client) *AddQANMySQLSlowlogAgentParams {
	return &AddQANMySQLSlowlogAgentParams{
		HTTPClient: client,
	}
}

/* AddQANMySQLSlowlogAgentParams contains all the parameters to send to the API endpoint
   for the add QAN my SQL slowlog agent operation.

   Typically these are written to a http.Request.
*/
type AddQANMySQLSlowlogAgentParams struct {

	// Body.
	Body AddQANMySQLSlowlogAgentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add QAN my SQL slowlog agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddQANMySQLSlowlogAgentParams) WithDefaults() *AddQANMySQLSlowlogAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add QAN my SQL slowlog agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddQANMySQLSlowlogAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) WithTimeout(timeout time.Duration) *AddQANMySQLSlowlogAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) WithContext(ctx context.Context) *AddQANMySQLSlowlogAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) WithHTTPClient(client *http.Client) *AddQANMySQLSlowlogAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) WithBody(body AddQANMySQLSlowlogAgentBody) *AddQANMySQLSlowlogAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add QAN my SQL slowlog agent params
func (o *AddQANMySQLSlowlogAgentParams) SetBody(body AddQANMySQLSlowlogAgentBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddQANMySQLSlowlogAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
