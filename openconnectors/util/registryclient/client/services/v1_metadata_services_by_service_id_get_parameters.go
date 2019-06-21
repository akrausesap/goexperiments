// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewV1MetadataServicesByServiceIDGetParams creates a new V1MetadataServicesByServiceIDGetParams object
// with the default values initialized.
func NewV1MetadataServicesByServiceIDGetParams() *V1MetadataServicesByServiceIDGetParams {
	var ()
	return &V1MetadataServicesByServiceIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MetadataServicesByServiceIDGetParamsWithTimeout creates a new V1MetadataServicesByServiceIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MetadataServicesByServiceIDGetParamsWithTimeout(timeout time.Duration) *V1MetadataServicesByServiceIDGetParams {
	var ()
	return &V1MetadataServicesByServiceIDGetParams{

		timeout: timeout,
	}
}

// NewV1MetadataServicesByServiceIDGetParamsWithContext creates a new V1MetadataServicesByServiceIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MetadataServicesByServiceIDGetParamsWithContext(ctx context.Context) *V1MetadataServicesByServiceIDGetParams {
	var ()
	return &V1MetadataServicesByServiceIDGetParams{

		Context: ctx,
	}
}

// NewV1MetadataServicesByServiceIDGetParamsWithHTTPClient creates a new V1MetadataServicesByServiceIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MetadataServicesByServiceIDGetParamsWithHTTPClient(client *http.Client) *V1MetadataServicesByServiceIDGetParams {
	var ()
	return &V1MetadataServicesByServiceIDGetParams{
		HTTPClient: client,
	}
}

/*V1MetadataServicesByServiceIDGetParams contains all the parameters to send to the API endpoint
for the v1 metadata services by service Id get operation typically these are written to a http.Request
*/
type V1MetadataServicesByServiceIDGetParams struct {

	/*ServiceID
	  ID of a service

	*/
	ServiceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) WithTimeout(timeout time.Duration) *V1MetadataServicesByServiceIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) WithContext(ctx context.Context) *V1MetadataServicesByServiceIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) WithHTTPClient(client *http.Client) *V1MetadataServicesByServiceIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) WithServiceID(serviceID strfmt.UUID) *V1MetadataServicesByServiceIDGetParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the v1 metadata services by service Id get params
func (o *V1MetadataServicesByServiceIDGetParams) SetServiceID(serviceID strfmt.UUID) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *V1MetadataServicesByServiceIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}