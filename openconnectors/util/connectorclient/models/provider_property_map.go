// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProviderPropertyMap provider property map
// swagger:model ProviderPropertyMap
type ProviderPropertyMap struct {

	// replace your provider property name
	ReplaceYourProviderPropertyName string `json:"<replace_your_provider_property_name>,omitempty"`
}

// Validate validates this provider property map
func (m *ProviderPropertyMap) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProviderPropertyMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderPropertyMap) UnmarshalBinary(b []byte) error {
	var res ProviderPropertyMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
