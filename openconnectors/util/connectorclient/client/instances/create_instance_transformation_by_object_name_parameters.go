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

	models "github.com/akrausesap/goexperiments/openconnectors/util/connectorclient/models"
)

// NewCreateInstanceTransformationByObjectNameParams creates a new CreateInstanceTransformationByObjectNameParams object
// with the default values initialized.
func NewCreateInstanceTransformationByObjectNameParams() *CreateInstanceTransformationByObjectNameParams {
	var ()
	return &CreateInstanceTransformationByObjectNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstanceTransformationByObjectNameParamsWithTimeout creates a new CreateInstanceTransformationByObjectNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInstanceTransformationByObjectNameParamsWithTimeout(timeout time.Duration) *CreateInstanceTransformationByObjectNameParams {
	var ()
	return &CreateInstanceTransformationByObjectNameParams{

		timeout: timeout,
	}
}

// NewCreateInstanceTransformationByObjectNameParamsWithContext creates a new CreateInstanceTransformationByObjectNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInstanceTransformationByObjectNameParamsWithContext(ctx context.Context) *CreateInstanceTransformationByObjectNameParams {
	var ()
	return &CreateInstanceTransformationByObjectNameParams{

		Context: ctx,
	}
}

// NewCreateInstanceTransformationByObjectNameParamsWithHTTPClient creates a new CreateInstanceTransformationByObjectNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInstanceTransformationByObjectNameParamsWithHTTPClient(client *http.Client) *CreateInstanceTransformationByObjectNameParams {
	var ()
	return &CreateInstanceTransformationByObjectNameParams{
		HTTPClient: client,
	}
}

/*CreateInstanceTransformationByObjectNameParams contains all the parameters to send to the API endpoint
for the create instance transformation by object name operation typically these are written to a http.Request
*/
type CreateInstanceTransformationByObjectNameParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string
	/*ObjectName
	  The name of the object

	*/
	ObjectName string
	/*Transformation
	  The transformation to create

	*/
	Transformation *models.Transformation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) WithTimeout(timeout time.Duration) *CreateInstanceTransformationByObjectNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) WithContext(ctx context.Context) *CreateInstanceTransformationByObjectNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) WithHTTPClient(client *http.Client) *CreateInstanceTransformationByObjectNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) WithAuthorization(authorization string) *CreateInstanceTransformationByObjectNameParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithObjectName adds the objectName to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) WithObjectName(objectName string) *CreateInstanceTransformationByObjectNameParams {
	o.SetObjectName(objectName)
	return o
}

// SetObjectName adds the objectName to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) SetObjectName(objectName string) {
	o.ObjectName = objectName
}

// WithTransformation adds the transformation to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) WithTransformation(transformation *models.Transformation) *CreateInstanceTransformationByObjectNameParams {
	o.SetTransformation(transformation)
	return o
}

// SetTransformation adds the transformation to the create instance transformation by object name params
func (o *CreateInstanceTransformationByObjectNameParams) SetTransformation(transformation *models.Transformation) {
	o.Transformation = transformation
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstanceTransformationByObjectNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Transformation != nil {
		if err := r.SetBodyParam(o.Transformation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
