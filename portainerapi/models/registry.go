// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Registry registry
// swagger:model Registry
type Registry struct {

	// Is authentication against this registry enabled
	Authentication bool `json:"Authentication,omitempty"`

	// List of team identifiers authorized to use this registry
	AuthorizedTeams []int64 `json:"AuthorizedTeams"`

	// List of user identifiers authorized to use this registry
	AuthorizedUsers []int64 `json:"AuthorizedUsers"`

	// Registry identifier
	ID int64 `json:"Id,omitempty"`

	// Registry name
	Name string `json:"Name,omitempty"`

	// Password used to authenticate against this registry
	Password string `json:"Password,omitempty"`

	// URL or IP address of the Docker registry
	URL string `json:"URL,omitempty"`

	// Username used to authenticate against this registry
	Username string `json:"Username,omitempty"`
}

// Validate validates this registry
func (m *Registry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Registry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Registry) UnmarshalBinary(b []byte) error {
	var res Registry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
