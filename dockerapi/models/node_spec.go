// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeSpec node spec
// swagger:model NodeSpec
type NodeSpec struct {

	// Availability of the node.
	// Enum: [active pause drain]
	Availability string `json:"Availability,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// Name for the node.
	Name string `json:"Name,omitempty"`

	// Role of the node.
	// Enum: [worker manager]
	Role string `json:"Role,omitempty"`
}

// Validate validates this node spec
func (m *NodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nodeSpecTypeAvailabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","pause","drain"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeSpecTypeAvailabilityPropEnum = append(nodeSpecTypeAvailabilityPropEnum, v)
	}
}

const (

	// NodeSpecAvailabilityActive captures enum value "active"
	NodeSpecAvailabilityActive string = "active"

	// NodeSpecAvailabilityPause captures enum value "pause"
	NodeSpecAvailabilityPause string = "pause"

	// NodeSpecAvailabilityDrain captures enum value "drain"
	NodeSpecAvailabilityDrain string = "drain"
)

// prop value enum
func (m *NodeSpec) validateAvailabilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nodeSpecTypeAvailabilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NodeSpec) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailabilityEnum("Availability", "body", m.Availability); err != nil {
		return err
	}

	return nil
}

var nodeSpecTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["worker","manager"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeSpecTypeRolePropEnum = append(nodeSpecTypeRolePropEnum, v)
	}
}

const (

	// NodeSpecRoleWorker captures enum value "worker"
	NodeSpecRoleWorker string = "worker"

	// NodeSpecRoleManager captures enum value "manager"
	NodeSpecRoleManager string = "manager"
)

// prop value enum
func (m *NodeSpec) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nodeSpecTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NodeSpec) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("Role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeSpec) UnmarshalBinary(b []byte) error {
	var res NodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
