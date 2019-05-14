// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EndpointSubset endpoint subset
// swagger:model EndpointSubset
type EndpointSubset struct {

	// List of team identifiers authorized to connect to this endpoint
	AuthorizedTeams []int64 `json:"AuthorizedTeams"`

	// List of user identifiers authorized to connect to this endpoint
	AuthorizedUsers []int64 `json:"AuthorizedUsers"`

	// Endpoint group identifier
	GroupID int64 `json:"GroupID,omitempty"`

	// Endpoint identifier
	ID int64 `json:"Id,omitempty"`

	// Endpoint name
	Name string `json:"Name,omitempty"`

	// URL or IP address where exposed containers will be reachable
	PublicURL string `json:"PublicURL,omitempty"`

	// TLS config
	TLSConfig *TLSConfiguration `json:"TLSConfig,omitempty"`

	// Endpoint environment type. 1 for a Docker environment, 2 for an agent on Docker environment, 3 for an Azure environment.
	Type int64 `json:"Type,omitempty"`

	// URL or IP address of the Docker host associated to this endpoint
	URL string `json:"URL,omitempty"`
}

// Validate validates this endpoint subset
func (m *EndpointSubset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointSubset) validateTLSConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointSubset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointSubset) UnmarshalBinary(b []byte) error {
	var res EndpointSubset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}