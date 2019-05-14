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

// ContainerUnpauseReader is a Reader for the ContainerUnpause structure.
type ContainerUnpauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerUnpauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewContainerUnpauseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerUnpauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerUnpauseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerUnpauseNoContent creates a ContainerUnpauseNoContent with default headers values
func NewContainerUnpauseNoContent() *ContainerUnpauseNoContent {
	return &ContainerUnpauseNoContent{}
}

/*ContainerUnpauseNoContent handles this case with default header values.

no error
*/
type ContainerUnpauseNoContent struct {
}

func (o *ContainerUnpauseNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/unpause][%d] containerUnpauseNoContent ", 204)
}

func (o *ContainerUnpauseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerUnpauseNotFound creates a ContainerUnpauseNotFound with default headers values
func NewContainerUnpauseNotFound() *ContainerUnpauseNotFound {
	return &ContainerUnpauseNotFound{}
}

/*ContainerUnpauseNotFound handles this case with default header values.

no such container
*/
type ContainerUnpauseNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerUnpauseNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/unpause][%d] containerUnpauseNotFound  %+v", 404, o.Payload)
}

func (o *ContainerUnpauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerUnpauseInternalServerError creates a ContainerUnpauseInternalServerError with default headers values
func NewContainerUnpauseInternalServerError() *ContainerUnpauseInternalServerError {
	return &ContainerUnpauseInternalServerError{}
}

/*ContainerUnpauseInternalServerError handles this case with default header values.

server error
*/
type ContainerUnpauseInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerUnpauseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/unpause][%d] containerUnpauseInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerUnpauseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}