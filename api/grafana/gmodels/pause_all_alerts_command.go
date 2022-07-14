// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PauseAllAlertsCommand pause all alerts command
//
// swagger:model PauseAllAlertsCommand
type PauseAllAlertsCommand struct {
	// paused
	Paused bool `json:"paused,omitempty"`
}

// Validate validates this pause all alerts command
func (m *PauseAllAlertsCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pause all alerts command based on context it is used
func (m *PauseAllAlertsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PauseAllAlertsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PauseAllAlertsCommand) UnmarshalBinary(b []byte) error {
	var res PauseAllAlertsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
