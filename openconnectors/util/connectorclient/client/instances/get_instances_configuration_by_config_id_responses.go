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

// GetInstancesConfigurationByConfigIDReader is a Reader for the GetInstancesConfigurationByConfigID structure.
type GetInstancesConfigurationByConfigIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesConfigurationByConfigIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesConfigurationByConfigIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesConfigurationByConfigIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesConfigurationByConfigIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesConfigurationByConfigIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesConfigurationByConfigIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesConfigurationByConfigIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesConfigurationByConfigIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesConfigurationByConfigIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesConfigurationByConfigIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesConfigurationByConfigIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesConfigurationByConfigIDOK creates a GetInstancesConfigurationByConfigIDOK with default headers values
func NewGetInstancesConfigurationByConfigIDOK() *GetInstancesConfigurationByConfigIDOK {
	return &GetInstancesConfigurationByConfigIDOK{}
}

/*GetInstancesConfigurationByConfigIDOK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesConfigurationByConfigIDOK struct {
	Payload *models.ElementInstanceConfig
}

func (o *GetInstancesConfigurationByConfigIDOK) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdOK  %+v", 200, o.Payload)
}

func (o *GetInstancesConfigurationByConfigIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElementInstanceConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesConfigurationByConfigIDBadRequest creates a GetInstancesConfigurationByConfigIDBadRequest with default headers values
func NewGetInstancesConfigurationByConfigIDBadRequest() *GetInstancesConfigurationByConfigIDBadRequest {
	return &GetInstancesConfigurationByConfigIDBadRequest{}
}

/*GetInstancesConfigurationByConfigIDBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesConfigurationByConfigIDBadRequest struct {
}

func (o *GetInstancesConfigurationByConfigIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdBadRequest ", 400)
}

func (o *GetInstancesConfigurationByConfigIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDUnauthorized creates a GetInstancesConfigurationByConfigIDUnauthorized with default headers values
func NewGetInstancesConfigurationByConfigIDUnauthorized() *GetInstancesConfigurationByConfigIDUnauthorized {
	return &GetInstancesConfigurationByConfigIDUnauthorized{}
}

/*GetInstancesConfigurationByConfigIDUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesConfigurationByConfigIDUnauthorized struct {
}

func (o *GetInstancesConfigurationByConfigIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdUnauthorized ", 401)
}

func (o *GetInstancesConfigurationByConfigIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDForbidden creates a GetInstancesConfigurationByConfigIDForbidden with default headers values
func NewGetInstancesConfigurationByConfigIDForbidden() *GetInstancesConfigurationByConfigIDForbidden {
	return &GetInstancesConfigurationByConfigIDForbidden{}
}

/*GetInstancesConfigurationByConfigIDForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesConfigurationByConfigIDForbidden struct {
}

func (o *GetInstancesConfigurationByConfigIDForbidden) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdForbidden ", 403)
}

func (o *GetInstancesConfigurationByConfigIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDNotFound creates a GetInstancesConfigurationByConfigIDNotFound with default headers values
func NewGetInstancesConfigurationByConfigIDNotFound() *GetInstancesConfigurationByConfigIDNotFound {
	return &GetInstancesConfigurationByConfigIDNotFound{}
}

/*GetInstancesConfigurationByConfigIDNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesConfigurationByConfigIDNotFound struct {
}

func (o *GetInstancesConfigurationByConfigIDNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdNotFound ", 404)
}

func (o *GetInstancesConfigurationByConfigIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDMethodNotAllowed creates a GetInstancesConfigurationByConfigIDMethodNotAllowed with default headers values
func NewGetInstancesConfigurationByConfigIDMethodNotAllowed() *GetInstancesConfigurationByConfigIDMethodNotAllowed {
	return &GetInstancesConfigurationByConfigIDMethodNotAllowed{}
}

/*GetInstancesConfigurationByConfigIDMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesConfigurationByConfigIDMethodNotAllowed struct {
}

func (o *GetInstancesConfigurationByConfigIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdMethodNotAllowed ", 405)
}

func (o *GetInstancesConfigurationByConfigIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDNotAcceptable creates a GetInstancesConfigurationByConfigIDNotAcceptable with default headers values
func NewGetInstancesConfigurationByConfigIDNotAcceptable() *GetInstancesConfigurationByConfigIDNotAcceptable {
	return &GetInstancesConfigurationByConfigIDNotAcceptable{}
}

/*GetInstancesConfigurationByConfigIDNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesConfigurationByConfigIDNotAcceptable struct {
}

func (o *GetInstancesConfigurationByConfigIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdNotAcceptable ", 406)
}

func (o *GetInstancesConfigurationByConfigIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDConflict creates a GetInstancesConfigurationByConfigIDConflict with default headers values
func NewGetInstancesConfigurationByConfigIDConflict() *GetInstancesConfigurationByConfigIDConflict {
	return &GetInstancesConfigurationByConfigIDConflict{}
}

/*GetInstancesConfigurationByConfigIDConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesConfigurationByConfigIDConflict struct {
}

func (o *GetInstancesConfigurationByConfigIDConflict) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdConflict ", 409)
}

func (o *GetInstancesConfigurationByConfigIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDUnsupportedMediaType creates a GetInstancesConfigurationByConfigIDUnsupportedMediaType with default headers values
func NewGetInstancesConfigurationByConfigIDUnsupportedMediaType() *GetInstancesConfigurationByConfigIDUnsupportedMediaType {
	return &GetInstancesConfigurationByConfigIDUnsupportedMediaType{}
}

/*GetInstancesConfigurationByConfigIDUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesConfigurationByConfigIDUnsupportedMediaType struct {
}

func (o *GetInstancesConfigurationByConfigIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdUnsupportedMediaType ", 415)
}

func (o *GetInstancesConfigurationByConfigIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesConfigurationByConfigIDInternalServerError creates a GetInstancesConfigurationByConfigIDInternalServerError with default headers values
func NewGetInstancesConfigurationByConfigIDInternalServerError() *GetInstancesConfigurationByConfigIDInternalServerError {
	return &GetInstancesConfigurationByConfigIDInternalServerError{}
}

/*GetInstancesConfigurationByConfigIDInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesConfigurationByConfigIDInternalServerError struct {
}

func (o *GetInstancesConfigurationByConfigIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/configuration/{configId}][%d] getInstancesConfigurationByConfigIdInternalServerError ", 500)
}

func (o *GetInstancesConfigurationByConfigIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
