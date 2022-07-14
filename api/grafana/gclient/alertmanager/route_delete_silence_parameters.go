// Code generated by go-swagger; DO NOT EDIT.

package alertmanager

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

// NewRouteDeleteSilenceParams creates a new RouteDeleteSilenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteDeleteSilenceParams() *RouteDeleteSilenceParams {
	return &RouteDeleteSilenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteDeleteSilenceParamsWithTimeout creates a new RouteDeleteSilenceParams object
// with the ability to set a timeout on a request.
func NewRouteDeleteSilenceParamsWithTimeout(timeout time.Duration) *RouteDeleteSilenceParams {
	return &RouteDeleteSilenceParams{
		timeout: timeout,
	}
}

// NewRouteDeleteSilenceParamsWithContext creates a new RouteDeleteSilenceParams object
// with the ability to set a context for a request.
func NewRouteDeleteSilenceParamsWithContext(ctx context.Context) *RouteDeleteSilenceParams {
	return &RouteDeleteSilenceParams{
		Context: ctx,
	}
}

// NewRouteDeleteSilenceParamsWithHTTPClient creates a new RouteDeleteSilenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteDeleteSilenceParamsWithHTTPClient(client *http.Client) *RouteDeleteSilenceParams {
	return &RouteDeleteSilenceParams{
		HTTPClient: client,
	}
}

/* RouteDeleteSilenceParams contains all the parameters to send to the API endpoint
   for the route delete silence operation.

   Typically these are written to a http.Request.
*/
type RouteDeleteSilenceParams struct {
	/* Recipient.

	     Recipient should be "grafana" for requests to be handled by grafana
	and the numeric datasource id for requests to be forwarded to a datasource
	*/
	Recipient string

	// SilenceID.
	SilenceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route delete silence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteDeleteSilenceParams) WithDefaults() *RouteDeleteSilenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route delete silence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteDeleteSilenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route delete silence params
func (o *RouteDeleteSilenceParams) WithTimeout(timeout time.Duration) *RouteDeleteSilenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route delete silence params
func (o *RouteDeleteSilenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route delete silence params
func (o *RouteDeleteSilenceParams) WithContext(ctx context.Context) *RouteDeleteSilenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route delete silence params
func (o *RouteDeleteSilenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route delete silence params
func (o *RouteDeleteSilenceParams) WithHTTPClient(client *http.Client) *RouteDeleteSilenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route delete silence params
func (o *RouteDeleteSilenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipient adds the recipient to the route delete silence params
func (o *RouteDeleteSilenceParams) WithRecipient(recipient string) *RouteDeleteSilenceParams {
	o.SetRecipient(recipient)
	return o
}

// SetRecipient adds the recipient to the route delete silence params
func (o *RouteDeleteSilenceParams) SetRecipient(recipient string) {
	o.Recipient = recipient
}

// WithSilenceID adds the silenceID to the route delete silence params
func (o *RouteDeleteSilenceParams) WithSilenceID(silenceID string) *RouteDeleteSilenceParams {
	o.SetSilenceID(silenceID)
	return o
}

// SetSilenceID adds the silenceId to the route delete silence params
func (o *RouteDeleteSilenceParams) SetSilenceID(silenceID string) {
	o.SilenceID = silenceID
}

// WriteToRequest writes these params to a swagger request
func (o *RouteDeleteSilenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Recipient
	if err := r.SetPathParam("Recipient", o.Recipient); err != nil {
		return err
	}

	// path param SilenceId
	if err := r.SetPathParam("SilenceId", o.SilenceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
