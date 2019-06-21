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

// APIUpdate ApiUpdate
// swagger:model ApiUpdate
type APIUpdate struct {

	// API type, for example OData
	APIType string `json:"ApiType,omitempty"`

	// specification Url
	SpecificationURL string `json:"SpecificationUrl,omitempty"`

	// credentials
	Credentials *APICredentialsUpdate `json:"credentials,omitempty"`

	// OpenApi v2 swagger file: https://github.com/OAI/OpenAPI-Specification/blob/master/schemas/v2.0/schema.json
	Spec interface{} `json:"spec,omitempty"`

	// target Url
	// Required: true
	TargetURL *string `json:"targetUrl"`
}

// Validate validates this Api update
func (m *APIUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
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

func (m *APIUpdate) validateCredentials(formats strfmt.Registry) error {

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

func (m *APIUpdate) validateTargetURL(formats strfmt.Registry) error {

	if err := validate.Required("targetUrl", "body", m.TargetURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIUpdate) UnmarshalBinary(b []byte) error {
	var res APIUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
