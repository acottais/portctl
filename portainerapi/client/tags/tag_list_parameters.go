// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewTagListParams creates a new TagListParams object
// with the default values initialized.
func NewTagListParams() *TagListParams {

	return &TagListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTagListParamsWithTimeout creates a new TagListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTagListParamsWithTimeout(timeout time.Duration) *TagListParams {

	return &TagListParams{

		timeout: timeout,
	}
}

// NewTagListParamsWithContext creates a new TagListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTagListParamsWithContext(ctx context.Context) *TagListParams {

	return &TagListParams{

		Context: ctx,
	}
}

// NewTagListParamsWithHTTPClient creates a new TagListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTagListParamsWithHTTPClient(client *http.Client) *TagListParams {

	return &TagListParams{
		HTTPClient: client,
	}
}

/*TagListParams contains all the parameters to send to the API endpoint
for the tag list operation typically these are written to a http.Request
*/
type TagListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tag list params
func (o *TagListParams) WithTimeout(timeout time.Duration) *TagListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tag list params
func (o *TagListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tag list params
func (o *TagListParams) WithContext(ctx context.Context) *TagListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tag list params
func (o *TagListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tag list params
func (o *TagListParams) WithHTTPClient(client *http.Client) *TagListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tag list params
func (o *TagListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TagListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}