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

	strfmt "github.com/go-openapi/strfmt"
)

// NewNodeInspectParams creates a new NodeInspectParams object
// with the default values initialized.
func NewNodeInspectParams() *NodeInspectParams {
	var ()
	return &NodeInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodeInspectParamsWithTimeout creates a new NodeInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodeInspectParamsWithTimeout(timeout time.Duration) *NodeInspectParams {
	var ()
	return &NodeInspectParams{

		timeout: timeout,
	}
}

// NewNodeInspectParamsWithContext creates a new NodeInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodeInspectParamsWithContext(ctx context.Context) *NodeInspectParams {
	var ()
	return &NodeInspectParams{

		Context: ctx,
	}
}

// NewNodeInspectParamsWithHTTPClient creates a new NodeInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodeInspectParamsWithHTTPClient(client *http.Client) *NodeInspectParams {
	var ()
	return &NodeInspectParams{
		HTTPClient: client,
	}
}

/*NodeInspectParams contains all the parameters to send to the API endpoint
for the node inspect operation typically these are written to a http.Request
*/
type NodeInspectParams struct {

	/*ID
	  The ID or name of the node

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the node inspect params
func (o *NodeInspectParams) WithTimeout(timeout time.Duration) *NodeInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node inspect params
func (o *NodeInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node inspect params
func (o *NodeInspectParams) WithContext(ctx context.Context) *NodeInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node inspect params
func (o *NodeInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node inspect params
func (o *NodeInspectParams) WithHTTPClient(client *http.Client) *NodeInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node inspect params
func (o *NodeInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the node inspect params
func (o *NodeInspectParams) WithID(id string) *NodeInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the node inspect params
func (o *NodeInspectParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NodeInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
