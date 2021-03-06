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

	models "github.com/akrausesap/goexperiments/openconnectors/util/registryclient/models"
)

// NewV1MetadataServicesPostParams creates a new V1MetadataServicesPostParams object
// with the default values initialized.
func NewV1MetadataServicesPostParams() *V1MetadataServicesPostParams {
	var ()
	return &V1MetadataServicesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MetadataServicesPostParamsWithTimeout creates a new V1MetadataServicesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MetadataServicesPostParamsWithTimeout(timeout time.Duration) *V1MetadataServicesPostParams {
	var ()
	return &V1MetadataServicesPostParams{

		timeout: timeout,
	}
}

// NewV1MetadataServicesPostParamsWithContext creates a new V1MetadataServicesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MetadataServicesPostParamsWithContext(ctx context.Context) *V1MetadataServicesPostParams {
	var ()
	return &V1MetadataServicesPostParams{

		Context: ctx,
	}
}

// NewV1MetadataServicesPostParamsWithHTTPClient creates a new V1MetadataServicesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MetadataServicesPostParamsWithHTTPClient(client *http.Client) *V1MetadataServicesPostParams {
	var ()
	return &V1MetadataServicesPostParams{
		HTTPClient: client,
	}
}

/*V1MetadataServicesPostParams contains all the parameters to send to the API endpoint
for the v1 metadata services post operation typically these are written to a http.Request
*/
type V1MetadataServicesPostParams struct {

	/*Body
	  Service object to be registered

	*/
	Body *models.ServiceDetails

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) WithTimeout(timeout time.Duration) *V1MetadataServicesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) WithContext(ctx context.Context) *V1MetadataServicesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) WithHTTPClient(client *http.Client) *V1MetadataServicesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) WithBody(body *models.ServiceDetails) *V1MetadataServicesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 metadata services post params
func (o *V1MetadataServicesPostParams) SetBody(body *models.ServiceDetails) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1MetadataServicesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
