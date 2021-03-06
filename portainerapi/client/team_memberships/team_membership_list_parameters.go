// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

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

// NewTeamMembershipListParams creates a new TeamMembershipListParams object
// with the default values initialized.
func NewTeamMembershipListParams() *TeamMembershipListParams {

	return &TeamMembershipListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamMembershipListParamsWithTimeout creates a new TeamMembershipListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamMembershipListParamsWithTimeout(timeout time.Duration) *TeamMembershipListParams {

	return &TeamMembershipListParams{

		timeout: timeout,
	}
}

// NewTeamMembershipListParamsWithContext creates a new TeamMembershipListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamMembershipListParamsWithContext(ctx context.Context) *TeamMembershipListParams {

	return &TeamMembershipListParams{

		Context: ctx,
	}
}

// NewTeamMembershipListParamsWithHTTPClient creates a new TeamMembershipListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamMembershipListParamsWithHTTPClient(client *http.Client) *TeamMembershipListParams {

	return &TeamMembershipListParams{
		HTTPClient: client,
	}
}

/*TeamMembershipListParams contains all the parameters to send to the API endpoint
for the team membership list operation typically these are written to a http.Request
*/
type TeamMembershipListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the team membership list params
func (o *TeamMembershipListParams) WithTimeout(timeout time.Duration) *TeamMembershipListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the team membership list params
func (o *TeamMembershipListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the team membership list params
func (o *TeamMembershipListParams) WithContext(ctx context.Context) *TeamMembershipListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the team membership list params
func (o *TeamMembershipListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the team membership list params
func (o *TeamMembershipListParams) WithHTTPClient(client *http.Client) *TeamMembershipListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the team membership list params
func (o *TeamMembershipListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TeamMembershipListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
