// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TeamMembership team membership
// swagger:model TeamMembership
type TeamMembership struct {

	// Membership identifier
	ID int64 `json:"Id,omitempty"`

	// Team role (1 for team leader and 2 for team member)
	Role int64 `json:"Role,omitempty"`

	// Team identifier
	TeamID int64 `json:"TeamID,omitempty"`

	// User identifier
	UserID int64 `json:"UserID,omitempty"`
}

// Validate validates this team membership
func (m *TeamMembership) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamMembership) UnmarshalBinary(b []byte) error {
	var res TeamMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}