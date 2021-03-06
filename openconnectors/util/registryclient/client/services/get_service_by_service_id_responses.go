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

// GetServiceByServiceIDReader is a Reader for the GetServiceByServiceID structure.
type GetServiceByServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceByServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceByServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetServiceByServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetServiceByServiceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceByServiceIDOK creates a GetServiceByServiceIDOK with default headers values
func NewGetServiceByServiceIDOK() *GetServiceByServiceIDOK {
	return &GetServiceByServiceIDOK{}
}

/*GetServiceByServiceIDOK handles this case with default header values.

Successful operation
*/
type GetServiceByServiceIDOK struct {
	Payload *models.ServiceDetails
}

func (o *GetServiceByServiceIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/metadata/services/{serviceId}][%d] getServiceByServiceIdOK  %+v", 200, o.Payload)
}

func (o *GetServiceByServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceByServiceIDNotFound creates a GetServiceByServiceIDNotFound with default headers values
func NewGetServiceByServiceIDNotFound() *GetServiceByServiceIDNotFound {
	return &GetServiceByServiceIDNotFound{}
}

/*GetServiceByServiceIDNotFound handles this case with default header values.

Service not found
*/
type GetServiceByServiceIDNotFound struct {
	Payload *models.MetadataErrorResponse
}

func (o *GetServiceByServiceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/metadata/services/{serviceId}][%d] getServiceByServiceIdNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceByServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceByServiceIDInternalServerError creates a GetServiceByServiceIDInternalServerError with default headers values
func NewGetServiceByServiceIDInternalServerError() *GetServiceByServiceIDInternalServerError {
	return &GetServiceByServiceIDInternalServerError{}
}

/*GetServiceByServiceIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetServiceByServiceIDInternalServerError struct {
	Payload *models.MetadataErrorResponse
}

func (o *GetServiceByServiceIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/metadata/services/{serviceId}][%d] getServiceByServiceIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServiceByServiceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
