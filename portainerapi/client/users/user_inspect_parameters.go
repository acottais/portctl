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

// NewUserInspectParams creates a new UserInspectParams object
// with the default values initialized.
func NewUserInspectParams() *UserInspectParams {
	var ()
	return &UserInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserInspectParamsWithTimeout creates a new UserInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserInspectParamsWithTimeout(timeout time.Duration) *UserInspectParams {
	var ()
	return &UserInspectParams{

		timeout: timeout,
	}
}

// NewUserInspectParamsWithContext creates a new UserInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserInspectParamsWithContext(ctx context.Context) *UserInspectParams {
	var ()
	return &UserInspectParams{

		Context: ctx,
	}
}

// NewUserInspectParamsWithHTTPClient creates a new UserInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserInspectParamsWithHTTPClient(client *http.Client) *UserInspectParams {
	var ()
	return &UserInspectParams{
		HTTPClient: client,
	}
}

/*UserInspectParams contains all the parameters to send to the API endpoint
for the user inspect operation typically these are written to a http.Request
*/
type UserInspectParams struct {

	/*ID
	  User identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user inspect params
func (o *UserInspectParams) WithTimeout(timeout time.Duration) *UserInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user inspect params
func (o *UserInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user inspect params
func (o *UserInspectParams) WithContext(ctx context.Context) *UserInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user inspect params
func (o *UserInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user inspect params
func (o *UserInspectParams) WithHTTPClient(client *http.Client) *UserInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user inspect params
func (o *UserInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user inspect params
func (o *UserInspectParams) WithID(id int64) *UserInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user inspect params
func (o *UserInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
