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

// GetInstancesDocsDefinitions2Reader is a Reader for the GetInstancesDocsDefinitions2 structure.
type GetInstancesDocsDefinitions2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesDocsDefinitions2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesDocsDefinitions2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesDocsDefinitions2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesDocsDefinitions2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesDocsDefinitions2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesDocsDefinitions2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesDocsDefinitions2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesDocsDefinitions2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesDocsDefinitions2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesDocsDefinitions2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesDocsDefinitions2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesDocsDefinitions2OK creates a GetInstancesDocsDefinitions2OK with default headers values
func NewGetInstancesDocsDefinitions2OK() *GetInstancesDocsDefinitions2OK {
	return &GetInstancesDocsDefinitions2OK{}
}

/*GetInstancesDocsDefinitions2OK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesDocsDefinitions2OK struct {
	Payload interface{}
}

func (o *GetInstancesDocsDefinitions2OK) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2OK  %+v", 200, o.Payload)
}

func (o *GetInstancesDocsDefinitions2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesDocsDefinitions2BadRequest creates a GetInstancesDocsDefinitions2BadRequest with default headers values
func NewGetInstancesDocsDefinitions2BadRequest() *GetInstancesDocsDefinitions2BadRequest {
	return &GetInstancesDocsDefinitions2BadRequest{}
}

/*GetInstancesDocsDefinitions2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesDocsDefinitions2BadRequest struct {
}

func (o *GetInstancesDocsDefinitions2BadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2BadRequest ", 400)
}

func (o *GetInstancesDocsDefinitions2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2Unauthorized creates a GetInstancesDocsDefinitions2Unauthorized with default headers values
func NewGetInstancesDocsDefinitions2Unauthorized() *GetInstancesDocsDefinitions2Unauthorized {
	return &GetInstancesDocsDefinitions2Unauthorized{}
}

/*GetInstancesDocsDefinitions2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesDocsDefinitions2Unauthorized struct {
}

func (o *GetInstancesDocsDefinitions2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2Unauthorized ", 401)
}

func (o *GetInstancesDocsDefinitions2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2Forbidden creates a GetInstancesDocsDefinitions2Forbidden with default headers values
func NewGetInstancesDocsDefinitions2Forbidden() *GetInstancesDocsDefinitions2Forbidden {
	return &GetInstancesDocsDefinitions2Forbidden{}
}

/*GetInstancesDocsDefinitions2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesDocsDefinitions2Forbidden struct {
}

func (o *GetInstancesDocsDefinitions2Forbidden) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2Forbidden ", 403)
}

func (o *GetInstancesDocsDefinitions2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2NotFound creates a GetInstancesDocsDefinitions2NotFound with default headers values
func NewGetInstancesDocsDefinitions2NotFound() *GetInstancesDocsDefinitions2NotFound {
	return &GetInstancesDocsDefinitions2NotFound{}
}

/*GetInstancesDocsDefinitions2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesDocsDefinitions2NotFound struct {
}

func (o *GetInstancesDocsDefinitions2NotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2NotFound ", 404)
}

func (o *GetInstancesDocsDefinitions2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2MethodNotAllowed creates a GetInstancesDocsDefinitions2MethodNotAllowed with default headers values
func NewGetInstancesDocsDefinitions2MethodNotAllowed() *GetInstancesDocsDefinitions2MethodNotAllowed {
	return &GetInstancesDocsDefinitions2MethodNotAllowed{}
}

/*GetInstancesDocsDefinitions2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesDocsDefinitions2MethodNotAllowed struct {
}

func (o *GetInstancesDocsDefinitions2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2MethodNotAllowed ", 405)
}

func (o *GetInstancesDocsDefinitions2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2NotAcceptable creates a GetInstancesDocsDefinitions2NotAcceptable with default headers values
func NewGetInstancesDocsDefinitions2NotAcceptable() *GetInstancesDocsDefinitions2NotAcceptable {
	return &GetInstancesDocsDefinitions2NotAcceptable{}
}

/*GetInstancesDocsDefinitions2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesDocsDefinitions2NotAcceptable struct {
}

func (o *GetInstancesDocsDefinitions2NotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2NotAcceptable ", 406)
}

func (o *GetInstancesDocsDefinitions2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2Conflict creates a GetInstancesDocsDefinitions2Conflict with default headers values
func NewGetInstancesDocsDefinitions2Conflict() *GetInstancesDocsDefinitions2Conflict {
	return &GetInstancesDocsDefinitions2Conflict{}
}

/*GetInstancesDocsDefinitions2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesDocsDefinitions2Conflict struct {
}

func (o *GetInstancesDocsDefinitions2Conflict) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2Conflict ", 409)
}

func (o *GetInstancesDocsDefinitions2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2UnsupportedMediaType creates a GetInstancesDocsDefinitions2UnsupportedMediaType with default headers values
func NewGetInstancesDocsDefinitions2UnsupportedMediaType() *GetInstancesDocsDefinitions2UnsupportedMediaType {
	return &GetInstancesDocsDefinitions2UnsupportedMediaType{}
}

/*GetInstancesDocsDefinitions2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesDocsDefinitions2UnsupportedMediaType struct {
}

func (o *GetInstancesDocsDefinitions2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2UnsupportedMediaType ", 415)
}

func (o *GetInstancesDocsDefinitions2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitions2InternalServerError creates a GetInstancesDocsDefinitions2InternalServerError with default headers values
func NewGetInstancesDocsDefinitions2InternalServerError() *GetInstancesDocsDefinitions2InternalServerError {
	return &GetInstancesDocsDefinitions2InternalServerError{}
}

/*GetInstancesDocsDefinitions2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesDocsDefinitions2InternalServerError struct {
}

func (o *GetInstancesDocsDefinitions2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/docs/{operationId}/definitions][%d] getInstancesDocsDefinitions2InternalServerError ", 500)
}

func (o *GetInstancesDocsDefinitions2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}