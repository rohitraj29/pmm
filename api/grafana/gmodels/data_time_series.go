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

// DataTimeSeries DataTimeSeries -- this structure is deprecated, all new work should use DataFrames from the SDK
//
// swagger:model DataTimeSeries
type DataTimeSeries struct {
	// name
	Name string `json:"name,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// points
	Points DataTimeSeriesPoints `json:"points,omitempty"`
}

// Validate validates this data time series
func (m *DataTimeSeries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTimeSeries) validatePoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Points) { // not required
		return nil
	}

	if err := m.Points.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("points")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("points")
		}
		return err
	}

	return nil
}

// ContextValidate validate this data time series based on the context it is used
func (m *DataTimeSeries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTimeSeries) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {
	if err := m.Points.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("points")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("points")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataTimeSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTimeSeries) UnmarshalBinary(b []byte) error {
	var res DataTimeSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
