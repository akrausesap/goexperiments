// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APICredentialsUpdate Api credentials update
// swagger:model ApiCredentialsUpdate
type APICredentialsUpdate struct {

	// basic
	Basic *Basic `json:"basic,omitempty"`

	// certificate gen
	CertificateGen *CertificateGenUpdate `json:"certificateGen,omitempty"`

	// oauth
	Oauth *OAuth `json:"oauth,omitempty"`
}

// Validate validates this Api credentials update
func (m *APICredentialsUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateGen(formats); err != nil {
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

func (m *APICredentialsUpdate) validateBasic(formats strfmt.Registry) error {

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

func (m *APICredentialsUpdate) validateCertificateGen(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificateGen) { // not required
		return nil
	}

	if m.CertificateGen != nil {
		if err := m.CertificateGen.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateGen")
			}
			return err
		}
	}

	return nil
}

func (m *APICredentialsUpdate) validateOauth(formats strfmt.Registry) error {

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
func (m *APICredentialsUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICredentialsUpdate) UnmarshalBinary(b []byte) error {
	var res APICredentialsUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
