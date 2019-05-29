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

// NewDeleteInstanceByIDParams creates a new DeleteInstanceByIDParams object
// with the default values initialized.
func NewDeleteInstanceByIDParams() *DeleteInstanceByIDParams {
	var ()
	return &DeleteInstanceByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceByIDParamsWithTimeout creates a new DeleteInstanceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstanceByIDParamsWithTimeout(timeout time.Duration) *DeleteInstanceByIDParams {
	var ()
	return &DeleteInstanceByIDParams{

		timeout: timeout,
	}
}

// NewDeleteInstanceByIDParamsWithContext creates a new DeleteInstanceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstanceByIDParamsWithContext(ctx context.Context) *DeleteInstanceByIDParams {
	var ()
	return &DeleteInstanceByIDParams{

		Context: ctx,
	}
}

// NewDeleteInstanceByIDParamsWithHTTPClient creates a new DeleteInstanceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstanceByIDParamsWithHTTPClient(client *http.Client) *DeleteInstanceByIDParams {
	var ()
	return &DeleteInstanceByIDParams{
		HTTPClient: client,
	}
}

/*DeleteInstanceByIDParams contains all the parameters to send to the API endpoint
for the delete instance by Id operation typically these are written to a http.Request
*/
type DeleteInstanceByIDParams struct {

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

// WithTimeout adds the timeout to the delete instance by Id params
func (o *DeleteInstanceByIDParams) WithTimeout(timeout time.Duration) *DeleteInstanceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance by Id params
func (o *DeleteInstanceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance by Id params
func (o *DeleteInstanceByIDParams) WithContext(ctx context.Context) *DeleteInstanceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance by Id params
func (o *DeleteInstanceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance by Id params
func (o *DeleteInstanceByIDParams) WithHTTPClient(client *http.Client) *DeleteInstanceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance by Id params
func (o *DeleteInstanceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete instance by Id params
func (o *DeleteInstanceByIDParams) WithAuthorization(authorization string) *DeleteInstanceByIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete instance by Id params
func (o *DeleteInstanceByIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the delete instance by Id params
func (o *DeleteInstanceByIDParams) WithID(id int64) *DeleteInstanceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instance by Id params
func (o *DeleteInstanceByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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