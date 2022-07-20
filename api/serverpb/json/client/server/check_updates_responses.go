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
	"github.com/go-openapi/validate"
)

// CheckUpdatesReader is a Reader for the CheckUpdates structure.
type CheckUpdatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckUpdatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckUpdatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckUpdatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckUpdatesOK creates a CheckUpdatesOK with default headers values
func NewCheckUpdatesOK() *CheckUpdatesOK {
	return &CheckUpdatesOK{}
}

/* CheckUpdatesOK describes a response with status code 200, with default header values.

A successful response.
*/
type CheckUpdatesOK struct {
	Payload *CheckUpdatesOKBody
}

func (o *CheckUpdatesOK) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Check][%d] checkUpdatesOk  %+v", 200, o.Payload)
}
func (o *CheckUpdatesOK) GetPayload() *CheckUpdatesOKBody {
	return o.Payload
}

func (o *CheckUpdatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckUpdatesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckUpdatesDefault creates a CheckUpdatesDefault with default headers values
func NewCheckUpdatesDefault(code int) *CheckUpdatesDefault {
	return &CheckUpdatesDefault{
		_statusCode: code,
	}
}

/* CheckUpdatesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CheckUpdatesDefault struct {
	_statusCode int

	Payload *CheckUpdatesDefaultBody
}

// Code gets the status code for the check updates default response
func (o *CheckUpdatesDefault) Code() int {
	return o._statusCode
}

func (o *CheckUpdatesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Check][%d] CheckUpdates default  %+v", o._statusCode, o.Payload)
}
func (o *CheckUpdatesDefault) GetPayload() *CheckUpdatesDefaultBody {
	return o.Payload
}

func (o *CheckUpdatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckUpdatesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckUpdatesBody check updates body
swagger:model CheckUpdatesBody
*/
type CheckUpdatesBody struct {

	// If false, cached information may be returned.
	Force bool `json:"force,omitempty"`

	// If true, only installed version will be in response.
	OnlyInstalledVersion bool `json:"only_installed_version,omitempty"`
}

// Validate validates this check updates body
func (o *CheckUpdatesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check updates body based on context it is used
func (o *CheckUpdatesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesBody) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesDefaultBody check updates default body
swagger:model CheckUpdatesDefaultBody
*/
type CheckUpdatesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*CheckUpdatesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this check updates default body
func (o *CheckUpdatesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this check updates default body based on the context it is used
func (o *CheckUpdatesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesDefaultBody) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
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
swagger:model CheckUpdatesDefaultBodyDetailsItems0
*/
type CheckUpdatesDefaultBodyDetailsItems0 struct {

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

// Validate validates this check updates default body details items0
func (o *CheckUpdatesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check updates default body details items0 based on context it is used
func (o *CheckUpdatesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesOKBody check updates OK body
swagger:model CheckUpdatesOKBody
*/
type CheckUpdatesOKBody struct {

	// True if there is a PMM Server update available.
	UpdateAvailable bool `json:"update_available,omitempty"`

	// Latest available PMM Server release announcement URL.
	LatestNewsURL string `json:"latest_news_url,omitempty"`

	// Last check time.
	// Format: date-time
	LastCheck strfmt.DateTime `json:"last_check,omitempty"`

	// installed
	Installed *CheckUpdatesOKBodyInstalled `json:"installed,omitempty"`

	// latest
	Latest *CheckUpdatesOKBodyLatest `json:"latest,omitempty"`
}

// Validate validates this check updates OK body
func (o *CheckUpdatesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLastCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInstalled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLatest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBody) validateLastCheck(formats strfmt.Registry) error {
	if swag.IsZero(o.LastCheck) { // not required
		return nil
	}

	if err := validate.FormatOf("checkUpdatesOk"+"."+"last_check", "body", "date-time", o.LastCheck.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CheckUpdatesOKBody) validateInstalled(formats strfmt.Registry) error {
	if swag.IsZero(o.Installed) { // not required
		return nil
	}

	if o.Installed != nil {
		if err := o.Installed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "installed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "installed")
			}
			return err
		}
	}

	return nil
}

func (o *CheckUpdatesOKBody) validateLatest(formats strfmt.Registry) error {
	if swag.IsZero(o.Latest) { // not required
		return nil
	}

	if o.Latest != nil {
		if err := o.Latest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "latest")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this check updates OK body based on the context it is used
func (o *CheckUpdatesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateInstalled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLatest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBody) contextValidateInstalled(ctx context.Context, formats strfmt.Registry) error {

	if o.Installed != nil {
		if err := o.Installed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "installed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "installed")
			}
			return err
		}
	}

	return nil
}

func (o *CheckUpdatesOKBody) contextValidateLatest(ctx context.Context, formats strfmt.Registry) error {

	if o.Latest != nil {
		if err := o.Latest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "latest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesOKBody) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesOKBodyInstalled VersionInfo describes component version, or PMM Server as a whole.
swagger:model CheckUpdatesOKBodyInstalled
*/
type CheckUpdatesOKBodyInstalled struct {

	// User-visible version.
	Version string `json:"version,omitempty"`

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this check updates OK body installed
func (o *CheckUpdatesOKBodyInstalled) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBodyInstalled) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("checkUpdatesOk"+"."+"installed"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this check updates OK body installed based on context it is used
func (o *CheckUpdatesOKBodyInstalled) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesOKBodyInstalled) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesOKBodyInstalled) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesOKBodyInstalled
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesOKBodyLatest VersionInfo describes component version, or PMM Server as a whole.
swagger:model CheckUpdatesOKBodyLatest
*/
type CheckUpdatesOKBodyLatest struct {

	// User-visible version.
	Version string `json:"version,omitempty"`

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this check updates OK body latest
func (o *CheckUpdatesOKBodyLatest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBodyLatest) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("checkUpdatesOk"+"."+"latest"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this check updates OK body latest based on context it is used
func (o *CheckUpdatesOKBodyLatest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesOKBodyLatest) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesOKBodyLatest) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesOKBodyLatest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
