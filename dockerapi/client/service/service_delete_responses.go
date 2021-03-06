// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// ServiceDeleteReader is a Reader for the ServiceDelete structure.
type ServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewServiceDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewServiceDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewServiceDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceDeleteOK creates a ServiceDeleteOK with default headers values
func NewServiceDeleteOK() *ServiceDeleteOK {
	return &ServiceDeleteOK{}
}

/*ServiceDeleteOK handles this case with default header values.

no error
*/
type ServiceDeleteOK struct {
}

func (o *ServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /services/{id}][%d] serviceDeleteOK ", 200)
}

func (o *ServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceDeleteNotFound creates a ServiceDeleteNotFound with default headers values
func NewServiceDeleteNotFound() *ServiceDeleteNotFound {
	return &ServiceDeleteNotFound{}
}

/*ServiceDeleteNotFound handles this case with default header values.

no such service
*/
type ServiceDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ServiceDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/{id}][%d] serviceDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ServiceDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDeleteInternalServerError creates a ServiceDeleteInternalServerError with default headers values
func NewServiceDeleteInternalServerError() *ServiceDeleteInternalServerError {
	return &ServiceDeleteInternalServerError{}
}

/*ServiceDeleteInternalServerError handles this case with default header values.

server error
*/
type ServiceDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ServiceDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /services/{id}][%d] serviceDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDeleteServiceUnavailable creates a ServiceDeleteServiceUnavailable with default headers values
func NewServiceDeleteServiceUnavailable() *ServiceDeleteServiceUnavailable {
	return &ServiceDeleteServiceUnavailable{}
}

/*ServiceDeleteServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type ServiceDeleteServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ServiceDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /services/{id}][%d] serviceDeleteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ServiceDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
