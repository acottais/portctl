// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Runtime Runtime describes an [OCI compliant](https://github.com/opencontainers/runtime-spec)
// runtime.
//
// The runtime is invoked by the daemon via the `containerd` daemon. OCI
// runtimes act as an interface to the Linux kernel namespaces, cgroups,
// and SELinux.
//
// swagger:model Runtime
type Runtime struct {

	// Name and, optional, path, of the OCI executable binary.
	//
	// If the path is omitted, the daemon searches the host's `$PATH` for the
	// binary and uses the first result.
	//
	Path string `json:"path,omitempty"`

	// List of command-line arguments to pass to the runtime when invoked.
	//
	RuntimeArgs []string `json:"runtimeArgs"`
}

// Validate validates this runtime
func (m *Runtime) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Runtime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Runtime) UnmarshalBinary(b []byte) error {
	var res Runtime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
