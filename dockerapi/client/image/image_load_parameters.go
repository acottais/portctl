// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewImageLoadParams creates a new ImageLoadParams object
// with the default values initialized.
func NewImageLoadParams() *ImageLoadParams {
	var (
		quietDefault = bool(false)
	)
	return &ImageLoadParams{
		Quiet: &quietDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewImageLoadParamsWithTimeout creates a new ImageLoadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageLoadParamsWithTimeout(timeout time.Duration) *ImageLoadParams {
	var (
		quietDefault = bool(false)
	)
	return &ImageLoadParams{
		Quiet: &quietDefault,

		timeout: timeout,
	}
}

// NewImageLoadParamsWithContext creates a new ImageLoadParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageLoadParamsWithContext(ctx context.Context) *ImageLoadParams {
	var (
		quietDefault = bool(false)
	)
	return &ImageLoadParams{
		Quiet: &quietDefault,

		Context: ctx,
	}
}

// NewImageLoadParamsWithHTTPClient creates a new ImageLoadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageLoadParamsWithHTTPClient(client *http.Client) *ImageLoadParams {
	var (
		quietDefault = bool(false)
	)
	return &ImageLoadParams{
		Quiet:      &quietDefault,
		HTTPClient: client,
	}
}

/*ImageLoadParams contains all the parameters to send to the API endpoint
for the image load operation typically these are written to a http.Request
*/
type ImageLoadParams struct {

	/*ImagesTarball
	  Tar archive containing images

	*/
	ImagesTarball io.ReadCloser
	/*Quiet
	  Suppress progress details during load.

	*/
	Quiet *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image load params
func (o *ImageLoadParams) WithTimeout(timeout time.Duration) *ImageLoadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image load params
func (o *ImageLoadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image load params
func (o *ImageLoadParams) WithContext(ctx context.Context) *ImageLoadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image load params
func (o *ImageLoadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image load params
func (o *ImageLoadParams) WithHTTPClient(client *http.Client) *ImageLoadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image load params
func (o *ImageLoadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImagesTarball adds the imagesTarball to the image load params
func (o *ImageLoadParams) WithImagesTarball(imagesTarball io.ReadCloser) *ImageLoadParams {
	o.SetImagesTarball(imagesTarball)
	return o
}

// SetImagesTarball adds the imagesTarball to the image load params
func (o *ImageLoadParams) SetImagesTarball(imagesTarball io.ReadCloser) {
	o.ImagesTarball = imagesTarball
}

// WithQuiet adds the quiet to the image load params
func (o *ImageLoadParams) WithQuiet(quiet *bool) *ImageLoadParams {
	o.SetQuiet(quiet)
	return o
}

// SetQuiet adds the quiet to the image load params
func (o *ImageLoadParams) SetQuiet(quiet *bool) {
	o.Quiet = quiet
}

// WriteToRequest writes these params to a swagger request
func (o *ImageLoadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImagesTarball != nil {
		if err := r.SetBodyParam(o.ImagesTarball); err != nil {
			return err
		}
	}

	if o.Quiet != nil {

		// query param quiet
		var qrQuiet bool
		if o.Quiet != nil {
			qrQuiet = *o.Quiet
		}
		qQuiet := swag.FormatBool(qrQuiet)
		if qQuiet != "" {
			if err := r.SetQueryParam("quiet", qQuiet); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
