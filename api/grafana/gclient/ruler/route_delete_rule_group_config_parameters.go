// Code generated by go-swagger; DO NOT EDIT.

package ruler

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

// NewRouteDeleteRuleGroupConfigParams creates a new RouteDeleteRuleGroupConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteDeleteRuleGroupConfigParams() *RouteDeleteRuleGroupConfigParams {
	return &RouteDeleteRuleGroupConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteDeleteRuleGroupConfigParamsWithTimeout creates a new RouteDeleteRuleGroupConfigParams object
// with the ability to set a timeout on a request.
func NewRouteDeleteRuleGroupConfigParamsWithTimeout(timeout time.Duration) *RouteDeleteRuleGroupConfigParams {
	return &RouteDeleteRuleGroupConfigParams{
		timeout: timeout,
	}
}

// NewRouteDeleteRuleGroupConfigParamsWithContext creates a new RouteDeleteRuleGroupConfigParams object
// with the ability to set a context for a request.
func NewRouteDeleteRuleGroupConfigParamsWithContext(ctx context.Context) *RouteDeleteRuleGroupConfigParams {
	return &RouteDeleteRuleGroupConfigParams{
		Context: ctx,
	}
}

// NewRouteDeleteRuleGroupConfigParamsWithHTTPClient creates a new RouteDeleteRuleGroupConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteDeleteRuleGroupConfigParamsWithHTTPClient(client *http.Client) *RouteDeleteRuleGroupConfigParams {
	return &RouteDeleteRuleGroupConfigParams{
		HTTPClient: client,
	}
}

/* RouteDeleteRuleGroupConfigParams contains all the parameters to send to the API endpoint
   for the route delete rule group config operation.

   Typically these are written to a http.Request.
*/
type RouteDeleteRuleGroupConfigParams struct {
	// Groupname.
	Groupname string

	// Namespace.
	Namespace string

	/* Recipient.

	     Recipient should be "grafana" for requests to be handled by grafana
	and the numeric datasource id for requests to be forwarded to a datasource
	*/
	Recipient string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route delete rule group config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteDeleteRuleGroupConfigParams) WithDefaults() *RouteDeleteRuleGroupConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route delete rule group config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteDeleteRuleGroupConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) WithTimeout(timeout time.Duration) *RouteDeleteRuleGroupConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) WithContext(ctx context.Context) *RouteDeleteRuleGroupConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) WithHTTPClient(client *http.Client) *RouteDeleteRuleGroupConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupname adds the groupname to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) WithGroupname(groupname string) *RouteDeleteRuleGroupConfigParams {
	o.SetGroupname(groupname)
	return o
}

// SetGroupname adds the groupname to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) SetGroupname(groupname string) {
	o.Groupname = groupname
}

// WithNamespace adds the namespace to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) WithNamespace(namespace string) *RouteDeleteRuleGroupConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRecipient adds the recipient to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) WithRecipient(recipient string) *RouteDeleteRuleGroupConfigParams {
	o.SetRecipient(recipient)
	return o
}

// SetRecipient adds the recipient to the route delete rule group config params
func (o *RouteDeleteRuleGroupConfigParams) SetRecipient(recipient string) {
	o.Recipient = recipient
}

// WriteToRequest writes these params to a swagger request
func (o *RouteDeleteRuleGroupConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Groupname
	if err := r.SetPathParam("Groupname", o.Groupname); err != nil {
		return err
	}

	// path param Namespace
	if err := r.SetPathParam("Namespace", o.Namespace); err != nil {
		return err
	}

	// path param Recipient
	if err := r.SetPathParam("Recipient", o.Recipient); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
