// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

// NewCreateRuleParams creates a new CreateRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRuleParams() *CreateRuleParams {
	return &CreateRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRuleParamsWithTimeout creates a new CreateRuleParams object
// with the ability to set a timeout on a request.
func NewCreateRuleParamsWithTimeout(timeout time.Duration) *CreateRuleParams {
	return &CreateRuleParams{
		timeout: timeout,
	}
}

// NewCreateRuleParamsWithContext creates a new CreateRuleParams object
// with the ability to set a context for a request.
func NewCreateRuleParamsWithContext(ctx context.Context) *CreateRuleParams {
	return &CreateRuleParams{
		Context: ctx,
	}
}

// NewCreateRuleParamsWithHTTPClient creates a new CreateRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRuleParamsWithHTTPClient(client *http.Client) *CreateRuleParams {
	return &CreateRuleParams{
		HTTPClient: client,
	}
}

/*
CreateRuleParams contains all the parameters to send to the API endpoint

	for the create rule operation.

	Typically these are written to a http.Request.
*/
type CreateRuleParams struct {

	// Body.
	Body CreateRuleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRuleParams) WithDefaults() *CreateRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create rule params
func (o *CreateRuleParams) WithTimeout(timeout time.Duration) *CreateRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create rule params
func (o *CreateRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create rule params
func (o *CreateRuleParams) WithContext(ctx context.Context) *CreateRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create rule params
func (o *CreateRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create rule params
func (o *CreateRuleParams) WithHTTPClient(client *http.Client) *CreateRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create rule params
func (o *CreateRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create rule params
func (o *CreateRuleParams) WithBody(body CreateRuleBody) *CreateRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create rule params
func (o *CreateRuleParams) SetBody(body CreateRuleBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
