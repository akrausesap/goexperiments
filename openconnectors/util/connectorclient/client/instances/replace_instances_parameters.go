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

// NewReplaceInstancesParams creates a new ReplaceInstancesParams object
// with the default values initialized.
func NewReplaceInstancesParams() *ReplaceInstancesParams {
	var ()
	return &ReplaceInstancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceInstancesParamsWithTimeout creates a new ReplaceInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceInstancesParamsWithTimeout(timeout time.Duration) *ReplaceInstancesParams {
	var ()
	return &ReplaceInstancesParams{

		timeout: timeout,
	}
}

// NewReplaceInstancesParamsWithContext creates a new ReplaceInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceInstancesParamsWithContext(ctx context.Context) *ReplaceInstancesParams {
	var ()
	return &ReplaceInstancesParams{

		Context: ctx,
	}
}

// NewReplaceInstancesParamsWithHTTPClient creates a new ReplaceInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceInstancesParamsWithHTTPClient(client *http.Client) *ReplaceInstancesParams {
	var ()
	return &ReplaceInstancesParams{
		HTTPClient: client,
	}
}

/*ReplaceInstancesParams contains all the parameters to send to the API endpoint
for the replace instances operation typically these are written to a http.Request
*/
type ReplaceInstancesParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string
	/*ElementInstance
	  The fields of the element instance to update

	*/
	ElementInstance *models.ElementInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace instances params
func (o *ReplaceInstancesParams) WithTimeout(timeout time.Duration) *ReplaceInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace instances params
func (o *ReplaceInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace instances params
func (o *ReplaceInstancesParams) WithContext(ctx context.Context) *ReplaceInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace instances params
func (o *ReplaceInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace instances params
func (o *ReplaceInstancesParams) WithHTTPClient(client *http.Client) *ReplaceInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace instances params
func (o *ReplaceInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the replace instances params
func (o *ReplaceInstancesParams) WithAuthorization(authorization string) *ReplaceInstancesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the replace instances params
func (o *ReplaceInstancesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithElementInstance adds the elementInstance to the replace instances params
func (o *ReplaceInstancesParams) WithElementInstance(elementInstance *models.ElementInstance) *ReplaceInstancesParams {
	o.SetElementInstance(elementInstance)
	return o
}

// SetElementInstance adds the elementInstance to the replace instances params
func (o *ReplaceInstancesParams) SetElementInstance(elementInstance *models.ElementInstance) {
	o.ElementInstance = elementInstance
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ElementInstance != nil {
		if err := r.SetBodyParam(o.ElementInstance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
