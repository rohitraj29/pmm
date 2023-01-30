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

// StartPostgreSQLShowCreateTableActionReader is a Reader for the StartPostgreSQLShowCreateTableAction structure.
type StartPostgreSQLShowCreateTableActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPostgreSQLShowCreateTableActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPostgreSQLShowCreateTableActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartPostgreSQLShowCreateTableActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPostgreSQLShowCreateTableActionOK creates a StartPostgreSQLShowCreateTableActionOK with default headers values
func NewStartPostgreSQLShowCreateTableActionOK() *StartPostgreSQLShowCreateTableActionOK {
	return &StartPostgreSQLShowCreateTableActionOK{}
}

/*
StartPostgreSQLShowCreateTableActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartPostgreSQLShowCreateTableActionOK struct {
	Payload *StartPostgreSQLShowCreateTableActionOKBody
}

func (o *StartPostgreSQLShowCreateTableActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPostgreSQLShowCreateTable][%d] startPostgreSqlShowCreateTableActionOk  %+v", 200, o.Payload)
}
func (o *StartPostgreSQLShowCreateTableActionOK) GetPayload() *StartPostgreSQLShowCreateTableActionOKBody {
	return o.Payload
}

func (o *StartPostgreSQLShowCreateTableActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPostgreSQLShowCreateTableActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPostgreSQLShowCreateTableActionDefault creates a StartPostgreSQLShowCreateTableActionDefault with default headers values
func NewStartPostgreSQLShowCreateTableActionDefault(code int) *StartPostgreSQLShowCreateTableActionDefault {
	return &StartPostgreSQLShowCreateTableActionDefault{
		_statusCode: code,
	}
}

/*
StartPostgreSQLShowCreateTableActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartPostgreSQLShowCreateTableActionDefault struct {
	_statusCode int

	Payload *StartPostgreSQLShowCreateTableActionDefaultBody
}

// Code gets the status code for the start postgre SQL show create table action default response
func (o *StartPostgreSQLShowCreateTableActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPostgreSQLShowCreateTableActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPostgreSQLShowCreateTable][%d] StartPostgreSQLShowCreateTableAction default  %+v", o._statusCode, o.Payload)
}
func (o *StartPostgreSQLShowCreateTableActionDefault) GetPayload() *StartPostgreSQLShowCreateTableActionDefaultBody {
	return o.Payload
}

func (o *StartPostgreSQLShowCreateTableActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPostgreSQLShowCreateTableActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
StartPostgreSQLShowCreateTableActionBody start postgre SQL show create table action body
swagger:model StartPostgreSQLShowCreateTableActionBody
*/
type StartPostgreSQLShowCreateTableActionBody struct {

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Table name. Required. May additionally contain a database name.
	TableName string `json:"table_name,omitempty"`

	// Database name. Required if not given in the table_name field.
	Database string `json:"database,omitempty"`
}

// Validate validates this start postgre SQL show create table action body
func (o *StartPostgreSQLShowCreateTableActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start postgre SQL show create table action body based on context it is used
func (o *StartPostgreSQLShowCreateTableActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowCreateTableActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPostgreSQLShowCreateTableActionDefaultBody start postgre SQL show create table action default body
swagger:model StartPostgreSQLShowCreateTableActionDefaultBody
*/
type StartPostgreSQLShowCreateTableActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start postgre SQL show create table action default body
func (o *StartPostgreSQLShowCreateTableActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPostgreSQLShowCreateTableActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartPostgreSQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPostgreSQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start postgre SQL show create table action default body based on the context it is used
func (o *StartPostgreSQLShowCreateTableActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPostgreSQLShowCreateTableActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartPostgreSQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPostgreSQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowCreateTableActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0 start postgre SQL show create table action default body details items0
swagger:model StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0
*/
type StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this start postgre SQL show create table action default body details items0
func (o *StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start postgre SQL show create table action default body details items0 based on context it is used
func (o *StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowCreateTableActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPostgreSQLShowCreateTableActionOKBody start postgre SQL show create table action OK body
swagger:model StartPostgreSQLShowCreateTableActionOKBody
*/
type StartPostgreSQLShowCreateTableActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start postgre SQL show create table action OK body
func (o *StartPostgreSQLShowCreateTableActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start postgre SQL show create table action OK body based on context it is used
func (o *StartPostgreSQLShowCreateTableActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowCreateTableActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
