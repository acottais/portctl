// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new endpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EndpointAccessUpdate manages accesses to an endpoint

Manage user and team accesses to an endpoint.
**Access policy**: administrator

*/
func (a *Client) EndpointAccessUpdate(params *EndpointAccessUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointAccessUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointAccessUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointAccessUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}/access",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointAccessUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointAccessUpdateOK), nil

}

/*
EndpointCreate creates a new endpoint

Create a new endpoint that will be used to manage a Docker environment.
**Access policy**: administrator

*/
func (a *Client) EndpointCreate(params *EndpointCreateParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointCreate",
		Method:             "POST",
		PathPattern:        "/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointCreateOK), nil

}

/*
EndpointDelete removes an endpoint

Remove an endpoint.
**Access policy**: administrator

*/
func (a *Client) EndpointDelete(params *EndpointDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointDelete",
		Method:             "DELETE",
		PathPattern:        "/endpoints/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointDeleteNoContent), nil

}

/*
EndpointInspect inspects an endpoint

Retrieve details abount an endpoint.
**Access policy**: restricted

*/
func (a *Client) EndpointInspect(params *EndpointInspectParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointInspect",
		Method:             "GET",
		PathPattern:        "/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointInspectOK), nil

}

/*
EndpointJob executes a job on the endpoint host

Execute a job (script) on the underlying host of the endpoint.
**Access policy**: administrator

*/
func (a *Client) EndpointJob(params *EndpointJobParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointJob",
		Method:             "POST",
		PathPattern:        "/endpoints/{id}/job",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointJobOK), nil

}

/*
EndpointList lists endpoints

List all endpoints based on the current user authorizations. Will
return all endpoints if using an administrator account otherwise it will
only return authorized endpoints.
**Access policy**: restricted

*/
func (a *Client) EndpointList(params *EndpointListParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointList",
		Method:             "GET",
		PathPattern:        "/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointListOK), nil

}

/*
EndpointUpdate updates an endpoint

Update an endpoint.
**Access policy**: administrator

*/
func (a *Client) EndpointUpdate(params *EndpointUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*EndpointUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
