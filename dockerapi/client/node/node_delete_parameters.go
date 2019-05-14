// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNodeDeleteParams creates a new NodeDeleteParams object
// with the default values initialized.
func NewNodeDeleteParams() *NodeDeleteParams {
	var (
		forceDefault = bool(false)
	)
	return &NodeDeleteParams{
		Force: &forceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewNodeDeleteParamsWithTimeout creates a new NodeDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodeDeleteParamsWithTimeout(timeout time.Duration) *NodeDeleteParams {
	var (
		forceDefault = bool(false)
	)
	return &NodeDeleteParams{
		Force: &forceDefault,

		timeout: timeout,
	}
}

// NewNodeDeleteParamsWithContext creates a new NodeDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodeDeleteParamsWithContext(ctx context.Context) *NodeDeleteParams {
	var (
		forceDefault = bool(false)
	)
	return &NodeDeleteParams{
		Force: &forceDefault,

		Context: ctx,
	}
}

// NewNodeDeleteParamsWithHTTPClient creates a new NodeDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodeDeleteParamsWithHTTPClient(client *http.Client) *NodeDeleteParams {
	var (
		forceDefault = bool(false)
	)
	return &NodeDeleteParams{
		Force:      &forceDefault,
		HTTPClient: client,
	}
}

/*NodeDeleteParams contains all the parameters to send to the API endpoint
for the node delete operation typically these are written to a http.Request
*/
type NodeDeleteParams struct {

	/*Force
	  Force remove a node from the swarm

	*/
	Force *bool
	/*ID
	  The ID or name of the node

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the node delete params
func (o *NodeDeleteParams) WithTimeout(timeout time.Duration) *NodeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node delete params
func (o *NodeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node delete params
func (o *NodeDeleteParams) WithContext(ctx context.Context) *NodeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node delete params
func (o *NodeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node delete params
func (o *NodeDeleteParams) WithHTTPClient(client *http.Client) *NodeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node delete params
func (o *NodeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the node delete params
func (o *NodeDeleteParams) WithForce(force *bool) *NodeDeleteParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the node delete params
func (o *NodeDeleteParams) SetForce(force *bool) {
	o.Force = force
}

// WithID adds the id to the node delete params
func (o *NodeDeleteParams) WithID(id string) *NodeDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the node delete params
func (o *NodeDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NodeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

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
