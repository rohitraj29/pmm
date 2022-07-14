// Code generated by go-swagger; DO NOT EDIT.

package prometheus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/grafana/gmodels"
)

// RouteGetRuleStatusesReader is a Reader for the RouteGetRuleStatuses structure.
type RouteGetRuleStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetRuleStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetRuleStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteGetRuleStatusesOK creates a RouteGetRuleStatusesOK with default headers values
func NewRouteGetRuleStatusesOK() *RouteGetRuleStatusesOK {
	return &RouteGetRuleStatusesOK{}
}

/* RouteGetRuleStatusesOK describes a response with status code 200, with default header values.

RuleResponse
*/
type RouteGetRuleStatusesOK struct {
	Payload *gmodels.RuleResponse
}

func (o *RouteGetRuleStatusesOK) Error() string {
	return fmt.Sprintf("[GET /api/prometheus/{Recipient}/api/v1/rules][%d] routeGetRuleStatusesOK  %+v", 200, o.Payload)
}

func (o *RouteGetRuleStatusesOK) GetPayload() *gmodels.RuleResponse {
	return o.Payload
}

func (o *RouteGetRuleStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(gmodels.RuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
