// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewAddPostgreSQLServiceParams creates a new AddPostgreSQLServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPostgreSQLServiceParams() *AddPostgreSQLServiceParams {
	return &AddPostgreSQLServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPostgreSQLServiceParamsWithTimeout creates a new AddPostgreSQLServiceParams object
// with the ability to set a timeout on a request.
func NewAddPostgreSQLServiceParamsWithTimeout(timeout time.Duration) *AddPostgreSQLServiceParams {
	return &AddPostgreSQLServiceParams{
		timeout: timeout,
	}
}

// NewAddPostgreSQLServiceParamsWithContext creates a new AddPostgreSQLServiceParams object
// with the ability to set a context for a request.
func NewAddPostgreSQLServiceParamsWithContext(ctx context.Context) *AddPostgreSQLServiceParams {
	return &AddPostgreSQLServiceParams{
		Context: ctx,
	}
}

// NewAddPostgreSQLServiceParamsWithHTTPClient creates a new AddPostgreSQLServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPostgreSQLServiceParamsWithHTTPClient(client *http.Client) *AddPostgreSQLServiceParams {
	return &AddPostgreSQLServiceParams{
		HTTPClient: client,
	}
}

/*
AddPostgreSQLServiceParams contains all the parameters to send to the API endpoint

	for the add postgre SQL service operation.

	Typically these are written to a http.Request.
*/
type AddPostgreSQLServiceParams struct {

	// Body.
	Body AddPostgreSQLServiceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add postgre SQL service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPostgreSQLServiceParams) WithDefaults() *AddPostgreSQLServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add postgre SQL service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPostgreSQLServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) WithTimeout(timeout time.Duration) *AddPostgreSQLServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) WithContext(ctx context.Context) *AddPostgreSQLServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) WithHTTPClient(client *http.Client) *AddPostgreSQLServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) WithBody(body AddPostgreSQLServiceBody) *AddPostgreSQLServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add postgre SQL service params
func (o *AddPostgreSQLServiceParams) SetBody(body AddPostgreSQLServiceBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddPostgreSQLServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
