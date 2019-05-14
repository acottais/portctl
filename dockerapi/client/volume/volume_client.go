// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new volume API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for volume API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
VolumeCreate creates a volume
*/
func (a *Client) VolumeCreate(params *VolumeCreateParams) (*VolumeCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VolumeCreate",
		Method:             "POST",
		PathPattern:        "/volumes/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VolumeCreateCreated), nil

}

/*
VolumeDelete removes a volume

Instruct the driver to remove the volume.
*/
func (a *Client) VolumeDelete(params *VolumeDeleteParams) (*VolumeDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VolumeDelete",
		Method:             "DELETE",
		PathPattern:        "/volumes/{name}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VolumeDeleteNoContent), nil

}

/*
VolumeInspect inspects a volume
*/
func (a *Client) VolumeInspect(params *VolumeInspectParams) (*VolumeInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VolumeInspect",
		Method:             "GET",
		PathPattern:        "/volumes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VolumeInspectOK), nil

}

/*
VolumeList lists volumes
*/
func (a *Client) VolumeList(params *VolumeListParams) (*VolumeListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VolumeList",
		Method:             "GET",
		PathPattern:        "/volumes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VolumeListOK), nil

}

/*
VolumePrune deletes unused volumes
*/
func (a *Client) VolumePrune(params *VolumePruneParams) (*VolumePruneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumePruneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VolumePrune",
		Method:             "POST",
		PathPattern:        "/volumes/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumePruneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VolumePruneOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
