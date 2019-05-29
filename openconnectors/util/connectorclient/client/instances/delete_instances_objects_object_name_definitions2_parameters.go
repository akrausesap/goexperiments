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

// NewDeleteInstancesObjectsObjectNameDefinitions2Params creates a new DeleteInstancesObjectsObjectNameDefinitions2Params object
// with the default values initialized.
func NewDeleteInstancesObjectsObjectNameDefinitions2Params() *DeleteInstancesObjectsObjectNameDefinitions2Params {
	var ()
	return &DeleteInstancesObjectsObjectNameDefinitions2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesObjectsObjectNameDefinitions2ParamsWithTimeout creates a new DeleteInstancesObjectsObjectNameDefinitions2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesObjectsObjectNameDefinitions2ParamsWithTimeout(timeout time.Duration) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	var ()
	return &DeleteInstancesObjectsObjectNameDefinitions2Params{

		timeout: timeout,
	}
}

// NewDeleteInstancesObjectsObjectNameDefinitions2ParamsWithContext creates a new DeleteInstancesObjectsObjectNameDefinitions2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesObjectsObjectNameDefinitions2ParamsWithContext(ctx context.Context) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	var ()
	return &DeleteInstancesObjectsObjectNameDefinitions2Params{

		Context: ctx,
	}
}

// NewDeleteInstancesObjectsObjectNameDefinitions2ParamsWithHTTPClient creates a new DeleteInstancesObjectsObjectNameDefinitions2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesObjectsObjectNameDefinitions2ParamsWithHTTPClient(client *http.Client) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	var ()
	return &DeleteInstancesObjectsObjectNameDefinitions2Params{
		HTTPClient: client,
	}
}

/*DeleteInstancesObjectsObjectNameDefinitions2Params contains all the parameters to send to the API endpoint
for the delete instances objects object name definitions2 operation typically these are written to a http.Request
*/
type DeleteInstancesObjectsObjectNameDefinitions2Params struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;'

	*/
	Authorization string
	/*ID
	  The ID of the instance

	*/
	ID int64
	/*ObjectName
	  The name of the object

	*/
	ObjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) WithTimeout(timeout time.Duration) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) WithContext(ctx context.Context) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) WithHTTPClient(client *http.Client) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) WithAuthorization(authorization string) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) WithID(id int64) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) SetID(id int64) {
	o.ID = id
}

// WithObjectName adds the objectName to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) WithObjectName(objectName string) *DeleteInstancesObjectsObjectNameDefinitions2Params {
	o.SetObjectName(objectName)
	return o
}

// SetObjectName adds the objectName to the delete instances objects object name definitions2 params
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) SetObjectName(objectName string) {
	o.ObjectName = objectName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesObjectsObjectNameDefinitions2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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
