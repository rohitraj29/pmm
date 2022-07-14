// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NamespaceConfigResponse namespace config response
//
// swagger:model NamespaceConfigResponse
type NamespaceConfigResponse map[string][]GettableRuleGroupConfig

// Validate validates this namespace config response
func (m NamespaceConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := validate.Required(k, "body", m[k]); err != nil {
			return err
		}

		for i := 0; i < len(m[k]); i++ {
			if err := m[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(k + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this namespace config response based on the context it is used
func (m NamespaceConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {
		for i := 0; i < len(m[k]); i++ {
			if err := m[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(k + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
