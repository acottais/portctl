// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewNetworkConnectParams creates a new NetworkConnectParams object
// with the default values initialized.
func NewNetworkConnectParams() *NetworkConnectParams {
	var ()
	return &NetworkConnectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkConnectParamsWithTimeout creates a new NetworkConnectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNetworkConnectParamsWithTimeout(timeout time.Duration) *NetworkConnectParams {
	var ()
	return &NetworkConnectParams{

		timeout: timeout,
	}
}

// NewNetworkConnectParamsWithContext creates a new NetworkConnectParams object
// with the default values initialized, and the ability to set a context for a request
func NewNetworkConnectParamsWithContext(ctx context.Context) *NetworkConnectParams {
	var ()
	return &NetworkConnectParams{

		Context: ctx,
	}
}

// NewNetworkConnectParamsWithHTTPClient creates a new NetworkConnectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNetworkConnectParamsWithHTTPClient(client *http.Client) *NetworkConnectParams {
	var ()
	return &NetworkConnectParams{
		HTTPClient: client,
	}
}

/*NetworkConnectParams contains all the parameters to send to the API endpoint
for the network connect operation typically these are written to a http.Request
*/
type NetworkConnectParams struct {

	/*Container*/
	Container NetworkConnectBody
	/*ID
	  Network ID or name

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the network connect params
func (o *NetworkConnectParams) WithTimeout(timeout time.Duration) *NetworkConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network connect params
func (o *NetworkConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network connect params
func (o *NetworkConnectParams) WithContext(ctx context.Context) *NetworkConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network connect params
func (o *NetworkConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network connect params
func (o *NetworkConnectParams) WithHTTPClient(client *http.Client) *NetworkConnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network connect params
func (o *NetworkConnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContainer adds the container to the network connect params
func (o *NetworkConnectParams) WithContainer(container NetworkConnectBody) *NetworkConnectParams {
	o.SetContainer(container)
	return o
}

// SetContainer adds the container to the network connect params
func (o *NetworkConnectParams) SetContainer(container NetworkConnectBody) {
	o.Container = container
}

// WithID adds the id to the network connect params
func (o *NetworkConnectParams) WithID(id string) *NetworkConnectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the network connect params
func (o *NetworkConnectParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Container); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
