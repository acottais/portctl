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

// Port An open port on a container
// swagger:model Port
type Port struct {

	// Host IP address that the container's port is mapped to
	IP string `json:"IP,omitempty"`

	// Port on the container
	// Required: true
	PrivatePort uint16 `json:"PrivatePort"`

	// Port exposed on the host
	PublicPort uint16 `json:"PublicPort,omitempty"`

	// type
	// Required: true
	// Enum: [tcp udp sctp]
	Type string `json:"Type"`
}

// Validate validates this port
func (m *Port) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Port) validatePrivatePort(formats strfmt.Registry) error {

	if err := validate.Required("PrivatePort", "body", uint16(m.PrivatePort)); err != nil {
		return err
	}

	return nil
}

var portTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp","sctp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portTypeTypePropEnum = append(portTypeTypePropEnum, v)
	}
}

const (

	// PortTypeTCP captures enum value "tcp"
	PortTypeTCP string = "tcp"

	// PortTypeUDP captures enum value "udp"
	PortTypeUDP string = "udp"

	// PortTypeSctp captures enum value "sctp"
	PortTypeSctp string = "sctp"
)

// prop value enum
func (m *Port) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, portTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Port) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("Type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Port) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Port) UnmarshalBinary(b []byte) error {
	var res Port
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
