// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackCreateRequest stack create request
// swagger:model StackCreateRequest
type StackCreateRequest struct {

	// Path to the Stack file inside the Git repository. Will default to 'docker-compose.yml' if not specified.
	ComposeFilePathInRepository string `json:"ComposeFilePathInRepository,omitempty"`

	// A list of environment variables used during stack deployment
	Env []*StackEnv `json:"Env"`

	// Name of the stack
	// Required: true
	Name *string `json:"Name"`

	// Use basic authentication to clone the Git repository.
	RepositoryAuthentication bool `json:"RepositoryAuthentication,omitempty"`

	// Password used in basic authentication. Required when RepositoryAuthentication is true.
	RepositoryPassword string `json:"RepositoryPassword,omitempty"`

	// Reference name of a Git repository hosting the Stack file. Used in 'repository' deployment method.
	RepositoryReferenceName string `json:"RepositoryReferenceName,omitempty"`

	// URL of a Git repository hosting the Stack file. Required when using the 'repository' deployment method.
	RepositoryURL string `json:"RepositoryURL,omitempty"`

	// Username used in basic authentication. Required when RepositoryAuthentication is true.
	RepositoryUsername string `json:"RepositoryUsername,omitempty"`

	// Content of the Stack file. Required when using the 'string' deployment method.
	StackFileContent string `json:"StackFileContent,omitempty"`

	// Swarm cluster identifier. Required when creating a Swarm stack (type 1).
	SwarmID string `json:"SwarmID,omitempty"`
}

// Validate validates this stack create request
func (m *StackCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackCreateRequest) validateEnv(formats strfmt.Registry) error {

	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackCreateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackCreateRequest) UnmarshalBinary(b []byte) error {
	var res StackCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
