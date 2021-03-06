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

// NewDeleteInstancesTransformationByObjectNameParams creates a new DeleteInstancesTransformationByObjectNameParams object
// with the default values initialized.
func NewDeleteInstancesTransformationByObjectNameParams() *DeleteInstancesTransformationByObjectNameParams {
	var ()
	return &DeleteInstancesTransformationByObjectNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesTransformationByObjectNameParamsWithTimeout creates a new DeleteInstancesTransformationByObjectNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesTransformationByObjectNameParamsWithTimeout(timeout time.Duration) *DeleteInstancesTransformationByObjectNameParams {
	var ()
	return &DeleteInstancesTransformationByObjectNameParams{

		timeout: timeout,
	}
}

// NewDeleteInstancesTransformationByObjectNameParamsWithContext creates a new DeleteInstancesTransformationByObjectNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesTransformationByObjectNameParamsWithContext(ctx context.Context) *DeleteInstancesTransformationByObjectNameParams {
	var ()
	return &DeleteInstancesTransformationByObjectNameParams{

		Context: ctx,
	}
}

// NewDeleteInstancesTransformationByObjectNameParamsWithHTTPClient creates a new DeleteInstancesTransformationByObjectNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesTransformationByObjectNameParamsWithHTTPClient(client *http.Client) *DeleteInstancesTransformationByObjectNameParams {
	var ()
	return &DeleteInstancesTransformationByObjectNameParams{
		HTTPClient: client,
	}
}

/*DeleteInstancesTransformationByObjectNameParams contains all the parameters to send to the API endpoint
for the delete instances transformation by object name operation typically these are written to a http.Request
*/
type DeleteInstancesTransformationByObjectNameParams struct {

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

// WithTimeout adds the timeout to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) WithTimeout(timeout time.Duration) *DeleteInstancesTransformationByObjectNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) WithContext(ctx context.Context) *DeleteInstancesTransformationByObjectNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) WithHTTPClient(client *http.Client) *DeleteInstancesTransformationByObjectNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) WithAuthorization(authorization string) *DeleteInstancesTransformationByObjectNameParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithObjectName adds the objectName to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) WithObjectName(objectName string) *DeleteInstancesTransformationByObjectNameParams {
	o.SetObjectName(objectName)
	return o
}

// SetObjectName adds the objectName to the delete instances transformation by object name params
func (o *DeleteInstancesTransformationByObjectNameParams) SetObjectName(objectName string) {
	o.ObjectName = objectName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesTransformationByObjectNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
