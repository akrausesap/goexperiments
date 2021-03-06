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

// NewGetInstancesDocByOperationIDParams creates a new GetInstancesDocByOperationIDParams object
// with the default values initialized.
func NewGetInstancesDocByOperationIDParams() *GetInstancesDocByOperationIDParams {
	var (
		versionDefault = string("-1")
	)
	return &GetInstancesDocByOperationIDParams{
		Version: &versionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesDocByOperationIDParamsWithTimeout creates a new GetInstancesDocByOperationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesDocByOperationIDParamsWithTimeout(timeout time.Duration) *GetInstancesDocByOperationIDParams {
	var (
		versionDefault = string("-1")
	)
	return &GetInstancesDocByOperationIDParams{
		Version: &versionDefault,

		timeout: timeout,
	}
}

// NewGetInstancesDocByOperationIDParamsWithContext creates a new GetInstancesDocByOperationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesDocByOperationIDParamsWithContext(ctx context.Context) *GetInstancesDocByOperationIDParams {
	var (
		versionDefault = string("-1")
	)
	return &GetInstancesDocByOperationIDParams{
		Version: &versionDefault,

		Context: ctx,
	}
}

// NewGetInstancesDocByOperationIDParamsWithHTTPClient creates a new GetInstancesDocByOperationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesDocByOperationIDParamsWithHTTPClient(client *http.Client) *GetInstancesDocByOperationIDParams {
	var (
		versionDefault = string("-1")
	)
	return &GetInstancesDocByOperationIDParams{
		Version:    &versionDefault,
		HTTPClient: client,
	}
}

/*GetInstancesDocByOperationIDParams contains all the parameters to send to the API endpoint
for the get instances doc by operation Id operation typically these are written to a http.Request
*/
type GetInstancesDocByOperationIDParams struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;, Element &lt;token&gt;'

	*/
	Authorization string
	/*Basic
	  Include only OpenAPI / Swagger properties in definitions

	*/
	Basic *bool
	/*Discovery
	  Include discovery metadata in definitions

	*/
	Discovery *bool
	/*OperationID
	  The id of the operation for which swagger docs is returned

	*/
	OperationID string
	/*Version
	  The element swagger version to get the corresponding element swagger, Passing in "-1" gives latest element swagger

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithTimeout(timeout time.Duration) *GetInstancesDocByOperationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithContext(ctx context.Context) *GetInstancesDocByOperationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithHTTPClient(client *http.Client) *GetInstancesDocByOperationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithAuthorization(authorization string) *GetInstancesDocByOperationIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBasic adds the basic to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithBasic(basic *bool) *GetInstancesDocByOperationIDParams {
	o.SetBasic(basic)
	return o
}

// SetBasic adds the basic to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetBasic(basic *bool) {
	o.Basic = basic
}

// WithDiscovery adds the discovery to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithDiscovery(discovery *bool) *GetInstancesDocByOperationIDParams {
	o.SetDiscovery(discovery)
	return o
}

// SetDiscovery adds the discovery to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetDiscovery(discovery *bool) {
	o.Discovery = discovery
}

// WithOperationID adds the operationID to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithOperationID(operationID string) *GetInstancesDocByOperationIDParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetOperationID(operationID string) {
	o.OperationID = operationID
}

// WithVersion adds the version to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) WithVersion(version *string) *GetInstancesDocByOperationIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances doc by operation Id params
func (o *GetInstancesDocByOperationIDParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesDocByOperationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Basic != nil {

		// query param basic
		var qrBasic bool
		if o.Basic != nil {
			qrBasic = *o.Basic
		}
		qBasic := swag.FormatBool(qrBasic)
		if qBasic != "" {
			if err := r.SetQueryParam("basic", qBasic); err != nil {
				return err
			}
		}

	}

	if o.Discovery != nil {

		// query param discovery
		var qrDiscovery bool
		if o.Discovery != nil {
			qrDiscovery = *o.Discovery
		}
		qDiscovery := swag.FormatBool(qrDiscovery)
		if qDiscovery != "" {
			if err := r.SetQueryParam("discovery", qDiscovery); err != nil {
				return err
			}
		}

	}

	// path param operationId
	if err := r.SetPathParam("operationId", o.OperationID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
