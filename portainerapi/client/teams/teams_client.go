// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new teams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for teams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TeamCreate creates a new team

Create a new team.
**Access policy**: administrator

*/
func (a *Client) TeamCreate(params *TeamCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamCreate",
		Method:             "POST",
		PathPattern:        "/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamCreateOK), nil

}

/*
TeamDelete removes a team

Remove a team.
**Access policy**: administrator

*/
func (a *Client) TeamDelete(params *TeamDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TeamDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamDelete",
		Method:             "DELETE",
		PathPattern:        "/teams/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamDeleteNoContent), nil

}

/*
TeamInspect inspects a team

Retrieve details about a team. Access is only available for administrator and leaders of that team.
**Access policy**: restricted

*/
func (a *Client) TeamInspect(params *TeamInspectParams, authInfo runtime.ClientAuthInfoWriter) (*TeamInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamInspect",
		Method:             "GET",
		PathPattern:        "/teams/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamInspectOK), nil

}

/*
TeamList lists teams

List teams. For non-administrator users, will only list the teams they are member of.
**Access policy**: restricted

*/
func (a *Client) TeamList(params *TeamListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamList",
		Method:             "GET",
		PathPattern:        "/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamListOK), nil

}

/*
TeamMembershipsInspect inspects a team memberships

Inspect a team memberships. Access is only available for administrator and leaders of that team.
**Access policy**: restricted

*/
func (a *Client) TeamMembershipsInspect(params *TeamMembershipsInspectParams, authInfo runtime.ClientAuthInfoWriter) (*TeamMembershipsInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipsInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamMembershipsInspect",
		Method:             "GET",
		PathPattern:        "/teams/{id}/memberships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipsInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamMembershipsInspectOK), nil

}

/*
TeamUpdate updates a team

Update a team.
**Access policy**: administrator

*/
func (a *Client) TeamUpdate(params *TeamUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamUpdate",
		Method:             "PUT",
		PathPattern:        "/teams/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
