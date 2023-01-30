// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewAddContainerNodeParams creates a new AddContainerNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddContainerNodeParams() *AddContainerNodeParams {
	return &AddContainerNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddContainerNodeParamsWithTimeout creates a new AddContainerNodeParams object
// with the ability to set a timeout on a request.
func NewAddContainerNodeParamsWithTimeout(timeout time.Duration) *AddContainerNodeParams {
	return &AddContainerNodeParams{
		timeout: timeout,
	}
}

// NewAddContainerNodeParamsWithContext creates a new AddContainerNodeParams object
// with the ability to set a context for a request.
func NewAddContainerNodeParamsWithContext(ctx context.Context) *AddContainerNodeParams {
	return &AddContainerNodeParams{
		Context: ctx,
	}
}

// NewAddContainerNodeParamsWithHTTPClient creates a new AddContainerNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddContainerNodeParamsWithHTTPClient(client *http.Client) *AddContainerNodeParams {
	return &AddContainerNodeParams{
		HTTPClient: client,
	}
}

/*
AddContainerNodeParams contains all the parameters to send to the API endpoint

	for the add container node operation.

	Typically these are written to a http.Request.
*/
type AddContainerNodeParams struct {

	// Body.
	Body AddContainerNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add container node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddContainerNodeParams) WithDefaults() *AddContainerNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add container node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddContainerNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add container node params
func (o *AddContainerNodeParams) WithTimeout(timeout time.Duration) *AddContainerNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add container node params
func (o *AddContainerNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add container node params
func (o *AddContainerNodeParams) WithContext(ctx context.Context) *AddContainerNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add container node params
func (o *AddContainerNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add container node params
func (o *AddContainerNodeParams) WithHTTPClient(client *http.Client) *AddContainerNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add container node params
func (o *AddContainerNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add container node params
func (o *AddContainerNodeParams) WithBody(body AddContainerNodeBody) *AddContainerNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add container node params
func (o *AddContainerNodeParams) SetBody(body AddContainerNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddContainerNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
