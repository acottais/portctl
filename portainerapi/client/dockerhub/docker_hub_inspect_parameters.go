// Code generated by go-swagger; DO NOT EDIT.

package dockerhub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDockerHubInspectParams creates a new DockerHubInspectParams object
// with the default values initialized.
func NewDockerHubInspectParams() *DockerHubInspectParams {

	return &DockerHubInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDockerHubInspectParamsWithTimeout creates a new DockerHubInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDockerHubInspectParamsWithTimeout(timeout time.Duration) *DockerHubInspectParams {

	return &DockerHubInspectParams{

		timeout: timeout,
	}
}

// NewDockerHubInspectParamsWithContext creates a new DockerHubInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewDockerHubInspectParamsWithContext(ctx context.Context) *DockerHubInspectParams {

	return &DockerHubInspectParams{

		Context: ctx,
	}
}

// NewDockerHubInspectParamsWithHTTPClient creates a new DockerHubInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDockerHubInspectParamsWithHTTPClient(client *http.Client) *DockerHubInspectParams {

	return &DockerHubInspectParams{
		HTTPClient: client,
	}
}

/*DockerHubInspectParams contains all the parameters to send to the API endpoint
for the docker hub inspect operation typically these are written to a http.Request
*/
type DockerHubInspectParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the docker hub inspect params
func (o *DockerHubInspectParams) WithTimeout(timeout time.Duration) *DockerHubInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the docker hub inspect params
func (o *DockerHubInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the docker hub inspect params
func (o *DockerHubInspectParams) WithContext(ctx context.Context) *DockerHubInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the docker hub inspect params
func (o *DockerHubInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the docker hub inspect params
func (o *DockerHubInspectParams) WithHTTPClient(client *http.Client) *DockerHubInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the docker hub inspect params
func (o *DockerHubInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DockerHubInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
