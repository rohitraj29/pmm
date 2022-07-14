// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TagsDTO TagsDTO is the frontend DTO for Tag.
//
// swagger:model TagsDTO
type TagsDTO struct {
	// count
	Count int64 `json:"count,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`
}

// Validate validates this tags d t o
func (m *TagsDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tags d t o based on context it is used
func (m *TagsDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagsDTO) UnmarshalBinary(b []byte) error {
	var res TagsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
