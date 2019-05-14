// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new registries API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for registries API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RegistryAccessUpdate manages accesses to a registry

Manage user and team accesses to a registry.
**Access policy**: administrator

*/
func (a *Client) RegistryAccessUpdate(params *RegistryAccessUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*RegistryAccessUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryAccessUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegistryAccessUpdate",
		Method:             "PUT",
		PathPattern:        "/registries/{id}/access",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryAccessUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegistryAccessUpdateOK), nil

}

/*
RegistryCreate creates a new registry

Create a new registry.
**Access policy**: administrator

*/
func (a *Client) RegistryCreate(params *RegistryCreateParams, authInfo runtime.ClientAuthInfoWriter) (*RegistryCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegistryCreate",
		Method:             "POST",
		PathPattern:        "/registries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegistryCreateOK), nil

}

/*
RegistryDelete removes a registry

Remove a registry.
**Access policy**: administrator

*/
func (a *Client) RegistryDelete(params *RegistryDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*RegistryDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegistryDelete",
		Method:             "DELETE",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegistryDeleteNoContent), nil

}

/*
RegistryInspect inspects a registry

Retrieve details about a registry.
**Access policy**: administrator

*/
func (a *Client) RegistryInspect(params *RegistryInspectParams, authInfo runtime.ClientAuthInfoWriter) (*RegistryInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegistryInspect",
		Method:             "GET",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegistryInspectOK), nil

}

/*
RegistryList lists registries

List all registries based on the current user authorizations.
Will return all registries if using an administrator account otherwise it
will only return authorized registries.
**Access policy**: restricted

*/
func (a *Client) RegistryList(params *RegistryListParams, authInfo runtime.ClientAuthInfoWriter) (*RegistryListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegistryList",
		Method:             "GET",
		PathPattern:        "/registries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegistryListOK), nil

}

/*
RegistryUpdate updates a registry

Update a registry.
**Access policy**: administrator

*/
func (a *Client) RegistryUpdate(params *RegistryUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*RegistryUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegistryUpdate",
		Method:             "PUT",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegistryUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}