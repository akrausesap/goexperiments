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

// GetInstancesEvents2Reader is a Reader for the GetInstancesEvents2 structure.
type GetInstancesEvents2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesEvents2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesEvents2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesEvents2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesEvents2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesEvents2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesEvents2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesEvents2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesEvents2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesEvents2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesEvents2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesEvents2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesEvents2OK creates a GetInstancesEvents2OK with default headers values
func NewGetInstancesEvents2OK() *GetInstancesEvents2OK {
	return &GetInstancesEvents2OK{}
}

/*GetInstancesEvents2OK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesEvents2OK struct {
	Payload []*models.Event
}

func (o *GetInstancesEvents2OK) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2OK  %+v", 200, o.Payload)
}

func (o *GetInstancesEvents2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesEvents2BadRequest creates a GetInstancesEvents2BadRequest with default headers values
func NewGetInstancesEvents2BadRequest() *GetInstancesEvents2BadRequest {
	return &GetInstancesEvents2BadRequest{}
}

/*GetInstancesEvents2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesEvents2BadRequest struct {
}

func (o *GetInstancesEvents2BadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2BadRequest ", 400)
}

func (o *GetInstancesEvents2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2Unauthorized creates a GetInstancesEvents2Unauthorized with default headers values
func NewGetInstancesEvents2Unauthorized() *GetInstancesEvents2Unauthorized {
	return &GetInstancesEvents2Unauthorized{}
}

/*GetInstancesEvents2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesEvents2Unauthorized struct {
}

func (o *GetInstancesEvents2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2Unauthorized ", 401)
}

func (o *GetInstancesEvents2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2Forbidden creates a GetInstancesEvents2Forbidden with default headers values
func NewGetInstancesEvents2Forbidden() *GetInstancesEvents2Forbidden {
	return &GetInstancesEvents2Forbidden{}
}

/*GetInstancesEvents2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesEvents2Forbidden struct {
}

func (o *GetInstancesEvents2Forbidden) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2Forbidden ", 403)
}

func (o *GetInstancesEvents2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2NotFound creates a GetInstancesEvents2NotFound with default headers values
func NewGetInstancesEvents2NotFound() *GetInstancesEvents2NotFound {
	return &GetInstancesEvents2NotFound{}
}

/*GetInstancesEvents2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesEvents2NotFound struct {
}

func (o *GetInstancesEvents2NotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2NotFound ", 404)
}

func (o *GetInstancesEvents2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2MethodNotAllowed creates a GetInstancesEvents2MethodNotAllowed with default headers values
func NewGetInstancesEvents2MethodNotAllowed() *GetInstancesEvents2MethodNotAllowed {
	return &GetInstancesEvents2MethodNotAllowed{}
}

/*GetInstancesEvents2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesEvents2MethodNotAllowed struct {
}

func (o *GetInstancesEvents2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2MethodNotAllowed ", 405)
}

func (o *GetInstancesEvents2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2NotAcceptable creates a GetInstancesEvents2NotAcceptable with default headers values
func NewGetInstancesEvents2NotAcceptable() *GetInstancesEvents2NotAcceptable {
	return &GetInstancesEvents2NotAcceptable{}
}

/*GetInstancesEvents2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesEvents2NotAcceptable struct {
}

func (o *GetInstancesEvents2NotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2NotAcceptable ", 406)
}

func (o *GetInstancesEvents2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2Conflict creates a GetInstancesEvents2Conflict with default headers values
func NewGetInstancesEvents2Conflict() *GetInstancesEvents2Conflict {
	return &GetInstancesEvents2Conflict{}
}

/*GetInstancesEvents2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesEvents2Conflict struct {
}

func (o *GetInstancesEvents2Conflict) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2Conflict ", 409)
}

func (o *GetInstancesEvents2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2UnsupportedMediaType creates a GetInstancesEvents2UnsupportedMediaType with default headers values
func NewGetInstancesEvents2UnsupportedMediaType() *GetInstancesEvents2UnsupportedMediaType {
	return &GetInstancesEvents2UnsupportedMediaType{}
}

/*GetInstancesEvents2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesEvents2UnsupportedMediaType struct {
}

func (o *GetInstancesEvents2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2UnsupportedMediaType ", 415)
}

func (o *GetInstancesEvents2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesEvents2InternalServerError creates a GetInstancesEvents2InternalServerError with default headers values
func NewGetInstancesEvents2InternalServerError() *GetInstancesEvents2InternalServerError {
	return &GetInstancesEvents2InternalServerError{}
}

/*GetInstancesEvents2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesEvents2InternalServerError struct {
}

func (o *GetInstancesEvents2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/events][%d] getInstancesEvents2InternalServerError ", 500)
}

func (o *GetInstancesEvents2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
