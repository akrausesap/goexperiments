// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BulkLoad BulkLoad response of bulk upload.
// swagger:model BulkLoad
type BulkLoad struct {

	// Id of the bulk job
	ID float64 `json:"id,omitempty"`

	// Status of the bulk job
	Status string `json:"status,omitempty"`
}

// Validate validates this bulk load
func (m *BulkLoad) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BulkLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkLoad) UnmarshalBinary(b []byte) error {
	var res BulkLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}