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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewContainerRestartParams creates a new ContainerRestartParams object
// with the default values initialized.
func NewContainerRestartParams() *ContainerRestartParams {
	var ()
	return &ContainerRestartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerRestartParamsWithTimeout creates a new ContainerRestartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerRestartParamsWithTimeout(timeout time.Duration) *ContainerRestartParams {
	var ()
	return &ContainerRestartParams{

		timeout: timeout,
	}
}

// NewContainerRestartParamsWithContext creates a new ContainerRestartParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerRestartParamsWithContext(ctx context.Context) *ContainerRestartParams {
	var ()
	return &ContainerRestartParams{

		Context: ctx,
	}
}

// NewContainerRestartParamsWithHTTPClient creates a new ContainerRestartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerRestartParamsWithHTTPClient(client *http.Client) *ContainerRestartParams {
	var ()
	return &ContainerRestartParams{
		HTTPClient: client,
	}
}

/*ContainerRestartParams contains all the parameters to send to the API endpoint
for the container restart operation typically these are written to a http.Request
*/
type ContainerRestartParams struct {

	/*ID
	  ID or name of the container

	*/
	ID string
	/*T
	  Number of seconds to wait before killing the container

	*/
	T *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container restart params
func (o *ContainerRestartParams) WithTimeout(timeout time.Duration) *ContainerRestartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container restart params
func (o *ContainerRestartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container restart params
func (o *ContainerRestartParams) WithContext(ctx context.Context) *ContainerRestartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container restart params
func (o *ContainerRestartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container restart params
func (o *ContainerRestartParams) WithHTTPClient(client *http.Client) *ContainerRestartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container restart params
func (o *ContainerRestartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the container restart params
func (o *ContainerRestartParams) WithID(id string) *ContainerRestartParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container restart params
func (o *ContainerRestartParams) SetID(id string) {
	o.ID = id
}

// WithT adds the t to the container restart params
func (o *ContainerRestartParams) WithT(t *int64) *ContainerRestartParams {
	o.SetT(t)
	return o
}

// SetT adds the t to the container restart params
func (o *ContainerRestartParams) SetT(t *int64) {
	o.T = t
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerRestartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.T != nil {

		// query param t
		var qrT int64
		if o.T != nil {
			qrT = *o.T
		}
		qT := swag.FormatInt64(qrT)
		if qT != "" {
			if err := r.SetQueryParam("t", qT); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
