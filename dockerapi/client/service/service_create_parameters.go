// Code generated by go-swagger; DO NOT EDIT.

package service

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

// NewServiceCreateParams creates a new ServiceCreateParams object
// with the default values initialized.
func NewServiceCreateParams() *ServiceCreateParams {
	var ()
	return &ServiceCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceCreateParamsWithTimeout creates a new ServiceCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceCreateParamsWithTimeout(timeout time.Duration) *ServiceCreateParams {
	var ()
	return &ServiceCreateParams{

		timeout: timeout,
	}
}

// NewServiceCreateParamsWithContext creates a new ServiceCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceCreateParamsWithContext(ctx context.Context) *ServiceCreateParams {
	var ()
	return &ServiceCreateParams{

		Context: ctx,
	}
}

// NewServiceCreateParamsWithHTTPClient creates a new ServiceCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceCreateParamsWithHTTPClient(client *http.Client) *ServiceCreateParams {
	var ()
	return &ServiceCreateParams{
		HTTPClient: client,
	}
}

/*ServiceCreateParams contains all the parameters to send to the API endpoint
for the service create operation typically these are written to a http.Request
*/
type ServiceCreateParams struct {

	/*XRegistryAuth
	  A base64-encoded auth configuration for pulling from private registries. [See the authentication section for details.](#section/Authentication)

	*/
	XRegistryAuth *string
	/*Body*/
	Body ServiceCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service create params
func (o *ServiceCreateParams) WithTimeout(timeout time.Duration) *ServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service create params
func (o *ServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service create params
func (o *ServiceCreateParams) WithContext(ctx context.Context) *ServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service create params
func (o *ServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service create params
func (o *ServiceCreateParams) WithHTTPClient(client *http.Client) *ServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service create params
func (o *ServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRegistryAuth adds the xRegistryAuth to the service create params
func (o *ServiceCreateParams) WithXRegistryAuth(xRegistryAuth *string) *ServiceCreateParams {
	o.SetXRegistryAuth(xRegistryAuth)
	return o
}

// SetXRegistryAuth adds the xRegistryAuth to the service create params
func (o *ServiceCreateParams) SetXRegistryAuth(xRegistryAuth *string) {
	o.XRegistryAuth = xRegistryAuth
}

// WithBody adds the body to the service create params
func (o *ServiceCreateParams) WithBody(body ServiceCreateBody) *ServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service create params
func (o *ServiceCreateParams) SetBody(body ServiceCreateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRegistryAuth != nil {

		// header param X-Registry-Auth
		if err := r.SetHeaderParam("X-Registry-Auth", *o.XRegistryAuth); err != nil {
			return err
		}

	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
