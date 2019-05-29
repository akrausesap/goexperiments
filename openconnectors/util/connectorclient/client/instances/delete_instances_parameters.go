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

// NewDeleteInstancesParams creates a new DeleteInstancesParams object
// with the default values initialized.
func NewDeleteInstancesParams() *DeleteInstancesParams {
	var ()
	return &DeleteInstancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesParamsWithTimeout creates a new DeleteInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesParamsWithTimeout(timeout time.Duration) *DeleteInstancesParams {
	var ()
	return &DeleteInstancesParams{

		timeout: timeout,
	}
}

// NewDeleteInstancesParamsWithContext creates a new DeleteInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesParamsWithContext(ctx context.Context) *DeleteInstancesParams {
	var ()
	return &DeleteInstancesParams{

		Context: ctx,
	}
}

// NewDeleteInstancesParamsWithHTTPClient creates a new DeleteInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesParamsWithHTTPClient(client *http.Client) *DeleteInstancesParams {
	var ()
	return &DeleteInstancesParams{
		HTTPClient: client,
	}
}

/*DeleteInstancesParams contains all the parameters to send to the API endpoint
for the delete instances operation typically these are written to a http.Request
*/
type DeleteInstancesParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instances params
func (o *DeleteInstancesParams) WithTimeout(timeout time.Duration) *DeleteInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances params
func (o *DeleteInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances params
func (o *DeleteInstancesParams) WithContext(ctx context.Context) *DeleteInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances params
func (o *DeleteInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances params
func (o *DeleteInstancesParams) WithHTTPClient(client *http.Client) *DeleteInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances params
func (o *DeleteInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete instances params
func (o *DeleteInstancesParams) WithAuthorization(authorization string) *DeleteInstancesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete instances params
func (o *DeleteInstancesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
