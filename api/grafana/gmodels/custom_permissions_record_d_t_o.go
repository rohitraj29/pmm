// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomPermissionsRecordDTO custom permissions record d t o
//
// swagger:model CustomPermissionsRecordDTO
type CustomPermissionsRecordDTO struct {
	// custom permissions
	CustomPermissions string `json:"customPermissions,omitempty"`

	// grantee name
	GranteeName string `json:"granteeName,omitempty"`

	// grantee type
	GranteeType string `json:"granteeType,omitempty"`

	// grantee URL
	GranteeURL string `json:"granteeUrl,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// is folder
	IsFolder bool `json:"isFolder,omitempty"`

	// org ID
	OrgID int64 `json:"orgId,omitempty"`

	// org role
	OrgRole string `json:"orgRole,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// UID
	UID string `json:"uid,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// users count
	UsersCount int64 `json:"usersCount,omitempty"`
}

// Validate validates this custom permissions record d t o
func (m *CustomPermissionsRecordDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custom permissions record d t o based on context it is used
func (m *CustomPermissionsRecordDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomPermissionsRecordDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomPermissionsRecordDTO) UnmarshalBinary(b []byte) error {
	var res CustomPermissionsRecordDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
