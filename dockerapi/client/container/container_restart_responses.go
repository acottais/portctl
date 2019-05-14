// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// ContainerRestartReader is a Reader for the ContainerRestart structure.
type ContainerRestartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRestartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewContainerRestartNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerRestartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerRestartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerRestartNoContent creates a ContainerRestartNoContent with default headers values
func NewContainerRestartNoContent() *ContainerRestartNoContent {
	return &ContainerRestartNoContent{}
}

/*ContainerRestartNoContent handles this case with default header values.

no error
*/
type ContainerRestartNoContent struct {
}

func (o *ContainerRestartNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/restart][%d] containerRestartNoContent ", 204)
}

func (o *ContainerRestartNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRestartNotFound creates a ContainerRestartNotFound with default headers values
func NewContainerRestartNotFound() *ContainerRestartNotFound {
	return &ContainerRestartNotFound{}
}

/*ContainerRestartNotFound handles this case with default header values.

no such container
*/
type ContainerRestartNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerRestartNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/restart][%d] containerRestartNotFound  %+v", 404, o.Payload)
}

func (o *ContainerRestartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRestartInternalServerError creates a ContainerRestartInternalServerError with default headers values
func NewContainerRestartInternalServerError() *ContainerRestartInternalServerError {
	return &ContainerRestartInternalServerError{}
}

/*ContainerRestartInternalServerError handles this case with default header values.

server error
*/
type ContainerRestartInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerRestartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/restart][%d] containerRestartInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerRestartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
