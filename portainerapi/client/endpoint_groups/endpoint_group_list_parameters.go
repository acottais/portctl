// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

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

// NewEndpointGroupListParams creates a new EndpointGroupListParams object
// with the default values initialized.
func NewEndpointGroupListParams() *EndpointGroupListParams {

	return &EndpointGroupListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointGroupListParamsWithTimeout creates a new EndpointGroupListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEndpointGroupListParamsWithTimeout(timeout time.Duration) *EndpointGroupListParams {

	return &EndpointGroupListParams{

		timeout: timeout,
	}
}

// NewEndpointGroupListParamsWithContext creates a new EndpointGroupListParams object
// with the default values initialized, and the ability to set a context for a request
func NewEndpointGroupListParamsWithContext(ctx context.Context) *EndpointGroupListParams {

	return &EndpointGroupListParams{

		Context: ctx,
	}
}

// NewEndpointGroupListParamsWithHTTPClient creates a new EndpointGroupListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEndpointGroupListParamsWithHTTPClient(client *http.Client) *EndpointGroupListParams {

	return &EndpointGroupListParams{
		HTTPClient: client,
	}
}

/*EndpointGroupListParams contains all the parameters to send to the API endpoint
for the endpoint group list operation typically these are written to a http.Request
*/
type EndpointGroupListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the endpoint group list params
func (o *EndpointGroupListParams) WithTimeout(timeout time.Duration) *EndpointGroupListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint group list params
func (o *EndpointGroupListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint group list params
func (o *EndpointGroupListParams) WithContext(ctx context.Context) *EndpointGroupListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint group list params
func (o *EndpointGroupListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint group list params
func (o *EndpointGroupListParams) WithHTTPClient(client *http.Client) *EndpointGroupListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint group list params
func (o *EndpointGroupListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointGroupListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
