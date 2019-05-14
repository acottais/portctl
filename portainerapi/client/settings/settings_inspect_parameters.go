// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewSettingsInspectParams creates a new SettingsInspectParams object
// with the default values initialized.
func NewSettingsInspectParams() *SettingsInspectParams {

	return &SettingsInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsInspectParamsWithTimeout creates a new SettingsInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSettingsInspectParamsWithTimeout(timeout time.Duration) *SettingsInspectParams {

	return &SettingsInspectParams{

		timeout: timeout,
	}
}

// NewSettingsInspectParamsWithContext creates a new SettingsInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewSettingsInspectParamsWithContext(ctx context.Context) *SettingsInspectParams {

	return &SettingsInspectParams{

		Context: ctx,
	}
}

// NewSettingsInspectParamsWithHTTPClient creates a new SettingsInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSettingsInspectParamsWithHTTPClient(client *http.Client) *SettingsInspectParams {

	return &SettingsInspectParams{
		HTTPClient: client,
	}
}

/*SettingsInspectParams contains all the parameters to send to the API endpoint
for the settings inspect operation typically these are written to a http.Request
*/
type SettingsInspectParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the settings inspect params
func (o *SettingsInspectParams) WithTimeout(timeout time.Duration) *SettingsInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings inspect params
func (o *SettingsInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings inspect params
func (o *SettingsInspectParams) WithContext(ctx context.Context) *SettingsInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings inspect params
func (o *SettingsInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings inspect params
func (o *SettingsInspectParams) WithHTTPClient(client *http.Client) *SettingsInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings inspect params
func (o *SettingsInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}