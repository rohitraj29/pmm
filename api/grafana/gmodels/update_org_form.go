// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateOrgForm update org form
//
// swagger:model UpdateOrgForm
type UpdateOrgForm struct {
	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this update org form
func (m *UpdateOrgForm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update org form based on context it is used
func (m *UpdateOrgForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOrgForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOrgForm) UnmarshalBinary(b []byte) error {
	var res UpdateOrgForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
