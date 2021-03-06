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

// NewGetInstancesEventByEventId2Params creates a new GetInstancesEventByEventId2Params object
// with the default values initialized.
func NewGetInstancesEventByEventId2Params() *GetInstancesEventByEventId2Params {
	var ()
	return &GetInstancesEventByEventId2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesEventByEventId2ParamsWithTimeout creates a new GetInstancesEventByEventId2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesEventByEventId2ParamsWithTimeout(timeout time.Duration) *GetInstancesEventByEventId2Params {
	var ()
	return &GetInstancesEventByEventId2Params{

		timeout: timeout,
	}
}

// NewGetInstancesEventByEventId2ParamsWithContext creates a new GetInstancesEventByEventId2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesEventByEventId2ParamsWithContext(ctx context.Context) *GetInstancesEventByEventId2Params {
	var ()
	return &GetInstancesEventByEventId2Params{

		Context: ctx,
	}
}

// NewGetInstancesEventByEventId2ParamsWithHTTPClient creates a new GetInstancesEventByEventId2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesEventByEventId2ParamsWithHTTPClient(client *http.Client) *GetInstancesEventByEventId2Params {
	var ()
	return &GetInstancesEventByEventId2Params{
		HTTPClient: client,
	}
}

/*GetInstancesEventByEventId2Params contains all the parameters to send to the API endpoint
for the get instances event by event id2 operation typically these are written to a http.Request
*/
type GetInstancesEventByEventId2Params struct {

	/*Authorization
	  The authorization tokens. The format for the header value is 'User &lt;user secret&gt;, Organization &lt;org secret&gt;'

	*/
	Authorization string
	/*EventID
	  The ID of the event

	*/
	EventID int64
	/*ID
	  The ID of the element instance

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) WithTimeout(timeout time.Duration) *GetInstancesEventByEventId2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) WithContext(ctx context.Context) *GetInstancesEventByEventId2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) WithHTTPClient(client *http.Client) *GetInstancesEventByEventId2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) WithAuthorization(authorization string) *GetInstancesEventByEventId2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithEventID adds the eventID to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) WithEventID(eventID int64) *GetInstancesEventByEventId2Params {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) SetEventID(eventID int64) {
	o.EventID = eventID
}

// WithID adds the id to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) WithID(id int64) *GetInstancesEventByEventId2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instances event by event id2 params
func (o *GetInstancesEventByEventId2Params) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesEventByEventId2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param eventId
	if err := r.SetPathParam("eventId", swag.FormatInt64(o.EventID)); err != nil {
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
