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

// UserCreateRequest user create request
// swagger:model UserCreateRequest
type UserCreateRequest struct {

	// Password
	// Required: true
	Password *string `json:"Password"`

	// User role (1 for administrator account and 2 for regular account)
	// Required: true
	Role *int64 `json:"Role"`

	// Username
	// Required: true
	Username *string `json:"Username"`
}

// Validate validates this user create request
func (m *UserCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCreateRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("Password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *UserCreateRequest) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("Role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *UserCreateRequest) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("Username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCreateRequest) UnmarshalBinary(b []byte) error {
	var res UserCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
