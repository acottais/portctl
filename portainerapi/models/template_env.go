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

// TemplateEnv template env
// swagger:model TemplateEnv
type TemplateEnv struct {

	// Default value that will be set for the variable
	Default string `json:"default,omitempty"`

	// Content of the tooltip that will be generated in the UI
	Description string `json:"description,omitempty"`

	// Text for the label that will be generated in the UI
	Label string `json:"label,omitempty"`

	// name of the environment variable
	Name string `json:"name,omitempty"`

	// If set to true, will not generate any input for this variable in the UI
	Preset bool `json:"preset,omitempty"`

	// A list of name/value that will be used to generate a dropdown in the UI
	Select []*TemplateEnvSelect `json:"select"`
}

// Validate validates this template env
func (m *TemplateEnv) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelect(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateEnv) validateSelect(formats strfmt.Registry) error {

	if swag.IsZero(m.Select) { // not required
		return nil
	}

	for i := 0; i < len(m.Select); i++ {
		if swag.IsZero(m.Select[i]) { // not required
			continue
		}

		if m.Select[i] != nil {
			if err := m.Select[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("select" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateEnv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateEnv) UnmarshalBinary(b []byte) error {
	var res TemplateEnv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
