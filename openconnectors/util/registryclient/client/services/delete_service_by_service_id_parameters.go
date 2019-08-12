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

// NewDeleteServiceByServiceIDParams creates a new DeleteServiceByServiceIDParams object
// with the default values initialized.
func NewDeleteServiceByServiceIDParams() *DeleteServiceByServiceIDParams {
	var ()
	return &DeleteServiceByServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceByServiceIDParamsWithTimeout creates a new DeleteServiceByServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServiceByServiceIDParamsWithTimeout(timeout time.Duration) *DeleteServiceByServiceIDParams {
	var ()
	return &DeleteServiceByServiceIDParams{

		timeout: timeout,
	}
}

// NewDeleteServiceByServiceIDParamsWithContext creates a new DeleteServiceByServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServiceByServiceIDParamsWithContext(ctx context.Context) *DeleteServiceByServiceIDParams {
	var ()
	return &DeleteServiceByServiceIDParams{

		Context: ctx,
	}
}

// NewDeleteServiceByServiceIDParamsWithHTTPClient creates a new DeleteServiceByServiceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServiceByServiceIDParamsWithHTTPClient(client *http.Client) *DeleteServiceByServiceIDParams {
	var ()
	return &DeleteServiceByServiceIDParams{
		HTTPClient: client,
	}
}

/*DeleteServiceByServiceIDParams contains all the parameters to send to the API endpoint
for the delete service by service Id operation typically these are written to a http.Request
*/
type DeleteServiceByServiceIDParams struct {

	/*ServiceID
	  ID of a service

	*/
	ServiceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) WithTimeout(timeout time.Duration) *DeleteServiceByServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) WithContext(ctx context.Context) *DeleteServiceByServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) WithHTTPClient(client *http.Client) *DeleteServiceByServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) WithServiceID(serviceID strfmt.UUID) *DeleteServiceByServiceIDParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the delete service by service Id params
func (o *DeleteServiceByServiceIDParams) SetServiceID(serviceID strfmt.UUID) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceByServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
