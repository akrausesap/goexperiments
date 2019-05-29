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

// GetInstancesObjectsObjectNameDefinitions2Reader is a Reader for the GetInstancesObjectsObjectNameDefinitions2 structure.
type GetInstancesObjectsObjectNameDefinitions2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesObjectsObjectNameDefinitions2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesObjectsObjectNameDefinitions2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesObjectsObjectNameDefinitions2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesObjectsObjectNameDefinitions2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesObjectsObjectNameDefinitions2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesObjectsObjectNameDefinitions2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesObjectsObjectNameDefinitions2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesObjectsObjectNameDefinitions2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesObjectsObjectNameDefinitions2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesObjectsObjectNameDefinitions2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesObjectsObjectNameDefinitions2OK creates a GetInstancesObjectsObjectNameDefinitions2OK with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2OK() *GetInstancesObjectsObjectNameDefinitions2OK {
	return &GetInstancesObjectsObjectNameDefinitions2OK{}
}

/*GetInstancesObjectsObjectNameDefinitions2OK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesObjectsObjectNameDefinitions2OK struct {
	Payload *models.Definition
}

func (o *GetInstancesObjectsObjectNameDefinitions2OK) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2OK  %+v", 200, o.Payload)
}

func (o *GetInstancesObjectsObjectNameDefinitions2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Definition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2BadRequest creates a GetInstancesObjectsObjectNameDefinitions2BadRequest with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2BadRequest() *GetInstancesObjectsObjectNameDefinitions2BadRequest {
	return &GetInstancesObjectsObjectNameDefinitions2BadRequest{}
}

/*GetInstancesObjectsObjectNameDefinitions2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesObjectsObjectNameDefinitions2BadRequest struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2BadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2BadRequest ", 400)
}

func (o *GetInstancesObjectsObjectNameDefinitions2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2Unauthorized creates a GetInstancesObjectsObjectNameDefinitions2Unauthorized with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2Unauthorized() *GetInstancesObjectsObjectNameDefinitions2Unauthorized {
	return &GetInstancesObjectsObjectNameDefinitions2Unauthorized{}
}

/*GetInstancesObjectsObjectNameDefinitions2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesObjectsObjectNameDefinitions2Unauthorized struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2Unauthorized ", 401)
}

func (o *GetInstancesObjectsObjectNameDefinitions2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2Forbidden creates a GetInstancesObjectsObjectNameDefinitions2Forbidden with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2Forbidden() *GetInstancesObjectsObjectNameDefinitions2Forbidden {
	return &GetInstancesObjectsObjectNameDefinitions2Forbidden{}
}

/*GetInstancesObjectsObjectNameDefinitions2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesObjectsObjectNameDefinitions2Forbidden struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2Forbidden) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2Forbidden ", 403)
}

func (o *GetInstancesObjectsObjectNameDefinitions2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2NotFound creates a GetInstancesObjectsObjectNameDefinitions2NotFound with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2NotFound() *GetInstancesObjectsObjectNameDefinitions2NotFound {
	return &GetInstancesObjectsObjectNameDefinitions2NotFound{}
}

/*GetInstancesObjectsObjectNameDefinitions2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesObjectsObjectNameDefinitions2NotFound struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2NotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2NotFound ", 404)
}

func (o *GetInstancesObjectsObjectNameDefinitions2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2MethodNotAllowed creates a GetInstancesObjectsObjectNameDefinitions2MethodNotAllowed with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2MethodNotAllowed() *GetInstancesObjectsObjectNameDefinitions2MethodNotAllowed {
	return &GetInstancesObjectsObjectNameDefinitions2MethodNotAllowed{}
}

/*GetInstancesObjectsObjectNameDefinitions2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesObjectsObjectNameDefinitions2MethodNotAllowed struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2MethodNotAllowed ", 405)
}

func (o *GetInstancesObjectsObjectNameDefinitions2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2NotAcceptable creates a GetInstancesObjectsObjectNameDefinitions2NotAcceptable with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2NotAcceptable() *GetInstancesObjectsObjectNameDefinitions2NotAcceptable {
	return &GetInstancesObjectsObjectNameDefinitions2NotAcceptable{}
}

/*GetInstancesObjectsObjectNameDefinitions2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesObjectsObjectNameDefinitions2NotAcceptable struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2NotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2NotAcceptable ", 406)
}

func (o *GetInstancesObjectsObjectNameDefinitions2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2Conflict creates a GetInstancesObjectsObjectNameDefinitions2Conflict with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2Conflict() *GetInstancesObjectsObjectNameDefinitions2Conflict {
	return &GetInstancesObjectsObjectNameDefinitions2Conflict{}
}

/*GetInstancesObjectsObjectNameDefinitions2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesObjectsObjectNameDefinitions2Conflict struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2Conflict) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2Conflict ", 409)
}

func (o *GetInstancesObjectsObjectNameDefinitions2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType creates a GetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType() *GetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType {
	return &GetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType{}
}

/*GetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2UnsupportedMediaType ", 415)
}

func (o *GetInstancesObjectsObjectNameDefinitions2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesObjectsObjectNameDefinitions2InternalServerError creates a GetInstancesObjectsObjectNameDefinitions2InternalServerError with default headers values
func NewGetInstancesObjectsObjectNameDefinitions2InternalServerError() *GetInstancesObjectsObjectNameDefinitions2InternalServerError {
	return &GetInstancesObjectsObjectNameDefinitions2InternalServerError{}
}

/*GetInstancesObjectsObjectNameDefinitions2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesObjectsObjectNameDefinitions2InternalServerError struct {
}

func (o *GetInstancesObjectsObjectNameDefinitions2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/objects/{objectName}/definitions][%d] getInstancesObjectsObjectNameDefinitions2InternalServerError ", 500)
}

func (o *GetInstancesObjectsObjectNameDefinitions2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}