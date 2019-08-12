// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CertificateGen certificate gen
// swagger:model CertificateGen
type CertificateGen struct {

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// common name
	CommonName string `json:"commonName,omitempty"`
}

// Validate validates this certificate gen
func (m *CertificateGen) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateGen) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateGen) UnmarshalBinary(b []byte) error {
	var res CertificateGen
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
