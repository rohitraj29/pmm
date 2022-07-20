// Code generated by go-swagger; DO NOT EDIT.

package pxc_clusters

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

// UpdatePXCClusterReader is a Reader for the UpdatePXCCluster structure.
type UpdatePXCClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePXCClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePXCClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdatePXCClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePXCClusterOK creates a UpdatePXCClusterOK with default headers values
func NewUpdatePXCClusterOK() *UpdatePXCClusterOK {
	return &UpdatePXCClusterOK{}
}

/* UpdatePXCClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdatePXCClusterOK struct {
	Payload interface{}
}

func (o *UpdatePXCClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCCluster/Update][%d] updatePxcClusterOk  %+v", 200, o.Payload)
}
func (o *UpdatePXCClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdatePXCClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePXCClusterDefault creates a UpdatePXCClusterDefault with default headers values
func NewUpdatePXCClusterDefault(code int) *UpdatePXCClusterDefault {
	return &UpdatePXCClusterDefault{
		_statusCode: code,
	}
}

/* UpdatePXCClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdatePXCClusterDefault struct {
	_statusCode int

	Payload *UpdatePXCClusterDefaultBody
}

// Code gets the status code for the update PXC cluster default response
func (o *UpdatePXCClusterDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePXCClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCCluster/Update][%d] UpdatePXCCluster default  %+v", o._statusCode, o.Payload)
}
func (o *UpdatePXCClusterDefault) GetPayload() *UpdatePXCClusterDefaultBody {
	return o.Payload
}

func (o *UpdatePXCClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdatePXCClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdatePXCClusterBody update PXC cluster body
swagger:model UpdatePXCClusterBody
*/
type UpdatePXCClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// PXC cluster name.
	Name string `json:"name,omitempty"`

	// params
	Params *UpdatePXCClusterParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this update PXC cluster body
func (o *UpdatePXCClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update PXC cluster body based on the context it is used
func (o *UpdatePXCClusterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	if o.Params != nil {
		if err := o.Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterBody) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterDefaultBody update PXC cluster default body
swagger:model UpdatePXCClusterDefaultBody
*/
type UpdatePXCClusterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UpdatePXCClusterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this update PXC cluster default body
func (o *UpdatePXCClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UpdatePXCCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdatePXCCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update PXC cluster default body based on the context it is used
func (o *UpdatePXCClusterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdatePXCCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdatePXCCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterDefaultBodyDetailsItems0 update PXC cluster default body details items0
swagger:model UpdatePXCClusterDefaultBodyDetailsItems0
*/
type UpdatePXCClusterDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this update PXC cluster default body details items0
func (o *UpdatePXCClusterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update PXC cluster default body details items0 based on context it is used
func (o *UpdatePXCClusterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterParamsBodyParams UpdatePXCClusterParams represents PXC cluster parameters that can be updated.
swagger:model UpdatePXCClusterParamsBodyParams
*/
type UpdatePXCClusterParamsBodyParams struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// Suspend cluster `pause: true`.
	Suspend bool `json:"suspend,omitempty"`

	// Resume cluster `pause: false`.
	Resume bool `json:"resume,omitempty"`

	// haproxy
	Haproxy *UpdatePXCClusterParamsBodyParamsHaproxy `json:"haproxy,omitempty"`

	// proxysql
	Proxysql *UpdatePXCClusterParamsBodyParamsProxysql `json:"proxysql,omitempty"`

	// pxc
	PXC *UpdatePXCClusterParamsBodyParamsPXC `json:"pxc,omitempty"`
}

// Validate validates this update PXC cluster params body params
func (o *UpdatePXCClusterParamsBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePXC(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParams) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	if o.Haproxy != nil {
		if err := o.Haproxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

func (o *UpdatePXCClusterParamsBodyParams) validateProxysql(formats strfmt.Registry) error {
	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *UpdatePXCClusterParamsBodyParams) validatePXC(formats strfmt.Registry) error {
	if swag.IsZero(o.PXC) { // not required
		return nil
	}

	if o.PXC != nil {
		if err := o.PXC.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update PXC cluster params body params based on the context it is used
func (o *UpdatePXCClusterParamsBodyParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHaproxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateProxysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePXC(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParams) contextValidateHaproxy(ctx context.Context, formats strfmt.Registry) error {

	if o.Haproxy != nil {
		if err := o.Haproxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

func (o *UpdatePXCClusterParamsBodyParams) contextValidateProxysql(ctx context.Context, formats strfmt.Registry) error {

	if o.Proxysql != nil {
		if err := o.Proxysql.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *UpdatePXCClusterParamsBodyParams) contextValidatePXC(ctx context.Context, formats strfmt.Registry) error {

	if o.PXC != nil {
		if err := o.PXC.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterParamsBodyParamsHaproxy HAProxy container parameters.
swagger:model UpdatePXCClusterParamsBodyParamsHaproxy
*/
type UpdatePXCClusterParamsBodyParamsHaproxy struct {

	// compute resources
	ComputeResources *UpdatePXCClusterParamsBodyParamsHaproxyComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this update PXC cluster params body params haproxy
func (o *UpdatePXCClusterParamsBodyParamsHaproxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParamsHaproxy) validateComputeResources(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update PXC cluster params body params haproxy based on the context it is used
func (o *UpdatePXCClusterParamsBodyParamsHaproxy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParamsHaproxy) contextValidateComputeResources(ctx context.Context, formats strfmt.Registry) error {

	if o.ComputeResources != nil {
		if err := o.ComputeResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsHaproxy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsHaproxy) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterParamsBodyParamsHaproxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterParamsBodyParamsHaproxyComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model UpdatePXCClusterParamsBodyParamsHaproxyComputeResources
*/
type UpdatePXCClusterParamsBodyParamsHaproxyComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this update PXC cluster params body params haproxy compute resources
func (o *UpdatePXCClusterParamsBodyParamsHaproxyComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update PXC cluster params body params haproxy compute resources based on context it is used
func (o *UpdatePXCClusterParamsBodyParamsHaproxyComputeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsHaproxyComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsHaproxyComputeResources) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterParamsBodyParamsHaproxyComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterParamsBodyParamsPXC PXC container parameters.
swagger:model UpdatePXCClusterParamsBodyParamsPXC
*/
type UpdatePXCClusterParamsBodyParamsPXC struct {

	// Image to use. If it's the same image but with different version tag, upgrade of database cluster to version
	// in given tag is triggered. If entirely different image is given, error is returned.
	Image string `json:"image,omitempty"`

	// compute resources
	ComputeResources *UpdatePXCClusterParamsBodyParamsPXCComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this update PXC cluster params body params PXC
func (o *UpdatePXCClusterParamsBodyParamsPXC) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParamsPXC) validateComputeResources(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update PXC cluster params body params PXC based on the context it is used
func (o *UpdatePXCClusterParamsBodyParamsPXC) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParamsPXC) contextValidateComputeResources(ctx context.Context, formats strfmt.Registry) error {

	if o.ComputeResources != nil {
		if err := o.ComputeResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsPXC) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsPXC) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterParamsBodyParamsPXC
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterParamsBodyParamsPXCComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model UpdatePXCClusterParamsBodyParamsPXCComputeResources
*/
type UpdatePXCClusterParamsBodyParamsPXCComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this update PXC cluster params body params PXC compute resources
func (o *UpdatePXCClusterParamsBodyParamsPXCComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update PXC cluster params body params PXC compute resources based on context it is used
func (o *UpdatePXCClusterParamsBodyParamsPXCComputeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsPXCComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsPXCComputeResources) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterParamsBodyParamsPXCComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterParamsBodyParamsProxysql ProxySQL container parameters.
swagger:model UpdatePXCClusterParamsBodyParamsProxysql
*/
type UpdatePXCClusterParamsBodyParamsProxysql struct {

	// compute resources
	ComputeResources *UpdatePXCClusterParamsBodyParamsProxysqlComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this update PXC cluster params body params proxysql
func (o *UpdatePXCClusterParamsBodyParamsProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParamsProxysql) validateComputeResources(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update PXC cluster params body params proxysql based on the context it is used
func (o *UpdatePXCClusterParamsBodyParamsProxysql) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePXCClusterParamsBodyParamsProxysql) contextValidateComputeResources(ctx context.Context, formats strfmt.Registry) error {

	if o.ComputeResources != nil {
		if err := o.ComputeResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsProxysql) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterParamsBodyParamsProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePXCClusterParamsBodyParamsProxysqlComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model UpdatePXCClusterParamsBodyParamsProxysqlComputeResources
*/
type UpdatePXCClusterParamsBodyParamsProxysqlComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this update PXC cluster params body params proxysql compute resources
func (o *UpdatePXCClusterParamsBodyParamsProxysqlComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update PXC cluster params body params proxysql compute resources based on context it is used
func (o *UpdatePXCClusterParamsBodyParamsProxysqlComputeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsProxysqlComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePXCClusterParamsBodyParamsProxysqlComputeResources) UnmarshalBinary(b []byte) error {
	var res UpdatePXCClusterParamsBodyParamsProxysqlComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
