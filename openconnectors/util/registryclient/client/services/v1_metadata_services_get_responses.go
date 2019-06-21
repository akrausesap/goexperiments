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

// V1MetadataServicesGetReader is a Reader for the V1MetadataServicesGet structure.
type V1MetadataServicesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MetadataServicesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV1MetadataServicesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewV1MetadataServicesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV1MetadataServicesGetOK creates a V1MetadataServicesGetOK with default headers values
func NewV1MetadataServicesGetOK() *V1MetadataServicesGetOK {
	return &V1MetadataServicesGetOK{}
}

/*V1MetadataServicesGetOK handles this case with default header values.

Successful operation
*/
type V1MetadataServicesGetOK struct {
	Payload []*models.Service
}

func (o *V1MetadataServicesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/metadata/services][%d] v1MetadataServicesGetOK  %+v", 200, o.Payload)
}

func (o *V1MetadataServicesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1MetadataServicesGetDefault creates a V1MetadataServicesGetDefault with default headers values
func NewV1MetadataServicesGetDefault(code int) *V1MetadataServicesGetDefault {
	return &V1MetadataServicesGetDefault{
		_statusCode: code,
	}
}

/*V1MetadataServicesGetDefault handles this case with default header values.

Internal server error
*/
type V1MetadataServicesGetDefault struct {
	_statusCode int

	Payload *models.MetadataErrorResponse
}

// Code gets the status code for the v1 metadata services get default response
func (o *V1MetadataServicesGetDefault) Code() int {
	return o._statusCode
}

func (o *V1MetadataServicesGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1/metadata/services][%d] V1MetadataServicesGet default  %+v", o._statusCode, o.Payload)
}

func (o *V1MetadataServicesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}