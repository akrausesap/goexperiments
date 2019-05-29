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

// GetInstancesConfiguration2Reader is a Reader for the GetInstancesConfiguration2 structure.
type GetInstancesConfiguration2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesConfiguration2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesConfiguration2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesConfiguration2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesConfiguration2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesConfiguration2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesConfiguration2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesConfiguration2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesConfiguration2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesConfiguration2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesConfiguration2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesConfiguration2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesConfiguration2OK creates a GetInstancesConfiguration2OK with default headers values
func NewGetInstancesConfiguration2OK() *GetInstancesConfiguration2OK {
	return &GetInstancesConfiguration2OK{}
}

/*GetInstancesConfiguration2OK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesConfiguration2OK struct {
	Payload []*models.ElementInstanceConfig
}

func (o *GetInstancesConfiguration2OK) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2OK  %+v", 200, o.Payload)
}

func (o *GetInstancesConfiguration2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesConfiguration2BadRequest creates a GetInstancesConfiguration2BadRequest with default headers values
func NewGetInstancesConfiguration2BadRequest() *GetInstancesConfiguration2BadRequest {
	return &GetInstancesConfiguration2BadRequest{}
}

/*GetInstancesConfiguration2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesConfiguration2BadRequest struct {
}

func (o *GetInstancesConfiguration2BadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2BadRequest ", 400)
}

func (o *GetInstancesConfiguration2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2Unauthorized creates a GetInstancesConfiguration2Unauthorized with default headers values
func NewGetInstancesConfiguration2Unauthorized() *GetInstancesConfiguration2Unauthorized {
	return &GetInstancesConfiguration2Unauthorized{}
}

/*GetInstancesConfiguration2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesConfiguration2Unauthorized struct {
}

func (o *GetInstancesConfiguration2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2Unauthorized ", 401)
}

func (o *GetInstancesConfiguration2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2Forbidden creates a GetInstancesConfiguration2Forbidden with default headers values
func NewGetInstancesConfiguration2Forbidden() *GetInstancesConfiguration2Forbidden {
	return &GetInstancesConfiguration2Forbidden{}
}

/*GetInstancesConfiguration2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesConfiguration2Forbidden struct {
}

func (o *GetInstancesConfiguration2Forbidden) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2Forbidden ", 403)
}

func (o *GetInstancesConfiguration2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2NotFound creates a GetInstancesConfiguration2NotFound with default headers values
func NewGetInstancesConfiguration2NotFound() *GetInstancesConfiguration2NotFound {
	return &GetInstancesConfiguration2NotFound{}
}

/*GetInstancesConfiguration2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesConfiguration2NotFound struct {
}

func (o *GetInstancesConfiguration2NotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2NotFound ", 404)
}

func (o *GetInstancesConfiguration2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2MethodNotAllowed creates a GetInstancesConfiguration2MethodNotAllowed with default headers values
func NewGetInstancesConfiguration2MethodNotAllowed() *GetInstancesConfiguration2MethodNotAllowed {
	return &GetInstancesConfiguration2MethodNotAllowed{}
}

/*GetInstancesConfiguration2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesConfiguration2MethodNotAllowed struct {
}

func (o *GetInstancesConfiguration2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2MethodNotAllowed ", 405)
}

func (o *GetInstancesConfiguration2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2NotAcceptable creates a GetInstancesConfiguration2NotAcceptable with default headers values
func NewGetInstancesConfiguration2NotAcceptable() *GetInstancesConfiguration2NotAcceptable {
	return &GetInstancesConfiguration2NotAcceptable{}
}

/*GetInstancesConfiguration2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesConfiguration2NotAcceptable struct {
}

func (o *GetInstancesConfiguration2NotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2NotAcceptable ", 406)
}

func (o *GetInstancesConfiguration2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2Conflict creates a GetInstancesConfiguration2Conflict with default headers values
func NewGetInstancesConfiguration2Conflict() *GetInstancesConfiguration2Conflict {
	return &GetInstancesConfiguration2Conflict{}
}

/*GetInstancesConfiguration2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesConfiguration2Conflict struct {
}

func (o *GetInstancesConfiguration2Conflict) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2Conflict ", 409)
}

func (o *GetInstancesConfiguration2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2UnsupportedMediaType creates a GetInstancesConfiguration2UnsupportedMediaType with default headers values
func NewGetInstancesConfiguration2UnsupportedMediaType() *GetInstancesConfiguration2UnsupportedMediaType {
	return &GetInstancesConfiguration2UnsupportedMediaType{}
}

/*GetInstancesConfiguration2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesConfiguration2UnsupportedMediaType struct {
}

func (o *GetInstancesConfiguration2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2UnsupportedMediaType ", 415)
}

func (o *GetInstancesConfiguration2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfiguration2InternalServerError creates a GetInstancesConfiguration2InternalServerError with default headers values
func NewGetInstancesConfiguration2InternalServerError() *GetInstancesConfiguration2InternalServerError {
	return &GetInstancesConfiguration2InternalServerError{}
}

/*GetInstancesConfiguration2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesConfiguration2InternalServerError struct {
}

func (o *GetInstancesConfiguration2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/configuration][%d] getInstancesConfiguration2InternalServerError ", 500)
}

func (o *GetInstancesConfiguration2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
