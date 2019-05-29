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

	models "github.com/akrausesap/goexperiments/openconnectors/util/connectorclient/models"
)

// NewCreateInstanceObjectObjectNameDefinition2Params creates a new CreateInstanceObjectObjectNameDefinition2Params object
// with the default values initialized.
func NewCreateInstanceObjectObjectNameDefinition2Params() *CreateInstanceObjectObjectNameDefinition2Params {
	var ()
	return &CreateInstanceObjectObjectNameDefinition2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstanceObjectObjectNameDefinition2ParamsWithTimeout creates a new CreateInstanceObjectObjectNameDefinition2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInstanceObjectObjectNameDefinition2ParamsWithTimeout(timeout time.Duration) *CreateInstanceObjectObjectNameDefinition2Params {
	var ()
	return &CreateInstanceObjectObjectNameDefinition2Params{

		timeout: timeout,
	}
}

// NewCreateInstanceObjectObjectNameDefinition2ParamsWithContext creates a new CreateInstanceObjectObjectNameDefinition2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInstanceObjectObjectNameDefinition2ParamsWithContext(ctx context.Context) *CreateInstanceObjectObjectNameDefinition2Params {
	var ()
	return &CreateInstanceObjectObjectNameDefinition2Params{

		Context: ctx,
	}
}

// NewCreateInstanceObjectObjectNameDefinition2ParamsWithHTTPClient creates a new CreateInstanceObjectObjectNameDefinition2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInstanceObjectObjectNameDefinition2ParamsWithHTTPClient(client *http.Client) *CreateInstanceObjectObjectNameDefinition2Params {
	var ()
	return &CreateInstanceObjectObjectNameDefinition2Params{
		HTTPClient: client,
	}
}

/*CreateInstanceObjectObjectNameDefinition2Params contains all the parameters to send to the API endpoint
for the create instance object object name definition2 operation typically these are written to a http.Request
*/
type CreateInstanceObjectObjectNameDefinition2Params struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;'

	*/
	Authorization string
	/*Body
	  The object definition

	*/
	Body *models.Definition
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

// WithTimeout adds the timeout to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) WithTimeout(timeout time.Duration) *CreateInstanceObjectObjectNameDefinition2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) WithContext(ctx context.Context) *CreateInstanceObjectObjectNameDefinition2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) WithHTTPClient(client *http.Client) *CreateInstanceObjectObjectNameDefinition2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) WithAuthorization(authorization string) *CreateInstanceObjectObjectNameDefinition2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) WithBody(body *models.Definition) *CreateInstanceObjectObjectNameDefinition2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) SetBody(body *models.Definition) {
	o.Body = body
}

// WithID adds the id to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) WithID(id int64) *CreateInstanceObjectObjectNameDefinition2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) SetID(id int64) {
	o.ID = id
}

// WithObjectName adds the objectName to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) WithObjectName(objectName string) *CreateInstanceObjectObjectNameDefinition2Params {
	o.SetObjectName(objectName)
	return o
}

// SetObjectName adds the objectName to the create instance object object name definition2 params
func (o *CreateInstanceObjectObjectNameDefinition2Params) SetObjectName(objectName string) {
	o.ObjectName = objectName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstanceObjectObjectNameDefinition2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
