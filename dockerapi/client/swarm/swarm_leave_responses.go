// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// SwarmLeaveReader is a Reader for the SwarmLeave structure.
type SwarmLeaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwarmLeaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSwarmLeaveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSwarmLeaveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewSwarmLeaveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSwarmLeaveOK creates a SwarmLeaveOK with default headers values
func NewSwarmLeaveOK() *SwarmLeaveOK {
	return &SwarmLeaveOK{}
}

/*SwarmLeaveOK handles this case with default header values.

no error
*/
type SwarmLeaveOK struct {
}

func (o *SwarmLeaveOK) Error() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveOK ", 200)
}

func (o *SwarmLeaveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSwarmLeaveInternalServerError creates a SwarmLeaveInternalServerError with default headers values
func NewSwarmLeaveInternalServerError() *SwarmLeaveInternalServerError {
	return &SwarmLeaveInternalServerError{}
}

/*SwarmLeaveInternalServerError handles this case with default header values.

server error
*/
type SwarmLeaveInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SwarmLeaveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveInternalServerError  %+v", 500, o.Payload)
}

func (o *SwarmLeaveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwarmLeaveServiceUnavailable creates a SwarmLeaveServiceUnavailable with default headers values
func NewSwarmLeaveServiceUnavailable() *SwarmLeaveServiceUnavailable {
	return &SwarmLeaveServiceUnavailable{}
}

/*SwarmLeaveServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type SwarmLeaveServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *SwarmLeaveServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SwarmLeaveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
