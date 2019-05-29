// Code generated by go-swagger; DO NOT EDIT.

package instances

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

// NewGetInstancesConfigurationParams creates a new GetInstancesConfigurationParams object
// with the default values initialized.
func NewGetInstancesConfigurationParams() *GetInstancesConfigurationParams {
	var ()
	return &GetInstancesConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesConfigurationParamsWithTimeout creates a new GetInstancesConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesConfigurationParamsWithTimeout(timeout time.Duration) *GetInstancesConfigurationParams {
	var ()
	return &GetInstancesConfigurationParams{

		timeout: timeout,
	}
}

// NewGetInstancesConfigurationParamsWithContext creates a new GetInstancesConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesConfigurationParamsWithContext(ctx context.Context) *GetInstancesConfigurationParams {
	var ()
	return &GetInstancesConfigurationParams{

		Context: ctx,
	}
}

// NewGetInstancesConfigurationParamsWithHTTPClient creates a new GetInstancesConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesConfigurationParamsWithHTTPClient(client *http.Client) *GetInstancesConfigurationParams {
	var ()
	return &GetInstancesConfigurationParams{
		HTTPClient: client,
	}
}

/*GetInstancesConfigurationParams contains all the parameters to send to the API endpoint
for the get instances configuration operation typically these are written to a http.Request
*/
type GetInstancesConfigurationParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances configuration params
func (o *GetInstancesConfigurationParams) WithTimeout(timeout time.Duration) *GetInstancesConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances configuration params
func (o *GetInstancesConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances configuration params
func (o *GetInstancesConfigurationParams) WithContext(ctx context.Context) *GetInstancesConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances configuration params
func (o *GetInstancesConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances configuration params
func (o *GetInstancesConfigurationParams) WithHTTPClient(client *http.Client) *GetInstancesConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances configuration params
func (o *GetInstancesConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get instances configuration params
func (o *GetInstancesConfigurationParams) WithAuthorization(authorization string) *GetInstancesConfigurationParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get instances configuration params
func (o *GetInstancesConfigurationParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
