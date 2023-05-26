// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// ScheduleBackupReader is a Reader for the ScheduleBackup structure.
type ScheduleBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleBackupOK creates a ScheduleBackupOK with default headers values
func NewScheduleBackupOK() *ScheduleBackupOK {
	return &ScheduleBackupOK{}
}

/*
ScheduleBackupOK describes a response with status code 200, with default header values.

A successful response.
*/
type ScheduleBackupOK struct {
	Payload *ScheduleBackupOKBody
}

func (o *ScheduleBackupOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Schedule][%d] scheduleBackupOk  %+v", 200, o.Payload)
}

func (o *ScheduleBackupOK) GetPayload() *ScheduleBackupOKBody {
	return o.Payload
}

func (o *ScheduleBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ScheduleBackupOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleBackupDefault creates a ScheduleBackupDefault with default headers values
func NewScheduleBackupDefault(code int) *ScheduleBackupDefault {
	return &ScheduleBackupDefault{
		_statusCode: code,
	}
}

/*
ScheduleBackupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ScheduleBackupDefault struct {
	_statusCode int

	Payload *ScheduleBackupDefaultBody
}

// Code gets the status code for the schedule backup default response
func (o *ScheduleBackupDefault) Code() int {
	return o._statusCode
}

func (o *ScheduleBackupDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Schedule][%d] ScheduleBackup default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleBackupDefault) GetPayload() *ScheduleBackupDefaultBody {
	return o.Payload
}

func (o *ScheduleBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ScheduleBackupDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ScheduleBackupBody schedule backup body
swagger:model ScheduleBackupBody
*/
type ScheduleBackupBody struct {
	// Service identifier where backup should be performed.
	ServiceID string `json:"service_id,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"location_id,omitempty"`

	// How often backup should be run in cron format.
	CronExpression string `json:"cron_expression,omitempty"`

	// First backup wouldn't happen before this time.
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// Name of backup.
	Name string `json:"name,omitempty"`

	// Human-readable description.
	Description string `json:"description,omitempty"`

	// Delay between each retry. Should have a suffix in JSON: 1s, 1m, 1h.
	RetryInterval string `json:"retry_interval,omitempty"`

	// How many times to retry a failed backup before giving up.
	Retries int64 `json:"retries,omitempty"`

	// If scheduling is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// How many artifacts keep. 0 - unlimited.
	Retention int64 `json:"retention,omitempty"`

	// BackupMode specifies backup mode.
	// Enum: [BACKUP_MODE_INVALID SNAPSHOT INCREMENTAL PITR]
	Mode *string `json:"mode,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_INVALID PHYSICAL LOGICAL]
	DataModel *string `json:"data_model,omitempty"`

	// Folder on storage for artifact.
	Folder string `json:"folder,omitempty"`
}

// Validate validates this schedule backup body
func (o *ScheduleBackupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleBackupBody) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"start_time", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var scheduleBackupBodyTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKUP_MODE_INVALID","SNAPSHOT","INCREMENTAL","PITR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleBackupBodyTypeModePropEnum = append(scheduleBackupBodyTypeModePropEnum, v)
	}
}

const (

	// ScheduleBackupBodyModeBACKUPMODEINVALID captures enum value "BACKUP_MODE_INVALID"
	ScheduleBackupBodyModeBACKUPMODEINVALID string = "BACKUP_MODE_INVALID"

	// ScheduleBackupBodyModeSNAPSHOT captures enum value "SNAPSHOT"
	ScheduleBackupBodyModeSNAPSHOT string = "SNAPSHOT"

	// ScheduleBackupBodyModeINCREMENTAL captures enum value "INCREMENTAL"
	ScheduleBackupBodyModeINCREMENTAL string = "INCREMENTAL"

	// ScheduleBackupBodyModePITR captures enum value "PITR"
	ScheduleBackupBodyModePITR string = "PITR"
)

// prop value enum
func (o *ScheduleBackupBody) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduleBackupBodyTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ScheduleBackupBody) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("body"+"."+"mode", "body", *o.Mode); err != nil {
		return err
	}

	return nil
}

var scheduleBackupBodyTypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_INVALID","PHYSICAL","LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleBackupBodyTypeDataModelPropEnum = append(scheduleBackupBodyTypeDataModelPropEnum, v)
	}
}

const (

	// ScheduleBackupBodyDataModelDATAMODELINVALID captures enum value "DATA_MODEL_INVALID"
	ScheduleBackupBodyDataModelDATAMODELINVALID string = "DATA_MODEL_INVALID"

	// ScheduleBackupBodyDataModelPHYSICAL captures enum value "PHYSICAL"
	ScheduleBackupBodyDataModelPHYSICAL string = "PHYSICAL"

	// ScheduleBackupBodyDataModelLOGICAL captures enum value "LOGICAL"
	ScheduleBackupBodyDataModelLOGICAL string = "LOGICAL"
)

// prop value enum
func (o *ScheduleBackupBody) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduleBackupBodyTypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ScheduleBackupBody) validateDataModel(formats strfmt.Registry) error {
	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("body"+"."+"data_model", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this schedule backup body based on context it is used
func (o *ScheduleBackupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ScheduleBackupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleBackupBody) UnmarshalBinary(b []byte) error {
	var res ScheduleBackupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleBackupDefaultBody schedule backup default body
swagger:model ScheduleBackupDefaultBody
*/
type ScheduleBackupDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ScheduleBackupDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this schedule backup default body
func (o *ScheduleBackupDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleBackupDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ScheduleBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ScheduleBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this schedule backup default body based on the context it is used
func (o *ScheduleBackupDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleBackupDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ScheduleBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ScheduleBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ScheduleBackupDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleBackupDefaultBody) UnmarshalBinary(b []byte) error {
	var res ScheduleBackupDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleBackupDefaultBodyDetailsItems0 schedule backup default body details items0
swagger:model ScheduleBackupDefaultBodyDetailsItems0
*/
type ScheduleBackupDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this schedule backup default body details items0
func (o *ScheduleBackupDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schedule backup default body details items0 based on context it is used
func (o *ScheduleBackupDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ScheduleBackupDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleBackupDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ScheduleBackupDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleBackupOKBody schedule backup OK body
swagger:model ScheduleBackupOKBody
*/
type ScheduleBackupOKBody struct {
	// scheduled backup id
	ScheduledBackupID string `json:"scheduled_backup_id,omitempty"`
}

// Validate validates this schedule backup OK body
func (o *ScheduleBackupOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schedule backup OK body based on context it is used
func (o *ScheduleBackupOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ScheduleBackupOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleBackupOKBody) UnmarshalBinary(b []byte) error {
	var res ScheduleBackupOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
