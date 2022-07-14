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

// VictorOpsConfig VictorOpsConfig configures notifications via VictorOps.
//
// swagger:model VictorOpsConfig
type VictorOpsConfig struct {
	// custom fields
	CustomFields map[string]string `json:"custom_fields,omitempty"`

	// entity display name
	EntityDisplayName string `json:"entity_display_name,omitempty"`

	// message type
	MessageType string `json:"message_type,omitempty"`

	// monitoring tool
	MonitoringTool string `json:"monitoring_tool,omitempty"`

	// routing key
	RoutingKey string `json:"routing_key,omitempty"`

	// state message
	StateMessage string `json:"state_message,omitempty"`

	// v send resolved
	VSendResolved bool `json:"send_resolved,omitempty"`

	// api key
	APIKey Secret `json:"api_key,omitempty"`

	// api key file
	APIKeyFile Secret `json:"api_key_file,omitempty"`

	// api url
	APIURL *URL `json:"api_url,omitempty"`

	// http config
	HTTPConfig *HTTPClientConfig `json:"http_config,omitempty"`
}

// Validate validates this victor ops config
func (m *VictorOpsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIKeyFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VictorOpsConfig) validateAPIKey(formats strfmt.Registry) error {
	if swag.IsZero(m.APIKey) { // not required
		return nil
	}

	if err := m.APIKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("api_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("api_key")
		}
		return err
	}

	return nil
}

func (m *VictorOpsConfig) validateAPIKeyFile(formats strfmt.Registry) error {
	if swag.IsZero(m.APIKeyFile) { // not required
		return nil
	}

	if err := m.APIKeyFile.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("api_key_file")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("api_key_file")
		}
		return err
	}

	return nil
}

func (m *VictorOpsConfig) validateAPIURL(formats strfmt.Registry) error {
	if swag.IsZero(m.APIURL) { // not required
		return nil
	}

	if m.APIURL != nil {
		if err := m.APIURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api_url")
			}
			return err
		}
	}

	return nil
}

func (m *VictorOpsConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConfig) { // not required
		return nil
	}

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this victor ops config based on the context it is used
func (m *VictorOpsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIKeyFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VictorOpsConfig) contextValidateAPIKey(ctx context.Context, formats strfmt.Registry) error {
	if err := m.APIKey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("api_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("api_key")
		}
		return err
	}

	return nil
}

func (m *VictorOpsConfig) contextValidateAPIKeyFile(ctx context.Context, formats strfmt.Registry) error {
	if err := m.APIKeyFile.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("api_key_file")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("api_key_file")
		}
		return err
	}

	return nil
}

func (m *VictorOpsConfig) contextValidateAPIURL(ctx context.Context, formats strfmt.Registry) error {
	if m.APIURL != nil {
		if err := m.APIURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api_url")
			}
			return err
		}
	}

	return nil
}

func (m *VictorOpsConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {
	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VictorOpsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VictorOpsConfig) UnmarshalBinary(b []byte) error {
	var res VictorOpsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
