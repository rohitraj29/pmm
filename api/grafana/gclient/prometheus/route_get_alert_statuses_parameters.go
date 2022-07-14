// Code generated by go-swagger; DO NOT EDIT.

package prometheus

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

// NewRouteGetAlertStatusesParams creates a new RouteGetAlertStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteGetAlertStatusesParams() *RouteGetAlertStatusesParams {
	return &RouteGetAlertStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteGetAlertStatusesParamsWithTimeout creates a new RouteGetAlertStatusesParams object
// with the ability to set a timeout on a request.
func NewRouteGetAlertStatusesParamsWithTimeout(timeout time.Duration) *RouteGetAlertStatusesParams {
	return &RouteGetAlertStatusesParams{
		timeout: timeout,
	}
}

// NewRouteGetAlertStatusesParamsWithContext creates a new RouteGetAlertStatusesParams object
// with the ability to set a context for a request.
func NewRouteGetAlertStatusesParamsWithContext(ctx context.Context) *RouteGetAlertStatusesParams {
	return &RouteGetAlertStatusesParams{
		Context: ctx,
	}
}

// NewRouteGetAlertStatusesParamsWithHTTPClient creates a new RouteGetAlertStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteGetAlertStatusesParamsWithHTTPClient(client *http.Client) *RouteGetAlertStatusesParams {
	return &RouteGetAlertStatusesParams{
		HTTPClient: client,
	}
}

/* RouteGetAlertStatusesParams contains all the parameters to send to the API endpoint
   for the route get alert statuses operation.

   Typically these are written to a http.Request.
*/
type RouteGetAlertStatusesParams struct {
	/* Recipient.

	     Recipient should be "grafana" for requests to be handled by grafana
	and the numeric datasource id for requests to be forwarded to a datasource
	*/
	Recipient string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route get alert statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetAlertStatusesParams) WithDefaults() *RouteGetAlertStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route get alert statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetAlertStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) WithTimeout(timeout time.Duration) *RouteGetAlertStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) WithContext(ctx context.Context) *RouteGetAlertStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) WithHTTPClient(client *http.Client) *RouteGetAlertStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipient adds the recipient to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) WithRecipient(recipient string) *RouteGetAlertStatusesParams {
	o.SetRecipient(recipient)
	return o
}

// SetRecipient adds the recipient to the route get alert statuses params
func (o *RouteGetAlertStatusesParams) SetRecipient(recipient string) {
	o.Recipient = recipient
}

// WriteToRequest writes these params to a swagger request
func (o *RouteGetAlertStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Recipient
	if err := r.SetPathParam("Recipient", o.Recipient); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
