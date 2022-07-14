// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TeamGroupMapping team group mapping
//
// swagger:model TeamGroupMapping
type TeamGroupMapping struct {
	// group Id
	GroupID string `json:"groupId,omitempty"`
}

// Validate validates this team group mapping
func (m *TeamGroupMapping) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this team group mapping based on context it is used
func (m *TeamGroupMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamGroupMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamGroupMapping) UnmarshalBinary(b []byte) error {
	var res TeamGroupMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
