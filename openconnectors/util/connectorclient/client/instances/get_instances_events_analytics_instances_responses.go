// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/akrausesap/goexperiments/openconnectors/util/connectorclient/models"
)

// GetInstancesEventsAnalyticsInstancesReader is a Reader for the GetInstancesEventsAnalyticsInstances structure.
type GetInstancesEventsAnalyticsInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesEventsAnalyticsInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesEventsAnalyticsInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesEventsAnalyticsInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesEventsAnalyticsInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesEventsAnalyticsInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesEventsAnalyticsInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesEventsAnalyticsInstancesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesEventsAnalyticsInstancesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesEventsAnalyticsInstancesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesEventsAnalyticsInstancesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesEventsAnalyticsInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesEventsAnalyticsInstancesOK creates a GetInstancesEventsAnalyticsInstancesOK with default headers values
func NewGetInstancesEventsAnalyticsInstancesOK() *GetInstancesEventsAnalyticsInstancesOK {
	return &GetInstancesEventsAnalyticsInstancesOK{}
}

/*GetInstancesEventsAnalyticsInstancesOK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesEventsAnalyticsInstancesOK struct {
	Payload []*models.InstanceAnalyticsEntry
}

func (o *GetInstancesEventsAnalyticsInstancesOK) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesOK  %+v", 200, o.Payload)
}

func (o *GetInstancesEventsAnalyticsInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesBadRequest creates a GetInstancesEventsAnalyticsInstancesBadRequest with default headers values
func NewGetInstancesEventsAnalyticsInstancesBadRequest() *GetInstancesEventsAnalyticsInstancesBadRequest {
	return &GetInstancesEventsAnalyticsInstancesBadRequest{}
}

/*GetInstancesEventsAnalyticsInstancesBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesEventsAnalyticsInstancesBadRequest struct {
}

func (o *GetInstancesEventsAnalyticsInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesBadRequest ", 400)
}

func (o *GetInstancesEventsAnalyticsInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesUnauthorized creates a GetInstancesEventsAnalyticsInstancesUnauthorized with default headers values
func NewGetInstancesEventsAnalyticsInstancesUnauthorized() *GetInstancesEventsAnalyticsInstancesUnauthorized {
	return &GetInstancesEventsAnalyticsInstancesUnauthorized{}
}

/*GetInstancesEventsAnalyticsInstancesUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesEventsAnalyticsInstancesUnauthorized struct {
}

func (o *GetInstancesEventsAnalyticsInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesUnauthorized ", 401)
}

func (o *GetInstancesEventsAnalyticsInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesForbidden creates a GetInstancesEventsAnalyticsInstancesForbidden with default headers values
func NewGetInstancesEventsAnalyticsInstancesForbidden() *GetInstancesEventsAnalyticsInstancesForbidden {
	return &GetInstancesEventsAnalyticsInstancesForbidden{}
}

/*GetInstancesEventsAnalyticsInstancesForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesEventsAnalyticsInstancesForbidden struct {
}

func (o *GetInstancesEventsAnalyticsInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesForbidden ", 403)
}

func (o *GetInstancesEventsAnalyticsInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesNotFound creates a GetInstancesEventsAnalyticsInstancesNotFound with default headers values
func NewGetInstancesEventsAnalyticsInstancesNotFound() *GetInstancesEventsAnalyticsInstancesNotFound {
	return &GetInstancesEventsAnalyticsInstancesNotFound{}
}

/*GetInstancesEventsAnalyticsInstancesNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesEventsAnalyticsInstancesNotFound struct {
}

func (o *GetInstancesEventsAnalyticsInstancesNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesNotFound ", 404)
}

func (o *GetInstancesEventsAnalyticsInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesMethodNotAllowed creates a GetInstancesEventsAnalyticsInstancesMethodNotAllowed with default headers values
func NewGetInstancesEventsAnalyticsInstancesMethodNotAllowed() *GetInstancesEventsAnalyticsInstancesMethodNotAllowed {
	return &GetInstancesEventsAnalyticsInstancesMethodNotAllowed{}
}

/*GetInstancesEventsAnalyticsInstancesMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesEventsAnalyticsInstancesMethodNotAllowed struct {
}

func (o *GetInstancesEventsAnalyticsInstancesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesMethodNotAllowed ", 405)
}

func (o *GetInstancesEventsAnalyticsInstancesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesNotAcceptable creates a GetInstancesEventsAnalyticsInstancesNotAcceptable with default headers values
func NewGetInstancesEventsAnalyticsInstancesNotAcceptable() *GetInstancesEventsAnalyticsInstancesNotAcceptable {
	return &GetInstancesEventsAnalyticsInstancesNotAcceptable{}
}

/*GetInstancesEventsAnalyticsInstancesNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesEventsAnalyticsInstancesNotAcceptable struct {
}

func (o *GetInstancesEventsAnalyticsInstancesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesNotAcceptable ", 406)
}

func (o *GetInstancesEventsAnalyticsInstancesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesConflict creates a GetInstancesEventsAnalyticsInstancesConflict with default headers values
func NewGetInstancesEventsAnalyticsInstancesConflict() *GetInstancesEventsAnalyticsInstancesConflict {
	return &GetInstancesEventsAnalyticsInstancesConflict{}
}

/*GetInstancesEventsAnalyticsInstancesConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesEventsAnalyticsInstancesConflict struct {
}

func (o *GetInstancesEventsAnalyticsInstancesConflict) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesConflict ", 409)
}

func (o *GetInstancesEventsAnalyticsInstancesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesUnsupportedMediaType creates a GetInstancesEventsAnalyticsInstancesUnsupportedMediaType with default headers values
func NewGetInstancesEventsAnalyticsInstancesUnsupportedMediaType() *GetInstancesEventsAnalyticsInstancesUnsupportedMediaType {
	return &GetInstancesEventsAnalyticsInstancesUnsupportedMediaType{}
}

/*GetInstancesEventsAnalyticsInstancesUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesEventsAnalyticsInstancesUnsupportedMediaType struct {
}

func (o *GetInstancesEventsAnalyticsInstancesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesUnsupportedMediaType ", 415)
}

func (o *GetInstancesEventsAnalyticsInstancesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsAnalyticsInstancesInternalServerError creates a GetInstancesEventsAnalyticsInstancesInternalServerError with default headers values
func NewGetInstancesEventsAnalyticsInstancesInternalServerError() *GetInstancesEventsAnalyticsInstancesInternalServerError {
	return &GetInstancesEventsAnalyticsInstancesInternalServerError{}
}

/*GetInstancesEventsAnalyticsInstancesInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesEventsAnalyticsInstancesInternalServerError struct {
}

func (o *GetInstancesEventsAnalyticsInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/events/analytics/instances][%d] getInstancesEventsAnalyticsInstancesInternalServerError ", 500)
}

func (o *GetInstancesEventsAnalyticsInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
