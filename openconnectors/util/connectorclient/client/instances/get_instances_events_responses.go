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

// GetInstancesEventsReader is a Reader for the GetInstancesEvents structure.
type GetInstancesEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesEventsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesEventsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesEventsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesEventsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesEventsOK creates a GetInstancesEventsOK with default headers values
func NewGetInstancesEventsOK() *GetInstancesEventsOK {
	return &GetInstancesEventsOK{}
}

/*GetInstancesEventsOK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesEventsOK struct {
	Payload []*models.Event
}

func (o *GetInstancesEventsOK) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsOK  %+v", 200, o.Payload)
}

func (o *GetInstancesEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesEventsBadRequest creates a GetInstancesEventsBadRequest with default headers values
func NewGetInstancesEventsBadRequest() *GetInstancesEventsBadRequest {
	return &GetInstancesEventsBadRequest{}
}

/*GetInstancesEventsBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesEventsBadRequest struct {
}

func (o *GetInstancesEventsBadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsBadRequest ", 400)
}

func (o *GetInstancesEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsUnauthorized creates a GetInstancesEventsUnauthorized with default headers values
func NewGetInstancesEventsUnauthorized() *GetInstancesEventsUnauthorized {
	return &GetInstancesEventsUnauthorized{}
}

/*GetInstancesEventsUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesEventsUnauthorized struct {
}

func (o *GetInstancesEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsUnauthorized ", 401)
}

func (o *GetInstancesEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsForbidden creates a GetInstancesEventsForbidden with default headers values
func NewGetInstancesEventsForbidden() *GetInstancesEventsForbidden {
	return &GetInstancesEventsForbidden{}
}

/*GetInstancesEventsForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesEventsForbidden struct {
}

func (o *GetInstancesEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsForbidden ", 403)
}

func (o *GetInstancesEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsNotFound creates a GetInstancesEventsNotFound with default headers values
func NewGetInstancesEventsNotFound() *GetInstancesEventsNotFound {
	return &GetInstancesEventsNotFound{}
}

/*GetInstancesEventsNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesEventsNotFound struct {
}

func (o *GetInstancesEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsNotFound ", 404)
}

func (o *GetInstancesEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsMethodNotAllowed creates a GetInstancesEventsMethodNotAllowed with default headers values
func NewGetInstancesEventsMethodNotAllowed() *GetInstancesEventsMethodNotAllowed {
	return &GetInstancesEventsMethodNotAllowed{}
}

/*GetInstancesEventsMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesEventsMethodNotAllowed struct {
}

func (o *GetInstancesEventsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsMethodNotAllowed ", 405)
}

func (o *GetInstancesEventsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsNotAcceptable creates a GetInstancesEventsNotAcceptable with default headers values
func NewGetInstancesEventsNotAcceptable() *GetInstancesEventsNotAcceptable {
	return &GetInstancesEventsNotAcceptable{}
}

/*GetInstancesEventsNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesEventsNotAcceptable struct {
}

func (o *GetInstancesEventsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsNotAcceptable ", 406)
}

func (o *GetInstancesEventsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsConflict creates a GetInstancesEventsConflict with default headers values
func NewGetInstancesEventsConflict() *GetInstancesEventsConflict {
	return &GetInstancesEventsConflict{}
}

/*GetInstancesEventsConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesEventsConflict struct {
}

func (o *GetInstancesEventsConflict) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsConflict ", 409)
}

func (o *GetInstancesEventsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsUnsupportedMediaType creates a GetInstancesEventsUnsupportedMediaType with default headers values
func NewGetInstancesEventsUnsupportedMediaType() *GetInstancesEventsUnsupportedMediaType {
	return &GetInstancesEventsUnsupportedMediaType{}
}

/*GetInstancesEventsUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesEventsUnsupportedMediaType struct {
}

func (o *GetInstancesEventsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsUnsupportedMediaType ", 415)
}

func (o *GetInstancesEventsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEventsInternalServerError creates a GetInstancesEventsInternalServerError with default headers values
func NewGetInstancesEventsInternalServerError() *GetInstancesEventsInternalServerError {
	return &GetInstancesEventsInternalServerError{}
}

/*GetInstancesEventsInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesEventsInternalServerError struct {
}

func (o *GetInstancesEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/events][%d] getInstancesEventsInternalServerError ", 500)
}

func (o *GetInstancesEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
