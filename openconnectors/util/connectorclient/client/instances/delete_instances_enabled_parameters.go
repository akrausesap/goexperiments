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

// NewDeleteInstancesEnabledParams creates a new DeleteInstancesEnabledParams object
// with the default values initialized.
func NewDeleteInstancesEnabledParams() *DeleteInstancesEnabledParams {
	var ()
	return &DeleteInstancesEnabledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesEnabledParamsWithTimeout creates a new DeleteInstancesEnabledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesEnabledParamsWithTimeout(timeout time.Duration) *DeleteInstancesEnabledParams {
	var ()
	return &DeleteInstancesEnabledParams{

		timeout: timeout,
	}
}

// NewDeleteInstancesEnabledParamsWithContext creates a new DeleteInstancesEnabledParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesEnabledParamsWithContext(ctx context.Context) *DeleteInstancesEnabledParams {
	var ()
	return &DeleteInstancesEnabledParams{

		Context: ctx,
	}
}

// NewDeleteInstancesEnabledParamsWithHTTPClient creates a new DeleteInstancesEnabledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesEnabledParamsWithHTTPClient(client *http.Client) *DeleteInstancesEnabledParams {
	var ()
	return &DeleteInstancesEnabledParams{
		HTTPClient: client,
	}
}

/*DeleteInstancesEnabledParams contains all the parameters to send to the API endpoint
for the delete instances enabled operation typically these are written to a http.Request
*/
type DeleteInstancesEnabledParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) WithTimeout(timeout time.Duration) *DeleteInstancesEnabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) WithContext(ctx context.Context) *DeleteInstancesEnabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) WithHTTPClient(client *http.Client) *DeleteInstancesEnabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) WithAuthorization(authorization string) *DeleteInstancesEnabledParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete instances enabled params
func (o *DeleteInstancesEnabledParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesEnabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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