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

// RestartPolicy The behavior to apply when the container exits. The default is not to restart.
//
// An ever increasing delay (double the previous delay, starting at 100ms) is added before each restart to prevent flooding the server.
//
// swagger:model RestartPolicy
type RestartPolicy struct {

	// If `on-failure` is used, the number of times to retry before giving up
	MaximumRetryCount int64 `json:"MaximumRetryCount,omitempty"`

	// - Empty string means not to restart
	// - `always` Always restart
	// - `unless-stopped` Restart always except when the user has manually stopped the container
	// - `on-failure` Restart only when the container exit code is non-zero
	//
	// Enum: [ always unless-stopped on-failure]
	Name string `json:"Name,omitempty"`
}

// Validate validates this restart policy
func (m *RestartPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var restartPolicyTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","always","unless-stopped","on-failure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restartPolicyTypeNamePropEnum = append(restartPolicyTypeNamePropEnum, v)
	}
}

const (

	// RestartPolicyName captures enum value ""
	RestartPolicyName string = ""

	// RestartPolicyNameAlways captures enum value "always"
	RestartPolicyNameAlways string = "always"

	// RestartPolicyNameUnlessStopped captures enum value "unless-stopped"
	RestartPolicyNameUnlessStopped string = "unless-stopped"

	// RestartPolicyNameOnFailure captures enum value "on-failure"
	RestartPolicyNameOnFailure string = "on-failure"
)

// prop value enum
func (m *RestartPolicy) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, restartPolicyTypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RestartPolicy) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestartPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestartPolicy) UnmarshalBinary(b []byte) error {
	var res RestartPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}