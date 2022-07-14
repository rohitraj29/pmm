// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchAnnotationsCmd patch annotations cmd
//
// swagger:model PatchAnnotationsCmd
type PatchAnnotationsCmd struct {
	// Id
	ID int64 `json:"id,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// text
	Text string `json:"text,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// time end
	TimeEnd int64 `json:"timeEnd,omitempty"`
}

// Validate validates this patch annotations cmd
func (m *PatchAnnotationsCmd) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch annotations cmd based on context it is used
func (m *PatchAnnotationsCmd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchAnnotationsCmd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchAnnotationsCmd) UnmarshalBinary(b []byte) error {
	var res PatchAnnotationsCmd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
