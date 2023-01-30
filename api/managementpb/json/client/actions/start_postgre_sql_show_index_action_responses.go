// Code generated by go-swagger; DO NOT EDIT.

package actions

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

// StartPostgreSQLShowIndexActionReader is a Reader for the StartPostgreSQLShowIndexAction structure.
type StartPostgreSQLShowIndexActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPostgreSQLShowIndexActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPostgreSQLShowIndexActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartPostgreSQLShowIndexActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPostgreSQLShowIndexActionOK creates a StartPostgreSQLShowIndexActionOK with default headers values
func NewStartPostgreSQLShowIndexActionOK() *StartPostgreSQLShowIndexActionOK {
	return &StartPostgreSQLShowIndexActionOK{}
}

/*
StartPostgreSQLShowIndexActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartPostgreSQLShowIndexActionOK struct {
	Payload *StartPostgreSQLShowIndexActionOKBody
}

func (o *StartPostgreSQLShowIndexActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPostgreSQLShowIndex][%d] startPostgreSqlShowIndexActionOk  %+v", 200, o.Payload)
}
func (o *StartPostgreSQLShowIndexActionOK) GetPayload() *StartPostgreSQLShowIndexActionOKBody {
	return o.Payload
}

func (o *StartPostgreSQLShowIndexActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPostgreSQLShowIndexActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPostgreSQLShowIndexActionDefault creates a StartPostgreSQLShowIndexActionDefault with default headers values
func NewStartPostgreSQLShowIndexActionDefault(code int) *StartPostgreSQLShowIndexActionDefault {
	return &StartPostgreSQLShowIndexActionDefault{
		_statusCode: code,
	}
}

/*
StartPostgreSQLShowIndexActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartPostgreSQLShowIndexActionDefault struct {
	_statusCode int

	Payload *StartPostgreSQLShowIndexActionDefaultBody
}

// Code gets the status code for the start postgre SQL show index action default response
func (o *StartPostgreSQLShowIndexActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPostgreSQLShowIndexActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPostgreSQLShowIndex][%d] StartPostgreSQLShowIndexAction default  %+v", o._statusCode, o.Payload)
}
func (o *StartPostgreSQLShowIndexActionDefault) GetPayload() *StartPostgreSQLShowIndexActionDefaultBody {
	return o.Payload
}

func (o *StartPostgreSQLShowIndexActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPostgreSQLShowIndexActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
StartPostgreSQLShowIndexActionBody start postgre SQL show index action body
swagger:model StartPostgreSQLShowIndexActionBody
*/
type StartPostgreSQLShowIndexActionBody struct {

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Table name. Required. May additionally contain a database name.
	TableName string `json:"table_name,omitempty"`

	// Database name. Required if not given in the table_name field.
	Database string `json:"database,omitempty"`
}

// Validate validates this start postgre SQL show index action body
func (o *StartPostgreSQLShowIndexActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start postgre SQL show index action body based on context it is used
func (o *StartPostgreSQLShowIndexActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowIndexActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPostgreSQLShowIndexActionDefaultBody start postgre SQL show index action default body
swagger:model StartPostgreSQLShowIndexActionDefaultBody
*/
type StartPostgreSQLShowIndexActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start postgre SQL show index action default body
func (o *StartPostgreSQLShowIndexActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPostgreSQLShowIndexActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartPostgreSQLShowIndexAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPostgreSQLShowIndexAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start postgre SQL show index action default body based on the context it is used
func (o *StartPostgreSQLShowIndexActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPostgreSQLShowIndexActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartPostgreSQLShowIndexAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPostgreSQLShowIndexAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowIndexActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0 start postgre SQL show index action default body details items0
swagger:model StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0
*/
type StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this start postgre SQL show index action default body details items0
func (o *StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start postgre SQL show index action default body details items0 based on context it is used
func (o *StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowIndexActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPostgreSQLShowIndexActionOKBody start postgre SQL show index action OK body
swagger:model StartPostgreSQLShowIndexActionOKBody
*/
type StartPostgreSQLShowIndexActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start postgre SQL show index action OK body
func (o *StartPostgreSQLShowIndexActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start postgre SQL show index action OK body based on context it is used
func (o *StartPostgreSQLShowIndexActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowIndexActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowIndexActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
