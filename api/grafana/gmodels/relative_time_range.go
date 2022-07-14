// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RelativeTimeRange RelativeTimeRange is the per query start and end time
// for requests.
//
// swagger:model RelativeTimeRange
type RelativeTimeRange struct {
	// from
	From int64 `json:"from,omitempty"`

	// to
	To int64 `json:"to,omitempty"`
}

// Validate validates this relative time range
func (m *RelativeTimeRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this relative time range based on context it is used
func (m *RelativeTimeRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RelativeTimeRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelativeTimeRange) UnmarshalBinary(b []byte) error {
	var res RelativeTimeRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
