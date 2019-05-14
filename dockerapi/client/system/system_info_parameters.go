// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewSystemInfoParams creates a new SystemInfoParams object
// with the default values initialized.
func NewSystemInfoParams() *SystemInfoParams {

	return &SystemInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemInfoParamsWithTimeout creates a new SystemInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemInfoParamsWithTimeout(timeout time.Duration) *SystemInfoParams {

	return &SystemInfoParams{

		timeout: timeout,
	}
}

// NewSystemInfoParamsWithContext creates a new SystemInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemInfoParamsWithContext(ctx context.Context) *SystemInfoParams {

	return &SystemInfoParams{

		Context: ctx,
	}
}

// NewSystemInfoParamsWithHTTPClient creates a new SystemInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemInfoParamsWithHTTPClient(client *http.Client) *SystemInfoParams {

	return &SystemInfoParams{
		HTTPClient: client,
	}
}

/*SystemInfoParams contains all the parameters to send to the API endpoint
for the system info operation typically these are written to a http.Request
*/
type SystemInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the system info params
func (o *SystemInfoParams) WithTimeout(timeout time.Duration) *SystemInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system info params
func (o *SystemInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system info params
func (o *SystemInfoParams) WithContext(ctx context.Context) *SystemInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system info params
func (o *SystemInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system info params
func (o *SystemInfoParams) WithHTTPClient(client *http.Client) *SystemInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system info params
func (o *SystemInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
