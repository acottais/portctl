// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UserAdminCheck checks administrator account existence

Check if an administrator account exists in the database.
**Access policy**: public

*/
func (a *Client) UserAdminCheck(params *UserAdminCheckParams, authInfo runtime.ClientAuthInfoWriter) (*UserAdminCheckNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAdminCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserAdminCheck",
		Method:             "GET",
		PathPattern:        "/users/admin/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserAdminCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserAdminCheckNoContent), nil

}

/*
UserAdminInit initializes administrator account

Initialize the 'admin' user account.
**Access policy**: public

*/
func (a *Client) UserAdminInit(params *UserAdminInitParams, authInfo runtime.ClientAuthInfoWriter) (*UserAdminInitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAdminInitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserAdminInit",
		Method:             "POST",
		PathPattern:        "/users/admin/init",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserAdminInitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserAdminInitOK), nil

}

/*
UserCreate creates a new user

Create a new Portainer user. Only team leaders and administrators can create users. Only administrators can
create an administrator user account.
**Access policy**: restricted

*/
func (a *Client) UserCreate(params *UserCreateParams, authInfo runtime.ClientAuthInfoWriter) (*UserCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserCreate",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCreateOK), nil

}

/*
UserDelete removes a user

Remove a user.
**Access policy**: administrator

*/
func (a *Client) UserDelete(params *UserDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*UserDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserDelete",
		Method:             "DELETE",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserDeleteNoContent), nil

}

/*
UserInspect inspects a user

Retrieve details about a user.
**Access policy**: administrator

*/
func (a *Client) UserInspect(params *UserInspectParams, authInfo runtime.ClientAuthInfoWriter) (*UserInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserInspect",
		Method:             "GET",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserInspectOK), nil

}

/*
UserList lists users

List Portainer users. Non-administrator users will only be able to list other non-administrator user accounts.
**Access policy**: restricted

*/
func (a *Client) UserList(params *UserListParams, authInfo runtime.ClientAuthInfoWriter) (*UserListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserList",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListOK), nil

}

/*
UserMembershipsInspect inspects a user memberships

Inspect a user memberships.
**Access policy**: authenticated

*/
func (a *Client) UserMembershipsInspect(params *UserMembershipsInspectParams, authInfo runtime.ClientAuthInfoWriter) (*UserMembershipsInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserMembershipsInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserMembershipsInspect",
		Method:             "GET",
		PathPattern:        "/users/{id}/memberships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserMembershipsInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserMembershipsInspectOK), nil

}

/*
UserPasswordCheck checks password validity for a user

Check if the submitted password is valid for the specified user.
**Access policy**: authenticated

*/
func (a *Client) UserPasswordCheck(params *UserPasswordCheckParams, authInfo runtime.ClientAuthInfoWriter) (*UserPasswordCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserPasswordCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserPasswordCheck",
		Method:             "POST",
		PathPattern:        "/users/{id}/passwd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserPasswordCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserPasswordCheckOK), nil

}

/*
UserUpdate updates a user

Update user details. A regular user account can only update his details.
**Access policy**: authenticated

*/
func (a *Client) UserUpdate(params *UserUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*UserUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserUpdate",
		Method:             "PUT",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
