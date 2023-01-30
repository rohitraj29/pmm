// Code generated by go-swagger; DO NOT EDIT.

package security_checks

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

// NewChangeSecurityChecksParams creates a new ChangeSecurityChecksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeSecurityChecksParams() *ChangeSecurityChecksParams {
	return &ChangeSecurityChecksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeSecurityChecksParamsWithTimeout creates a new ChangeSecurityChecksParams object
// with the ability to set a timeout on a request.
func NewChangeSecurityChecksParamsWithTimeout(timeout time.Duration) *ChangeSecurityChecksParams {
	return &ChangeSecurityChecksParams{
		timeout: timeout,
	}
}

// NewChangeSecurityChecksParamsWithContext creates a new ChangeSecurityChecksParams object
// with the ability to set a context for a request.
func NewChangeSecurityChecksParamsWithContext(ctx context.Context) *ChangeSecurityChecksParams {
	return &ChangeSecurityChecksParams{
		Context: ctx,
	}
}

// NewChangeSecurityChecksParamsWithHTTPClient creates a new ChangeSecurityChecksParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeSecurityChecksParamsWithHTTPClient(client *http.Client) *ChangeSecurityChecksParams {
	return &ChangeSecurityChecksParams{
		HTTPClient: client,
	}
}

/*
ChangeSecurityChecksParams contains all the parameters to send to the API endpoint

	for the change security checks operation.

	Typically these are written to a http.Request.
*/
type ChangeSecurityChecksParams struct {

	// Body.
	Body ChangeSecurityChecksBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change security checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeSecurityChecksParams) WithDefaults() *ChangeSecurityChecksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change security checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeSecurityChecksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change security checks params
func (o *ChangeSecurityChecksParams) WithTimeout(timeout time.Duration) *ChangeSecurityChecksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change security checks params
func (o *ChangeSecurityChecksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change security checks params
func (o *ChangeSecurityChecksParams) WithContext(ctx context.Context) *ChangeSecurityChecksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change security checks params
func (o *ChangeSecurityChecksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change security checks params
func (o *ChangeSecurityChecksParams) WithHTTPClient(client *http.Client) *ChangeSecurityChecksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change security checks params
func (o *ChangeSecurityChecksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change security checks params
func (o *ChangeSecurityChecksParams) WithBody(body ChangeSecurityChecksBody) *ChangeSecurityChecksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change security checks params
func (o *ChangeSecurityChecksParams) SetBody(body ChangeSecurityChecksBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeSecurityChecksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
