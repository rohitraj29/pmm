// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataSourceListItemDTO data source list item d t o
//
// swagger:model DataSourceListItemDTO
type DataSourceListItemDTO struct {
	// basic auth
	BasicAuth bool `json:"basicAuth,omitempty"`

	// database
	Database string `json:"database,omitempty"`

	// Id
	ID int64 `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// type logo Url
	TypeLogoURL string `json:"typeLogoUrl,omitempty"`

	// type name
	TypeName string `json:"typeName,omitempty"`

	// UID
	UID string `json:"uid,omitempty"`

	// Url
	URL string `json:"url,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// access
	Access DsAccess `json:"access,omitempty"`

	// json data
	JSONData JSON `json:"jsonData,omitempty"`
}

// Validate validates this data source list item d t o
func (m *DataSourceListItemDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSourceListItemDTO) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if err := m.Access.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

// ContextValidate validate this data source list item d t o based on the context it is used
func (m *DataSourceListItemDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSourceListItemDTO) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {
	if err := m.Access.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSourceListItemDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSourceListItemDTO) UnmarshalBinary(b []byte) error {
	var res DataSourceListItemDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
