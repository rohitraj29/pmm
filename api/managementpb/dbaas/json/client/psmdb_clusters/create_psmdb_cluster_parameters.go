// Code generated by go-swagger; DO NOT EDIT.

package psmdb_clusters

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

// NewCreatePSMDBClusterParams creates a new CreatePSMDBClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePSMDBClusterParams() *CreatePSMDBClusterParams {
	return &CreatePSMDBClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePSMDBClusterParamsWithTimeout creates a new CreatePSMDBClusterParams object
// with the ability to set a timeout on a request.
func NewCreatePSMDBClusterParamsWithTimeout(timeout time.Duration) *CreatePSMDBClusterParams {
	return &CreatePSMDBClusterParams{
		timeout: timeout,
	}
}

// NewCreatePSMDBClusterParamsWithContext creates a new CreatePSMDBClusterParams object
// with the ability to set a context for a request.
func NewCreatePSMDBClusterParamsWithContext(ctx context.Context) *CreatePSMDBClusterParams {
	return &CreatePSMDBClusterParams{
		Context: ctx,
	}
}

// NewCreatePSMDBClusterParamsWithHTTPClient creates a new CreatePSMDBClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePSMDBClusterParamsWithHTTPClient(client *http.Client) *CreatePSMDBClusterParams {
	return &CreatePSMDBClusterParams{
		HTTPClient: client,
	}
}

/*
CreatePSMDBClusterParams contains all the parameters to send to the API endpoint

	for the create PSMDB cluster operation.

	Typically these are written to a http.Request.
*/
type CreatePSMDBClusterParams struct {

	// Body.
	Body CreatePSMDBClusterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create PSMDB cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePSMDBClusterParams) WithDefaults() *CreatePSMDBClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create PSMDB cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePSMDBClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) WithTimeout(timeout time.Duration) *CreatePSMDBClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) WithContext(ctx context.Context) *CreatePSMDBClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) WithHTTPClient(client *http.Client) *CreatePSMDBClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) WithBody(body CreatePSMDBClusterBody) *CreatePSMDBClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create PSMDB cluster params
func (o *CreatePSMDBClusterParams) SetBody(body CreatePSMDBClusterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePSMDBClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
