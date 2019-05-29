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

// NewGetInstancesObjectsObjectNameDefinitionsParams creates a new GetInstancesObjectsObjectNameDefinitionsParams object
// with the default values initialized.
func NewGetInstancesObjectsObjectNameDefinitionsParams() *GetInstancesObjectsObjectNameDefinitionsParams {
	var ()
	return &GetInstancesObjectsObjectNameDefinitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesObjectsObjectNameDefinitionsParamsWithTimeout creates a new GetInstancesObjectsObjectNameDefinitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesObjectsObjectNameDefinitionsParamsWithTimeout(timeout time.Duration) *GetInstancesObjectsObjectNameDefinitionsParams {
	var ()
	return &GetInstancesObjectsObjectNameDefinitionsParams{

		timeout: timeout,
	}
}

// NewGetInstancesObjectsObjectNameDefinitionsParamsWithContext creates a new GetInstancesObjectsObjectNameDefinitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesObjectsObjectNameDefinitionsParamsWithContext(ctx context.Context) *GetInstancesObjectsObjectNameDefinitionsParams {
	var ()
	return &GetInstancesObjectsObjectNameDefinitionsParams{

		Context: ctx,
	}
}

// NewGetInstancesObjectsObjectNameDefinitionsParamsWithHTTPClient creates a new GetInstancesObjectsObjectNameDefinitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesObjectsObjectNameDefinitionsParamsWithHTTPClient(client *http.Client) *GetInstancesObjectsObjectNameDefinitionsParams {
	var ()
	return &GetInstancesObjectsObjectNameDefinitionsParams{
		HTTPClient: client,
	}
}

/*GetInstancesObjectsObjectNameDefinitionsParams contains all the parameters to send to the API endpoint
for the get instances objects object name definitions operation typically these are written to a http.Request
*/
type GetInstancesObjectsObjectNameDefinitionsParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string
	/*ObjectName
	  The name of the object

	*/
	ObjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) WithTimeout(timeout time.Duration) *GetInstancesObjectsObjectNameDefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) WithContext(ctx context.Context) *GetInstancesObjectsObjectNameDefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) WithHTTPClient(client *http.Client) *GetInstancesObjectsObjectNameDefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) WithAuthorization(authorization string) *GetInstancesObjectsObjectNameDefinitionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithObjectName adds the objectName to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) WithObjectName(objectName string) *GetInstancesObjectsObjectNameDefinitionsParams {
	o.SetObjectName(objectName)
	return o
}

// SetObjectName adds the objectName to the get instances objects object name definitions params
func (o *GetInstancesObjectsObjectNameDefinitionsParams) SetObjectName(objectName string) {
	o.ObjectName = objectName
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesObjectsObjectNameDefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param objectName
	if err := r.SetPathParam("objectName", o.ObjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}