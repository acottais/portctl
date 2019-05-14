// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new team memberships API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for team memberships API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TeamMembershipCreate creates a new team membership

Create a new team memberships. Access is only available to administrators leaders of the associated team.
**Access policy**: restricted

*/
func (a *Client) TeamMembershipCreate(params *TeamMembershipCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamMembershipCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamMembershipCreate",
		Method:             "POST",
		PathPattern:        "/team_memberships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamMembershipCreateOK), nil

}

/*
TeamMembershipDelete removes a team membership

Remove a team membership. Access is only available to administrators leaders of the associated team.
**Access policy**: restricted

*/
func (a *Client) TeamMembershipDelete(params *TeamMembershipDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TeamMembershipDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamMembershipDelete",
		Method:             "DELETE",
		PathPattern:        "/team_memberships/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamMembershipDeleteNoContent), nil

}

/*
TeamMembershipList lists team memberships

List team memberships. Access is only available to administrators and team leaders.
**Access policy**: restricted

*/
func (a *Client) TeamMembershipList(params *TeamMembershipListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamMembershipListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamMembershipList",
		Method:             "GET",
		PathPattern:        "/team_memberships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamMembershipListOK), nil

}

/*
TeamMembershipUpdate updates a team membership

Update a team membership. Access is only available to administrators leaders of the associated team.
**Access policy**: restricted

*/
func (a *Client) TeamMembershipUpdate(params *TeamMembershipUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamMembershipUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TeamMembershipUpdate",
		Method:             "PUT",
		PathPattern:        "/team_memberships/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamMembershipUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}