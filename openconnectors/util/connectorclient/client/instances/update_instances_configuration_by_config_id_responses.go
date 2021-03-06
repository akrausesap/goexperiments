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

// UpdateInstancesConfigurationByConfigIDReader is a Reader for the UpdateInstancesConfigurationByConfigID structure.
type UpdateInstancesConfigurationByConfigIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstancesConfigurationByConfigIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateInstancesConfigurationByConfigIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateInstancesConfigurationByConfigIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateInstancesConfigurationByConfigIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateInstancesConfigurationByConfigIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateInstancesConfigurationByConfigIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewUpdateInstancesConfigurationByConfigIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewUpdateInstancesConfigurationByConfigIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateInstancesConfigurationByConfigIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewUpdateInstancesConfigurationByConfigIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateInstancesConfigurationByConfigIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateInstancesConfigurationByConfigIDOK creates a UpdateInstancesConfigurationByConfigIDOK with default headers values
func NewUpdateInstancesConfigurationByConfigIDOK() *UpdateInstancesConfigurationByConfigIDOK {
	return &UpdateInstancesConfigurationByConfigIDOK{}
}

/*UpdateInstancesConfigurationByConfigIDOK handles this case with default header values.

OK - Everything worked as expected
*/
type UpdateInstancesConfigurationByConfigIDOK struct {
	Payload *models.ElementInstanceConfig
}

func (o *UpdateInstancesConfigurationByConfigIDOK) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdOK  %+v", 200, o.Payload)
}

func (o *UpdateInstancesConfigurationByConfigIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElementInstanceConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDBadRequest creates a UpdateInstancesConfigurationByConfigIDBadRequest with default headers values
func NewUpdateInstancesConfigurationByConfigIDBadRequest() *UpdateInstancesConfigurationByConfigIDBadRequest {
	return &UpdateInstancesConfigurationByConfigIDBadRequest{}
}

/*UpdateInstancesConfigurationByConfigIDBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type UpdateInstancesConfigurationByConfigIDBadRequest struct {
}

func (o *UpdateInstancesConfigurationByConfigIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdBadRequest ", 400)
}

func (o *UpdateInstancesConfigurationByConfigIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDUnauthorized creates a UpdateInstancesConfigurationByConfigIDUnauthorized with default headers values
func NewUpdateInstancesConfigurationByConfigIDUnauthorized() *UpdateInstancesConfigurationByConfigIDUnauthorized {
	return &UpdateInstancesConfigurationByConfigIDUnauthorized{}
}

/*UpdateInstancesConfigurationByConfigIDUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type UpdateInstancesConfigurationByConfigIDUnauthorized struct {
}

func (o *UpdateInstancesConfigurationByConfigIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdUnauthorized ", 401)
}

func (o *UpdateInstancesConfigurationByConfigIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDForbidden creates a UpdateInstancesConfigurationByConfigIDForbidden with default headers values
func NewUpdateInstancesConfigurationByConfigIDForbidden() *UpdateInstancesConfigurationByConfigIDForbidden {
	return &UpdateInstancesConfigurationByConfigIDForbidden{}
}

/*UpdateInstancesConfigurationByConfigIDForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type UpdateInstancesConfigurationByConfigIDForbidden struct {
}

func (o *UpdateInstancesConfigurationByConfigIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdForbidden ", 403)
}

func (o *UpdateInstancesConfigurationByConfigIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDNotFound creates a UpdateInstancesConfigurationByConfigIDNotFound with default headers values
func NewUpdateInstancesConfigurationByConfigIDNotFound() *UpdateInstancesConfigurationByConfigIDNotFound {
	return &UpdateInstancesConfigurationByConfigIDNotFound{}
}

/*UpdateInstancesConfigurationByConfigIDNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type UpdateInstancesConfigurationByConfigIDNotFound struct {
}

func (o *UpdateInstancesConfigurationByConfigIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdNotFound ", 404)
}

func (o *UpdateInstancesConfigurationByConfigIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDMethodNotAllowed creates a UpdateInstancesConfigurationByConfigIDMethodNotAllowed with default headers values
func NewUpdateInstancesConfigurationByConfigIDMethodNotAllowed() *UpdateInstancesConfigurationByConfigIDMethodNotAllowed {
	return &UpdateInstancesConfigurationByConfigIDMethodNotAllowed{}
}

/*UpdateInstancesConfigurationByConfigIDMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type UpdateInstancesConfigurationByConfigIDMethodNotAllowed struct {
}

func (o *UpdateInstancesConfigurationByConfigIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdMethodNotAllowed ", 405)
}

func (o *UpdateInstancesConfigurationByConfigIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDNotAcceptable creates a UpdateInstancesConfigurationByConfigIDNotAcceptable with default headers values
func NewUpdateInstancesConfigurationByConfigIDNotAcceptable() *UpdateInstancesConfigurationByConfigIDNotAcceptable {
	return &UpdateInstancesConfigurationByConfigIDNotAcceptable{}
}

/*UpdateInstancesConfigurationByConfigIDNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type UpdateInstancesConfigurationByConfigIDNotAcceptable struct {
}

func (o *UpdateInstancesConfigurationByConfigIDNotAcceptable) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdNotAcceptable ", 406)
}

func (o *UpdateInstancesConfigurationByConfigIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDConflict creates a UpdateInstancesConfigurationByConfigIDConflict with default headers values
func NewUpdateInstancesConfigurationByConfigIDConflict() *UpdateInstancesConfigurationByConfigIDConflict {
	return &UpdateInstancesConfigurationByConfigIDConflict{}
}

/*UpdateInstancesConfigurationByConfigIDConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type UpdateInstancesConfigurationByConfigIDConflict struct {
}

func (o *UpdateInstancesConfigurationByConfigIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdConflict ", 409)
}

func (o *UpdateInstancesConfigurationByConfigIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDUnsupportedMediaType creates a UpdateInstancesConfigurationByConfigIDUnsupportedMediaType with default headers values
func NewUpdateInstancesConfigurationByConfigIDUnsupportedMediaType() *UpdateInstancesConfigurationByConfigIDUnsupportedMediaType {
	return &UpdateInstancesConfigurationByConfigIDUnsupportedMediaType{}
}

/*UpdateInstancesConfigurationByConfigIDUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type UpdateInstancesConfigurationByConfigIDUnsupportedMediaType struct {
}

func (o *UpdateInstancesConfigurationByConfigIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdUnsupportedMediaType ", 415)
}

func (o *UpdateInstancesConfigurationByConfigIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstancesConfigurationByConfigIDInternalServerError creates a UpdateInstancesConfigurationByConfigIDInternalServerError with default headers values
func NewUpdateInstancesConfigurationByConfigIDInternalServerError() *UpdateInstancesConfigurationByConfigIDInternalServerError {
	return &UpdateInstancesConfigurationByConfigIDInternalServerError{}
}

/*UpdateInstancesConfigurationByConfigIDInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type UpdateInstancesConfigurationByConfigIDInternalServerError struct {
}

func (o *UpdateInstancesConfigurationByConfigIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /instances/configuration/{configId}][%d] updateInstancesConfigurationByConfigIdInternalServerError ", 500)
}

func (o *UpdateInstancesConfigurationByConfigIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
