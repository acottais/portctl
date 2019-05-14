// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DeviceRequest A request for devices to be sent to device drivers
// swagger:model DeviceRequest
type DeviceRequest struct {

	// A list of capabilities; an OR list of AND lists of capabilities.
	//
	Capabilities [][]string `json:"Capabilities"`

	// count
	Count int64 `json:"Count,omitempty"`

	// device ids
	DeviceIds []string `json:"DeviceIDs"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// Driver-specific options, specified as a key/value pairs. These options
	// are passed directly to the driver.
	//
	Options map[string]string `json:"Options,omitempty"`
}

// Validate validates this device request
func (m *DeviceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceRequest) UnmarshalBinary(b []byte) error {
	var res DeviceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
