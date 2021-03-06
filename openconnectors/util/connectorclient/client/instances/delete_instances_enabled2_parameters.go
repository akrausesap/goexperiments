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

// NewDeleteInstancesEnabled2Params creates a new DeleteInstancesEnabled2Params object
// with the default values initialized.
func NewDeleteInstancesEnabled2Params() *DeleteInstancesEnabled2Params {
	var ()
	return &DeleteInstancesEnabled2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesEnabled2ParamsWithTimeout creates a new DeleteInstancesEnabled2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesEnabled2ParamsWithTimeout(timeout time.Duration) *DeleteInstancesEnabled2Params {
	var ()
	return &DeleteInstancesEnabled2Params{

		timeout: timeout,
	}
}

// NewDeleteInstancesEnabled2ParamsWithContext creates a new DeleteInstancesEnabled2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesEnabled2ParamsWithContext(ctx context.Context) *DeleteInstancesEnabled2Params {
	var ()
	return &DeleteInstancesEnabled2Params{

		Context: ctx,
	}
}

// NewDeleteInstancesEnabled2ParamsWithHTTPClient creates a new DeleteInstancesEnabled2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesEnabled2ParamsWithHTTPClient(client *http.Client) *DeleteInstancesEnabled2Params {
	var ()
	return &DeleteInstancesEnabled2Params{
		HTTPClient: client,
	}
}

/*DeleteInstancesEnabled2Params contains all the parameters to send to the API endpoint
for the delete instances enabled2 operation typically these are written to a http.Request
*/
type DeleteInstancesEnabled2Params struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;'

	*/
	Authorization string
	/*ID
	  The ID of the element instance

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) WithTimeout(timeout time.Duration) *DeleteInstancesEnabled2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) WithContext(ctx context.Context) *DeleteInstancesEnabled2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) WithHTTPClient(client *http.Client) *DeleteInstancesEnabled2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) WithAuthorization(authorization string) *DeleteInstancesEnabled2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) WithID(id int64) *DeleteInstancesEnabled2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instances enabled2 params
func (o *DeleteInstancesEnabled2Params) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesEnabled2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
