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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstancesConfigurationByConfigIDParams creates a new GetInstancesConfigurationByConfigIDParams object
// with the default values initialized.
func NewGetInstancesConfigurationByConfigIDParams() *GetInstancesConfigurationByConfigIDParams {
	var ()
	return &GetInstancesConfigurationByConfigIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesConfigurationByConfigIDParamsWithTimeout creates a new GetInstancesConfigurationByConfigIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesConfigurationByConfigIDParamsWithTimeout(timeout time.Duration) *GetInstancesConfigurationByConfigIDParams {
	var ()
	return &GetInstancesConfigurationByConfigIDParams{

		timeout: timeout,
	}
}

// NewGetInstancesConfigurationByConfigIDParamsWithContext creates a new GetInstancesConfigurationByConfigIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesConfigurationByConfigIDParamsWithContext(ctx context.Context) *GetInstancesConfigurationByConfigIDParams {
	var ()
	return &GetInstancesConfigurationByConfigIDParams{

		Context: ctx,
	}
}

// NewGetInstancesConfigurationByConfigIDParamsWithHTTPClient creates a new GetInstancesConfigurationByConfigIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesConfigurationByConfigIDParamsWithHTTPClient(client *http.Client) *GetInstancesConfigurationByConfigIDParams {
	var ()
	return &GetInstancesConfigurationByConfigIDParams{
		HTTPClient: client,
	}
}

/*GetInstancesConfigurationByConfigIDParams contains all the parameters to send to the API endpoint
for the get instances configuration by config Id operation typically these are written to a http.Request
*/
type GetInstancesConfigurationByConfigIDParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string
	/*ConfigID
	  The ID of the element instance config

	*/
	ConfigID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) WithTimeout(timeout time.Duration) *GetInstancesConfigurationByConfigIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) WithContext(ctx context.Context) *GetInstancesConfigurationByConfigIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) WithHTTPClient(client *http.Client) *GetInstancesConfigurationByConfigIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) WithAuthorization(authorization string) *GetInstancesConfigurationByConfigIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithConfigID adds the configID to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) WithConfigID(configID int64) *GetInstancesConfigurationByConfigIDParams {
	o.SetConfigID(configID)
	return o
}

// SetConfigID adds the configId to the get instances configuration by config Id params
func (o *GetInstancesConfigurationByConfigIDParams) SetConfigID(configID int64) {
	o.ConfigID = configID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesConfigurationByConfigIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param configId
	if err := r.SetPathParam("configId", swag.FormatInt64(o.ConfigID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
