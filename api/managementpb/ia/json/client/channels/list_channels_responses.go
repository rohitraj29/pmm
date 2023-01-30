// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListChannelsReader is a Reader for the ListChannels structure.
type ListChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListChannelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListChannelsOK creates a ListChannelsOK with default headers values
func NewListChannelsOK() *ListChannelsOK {
	return &ListChannelsOK{}
}

/*
ListChannelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListChannelsOK struct {
	Payload *ListChannelsOKBody
}

func (o *ListChannelsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/List][%d] listChannelsOk  %+v", 200, o.Payload)
}
func (o *ListChannelsOK) GetPayload() *ListChannelsOKBody {
	return o.Payload
}

func (o *ListChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListChannelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListChannelsDefault creates a ListChannelsDefault with default headers values
func NewListChannelsDefault(code int) *ListChannelsDefault {
	return &ListChannelsDefault{
		_statusCode: code,
	}
}

/*
ListChannelsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListChannelsDefault struct {
	_statusCode int

	Payload *ListChannelsDefaultBody
}

// Code gets the status code for the list channels default response
func (o *ListChannelsDefault) Code() int {
	return o._statusCode
}

func (o *ListChannelsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/List][%d] ListChannels default  %+v", o._statusCode, o.Payload)
}
func (o *ListChannelsDefault) GetPayload() *ListChannelsDefaultBody {
	return o.Payload
}

func (o *ListChannelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListChannelsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListChannelsBody list channels body
swagger:model ListChannelsBody
*/
type ListChannelsBody struct {

	// page params
	PageParams *ListChannelsParamsBodyPageParams `json:"page_params,omitempty"`
}

// Validate validates this list channels body
func (o *ListChannelsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePageParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsBody) validatePageParams(formats strfmt.Registry) error {
	if swag.IsZero(o.PageParams) { // not required
		return nil
	}

	if o.PageParams != nil {
		if err := o.PageParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "page_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "page_params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list channels body based on the context it is used
func (o *ListChannelsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePageParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsBody) contextValidatePageParams(ctx context.Context, formats strfmt.Registry) error {

	if o.PageParams != nil {
		if err := o.PageParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "page_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "page_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsBody) UnmarshalBinary(b []byte) error {
	var res ListChannelsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsDefaultBody list channels default body
swagger:model ListChannelsDefaultBody
*/
type ListChannelsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListChannelsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list channels default body
func (o *ListChannelsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListChannels default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListChannels default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list channels default body based on the context it is used
func (o *ListChannelsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListChannels default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListChannels default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListChannelsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsDefaultBodyDetailsItems0 list channels default body details items0
swagger:model ListChannelsDefaultBodyDetailsItems0
*/
type ListChannelsDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list channels default body details items0
func (o *ListChannelsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels default body details items0 based on context it is used
func (o *ListChannelsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListChannelsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBody list channels OK body
swagger:model ListChannelsOKBody
*/
type ListChannelsOKBody struct {

	// channels
	Channels []*ListChannelsOKBodyChannelsItems0 `json:"channels"`

	// totals
	Totals *ListChannelsOKBodyTotals `json:"totals,omitempty"`
}

// Validate validates this list channels OK body
func (o *ListChannelsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBody) validateChannels(formats strfmt.Registry) error {
	if swag.IsZero(o.Channels) { // not required
		return nil
	}

	for i := 0; i < len(o.Channels); i++ {
		if swag.IsZero(o.Channels[i]) { // not required
			continue
		}

		if o.Channels[i] != nil {
			if err := o.Channels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listChannelsOk" + "." + "channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listChannelsOk" + "." + "channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListChannelsOKBody) validateTotals(formats strfmt.Registry) error {
	if swag.IsZero(o.Totals) { // not required
		return nil
	}

	if o.Totals != nil {
		if err := o.Totals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listChannelsOk" + "." + "totals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listChannelsOk" + "." + "totals")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list channels OK body based on the context it is used
func (o *ListChannelsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTotals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBody) contextValidateChannels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Channels); i++ {

		if o.Channels[i] != nil {
			if err := o.Channels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listChannelsOk" + "." + "channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listChannelsOk" + "." + "channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListChannelsOKBody) contextValidateTotals(ctx context.Context, formats strfmt.Registry) error {

	if o.Totals != nil {
		if err := o.Totals.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listChannelsOk" + "." + "totals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listChannelsOk" + "." + "totals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBody) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0 Channel represents a single Notification Channel.
//
// reserved channels
//  pushover_config = 5;
//  opsgenie_config = 7;
//  victorops_config = 9;
//  wechat_config = 10;
swagger:model ListChannelsOKBodyChannelsItems0
*/
type ListChannelsOKBodyChannelsItems0 struct {

	// Machine-readable ID.
	ChannelID string `json:"channel_id,omitempty"`

	// Short human-readable summary.
	Summary string `json:"summary,omitempty"`

	// True if that channel is disabled.
	Disabled bool `json:"disabled,omitempty"`

	// email config
	EmailConfig *ListChannelsOKBodyChannelsItems0EmailConfig `json:"email_config,omitempty"`

	// pagerduty config
	PagerdutyConfig *ListChannelsOKBodyChannelsItems0PagerdutyConfig `json:"pagerduty_config,omitempty"`

	// slack config
	SlackConfig *ListChannelsOKBodyChannelsItems0SlackConfig `json:"slack_config,omitempty"`

	// webhook config
	WebhookConfig *ListChannelsOKBodyChannelsItems0WebhookConfig `json:"webhook_config,omitempty"`
}

// Validate validates this list channels OK body channels items0
func (o *ListChannelsOKBodyChannelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagerdutyConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlackConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWebhookConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) validateEmailConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.EmailConfig) { // not required
		return nil
	}

	if o.EmailConfig != nil {
		if err := o.EmailConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email_config")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) validatePagerdutyConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.PagerdutyConfig) { // not required
		return nil
	}

	if o.PagerdutyConfig != nil {
		if err := o.PagerdutyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerduty_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerduty_config")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) validateSlackConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.SlackConfig) { // not required
		return nil
	}

	if o.SlackConfig != nil {
		if err := o.SlackConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_config")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) validateWebhookConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.WebhookConfig) { // not required
		return nil
	}

	if o.WebhookConfig != nil {
		if err := o.WebhookConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list channels OK body channels items0 based on the context it is used
func (o *ListChannelsOKBodyChannelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEmailConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePagerdutyConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSlackConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWebhookConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) contextValidateEmailConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.EmailConfig != nil {
		if err := o.EmailConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email_config")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) contextValidatePagerdutyConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.PagerdutyConfig != nil {
		if err := o.PagerdutyConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerduty_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerduty_config")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) contextValidateSlackConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.SlackConfig != nil {
		if err := o.SlackConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_config")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0) contextValidateWebhookConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.WebhookConfig != nil {
		if err := o.WebhookConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0EmailConfig EmailConfig represents email configuration.
swagger:model ListChannelsOKBodyChannelsItems0EmailConfig
*/
type ListChannelsOKBodyChannelsItems0EmailConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// to
	To []string `json:"to"`
}

// Validate validates this list channels OK body channels items0 email config
func (o *ListChannelsOKBodyChannelsItems0EmailConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels OK body channels items0 email config based on context it is used
func (o *ListChannelsOKBodyChannelsItems0EmailConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0EmailConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0EmailConfig) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0EmailConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0PagerdutyConfig PagerDutyConfig represents PagerDuty configuration.
swagger:model ListChannelsOKBodyChannelsItems0PagerdutyConfig
*/
type ListChannelsOKBodyChannelsItems0PagerdutyConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// The PagerDuty key for "Events API v2" integration type. Exactly one key should be set.
	RoutingKey string `json:"routing_key,omitempty"`

	// The PagerDuty key for "Prometheus" integration type. Exactly one key should be set.
	ServiceKey string `json:"service_key,omitempty"`
}

// Validate validates this list channels OK body channels items0 pagerduty config
func (o *ListChannelsOKBodyChannelsItems0PagerdutyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels OK body channels items0 pagerduty config based on context it is used
func (o *ListChannelsOKBodyChannelsItems0PagerdutyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0PagerdutyConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0PagerdutyConfig) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0PagerdutyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0SlackConfig SlackConfig represents Slack configuration.
swagger:model ListChannelsOKBodyChannelsItems0SlackConfig
*/
type ListChannelsOKBodyChannelsItems0SlackConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`
}

// Validate validates this list channels OK body channels items0 slack config
func (o *ListChannelsOKBodyChannelsItems0SlackConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels OK body channels items0 slack config based on context it is used
func (o *ListChannelsOKBodyChannelsItems0SlackConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0SlackConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0SlackConfig) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0SlackConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0WebhookConfig WebhookConfig represents webhook configuration.
swagger:model ListChannelsOKBodyChannelsItems0WebhookConfig
*/
type ListChannelsOKBodyChannelsItems0WebhookConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// max alerts
	MaxAlerts int32 `json:"max_alerts,omitempty"`

	// http config
	HTTPConfig *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig `json:"http_config,omitempty"`
}

// Validate validates this list channels OK body channels items0 webhook config
func (o *ListChannelsOKBodyChannelsItems0WebhookConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBodyChannelsItems0WebhookConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.HTTPConfig) { // not required
		return nil
	}

	if o.HTTPConfig != nil {
		if err := o.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config" + "." + "http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config" + "." + "http_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list channels OK body channels items0 webhook config based on the context it is used
func (o *ListChannelsOKBodyChannelsItems0WebhookConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBodyChannelsItems0WebhookConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.HTTPConfig != nil {
		if err := o.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config" + "." + "http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config" + "." + "http_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfig) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0WebhookConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig HTTPConfig represents HTTP client configuration.
swagger:model ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig
*/
type ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig struct {

	// bearer token
	BearerToken string `json:"bearer_token,omitempty"`

	// bearer token file
	BearerTokenFile string `json:"bearer_token_file,omitempty"`

	// proxy url
	ProxyURL string `json:"proxy_url,omitempty"`

	// basic auth
	BasicAuth *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth `json:"basic_auth,omitempty"`

	// tls config
	TLSConfig *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig `json:"tls_config,omitempty"`
}

// Validate validates this list channels OK body channels items0 webhook config HTTP config
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) validateBasicAuth(formats strfmt.Registry) error {
	if swag.IsZero(o.BasicAuth) { // not required
		return nil
	}

	if o.BasicAuth != nil {
		if err := o.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config" + "." + "http_config" + "." + "basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config" + "." + "http_config" + "." + "basic_auth")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.TLSConfig) { // not required
		return nil
	}

	if o.TLSConfig != nil {
		if err := o.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config" + "." + "http_config" + "." + "tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config" + "." + "http_config" + "." + "tls_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list channels OK body channels items0 webhook config HTTP config based on the context it is used
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBasicAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) contextValidateBasicAuth(ctx context.Context, formats strfmt.Registry) error {

	if o.BasicAuth != nil {
		if err := o.BasicAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config" + "." + "http_config" + "." + "basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config" + "." + "http_config" + "." + "basic_auth")
			}
			return err
		}
	}

	return nil
}

func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.TLSConfig != nil {
		if err := o.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_config" + "." + "http_config" + "." + "tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_config" + "." + "http_config" + "." + "tls_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth BasicAuth represents basic HTTP auth configuration.
swagger:model ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth
*/
type ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth struct {

	// username
	Username string `json:"username,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// password file
	PasswordFile string `json:"password_file,omitempty"`
}

// Validate validates this list channels OK body channels items0 webhook config HTTP config basic auth
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels OK body channels items0 webhook config HTTP config basic auth based on context it is used
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigBasicAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig TLSConfig represents TLS configuration for alertmanager
// https://prometheus.io/docs/alerting/latest/configuration/#tls_config
swagger:model ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig
*/
type ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig struct {

	// A path to the CA certificate file to validate the server certificate with.
	// ca_file and ca_file_content should not be set at the same time.
	CaFile string `json:"ca_file,omitempty"`

	// A path to the certificate file for client cert authentication to the server.
	// cert_file and cert_file_content should not be set at the same time.
	CertFile string `json:"cert_file,omitempty"`

	// A path to the key file for client cert authentication to the server.
	// key_file and key_file_content should not be set at the same time.
	KeyFile string `json:"key_file,omitempty"`

	// Name of the server.
	ServerName string `json:"server_name,omitempty"`

	// Disable validation of the server certificate.
	InsecureSkipVerify bool `json:"insecure_skip_verify,omitempty"`

	// CA certificate to validate the server certificate with.
	// ca_file and ca_file_content should not be set at the same time.
	CaFileContent string `json:"ca_file_content,omitempty"`

	// A certificate for client cert authentication to the server.
	// cert_file and cert_file_content should not be set at the same time.
	CertFileContent string `json:"cert_file_content,omitempty"`

	// A key for client cert authentication to the server.
	// key_file and key_file_content should not be set at the same time.
	KeyFileContent string `json:"key_file_content,omitempty"`
}

// Validate validates this list channels OK body channels items0 webhook config HTTP config TLS config
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels OK body channels items0 webhook config HTTP config TLS config based on context it is used
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyChannelsItems0WebhookConfigHTTPConfigTLSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsOKBodyTotals PageTotals represents total values for pagination.
swagger:model ListChannelsOKBodyTotals
*/
type ListChannelsOKBodyTotals struct {

	// Total number of results.
	TotalItems int32 `json:"total_items,omitempty"`

	// Total number of pages.
	TotalPages int32 `json:"total_pages,omitempty"`
}

// Validate validates this list channels OK body totals
func (o *ListChannelsOKBodyTotals) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels OK body totals based on context it is used
func (o *ListChannelsOKBodyTotals) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsOKBodyTotals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsOKBodyTotals) UnmarshalBinary(b []byte) error {
	var res ListChannelsOKBodyTotals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListChannelsParamsBodyPageParams PageParams represents page request parameters for pagination.
swagger:model ListChannelsParamsBodyPageParams
*/
type ListChannelsParamsBodyPageParams struct {

	// Maximum number of results per page.
	PageSize int32 `json:"page_size,omitempty"`

	// Index of the requested page, starts from 0.
	Index int32 `json:"index,omitempty"`
}

// Validate validates this list channels params body page params
func (o *ListChannelsParamsBodyPageParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list channels params body page params based on context it is used
func (o *ListChannelsParamsBodyPageParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListChannelsParamsBodyPageParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListChannelsParamsBodyPageParams) UnmarshalBinary(b []byte) error {
	var res ListChannelsParamsBodyPageParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
