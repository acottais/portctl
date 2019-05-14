// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUserUpdateParams creates a new UserUpdateParams object
// with the default values initialized.
func NewUserUpdateParams() *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserUpdateParamsWithTimeout creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserUpdateParamsWithTimeout(timeout time.Duration) *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		timeout: timeout,
	}
}

// NewUserUpdateParamsWithContext creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserUpdateParamsWithContext(ctx context.Context) *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		Context: ctx,
	}
}

// NewUserUpdateParamsWithHTTPClient creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserUpdateParamsWithHTTPClient(client *http.Client) *UserUpdateParams {
	var ()
	return &UserUpdateParams{
		HTTPClient: client,
	}
}

/*UserUpdateParams contains all the parameters to send to the API endpoint
for the user update operation typically these are written to a http.Request
*/
type UserUpdateParams struct {

	/*Body
	  User details

	*/
	Body *models.UserUpdateRequest
	/*ID
	  User identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user update params
func (o *UserUpdateParams) WithTimeout(timeout time.Duration) *UserUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user update params
func (o *UserUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user update params
func (o *UserUpdateParams) WithContext(ctx context.Context) *UserUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user update params
func (o *UserUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user update params
func (o *UserUpdateParams) WithHTTPClient(client *http.Client) *UserUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user update params
func (o *UserUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user update params
func (o *UserUpdateParams) WithBody(body *models.UserUpdateRequest) *UserUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user update params
func (o *UserUpdateParams) SetBody(body *models.UserUpdateRequest) {
	o.Body = body
}

// WithID adds the id to the user update params
func (o *UserUpdateParams) WithID(id int64) *UserUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user update params
func (o *UserUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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