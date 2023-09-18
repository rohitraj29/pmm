// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// AddNodeReader is a Reader for the AddNode structure.
type AddNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddNodeOK creates a AddNodeOK with default headers values
func NewAddNodeOK() *AddNodeOK {
	return &AddNodeOK{}
}

/*
AddNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddNodeOK struct {
	Payload *AddNodeOKBody
}

func (o *AddNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Add][%d] addNodeOk  %+v", 200, o.Payload)
}

func (o *AddNodeOK) GetPayload() *AddNodeOKBody {
	return o.Payload
}

func (o *AddNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNodeDefault creates a AddNodeDefault with default headers values
func NewAddNodeDefault(code int) *AddNodeDefault {
	return &AddNodeDefault{
		_statusCode: code,
	}
}

/*
AddNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddNodeDefault struct {
	_statusCode int

	Payload *AddNodeDefaultBody
}

// Code gets the status code for the add node default response
func (o *AddNodeDefault) Code() int {
	return o._statusCode
}

func (o *AddNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Add][%d] AddNode default  %+v", o._statusCode, o.Payload)
}

func (o *AddNodeDefault) GetPayload() *AddNodeDefaultBody {
	return o.Payload
}

func (o *AddNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddNodeBody add node body
swagger:model AddNodeBody
*/
type AddNodeBody struct {
	// container
	Container *AddNodeParamsBodyContainer `json:"container,omitempty"`

	// generic
	Generic *AddNodeParamsBodyGeneric `json:"generic,omitempty"`

	// remote
	Remote *AddNodeParamsBodyRemote `json:"remote,omitempty"`

	// remote azure
	RemoteAzure *AddNodeParamsBodyRemoteAzure `json:"remote_azure,omitempty"`

	// remote rds
	RemoteRDS *AddNodeParamsBodyRemoteRDS `json:"remote_rds,omitempty"`
}

// Validate validates this add node body
func (o *AddNodeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteRDS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeBody) validateContainer(formats strfmt.Registry) error {
	if swag.IsZero(o.Container) { // not required
		return nil
	}

	if o.Container != nil {
		if err := o.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "container")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) validateGeneric(formats strfmt.Registry) error {
	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	if o.Generic != nil {
		if err := o.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) validateRemote(formats strfmt.Registry) error {
	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	if o.Remote != nil {
		if err := o.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) validateRemoteAzure(formats strfmt.Registry) error {
	if swag.IsZero(o.RemoteAzure) { // not required
		return nil
	}

	if o.RemoteAzure != nil {
		if err := o.RemoteAzure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "remote_azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "remote_azure")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) validateRemoteRDS(formats strfmt.Registry) error {
	if swag.IsZero(o.RemoteRDS) { // not required
		return nil
	}

	if o.RemoteRDS != nil {
		if err := o.RemoteRDS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "remote_rds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "remote_rds")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add node body based on the context it is used
func (o *AddNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateContainer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGeneric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemoteAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemoteRDS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeBody) contextValidateContainer(ctx context.Context, formats strfmt.Registry) error {
	if o.Container != nil {
		if err := o.Container.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "container")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) contextValidateGeneric(ctx context.Context, formats strfmt.Registry) error {
	if o.Generic != nil {
		if err := o.Generic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) contextValidateRemote(ctx context.Context, formats strfmt.Registry) error {
	if o.Remote != nil {
		if err := o.Remote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) contextValidateRemoteAzure(ctx context.Context, formats strfmt.Registry) error {
	if o.RemoteAzure != nil {
		if err := o.RemoteAzure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "remote_azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "remote_azure")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeBody) contextValidateRemoteRDS(ctx context.Context, formats strfmt.Registry) error {
	if o.RemoteRDS != nil {
		if err := o.RemoteRDS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "remote_rds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "remote_rds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeBody) UnmarshalBinary(b []byte) error {
	var res AddNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeDefaultBody add node default body
swagger:model AddNodeDefaultBody
*/
type AddNodeDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add node default body
func (o *AddNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add node default body based on the context it is used
func (o *AddNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeDefaultBodyDetailsItems0 add node default body details items0
swagger:model AddNodeDefaultBodyDetailsItems0
*/
type AddNodeDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add node default body details items0
func (o *AddNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node default body details items0 based on context it is used
func (o *AddNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeOKBody add node OK body
swagger:model AddNodeOKBody
*/
type AddNodeOKBody struct {
	// container
	Container *AddNodeOKBodyContainer `json:"container,omitempty"`

	// generic
	Generic *AddNodeOKBodyGeneric `json:"generic,omitempty"`

	// remote
	Remote *AddNodeOKBodyRemote `json:"remote,omitempty"`

	// remote azure database
	RemoteAzureDatabase *AddNodeOKBodyRemoteAzureDatabase `json:"remote_azure_database,omitempty"`

	// remote rds
	RemoteRDS *AddNodeOKBodyRemoteRDS `json:"remote_rds,omitempty"`
}

// Validate validates this add node OK body
func (o *AddNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteAzureDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteRDS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeOKBody) validateContainer(formats strfmt.Registry) error {
	if swag.IsZero(o.Container) { // not required
		return nil
	}

	if o.Container != nil {
		if err := o.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "container")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) validateGeneric(formats strfmt.Registry) error {
	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	if o.Generic != nil {
		if err := o.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) validateRemote(formats strfmt.Registry) error {
	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	if o.Remote != nil {
		if err := o.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) validateRemoteAzureDatabase(formats strfmt.Registry) error {
	if swag.IsZero(o.RemoteAzureDatabase) { // not required
		return nil
	}

	if o.RemoteAzureDatabase != nil {
		if err := o.RemoteAzureDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "remote_azure_database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "remote_azure_database")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) validateRemoteRDS(formats strfmt.Registry) error {
	if swag.IsZero(o.RemoteRDS) { // not required
		return nil
	}

	if o.RemoteRDS != nil {
		if err := o.RemoteRDS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "remote_rds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "remote_rds")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add node OK body based on the context it is used
func (o *AddNodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateContainer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGeneric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemoteAzureDatabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemoteRDS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeOKBody) contextValidateContainer(ctx context.Context, formats strfmt.Registry) error {
	if o.Container != nil {
		if err := o.Container.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "container")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) contextValidateGeneric(ctx context.Context, formats strfmt.Registry) error {
	if o.Generic != nil {
		if err := o.Generic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) contextValidateRemote(ctx context.Context, formats strfmt.Registry) error {
	if o.Remote != nil {
		if err := o.Remote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) contextValidateRemoteAzureDatabase(ctx context.Context, formats strfmt.Registry) error {
	if o.RemoteAzureDatabase != nil {
		if err := o.RemoteAzureDatabase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "remote_azure_database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "remote_azure_database")
			}
			return err
		}
	}

	return nil
}

func (o *AddNodeOKBody) contextValidateRemoteRDS(ctx context.Context, formats strfmt.Registry) error {
	if o.RemoteRDS != nil {
		if err := o.RemoteRDS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeOk" + "." + "remote_rds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addNodeOk" + "." + "remote_rds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeOKBody) UnmarshalBinary(b []byte) error {
	var res AddNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeOKBodyContainer ContainerNode represents a Docker container.
swagger:model AddNodeOKBodyContainer
*/
type AddNodeOKBodyContainer struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs.
	MachineID string `json:"machine_id,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node OK body container
func (o *AddNodeOKBodyContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node OK body container based on context it is used
func (o *AddNodeOKBodyContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeOKBodyContainer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeOKBodyContainer) UnmarshalBinary(b []byte) error {
	var res AddNodeOKBodyContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeOKBodyGeneric GenericNode represents a bare metal server or virtual machine.
swagger:model AddNodeOKBodyGeneric
*/
type AddNodeOKBodyGeneric struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node OK body generic
func (o *AddNodeOKBodyGeneric) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node OK body generic based on context it is used
func (o *AddNodeOKBodyGeneric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeOKBodyGeneric) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeOKBodyGeneric) UnmarshalBinary(b []byte) error {
	var res AddNodeOKBodyGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeOKBodyRemote RemoteNode represents generic remote Node. It's a node where we don't run pmm-agents. Only external exporters can run on Remote Nodes.
swagger:model AddNodeOKBodyRemote
*/
type AddNodeOKBodyRemote struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node OK body remote
func (o *AddNodeOKBodyRemote) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node OK body remote based on context it is used
func (o *AddNodeOKBodyRemote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeOKBodyRemote) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeOKBodyRemote) UnmarshalBinary(b []byte) error {
	var res AddNodeOKBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeOKBodyRemoteAzureDatabase RemoteAzureDatabaseNode represents remote AzureDatabase Node. Agents can't run on Remote AzureDatabase Nodes.
swagger:model AddNodeOKBodyRemoteAzureDatabase
*/
type AddNodeOKBodyRemoteAzureDatabase struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node OK body remote azure database
func (o *AddNodeOKBodyRemoteAzureDatabase) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node OK body remote azure database based on context it is used
func (o *AddNodeOKBodyRemoteAzureDatabase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeOKBodyRemoteAzureDatabase) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeOKBodyRemoteAzureDatabase) UnmarshalBinary(b []byte) error {
	var res AddNodeOKBodyRemoteAzureDatabase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeOKBodyRemoteRDS RemoteRDSNode represents remote RDS Node. Agents can't run on Remote RDS Nodes.
swagger:model AddNodeOKBodyRemoteRDS
*/
type AddNodeOKBodyRemoteRDS struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node OK body remote RDS
func (o *AddNodeOKBodyRemoteRDS) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node OK body remote RDS based on context it is used
func (o *AddNodeOKBodyRemoteRDS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeOKBodyRemoteRDS) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeOKBodyRemoteRDS) UnmarshalBinary(b []byte) error {
	var res AddNodeOKBodyRemoteRDS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeParamsBodyContainer add node params body container
swagger:model AddNodeParamsBodyContainer
*/
type AddNodeParamsBodyContainer struct {
	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs.
	MachineID string `json:"machine_id,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node params body container
func (o *AddNodeParamsBodyContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node params body container based on context it is used
func (o *AddNodeParamsBodyContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeParamsBodyContainer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeParamsBodyContainer) UnmarshalBinary(b []byte) error {
	var res AddNodeParamsBodyContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeParamsBodyGeneric add node params body generic
swagger:model AddNodeParamsBodyGeneric
*/
type AddNodeParamsBodyGeneric struct {
	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node params body generic
func (o *AddNodeParamsBodyGeneric) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node params body generic based on context it is used
func (o *AddNodeParamsBodyGeneric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeParamsBodyGeneric) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeParamsBodyGeneric) UnmarshalBinary(b []byte) error {
	var res AddNodeParamsBodyGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeParamsBodyRemote add node params body remote
swagger:model AddNodeParamsBodyRemote
*/
type AddNodeParamsBodyRemote struct {
	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node params body remote
func (o *AddNodeParamsBodyRemote) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node params body remote based on context it is used
func (o *AddNodeParamsBodyRemote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeParamsBodyRemote) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeParamsBodyRemote) UnmarshalBinary(b []byte) error {
	var res AddNodeParamsBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeParamsBodyRemoteAzure add node params body remote azure
swagger:model AddNodeParamsBodyRemoteAzure
*/
type AddNodeParamsBodyRemoteAzure struct {
	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node params body remote azure
func (o *AddNodeParamsBodyRemoteAzure) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node params body remote azure based on context it is used
func (o *AddNodeParamsBodyRemoteAzure) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeParamsBodyRemoteAzure) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeParamsBodyRemoteAzure) UnmarshalBinary(b []byte) error {
	var res AddNodeParamsBodyRemoteAzure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddNodeParamsBodyRemoteRDS add node params body remote RDS
swagger:model AddNodeParamsBodyRemoteRDS
*/
type AddNodeParamsBodyRemoteRDS struct {
	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add node params body remote RDS
func (o *AddNodeParamsBodyRemoteRDS) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add node params body remote RDS based on context it is used
func (o *AddNodeParamsBodyRemoteRDS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeParamsBodyRemoteRDS) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeParamsBodyRemoteRDS) UnmarshalBinary(b []byte) error {
	var res AddNodeParamsBodyRemoteRDS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
