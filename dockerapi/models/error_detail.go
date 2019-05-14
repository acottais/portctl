// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ErrorDetail error detail
// swagger:model ErrorDetail
type ErrorDetail struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this error detail
func (m *ErrorDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorDetail) UnmarshalBinary(b []byte) error {
	var res ErrorDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}