// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SettingsLDAPCheckRequest settings l d a p check request
// swagger:model SettingsLDAPCheckRequest
type SettingsLDAPCheckRequest struct {

	// l d a p settings
	LDAPSettings *LDAPSettings `json:"LDAPSettings,omitempty"`
}

// Validate validates this settings l d a p check request
func (m *SettingsLDAPCheckRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLDAPSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsLDAPCheckRequest) validateLDAPSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.LDAPSettings) { // not required
		return nil
	}

	if m.LDAPSettings != nil {
		if err := m.LDAPSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LDAPSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingsLDAPCheckRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsLDAPCheckRequest) UnmarshalBinary(b []byte) error {
	var res SettingsLDAPCheckRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
