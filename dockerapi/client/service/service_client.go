// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServiceCreate creates a service
*/
func (a *Client) ServiceCreate(params *ServiceCreateParams) (*ServiceCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceCreate",
		Method:             "POST",
		PathPattern:        "/services/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceCreateCreated), nil

}

/*
ServiceDelete deletes a service
*/
func (a *Client) ServiceDelete(params *ServiceDeleteParams) (*ServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceDelete",
		Method:             "DELETE",
		PathPattern:        "/services/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceDeleteOK), nil

}

/*
ServiceInspect inspects a service
*/
func (a *Client) ServiceInspect(params *ServiceInspectParams) (*ServiceInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceInspect",
		Method:             "GET",
		PathPattern:        "/services/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceInspectOK), nil

}

/*
ServiceList lists services
*/
func (a *Client) ServiceList(params *ServiceListParams) (*ServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceList",
		Method:             "GET",
		PathPattern:        "/services",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceListOK), nil

}

/*
ServiceLogs gets service logs

Get `stdout` and `stderr` logs from a service. See also [`/containers/{id}/logs`](#operation/ContainerLogs).

**Note**: This endpoint works only for services with the `local`, `json-file` or `journald` logging drivers.

*/
func (a *Client) ServiceLogs(params *ServiceLogsParams, writer io.Writer) (*ServiceLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceLogs",
		Method:             "GET",
		PathPattern:        "/services/{id}/logs",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceLogsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceLogsOK), nil

}

/*
ServiceUpdate updates a service
*/
func (a *Client) ServiceUpdate(params *ServiceUpdateParams) (*ServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceUpdate",
		Method:             "POST",
		PathPattern:        "/services/{id}/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
