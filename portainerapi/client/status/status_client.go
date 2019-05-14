// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
StatusInspect checks portainer status

Retrieve Portainer status.
**Access policy**: public

*/
func (a *Client) StatusInspect(params *StatusInspectParams, authInfo runtime.ClientAuthInfoWriter) (*StatusInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StatusInspect",
		Method:             "GET",
		PathPattern:        "/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusInspectOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}