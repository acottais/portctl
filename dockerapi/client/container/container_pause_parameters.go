// Code generated by go-swagger; DO NOT EDIT.

package container

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

// NewContainerPauseParams creates a new ContainerPauseParams object
// with the default values initialized.
func NewContainerPauseParams() *ContainerPauseParams {
	var ()
	return &ContainerPauseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerPauseParamsWithTimeout creates a new ContainerPauseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerPauseParamsWithTimeout(timeout time.Duration) *ContainerPauseParams {
	var ()
	return &ContainerPauseParams{

		timeout: timeout,
	}
}

// NewContainerPauseParamsWithContext creates a new ContainerPauseParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerPauseParamsWithContext(ctx context.Context) *ContainerPauseParams {
	var ()
	return &ContainerPauseParams{

		Context: ctx,
	}
}

// NewContainerPauseParamsWithHTTPClient creates a new ContainerPauseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerPauseParamsWithHTTPClient(client *http.Client) *ContainerPauseParams {
	var ()
	return &ContainerPauseParams{
		HTTPClient: client,
	}
}

/*ContainerPauseParams contains all the parameters to send to the API endpoint
for the container pause operation typically these are written to a http.Request
*/
type ContainerPauseParams struct {

	/*ID
	  ID or name of the container

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container pause params
func (o *ContainerPauseParams) WithTimeout(timeout time.Duration) *ContainerPauseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container pause params
func (o *ContainerPauseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container pause params
func (o *ContainerPauseParams) WithContext(ctx context.Context) *ContainerPauseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container pause params
func (o *ContainerPauseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container pause params
func (o *ContainerPauseParams) WithHTTPClient(client *http.Client) *ContainerPauseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container pause params
func (o *ContainerPauseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the container pause params
func (o *ContainerPauseParams) WithID(id string) *ContainerPauseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container pause params
func (o *ContainerPauseParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerPauseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
