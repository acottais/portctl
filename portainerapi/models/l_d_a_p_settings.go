// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LDAPSettings l d a p settings
// swagger:model LDAPSettings
type LDAPSettings struct {

	// Automatically provision users and assign them to matching LDAP group names
	AutoCreateUsers bool `json:"AutoCreateUsers,omitempty"`

	// group search settings
	GroupSearchSettings []*LDAPGroupSearchSettings `json:"GroupSearchSettings"`

	// Password of the account that will be used to search users
	Password string `json:"Password,omitempty"`

	// Account that will be used to search for users
	ReaderDN string `json:"ReaderDN,omitempty"`

	// search settings
	SearchSettings []*LDAPSearchSettings `json:"SearchSettings"`

	// Whether LDAP connection should use StartTLS
	StartTLS bool `json:"StartTLS,omitempty"`

	// TLS config
	TLSConfig *TLSConfiguration `json:"TLSConfig,omitempty"`

	// URL or IP address of the LDAP server
	URL string `json:"URL,omitempty"`
}

// Validate validates this l d a p settings
func (m *LDAPSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupSearchSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LDAPSettings) validateGroupSearchSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupSearchSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupSearchSettings); i++ {
		if swag.IsZero(m.GroupSearchSettings[i]) { // not required
			continue
		}

		if m.GroupSearchSettings[i] != nil {
			if err := m.GroupSearchSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GroupSearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LDAPSettings) validateSearchSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.SearchSettings); i++ {
		if swag.IsZero(m.SearchSettings[i]) { // not required
			continue
		}

		if m.SearchSettings[i] != nil {
			if err := m.SearchSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LDAPSettings) validateTLSConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LDAPSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LDAPSettings) UnmarshalBinary(b []byte) error {
	var res LDAPSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
