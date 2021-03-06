// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ElementInstanceConfigUpdate element instance config update
// swagger:model ElementInstanceConfigUpdate
type ElementInstanceConfigUpdate struct {

	// property value
	PropertyValue string `json:"propertyValue,omitempty"`
}

// Validate validates this element instance config update
func (m *ElementInstanceConfigUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElementInstanceConfigUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElementInstanceConfigUpdate) UnmarshalBinary(b []byte) error {
	var res ElementInstanceConfigUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
