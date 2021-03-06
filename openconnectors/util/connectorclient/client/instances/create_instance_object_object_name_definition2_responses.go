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

// CreateInstanceObjectObjectNameDefinition2Reader is a Reader for the CreateInstanceObjectObjectNameDefinition2 structure.
type CreateInstanceObjectObjectNameDefinition2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInstanceObjectObjectNameDefinition2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateInstanceObjectObjectNameDefinition2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateInstanceObjectObjectNameDefinition2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateInstanceObjectObjectNameDefinition2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateInstanceObjectObjectNameDefinition2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateInstanceObjectObjectNameDefinition2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCreateInstanceObjectObjectNameDefinition2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewCreateInstanceObjectObjectNameDefinition2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateInstanceObjectObjectNameDefinition2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewCreateInstanceObjectObjectNameDefinition2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateInstanceObjectObjectNameDefinition2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateInstanceObjectObjectNameDefinition2OK creates a CreateInstanceObjectObjectNameDefinition2OK with default headers values
func NewCreateInstanceObjectObjectNameDefinition2OK() *CreateInstanceObjectObjectNameDefinition2OK {
	return &CreateInstanceObjectObjectNameDefinition2OK{}
}

/*CreateInstanceObjectObjectNameDefinition2OK handles this case with default header values.

OK - Everything worked as expected
*/
type CreateInstanceObjectObjectNameDefinition2OK struct {
	Payload *models.Definition
}

func (o *CreateInstanceObjectObjectNameDefinition2OK) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2OK  %+v", 200, o.Payload)
}

func (o *CreateInstanceObjectObjectNameDefinition2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Definition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2BadRequest creates a CreateInstanceObjectObjectNameDefinition2BadRequest with default headers values
func NewCreateInstanceObjectObjectNameDefinition2BadRequest() *CreateInstanceObjectObjectNameDefinition2BadRequest {
	return &CreateInstanceObjectObjectNameDefinition2BadRequest{}
}

/*CreateInstanceObjectObjectNameDefinition2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type CreateInstanceObjectObjectNameDefinition2BadRequest struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2BadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2BadRequest ", 400)
}

func (o *CreateInstanceObjectObjectNameDefinition2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2Unauthorized creates a CreateInstanceObjectObjectNameDefinition2Unauthorized with default headers values
func NewCreateInstanceObjectObjectNameDefinition2Unauthorized() *CreateInstanceObjectObjectNameDefinition2Unauthorized {
	return &CreateInstanceObjectObjectNameDefinition2Unauthorized{}
}

/*CreateInstanceObjectObjectNameDefinition2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type CreateInstanceObjectObjectNameDefinition2Unauthorized struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2Unauthorized ", 401)
}

func (o *CreateInstanceObjectObjectNameDefinition2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2Forbidden creates a CreateInstanceObjectObjectNameDefinition2Forbidden with default headers values
func NewCreateInstanceObjectObjectNameDefinition2Forbidden() *CreateInstanceObjectObjectNameDefinition2Forbidden {
	return &CreateInstanceObjectObjectNameDefinition2Forbidden{}
}

/*CreateInstanceObjectObjectNameDefinition2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type CreateInstanceObjectObjectNameDefinition2Forbidden struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2Forbidden) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2Forbidden ", 403)
}

func (o *CreateInstanceObjectObjectNameDefinition2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2NotFound creates a CreateInstanceObjectObjectNameDefinition2NotFound with default headers values
func NewCreateInstanceObjectObjectNameDefinition2NotFound() *CreateInstanceObjectObjectNameDefinition2NotFound {
	return &CreateInstanceObjectObjectNameDefinition2NotFound{}
}

/*CreateInstanceObjectObjectNameDefinition2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type CreateInstanceObjectObjectNameDefinition2NotFound struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2NotFound) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2NotFound ", 404)
}

func (o *CreateInstanceObjectObjectNameDefinition2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2MethodNotAllowed creates a CreateInstanceObjectObjectNameDefinition2MethodNotAllowed with default headers values
func NewCreateInstanceObjectObjectNameDefinition2MethodNotAllowed() *CreateInstanceObjectObjectNameDefinition2MethodNotAllowed {
	return &CreateInstanceObjectObjectNameDefinition2MethodNotAllowed{}
}

/*CreateInstanceObjectObjectNameDefinition2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type CreateInstanceObjectObjectNameDefinition2MethodNotAllowed struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2MethodNotAllowed ", 405)
}

func (o *CreateInstanceObjectObjectNameDefinition2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2NotAcceptable creates a CreateInstanceObjectObjectNameDefinition2NotAcceptable with default headers values
func NewCreateInstanceObjectObjectNameDefinition2NotAcceptable() *CreateInstanceObjectObjectNameDefinition2NotAcceptable {
	return &CreateInstanceObjectObjectNameDefinition2NotAcceptable{}
}

/*CreateInstanceObjectObjectNameDefinition2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type CreateInstanceObjectObjectNameDefinition2NotAcceptable struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2NotAcceptable) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2NotAcceptable ", 406)
}

func (o *CreateInstanceObjectObjectNameDefinition2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2Conflict creates a CreateInstanceObjectObjectNameDefinition2Conflict with default headers values
func NewCreateInstanceObjectObjectNameDefinition2Conflict() *CreateInstanceObjectObjectNameDefinition2Conflict {
	return &CreateInstanceObjectObjectNameDefinition2Conflict{}
}

/*CreateInstanceObjectObjectNameDefinition2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type CreateInstanceObjectObjectNameDefinition2Conflict struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2Conflict) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2Conflict ", 409)
}

func (o *CreateInstanceObjectObjectNameDefinition2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2UnsupportedMediaType creates a CreateInstanceObjectObjectNameDefinition2UnsupportedMediaType with default headers values
func NewCreateInstanceObjectObjectNameDefinition2UnsupportedMediaType() *CreateInstanceObjectObjectNameDefinition2UnsupportedMediaType {
	return &CreateInstanceObjectObjectNameDefinition2UnsupportedMediaType{}
}

/*CreateInstanceObjectObjectNameDefinition2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type CreateInstanceObjectObjectNameDefinition2UnsupportedMediaType struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2UnsupportedMediaType ", 415)
}

func (o *CreateInstanceObjectObjectNameDefinition2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstanceObjectObjectNameDefinition2InternalServerError creates a CreateInstanceObjectObjectNameDefinition2InternalServerError with default headers values
func NewCreateInstanceObjectObjectNameDefinition2InternalServerError() *CreateInstanceObjectObjectNameDefinition2InternalServerError {
	return &CreateInstanceObjectObjectNameDefinition2InternalServerError{}
}

/*CreateInstanceObjectObjectNameDefinition2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type CreateInstanceObjectObjectNameDefinition2InternalServerError struct {
}

func (o *CreateInstanceObjectObjectNameDefinition2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /instances/{id}/objects/{objectName}/definitions][%d] createInstanceObjectObjectNameDefinition2InternalServerError ", 500)
}

func (o *CreateInstanceObjectObjectNameDefinition2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
