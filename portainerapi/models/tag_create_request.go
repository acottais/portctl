// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TagCreateRequest tag create request
// swagger:model TagCreateRequest
type TagCreateRequest struct {

	// Name
	// Required: true
	Name *string `json:"Name"`
}

// Validate validates this tag create request
func (m *TagCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagCreateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TagCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagCreateRequest) UnmarshalBinary(b []byte) error {
	var res TagCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}