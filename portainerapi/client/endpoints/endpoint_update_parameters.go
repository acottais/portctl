// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

// NewEndpointUpdateParams creates a new EndpointUpdateParams object
// with the default values initialized.
func NewEndpointUpdateParams() *EndpointUpdateParams {
	var ()
	return &EndpointUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointUpdateParamsWithTimeout creates a new EndpointUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEndpointUpdateParamsWithTimeout(timeout time.Duration) *EndpointUpdateParams {
	var ()
	return &EndpointUpdateParams{

		timeout: timeout,
	}
}

// NewEndpointUpdateParamsWithContext creates a new EndpointUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewEndpointUpdateParamsWithContext(ctx context.Context) *EndpointUpdateParams {
	var ()
	return &EndpointUpdateParams{

		Context: ctx,
	}
}

// NewEndpointUpdateParamsWithHTTPClient creates a new EndpointUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEndpointUpdateParamsWithHTTPClient(client *http.Client) *EndpointUpdateParams {
	var ()
	return &EndpointUpdateParams{
		HTTPClient: client,
	}
}

/*EndpointUpdateParams contains all the parameters to send to the API endpoint
for the endpoint update operation typically these are written to a http.Request
*/
type EndpointUpdateParams struct {

	/*Body
	  Endpoint details

	*/
	Body *models.EndpointUpdateRequest
	/*ID
	  Endpoint identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the endpoint update params
func (o *EndpointUpdateParams) WithTimeout(timeout time.Duration) *EndpointUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint update params
func (o *EndpointUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint update params
func (o *EndpointUpdateParams) WithContext(ctx context.Context) *EndpointUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint update params
func (o *EndpointUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint update params
func (o *EndpointUpdateParams) WithHTTPClient(client *http.Client) *EndpointUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint update params
func (o *EndpointUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the endpoint update params
func (o *EndpointUpdateParams) WithBody(body *models.EndpointUpdateRequest) *EndpointUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the endpoint update params
func (o *EndpointUpdateParams) SetBody(body *models.EndpointUpdateRequest) {
	o.Body = body
}

// WithID adds the id to the endpoint update params
func (o *EndpointUpdateParams) WithID(id int64) *EndpointUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint update params
func (o *EndpointUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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