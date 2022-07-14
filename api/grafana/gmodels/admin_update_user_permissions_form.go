// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminUpdateUserPermissionsForm admin update user permissions form
//
// swagger:model AdminUpdateUserPermissionsForm
type AdminUpdateUserPermissionsForm struct {
	// is grafana admin
	IsGrafanaAdmin bool `json:"isGrafanaAdmin,omitempty"`
}

// Validate validates this admin update user permissions form
func (m *AdminUpdateUserPermissionsForm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin update user permissions form based on context it is used
func (m *AdminUpdateUserPermissionsForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminUpdateUserPermissionsForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminUpdateUserPermissionsForm) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserPermissionsForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
