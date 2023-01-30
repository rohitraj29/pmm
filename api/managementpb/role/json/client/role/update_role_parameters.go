// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewUpdateRoleParams creates a new UpdateRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRoleParams() *UpdateRoleParams {
	return &UpdateRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoleParamsWithTimeout creates a new UpdateRoleParams object
// with the ability to set a timeout on a request.
func NewUpdateRoleParamsWithTimeout(timeout time.Duration) *UpdateRoleParams {
	return &UpdateRoleParams{
		timeout: timeout,
	}
}

// NewUpdateRoleParamsWithContext creates a new UpdateRoleParams object
// with the ability to set a context for a request.
func NewUpdateRoleParamsWithContext(ctx context.Context) *UpdateRoleParams {
	return &UpdateRoleParams{
		Context: ctx,
	}
}

// NewUpdateRoleParamsWithHTTPClient creates a new UpdateRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRoleParamsWithHTTPClient(client *http.Client) *UpdateRoleParams {
	return &UpdateRoleParams{
		HTTPClient: client,
	}
}

/*
UpdateRoleParams contains all the parameters to send to the API endpoint

	for the update role operation.

	Typically these are written to a http.Request.
*/
type UpdateRoleParams struct {

	// Body.
	Body UpdateRoleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoleParams) WithDefaults() *UpdateRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update role params
func (o *UpdateRoleParams) WithTimeout(timeout time.Duration) *UpdateRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update role params
func (o *UpdateRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update role params
func (o *UpdateRoleParams) WithContext(ctx context.Context) *UpdateRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update role params
func (o *UpdateRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update role params
func (o *UpdateRoleParams) WithHTTPClient(client *http.Client) *UpdateRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update role params
func (o *UpdateRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update role params
func (o *UpdateRoleParams) WithBody(body UpdateRoleBody) *UpdateRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update role params
func (o *UpdateRoleParams) SetBody(body UpdateRoleBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
