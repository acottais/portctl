// Code generated by go-swagger; DO NOT EDIT.

package container

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

// NewContainerAttachParams creates a new ContainerAttachParams object
// with the default values initialized.
func NewContainerAttachParams() *ContainerAttachParams {
	var (
		logsDefault   = bool(false)
		stderrDefault = bool(false)
		stdinDefault  = bool(false)
		stdoutDefault = bool(false)
		streamDefault = bool(false)
	)
	return &ContainerAttachParams{
		Logs:   &logsDefault,
		Stderr: &stderrDefault,
		Stdin:  &stdinDefault,
		Stdout: &stdoutDefault,
		Stream: &streamDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerAttachParamsWithTimeout creates a new ContainerAttachParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerAttachParamsWithTimeout(timeout time.Duration) *ContainerAttachParams {
	var (
		logsDefault   = bool(false)
		stderrDefault = bool(false)
		stdinDefault  = bool(false)
		stdoutDefault = bool(false)
		streamDefault = bool(false)
	)
	return &ContainerAttachParams{
		Logs:   &logsDefault,
		Stderr: &stderrDefault,
		Stdin:  &stdinDefault,
		Stdout: &stdoutDefault,
		Stream: &streamDefault,

		timeout: timeout,
	}
}

// NewContainerAttachParamsWithContext creates a new ContainerAttachParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerAttachParamsWithContext(ctx context.Context) *ContainerAttachParams {
	var (
		logsDefault   = bool(false)
		stderrDefault = bool(false)
		stdinDefault  = bool(false)
		stdoutDefault = bool(false)
		streamDefault = bool(false)
	)
	return &ContainerAttachParams{
		Logs:   &logsDefault,
		Stderr: &stderrDefault,
		Stdin:  &stdinDefault,
		Stdout: &stdoutDefault,
		Stream: &streamDefault,

		Context: ctx,
	}
}

// NewContainerAttachParamsWithHTTPClient creates a new ContainerAttachParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerAttachParamsWithHTTPClient(client *http.Client) *ContainerAttachParams {
	var (
		logsDefault   = bool(false)
		stderrDefault = bool(false)
		stdinDefault  = bool(false)
		stdoutDefault = bool(false)
		streamDefault = bool(false)
	)
	return &ContainerAttachParams{
		Logs:       &logsDefault,
		Stderr:     &stderrDefault,
		Stdin:      &stdinDefault,
		Stdout:     &stdoutDefault,
		Stream:     &streamDefault,
		HTTPClient: client,
	}
}

/*ContainerAttachParams contains all the parameters to send to the API endpoint
for the container attach operation typically these are written to a http.Request
*/
type ContainerAttachParams struct {

	/*DetachKeys
	  Override the key sequence for detaching a container.Format is a single character `[a-Z]` or `ctrl-<value>` where `<value>` is one of: `a-z`, `@`, `^`, `[`, `,` or `_`.

	*/
	DetachKeys *string
	/*ID
	  ID or name of the container

	*/
	ID string
	/*Logs
	  Replay previous logs from the container.

	This is useful for attaching to a container that has started and you want to output everything since the container started.

	If `stream` is also enabled, once all the previous output has been returned, it will seamlessly transition into streaming current output.


	*/
	Logs *bool
	/*Stderr
	  Attach to `stderr`

	*/
	Stderr *bool
	/*Stdin
	  Attach to `stdin`

	*/
	Stdin *bool
	/*Stdout
	  Attach to `stdout`

	*/
	Stdout *bool
	/*Stream
	  Stream attached streams from the time the request was made onwards

	*/
	Stream *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container attach params
func (o *ContainerAttachParams) WithTimeout(timeout time.Duration) *ContainerAttachParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container attach params
func (o *ContainerAttachParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container attach params
func (o *ContainerAttachParams) WithContext(ctx context.Context) *ContainerAttachParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container attach params
func (o *ContainerAttachParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container attach params
func (o *ContainerAttachParams) WithHTTPClient(client *http.Client) *ContainerAttachParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container attach params
func (o *ContainerAttachParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetachKeys adds the detachKeys to the container attach params
func (o *ContainerAttachParams) WithDetachKeys(detachKeys *string) *ContainerAttachParams {
	o.SetDetachKeys(detachKeys)
	return o
}

// SetDetachKeys adds the detachKeys to the container attach params
func (o *ContainerAttachParams) SetDetachKeys(detachKeys *string) {
	o.DetachKeys = detachKeys
}

// WithID adds the id to the container attach params
func (o *ContainerAttachParams) WithID(id string) *ContainerAttachParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container attach params
func (o *ContainerAttachParams) SetID(id string) {
	o.ID = id
}

// WithLogs adds the logs to the container attach params
func (o *ContainerAttachParams) WithLogs(logs *bool) *ContainerAttachParams {
	o.SetLogs(logs)
	return o
}

// SetLogs adds the logs to the container attach params
func (o *ContainerAttachParams) SetLogs(logs *bool) {
	o.Logs = logs
}

// WithStderr adds the stderr to the container attach params
func (o *ContainerAttachParams) WithStderr(stderr *bool) *ContainerAttachParams {
	o.SetStderr(stderr)
	return o
}

// SetStderr adds the stderr to the container attach params
func (o *ContainerAttachParams) SetStderr(stderr *bool) {
	o.Stderr = stderr
}

// WithStdin adds the stdin to the container attach params
func (o *ContainerAttachParams) WithStdin(stdin *bool) *ContainerAttachParams {
	o.SetStdin(stdin)
	return o
}

// SetStdin adds the stdin to the container attach params
func (o *ContainerAttachParams) SetStdin(stdin *bool) {
	o.Stdin = stdin
}

// WithStdout adds the stdout to the container attach params
func (o *ContainerAttachParams) WithStdout(stdout *bool) *ContainerAttachParams {
	o.SetStdout(stdout)
	return o
}

// SetStdout adds the stdout to the container attach params
func (o *ContainerAttachParams) SetStdout(stdout *bool) {
	o.Stdout = stdout
}

// WithStream adds the stream to the container attach params
func (o *ContainerAttachParams) WithStream(stream *bool) *ContainerAttachParams {
	o.SetStream(stream)
	return o
}

// SetStream adds the stream to the container attach params
func (o *ContainerAttachParams) SetStream(stream *bool) {
	o.Stream = stream
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerAttachParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DetachKeys != nil {

		// query param detachKeys
		var qrDetachKeys string
		if o.DetachKeys != nil {
			qrDetachKeys = *o.DetachKeys
		}
		qDetachKeys := qrDetachKeys
		if qDetachKeys != "" {
			if err := r.SetQueryParam("detachKeys", qDetachKeys); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Logs != nil {

		// query param logs
		var qrLogs bool
		if o.Logs != nil {
			qrLogs = *o.Logs
		}
		qLogs := swag.FormatBool(qrLogs)
		if qLogs != "" {
			if err := r.SetQueryParam("logs", qLogs); err != nil {
				return err
			}
		}

	}

	if o.Stderr != nil {

		// query param stderr
		var qrStderr bool
		if o.Stderr != nil {
			qrStderr = *o.Stderr
		}
		qStderr := swag.FormatBool(qrStderr)
		if qStderr != "" {
			if err := r.SetQueryParam("stderr", qStderr); err != nil {
				return err
			}
		}

	}

	if o.Stdin != nil {

		// query param stdin
		var qrStdin bool
		if o.Stdin != nil {
			qrStdin = *o.Stdin
		}
		qStdin := swag.FormatBool(qrStdin)
		if qStdin != "" {
			if err := r.SetQueryParam("stdin", qStdin); err != nil {
				return err
			}
		}

	}

	if o.Stdout != nil {

		// query param stdout
		var qrStdout bool
		if o.Stdout != nil {
			qrStdout = *o.Stdout
		}
		qStdout := swag.FormatBool(qrStdout)
		if qStdout != "" {
			if err := r.SetQueryParam("stdout", qStdout); err != nil {
				return err
			}
		}

	}

	if o.Stream != nil {

		// query param stream
		var qrStream bool
		if o.Stream != nil {
			qrStream = *o.Stream
		}
		qStream := swag.FormatBool(qrStream)
		if qStream != "" {
			if err := r.SetQueryParam("stream", qStream); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
