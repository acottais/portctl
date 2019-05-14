// Code generated by go-swagger; DO NOT EDIT.

package task

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

// NewTaskListParams creates a new TaskListParams object
// with the default values initialized.
func NewTaskListParams() *TaskListParams {
	var ()
	return &TaskListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTaskListParamsWithTimeout creates a new TaskListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaskListParamsWithTimeout(timeout time.Duration) *TaskListParams {
	var ()
	return &TaskListParams{

		timeout: timeout,
	}
}

// NewTaskListParamsWithContext creates a new TaskListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTaskListParamsWithContext(ctx context.Context) *TaskListParams {
	var ()
	return &TaskListParams{

		Context: ctx,
	}
}

// NewTaskListParamsWithHTTPClient creates a new TaskListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaskListParamsWithHTTPClient(client *http.Client) *TaskListParams {
	var ()
	return &TaskListParams{
		HTTPClient: client,
	}
}

/*TaskListParams contains all the parameters to send to the API endpoint
for the task list operation typically these are written to a http.Request
*/
type TaskListParams struct {

	/*Filters
	  A JSON encoded value of the filters (a `map[string][]string`) to process on the tasks list. Available filters:

	- `desired-state=(running | shutdown | accepted)`
	- `id=<task id>`
	- `label=key` or `label="key=value"`
	- `name=<task name>`
	- `node=<node id or name>`
	- `service=<service name>`


	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the task list params
func (o *TaskListParams) WithTimeout(timeout time.Duration) *TaskListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the task list params
func (o *TaskListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the task list params
func (o *TaskListParams) WithContext(ctx context.Context) *TaskListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the task list params
func (o *TaskListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the task list params
func (o *TaskListParams) WithHTTPClient(client *http.Client) *TaskListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the task list params
func (o *TaskListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the task list params
func (o *TaskListParams) WithFilters(filters *string) *TaskListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the task list params
func (o *TaskListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *TaskListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
