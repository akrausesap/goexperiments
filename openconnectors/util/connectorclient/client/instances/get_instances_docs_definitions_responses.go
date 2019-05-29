// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetInstancesDocsDefinitionsReader is a Reader for the GetInstancesDocsDefinitions structure.
type GetInstancesDocsDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesDocsDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesDocsDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstancesDocsDefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstancesDocsDefinitionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstancesDocsDefinitionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstancesDocsDefinitionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstancesDocsDefinitionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetInstancesDocsDefinitionsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInstancesDocsDefinitionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetInstancesDocsDefinitionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstancesDocsDefinitionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesDocsDefinitionsOK creates a GetInstancesDocsDefinitionsOK with default headers values
func NewGetInstancesDocsDefinitionsOK() *GetInstancesDocsDefinitionsOK {
	return &GetInstancesDocsDefinitionsOK{}
}

/*GetInstancesDocsDefinitionsOK handles this case with default header values.

OK - Everything worked as expected
*/
type GetInstancesDocsDefinitionsOK struct {
	Payload interface{}
}

func (o *GetInstancesDocsDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetInstancesDocsDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesDocsDefinitionsBadRequest creates a GetInstancesDocsDefinitionsBadRequest with default headers values
func NewGetInstancesDocsDefinitionsBadRequest() *GetInstancesDocsDefinitionsBadRequest {
	return &GetInstancesDocsDefinitionsBadRequest{}
}

/*GetInstancesDocsDefinitionsBadRequest handles this case with default header values.

Bad Request - Often due to a missing request parameter
*/
type GetInstancesDocsDefinitionsBadRequest struct {
}

func (o *GetInstancesDocsDefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsBadRequest ", 400)
}

func (o *GetInstancesDocsDefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsUnauthorized creates a GetInstancesDocsDefinitionsUnauthorized with default headers values
func NewGetInstancesDocsDefinitionsUnauthorized() *GetInstancesDocsDefinitionsUnauthorized {
	return &GetInstancesDocsDefinitionsUnauthorized{}
}

/*GetInstancesDocsDefinitionsUnauthorized handles this case with default header values.

Unauthorized - An invalid element token, user secret and/or org secret provided
*/
type GetInstancesDocsDefinitionsUnauthorized struct {
}

func (o *GetInstancesDocsDefinitionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsUnauthorized ", 401)
}

func (o *GetInstancesDocsDefinitionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsForbidden creates a GetInstancesDocsDefinitionsForbidden with default headers values
func NewGetInstancesDocsDefinitionsForbidden() *GetInstancesDocsDefinitionsForbidden {
	return &GetInstancesDocsDefinitionsForbidden{}
}

/*GetInstancesDocsDefinitionsForbidden handles this case with default header values.

Forbidden - Access to the resource by the provider is forbidden
*/
type GetInstancesDocsDefinitionsForbidden struct {
}

func (o *GetInstancesDocsDefinitionsForbidden) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsForbidden ", 403)
}

func (o *GetInstancesDocsDefinitionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsNotFound creates a GetInstancesDocsDefinitionsNotFound with default headers values
func NewGetInstancesDocsDefinitionsNotFound() *GetInstancesDocsDefinitionsNotFound {
	return &GetInstancesDocsDefinitionsNotFound{}
}

/*GetInstancesDocsDefinitionsNotFound handles this case with default header values.

Not found - The requested resource is not found
*/
type GetInstancesDocsDefinitionsNotFound struct {
}

func (o *GetInstancesDocsDefinitionsNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsNotFound ", 404)
}

func (o *GetInstancesDocsDefinitionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsMethodNotAllowed creates a GetInstancesDocsDefinitionsMethodNotAllowed with default headers values
func NewGetInstancesDocsDefinitionsMethodNotAllowed() *GetInstancesDocsDefinitionsMethodNotAllowed {
	return &GetInstancesDocsDefinitionsMethodNotAllowed{}
}

/*GetInstancesDocsDefinitionsMethodNotAllowed handles this case with default header values.

Method not allowed - Incorrect HTTP verb used, e.g., GET used when POST expected
*/
type GetInstancesDocsDefinitionsMethodNotAllowed struct {
}

func (o *GetInstancesDocsDefinitionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsMethodNotAllowed ", 405)
}

func (o *GetInstancesDocsDefinitionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsNotAcceptable creates a GetInstancesDocsDefinitionsNotAcceptable with default headers values
func NewGetInstancesDocsDefinitionsNotAcceptable() *GetInstancesDocsDefinitionsNotAcceptable {
	return &GetInstancesDocsDefinitionsNotAcceptable{}
}

/*GetInstancesDocsDefinitionsNotAcceptable handles this case with default header values.

Not acceptable - The response content type does not match the 'Accept' header value
*/
type GetInstancesDocsDefinitionsNotAcceptable struct {
}

func (o *GetInstancesDocsDefinitionsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsNotAcceptable ", 406)
}

func (o *GetInstancesDocsDefinitionsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsConflict creates a GetInstancesDocsDefinitionsConflict with default headers values
func NewGetInstancesDocsDefinitionsConflict() *GetInstancesDocsDefinitionsConflict {
	return &GetInstancesDocsDefinitionsConflict{}
}

/*GetInstancesDocsDefinitionsConflict handles this case with default header values.

Conflict - If a resource being created already exists
*/
type GetInstancesDocsDefinitionsConflict struct {
}

func (o *GetInstancesDocsDefinitionsConflict) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsConflict ", 409)
}

func (o *GetInstancesDocsDefinitionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsUnsupportedMediaType creates a GetInstancesDocsDefinitionsUnsupportedMediaType with default headers values
func NewGetInstancesDocsDefinitionsUnsupportedMediaType() *GetInstancesDocsDefinitionsUnsupportedMediaType {
	return &GetInstancesDocsDefinitionsUnsupportedMediaType{}
}

/*GetInstancesDocsDefinitionsUnsupportedMediaType handles this case with default header values.

Unsupported media type - The server cannot handle the requested Content-Type
*/
type GetInstancesDocsDefinitionsUnsupportedMediaType struct {
}

func (o *GetInstancesDocsDefinitionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsUnsupportedMediaType ", 415)
}

func (o *GetInstancesDocsDefinitionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstancesDocsDefinitionsInternalServerError creates a GetInstancesDocsDefinitionsInternalServerError with default headers values
func NewGetInstancesDocsDefinitionsInternalServerError() *GetInstancesDocsDefinitionsInternalServerError {
	return &GetInstancesDocsDefinitionsInternalServerError{}
}

/*GetInstancesDocsDefinitionsInternalServerError handles this case with default header values.

Server error - Something went wrong on the Cloud Elements server
*/
type GetInstancesDocsDefinitionsInternalServerError struct {
}

func (o *GetInstancesDocsDefinitionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/docs/{operationId}/definitions][%d] getInstancesDocsDefinitionsInternalServerError ", 500)
}

func (o *GetInstancesDocsDefinitionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
