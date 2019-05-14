// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// LocalNodeState Current local status of this node.
// swagger:model LocalNodeState
type LocalNodeState string

const (

	// LocalNodeState captures enum value ""
	LocalNodeStateEmpty LocalNodeState = ""

	// LocalNodeStateInactive captures enum value "inactive"
	LocalNodeStateInactive LocalNodeState = "inactive"

	// LocalNodeStatePending captures enum value "pending"
	LocalNodeStatePending LocalNodeState = "pending"

	// LocalNodeStateActive captures enum value "active"
	LocalNodeStateActive LocalNodeState = "active"

	// LocalNodeStateError captures enum value "error"
	LocalNodeStateError LocalNodeState = "error"

	// LocalNodeStateLocked captures enum value "locked"
	LocalNodeStateLocked LocalNodeState = "locked"
)

// for schema
var localNodeStateEnum []interface{}

func init() {
	var res []LocalNodeState
	if err := json.Unmarshal([]byte(`["","inactive","pending","active","error","locked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		localNodeStateEnum = append(localNodeStateEnum, v)
	}
}

func (m LocalNodeState) validateLocalNodeStateEnum(path, location string, value LocalNodeState) error {
	if err := validate.Enum(path, location, value, localNodeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this local node state
func (m LocalNodeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLocalNodeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
