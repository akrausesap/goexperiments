// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteInstanceByIDReader is a Reader for the DeleteInstanceByID structure.
type DeleteInstanceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstanceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteInstanceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteInstanceByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteInstanceByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteInstanceByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteInstanceByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDeleteInstanceByIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewDeleteInstanceByIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteInstanceByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewDeleteInstanceByIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteInstanceByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteInstanceByIDOK creates a DeleteInstanceByIDOK with default headers values
func NewDeleteInstanceByIDOK() *DeleteInstanceByIDOK {
	return &DeleteInstanceByIDOK{}
}

/*DeleteInstanceByIDOK handles this case with default header values.

OK - Everything worked as expected
*/
type DeleteInstanceByIDOK struct {
}

func (o *DeleteInstanceByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdOK ", 200)
}

func (o *DeleteInstanceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDBadRequest creates a DeleteInstanceByIDBadRequest with default headers values
func NewDeleteInstanceByIDBadRequest() *DeleteInstanceByIDBadRequest {
	return &DeleteInstanceByIDBadRequest{}
}

/*DeleteInstanceByIDBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type DeleteInstanceByIDBadRequest struct {
}

func (o *DeleteInstanceByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdBadRequest ", 400)
}

func (o *DeleteInstanceByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDUnauthorized creates a DeleteInstanceByIDUnauthorized with default headers values
func NewDeleteInstanceByIDUnauthorized() *DeleteInstanceByIDUnauthorized {
	return &DeleteInstanceByIDUnauthorized{}
}

/*DeleteInstanceByIDUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type DeleteInstanceByIDUnauthorized struct {
}

func (o *DeleteInstanceByIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdUnauthorized ", 401)
}

func (o *DeleteInstanceByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDForbidden creates a DeleteInstanceByIDForbidden with default headers values
func NewDeleteInstanceByIDForbidden() *DeleteInstanceByIDForbidden {
	return &DeleteInstanceByIDForbidden{}
}

/*DeleteInstanceByIDForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type DeleteInstanceByIDForbidden struct {
}

func (o *DeleteInstanceByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdForbidden ", 403)
}

func (o *DeleteInstanceByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDNotFound creates a DeleteInstanceByIDNotFound with default headers values
func NewDeleteInstanceByIDNotFound() *DeleteInstanceByIDNotFound {
	return &DeleteInstanceByIDNotFound{}
}

/*DeleteInstanceByIDNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type DeleteInstanceByIDNotFound struct {
}

func (o *DeleteInstanceByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdNotFound ", 404)
}

func (o *DeleteInstanceByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDMethodNotAllowed creates a DeleteInstanceByIDMethodNotAllowed with default headers values
func NewDeleteInstanceByIDMethodNotAllowed() *DeleteInstanceByIDMethodNotAllowed {
	return &DeleteInstanceByIDMethodNotAllowed{}
}

/*DeleteInstanceByIDMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type DeleteInstanceByIDMethodNotAllowed struct {
}

func (o *DeleteInstanceByIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdMethodNotAllowed ", 405)
}

func (o *DeleteInstanceByIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDNotAcceptable creates a DeleteInstanceByIDNotAcceptable with default headers values
func NewDeleteInstanceByIDNotAcceptable() *DeleteInstanceByIDNotAcceptable {
	return &DeleteInstanceByIDNotAcceptable{}
}

/*DeleteInstanceByIDNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type DeleteInstanceByIDNotAcceptable struct {
}

func (o *DeleteInstanceByIDNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdNotAcceptable ", 406)
}

func (o *DeleteInstanceByIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDConflict creates a DeleteInstanceByIDConflict with default headers values
func NewDeleteInstanceByIDConflict() *DeleteInstanceByIDConflict {
	return &DeleteInstanceByIDConflict{}
}

/*DeleteInstanceByIDConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type DeleteInstanceByIDConflict struct {
}

func (o *DeleteInstanceByIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdConflict ", 409)
}

func (o *DeleteInstanceByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDUnsupportedMediaType creates a DeleteInstanceByIDUnsupportedMediaType with default headers values
func NewDeleteInstanceByIDUnsupportedMediaType() *DeleteInstanceByIDUnsupportedMediaType {
	return &DeleteInstanceByIDUnsupportedMediaType{}
}

/*DeleteInstanceByIDUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type DeleteInstanceByIDUnsupportedMediaType struct {
}

func (o *DeleteInstanceByIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdUnsupportedMediaType ", 415)
}

func (o *DeleteInstanceByIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceByIDInternalServerError creates a DeleteInstanceByIDInternalServerError with default headers values
func NewDeleteInstanceByIDInternalServerError() *DeleteInstanceByIDInternalServerError {
	return &DeleteInstanceByIDInternalServerError{}
}

/*DeleteInstanceByIDInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type DeleteInstanceByIDInternalServerError struct {
}

func (o *DeleteInstanceByIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstanceByIdInternalServerError ", 500)
}

func (o *DeleteInstanceByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}