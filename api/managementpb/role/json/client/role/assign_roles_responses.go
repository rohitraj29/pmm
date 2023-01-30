// Code generated by go-swagger; DO NOT EDIT.

package role

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

// AssignRolesReader is a Reader for the AssignRoles structure.
type AssignRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignRolesOK creates a AssignRolesOK with default headers values
func NewAssignRolesOK() *AssignRolesOK {
	return &AssignRolesOK{}
}

/*
AssignRolesOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssignRolesOK struct {
	Payload interface{}
}

func (o *AssignRolesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Role/Assign][%d] assignRolesOk  %+v", 200, o.Payload)
}
func (o *AssignRolesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AssignRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignRolesDefault creates a AssignRolesDefault with default headers values
func NewAssignRolesDefault(code int) *AssignRolesDefault {
	return &AssignRolesDefault{
		_statusCode: code,
	}
}

/*
AssignRolesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssignRolesDefault struct {
	_statusCode int

	Payload *AssignRolesDefaultBody
}

// Code gets the status code for the assign roles default response
func (o *AssignRolesDefault) Code() int {
	return o._statusCode
}

func (o *AssignRolesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Role/Assign][%d] AssignRoles default  %+v", o._statusCode, o.Payload)
}
func (o *AssignRolesDefault) GetPayload() *AssignRolesDefaultBody {
	return o.Payload
}

func (o *AssignRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AssignRolesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AssignRolesBody assign roles body
swagger:model AssignRolesBody
*/
type AssignRolesBody struct {

	// role ids
	RoleIds []int64 `json:"role_ids"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this assign roles body
func (o *AssignRolesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assign roles body based on context it is used
func (o *AssignRolesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignRolesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignRolesBody) UnmarshalBinary(b []byte) error {
	var res AssignRolesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AssignRolesDefaultBody assign roles default body
swagger:model AssignRolesDefaultBody
*/
type AssignRolesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AssignRolesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this assign roles default body
func (o *AssignRolesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignRolesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AssignRoles default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AssignRoles default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this assign roles default body based on the context it is used
func (o *AssignRolesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignRolesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AssignRoles default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AssignRoles default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssignRolesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignRolesDefaultBody) UnmarshalBinary(b []byte) error {
	var res AssignRolesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AssignRolesDefaultBodyDetailsItems0 assign roles default body details items0
swagger:model AssignRolesDefaultBodyDetailsItems0
*/
type AssignRolesDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this assign roles default body details items0
func (o *AssignRolesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assign roles default body details items0 based on context it is used
func (o *AssignRolesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignRolesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignRolesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AssignRolesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
