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

// NewToggleCheckAlertParams creates a new ToggleCheckAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewToggleCheckAlertParams() *ToggleCheckAlertParams {
	return &ToggleCheckAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewToggleCheckAlertParamsWithTimeout creates a new ToggleCheckAlertParams object
// with the ability to set a timeout on a request.
func NewToggleCheckAlertParamsWithTimeout(timeout time.Duration) *ToggleCheckAlertParams {
	return &ToggleCheckAlertParams{
		timeout: timeout,
	}
}

// NewToggleCheckAlertParamsWithContext creates a new ToggleCheckAlertParams object
// with the ability to set a context for a request.
func NewToggleCheckAlertParamsWithContext(ctx context.Context) *ToggleCheckAlertParams {
	return &ToggleCheckAlertParams{
		Context: ctx,
	}
}

// NewToggleCheckAlertParamsWithHTTPClient creates a new ToggleCheckAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewToggleCheckAlertParamsWithHTTPClient(client *http.Client) *ToggleCheckAlertParams {
	return &ToggleCheckAlertParams{
		HTTPClient: client,
	}
}

/*
ToggleCheckAlertParams contains all the parameters to send to the API endpoint

	for the toggle check alert operation.

	Typically these are written to a http.Request.
*/
type ToggleCheckAlertParams struct {

	// Body.
	Body ToggleCheckAlertBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the toggle check alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleCheckAlertParams) WithDefaults() *ToggleCheckAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the toggle check alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleCheckAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the toggle check alert params
func (o *ToggleCheckAlertParams) WithTimeout(timeout time.Duration) *ToggleCheckAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the toggle check alert params
func (o *ToggleCheckAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the toggle check alert params
func (o *ToggleCheckAlertParams) WithContext(ctx context.Context) *ToggleCheckAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the toggle check alert params
func (o *ToggleCheckAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the toggle check alert params
func (o *ToggleCheckAlertParams) WithHTTPClient(client *http.Client) *ToggleCheckAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the toggle check alert params
func (o *ToggleCheckAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the toggle check alert params
func (o *ToggleCheckAlertParams) WithBody(body ToggleCheckAlertBody) *ToggleCheckAlertParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the toggle check alert params
func (o *ToggleCheckAlertParams) SetBody(body ToggleCheckAlertBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ToggleCheckAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
