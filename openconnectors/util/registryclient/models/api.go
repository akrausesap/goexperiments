// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// API Api
// swagger:model Api
type API struct {

	// API type, for example OData
	APIType string `json:"ApiType,omitempty"`

	// specification Url
	// Format: uri
	SpecificationURL strfmt.URI `json:"SpecificationUrl,omitempty"`

	// credentials
	Credentials *APICredentials `json:"credentials,omitempty"`

	// request parameters
	RequestParameters *RequestParameters `json:"requestParameters,omitempty"`

	// OpenApi v2 swagger file: https://github.com/OAI/OpenAPI-Specification/blob/master/schemas/v2.0/schema.json
	Spec interface{} `json:"spec,omitempty"`

	// specification credentials
	SpecificationCredentials *SpecificationCredentials `json:"specificationCredentials,omitempty"`

	// specification request parameters
	SpecificationRequestParameters *RequestParameters `json:"specificationRequestParameters,omitempty"`

	// target Url
	// Required: true
	// Format: uri
	TargetURL *strfmt.URI `json:"targetUrl"`
}

// Validate validates this Api
func (m *API) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpecificationURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecificationCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecificationRequestParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *API) validateSpecificationURL(formats strfmt.Registry) error {

	if swag.IsZero(m.SpecificationURL) { // not required
		return nil
	}

	if err := validate.FormatOf("SpecificationUrl", "body", "uri", m.SpecificationURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *API) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *API) validateRequestParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestParameters) { // not required
		return nil
	}

	if m.RequestParameters != nil {
		if err := m.RequestParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestParameters")
			}
			return err
		}
	}

	return nil
}

func (m *API) validateSpecificationCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.SpecificationCredentials) { // not required
		return nil
	}

	if m.SpecificationCredentials != nil {
		if err := m.SpecificationCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specificationCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *API) validateSpecificationRequestParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.SpecificationRequestParameters) { // not required
		return nil
	}

	if m.SpecificationRequestParameters != nil {
		if err := m.SpecificationRequestParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specificationRequestParameters")
			}
			return err
		}
	}

	return nil
}

func (m *API) validateTargetURL(formats strfmt.Registry) error {

	if err := validate.Required("targetUrl", "body", m.TargetURL); err != nil {
		return err
	}

	if err := validate.FormatOf("targetUrl", "body", "uri", m.TargetURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *API) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *API) UnmarshalBinary(b []byte) error {
	var res API
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
