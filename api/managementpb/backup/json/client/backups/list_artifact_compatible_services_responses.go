// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// ListArtifactCompatibleServicesReader is a Reader for the ListArtifactCompatibleServices structure.
type ListArtifactCompatibleServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArtifactCompatibleServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArtifactCompatibleServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListArtifactCompatibleServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArtifactCompatibleServicesOK creates a ListArtifactCompatibleServicesOK with default headers values
func NewListArtifactCompatibleServicesOK() *ListArtifactCompatibleServicesOK {
	return &ListArtifactCompatibleServicesOK{}
}

/*
ListArtifactCompatibleServicesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListArtifactCompatibleServicesOK struct {
	Payload *ListArtifactCompatibleServicesOKBody
}

func (o *ListArtifactCompatibleServicesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ListArtifactCompatibleServices][%d] listArtifactCompatibleServicesOk  %+v", 200, o.Payload)
}
func (o *ListArtifactCompatibleServicesOK) GetPayload() *ListArtifactCompatibleServicesOKBody {
	return o.Payload
}

func (o *ListArtifactCompatibleServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListArtifactCompatibleServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactCompatibleServicesDefault creates a ListArtifactCompatibleServicesDefault with default headers values
func NewListArtifactCompatibleServicesDefault(code int) *ListArtifactCompatibleServicesDefault {
	return &ListArtifactCompatibleServicesDefault{
		_statusCode: code,
	}
}

/*
ListArtifactCompatibleServicesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListArtifactCompatibleServicesDefault struct {
	_statusCode int

	Payload *ListArtifactCompatibleServicesDefaultBody
}

// Code gets the status code for the list artifact compatible services default response
func (o *ListArtifactCompatibleServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListArtifactCompatibleServicesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ListArtifactCompatibleServices][%d] ListArtifactCompatibleServices default  %+v", o._statusCode, o.Payload)
}
func (o *ListArtifactCompatibleServicesDefault) GetPayload() *ListArtifactCompatibleServicesDefaultBody {
	return o.Payload
}

func (o *ListArtifactCompatibleServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListArtifactCompatibleServicesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListArtifactCompatibleServicesBody list artifact compatible services body
swagger:model ListArtifactCompatibleServicesBody
*/
type ListArtifactCompatibleServicesBody struct {

	// Artifact id used to determine restore compatibility.
	ArtifactID string `json:"artifact_id,omitempty"`
}

// Validate validates this list artifact compatible services body
func (o *ListArtifactCompatibleServicesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list artifact compatible services body based on context it is used
func (o *ListArtifactCompatibleServicesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactCompatibleServicesDefaultBody list artifact compatible services default body
swagger:model ListArtifactCompatibleServicesDefaultBody
*/
type ListArtifactCompatibleServicesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListArtifactCompatibleServicesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list artifact compatible services default body
func (o *ListArtifactCompatibleServicesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactCompatibleServicesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListArtifactCompatibleServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListArtifactCompatibleServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list artifact compatible services default body based on the context it is used
func (o *ListArtifactCompatibleServicesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactCompatibleServicesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListArtifactCompatibleServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListArtifactCompatibleServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactCompatibleServicesDefaultBodyDetailsItems0 list artifact compatible services default body details items0
swagger:model ListArtifactCompatibleServicesDefaultBodyDetailsItems0
*/
type ListArtifactCompatibleServicesDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list artifact compatible services default body details items0
func (o *ListArtifactCompatibleServicesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list artifact compatible services default body details items0 based on context it is used
func (o *ListArtifactCompatibleServicesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactCompatibleServicesOKBody list artifact compatible services OK body
swagger:model ListArtifactCompatibleServicesOKBody
*/
type ListArtifactCompatibleServicesOKBody struct {

	// mysql
	Mysql []*ListArtifactCompatibleServicesOKBodyMysqlItems0 `json:"mysql"`

	// mongodb
	Mongodb []*ListArtifactCompatibleServicesOKBodyMongodbItems0 `json:"mongodb"`
}

// Validate validates this list artifact compatible services OK body
func (o *ListArtifactCompatibleServicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactCompatibleServicesOKBody) validateMysql(formats strfmt.Registry) error {
	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	for i := 0; i < len(o.Mysql); i++ {
		if swag.IsZero(o.Mysql[i]) { // not required
			continue
		}

		if o.Mysql[i] != nil {
			if err := o.Mysql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactCompatibleServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listArtifactCompatibleServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListArtifactCompatibleServicesOKBody) validateMongodb(formats strfmt.Registry) error {
	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	for i := 0; i < len(o.Mongodb); i++ {
		if swag.IsZero(o.Mongodb[i]) { // not required
			continue
		}

		if o.Mongodb[i] != nil {
			if err := o.Mongodb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactCompatibleServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listArtifactCompatibleServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list artifact compatible services OK body based on the context it is used
func (o *ListArtifactCompatibleServicesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMongodb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactCompatibleServicesOKBody) contextValidateMysql(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Mysql); i++ {

		if o.Mysql[i] != nil {
			if err := o.Mysql[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactCompatibleServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listArtifactCompatibleServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListArtifactCompatibleServicesOKBody) contextValidateMongodb(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Mongodb); i++ {

		if o.Mongodb[i] != nil {
			if err := o.Mongodb[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactCompatibleServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listArtifactCompatibleServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactCompatibleServicesOKBodyMongodbItems0 MongoDBService represents a generic MongoDB instance.
swagger:model ListArtifactCompatibleServicesOKBodyMongodbItems0
*/
type ListArtifactCompatibleServicesOKBodyMongodbItems0 struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

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

// Validate validates this list artifact compatible services OK body mongodb items0
func (o *ListArtifactCompatibleServicesOKBodyMongodbItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list artifact compatible services OK body mongodb items0 based on context it is used
func (o *ListArtifactCompatibleServicesOKBodyMongodbItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBodyMongodbItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBodyMongodbItems0) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesOKBodyMongodbItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactCompatibleServicesOKBodyMysqlItems0 MySQLService represents a generic MySQL instance.
swagger:model ListArtifactCompatibleServicesOKBodyMysqlItems0
*/
type ListArtifactCompatibleServicesOKBodyMysqlItems0 struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

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

// Validate validates this list artifact compatible services OK body mysql items0
func (o *ListArtifactCompatibleServicesOKBodyMysqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list artifact compatible services OK body mysql items0 based on context it is used
func (o *ListArtifactCompatibleServicesOKBodyMysqlItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBodyMysqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBodyMysqlItems0) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesOKBodyMysqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
