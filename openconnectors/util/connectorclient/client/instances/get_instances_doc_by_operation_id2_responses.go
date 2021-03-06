// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetInstancesDocByOperationId2Reader is a Reader for the GetInstancesDocByOperationId2 structure.
type GetInstancesDocByOperationId2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesDocByOperationId2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesDocByOperationId2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesDocByOperationId2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesDocByOperationId2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesDocByOperationId2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesDocByOperationId2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesDocByOperationId2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesDocByOperationId2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesDocByOperationId2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesDocByOperationId2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesDocByOperationId2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesDocByOperationId2OK creates a GetInstancesDocByOperationId2OK with default headers values
func NewGetInstancesDocByOperationId2OK() *GetInstancesDocByOperationId2OK {
	return &GetInstancesDocByOperationId2OK{}
}

/*GetInstancesDocByOperationId2OK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesDocByOperationId2OK struct {
	Payload interface{}
}

func (o *GetInstancesDocByOperationId2OK) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2OK  %+v", 200, o.Payload)
}

func (o *GetInstancesDocByOperationId2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesDocByOperationId2BadRequest creates a GetInstancesDocByOperationId2BadRequest with default headers values
func NewGetInstancesDocByOperationId2BadRequest() *GetInstancesDocByOperationId2BadRequest {
	return &GetInstancesDocByOperationId2BadRequest{}
}

/*GetInstancesDocByOperationId2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesDocByOperationId2BadRequest struct {
}

func (o *GetInstancesDocByOperationId2BadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2BadRequest ", 400)
}

func (o *GetInstancesDocByOperationId2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2Unauthorized creates a GetInstancesDocByOperationId2Unauthorized with default headers values
func NewGetInstancesDocByOperationId2Unauthorized() *GetInstancesDocByOperationId2Unauthorized {
	return &GetInstancesDocByOperationId2Unauthorized{}
}

/*GetInstancesDocByOperationId2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesDocByOperationId2Unauthorized struct {
}

func (o *GetInstancesDocByOperationId2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2Unauthorized ", 401)
}

func (o *GetInstancesDocByOperationId2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2Forbidden creates a GetInstancesDocByOperationId2Forbidden with default headers values
func NewGetInstancesDocByOperationId2Forbidden() *GetInstancesDocByOperationId2Forbidden {
	return &GetInstancesDocByOperationId2Forbidden{}
}

/*GetInstancesDocByOperationId2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesDocByOperationId2Forbidden struct {
}

func (o *GetInstancesDocByOperationId2Forbidden) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2Forbidden ", 403)
}

func (o *GetInstancesDocByOperationId2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2NotFound creates a GetInstancesDocByOperationId2NotFound with default headers values
func NewGetInstancesDocByOperationId2NotFound() *GetInstancesDocByOperationId2NotFound {
	return &GetInstancesDocByOperationId2NotFound{}
}

/*GetInstancesDocByOperationId2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesDocByOperationId2NotFound struct {
}

func (o *GetInstancesDocByOperationId2NotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2NotFound ", 404)
}

func (o *GetInstancesDocByOperationId2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2MethodNotAllowed creates a GetInstancesDocByOperationId2MethodNotAllowed with default headers values
func NewGetInstancesDocByOperationId2MethodNotAllowed() *GetInstancesDocByOperationId2MethodNotAllowed {
	return &GetInstancesDocByOperationId2MethodNotAllowed{}
}

/*GetInstancesDocByOperationId2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesDocByOperationId2MethodNotAllowed struct {
}

func (o *GetInstancesDocByOperationId2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2MethodNotAllowed ", 405)
}

func (o *GetInstancesDocByOperationId2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2NotAcceptable creates a GetInstancesDocByOperationId2NotAcceptable with default headers values
func NewGetInstancesDocByOperationId2NotAcceptable() *GetInstancesDocByOperationId2NotAcceptable {
	return &GetInstancesDocByOperationId2NotAcceptable{}
}

/*GetInstancesDocByOperationId2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesDocByOperationId2NotAcceptable struct {
}

func (o *GetInstancesDocByOperationId2NotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2NotAcceptable ", 406)
}

func (o *GetInstancesDocByOperationId2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2Conflict creates a GetInstancesDocByOperationId2Conflict with default headers values
func NewGetInstancesDocByOperationId2Conflict() *GetInstancesDocByOperationId2Conflict {
	return &GetInstancesDocByOperationId2Conflict{}
}

/*GetInstancesDocByOperationId2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesDocByOperationId2Conflict struct {
}

func (o *GetInstancesDocByOperationId2Conflict) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2Conflict ", 409)
}

func (o *GetInstancesDocByOperationId2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2UnsupportedMediaType creates a GetInstancesDocByOperationId2UnsupportedMediaType with default headers values
func NewGetInstancesDocByOperationId2UnsupportedMediaType() *GetInstancesDocByOperationId2UnsupportedMediaType {
	return &GetInstancesDocByOperationId2UnsupportedMediaType{}
}

/*GetInstancesDocByOperationId2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesDocByOperationId2UnsupportedMediaType struct {
}

func (o *GetInstancesDocByOperationId2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2UnsupportedMediaType ", 415)
}

func (o *GetInstancesDocByOperationId2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocByOperationId2InternalServerError creates a GetInstancesDocByOperationId2InternalServerError with default headers values
func NewGetInstancesDocByOperationId2InternalServerError() *GetInstancesDocByOperationId2InternalServerError {
	return &GetInstancesDocByOperationId2InternalServerError{}
}

/*GetInstancesDocByOperationId2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesDocByOperationId2InternalServerError struct {
}

func (o *GetInstancesDocByOperationId2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}][%d] getInstancesDocByOperationId2InternalServerError ", 500)
}

func (o *GetInstancesDocByOperationId2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
