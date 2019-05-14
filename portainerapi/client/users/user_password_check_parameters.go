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

// NewUserPasswordCheckParams creates a new UserPasswordCheckParams object
// with the default values initialized.
func NewUserPasswordCheckParams() *UserPasswordCheckParams {
	var ()
	return &UserPasswordCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserPasswordCheckParamsWithTimeout creates a new UserPasswordCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserPasswordCheckParamsWithTimeout(timeout time.Duration) *UserPasswordCheckParams {
	var ()
	return &UserPasswordCheckParams{

		timeout: timeout,
	}
}

// NewUserPasswordCheckParamsWithContext creates a new UserPasswordCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserPasswordCheckParamsWithContext(ctx context.Context) *UserPasswordCheckParams {
	var ()
	return &UserPasswordCheckParams{

		Context: ctx,
	}
}

// NewUserPasswordCheckParamsWithHTTPClient creates a new UserPasswordCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserPasswordCheckParamsWithHTTPClient(client *http.Client) *UserPasswordCheckParams {
	var ()
	return &UserPasswordCheckParams{
		HTTPClient: client,
	}
}

/*UserPasswordCheckParams contains all the parameters to send to the API endpoint
for the user password check operation typically these are written to a http.Request
*/
type UserPasswordCheckParams struct {

	/*Body
	  User details

	*/
	Body *models.UserPasswordCheckRequest
	/*ID
	  User identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user password check params
func (o *UserPasswordCheckParams) WithTimeout(timeout time.Duration) *UserPasswordCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user password check params
func (o *UserPasswordCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user password check params
func (o *UserPasswordCheckParams) WithContext(ctx context.Context) *UserPasswordCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user password check params
func (o *UserPasswordCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user password check params
func (o *UserPasswordCheckParams) WithHTTPClient(client *http.Client) *UserPasswordCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user password check params
func (o *UserPasswordCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user password check params
func (o *UserPasswordCheckParams) WithBody(body *models.UserPasswordCheckRequest) *UserPasswordCheckParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user password check params
func (o *UserPasswordCheckParams) SetBody(body *models.UserPasswordCheckRequest) {
	o.Body = body
}

// WithID adds the id to the user password check params
func (o *UserPasswordCheckParams) WithID(id int64) *UserPasswordCheckParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user password check params
func (o *UserPasswordCheckParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserPasswordCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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