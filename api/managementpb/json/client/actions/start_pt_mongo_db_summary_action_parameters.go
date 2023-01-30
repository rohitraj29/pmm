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

// NewStartPTMongoDBSummaryActionParams creates a new StartPTMongoDBSummaryActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartPTMongoDBSummaryActionParams() *StartPTMongoDBSummaryActionParams {
	return &StartPTMongoDBSummaryActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartPTMongoDBSummaryActionParamsWithTimeout creates a new StartPTMongoDBSummaryActionParams object
// with the ability to set a timeout on a request.
func NewStartPTMongoDBSummaryActionParamsWithTimeout(timeout time.Duration) *StartPTMongoDBSummaryActionParams {
	return &StartPTMongoDBSummaryActionParams{
		timeout: timeout,
	}
}

// NewStartPTMongoDBSummaryActionParamsWithContext creates a new StartPTMongoDBSummaryActionParams object
// with the ability to set a context for a request.
func NewStartPTMongoDBSummaryActionParamsWithContext(ctx context.Context) *StartPTMongoDBSummaryActionParams {
	return &StartPTMongoDBSummaryActionParams{
		Context: ctx,
	}
}

// NewStartPTMongoDBSummaryActionParamsWithHTTPClient creates a new StartPTMongoDBSummaryActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartPTMongoDBSummaryActionParamsWithHTTPClient(client *http.Client) *StartPTMongoDBSummaryActionParams {
	return &StartPTMongoDBSummaryActionParams{
		HTTPClient: client,
	}
}

/*
StartPTMongoDBSummaryActionParams contains all the parameters to send to the API endpoint

	for the start PT mongo DB summary action operation.

	Typically these are written to a http.Request.
*/
type StartPTMongoDBSummaryActionParams struct {

	// Body.
	Body StartPTMongoDBSummaryActionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start PT mongo DB summary action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartPTMongoDBSummaryActionParams) WithDefaults() *StartPTMongoDBSummaryActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start PT mongo DB summary action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartPTMongoDBSummaryActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) WithTimeout(timeout time.Duration) *StartPTMongoDBSummaryActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) WithContext(ctx context.Context) *StartPTMongoDBSummaryActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) WithHTTPClient(client *http.Client) *StartPTMongoDBSummaryActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) WithBody(body StartPTMongoDBSummaryActionBody) *StartPTMongoDBSummaryActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start PT mongo DB summary action params
func (o *StartPTMongoDBSummaryActionParams) SetBody(body StartPTMongoDBSummaryActionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartPTMongoDBSummaryActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
