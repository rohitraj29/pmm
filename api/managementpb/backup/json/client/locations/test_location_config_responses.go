// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// TestLocationConfigReader is a Reader for the TestLocationConfig structure.
type TestLocationConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestLocationConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestLocationConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTestLocationConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTestLocationConfigOK creates a TestLocationConfigOK with default headers values
func NewTestLocationConfigOK() *TestLocationConfigOK {
	return &TestLocationConfigOK{}
}

/* TestLocationConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type TestLocationConfigOK struct {
	Payload interface{}
}

func (o *TestLocationConfigOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/TestConfig][%d] testLocationConfigOk  %+v", 200, o.Payload)
}
func (o *TestLocationConfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TestLocationConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLocationConfigDefault creates a TestLocationConfigDefault with default headers values
func NewTestLocationConfigDefault(code int) *TestLocationConfigDefault {
	return &TestLocationConfigDefault{
		_statusCode: code,
	}
}

/* TestLocationConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TestLocationConfigDefault struct {
	_statusCode int

	Payload *TestLocationConfigDefaultBody
}

// Code gets the status code for the test location config default response
func (o *TestLocationConfigDefault) Code() int {
	return o._statusCode
}

func (o *TestLocationConfigDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/TestConfig][%d] TestLocationConfig default  %+v", o._statusCode, o.Payload)
}
func (o *TestLocationConfigDefault) GetPayload() *TestLocationConfigDefaultBody {
	return o.Payload
}

func (o *TestLocationConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TestLocationConfigDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TestLocationConfigBody test location config body
swagger:model TestLocationConfigBody
*/
type TestLocationConfigBody struct {

	// pmm client config
	PMMClientConfig *TestLocationConfigParamsBodyPMMClientConfig `json:"pmm_client_config,omitempty"`

	// pmm server config
	PMMServerConfig *TestLocationConfigParamsBodyPMMServerConfig `json:"pmm_server_config,omitempty"`

	// s3 config
	S3Config *TestLocationConfigParamsBodyS3Config `json:"s3_config,omitempty"`
}

// Validate validates this test location config body
func (o *TestLocationConfigBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePMMClientConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateS3Config(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TestLocationConfigBody) validatePMMClientConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.PMMClientConfig) { // not required
		return nil
	}

	if o.PMMClientConfig != nil {
		if err := o.PMMClientConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_client_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_client_config")
			}
			return err
		}
	}

	return nil
}

func (o *TestLocationConfigBody) validatePMMServerConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.PMMServerConfig) { // not required
		return nil
	}

	if o.PMMServerConfig != nil {
		if err := o.PMMServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_server_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_server_config")
			}
			return err
		}
	}

	return nil
}

func (o *TestLocationConfigBody) validateS3Config(formats strfmt.Registry) error {
	if swag.IsZero(o.S3Config) { // not required
		return nil
	}

	if o.S3Config != nil {
		if err := o.S3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "s3_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "s3_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this test location config body based on the context it is used
func (o *TestLocationConfigBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePMMClientConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePMMServerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateS3Config(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TestLocationConfigBody) contextValidatePMMClientConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.PMMClientConfig != nil {
		if err := o.PMMClientConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_client_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_client_config")
			}
			return err
		}
	}

	return nil
}

func (o *TestLocationConfigBody) contextValidatePMMServerConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.PMMServerConfig != nil {
		if err := o.PMMServerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_server_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_server_config")
			}
			return err
		}
	}

	return nil
}

func (o *TestLocationConfigBody) contextValidateS3Config(ctx context.Context, formats strfmt.Registry) error {

	if o.S3Config != nil {
		if err := o.S3Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "s3_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "s3_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TestLocationConfigBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestLocationConfigBody) UnmarshalBinary(b []byte) error {
	var res TestLocationConfigBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TestLocationConfigDefaultBody test location config default body
swagger:model TestLocationConfigDefaultBody
*/
type TestLocationConfigDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*TestLocationConfigDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this test location config default body
func (o *TestLocationConfigDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TestLocationConfigDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("TestLocationConfig default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TestLocationConfig default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this test location config default body based on the context it is used
func (o *TestLocationConfigDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TestLocationConfigDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TestLocationConfig default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TestLocationConfig default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TestLocationConfigDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestLocationConfigDefaultBody) UnmarshalBinary(b []byte) error {
	var res TestLocationConfigDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TestLocationConfigDefaultBodyDetailsItems0 test location config default body details items0
swagger:model TestLocationConfigDefaultBodyDetailsItems0
*/
type TestLocationConfigDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this test location config default body details items0
func (o *TestLocationConfigDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this test location config default body details items0 based on context it is used
func (o *TestLocationConfigDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TestLocationConfigDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestLocationConfigDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res TestLocationConfigDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TestLocationConfigParamsBodyPMMClientConfig PMMClientLocationConfig represents file system config inside pmm-client.
swagger:model TestLocationConfigParamsBodyPMMClientConfig
*/
type TestLocationConfigParamsBodyPMMClientConfig struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this test location config params body PMM client config
func (o *TestLocationConfigParamsBodyPMMClientConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this test location config params body PMM client config based on context it is used
func (o *TestLocationConfigParamsBodyPMMClientConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TestLocationConfigParamsBodyPMMClientConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestLocationConfigParamsBodyPMMClientConfig) UnmarshalBinary(b []byte) error {
	var res TestLocationConfigParamsBodyPMMClientConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TestLocationConfigParamsBodyPMMServerConfig PMMServerLocationConfig represents file system config inside pmm-server.
swagger:model TestLocationConfigParamsBodyPMMServerConfig
*/
type TestLocationConfigParamsBodyPMMServerConfig struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this test location config params body PMM server config
func (o *TestLocationConfigParamsBodyPMMServerConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this test location config params body PMM server config based on context it is used
func (o *TestLocationConfigParamsBodyPMMServerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TestLocationConfigParamsBodyPMMServerConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestLocationConfigParamsBodyPMMServerConfig) UnmarshalBinary(b []byte) error {
	var res TestLocationConfigParamsBodyPMMServerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TestLocationConfigParamsBodyS3Config S3LocationConfig represents S3 bucket configuration.
swagger:model TestLocationConfigParamsBodyS3Config
*/
type TestLocationConfigParamsBodyS3Config struct {

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// access key
	AccessKey string `json:"access_key,omitempty"`

	// secret key
	SecretKey string `json:"secret_key,omitempty"`

	// bucket name
	BucketName string `json:"bucket_name,omitempty"`
}

// Validate validates this test location config params body s3 config
func (o *TestLocationConfigParamsBodyS3Config) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this test location config params body s3 config based on context it is used
func (o *TestLocationConfigParamsBodyS3Config) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TestLocationConfigParamsBodyS3Config) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestLocationConfigParamsBodyS3Config) UnmarshalBinary(b []byte) error {
	var res TestLocationConfigParamsBodyS3Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
