// Code generated by go-swagger; DO NOT EDIT.

package services

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

// AddPostgreSQLServiceReader is a Reader for the AddPostgreSQLService structure.
type AddPostgreSQLServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPostgreSQLServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPostgreSQLServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddPostgreSQLServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPostgreSQLServiceOK creates a AddPostgreSQLServiceOK with default headers values
func NewAddPostgreSQLServiceOK() *AddPostgreSQLServiceOK {
	return &AddPostgreSQLServiceOK{}
}

/*
AddPostgreSQLServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddPostgreSQLServiceOK struct {
	Payload *AddPostgreSQLServiceOKBody
}

func (o *AddPostgreSQLServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddPostgreSQL][%d] addPostgreSqlServiceOk  %+v", 200, o.Payload)
}
func (o *AddPostgreSQLServiceOK) GetPayload() *AddPostgreSQLServiceOKBody {
	return o.Payload
}

func (o *AddPostgreSQLServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPostgreSQLServiceDefault creates a AddPostgreSQLServiceDefault with default headers values
func NewAddPostgreSQLServiceDefault(code int) *AddPostgreSQLServiceDefault {
	return &AddPostgreSQLServiceDefault{
		_statusCode: code,
	}
}

/*
AddPostgreSQLServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddPostgreSQLServiceDefault struct {
	_statusCode int

	Payload *AddPostgreSQLServiceDefaultBody
}

// Code gets the status code for the add postgre SQL service default response
func (o *AddPostgreSQLServiceDefault) Code() int {
	return o._statusCode
}

func (o *AddPostgreSQLServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddPostgreSQL][%d] AddPostgreSQLService default  %+v", o._statusCode, o.Payload)
}
func (o *AddPostgreSQLServiceDefault) GetPayload() *AddPostgreSQLServiceDefaultBody {
	return o.Payload
}

func (o *AddPostgreSQLServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddPostgreSQLServiceBody add postgre SQL service body
swagger:model AddPostgreSQLServiceBody
*/
type AddPostgreSQLServiceBody struct {

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add postgre SQL service body
func (o *AddPostgreSQLServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add postgre SQL service body based on context it is used
func (o *AddPostgreSQLServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddPostgreSQLServiceDefaultBody add postgre SQL service default body
swagger:model AddPostgreSQLServiceDefaultBody
*/
type AddPostgreSQLServiceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddPostgreSQLServiceDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add postgre SQL service default body
func (o *AddPostgreSQLServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLServiceDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddPostgreSQLService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddPostgreSQLService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add postgre SQL service default body based on the context it is used
func (o *AddPostgreSQLServiceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLServiceDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddPostgreSQLService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddPostgreSQLService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddPostgreSQLServiceDefaultBodyDetailsItems0 add postgre SQL service default body details items0
swagger:model AddPostgreSQLServiceDefaultBodyDetailsItems0
*/
type AddPostgreSQLServiceDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add postgre SQL service default body details items0
func (o *AddPostgreSQLServiceDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add postgre SQL service default body details items0 based on context it is used
func (o *AddPostgreSQLServiceDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddPostgreSQLServiceOKBody add postgre SQL service OK body
swagger:model AddPostgreSQLServiceOKBody
*/
type AddPostgreSQLServiceOKBody struct {

	// postgresql
	Postgresql *AddPostgreSQLServiceOKBodyPostgresql `json:"postgresql,omitempty"`
}

// Validate validates this add postgre SQL service OK body
func (o *AddPostgreSQLServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePostgresql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLServiceOKBody) validatePostgresql(formats strfmt.Registry) error {
	if swag.IsZero(o.Postgresql) { // not required
		return nil
	}

	if o.Postgresql != nil {
		if err := o.Postgresql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlServiceOk" + "." + "postgresql")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addPostgreSqlServiceOk" + "." + "postgresql")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add postgre SQL service OK body based on the context it is used
func (o *AddPostgreSQLServiceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePostgresql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLServiceOKBody) contextValidatePostgresql(ctx context.Context, formats strfmt.Registry) error {

	if o.Postgresql != nil {
		if err := o.Postgresql.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlServiceOk" + "." + "postgresql")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addPostgreSqlServiceOk" + "." + "postgresql")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddPostgreSQLServiceOKBodyPostgresql PostgreSQLService represents a generic PostgreSQL instance.
swagger:model AddPostgreSQLServiceOKBodyPostgresql
*/
type AddPostgreSQLServiceOKBodyPostgresql struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Database name.
	DatabaseName string `json:"database_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add postgre SQL service OK body postgresql
func (o *AddPostgreSQLServiceOKBodyPostgresql) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add postgre SQL service OK body postgresql based on context it is used
func (o *AddPostgreSQLServiceOKBodyPostgresql) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBodyPostgresql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBodyPostgresql) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceOKBodyPostgresql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
