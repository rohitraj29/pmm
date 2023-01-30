// Code generated by go-swagger; DO NOT EDIT.

package artifacts

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
	"github.com/go-openapi/validate"
)

// ListPitrTimerangesReader is a Reader for the ListPitrTimeranges structure.
type ListPitrTimerangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPitrTimerangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPitrTimerangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPitrTimerangesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPitrTimerangesOK creates a ListPitrTimerangesOK with default headers values
func NewListPitrTimerangesOK() *ListPitrTimerangesOK {
	return &ListPitrTimerangesOK{}
}

/*
ListPitrTimerangesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListPitrTimerangesOK struct {
	Payload *ListPitrTimerangesOKBody
}

func (o *ListPitrTimerangesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Artifacts/ListPITRTimeranges][%d] listPitrTimerangesOk  %+v", 200, o.Payload)
}
func (o *ListPitrTimerangesOK) GetPayload() *ListPitrTimerangesOKBody {
	return o.Payload
}

func (o *ListPitrTimerangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListPitrTimerangesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPitrTimerangesDefault creates a ListPitrTimerangesDefault with default headers values
func NewListPitrTimerangesDefault(code int) *ListPitrTimerangesDefault {
	return &ListPitrTimerangesDefault{
		_statusCode: code,
	}
}

/*
ListPitrTimerangesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListPitrTimerangesDefault struct {
	_statusCode int

	Payload *ListPitrTimerangesDefaultBody
}

// Code gets the status code for the list pitr timeranges default response
func (o *ListPitrTimerangesDefault) Code() int {
	return o._statusCode
}

func (o *ListPitrTimerangesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Artifacts/ListPITRTimeranges][%d] ListPitrTimeranges default  %+v", o._statusCode, o.Payload)
}
func (o *ListPitrTimerangesDefault) GetPayload() *ListPitrTimerangesDefaultBody {
	return o.Payload
}

func (o *ListPitrTimerangesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListPitrTimerangesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListPitrTimerangesBody list pitr timeranges body
swagger:model ListPitrTimerangesBody
*/
type ListPitrTimerangesBody struct {

	// Artifact ID represents artifact whose location has PITR timeranges to be retrieved.
	ArtifactID string `json:"artifact_id,omitempty"`
}

// Validate validates this list pitr timeranges body
func (o *ListPitrTimerangesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list pitr timeranges body based on context it is used
func (o *ListPitrTimerangesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPitrTimerangesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPitrTimerangesBody) UnmarshalBinary(b []byte) error {
	var res ListPitrTimerangesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListPitrTimerangesDefaultBody list pitr timeranges default body
swagger:model ListPitrTimerangesDefaultBody
*/
type ListPitrTimerangesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListPitrTimerangesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list pitr timeranges default body
func (o *ListPitrTimerangesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPitrTimerangesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListPitrTimeranges default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListPitrTimeranges default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list pitr timeranges default body based on the context it is used
func (o *ListPitrTimerangesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPitrTimerangesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListPitrTimeranges default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListPitrTimeranges default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPitrTimerangesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPitrTimerangesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListPitrTimerangesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListPitrTimerangesDefaultBodyDetailsItems0 list pitr timeranges default body details items0
swagger:model ListPitrTimerangesDefaultBodyDetailsItems0
*/
type ListPitrTimerangesDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list pitr timeranges default body details items0
func (o *ListPitrTimerangesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list pitr timeranges default body details items0 based on context it is used
func (o *ListPitrTimerangesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPitrTimerangesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPitrTimerangesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListPitrTimerangesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListPitrTimerangesOKBody list pitr timeranges OK body
swagger:model ListPitrTimerangesOKBody
*/
type ListPitrTimerangesOKBody struct {

	// timeranges
	Timeranges []*ListPitrTimerangesOKBodyTimerangesItems0 `json:"timeranges"`
}

// Validate validates this list pitr timeranges OK body
func (o *ListPitrTimerangesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimeranges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPitrTimerangesOKBody) validateTimeranges(formats strfmt.Registry) error {
	if swag.IsZero(o.Timeranges) { // not required
		return nil
	}

	for i := 0; i < len(o.Timeranges); i++ {
		if swag.IsZero(o.Timeranges[i]) { // not required
			continue
		}

		if o.Timeranges[i] != nil {
			if err := o.Timeranges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listPitrTimerangesOk" + "." + "timeranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listPitrTimerangesOk" + "." + "timeranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list pitr timeranges OK body based on the context it is used
func (o *ListPitrTimerangesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTimeranges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPitrTimerangesOKBody) contextValidateTimeranges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Timeranges); i++ {

		if o.Timeranges[i] != nil {
			if err := o.Timeranges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listPitrTimerangesOk" + "." + "timeranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listPitrTimerangesOk" + "." + "timeranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPitrTimerangesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPitrTimerangesOKBody) UnmarshalBinary(b []byte) error {
	var res ListPitrTimerangesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListPitrTimerangesOKBodyTimerangesItems0 list pitr timeranges OK body timeranges items0
swagger:model ListPitrTimerangesOKBodyTimerangesItems0
*/
type ListPitrTimerangesOKBodyTimerangesItems0 struct {

	// start_timestamp is the time of the first event in the PITR chunk.
	// Format: date-time
	StartTimestamp strfmt.DateTime `json:"start_timestamp,omitempty"`

	// end_timestamp is the time of the last event in the PITR chunk.
	// Format: date-time
	EndTimestamp strfmt.DateTime `json:"end_timestamp,omitempty"`
}

// Validate validates this list pitr timeranges OK body timeranges items0
func (o *ListPitrTimerangesOKBodyTimerangesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStartTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPitrTimerangesOKBodyTimerangesItems0) validateStartTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("start_timestamp", "body", "date-time", o.StartTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListPitrTimerangesOKBodyTimerangesItems0) validateEndTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.EndTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("end_timestamp", "body", "date-time", o.EndTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list pitr timeranges OK body timeranges items0 based on context it is used
func (o *ListPitrTimerangesOKBodyTimerangesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPitrTimerangesOKBodyTimerangesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPitrTimerangesOKBodyTimerangesItems0) UnmarshalBinary(b []byte) error {
	var res ListPitrTimerangesOKBodyTimerangesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
