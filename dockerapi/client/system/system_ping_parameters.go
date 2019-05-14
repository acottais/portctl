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

// NewSystemPingParams creates a new SystemPingParams object
// with the default values initialized.
func NewSystemPingParams() *SystemPingParams {

	return &SystemPingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemPingParamsWithTimeout creates a new SystemPingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemPingParamsWithTimeout(timeout time.Duration) *SystemPingParams {

	return &SystemPingParams{

		timeout: timeout,
	}
}

// NewSystemPingParamsWithContext creates a new SystemPingParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemPingParamsWithContext(ctx context.Context) *SystemPingParams {

	return &SystemPingParams{

		Context: ctx,
	}
}

// NewSystemPingParamsWithHTTPClient creates a new SystemPingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemPingParamsWithHTTPClient(client *http.Client) *SystemPingParams {

	return &SystemPingParams{
		HTTPClient: client,
	}
}

/*SystemPingParams contains all the parameters to send to the API endpoint
for the system ping operation typically these are written to a http.Request
*/
type SystemPingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the system ping params
func (o *SystemPingParams) WithTimeout(timeout time.Duration) *SystemPingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system ping params
func (o *SystemPingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system ping params
func (o *SystemPingParams) WithContext(ctx context.Context) *SystemPingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system ping params
func (o *SystemPingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system ping params
func (o *SystemPingParams) WithHTTPClient(client *http.Client) *SystemPingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system ping params
func (o *SystemPingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemPingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}