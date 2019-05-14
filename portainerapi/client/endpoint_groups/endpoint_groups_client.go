// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new endpoint groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoint groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EndpointGroupAccessUpdate manages accesses to an endpoint group

Manage user and team accesses to an endpoint group.
**Access policy**: administrator

*/
func (a *Client) EndpointGroupAccessUpdate(params *EndpointGroupAccessUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointGroupAccessUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupAccessUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointGroupAccessUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoint_groups/{id}/access",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupAccessUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointGroupAccessUpdateOK), nil

}

/*
EndpointGroupCreate creates a new endpoint

Create a new endpoint group.
**Access policy**: administrator

*/
func (a *Client) EndpointGroupCreate(params *EndpointGroupCreateParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointGroupCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointGroupCreate",
		Method:             "POST",
		PathPattern:        "/endpoint_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointGroupCreateOK), nil

}

/*
EndpointGroupDelete removes an endpoint group

Remove an endpoint group.
**Access policy**: administrator

*/
func (a *Client) EndpointGroupDelete(params *EndpointGroupDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointGroupDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointGroupDelete",
		Method:             "DELETE",
		PathPattern:        "/endpoint_groups/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointGroupDeleteNoContent), nil

}

/*
EndpointGroupInspect inspects an endpoint group

Retrieve details abount an endpoint group.
**Access policy**: administrator

*/
func (a *Client) EndpointGroupInspect(params *EndpointGroupInspectParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointGroupInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointGroupInspect",
		Method:             "GET",
		PathPattern:        "/endpoint_groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointGroupInspectOK), nil

}

/*
EndpointGroupList lists endpoint groups

List all endpoint groups based on the current user authorizations. Will
return all endpoint groups if using an administrator account otherwise it will
only return authorized endpoint groups.
**Access policy**: restricted

*/
func (a *Client) EndpointGroupList(params *EndpointGroupListParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointGroupListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointGroupList",
		Method:             "GET",
		PathPattern:        "/endpoint_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointGroupListOK), nil

}

/*
EndpointGroupUpdate updates an endpoint group

Update an endpoint group.
**Access policy**: administrator

*/
func (a *Client) EndpointGroupUpdate(params *EndpointGroupUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointGroupUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointGroupUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoint_groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointGroupUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
