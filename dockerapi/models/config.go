// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Config config
// swagger:model Config
type Config struct {

	// created at
	CreatedAt string `json:"CreatedAt,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// spec
	Spec *ConfigSpec `json:"Spec,omitempty"`

	// updated at
	UpdatedAt string `json:"UpdatedAt,omitempty"`

	// version
	Version *ObjectVersion `json:"Version,omitempty"`
}

// Validate validates this config
func (m *Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Config) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

func (m *Config) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Config) UnmarshalBinary(b []byte) error {
	var res Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (m *Config) String() string {
	return m.ID + " " + m.Spec.Name
}
