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

// ReplaceInstancesObjectsObjectNameDefinitions2Reader is a Reader for the ReplaceInstancesObjectsObjectNameDefinitions2 structure.
type ReplaceInstancesObjectsObjectNameDefinitions2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceInstancesObjectsObjectNameDefinitions2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewReplaceInstancesObjectsObjectNameDefinitions2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceInstancesObjectsObjectNameDefinitions2OK creates a ReplaceInstancesObjectsObjectNameDefinitions2OK with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2OK() *ReplaceInstancesObjectsObjectNameDefinitions2OK {
	return &ReplaceInstancesObjectsObjectNameDefinitions2OK{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2OK handles this case with default header values.

OK - Everything worked as expected
*/
type ReplaceInstancesObjectsObjectNameDefinitions2OK struct {
	Payload *models.Definition
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2OK) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2OK  %+v", 200, o.Payload)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Definition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2BadRequest creates a ReplaceInstancesObjectsObjectNameDefinitions2BadRequest with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2BadRequest() *ReplaceInstancesObjectsObjectNameDefinitions2BadRequest {
	return &ReplaceInstancesObjectsObjectNameDefinitions2BadRequest{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type ReplaceInstancesObjectsObjectNameDefinitions2BadRequest struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2BadRequest ", 400)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2Unauthorized creates a ReplaceInstancesObjectsObjectNameDefinitions2Unauthorized with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2Unauthorized() *ReplaceInstancesObjectsObjectNameDefinitions2Unauthorized {
	return &ReplaceInstancesObjectsObjectNameDefinitions2Unauthorized{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type ReplaceInstancesObjectsObjectNameDefinitions2Unauthorized struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2Unauthorized ", 401)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2Forbidden creates a ReplaceInstancesObjectsObjectNameDefinitions2Forbidden with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2Forbidden() *ReplaceInstancesObjectsObjectNameDefinitions2Forbidden {
	return &ReplaceInstancesObjectsObjectNameDefinitions2Forbidden{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type ReplaceInstancesObjectsObjectNameDefinitions2Forbidden struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2Forbidden ", 403)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2NotFound creates a ReplaceInstancesObjectsObjectNameDefinitions2NotFound with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2NotFound() *ReplaceInstancesObjectsObjectNameDefinitions2NotFound {
	return &ReplaceInstancesObjectsObjectNameDefinitions2NotFound{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type ReplaceInstancesObjectsObjectNameDefinitions2NotFound struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2NotFound) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2NotFound ", 404)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed creates a ReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed() *ReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed {
	return &ReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type ReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed ", 405)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable creates a ReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable() *ReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable {
	return &ReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type ReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2NotAcceptable ", 406)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2Conflict creates a ReplaceInstancesObjectsObjectNameDefinitions2Conflict with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2Conflict() *ReplaceInstancesObjectsObjectNameDefinitions2Conflict {
	return &ReplaceInstancesObjectsObjectNameDefinitions2Conflict{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type ReplaceInstancesObjectsObjectNameDefinitions2Conflict struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2Conflict) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2Conflict ", 409)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType creates a ReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType() *ReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType {
	return &ReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type ReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType ", 415)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceInstancesObjectsObjectNameDefinitions2InternalServerError creates a ReplaceInstancesObjectsObjectNameDefinitions2InternalServerError with default headers values
func NewReplaceInstancesObjectsObjectNameDefinitions2InternalServerError() *ReplaceInstancesObjectsObjectNameDefinitions2InternalServerError {
	return &ReplaceInstancesObjectsObjectNameDefinitions2InternalServerError{}
}

/*ReplaceInstancesObjectsObjectNameDefinitions2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type ReplaceInstancesObjectsObjectNameDefinitions2InternalServerError struct {
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /instances/{id}/objects/{objectName}/definitions][%d] replaceInstancesObjectsObjectNameDefinitions2InternalServerError ", 500)
}

func (o *ReplaceInstancesObjectsObjectNameDefinitions2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
