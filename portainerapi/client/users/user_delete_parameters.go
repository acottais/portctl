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
)

// NewUserDeleteParams creates a new UserDeleteParams object
// with the default values initialized.
func NewUserDeleteParams() *UserDeleteParams {
	var ()
	return &UserDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserDeleteParamsWithTimeout creates a new UserDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserDeleteParamsWithTimeout(timeout time.Duration) *UserDeleteParams {
	var ()
	return &UserDeleteParams{

		timeout: timeout,
	}
}

// NewUserDeleteParamsWithContext creates a new UserDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserDeleteParamsWithContext(ctx context.Context) *UserDeleteParams {
	var ()
	return &UserDeleteParams{

		Context: ctx,
	}
}

// NewUserDeleteParamsWithHTTPClient creates a new UserDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserDeleteParamsWithHTTPClient(client *http.Client) *UserDeleteParams {
	var ()
	return &UserDeleteParams{
		HTTPClient: client,
	}
}

/*UserDeleteParams contains all the parameters to send to the API endpoint
for the user delete operation typically these are written to a http.Request
*/
type UserDeleteParams struct {

	/*ID
	  User identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user delete params
func (o *UserDeleteParams) WithTimeout(timeout time.Duration) *UserDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user delete params
func (o *UserDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user delete params
func (o *UserDeleteParams) WithContext(ctx context.Context) *UserDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user delete params
func (o *UserDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user delete params
func (o *UserDeleteParams) WithHTTPClient(client *http.Client) *UserDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user delete params
func (o *UserDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user delete params
func (o *UserDeleteParams) WithID(id int64) *UserDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user delete params
func (o *UserDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
