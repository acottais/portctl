// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewConfigListParams creates a new ConfigListParams object
// with the default values initialized.
func NewConfigListParams() *ConfigListParams {
	var ()
	return &ConfigListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfigListParamsWithTimeout creates a new ConfigListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfigListParamsWithTimeout(timeout time.Duration) *ConfigListParams {
	var ()
	return &ConfigListParams{

		timeout: timeout,
	}
}

// NewConfigListParamsWithContext creates a new ConfigListParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfigListParamsWithContext(ctx context.Context) *ConfigListParams {
	var ()
	return &ConfigListParams{

		Context: ctx,
	}
}

// NewConfigListParamsWithHTTPClient creates a new ConfigListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfigListParamsWithHTTPClient(client *http.Client) *ConfigListParams {
	var ()
	return &ConfigListParams{
		HTTPClient: client,
	}
}

/*ConfigListParams contains all the parameters to send to the API endpoint
for the config list operation typically these are written to a http.Request
*/
type ConfigListParams struct {

	/*Filters
	  A JSON encoded value of the filters (a `map[string][]string`) to process on the configs list. Available filters:

	- `id=<config id>`
	- `label=<key> or label=<key>=value`
	- `name=<config name>`
	- `names=<config name>`


	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the config list params
func (o *ConfigListParams) WithTimeout(timeout time.Duration) *ConfigListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config list params
func (o *ConfigListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config list params
func (o *ConfigListParams) WithContext(ctx context.Context) *ConfigListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config list params
func (o *ConfigListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config list params
func (o *ConfigListParams) WithHTTPClient(client *http.Client) *ConfigListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config list params
func (o *ConfigListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the config list params
func (o *ConfigListParams) WithFilters(filters *string) *ConfigListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the config list params
func (o *ConfigListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}