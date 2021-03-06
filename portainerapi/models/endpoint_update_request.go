// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EndpointUpdateRequest endpoint update request
// swagger:model EndpointUpdateRequest
type EndpointUpdateRequest struct {

	// Azure application ID
	ApplicationID string `json:"ApplicationID,omitempty"`

	// Azure authentication key
	AuthenticationKey string `json:"AuthenticationKey,omitempty"`

	// Group identifier
	GroupID int64 `json:"GroupID,omitempty"`

	// Name that will be used to identify this endpoint
	Name string `json:"Name,omitempty"`

	// URL or IP address where exposed containers will be reachable. Defaults to URL if not specified
	PublicURL string `json:"PublicURL,omitempty"`

	// Require TLS to connect against this endpoint
	TLS bool `json:"TLS,omitempty"`

	// Skip client verification when using TLS
	TLSSkipClientVerify bool `json:"TLSSkipClientVerify,omitempty"`

	// Skip server verification when using TLS
	TLSSkipVerify bool `json:"TLSSkipVerify,omitempty"`

	// Azure tenant ID
	TenantID string `json:"TenantID,omitempty"`

	// URL or IP address of a Docker host
	URL string `json:"URL,omitempty"`
}

// Validate validates this endpoint update request
func (m *EndpointUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointUpdateRequest) UnmarshalBinary(b []byte) error {
	var res EndpointUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
