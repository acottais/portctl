// Code generated by go-swagger; DO NOT EDIT.

package image

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

// NewBuildPruneParams creates a new BuildPruneParams object
// with the default values initialized.
func NewBuildPruneParams() *BuildPruneParams {
	var ()
	return &BuildPruneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildPruneParamsWithTimeout creates a new BuildPruneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildPruneParamsWithTimeout(timeout time.Duration) *BuildPruneParams {
	var ()
	return &BuildPruneParams{

		timeout: timeout,
	}
}

// NewBuildPruneParamsWithContext creates a new BuildPruneParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildPruneParamsWithContext(ctx context.Context) *BuildPruneParams {
	var ()
	return &BuildPruneParams{

		Context: ctx,
	}
}

// NewBuildPruneParamsWithHTTPClient creates a new BuildPruneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildPruneParamsWithHTTPClient(client *http.Client) *BuildPruneParams {
	var ()
	return &BuildPruneParams{
		HTTPClient: client,
	}
}

/*BuildPruneParams contains all the parameters to send to the API endpoint
for the build prune operation typically these are written to a http.Request
*/
type BuildPruneParams struct {

	/*All
	  Remove all types of build cache

	*/
	All *bool
	/*Filters
	  A JSON encoded value of the filters (a `map[string][]string`) to process on the list of build cache objects. Available filters:
	- `until=<duration>`: duration relative to daemon's time, during which build cache was not used, in Go's duration format (e.g., '24h')
	- `id=<id>`
	- `parent=<id>`
	- `type=<string>`
	- `description=<string>`
	- `inuse`
	- `shared`
	- `private`


	*/
	Filters *string
	/*KeepStorage
	  Amount of disk space in bytes to keep for cache

	*/
	KeepStorage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the build prune params
func (o *BuildPruneParams) WithTimeout(timeout time.Duration) *BuildPruneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build prune params
func (o *BuildPruneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build prune params
func (o *BuildPruneParams) WithContext(ctx context.Context) *BuildPruneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build prune params
func (o *BuildPruneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build prune params
func (o *BuildPruneParams) WithHTTPClient(client *http.Client) *BuildPruneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build prune params
func (o *BuildPruneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the build prune params
func (o *BuildPruneParams) WithAll(all *bool) *BuildPruneParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the build prune params
func (o *BuildPruneParams) SetAll(all *bool) {
	o.All = all
}

// WithFilters adds the filters to the build prune params
func (o *BuildPruneParams) WithFilters(filters *string) *BuildPruneParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the build prune params
func (o *BuildPruneParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithKeepStorage adds the keepStorage to the build prune params
func (o *BuildPruneParams) WithKeepStorage(keepStorage *int64) *BuildPruneParams {
	o.SetKeepStorage(keepStorage)
	return o
}

// SetKeepStorage adds the keepStorage to the build prune params
func (o *BuildPruneParams) SetKeepStorage(keepStorage *int64) {
	o.KeepStorage = keepStorage
}

// WriteToRequest writes these params to a swagger request
func (o *BuildPruneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool
		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {
			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}

	}

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

	if o.KeepStorage != nil {

		// query param keep-storage
		var qrKeepStorage int64
		if o.KeepStorage != nil {
			qrKeepStorage = *o.KeepStorage
		}
		qKeepStorage := swag.FormatInt64(qrKeepStorage)
		if qKeepStorage != "" {
			if err := r.SetQueryParam("keep-storage", qKeepStorage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
