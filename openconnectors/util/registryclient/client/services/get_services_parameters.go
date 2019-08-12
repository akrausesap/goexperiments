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

// NewGetServicesParams creates a new GetServicesParams object
// with the default values initialized.
func NewGetServicesParams() *GetServicesParams {

	return &GetServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicesParamsWithTimeout creates a new GetServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicesParamsWithTimeout(timeout time.Duration) *GetServicesParams {

	return &GetServicesParams{

		timeout: timeout,
	}
}

// NewGetServicesParamsWithContext creates a new GetServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServicesParamsWithContext(ctx context.Context) *GetServicesParams {

	return &GetServicesParams{

		Context: ctx,
	}
}

// NewGetServicesParamsWithHTTPClient creates a new GetServicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServicesParamsWithHTTPClient(client *http.Client) *GetServicesParams {

	return &GetServicesParams{
		HTTPClient: client,
	}
}

/*GetServicesParams contains all the parameters to send to the API endpoint
for the get services operation typically these are written to a http.Request
*/
type GetServicesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get services params
func (o *GetServicesParams) WithTimeout(timeout time.Duration) *GetServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get services params
func (o *GetServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get services params
func (o *GetServicesParams) WithContext(ctx context.Context) *GetServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get services params
func (o *GetServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get services params
func (o *GetServicesParams) WithHTTPClient(client *http.Client) *GetServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get services params
func (o *GetServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
