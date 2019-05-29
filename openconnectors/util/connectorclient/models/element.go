// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Element element
// swagger:model Element
type Element struct {

	// active
	Active bool `json:"active,omitempty"`

	// authentication type
	AuthenticationType string `json:"authenticationType,omitempty"`

	// config description
	ConfigDescription string `json:"configDescription,omitempty"`

	// created date
	CreatedDate float64 `json:"createdDate,omitempty"`

	// deleted
	Deleted bool `json:"deleted,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// existing account description
	ExistingAccountDescription string `json:"existingAccountDescription,omitempty"`

	// hub
	Hub string `json:"hub,omitempty"`

	// id
	ID float64 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// signup URL
	SignupURL string `json:"signupURL,omitempty"`

	// transformations enabled
	TransformationsEnabled bool `json:"transformationsEnabled,omitempty"`

	// trial account
	TrialAccount bool `json:"trialAccount,omitempty"`

	// trial account description
	TrialAccountDescription string `json:"trialAccountDescription,omitempty"`

	// type oauth
	TypeOauth bool `json:"typeOauth,omitempty"`

	// updated date
	UpdatedDate float64 `json:"updatedDate,omitempty"`
}

// Validate validates this element
func (m *Element) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Element) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Element) UnmarshalBinary(b []byte) error {
	var res Element
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
