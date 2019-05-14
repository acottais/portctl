// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PushImageInfo push image info
// swagger:model PushImageInfo
type PushImageInfo struct {

	// error
	Error string `json:"error,omitempty"`

	// progress
	Progress string `json:"progress,omitempty"`

	// progress detail
	ProgressDetail *ProgressDetail `json:"progressDetail,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this push image info
func (m *PushImageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProgressDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PushImageInfo) validateProgressDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgressDetail) { // not required
		return nil
	}

	if m.ProgressDetail != nil {
		if err := m.ProgressDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progressDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PushImageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PushImageInfo) UnmarshalBinary(b []byte) error {
	var res PushImageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
