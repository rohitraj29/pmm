// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// GetResourcesReader is a Reader for the GetResources structure.
type GetResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResourcesOK creates a GetResourcesOK with default headers values
func NewGetResourcesOK() *GetResourcesOK {
	return &GetResourcesOK{}
}

/* GetResourcesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetResourcesOK struct {
	Payload *GetResourcesOKBody
}

func (o *GetResourcesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/Resources/Get][%d] getResourcesOk  %+v", 200, o.Payload)
}
func (o *GetResourcesOK) GetPayload() *GetResourcesOKBody {
	return o.Payload
}

func (o *GetResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesDefault creates a GetResourcesDefault with default headers values
func NewGetResourcesDefault(code int) *GetResourcesDefault {
	return &GetResourcesDefault{
		_statusCode: code,
	}
}

/* GetResourcesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetResourcesDefault struct {
	_statusCode int

	Payload *GetResourcesDefaultBody
}

// Code gets the status code for the get resources default response
func (o *GetResourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetResourcesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/Resources/Get][%d] GetResources default  %+v", o._statusCode, o.Payload)
}
func (o *GetResourcesDefault) GetPayload() *GetResourcesDefaultBody {
	return o.Payload
}

func (o *GetResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourcesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetResourcesBody get resources body
swagger:model GetResourcesBody
*/
type GetResourcesBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`
}

// Validate validates this get resources body
func (o *GetResourcesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get resources body based on context it is used
func (o *GetResourcesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesBody) UnmarshalBinary(b []byte) error {
	var res GetResourcesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourcesDefaultBody get resources default body
swagger:model GetResourcesDefaultBody
*/
type GetResourcesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetResourcesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get resources default body
func (o *GetResourcesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourcesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetResources default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetResources default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get resources default body based on the context it is used
func (o *GetResourcesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourcesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetResources default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetResources default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetResourcesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourcesDefaultBodyDetailsItems0 get resources default body details items0
swagger:model GetResourcesDefaultBodyDetailsItems0
*/
type GetResourcesDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get resources default body details items0
func (o *GetResourcesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get resources default body details items0 based on context it is used
func (o *GetResourcesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetResourcesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourcesOKBody get resources OK body
swagger:model GetResourcesOKBody
*/
type GetResourcesOKBody struct {

	// all
	All *GetResourcesOKBodyAll `json:"all,omitempty"`

	// available
	Available *GetResourcesOKBodyAvailable `json:"available,omitempty"`
}

// Validate validates this get resources OK body
func (o *GetResourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAll(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourcesOKBody) validateAll(formats strfmt.Registry) error {
	if swag.IsZero(o.All) { // not required
		return nil
	}

	if o.All != nil {
		if err := o.All.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResourcesOk" + "." + "all")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getResourcesOk" + "." + "all")
			}
			return err
		}
	}

	return nil
}

func (o *GetResourcesOKBody) validateAvailable(formats strfmt.Registry) error {
	if swag.IsZero(o.Available) { // not required
		return nil
	}

	if o.Available != nil {
		if err := o.Available.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResourcesOk" + "." + "available")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getResourcesOk" + "." + "available")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get resources OK body based on the context it is used
func (o *GetResourcesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAll(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAvailable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourcesOKBody) contextValidateAll(ctx context.Context, formats strfmt.Registry) error {

	if o.All != nil {
		if err := o.All.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResourcesOk" + "." + "all")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getResourcesOk" + "." + "all")
			}
			return err
		}
	}

	return nil
}

func (o *GetResourcesOKBody) contextValidateAvailable(ctx context.Context, formats strfmt.Registry) error {

	if o.Available != nil {
		if err := o.Available.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResourcesOk" + "." + "available")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getResourcesOk" + "." + "available")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourcesOKBodyAll Resources contains Kubernetes cluster resources.
swagger:model GetResourcesOKBodyAll
*/
type GetResourcesOKBodyAll struct {

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`

	// CPU in millicpus. For example 0.1 of CPU is equivalent to 100 millicpus.
	// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.
	CPUm string `json:"cpu_m,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`
}

// Validate validates this get resources OK body all
func (o *GetResourcesOKBodyAll) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get resources OK body all based on context it is used
func (o *GetResourcesOKBodyAll) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesOKBodyAll) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesOKBodyAll) UnmarshalBinary(b []byte) error {
	var res GetResourcesOKBodyAll
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourcesOKBodyAvailable Resources contains Kubernetes cluster resources.
swagger:model GetResourcesOKBodyAvailable
*/
type GetResourcesOKBodyAvailable struct {

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`

	// CPU in millicpus. For example 0.1 of CPU is equivalent to 100 millicpus.
	// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.
	CPUm string `json:"cpu_m,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`
}

// Validate validates this get resources OK body available
func (o *GetResourcesOKBodyAvailable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get resources OK body available based on context it is used
func (o *GetResourcesOKBodyAvailable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesOKBodyAvailable) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesOKBodyAvailable) UnmarshalBinary(b []byte) error {
	var res GetResourcesOKBodyAvailable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
