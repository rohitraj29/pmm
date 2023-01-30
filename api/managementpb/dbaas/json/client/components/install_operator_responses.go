// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstallOperatorReader is a Reader for the InstallOperator structure.
type InstallOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstallOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInstallOperatorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInstallOperatorOK creates a InstallOperatorOK with default headers values
func NewInstallOperatorOK() *InstallOperatorOK {
	return &InstallOperatorOK{}
}

/*
InstallOperatorOK describes a response with status code 200, with default header values.

A successful response.
*/
type InstallOperatorOK struct {
	Payload *InstallOperatorOKBody
}

func (o *InstallOperatorOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/InstallOperator][%d] installOperatorOk  %+v", 200, o.Payload)
}
func (o *InstallOperatorOK) GetPayload() *InstallOperatorOKBody {
	return o.Payload
}

func (o *InstallOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InstallOperatorOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallOperatorDefault creates a InstallOperatorDefault with default headers values
func NewInstallOperatorDefault(code int) *InstallOperatorDefault {
	return &InstallOperatorDefault{
		_statusCode: code,
	}
}

/*
InstallOperatorDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type InstallOperatorDefault struct {
	_statusCode int

	Payload *InstallOperatorDefaultBody
}

// Code gets the status code for the install operator default response
func (o *InstallOperatorDefault) Code() int {
	return o._statusCode
}

func (o *InstallOperatorDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/InstallOperator][%d] InstallOperator default  %+v", o._statusCode, o.Payload)
}
func (o *InstallOperatorDefault) GetPayload() *InstallOperatorDefaultBody {
	return o.Payload
}

func (o *InstallOperatorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InstallOperatorDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
InstallOperatorBody install operator body
swagger:model InstallOperatorBody
*/
type InstallOperatorBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// operator_type tells what operator we are interested in updating.
	OperatorType string `json:"operator_type,omitempty"`

	// version tells what version of the operator we should update to.
	Version string `json:"version,omitempty"`
}

// Validate validates this install operator body
func (o *InstallOperatorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this install operator body based on context it is used
func (o *InstallOperatorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InstallOperatorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstallOperatorBody) UnmarshalBinary(b []byte) error {
	var res InstallOperatorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InstallOperatorDefaultBody install operator default body
swagger:model InstallOperatorDefaultBody
*/
type InstallOperatorDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*InstallOperatorDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this install operator default body
func (o *InstallOperatorDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InstallOperatorDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("InstallOperator default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("InstallOperator default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this install operator default body based on the context it is used
func (o *InstallOperatorDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InstallOperatorDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InstallOperator default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("InstallOperator default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *InstallOperatorDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstallOperatorDefaultBody) UnmarshalBinary(b []byte) error {
	var res InstallOperatorDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InstallOperatorDefaultBodyDetailsItems0 install operator default body details items0
swagger:model InstallOperatorDefaultBodyDetailsItems0
*/
type InstallOperatorDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this install operator default body details items0
func (o *InstallOperatorDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this install operator default body details items0 based on context it is used
func (o *InstallOperatorDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InstallOperatorDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstallOperatorDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res InstallOperatorDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InstallOperatorOKBody install operator OK body
swagger:model InstallOperatorOKBody
*/
type InstallOperatorOKBody struct {

	// OperatorsStatus defines status of operators installed in Kubernetes cluster.
	//
	//  - OPERATORS_STATUS_INVALID: OPERATORS_STATUS_INVALID represents unknown state.
	//  - OPERATORS_STATUS_OK: OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	//  - OPERATORS_STATUS_UNSUPPORTED: OPERATORS_STATUS_UNSUPPORTED represents that operators are installed, but doesn't have supported API version.
	//  - OPERATORS_STATUS_NOT_INSTALLED: OPERATORS_STATUS_NOT_INSTALLED represents that operators are not installed.
	// Enum: [OPERATORS_STATUS_INVALID OPERATORS_STATUS_OK OPERATORS_STATUS_UNSUPPORTED OPERATORS_STATUS_NOT_INSTALLED]
	Status *string `json:"status,omitempty"`
}

// Validate validates this install operator OK body
func (o *InstallOperatorOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var installOperatorOkBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPERATORS_STATUS_INVALID","OPERATORS_STATUS_OK","OPERATORS_STATUS_UNSUPPORTED","OPERATORS_STATUS_NOT_INSTALLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		installOperatorOkBodyTypeStatusPropEnum = append(installOperatorOkBodyTypeStatusPropEnum, v)
	}
}

const (

	// InstallOperatorOKBodyStatusOPERATORSSTATUSINVALID captures enum value "OPERATORS_STATUS_INVALID"
	InstallOperatorOKBodyStatusOPERATORSSTATUSINVALID string = "OPERATORS_STATUS_INVALID"

	// InstallOperatorOKBodyStatusOPERATORSSTATUSOK captures enum value "OPERATORS_STATUS_OK"
	InstallOperatorOKBodyStatusOPERATORSSTATUSOK string = "OPERATORS_STATUS_OK"

	// InstallOperatorOKBodyStatusOPERATORSSTATUSUNSUPPORTED captures enum value "OPERATORS_STATUS_UNSUPPORTED"
	InstallOperatorOKBodyStatusOPERATORSSTATUSUNSUPPORTED string = "OPERATORS_STATUS_UNSUPPORTED"

	// InstallOperatorOKBodyStatusOPERATORSSTATUSNOTINSTALLED captures enum value "OPERATORS_STATUS_NOT_INSTALLED"
	InstallOperatorOKBodyStatusOPERATORSSTATUSNOTINSTALLED string = "OPERATORS_STATUS_NOT_INSTALLED"
)

// prop value enum
func (o *InstallOperatorOKBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, installOperatorOkBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *InstallOperatorOKBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("installOperatorOk"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this install operator OK body based on context it is used
func (o *InstallOperatorOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InstallOperatorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstallOperatorOKBody) UnmarshalBinary(b []byte) error {
	var res InstallOperatorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
