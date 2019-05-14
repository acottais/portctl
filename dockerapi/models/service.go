// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Service service
// swagger:model Service
type Service struct {

	// created at
	CreatedAt string `json:"CreatedAt,omitempty"`

	// endpoint
	Endpoint *ServiceEndpoint `json:"Endpoint,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// spec
	Spec *ServiceSpec `json:"Spec,omitempty"`

	// update status
	UpdateStatus *ServiceUpdateStatus `json:"UpdateStatus,omitempty"`

	// updated at
	UpdatedAt string `json:"UpdatedAt,omitempty"`

	// version
	Version *ObjectVersion `json:"Version,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) validateEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	if m.Endpoint != nil {
		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *Service) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

func (m *Service) validateUpdateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateStatus) { // not required
		return nil
	}

	if m.UpdateStatus != nil {
		if err := m.UpdateStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UpdateStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Service) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Service) UnmarshalBinary(b []byte) error {
	var res Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceEndpoint service endpoint
// swagger:model ServiceEndpoint
type ServiceEndpoint struct {

	// ports
	Ports []*EndpointPortConfig `json:"Ports"`

	// spec
	Spec *EndpointSpec `json:"Spec,omitempty"`

	// virtual ips
	VirtualIps []*ServiceEndpointVirtualIpsItems0 `json:"VirtualIPs"`
}

// Validate validates this service endpoint
func (m *ServiceEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualIps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEndpoint) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Endpoint" + "." + "Ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceEndpoint) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Endpoint" + "." + "Spec")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEndpoint) validateVirtualIps(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualIps) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualIps); i++ {
		if swag.IsZero(m.VirtualIps[i]) { // not required
			continue
		}

		if m.VirtualIps[i] != nil {
			if err := m.VirtualIps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Endpoint" + "." + "VirtualIPs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceEndpoint) UnmarshalBinary(b []byte) error {
	var res ServiceEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceEndpointVirtualIpsItems0 service endpoint virtual ips items0
// swagger:model ServiceEndpointVirtualIpsItems0
type ServiceEndpointVirtualIpsItems0 struct {

	// addr
	Addr string `json:"Addr,omitempty"`

	// network ID
	NetworkID string `json:"NetworkID,omitempty"`
}

// Validate validates this service endpoint virtual ips items0
func (m *ServiceEndpointVirtualIpsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceEndpointVirtualIpsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceEndpointVirtualIpsItems0) UnmarshalBinary(b []byte) error {
	var res ServiceEndpointVirtualIpsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceUpdateStatus The status of a service update.
// swagger:model ServiceUpdateStatus
type ServiceUpdateStatus struct {

	// completed at
	CompletedAt string `json:"CompletedAt,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// started at
	StartedAt string `json:"StartedAt,omitempty"`

	// state
	// Enum: [updating paused completed]
	State string `json:"State,omitempty"`
}

// Validate validates this service update status
func (m *ServiceUpdateStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceUpdateStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["updating","paused","completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceUpdateStatusTypeStatePropEnum = append(serviceUpdateStatusTypeStatePropEnum, v)
	}
}

const (

	// ServiceUpdateStatusStateUpdating captures enum value "updating"
	ServiceUpdateStatusStateUpdating string = "updating"

	// ServiceUpdateStatusStatePaused captures enum value "paused"
	ServiceUpdateStatusStatePaused string = "paused"

	// ServiceUpdateStatusStateCompleted captures enum value "completed"
	ServiceUpdateStatusStateCompleted string = "completed"
)

// prop value enum
func (m *ServiceUpdateStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceUpdateStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceUpdateStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("UpdateStatus"+"."+"State", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceUpdateStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceUpdateStatus) UnmarshalBinary(b []byte) error {
	var res ServiceUpdateStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}