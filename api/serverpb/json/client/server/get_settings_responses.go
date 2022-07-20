// Code generated by go-swagger; DO NOT EDIT.

package server

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

// GetSettingsReader is a Reader for the GetSettings structure.
type GetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSettingsOK creates a GetSettingsOK with default headers values
func NewGetSettingsOK() *GetSettingsOK {
	return &GetSettingsOK{}
}

/* GetSettingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetSettingsOK struct {
	Payload *GetSettingsOKBody
}

func (o *GetSettingsOK) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Get][%d] getSettingsOk  %+v", 200, o.Payload)
}
func (o *GetSettingsOK) GetPayload() *GetSettingsOKBody {
	return o.Payload
}

func (o *GetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsDefault creates a GetSettingsDefault with default headers values
func NewGetSettingsDefault(code int) *GetSettingsDefault {
	return &GetSettingsDefault{
		_statusCode: code,
	}
}

/* GetSettingsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetSettingsDefault struct {
	_statusCode int

	Payload *GetSettingsDefaultBody
}

// Code gets the status code for the get settings default response
func (o *GetSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetSettingsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Get][%d] GetSettings default  %+v", o._statusCode, o.Payload)
}
func (o *GetSettingsDefault) GetPayload() *GetSettingsDefaultBody {
	return o.Payload
}

func (o *GetSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSettingsDefaultBody get settings default body
swagger:model GetSettingsDefaultBody
*/
type GetSettingsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetSettingsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get settings default body
func (o *GetSettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetSettings default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get settings default body based on the context it is used
func (o *GetSettingsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetSettings default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetSettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
// Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
// Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := anypb.New(foo)
//      if err != nil {
//        ...
//      }
//      ...
//      foo := &pb.Foo{}
//      if err := any.UnmarshalTo(foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
//
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model GetSettingsDefaultBodyDetailsItems0
*/
type GetSettingsDefaultBodyDetailsItems0 struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get settings default body details items0
func (o *GetSettingsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get settings default body details items0 based on context it is used
func (o *GetSettingsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetSettingsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBody get settings OK body
swagger:model GetSettingsOKBody
*/
type GetSettingsOKBody struct {

	// settings
	Settings *GetSettingsOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this get settings OK body
func (o *GetSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBody) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get settings OK body based on the context it is used
func (o *GetSettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBody) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.Settings != nil {
		if err := o.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettings Settings represents PMM Server settings.
swagger:model GetSettingsOKBodySettings
*/
type GetSettingsOKBodySettings struct {

	// True if updates are disabled.
	UpdatesDisabled bool `json:"updates_disabled,omitempty"`

	// True if telemetry is enabled.
	TelemetryEnabled bool `json:"telemetry_enabled,omitempty"`

	// data retention
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// External AlertManager URL (e.g., https://username:password@1.2.3.4/path).
	AlertManagerURL string `json:"alert_manager_url,omitempty"`

	// Custom alerting or recording rules.
	AlertManagerRules string `json:"alert_manager_rules,omitempty"`

	// True if Security Threat Tool is enabled.
	SttEnabled bool `json:"stt_enabled,omitempty"`

	// Percona Platform user's email, if this PMM instance is linked to the Platform.
	PlatformEmail string `json:"platform_email,omitempty"`

	// True if DBaaS is enabled.
	DbaasEnabled bool `json:"dbaas_enabled,omitempty"`

	// True if Alerting is enabled.
	AlertingEnabled bool `json:"alerting_enabled,omitempty"`

	// PMM Server public address.
	PMMPublicAddress string `json:"pmm_public_address,omitempty"`

	// True if Backup Management is enabled.
	BackupManagementEnabled bool `json:"backup_management_enabled,omitempty"`

	// True if Azure Discover is enabled.
	AzurediscoverEnabled bool `json:"azurediscover_enabled,omitempty"`

	// True if the PMM instance is connected to Platform
	ConnectedToPlatform bool `json:"connected_to_platform,omitempty"`

	// email alerting settings
	EmailAlertingSettings *GetSettingsOKBodySettingsEmailAlertingSettings `json:"email_alerting_settings,omitempty"`

	// metrics resolutions
	MetricsResolutions *GetSettingsOKBodySettingsMetricsResolutions `json:"metrics_resolutions,omitempty"`

	// slack alerting settings
	SlackAlertingSettings *GetSettingsOKBodySettingsSlackAlertingSettings `json:"slack_alerting_settings,omitempty"`

	// stt check intervals
	SttCheckIntervals *GetSettingsOKBodySettingsSttCheckIntervals `json:"stt_check_intervals,omitempty"`
}

// Validate validates this get settings OK body settings
func (o *GetSettingsOKBodySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlackAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSttCheckIntervals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBodySettings) validateEmailAlertingSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.EmailAlertingSettings) { // not required
		return nil
	}

	if o.EmailAlertingSettings != nil {
		if err := o.EmailAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "email_alerting_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "email_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) validateMetricsResolutions(formats strfmt.Registry) error {
	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) validateSlackAlertingSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.SlackAlertingSettings) { // not required
		return nil
	}

	if o.SlackAlertingSettings != nil {
		if err := o.SlackAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "slack_alerting_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "slack_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) validateSttCheckIntervals(formats strfmt.Registry) error {
	if swag.IsZero(o.SttCheckIntervals) { // not required
		return nil
	}

	if o.SttCheckIntervals != nil {
		if err := o.SttCheckIntervals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "stt_check_intervals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "stt_check_intervals")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get settings OK body settings based on the context it is used
func (o *GetSettingsOKBodySettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEmailAlertingSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMetricsResolutions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSlackAlertingSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSttCheckIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBodySettings) contextValidateEmailAlertingSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.EmailAlertingSettings != nil {
		if err := o.EmailAlertingSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "email_alerting_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "email_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) contextValidateMetricsResolutions(ctx context.Context, formats strfmt.Registry) error {

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) contextValidateSlackAlertingSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.SlackAlertingSettings != nil {
		if err := o.SlackAlertingSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "slack_alerting_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "slack_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) contextValidateSttCheckIntervals(ctx context.Context, formats strfmt.Registry) error {

	if o.SttCheckIntervals != nil {
		if err := o.SttCheckIntervals.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "stt_check_intervals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSettingsOk" + "." + "settings" + "." + "stt_check_intervals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettings) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsEmailAlertingSettings EmailAlertingSettings represents email (SMTP) configuration for Alerting.
swagger:model GetSettingsOKBodySettingsEmailAlertingSettings
*/
type GetSettingsOKBodySettingsEmailAlertingSettings struct {

	// SMTP From header field.
	From string `json:"from,omitempty"`

	// SMTP host and port.
	Smarthost string `json:"smarthost,omitempty"`

	// Hostname to identify to the SMTP server.
	Hello string `json:"hello,omitempty"`

	// Auth using CRAM-MD5, LOGIN and PLAIN.
	Username string `json:"username,omitempty"`

	// Auth using LOGIN and PLAIN.
	Password string `json:"password,omitempty"`

	// Auth using PLAIN.
	Identity string `json:"identity,omitempty"`

	// Auth using CRAM-MD5.
	Secret string `json:"secret,omitempty"`

	// Require TLS.
	RequireTLS bool `json:"require_tls,omitempty"`
}

// Validate validates this get settings OK body settings email alerting settings
func (o *GetSettingsOKBodySettingsEmailAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get settings OK body settings email alerting settings based on context it is used
func (o *GetSettingsOKBodySettingsEmailAlertingSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsEmailAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsEmailAlertingSettings) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsEmailAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model GetSettingsOKBodySettingsMetricsResolutions
*/
type GetSettingsOKBodySettingsMetricsResolutions struct {

	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this get settings OK body settings metrics resolutions
func (o *GetSettingsOKBodySettingsMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get settings OK body settings metrics resolutions based on context it is used
func (o *GetSettingsOKBodySettingsMetricsResolutions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsSlackAlertingSettings SlackAlertingSettings represents Slack configuration for Alerting.
swagger:model GetSettingsOKBodySettingsSlackAlertingSettings
*/
type GetSettingsOKBodySettingsSlackAlertingSettings struct {

	// Slack API (webhook) URL.
	URL string `json:"url,omitempty"`
}

// Validate validates this get settings OK body settings slack alerting settings
func (o *GetSettingsOKBodySettingsSlackAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get settings OK body settings slack alerting settings based on context it is used
func (o *GetSettingsOKBodySettingsSlackAlertingSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSlackAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSlackAlertingSettings) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsSlackAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsSttCheckIntervals STTCheckIntervals represents intervals between STT checks.
swagger:model GetSettingsOKBodySettingsSttCheckIntervals
*/
type GetSettingsOKBodySettingsSttCheckIntervals struct {

	// Standard check interval.
	StandardInterval string `json:"standard_interval,omitempty"`

	// Interval for rare check runs.
	RareInterval string `json:"rare_interval,omitempty"`

	// Interval for frequent check runs.
	FrequentInterval string `json:"frequent_interval,omitempty"`
}

// Validate validates this get settings OK body settings stt check intervals
func (o *GetSettingsOKBodySettingsSttCheckIntervals) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get settings OK body settings stt check intervals based on context it is used
func (o *GetSettingsOKBodySettingsSttCheckIntervals) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSttCheckIntervals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSttCheckIntervals) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsSttCheckIntervals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
