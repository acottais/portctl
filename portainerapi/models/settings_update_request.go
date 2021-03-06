// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SettingsUpdateRequest settings update request
// swagger:model SettingsUpdateRequest
type SettingsUpdateRequest struct {

	// Whether non-administrator users should be able to use bind mounts when creating containers
	AllowBindMountsForRegularUsers bool `json:"AllowBindMountsForRegularUsers,omitempty"`

	// Whether non-administrator users should be able to use privileged mode when creating containers
	AllowPrivilegedModeForRegularUsers bool `json:"AllowPrivilegedModeForRegularUsers,omitempty"`

	// Active authentication method for the Portainer instance. Valid values are: 1 for managed or 2 for LDAP.
	// Required: true
	AuthenticationMethod *int64 `json:"AuthenticationMethod"`

	// A list of label name & value that will be used to hide containers when querying containers
	BlackListedLabels []*SettingsBlackListedLabels `json:"BlackListedLabels"`

	// Whether to display or not external templates contributions as sub-menus in the UI.
	DisplayExternalContributors bool `json:"DisplayExternalContributors,omitempty"`

	// l d a p settings
	LDAPSettings *LDAPSettings `json:"LDAPSettings,omitempty"`

	// URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string
	LogoURL string `json:"LogoURL,omitempty"`

	// URL to the templates that will be displayed in the UI when navigating to App Templates
	// Required: true
	TemplatesURL *string `json:"TemplatesURL"`
}

// Validate validates this settings update request
func (m *SettingsUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlackListedLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLDAPSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplatesURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsUpdateRequest) validateAuthenticationMethod(formats strfmt.Registry) error {

	if err := validate.Required("AuthenticationMethod", "body", m.AuthenticationMethod); err != nil {
		return err
	}

	return nil
}

func (m *SettingsUpdateRequest) validateBlackListedLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.BlackListedLabels) { // not required
		return nil
	}

	for i := 0; i < len(m.BlackListedLabels); i++ {
		if swag.IsZero(m.BlackListedLabels[i]) { // not required
			continue
		}

		if m.BlackListedLabels[i] != nil {
			if err := m.BlackListedLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlackListedLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SettingsUpdateRequest) validateLDAPSettings(formats strfmt.Registry) error {

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

func (m *SettingsUpdateRequest) validateTemplatesURL(formats strfmt.Registry) error {

	if err := validate.Required("TemplatesURL", "body", m.TemplatesURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingsUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsUpdateRequest) UnmarshalBinary(b []byte) error {
	var res SettingsUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
