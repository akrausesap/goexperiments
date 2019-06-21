// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/akrausesap/goexperiments/openconnectors/util/registryclient/models"
)

// V1MetadataServicesPostReader is a Reader for the V1MetadataServicesPost structure.
type V1MetadataServicesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MetadataServicesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV1MetadataServicesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewV1MetadataServicesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV1MetadataServicesPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV1MetadataServicesPostOK creates a V1MetadataServicesPostOK with default headers values
func NewV1MetadataServicesPostOK() *V1MetadataServicesPostOK {
	return &V1MetadataServicesPostOK{}
}

/*V1MetadataServicesPostOK handles this case with default header values.

Successful operation
*/
type V1MetadataServicesPostOK struct {
	Payload *models.ServiceID
}

func (o *V1MetadataServicesPostOK) Error() string {
	return fmt.Sprintf("[POST /v1/metadata/services][%d] v1MetadataServicesPostOK  %+v", 200, o.Payload)
}

func (o *V1MetadataServicesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1MetadataServicesPostBadRequest creates a V1MetadataServicesPostBadRequest with default headers values
func NewV1MetadataServicesPostBadRequest() *V1MetadataServicesPostBadRequest {
	return &V1MetadataServicesPostBadRequest{}
}

/*V1MetadataServicesPostBadRequest handles this case with default header values.

Bad request
*/
type V1MetadataServicesPostBadRequest struct {
	Payload *models.MetadataErrorResponse
}

func (o *V1MetadataServicesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/metadata/services][%d] v1MetadataServicesPostBadRequest  %+v", 400, o.Payload)
}

func (o *V1MetadataServicesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1MetadataServicesPostDefault creates a V1MetadataServicesPostDefault with default headers values
func NewV1MetadataServicesPostDefault(code int) *V1MetadataServicesPostDefault {
	return &V1MetadataServicesPostDefault{
		_statusCode: code,
	}
}

/*V1MetadataServicesPostDefault handles this case with default header values.

Internal server error
*/
type V1MetadataServicesPostDefault struct {
	_statusCode int

	Payload *models.MetadataErrorResponse
}

// Code gets the status code for the v1 metadata services post default response
func (o *V1MetadataServicesPostDefault) Code() int {
	return o._statusCode
}

func (o *V1MetadataServicesPostDefault) Error() string {
	return fmt.Sprintf("[POST /v1/metadata/services][%d] V1MetadataServicesPost default  %+v", o._statusCode, o.Payload)
}

func (o *V1MetadataServicesPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}