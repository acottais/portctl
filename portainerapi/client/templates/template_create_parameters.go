// Code generated by go-swagger; DO NOT EDIT.

package templates

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

	models "github.com/acottais/portctl/portainerapi/models"
)

// NewTemplateCreateParams creates a new TemplateCreateParams object
// with the default values initialized.
func NewTemplateCreateParams() *TemplateCreateParams {
	var ()
	return &TemplateCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTemplateCreateParamsWithTimeout creates a new TemplateCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTemplateCreateParamsWithTimeout(timeout time.Duration) *TemplateCreateParams {
	var ()
	return &TemplateCreateParams{

		timeout: timeout,
	}
}

// NewTemplateCreateParamsWithContext creates a new TemplateCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTemplateCreateParamsWithContext(ctx context.Context) *TemplateCreateParams {
	var ()
	return &TemplateCreateParams{

		Context: ctx,
	}
}

// NewTemplateCreateParamsWithHTTPClient creates a new TemplateCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTemplateCreateParamsWithHTTPClient(client *http.Client) *TemplateCreateParams {
	var ()
	return &TemplateCreateParams{
		HTTPClient: client,
	}
}

/*TemplateCreateParams contains all the parameters to send to the API endpoint
for the template create operation typically these are written to a http.Request
*/
type TemplateCreateParams struct {

	/*Body
	  Template details

	*/
	Body *models.TemplateCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the template create params
func (o *TemplateCreateParams) WithTimeout(timeout time.Duration) *TemplateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the template create params
func (o *TemplateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the template create params
func (o *TemplateCreateParams) WithContext(ctx context.Context) *TemplateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the template create params
func (o *TemplateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the template create params
func (o *TemplateCreateParams) WithHTTPClient(client *http.Client) *TemplateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the template create params
func (o *TemplateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the template create params
func (o *TemplateCreateParams) WithBody(body *models.TemplateCreateRequest) *TemplateCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the template create params
func (o *TemplateCreateParams) SetBody(body *models.TemplateCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TemplateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
