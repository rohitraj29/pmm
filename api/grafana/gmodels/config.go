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
)

// Config Config is the top-level configuration for Alertmanager's config files.
//
// swagger:model Config
type Config struct {
	// inhibit rules
	InhibitRules []*InhibitRule `json:"inhibit_rules"`

	// mute time intervals
	MuteTimeIntervals []*MuteTimeInterval `json:"mute_time_intervals"`

	// templates
	Templates []string `json:"templates"`

	// global
	Global *GlobalConfig `json:"global,omitempty"`

	// route
	Route *Route `json:"route,omitempty"`
}

// Validate validates this config
func (m *Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInhibitRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMuteTimeIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Config) validateInhibitRules(formats strfmt.Registry) error {
	if swag.IsZero(m.InhibitRules) { // not required
		return nil
	}

	for i := 0; i < len(m.InhibitRules); i++ {
		if swag.IsZero(m.InhibitRules[i]) { // not required
			continue
		}

		if m.InhibitRules[i] != nil {
			if err := m.InhibitRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Config) validateMuteTimeIntervals(formats strfmt.Registry) error {
	if swag.IsZero(m.MuteTimeIntervals) { // not required
		return nil
	}

	for i := 0; i < len(m.MuteTimeIntervals); i++ {
		if swag.IsZero(m.MuteTimeIntervals[i]) { // not required
			continue
		}

		if m.MuteTimeIntervals[i] != nil {
			if err := m.MuteTimeIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Config) validateGlobal(formats strfmt.Registry) error {
	if swag.IsZero(m.Global) { // not required
		return nil
	}

	if m.Global != nil {
		if err := m.Global.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global")
			}
			return err
		}
	}

	return nil
}

func (m *Config) validateRoute(formats strfmt.Registry) error {
	if swag.IsZero(m.Route) { // not required
		return nil
	}

	if m.Route != nil {
		if err := m.Route.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config based on the context it is used
func (m *Config) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInhibitRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMuteTimeIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Config) contextValidateInhibitRules(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(m.InhibitRules); i++ {
		if m.InhibitRules[i] != nil {
			if err := m.InhibitRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (m *Config) contextValidateMuteTimeIntervals(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(m.MuteTimeIntervals); i++ {
		if m.MuteTimeIntervals[i] != nil {
			if err := m.MuteTimeIntervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (m *Config) contextValidateGlobal(ctx context.Context, formats strfmt.Registry) error {
	if m.Global != nil {
		if err := m.Global.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global")
			}
			return err
		}
	}

	return nil
}

func (m *Config) contextValidateRoute(ctx context.Context, formats strfmt.Registry) error {
	if m.Route != nil {
		if err := m.Route.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Config) UnmarshalBinary(b []byte) error {
	var res Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
