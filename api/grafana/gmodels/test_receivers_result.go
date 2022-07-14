// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestReceiversResult test receivers result
//
// swagger:model TestReceiversResult
type TestReceiversResult struct {
	// notified at
	// Format: date-time
	NotifiedAt strfmt.DateTime `json:"notified_at,omitempty"`

	// receivers
	Receivers []*TestReceiverResult `json:"receivers"`

	// alert
	Alert *TestReceiversConfigAlertParams `json:"alert,omitempty"`
}

// Validate validates this test receivers result
func (m *TestReceiversResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifiedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestReceiversResult) validateNotifiedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.NotifiedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("notified_at", "body", "date-time", m.NotifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestReceiversResult) validateReceivers(formats strfmt.Registry) error {
	if swag.IsZero(m.Receivers) { // not required
		return nil
	}

	for i := 0; i < len(m.Receivers); i++ {
		if swag.IsZero(m.Receivers[i]) { // not required
			continue
		}

		if m.Receivers[i] != nil {
			if err := m.Receivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestReceiversResult) validateAlert(formats strfmt.Registry) error {
	if swag.IsZero(m.Alert) { // not required
		return nil
	}

	if m.Alert != nil {
		if err := m.Alert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alert")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this test receivers result based on the context it is used
func (m *TestReceiversResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReceivers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestReceiversResult) contextValidateReceivers(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(m.Receivers); i++ {
		if m.Receivers[i] != nil {
			if err := m.Receivers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (m *TestReceiversResult) contextValidateAlert(ctx context.Context, formats strfmt.Registry) error {
	if m.Alert != nil {
		if err := m.Alert.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alert")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestReceiversResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestReceiversResult) UnmarshalBinary(b []byte) error {
	var res TestReceiversResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
