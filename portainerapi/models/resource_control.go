// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ResourceControl resource control
// swagger:model ResourceControl
type ResourceControl struct {

	// Permit access to the associated resource to any user
	Public bool `json:"Public,omitempty"`

	// Docker resource identifier on which access control will be applied. In the case of a resource control applied to a stack, use the stack name as identifier
	ResourceID string `json:"ResourceID,omitempty"`

	// List of Docker resources that will inherit this access control
	SubResourceIds []string `json:"SubResourceIDs"`

	// List of team identifiers with access to the associated resource
	Teams []int64 `json:"Teams"`

	// Type of Docker resource. Valid values are: container, volume service, secret, config or stack
	Type string `json:"Type,omitempty"`

	// List of user identifiers with access to the associated resource
	Users []int64 `json:"Users"`
}

// Validate validates this resource control
func (m *ResourceControl) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceControl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceControl) UnmarshalBinary(b []byte) error {
	var res ResourceControl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
