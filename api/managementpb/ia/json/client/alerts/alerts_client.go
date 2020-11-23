// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAlert(params *DeleteAlertParams) (*DeleteAlertOK, error)

	ListAlerts(params *ListAlertsParams) (*ListAlertsOK, error)

	ToggleAlert(params *ToggleAlertParams) (*ToggleAlertOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAlert deletes alert deletes alert
*/
func (a *Client) DeleteAlert(params *DeleteAlertParams) (*DeleteAlertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAlertParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAlert",
		Method:             "POST",
		PathPattern:        "/v1/management/ia/Alerts/Delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAlertReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAlertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAlertDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAlerts lists alerts returns a list of all alerts
*/
func (a *Client) ListAlerts(params *ListAlertsParams) (*ListAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAlertsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAlerts",
		Method:             "POST",
		PathPattern:        "/v1/management/ia/Alerts/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAlertsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAlertsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ToggleAlert toggles alert allows to switch between silenced and unsilenced states of an alert
*/
func (a *Client) ToggleAlert(params *ToggleAlertParams) (*ToggleAlertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewToggleAlertParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ToggleAlert",
		Method:             "POST",
		PathPattern:        "/v1/management/ia/Alerts/Toggle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ToggleAlertReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ToggleAlertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ToggleAlertDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
