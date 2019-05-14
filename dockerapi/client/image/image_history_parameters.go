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

	strfmt "github.com/go-openapi/strfmt"
)

// NewImageHistoryParams creates a new ImageHistoryParams object
// with the default values initialized.
func NewImageHistoryParams() *ImageHistoryParams {
	var ()
	return &ImageHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageHistoryParamsWithTimeout creates a new ImageHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageHistoryParamsWithTimeout(timeout time.Duration) *ImageHistoryParams {
	var ()
	return &ImageHistoryParams{

		timeout: timeout,
	}
}

// NewImageHistoryParamsWithContext creates a new ImageHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageHistoryParamsWithContext(ctx context.Context) *ImageHistoryParams {
	var ()
	return &ImageHistoryParams{

		Context: ctx,
	}
}

// NewImageHistoryParamsWithHTTPClient creates a new ImageHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageHistoryParamsWithHTTPClient(client *http.Client) *ImageHistoryParams {
	var ()
	return &ImageHistoryParams{
		HTTPClient: client,
	}
}

/*ImageHistoryParams contains all the parameters to send to the API endpoint
for the image history operation typically these are written to a http.Request
*/
type ImageHistoryParams struct {

	/*Name
	  Image name or ID

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image history params
func (o *ImageHistoryParams) WithTimeout(timeout time.Duration) *ImageHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image history params
func (o *ImageHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image history params
func (o *ImageHistoryParams) WithContext(ctx context.Context) *ImageHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image history params
func (o *ImageHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image history params
func (o *ImageHistoryParams) WithHTTPClient(client *http.Client) *ImageHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image history params
func (o *ImageHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the image history params
func (o *ImageHistoryParams) WithName(name string) *ImageHistoryParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image history params
func (o *ImageHistoryParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ImageHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
