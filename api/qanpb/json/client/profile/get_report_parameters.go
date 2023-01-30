// Code generated by go-swagger; DO NOT EDIT.

package profile

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

// NewGetReportParams creates a new GetReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReportParams() *GetReportParams {
	return &GetReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportParamsWithTimeout creates a new GetReportParams object
// with the ability to set a timeout on a request.
func NewGetReportParamsWithTimeout(timeout time.Duration) *GetReportParams {
	return &GetReportParams{
		timeout: timeout,
	}
}

// NewGetReportParamsWithContext creates a new GetReportParams object
// with the ability to set a context for a request.
func NewGetReportParamsWithContext(ctx context.Context) *GetReportParams {
	return &GetReportParams{
		Context: ctx,
	}
}

// NewGetReportParamsWithHTTPClient creates a new GetReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReportParamsWithHTTPClient(client *http.Client) *GetReportParams {
	return &GetReportParams{
		HTTPClient: client,
	}
}

/*
GetReportParams contains all the parameters to send to the API endpoint

	for the get report operation.

	Typically these are written to a http.Request.
*/
type GetReportParams struct {

	/* Body.

	   ReportRequest defines filtering of metrics report for db server or other dimentions.
	*/
	Body GetReportBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportParams) WithDefaults() *GetReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get report params
func (o *GetReportParams) WithTimeout(timeout time.Duration) *GetReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get report params
func (o *GetReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get report params
func (o *GetReportParams) WithContext(ctx context.Context) *GetReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get report params
func (o *GetReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get report params
func (o *GetReportParams) WithHTTPClient(client *http.Client) *GetReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get report params
func (o *GetReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get report params
func (o *GetReportParams) WithBody(body GetReportBody) *GetReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get report params
func (o *GetReportParams) SetBody(body GetReportBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
