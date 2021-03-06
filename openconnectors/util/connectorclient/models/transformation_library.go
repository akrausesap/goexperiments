// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TransformationLibrary transformation library
// swagger:model TransformationLibrary
type TransformationLibrary struct {

	// object name
	ObjectName *Transformation `json:"<object_name>,omitempty"`
}

// Validate validates this transformation library
func (m *TransformationLibrary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransformationLibrary) validateObjectName(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectName) { // not required
		return nil
	}

	if m.ObjectName != nil {
		if err := m.ObjectName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("<object_name>")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransformationLibrary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransformationLibrary) UnmarshalBinary(b []byte) error {
	var res TransformationLibrary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
