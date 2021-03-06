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

// NewPublicSettingsInspectParams creates a new PublicSettingsInspectParams object
// with the default values initialized.
func NewPublicSettingsInspectParams() *PublicSettingsInspectParams {

	return &PublicSettingsInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicSettingsInspectParamsWithTimeout creates a new PublicSettingsInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicSettingsInspectParamsWithTimeout(timeout time.Duration) *PublicSettingsInspectParams {

	return &PublicSettingsInspectParams{

		timeout: timeout,
	}
}

// NewPublicSettingsInspectParamsWithContext creates a new PublicSettingsInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicSettingsInspectParamsWithContext(ctx context.Context) *PublicSettingsInspectParams {

	return &PublicSettingsInspectParams{

		Context: ctx,
	}
}

// NewPublicSettingsInspectParamsWithHTTPClient creates a new PublicSettingsInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicSettingsInspectParamsWithHTTPClient(client *http.Client) *PublicSettingsInspectParams {

	return &PublicSettingsInspectParams{
		HTTPClient: client,
	}
}

/*PublicSettingsInspectParams contains all the parameters to send to the API endpoint
for the public settings inspect operation typically these are written to a http.Request
*/
type PublicSettingsInspectParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public settings inspect params
func (o *PublicSettingsInspectParams) WithTimeout(timeout time.Duration) *PublicSettingsInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public settings inspect params
func (o *PublicSettingsInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public settings inspect params
func (o *PublicSettingsInspectParams) WithContext(ctx context.Context) *PublicSettingsInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public settings inspect params
func (o *PublicSettingsInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public settings inspect params
func (o *PublicSettingsInspectParams) WithHTTPClient(client *http.Client) *PublicSettingsInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public settings inspect params
func (o *PublicSettingsInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PublicSettingsInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
