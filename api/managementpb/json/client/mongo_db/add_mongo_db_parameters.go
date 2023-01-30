// Code generated by go-swagger; DO NOT EDIT.

package mongo_db

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

// NewAddMongoDBParams creates a new AddMongoDBParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddMongoDBParams() *AddMongoDBParams {
	return &AddMongoDBParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddMongoDBParamsWithTimeout creates a new AddMongoDBParams object
// with the ability to set a timeout on a request.
func NewAddMongoDBParamsWithTimeout(timeout time.Duration) *AddMongoDBParams {
	return &AddMongoDBParams{
		timeout: timeout,
	}
}

// NewAddMongoDBParamsWithContext creates a new AddMongoDBParams object
// with the ability to set a context for a request.
func NewAddMongoDBParamsWithContext(ctx context.Context) *AddMongoDBParams {
	return &AddMongoDBParams{
		Context: ctx,
	}
}

// NewAddMongoDBParamsWithHTTPClient creates a new AddMongoDBParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddMongoDBParamsWithHTTPClient(client *http.Client) *AddMongoDBParams {
	return &AddMongoDBParams{
		HTTPClient: client,
	}
}

/*
AddMongoDBParams contains all the parameters to send to the API endpoint

	for the add mongo DB operation.

	Typically these are written to a http.Request.
*/
type AddMongoDBParams struct {

	// Body.
	Body AddMongoDBBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add mongo DB params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddMongoDBParams) WithDefaults() *AddMongoDBParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add mongo DB params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddMongoDBParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add mongo DB params
func (o *AddMongoDBParams) WithTimeout(timeout time.Duration) *AddMongoDBParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add mongo DB params
func (o *AddMongoDBParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add mongo DB params
func (o *AddMongoDBParams) WithContext(ctx context.Context) *AddMongoDBParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add mongo DB params
func (o *AddMongoDBParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add mongo DB params
func (o *AddMongoDBParams) WithHTTPClient(client *http.Client) *AddMongoDBParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add mongo DB params
func (o *AddMongoDBParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add mongo DB params
func (o *AddMongoDBParams) WithBody(body AddMongoDBBody) *AddMongoDBParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add mongo DB params
func (o *AddMongoDBParams) SetBody(body AddMongoDBBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddMongoDBParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
