// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/alertmanager/ammodels"
)

// GetAlertsReader is a Reader for the GetAlerts structure.
type GetAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertsOK creates a GetAlertsOK with default headers values
func NewGetAlertsOK() *GetAlertsOK {
	return &GetAlertsOK{}
}

/*
GetAlertsOK describes a response with status code 200, with default header values.

Get alerts response
*/
type GetAlertsOK struct {
	Payload ammodels.GettableAlerts
}

func (o *GetAlertsOK) Error() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsOK  %+v", 200, o.Payload)
}

func (o *GetAlertsOK) GetPayload() ammodels.GettableAlerts {
	return o.Payload
}

func (o *GetAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertsBadRequest creates a GetAlertsBadRequest with default headers values
func NewGetAlertsBadRequest() *GetAlertsBadRequest {
	return &GetAlertsBadRequest{}
}

/*
GetAlertsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlertsBadRequest struct {
	Payload string
}

func (o *GetAlertsBadRequest) Error() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertsInternalServerError creates a GetAlertsInternalServerError with default headers values
func NewGetAlertsInternalServerError() *GetAlertsInternalServerError {
	return &GetAlertsInternalServerError{}
}

/*
GetAlertsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAlertsInternalServerError struct {
	Payload string
}

func (o *GetAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertsInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
