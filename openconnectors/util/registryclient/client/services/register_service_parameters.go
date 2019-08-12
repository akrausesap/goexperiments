// Code generated by go-swagger; DO NOT EDIT.

package services

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

	models "github.com/akrausesap/goexperiments/openconnectors/util/registryclient/models"
)

// NewRegisterServiceParams creates a new RegisterServiceParams object
// with the default values initialized.
func NewRegisterServiceParams() *RegisterServiceParams {
	var ()
	return &RegisterServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterServiceParamsWithTimeout creates a new RegisterServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterServiceParamsWithTimeout(timeout time.Duration) *RegisterServiceParams {
	var ()
	return &RegisterServiceParams{

		timeout: timeout,
	}
}

// NewRegisterServiceParamsWithContext creates a new RegisterServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterServiceParamsWithContext(ctx context.Context) *RegisterServiceParams {
	var ()
	return &RegisterServiceParams{

		Context: ctx,
	}
}

// NewRegisterServiceParamsWithHTTPClient creates a new RegisterServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterServiceParamsWithHTTPClient(client *http.Client) *RegisterServiceParams {
	var ()
	return &RegisterServiceParams{
		HTTPClient: client,
	}
}

/*RegisterServiceParams contains all the parameters to send to the API endpoint
for the register service operation typically these are written to a http.Request
*/
type RegisterServiceParams struct {

	/*Body
	  Service object to be registered

	*/
	Body *models.ServiceDetails

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register service params
func (o *RegisterServiceParams) WithTimeout(timeout time.Duration) *RegisterServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register service params
func (o *RegisterServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register service params
func (o *RegisterServiceParams) WithContext(ctx context.Context) *RegisterServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register service params
func (o *RegisterServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register service params
func (o *RegisterServiceParams) WithHTTPClient(client *http.Client) *RegisterServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register service params
func (o *RegisterServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register service params
func (o *RegisterServiceParams) WithBody(body *models.ServiceDetails) *RegisterServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register service params
func (o *RegisterServiceParams) SetBody(body *models.ServiceDetails) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
