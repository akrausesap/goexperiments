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

// UpdateInstancesReader is a Reader for the UpdateInstances structure.
type UpdateInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewUpdateInstancesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewUpdateInstancesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateInstancesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewUpdateInstancesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateInstancesOK creates a UpdateInstancesOK with default headers values
func NewUpdateInstancesOK() *UpdateInstancesOK {
	return &UpdateInstancesOK{}
}

/*UpdateInstancesOK handles this case with default header values.

OK - Everything worked as expected
*/
type UpdateInstancesOK struct {
	Payload *models.ElementInstance
}

func (o *UpdateInstancesOK) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesOK  %+v", 200, o.Payload)
}

func (o *UpdateInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElementInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstancesBadRequest creates a UpdateInstancesBadRequest with default headers values
func NewUpdateInstancesBadRequest() *UpdateInstancesBadRequest {
	return &UpdateInstancesBadRequest{}
}

/*UpdateInstancesBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type UpdateInstancesBadRequest struct {
}

func (o *UpdateInstancesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesBadRequest ", 400)
}

func (o *UpdateInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesUnauthorized creates a UpdateInstancesUnauthorized with default headers values
func NewUpdateInstancesUnauthorized() *UpdateInstancesUnauthorized {
	return &UpdateInstancesUnauthorized{}
}

/*UpdateInstancesUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type UpdateInstancesUnauthorized struct {
}

func (o *UpdateInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesUnauthorized ", 401)
}

func (o *UpdateInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesForbidden creates a UpdateInstancesForbidden with default headers values
func NewUpdateInstancesForbidden() *UpdateInstancesForbidden {
	return &UpdateInstancesForbidden{}
}

/*UpdateInstancesForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type UpdateInstancesForbidden struct {
}

func (o *UpdateInstancesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesForbidden ", 403)
}

func (o *UpdateInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesNotFound creates a UpdateInstancesNotFound with default headers values
func NewUpdateInstancesNotFound() *UpdateInstancesNotFound {
	return &UpdateInstancesNotFound{}
}

/*UpdateInstancesNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type UpdateInstancesNotFound struct {
}

func (o *UpdateInstancesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesNotFound ", 404)
}

func (o *UpdateInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesMethodNotAllowed creates a UpdateInstancesMethodNotAllowed with default headers values
func NewUpdateInstancesMethodNotAllowed() *UpdateInstancesMethodNotAllowed {
	return &UpdateInstancesMethodNotAllowed{}
}

/*UpdateInstancesMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type UpdateInstancesMethodNotAllowed struct {
}

func (o *UpdateInstancesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesMethodNotAllowed ", 405)
}

func (o *UpdateInstancesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesNotAcceptable creates a UpdateInstancesNotAcceptable with default headers values
func NewUpdateInstancesNotAcceptable() *UpdateInstancesNotAcceptable {
	return &UpdateInstancesNotAcceptable{}
}

/*UpdateInstancesNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type UpdateInstancesNotAcceptable struct {
}

func (o *UpdateInstancesNotAcceptable) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesNotAcceptable ", 406)
}

func (o *UpdateInstancesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConflict creates a UpdateInstancesConflict with default headers values
func NewUpdateInstancesConflict() *UpdateInstancesConflict {
	return &UpdateInstancesConflict{}
}

/*UpdateInstancesConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type UpdateInstancesConflict struct {
}

func (o *UpdateInstancesConflict) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesConflict ", 409)
}

func (o *UpdateInstancesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesUnsupportedMediaType creates a UpdateInstancesUnsupportedMediaType with default headers values
func NewUpdateInstancesUnsupportedMediaType() *UpdateInstancesUnsupportedMediaType {
	return &UpdateInstancesUnsupportedMediaType{}
}

/*UpdateInstancesUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type UpdateInstancesUnsupportedMediaType struct {
}

func (o *UpdateInstancesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesUnsupportedMediaType ", 415)
}

func (o *UpdateInstancesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesInternalServerError creates a UpdateInstancesInternalServerError with default headers values
func NewUpdateInstancesInternalServerError() *UpdateInstancesInternalServerError {
	return &UpdateInstancesInternalServerError{}
}

/*UpdateInstancesInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type UpdateInstancesInternalServerError struct {
}

func (o *UpdateInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /instances][%d] updateInstancesInternalServerError ", 500)
}

func (o *UpdateInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
