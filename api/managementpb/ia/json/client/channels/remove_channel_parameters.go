// Code generated by go-swagger; DO NOT EDIT.

package channels

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

// NewRemoveChannelParams creates a new RemoveChannelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveChannelParams() *RemoveChannelParams {
	return &RemoveChannelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveChannelParamsWithTimeout creates a new RemoveChannelParams object
// with the ability to set a timeout on a request.
func NewRemoveChannelParamsWithTimeout(timeout time.Duration) *RemoveChannelParams {
	return &RemoveChannelParams{
		timeout: timeout,
	}
}

// NewRemoveChannelParamsWithContext creates a new RemoveChannelParams object
// with the ability to set a context for a request.
func NewRemoveChannelParamsWithContext(ctx context.Context) *RemoveChannelParams {
	return &RemoveChannelParams{
		Context: ctx,
	}
}

// NewRemoveChannelParamsWithHTTPClient creates a new RemoveChannelParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveChannelParamsWithHTTPClient(client *http.Client) *RemoveChannelParams {
	return &RemoveChannelParams{
		HTTPClient: client,
	}
}

/*
RemoveChannelParams contains all the parameters to send to the API endpoint

	for the remove channel operation.

	Typically these are written to a http.Request.
*/
type RemoveChannelParams struct {

	// Body.
	Body RemoveChannelBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveChannelParams) WithDefaults() *RemoveChannelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveChannelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove channel params
func (o *RemoveChannelParams) WithTimeout(timeout time.Duration) *RemoveChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove channel params
func (o *RemoveChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove channel params
func (o *RemoveChannelParams) WithContext(ctx context.Context) *RemoveChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove channel params
func (o *RemoveChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove channel params
func (o *RemoveChannelParams) WithHTTPClient(client *http.Client) *RemoveChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove channel params
func (o *RemoveChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove channel params
func (o *RemoveChannelParams) WithBody(body RemoveChannelBody) *RemoveChannelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove channel params
func (o *RemoveChannelParams) SetBody(body RemoveChannelBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
