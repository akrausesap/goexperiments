// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteInstancesTraceLogging2Reader is a Reader for the DeleteInstancesTraceLogging2 structure.
type DeleteInstancesTraceLogging2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstancesTraceLogging2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteInstancesTraceLogging2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteInstancesTraceLogging2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteInstancesTraceLogging2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteInstancesTraceLogging2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteInstancesTraceLogging2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDeleteInstancesTraceLogging2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewDeleteInstancesTraceLogging2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteInstancesTraceLogging2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewDeleteInstancesTraceLogging2UnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteInstancesTraceLogging2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteInstancesTraceLogging2OK creates a DeleteInstancesTraceLogging2OK with default headers values
func NewDeleteInstancesTraceLogging2OK() *DeleteInstancesTraceLogging2OK {
	return &DeleteInstancesTraceLogging2OK{}
}

/*DeleteInstancesTraceLogging2OK handles this case with default header values.

OK - Everything worked as expected
*/
type DeleteInstancesTraceLogging2OK struct {
}

func (o *DeleteInstancesTraceLogging2OK) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2OK ", 200)
}

func (o *DeleteInstancesTraceLogging2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2BadRequest creates a DeleteInstancesTraceLogging2BadRequest with default headers values
func NewDeleteInstancesTraceLogging2BadRequest() *DeleteInstancesTraceLogging2BadRequest {
	return &DeleteInstancesTraceLogging2BadRequest{}
}

/*DeleteInstancesTraceLogging2BadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type DeleteInstancesTraceLogging2BadRequest struct {
}

func (o *DeleteInstancesTraceLogging2BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2BadRequest ", 400)
}

func (o *DeleteInstancesTraceLogging2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2Unauthorized creates a DeleteInstancesTraceLogging2Unauthorized with default headers values
func NewDeleteInstancesTraceLogging2Unauthorized() *DeleteInstancesTraceLogging2Unauthorized {
	return &DeleteInstancesTraceLogging2Unauthorized{}
}

/*DeleteInstancesTraceLogging2Unauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type DeleteInstancesTraceLogging2Unauthorized struct {
}

func (o *DeleteInstancesTraceLogging2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2Unauthorized ", 401)
}

func (o *DeleteInstancesTraceLogging2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2Forbidden creates a DeleteInstancesTraceLogging2Forbidden with default headers values
func NewDeleteInstancesTraceLogging2Forbidden() *DeleteInstancesTraceLogging2Forbidden {
	return &DeleteInstancesTraceLogging2Forbidden{}
}

/*DeleteInstancesTraceLogging2Forbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type DeleteInstancesTraceLogging2Forbidden struct {
}

func (o *DeleteInstancesTraceLogging2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2Forbidden ", 403)
}

func (o *DeleteInstancesTraceLogging2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2NotFound creates a DeleteInstancesTraceLogging2NotFound with default headers values
func NewDeleteInstancesTraceLogging2NotFound() *DeleteInstancesTraceLogging2NotFound {
	return &DeleteInstancesTraceLogging2NotFound{}
}

/*DeleteInstancesTraceLogging2NotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type DeleteInstancesTraceLogging2NotFound struct {
}

func (o *DeleteInstancesTraceLogging2NotFound) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2NotFound ", 404)
}

func (o *DeleteInstancesTraceLogging2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2MethodNotAllowed creates a DeleteInstancesTraceLogging2MethodNotAllowed with default headers values
func NewDeleteInstancesTraceLogging2MethodNotAllowed() *DeleteInstancesTraceLogging2MethodNotAllowed {
	return &DeleteInstancesTraceLogging2MethodNotAllowed{}
}

/*DeleteInstancesTraceLogging2MethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type DeleteInstancesTraceLogging2MethodNotAllowed struct {
}

func (o *DeleteInstancesTraceLogging2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2MethodNotAllowed ", 405)
}

func (o *DeleteInstancesTraceLogging2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2NotAcceptable creates a DeleteInstancesTraceLogging2NotAcceptable with default headers values
func NewDeleteInstancesTraceLogging2NotAcceptable() *DeleteInstancesTraceLogging2NotAcceptable {
	return &DeleteInstancesTraceLogging2NotAcceptable{}
}

/*DeleteInstancesTraceLogging2NotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type DeleteInstancesTraceLogging2NotAcceptable struct {
}

func (o *DeleteInstancesTraceLogging2NotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2NotAcceptable ", 406)
}

func (o *DeleteInstancesTraceLogging2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2Conflict creates a DeleteInstancesTraceLogging2Conflict with default headers values
func NewDeleteInstancesTraceLogging2Conflict() *DeleteInstancesTraceLogging2Conflict {
	return &DeleteInstancesTraceLogging2Conflict{}
}

/*DeleteInstancesTraceLogging2Conflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type DeleteInstancesTraceLogging2Conflict struct {
}

func (o *DeleteInstancesTraceLogging2Conflict) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2Conflict ", 409)
}

func (o *DeleteInstancesTraceLogging2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2UnsupportedMediaType creates a DeleteInstancesTraceLogging2UnsupportedMediaType with default headers values
func NewDeleteInstancesTraceLogging2UnsupportedMediaType() *DeleteInstancesTraceLogging2UnsupportedMediaType {
	return &DeleteInstancesTraceLogging2UnsupportedMediaType{}
}

/*DeleteInstancesTraceLogging2UnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type DeleteInstancesTraceLogging2UnsupportedMediaType struct {
}

func (o *DeleteInstancesTraceLogging2UnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2UnsupportedMediaType ", 415)
}

func (o *DeleteInstancesTraceLogging2UnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLogging2InternalServerError creates a DeleteInstancesTraceLogging2InternalServerError with default headers values
func NewDeleteInstancesTraceLogging2InternalServerError() *DeleteInstancesTraceLogging2InternalServerError {
	return &DeleteInstancesTraceLogging2InternalServerError{}
}

/*DeleteInstancesTraceLogging2InternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type DeleteInstancesTraceLogging2InternalServerError struct {
}

func (o *DeleteInstancesTraceLogging2InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}/trace-logging][%d] deleteInstancesTraceLogging2InternalServerError ", 500)
}

func (o *DeleteInstancesTraceLogging2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}