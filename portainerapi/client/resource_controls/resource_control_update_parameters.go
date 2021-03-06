// Code generated by go-swagger; DO NOT EDIT.

package resource_controls

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

	models "github.com/acottais/portctl/portainerapi/models"
)

// NewResourceControlUpdateParams creates a new ResourceControlUpdateParams object
// with the default values initialized.
func NewResourceControlUpdateParams() *ResourceControlUpdateParams {
	var ()
	return &ResourceControlUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResourceControlUpdateParamsWithTimeout creates a new ResourceControlUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResourceControlUpdateParamsWithTimeout(timeout time.Duration) *ResourceControlUpdateParams {
	var ()
	return &ResourceControlUpdateParams{

		timeout: timeout,
	}
}

// NewResourceControlUpdateParamsWithContext creates a new ResourceControlUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewResourceControlUpdateParamsWithContext(ctx context.Context) *ResourceControlUpdateParams {
	var ()
	return &ResourceControlUpdateParams{

		Context: ctx,
	}
}

// NewResourceControlUpdateParamsWithHTTPClient creates a new ResourceControlUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResourceControlUpdateParamsWithHTTPClient(client *http.Client) *ResourceControlUpdateParams {
	var ()
	return &ResourceControlUpdateParams{
		HTTPClient: client,
	}
}

/*ResourceControlUpdateParams contains all the parameters to send to the API endpoint
for the resource control update operation typically these are written to a http.Request
*/
type ResourceControlUpdateParams struct {

	/*Body
	  Resource control details

	*/
	Body *models.ResourceControlUpdateRequest
	/*ID
	  Resource control identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resource control update params
func (o *ResourceControlUpdateParams) WithTimeout(timeout time.Duration) *ResourceControlUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource control update params
func (o *ResourceControlUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource control update params
func (o *ResourceControlUpdateParams) WithContext(ctx context.Context) *ResourceControlUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource control update params
func (o *ResourceControlUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource control update params
func (o *ResourceControlUpdateParams) WithHTTPClient(client *http.Client) *ResourceControlUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource control update params
func (o *ResourceControlUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the resource control update params
func (o *ResourceControlUpdateParams) WithBody(body *models.ResourceControlUpdateRequest) *ResourceControlUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the resource control update params
func (o *ResourceControlUpdateParams) SetBody(body *models.ResourceControlUpdateRequest) {
	o.Body = body
}

// WithID adds the id to the resource control update params
func (o *ResourceControlUpdateParams) WithID(id int64) *ResourceControlUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resource control update params
func (o *ResourceControlUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceControlUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
