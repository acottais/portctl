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

// NewUserMembershipsInspectParams creates a new UserMembershipsInspectParams object
// with the default values initialized.
func NewUserMembershipsInspectParams() *UserMembershipsInspectParams {
	var ()
	return &UserMembershipsInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserMembershipsInspectParamsWithTimeout creates a new UserMembershipsInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserMembershipsInspectParamsWithTimeout(timeout time.Duration) *UserMembershipsInspectParams {
	var ()
	return &UserMembershipsInspectParams{

		timeout: timeout,
	}
}

// NewUserMembershipsInspectParamsWithContext creates a new UserMembershipsInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserMembershipsInspectParamsWithContext(ctx context.Context) *UserMembershipsInspectParams {
	var ()
	return &UserMembershipsInspectParams{

		Context: ctx,
	}
}

// NewUserMembershipsInspectParamsWithHTTPClient creates a new UserMembershipsInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserMembershipsInspectParamsWithHTTPClient(client *http.Client) *UserMembershipsInspectParams {
	var ()
	return &UserMembershipsInspectParams{
		HTTPClient: client,
	}
}

/*UserMembershipsInspectParams contains all the parameters to send to the API endpoint
for the user memberships inspect operation typically these are written to a http.Request
*/
type UserMembershipsInspectParams struct {

	/*ID
	  User identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user memberships inspect params
func (o *UserMembershipsInspectParams) WithTimeout(timeout time.Duration) *UserMembershipsInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user memberships inspect params
func (o *UserMembershipsInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user memberships inspect params
func (o *UserMembershipsInspectParams) WithContext(ctx context.Context) *UserMembershipsInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user memberships inspect params
func (o *UserMembershipsInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user memberships inspect params
func (o *UserMembershipsInspectParams) WithHTTPClient(client *http.Client) *UserMembershipsInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user memberships inspect params
func (o *UserMembershipsInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user memberships inspect params
func (o *UserMembershipsInspectParams) WithID(id int64) *UserMembershipsInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user memberships inspect params
func (o *UserMembershipsInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserMembershipsInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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