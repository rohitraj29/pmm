// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// AddMongoDBExporterReader is a Reader for the AddMongoDBExporter structure.
type AddMongoDBExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMongoDBExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMongoDBExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMongoDBExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMongoDBExporterOK creates a AddMongoDBExporterOK with default headers values
func NewAddMongoDBExporterOK() *AddMongoDBExporterOK {
	return &AddMongoDBExporterOK{}
}

/*
AddMongoDBExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddMongoDBExporterOK struct {
	Payload *AddMongoDBExporterOKBody
}

func (o *AddMongoDBExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMongoDBExporter][%d] addMongoDbExporterOk  %+v", 200, o.Payload)
}
func (o *AddMongoDBExporterOK) GetPayload() *AddMongoDBExporterOKBody {
	return o.Payload
}

func (o *AddMongoDBExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMongoDBExporterDefault creates a AddMongoDBExporterDefault with default headers values
func NewAddMongoDBExporterDefault(code int) *AddMongoDBExporterDefault {
	return &AddMongoDBExporterDefault{
		_statusCode: code,
	}
}

/*
AddMongoDBExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddMongoDBExporterDefault struct {
	_statusCode int

	Payload *AddMongoDBExporterDefaultBody
}

// Code gets the status code for the add mongo DB exporter default response
func (o *AddMongoDBExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddMongoDBExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMongoDBExporter][%d] AddMongoDBExporter default  %+v", o._statusCode, o.Payload)
}
func (o *AddMongoDBExporterDefault) GetPayload() *AddMongoDBExporterDefaultBody {
	return o.Payload
}

func (o *AddMongoDBExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddMongoDBExporterBody add mongo DB exporter body
swagger:model AddMongoDBExporterBody
*/
type AddMongoDBExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`

	// MongoDB password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Client certificate and key.
	TLSCertificateKey string `json:"tls_certificate_key,omitempty"`

	// Password for decrypting tls_certificate_key.
	TLSCertificateKeyFilePassword string `json:"tls_certificate_key_file_password,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// Authentication mechanism.
	// See https://docs.mongodb.com/manual/reference/connection-string/#mongodb-urioption-urioption.authMechanism
	// for details.
	AuthenticationMechanism string `json:"authentication_mechanism,omitempty"`

	// Authentication database.
	AuthenticationDatabase string `json:"authentication_database,omitempty"`

	// Custom password for exporter endpoint /metrics.
	AgentPassword string `json:"agent_password,omitempty"`

	// List of colletions to get stats from. Can use *
	StatsCollections []string `json:"stats_collections"`

	// Collections limit. Only get Databases and collection stats if the total number of collections in the server
	// is less than this value. 0: no limit
	CollectionsLimit int32 `json:"collections_limit,omitempty"`

	// Log level for exporters
	// Enum: [auto fatal error warn info debug]
	LogLevel *string `json:"log_level,omitempty"`
}

// Validate validates this add mongo DB exporter body
func (o *AddMongoDBExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMongoDbExporterBodyTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","fatal","error","warn","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbExporterBodyTypeLogLevelPropEnum = append(addMongoDbExporterBodyTypeLogLevelPropEnum, v)
	}
}

const (

	// AddMongoDBExporterBodyLogLevelAuto captures enum value "auto"
	AddMongoDBExporterBodyLogLevelAuto string = "auto"

	// AddMongoDBExporterBodyLogLevelFatal captures enum value "fatal"
	AddMongoDBExporterBodyLogLevelFatal string = "fatal"

	// AddMongoDBExporterBodyLogLevelError captures enum value "error"
	AddMongoDBExporterBodyLogLevelError string = "error"

	// AddMongoDBExporterBodyLogLevelWarn captures enum value "warn"
	AddMongoDBExporterBodyLogLevelWarn string = "warn"

	// AddMongoDBExporterBodyLogLevelInfo captures enum value "info"
	AddMongoDBExporterBodyLogLevelInfo string = "info"

	// AddMongoDBExporterBodyLogLevelDebug captures enum value "debug"
	AddMongoDBExporterBodyLogLevelDebug string = "debug"
)

// prop value enum
func (o *AddMongoDBExporterBody) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addMongoDbExporterBodyTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBExporterBody) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("body"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add mongo DB exporter body based on context it is used
func (o *AddMongoDBExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBExporterDefaultBody add mongo DB exporter default body
swagger:model AddMongoDBExporterDefaultBody
*/
type AddMongoDBExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddMongoDBExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add mongo DB exporter default body
func (o *AddMongoDBExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add mongo DB exporter default body based on the context it is used
func (o *AddMongoDBExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBExporterDefaultBodyDetailsItems0 add mongo DB exporter default body details items0
swagger:model AddMongoDBExporterDefaultBodyDetailsItems0
*/
type AddMongoDBExporterDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add mongo DB exporter default body details items0
func (o *AddMongoDBExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add mongo DB exporter default body details items0 based on context it is used
func (o *AddMongoDBExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBExporterOKBody add mongo DB exporter OK body
swagger:model AddMongoDBExporterOKBody
*/
type AddMongoDBExporterOKBody struct {

	// mongodb exporter
	MongodbExporter *AddMongoDBExporterOKBodyMongodbExporter `json:"mongodb_exporter,omitempty"`
}

// Validate validates this add mongo DB exporter OK body
func (o *AddMongoDBExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterOKBody) validateMongodbExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbExporterOk" + "." + "mongodb_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addMongoDbExporterOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add mongo DB exporter OK body based on the context it is used
func (o *AddMongoDBExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMongodbExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterOKBody) contextValidateMongodbExporter(ctx context.Context, formats strfmt.Registry) error {

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbExporterOk" + "." + "mongodb_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addMongoDbExporterOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBExporterOKBodyMongodbExporter MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model AddMongoDBExporterOKBodyMongodbExporter
*/
type AddMongoDBExporterOKBodyMongodbExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	//
	// Status fields below.
	DisabledCollectors []string `json:"disabled_collectors"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// List of colletions to get stats from. Can use *
	StatsCollections []string `json:"stats_collections"`

	// Collections limit. Only get Databases and collection stats if the total number of collections in the server
	// is less than this value. 0: no limit
	CollectionsLimit int32 `json:"collections_limit,omitempty"`

	// Enable All collectors.
	EnableAllCollectors bool `json:"enable_all_collectors,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// Log level for exporters
	// Enum: [auto fatal error warn info debug]
	LogLevel *string `json:"log_level,omitempty"`
}

// Validate validates this add mongo DB exporter OK body mongodb exporter
func (o *AddMongoDBExporterOKBodyMongodbExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum = append(addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMongoDBExporterOKBodyMongodbExporterStatusSTARTING captures enum value "STARTING"
	AddMongoDBExporterOKBodyMongodbExporterStatusSTARTING string = "STARTING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusRUNNING captures enum value "RUNNING"
	AddMongoDBExporterOKBodyMongodbExporterStatusRUNNING string = "RUNNING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusWAITING captures enum value "WAITING"
	AddMongoDBExporterOKBodyMongodbExporterStatusWAITING string = "WAITING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusSTOPPING captures enum value "STOPPING"
	AddMongoDBExporterOKBodyMongodbExporterStatusSTOPPING string = "STOPPING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusDONE captures enum value "DONE"
	AddMongoDBExporterOKBodyMongodbExporterStatusDONE string = "DONE"

	// AddMongoDBExporterOKBodyMongodbExporterStatusUNKNOWN captures enum value "UNKNOWN"
	AddMongoDBExporterOKBodyMongodbExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddMongoDBExporterOKBodyMongodbExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBExporterOKBodyMongodbExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMongoDbExporterOk"+"."+"mongodb_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var addMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","fatal","error","warn","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum = append(addMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum, v)
	}
}

const (

	// AddMongoDBExporterOKBodyMongodbExporterLogLevelAuto captures enum value "auto"
	AddMongoDBExporterOKBodyMongodbExporterLogLevelAuto string = "auto"

	// AddMongoDBExporterOKBodyMongodbExporterLogLevelFatal captures enum value "fatal"
	AddMongoDBExporterOKBodyMongodbExporterLogLevelFatal string = "fatal"

	// AddMongoDBExporterOKBodyMongodbExporterLogLevelError captures enum value "error"
	AddMongoDBExporterOKBodyMongodbExporterLogLevelError string = "error"

	// AddMongoDBExporterOKBodyMongodbExporterLogLevelWarn captures enum value "warn"
	AddMongoDBExporterOKBodyMongodbExporterLogLevelWarn string = "warn"

	// AddMongoDBExporterOKBodyMongodbExporterLogLevelInfo captures enum value "info"
	AddMongoDBExporterOKBodyMongodbExporterLogLevelInfo string = "info"

	// AddMongoDBExporterOKBodyMongodbExporterLogLevelDebug captures enum value "debug"
	AddMongoDBExporterOKBodyMongodbExporterLogLevelDebug string = "debug"
)

// prop value enum
func (o *AddMongoDBExporterOKBodyMongodbExporter) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBExporterOKBodyMongodbExporter) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("addMongoDbExporterOk"+"."+"mongodb_exporter"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add mongo DB exporter OK body mongodb exporter based on context it is used
func (o *AddMongoDBExporterOKBodyMongodbExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterOKBodyMongodbExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterOKBodyMongodbExporter) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterOKBodyMongodbExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
