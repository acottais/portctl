// Code generated by go-swagger; DO NOT EDIT.

package registries

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

// NewRegistryInspectParams creates a new RegistryInspectParams object
// with the default values initialized.
func NewRegistryInspectParams() *RegistryInspectParams {
	var ()
	return &RegistryInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegistryInspectParamsWithTimeout creates a new RegistryInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegistryInspectParamsWithTimeout(timeout time.Duration) *RegistryInspectParams {
	var ()
	return &RegistryInspectParams{

		timeout: timeout,
	}
}

// NewRegistryInspectParamsWithContext creates a new RegistryInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegistryInspectParamsWithContext(ctx context.Context) *RegistryInspectParams {
	var ()
	return &RegistryInspectParams{

		Context: ctx,
	}
}

// NewRegistryInspectParamsWithHTTPClient creates a new RegistryInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegistryInspectParamsWithHTTPClient(client *http.Client) *RegistryInspectParams {
	var ()
	return &RegistryInspectParams{
		HTTPClient: client,
	}
}

/*RegistryInspectParams contains all the parameters to send to the API endpoint
for the registry inspect operation typically these are written to a http.Request
*/
type RegistryInspectParams struct {

	/*ID
	  Registry identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the registry inspect params
func (o *RegistryInspectParams) WithTimeout(timeout time.Duration) *RegistryInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the registry inspect params
func (o *RegistryInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the registry inspect params
func (o *RegistryInspectParams) WithContext(ctx context.Context) *RegistryInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the registry inspect params
func (o *RegistryInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the registry inspect params
func (o *RegistryInspectParams) WithHTTPClient(client *http.Client) *RegistryInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the registry inspect params
func (o *RegistryInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the registry inspect params
func (o *RegistryInspectParams) WithID(id int64) *RegistryInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the registry inspect params
func (o *RegistryInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RegistryInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}