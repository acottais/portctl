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

// ContainerLogsReader is a Reader for the ContainerLogs structure.
type ContainerLogsReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ContainerLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerLogsOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerLogsOK creates a ContainerLogsOK with default headers values
func NewContainerLogsOK(writer io.Writer) *ContainerLogsOK {
	return &ContainerLogsOK{
		Payload: writer,
	}
}

/*ContainerLogsOK handles this case with default header values.

logs returned as a stream in response body.
For the stream format, [see the documentation for the attach endpoint](#operation/ContainerAttach).
Note that unlike the attach endpoint, the logs endpoint does not upgrade the connection and does not
set Content-Type.

*/
type ContainerLogsOK struct {
	Payload io.Writer
}

func (o *ContainerLogsOK) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/logs][%d] containerLogsOK  %+v", 200, o.Payload)
}

func (o *ContainerLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerLogsNotFound creates a ContainerLogsNotFound with default headers values
func NewContainerLogsNotFound() *ContainerLogsNotFound {
	return &ContainerLogsNotFound{}
}

/*ContainerLogsNotFound handles this case with default header values.

no such container
*/
type ContainerLogsNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/logs][%d] containerLogsNotFound  %+v", 404, o.Payload)
}

func (o *ContainerLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerLogsInternalServerError creates a ContainerLogsInternalServerError with default headers values
func NewContainerLogsInternalServerError() *ContainerLogsInternalServerError {
	return &ContainerLogsInternalServerError{}
}

/*ContainerLogsInternalServerError handles this case with default header values.

server error
*/
type ContainerLogsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/logs][%d] containerLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
