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

// InhibitRule InhibitRule defines an inhibition rule that mutes alerts that match the
// target labels if an alert matching the source labels exists.
// Both alerts have to have a set of labels being equal.
//
// swagger:model InhibitRule
type InhibitRule struct {
	// SourceMatch defines a set of labels that have to equal the given
	// value for source alerts. Deprecated. Remove before v1.0 release.
	SourceMatch map[string]string `json:"source_match,omitempty"`

	// TargetMatch defines a set of labels that have to equal the given
	// value for target alerts. Deprecated. Remove before v1.0 release.
	TargetMatch map[string]string `json:"target_match,omitempty"`

	// equal
	Equal LabelNames `json:"equal,omitempty"`

	// source match re
	SourceMatchRe MatchRegexps `json:"source_match_re,omitempty"`

	// source matchers
	SourceMatchers Matchers `json:"source_matchers,omitempty"`

	// target match re
	TargetMatchRe MatchRegexps `json:"target_match_re,omitempty"`

	// target matchers
	TargetMatchers Matchers `json:"target_matchers,omitempty"`
}

// Validate validates this inhibit rule
func (m *InhibitRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEqual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceMatchRe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceMatchers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetMatchRe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetMatchers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InhibitRule) validateEqual(formats strfmt.Registry) error {
	if swag.IsZero(m.Equal) { // not required
		return nil
	}

	if err := m.Equal.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("equal")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("equal")
		}
		return err
	}

	return nil
}

func (m *InhibitRule) validateSourceMatchRe(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceMatchRe) { // not required
		return nil
	}

	if m.SourceMatchRe != nil {
		if err := m.SourceMatchRe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_match_re")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_match_re")
			}
			return err
		}
	}

	return nil
}

func (m *InhibitRule) validateSourceMatchers(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceMatchers) { // not required
		return nil
	}

	if err := m.SourceMatchers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source_matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source_matchers")
		}
		return err
	}

	return nil
}

func (m *InhibitRule) validateTargetMatchRe(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetMatchRe) { // not required
		return nil
	}

	if m.TargetMatchRe != nil {
		if err := m.TargetMatchRe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target_match_re")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target_match_re")
			}
			return err
		}
	}

	return nil
}

func (m *InhibitRule) validateTargetMatchers(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetMatchers) { // not required
		return nil
	}

	if err := m.TargetMatchers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("target_matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("target_matchers")
		}
		return err
	}

	return nil
}

// ContextValidate validate this inhibit rule based on the context it is used
func (m *InhibitRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEqual(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceMatchRe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceMatchers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetMatchRe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetMatchers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InhibitRule) contextValidateEqual(ctx context.Context, formats strfmt.Registry) error {
	if err := m.Equal.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("equal")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("equal")
		}
		return err
	}

	return nil
}

func (m *InhibitRule) contextValidateSourceMatchRe(ctx context.Context, formats strfmt.Registry) error {
	if err := m.SourceMatchRe.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source_match_re")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source_match_re")
		}
		return err
	}

	return nil
}

func (m *InhibitRule) contextValidateSourceMatchers(ctx context.Context, formats strfmt.Registry) error {
	if err := m.SourceMatchers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source_matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source_matchers")
		}
		return err
	}

	return nil
}

func (m *InhibitRule) contextValidateTargetMatchRe(ctx context.Context, formats strfmt.Registry) error {
	if err := m.TargetMatchRe.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("target_match_re")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("target_match_re")
		}
		return err
	}

	return nil
}

func (m *InhibitRule) contextValidateTargetMatchers(ctx context.Context, formats strfmt.Registry) error {
	if err := m.TargetMatchers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("target_matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("target_matchers")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InhibitRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InhibitRule) UnmarshalBinary(b []byte) error {
	var res InhibitRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
