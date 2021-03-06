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

// NewGetInstancesConfigurationByConfigId2Params creates a new GetInstancesConfigurationByConfigId2Params object
// with the default values initialized.
func NewGetInstancesConfigurationByConfigId2Params() *GetInstancesConfigurationByConfigId2Params {
	var ()
	return &GetInstancesConfigurationByConfigId2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesConfigurationByConfigId2ParamsWithTimeout creates a new GetInstancesConfigurationByConfigId2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesConfigurationByConfigId2ParamsWithTimeout(timeout time.Duration) *GetInstancesConfigurationByConfigId2Params {
	var ()
	return &GetInstancesConfigurationByConfigId2Params{

		timeout: timeout,
	}
}

// NewGetInstancesConfigurationByConfigId2ParamsWithContext creates a new GetInstancesConfigurationByConfigId2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesConfigurationByConfigId2ParamsWithContext(ctx context.Context) *GetInstancesConfigurationByConfigId2Params {
	var ()
	return &GetInstancesConfigurationByConfigId2Params{

		Context: ctx,
	}
}

// NewGetInstancesConfigurationByConfigId2ParamsWithHTTPClient creates a new GetInstancesConfigurationByConfigId2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesConfigurationByConfigId2ParamsWithHTTPClient(client *http.Client) *GetInstancesConfigurationByConfigId2Params {
	var ()
	return &GetInstancesConfigurationByConfigId2Params{
		HTTPClient: client,
	}
}

/*GetInstancesConfigurationByConfigId2Params contains all the parameters to send to the API endpoint
for the get instances configuration by config id2 operation typically these are written to a http.Request
*/
type GetInstancesConfigurationByConfigId2Params struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;'

	*/
	Authorization string
	/*ConfigID
	  The ID of the element instance config

	*/
	ConfigID int64
	/*ID
	  The ID of the element instance

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) WithTimeout(timeout time.Duration) *GetInstancesConfigurationByConfigId2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) WithContext(ctx context.Context) *GetInstancesConfigurationByConfigId2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) WithHTTPClient(client *http.Client) *GetInstancesConfigurationByConfigId2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) WithAuthorization(authorization string) *GetInstancesConfigurationByConfigId2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithConfigID adds the configID to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) WithConfigID(configID int64) *GetInstancesConfigurationByConfigId2Params {
	o.SetConfigID(configID)
	return o
}

// SetConfigID adds the configId to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) SetConfigID(configID int64) {
	o.ConfigID = configID
}

// WithID adds the id to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) WithID(id int64) *GetInstancesConfigurationByConfigId2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instances configuration by config id2 params
func (o *GetInstancesConfigurationByConfigId2Params) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesConfigurationByConfigId2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
