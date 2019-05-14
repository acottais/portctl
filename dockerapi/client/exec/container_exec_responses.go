// Code generated by go-swagger; DO NOT EDIT.

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// ContainerExecReader is a Reader for the ContainerExec structure.
type ContainerExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewContainerExecCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerExecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewContainerExecConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerExecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerExecCreated creates a ContainerExecCreated with default headers values
func NewContainerExecCreated() *ContainerExecCreated {
	return &ContainerExecCreated{}
}

/*ContainerExecCreated handles this case with default header values.

no error
*/
type ContainerExecCreated struct {
	Payload *models.IDResponse
}

func (o *ContainerExecCreated) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecCreated  %+v", 201, o.Payload)
}

func (o *ContainerExecCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExecNotFound creates a ContainerExecNotFound with default headers values
func NewContainerExecNotFound() *ContainerExecNotFound {
	return &ContainerExecNotFound{}
}

/*ContainerExecNotFound handles this case with default header values.

no such container
*/
type ContainerExecNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerExecNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecNotFound  %+v", 404, o.Payload)
}

func (o *ContainerExecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExecConflict creates a ContainerExecConflict with default headers values
func NewContainerExecConflict() *ContainerExecConflict {
	return &ContainerExecConflict{}
}

/*ContainerExecConflict handles this case with default header values.

container is paused
*/
type ContainerExecConflict struct {
	Payload *models.ErrorResponse
}

func (o *ContainerExecConflict) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecConflict  %+v", 409, o.Payload)
}

func (o *ContainerExecConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExecInternalServerError creates a ContainerExecInternalServerError with default headers values
func NewContainerExecInternalServerError() *ContainerExecInternalServerError {
	return &ContainerExecInternalServerError{}
}

/*ContainerExecInternalServerError handles this case with default header values.

Server error
*/
type ContainerExecInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerExecInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerExecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerExecBody container exec body
swagger:model ContainerExecBody
*/
type ContainerExecBody struct {

	// Attach to `stderr` of the exec command.
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// Attach to `stdin` of the exec command.
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Attach to `stdout` of the exec command.
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Command to run, as a string or array of strings.
	Cmd []string `json:"Cmd"`

	// Override the key sequence for detaching a container. Format is a single character `[a-Z]` or `ctrl-<value>` where `<value>` is one of: `a-z`, `@`, `^`, `[`, `,` or `_`.
	DetachKeys string `json:"DetachKeys,omitempty"`

	// A list of environment variables in the form `["VAR=value", ...]`.
	Env []string `json:"Env"`

	// Runs the exec process with extended privileges.
	Privileged *bool `json:"Privileged,omitempty"`

	// Allocate a pseudo-TTY.
	Tty bool `json:"Tty,omitempty"`

	// The user, and optionally, group to run the exec process inside the container. Format is one of: `user`, `user:group`, `uid`, or `uid:gid`.
	User string `json:"User,omitempty"`

	// The working directory for the exec process inside the container.
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this container exec body
func (o *ContainerExecBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExecBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExecBody) UnmarshalBinary(b []byte) error {
	var res ContainerExecBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
