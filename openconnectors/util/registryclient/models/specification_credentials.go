// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SpecificationCredentials specification credentials
// swagger:model SpecificationCredentials
type SpecificationCredentials struct {

	// basic
	Basic *Basic `json:"basic,omitempty"`

	// oauth
	Oauth *OAuth `json:"oauth,omitempty"`
}

// Validate validates this specification credentials
func (m *SpecificationCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOauth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpecificationCredentials) validateBasic(formats strfmt.Registry) error {

	if swag.IsZero(m.Basic) { // not required
		return nil
	}

	if m.Basic != nil {
		if err := m.Basic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic")
			}
			return err
		}
	}

	return nil
}

func (m *SpecificationCredentials) validateOauth(formats strfmt.Registry) error {

	if swag.IsZero(m.Oauth) { // not required
		return nil
	}

	if m.Oauth != nil {
		if err := m.Oauth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpecificationCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpecificationCredentials) UnmarshalBinary(b []byte) error {
	var res SpecificationCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
