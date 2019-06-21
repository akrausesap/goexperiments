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

// NewV1MetadataServicesGetParams creates a new V1MetadataServicesGetParams object
// with the default values initialized.
func NewV1MetadataServicesGetParams() *V1MetadataServicesGetParams {

	return &V1MetadataServicesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MetadataServicesGetParamsWithTimeout creates a new V1MetadataServicesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MetadataServicesGetParamsWithTimeout(timeout time.Duration) *V1MetadataServicesGetParams {

	return &V1MetadataServicesGetParams{

		timeout: timeout,
	}
}

// NewV1MetadataServicesGetParamsWithContext creates a new V1MetadataServicesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MetadataServicesGetParamsWithContext(ctx context.Context) *V1MetadataServicesGetParams {

	return &V1MetadataServicesGetParams{

		Context: ctx,
	}
}

// NewV1MetadataServicesGetParamsWithHTTPClient creates a new V1MetadataServicesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MetadataServicesGetParamsWithHTTPClient(client *http.Client) *V1MetadataServicesGetParams {

	return &V1MetadataServicesGetParams{
		HTTPClient: client,
	}
}

/*V1MetadataServicesGetParams contains all the parameters to send to the API endpoint
for the v1 metadata services get operation typically these are written to a http.Request
*/
type V1MetadataServicesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 metadata services get params
func (o *V1MetadataServicesGetParams) WithTimeout(timeout time.Duration) *V1MetadataServicesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 metadata services get params
func (o *V1MetadataServicesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 metadata services get params
func (o *V1MetadataServicesGetParams) WithContext(ctx context.Context) *V1MetadataServicesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 metadata services get params
func (o *V1MetadataServicesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 metadata services get params
func (o *V1MetadataServicesGetParams) WithHTTPClient(client *http.Client) *V1MetadataServicesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 metadata services get params
func (o *V1MetadataServicesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1MetadataServicesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
