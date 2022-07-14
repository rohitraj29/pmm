// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserSearchHitDTO user search hit d t o
//
// swagger:model UserSearchHitDTO
type UserSearchHitDTO struct {
	// auth labels
	AuthLabels []string `json:"authLabels"`

	// avatar Url
	AvatarURL string `json:"avatarUrl,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// Id
	ID int64 `json:"id,omitempty"`

	// is admin
	IsAdmin bool `json:"isAdmin,omitempty"`

	// is disabled
	IsDisabled bool `json:"isDisabled,omitempty"`

	// last seen at
	// Format: date-time
	LastSeenAt strfmt.DateTime `json:"lastSeenAt,omitempty"`

	// last seen at age
	LastSeenAtAge string `json:"lastSeenAtAge,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this user search hit d t o
func (m *UserSearchHitDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastSeenAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSearchHitDTO) validateLastSeenAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSeenAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSeenAt", "body", "date-time", m.LastSeenAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user search hit d t o based on context it is used
func (m *UserSearchHitDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserSearchHitDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSearchHitDTO) UnmarshalBinary(b []byte) error {
	var res UserSearchHitDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
