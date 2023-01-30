// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// ListKubernetesClustersReader is a Reader for the ListKubernetesClusters structure.
type ListKubernetesClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubernetesClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubernetesClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubernetesClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubernetesClustersOK creates a ListKubernetesClustersOK with default headers values
func NewListKubernetesClustersOK() *ListKubernetesClustersOK {
	return &ListKubernetesClustersOK{}
}

/*
ListKubernetesClustersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListKubernetesClustersOK struct {
	Payload *ListKubernetesClustersOKBody
}

func (o *ListKubernetesClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/List][%d] listKubernetesClustersOk  %+v", 200, o.Payload)
}
func (o *ListKubernetesClustersOK) GetPayload() *ListKubernetesClustersOKBody {
	return o.Payload
}

func (o *ListKubernetesClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListKubernetesClustersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubernetesClustersDefault creates a ListKubernetesClustersDefault with default headers values
func NewListKubernetesClustersDefault(code int) *ListKubernetesClustersDefault {
	return &ListKubernetesClustersDefault{
		_statusCode: code,
	}
}

/*
ListKubernetesClustersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListKubernetesClustersDefault struct {
	_statusCode int

	Payload *ListKubernetesClustersDefaultBody
}

// Code gets the status code for the list kubernetes clusters default response
func (o *ListKubernetesClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListKubernetesClustersDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/List][%d] ListKubernetesClusters default  %+v", o._statusCode, o.Payload)
}
func (o *ListKubernetesClustersDefault) GetPayload() *ListKubernetesClustersDefaultBody {
	return o.Payload
}

func (o *ListKubernetesClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListKubernetesClustersDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListKubernetesClustersDefaultBody list kubernetes clusters default body
swagger:model ListKubernetesClustersDefaultBody
*/
type ListKubernetesClustersDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListKubernetesClustersDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list kubernetes clusters default body
func (o *ListKubernetesClustersDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListKubernetesClusters default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListKubernetesClusters default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list kubernetes clusters default body based on the context it is used
func (o *ListKubernetesClustersDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListKubernetesClusters default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListKubernetesClusters default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListKubernetesClustersDefaultBodyDetailsItems0 list kubernetes clusters default body details items0
swagger:model ListKubernetesClustersDefaultBodyDetailsItems0
*/
type ListKubernetesClustersDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list kubernetes clusters default body details items0
func (o *ListKubernetesClustersDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list kubernetes clusters default body details items0 based on context it is used
func (o *ListKubernetesClustersDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListKubernetesClustersOKBody list kubernetes clusters OK body
swagger:model ListKubernetesClustersOKBody
*/
type ListKubernetesClustersOKBody struct {

	// Kubernetes clusters.
	KubernetesClusters []*ListKubernetesClustersOKBodyKubernetesClustersItems0 `json:"kubernetes_clusters"`
}

// Validate validates this list kubernetes clusters OK body
func (o *ListKubernetesClustersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKubernetesClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersOKBody) validateKubernetesClusters(formats strfmt.Registry) error {
	if swag.IsZero(o.KubernetesClusters) { // not required
		return nil
	}

	for i := 0; i < len(o.KubernetesClusters); i++ {
		if swag.IsZero(o.KubernetesClusters[i]) { // not required
			continue
		}

		if o.KubernetesClusters[i] != nil {
			if err := o.KubernetesClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listKubernetesClustersOk" + "." + "kubernetes_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listKubernetesClustersOk" + "." + "kubernetes_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list kubernetes clusters OK body based on the context it is used
func (o *ListKubernetesClustersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateKubernetesClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersOKBody) contextValidateKubernetesClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.KubernetesClusters); i++ {

		if o.KubernetesClusters[i] != nil {
			if err := o.KubernetesClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listKubernetesClustersOk" + "." + "kubernetes_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listKubernetesClustersOk" + "." + "kubernetes_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBody) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListKubernetesClustersOKBodyKubernetesClustersItems0 Cluster contains public info about Kubernetes cluster.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model ListKubernetesClustersOKBodyKubernetesClustersItems0
*/
type ListKubernetesClustersOKBodyKubernetesClustersItems0 struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// KubernetesClusterStatus defines status of Kubernetes cluster.
	//
	//  - KUBERNETES_CLUSTER_STATUS_INVALID: KUBERNETES_CLUSTER_STATUS_INVALID represents unknown state.
	//  - KUBERNETES_CLUSTER_STATUS_OK: KUBERNETES_CLUSTER_STATUS_OK represents that Kubernetes cluster is accessible.
	//  - KUBERNETES_CLUSTER_STATUS_UNAVAILABLE: KUBERNETES_CLUSTER_STATUS_UNAVAILABLE represents that Kubernetes cluster is not accessible.
	//  - KUBERNETES_CLUSTER_STATUS_PROVISIONING: KUBERNETES_CLUSTER_STATUS_PROVISIONING represents that Kubernetes cluster is privisioning.
	// Enum: [KUBERNETES_CLUSTER_STATUS_INVALID KUBERNETES_CLUSTER_STATUS_OK KUBERNETES_CLUSTER_STATUS_UNAVAILABLE KUBERNETES_CLUSTER_STATUS_PROVISIONING]
	Status *string `json:"status,omitempty"`

	// operators
	Operators *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators `json:"operators,omitempty"`
}

// Validate validates this list kubernetes clusters OK body kubernetes clusters items0
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listKubernetesClustersOkBodyKubernetesClustersItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["KUBERNETES_CLUSTER_STATUS_INVALID","KUBERNETES_CLUSTER_STATUS_OK","KUBERNETES_CLUSTER_STATUS_UNAVAILABLE","KUBERNETES_CLUSTER_STATUS_PROVISIONING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listKubernetesClustersOkBodyKubernetesClustersItems0TypeStatusPropEnum = append(listKubernetesClustersOkBodyKubernetesClustersItems0TypeStatusPropEnum, v)
	}
}

const (

	// ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSINVALID captures enum value "KUBERNETES_CLUSTER_STATUS_INVALID"
	ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSINVALID string = "KUBERNETES_CLUSTER_STATUS_INVALID"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSOK captures enum value "KUBERNETES_CLUSTER_STATUS_OK"
	ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSOK string = "KUBERNETES_CLUSTER_STATUS_OK"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSUNAVAILABLE captures enum value "KUBERNETES_CLUSTER_STATUS_UNAVAILABLE"
	ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSUNAVAILABLE string = "KUBERNETES_CLUSTER_STATUS_UNAVAILABLE"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSPROVISIONING captures enum value "KUBERNETES_CLUSTER_STATUS_PROVISIONING"
	ListKubernetesClustersOKBodyKubernetesClustersItems0StatusKUBERNETESCLUSTERSTATUSPROVISIONING string = "KUBERNETES_CLUSTER_STATUS_PROVISIONING"
)

// prop value enum
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listKubernetesClustersOkBodyKubernetesClustersItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) validateOperators(formats strfmt.Registry) error {
	if swag.IsZero(o.Operators) { // not required
		return nil
	}

	if o.Operators != nil {
		if err := o.Operators.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list kubernetes clusters OK body kubernetes clusters items0 based on the context it is used
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOperators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) contextValidateOperators(ctx context.Context, formats strfmt.Registry) error {

	if o.Operators != nil {
		if err := o.Operators.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBodyKubernetesClustersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListKubernetesClustersOKBodyKubernetesClustersItems0Operators Operators contains list of operators installed in Kubernetes cluster.
swagger:model ListKubernetesClustersOKBodyKubernetesClustersItems0Operators
*/
type ListKubernetesClustersOKBodyKubernetesClustersItems0Operators struct {

	// dbaas
	Dbaas *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas `json:"dbaas,omitempty"`

	// psmdb
	PSMDB *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB `json:"psmdb,omitempty"`

	// pxc
	PXC *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC `json:"pxc,omitempty"`
}

// Validate validates this list kubernetes clusters OK body kubernetes clusters items0 operators
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDbaas(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePSMDB(formats); err != nil {
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

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) validateDbaas(formats strfmt.Registry) error {
	if swag.IsZero(o.Dbaas) { // not required
		return nil
	}

	if o.Dbaas != nil {
		if err := o.Dbaas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "dbaas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators" + "." + "dbaas")
			}
			return err
		}
	}

	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) validatePSMDB(formats strfmt.Registry) error {
	if swag.IsZero(o.PSMDB) { // not required
		return nil
	}

	if o.PSMDB != nil {
		if err := o.PSMDB.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "psmdb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators" + "." + "psmdb")
			}
			return err
		}
	}

	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) validatePXC(formats strfmt.Registry) error {
	if swag.IsZero(o.PXC) { // not required
		return nil
	}

	if o.PXC != nil {
		if err := o.PXC.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "pxc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list kubernetes clusters OK body kubernetes clusters items0 operators based on the context it is used
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDbaas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePSMDB(ctx, formats); err != nil {
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

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) contextValidateDbaas(ctx context.Context, formats strfmt.Registry) error {

	if o.Dbaas != nil {
		if err := o.Dbaas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "dbaas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators" + "." + "dbaas")
			}
			return err
		}
	}

	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) contextValidatePSMDB(ctx context.Context, formats strfmt.Registry) error {

	if o.PSMDB != nil {
		if err := o.PSMDB.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "psmdb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators" + "." + "psmdb")
			}
			return err
		}
	}

	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) contextValidatePXC(ctx context.Context, formats strfmt.Registry) error {

	if o.PXC != nil {
		if err := o.PXC.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "pxc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operators" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0Operators) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBodyKubernetesClustersItems0Operators
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas Operator contains all information about operator installed in Kubernetes cluster.
swagger:model ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas
*/
type ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas struct {

	// OperatorsStatus defines status of operators installed in Kubernetes cluster.
	//
	//  - OPERATORS_STATUS_INVALID: OPERATORS_STATUS_INVALID represents unknown state.
	//  - OPERATORS_STATUS_OK: OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	//  - OPERATORS_STATUS_UNSUPPORTED: OPERATORS_STATUS_UNSUPPORTED represents that operators are installed, but doesn't have supported API version.
	//  - OPERATORS_STATUS_NOT_INSTALLED: OPERATORS_STATUS_NOT_INSTALLED represents that operators are not installed.
	// Enum: [OPERATORS_STATUS_INVALID OPERATORS_STATUS_OK OPERATORS_STATUS_UNSUPPORTED OPERATORS_STATUS_NOT_INSTALLED]
	Status *string `json:"status,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this list kubernetes clusters OK body kubernetes clusters items0 operators dbaas
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsDbaasTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPERATORS_STATUS_INVALID","OPERATORS_STATUS_OK","OPERATORS_STATUS_UNSUPPORTED","OPERATORS_STATUS_NOT_INSTALLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsDbaasTypeStatusPropEnum = append(listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsDbaasTypeStatusPropEnum, v)
	}
}

const (

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSINVALID captures enum value "OPERATORS_STATUS_INVALID"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSINVALID string = "OPERATORS_STATUS_INVALID"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSOK captures enum value "OPERATORS_STATUS_OK"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSOK string = "OPERATORS_STATUS_OK"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSUNSUPPORTED captures enum value "OPERATORS_STATUS_UNSUPPORTED"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSUNSUPPORTED string = "OPERATORS_STATUS_UNSUPPORTED"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSNOTINSTALLED captures enum value "OPERATORS_STATUS_NOT_INSTALLED"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaasStatusOPERATORSSTATUSNOTINSTALLED string = "OPERATORS_STATUS_NOT_INSTALLED"
)

// prop value enum
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsDbaasTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("operators"+"."+"dbaas"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list kubernetes clusters OK body kubernetes clusters items0 operators dbaas based on context it is used
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsDbaas
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB Operator contains all information about operator installed in Kubernetes cluster.
swagger:model ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB
*/
type ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB struct {

	// OperatorsStatus defines status of operators installed in Kubernetes cluster.
	//
	//  - OPERATORS_STATUS_INVALID: OPERATORS_STATUS_INVALID represents unknown state.
	//  - OPERATORS_STATUS_OK: OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	//  - OPERATORS_STATUS_UNSUPPORTED: OPERATORS_STATUS_UNSUPPORTED represents that operators are installed, but doesn't have supported API version.
	//  - OPERATORS_STATUS_NOT_INSTALLED: OPERATORS_STATUS_NOT_INSTALLED represents that operators are not installed.
	// Enum: [OPERATORS_STATUS_INVALID OPERATORS_STATUS_OK OPERATORS_STATUS_UNSUPPORTED OPERATORS_STATUS_NOT_INSTALLED]
	Status *string `json:"status,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this list kubernetes clusters OK body kubernetes clusters items0 operators PSMDB
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPERATORS_STATUS_INVALID","OPERATORS_STATUS_OK","OPERATORS_STATUS_UNSUPPORTED","OPERATORS_STATUS_NOT_INSTALLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum = append(listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum, v)
	}
}

const (

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSINVALID captures enum value "OPERATORS_STATUS_INVALID"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSINVALID string = "OPERATORS_STATUS_INVALID"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSOK captures enum value "OPERATORS_STATUS_OK"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSOK string = "OPERATORS_STATUS_OK"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSUNSUPPORTED captures enum value "OPERATORS_STATUS_UNSUPPORTED"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSUNSUPPORTED string = "OPERATORS_STATUS_UNSUPPORTED"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSNOTINSTALLED captures enum value "OPERATORS_STATUS_NOT_INSTALLED"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSNOTINSTALLED string = "OPERATORS_STATUS_NOT_INSTALLED"
)

// prop value enum
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("operators"+"."+"psmdb"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list kubernetes clusters OK body kubernetes clusters items0 operators PSMDB based on context it is used
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPSMDB
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC Operator contains all information about operator installed in Kubernetes cluster.
swagger:model ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC
*/
type ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC struct {

	// OperatorsStatus defines status of operators installed in Kubernetes cluster.
	//
	//  - OPERATORS_STATUS_INVALID: OPERATORS_STATUS_INVALID represents unknown state.
	//  - OPERATORS_STATUS_OK: OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	//  - OPERATORS_STATUS_UNSUPPORTED: OPERATORS_STATUS_UNSUPPORTED represents that operators are installed, but doesn't have supported API version.
	//  - OPERATORS_STATUS_NOT_INSTALLED: OPERATORS_STATUS_NOT_INSTALLED represents that operators are not installed.
	// Enum: [OPERATORS_STATUS_INVALID OPERATORS_STATUS_OK OPERATORS_STATUS_UNSUPPORTED OPERATORS_STATUS_NOT_INSTALLED]
	Status *string `json:"status,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this list kubernetes clusters OK body kubernetes clusters items0 operators PXC
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPxcTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPERATORS_STATUS_INVALID","OPERATORS_STATUS_OK","OPERATORS_STATUS_UNSUPPORTED","OPERATORS_STATUS_NOT_INSTALLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPxcTypeStatusPropEnum = append(listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPxcTypeStatusPropEnum, v)
	}
}

const (

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSINVALID captures enum value "OPERATORS_STATUS_INVALID"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSINVALID string = "OPERATORS_STATUS_INVALID"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSOK captures enum value "OPERATORS_STATUS_OK"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSOK string = "OPERATORS_STATUS_OK"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSUNSUPPORTED captures enum value "OPERATORS_STATUS_UNSUPPORTED"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSUNSUPPORTED string = "OPERATORS_STATUS_UNSUPPORTED"

	// ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSNOTINSTALLED captures enum value "OPERATORS_STATUS_NOT_INSTALLED"
	ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXCStatusOPERATORSSTATUSNOTINSTALLED string = "OPERATORS_STATUS_NOT_INSTALLED"
)

// prop value enum
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listKubernetesClustersOkBodyKubernetesClustersItems0OperatorsPxcTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("operators"+"."+"pxc"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list kubernetes clusters OK body kubernetes clusters items0 operators PXC based on context it is used
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBodyKubernetesClustersItems0OperatorsPXC
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
