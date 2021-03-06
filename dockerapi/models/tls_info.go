// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TLSInfo Information about the issuer of leaf TLS certificates and the trusted root CA certificate
// swagger:model TLSInfo
type TLSInfo struct {

	// The base64-url-safe-encoded raw public key bytes of the issuer
	CertIssuerPublicKey string `json:"CertIssuerPublicKey,omitempty"`

	// The base64-url-safe-encoded raw subject bytes of the issuer
	CertIssuerSubject string `json:"CertIssuerSubject,omitempty"`

	// The root CA certificate(s) that are used to validate leaf TLS certificates
	TrustRoot string `json:"TrustRoot,omitempty"`
}

// Validate validates this TLS info
func (m *TLSInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TLSInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TLSInfo) UnmarshalBinary(b []byte) error {
	var res TLSInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
