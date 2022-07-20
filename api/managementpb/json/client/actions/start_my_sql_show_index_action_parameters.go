// Code generated by go-swagger; DO NOT EDIT.

package actions

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

// NewStartMySQLShowIndexActionParams creates a new StartMySQLShowIndexActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartMySQLShowIndexActionParams() *StartMySQLShowIndexActionParams {
	return &StartMySQLShowIndexActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartMySQLShowIndexActionParamsWithTimeout creates a new StartMySQLShowIndexActionParams object
// with the ability to set a timeout on a request.
func NewStartMySQLShowIndexActionParamsWithTimeout(timeout time.Duration) *StartMySQLShowIndexActionParams {
	return &StartMySQLShowIndexActionParams{
		timeout: timeout,
	}
}

// NewStartMySQLShowIndexActionParamsWithContext creates a new StartMySQLShowIndexActionParams object
// with the ability to set a context for a request.
func NewStartMySQLShowIndexActionParamsWithContext(ctx context.Context) *StartMySQLShowIndexActionParams {
	return &StartMySQLShowIndexActionParams{
		Context: ctx,
	}
}

// NewStartMySQLShowIndexActionParamsWithHTTPClient creates a new StartMySQLShowIndexActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartMySQLShowIndexActionParamsWithHTTPClient(client *http.Client) *StartMySQLShowIndexActionParams {
	return &StartMySQLShowIndexActionParams{
		HTTPClient: client,
	}
}

/* StartMySQLShowIndexActionParams contains all the parameters to send to the API endpoint
   for the start my SQL show index action operation.

   Typically these are written to a http.Request.
*/
type StartMySQLShowIndexActionParams struct {

	// Body.
	Body StartMySQLShowIndexActionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start my SQL show index action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartMySQLShowIndexActionParams) WithDefaults() *StartMySQLShowIndexActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start my SQL show index action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartMySQLShowIndexActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) WithTimeout(timeout time.Duration) *StartMySQLShowIndexActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) WithContext(ctx context.Context) *StartMySQLShowIndexActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) WithHTTPClient(client *http.Client) *StartMySQLShowIndexActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) WithBody(body StartMySQLShowIndexActionBody) *StartMySQLShowIndexActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start my SQL show index action params
func (o *StartMySQLShowIndexActionParams) SetBody(body StartMySQLShowIndexActionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartMySQLShowIndexActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
