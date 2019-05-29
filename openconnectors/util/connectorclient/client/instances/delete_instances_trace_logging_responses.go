// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteInstancesTraceLoggingReader is a Reader for the DeleteInstancesTraceLogging structure.
type DeleteInstancesTraceLoggingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstancesTraceLoggingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteInstancesTraceLoggingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteInstancesTraceLoggingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteInstancesTraceLoggingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteInstancesTraceLoggingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteInstancesTraceLoggingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDeleteInstancesTraceLoggingMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewDeleteInstancesTraceLoggingNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteInstancesTraceLoggingConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewDeleteInstancesTraceLoggingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteInstancesTraceLoggingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteInstancesTraceLoggingOK creates a DeleteInstancesTraceLoggingOK with default headers values
func NewDeleteInstancesTraceLoggingOK() *DeleteInstancesTraceLoggingOK {
	return &DeleteInstancesTraceLoggingOK{}
}

/*DeleteInstancesTraceLoggingOK handles this case with default header values.

OK - Everything worked as expected
*/
type DeleteInstancesTraceLoggingOK struct {
}

func (o *DeleteInstancesTraceLoggingOK) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingOK ", 200)
}

func (o *DeleteInstancesTraceLoggingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingBadRequest creates a DeleteInstancesTraceLoggingBadRequest with default headers values
func NewDeleteInstancesTraceLoggingBadRequest() *DeleteInstancesTraceLoggingBadRequest {
	return &DeleteInstancesTraceLoggingBadRequest{}
}

/*DeleteInstancesTraceLoggingBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type DeleteInstancesTraceLoggingBadRequest struct {
}

func (o *DeleteInstancesTraceLoggingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingBadRequest ", 400)
}

func (o *DeleteInstancesTraceLoggingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingUnauthorized creates a DeleteInstancesTraceLoggingUnauthorized with default headers values
func NewDeleteInstancesTraceLoggingUnauthorized() *DeleteInstancesTraceLoggingUnauthorized {
	return &DeleteInstancesTraceLoggingUnauthorized{}
}

/*DeleteInstancesTraceLoggingUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type DeleteInstancesTraceLoggingUnauthorized struct {
}

func (o *DeleteInstancesTraceLoggingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingUnauthorized ", 401)
}

func (o *DeleteInstancesTraceLoggingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingForbidden creates a DeleteInstancesTraceLoggingForbidden with default headers values
func NewDeleteInstancesTraceLoggingForbidden() *DeleteInstancesTraceLoggingForbidden {
	return &DeleteInstancesTraceLoggingForbidden{}
}

/*DeleteInstancesTraceLoggingForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type DeleteInstancesTraceLoggingForbidden struct {
}

func (o *DeleteInstancesTraceLoggingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingForbidden ", 403)
}

func (o *DeleteInstancesTraceLoggingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingNotFound creates a DeleteInstancesTraceLoggingNotFound with default headers values
func NewDeleteInstancesTraceLoggingNotFound() *DeleteInstancesTraceLoggingNotFound {
	return &DeleteInstancesTraceLoggingNotFound{}
}

/*DeleteInstancesTraceLoggingNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type DeleteInstancesTraceLoggingNotFound struct {
}

func (o *DeleteInstancesTraceLoggingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingNotFound ", 404)
}

func (o *DeleteInstancesTraceLoggingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingMethodNotAllowed creates a DeleteInstancesTraceLoggingMethodNotAllowed with default headers values
func NewDeleteInstancesTraceLoggingMethodNotAllowed() *DeleteInstancesTraceLoggingMethodNotAllowed {
	return &DeleteInstancesTraceLoggingMethodNotAllowed{}
}

/*DeleteInstancesTraceLoggingMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type DeleteInstancesTraceLoggingMethodNotAllowed struct {
}

func (o *DeleteInstancesTraceLoggingMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingMethodNotAllowed ", 405)
}

func (o *DeleteInstancesTraceLoggingMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingNotAcceptable creates a DeleteInstancesTraceLoggingNotAcceptable with default headers values
func NewDeleteInstancesTraceLoggingNotAcceptable() *DeleteInstancesTraceLoggingNotAcceptable {
	return &DeleteInstancesTraceLoggingNotAcceptable{}
}

/*DeleteInstancesTraceLoggingNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type DeleteInstancesTraceLoggingNotAcceptable struct {
}

func (o *DeleteInstancesTraceLoggingNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingNotAcceptable ", 406)
}

func (o *DeleteInstancesTraceLoggingNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingConflict creates a DeleteInstancesTraceLoggingConflict with default headers values
func NewDeleteInstancesTraceLoggingConflict() *DeleteInstancesTraceLoggingConflict {
	return &DeleteInstancesTraceLoggingConflict{}
}

/*DeleteInstancesTraceLoggingConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type DeleteInstancesTraceLoggingConflict struct {
}

func (o *DeleteInstancesTraceLoggingConflict) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingConflict ", 409)
}

func (o *DeleteInstancesTraceLoggingConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingUnsupportedMediaType creates a DeleteInstancesTraceLoggingUnsupportedMediaType with default headers values
func NewDeleteInstancesTraceLoggingUnsupportedMediaType() *DeleteInstancesTraceLoggingUnsupportedMediaType {
	return &DeleteInstancesTraceLoggingUnsupportedMediaType{}
}

/*DeleteInstancesTraceLoggingUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type DeleteInstancesTraceLoggingUnsupportedMediaType struct {
}

func (o *DeleteInstancesTraceLoggingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingUnsupportedMediaType ", 415)
}

func (o *DeleteInstancesTraceLoggingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstancesTraceLoggingInternalServerError creates a DeleteInstancesTraceLoggingInternalServerError with default headers values
func NewDeleteInstancesTraceLoggingInternalServerError() *DeleteInstancesTraceLoggingInternalServerError {
	return &DeleteInstancesTraceLoggingInternalServerError{}
}

/*DeleteInstancesTraceLoggingInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type DeleteInstancesTraceLoggingInternalServerError struct {
}

func (o *DeleteInstancesTraceLoggingInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /instances/trace-logging][%d] deleteInstancesTraceLoggingInternalServerError ", 500)
}

func (o *DeleteInstancesTraceLoggingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}