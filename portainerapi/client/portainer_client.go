// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/acottais/portctl/portainerapi/client/auth"
	"github.com/acottais/portctl/portainerapi/client/dockerhub"
	"github.com/acottais/portctl/portainerapi/client/endpoint_groups"
	"github.com/acottais/portctl/portainerapi/client/endpoints"
	"github.com/acottais/portctl/portainerapi/client/registries"
	"github.com/acottais/portctl/portainerapi/client/resource_controls"
	"github.com/acottais/portctl/portainerapi/client/settings"
	"github.com/acottais/portctl/portainerapi/client/stacks"
	"github.com/acottais/portctl/portainerapi/client/status"
	"github.com/acottais/portctl/portainerapi/client/tags"
	"github.com/acottais/portctl/portainerapi/client/team_memberships"
	"github.com/acottais/portctl/portainerapi/client/teams"
	"github.com/acottais/portctl/portainerapi/client/templates"
	"github.com/acottais/portctl/portainerapi/client/upload"
	"github.com/acottais/portctl/portainerapi/client/users"
)

// Default portainer HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "portainer.domain"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new portainer HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Portainer {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new portainer HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Portainer {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new portainer client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Portainer {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Portainer)
	cli.Transport = transport

	cli.Auth = auth.New(transport, formats)

	cli.Dockerhub = dockerhub.New(transport, formats)

	cli.EndpointGroups = endpoint_groups.New(transport, formats)

	cli.Endpoints = endpoints.New(transport, formats)

	cli.Registries = registries.New(transport, formats)

	cli.ResourceControls = resource_controls.New(transport, formats)

	cli.Settings = settings.New(transport, formats)

	cli.Stacks = stacks.New(transport, formats)

	cli.Status = status.New(transport, formats)

	cli.Tags = tags.New(transport, formats)

	cli.TeamMemberships = team_memberships.New(transport, formats)

	cli.Teams = teams.New(transport, formats)

	cli.Templates = templates.New(transport, formats)

	cli.Upload = upload.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Portainer is a client for portainer
type Portainer struct {
	Auth *auth.Client

	Dockerhub *dockerhub.Client

	EndpointGroups *endpoint_groups.Client

	Endpoints *endpoints.Client

	Registries *registries.Client

	ResourceControls *resource_controls.Client

	Settings *settings.Client

	Stacks *stacks.Client

	Status *status.Client

	Tags *tags.Client

	TeamMemberships *team_memberships.Client

	Teams *teams.Client

	Templates *templates.Client

	Upload *upload.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Portainer) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Auth.SetTransport(transport)

	c.Dockerhub.SetTransport(transport)

	c.EndpointGroups.SetTransport(transport)

	c.Endpoints.SetTransport(transport)

	c.Registries.SetTransport(transport)

	c.ResourceControls.SetTransport(transport)

	c.Settings.SetTransport(transport)

	c.Stacks.SetTransport(transport)

	c.Status.SetTransport(transport)

	c.Tags.SetTransport(transport)

	c.TeamMemberships.SetTransport(transport)

	c.Teams.SetTransport(transport)

	c.Templates.SetTransport(transport)

	c.Upload.SetTransport(transport)

	c.Users.SetTransport(transport)

}
