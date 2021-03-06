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

// EndpointGroup endpoint group
// swagger:model EndpointGroup
type EndpointGroup struct {

	// List of team identifiers authorized to connect to this endpoint. Will be inherited by endpoints that are part of the group
	AuthorizedTeams []int64 `json:"AuthorizedTeams"`

	// List of user identifiers authorized to connect to this endpoint group. Will be inherited by endpoints that are part of the group
	AuthorizedUsers []int64 `json:"AuthorizedUsers"`

	// Endpoint group description
	Description string `json:"Description,omitempty"`

	// Endpoint group identifier
	ID int64 `json:"Id,omitempty"`

	// labels
	Labels []*Pair `json:"Labels"`

	// Endpoint group name
	Name string `json:"Name,omitempty"`
}

// Validate validates this endpoint group
func (m *EndpointGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointGroup) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointGroup) UnmarshalBinary(b []byte) error {
	var res EndpointGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
