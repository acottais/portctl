// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/portainerapi/models"
)

// RegistryDeleteReader is a Reader for the RegistryDelete structure.
type RegistryDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRegistryDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRegistryDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRegistryDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRegistryDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegistryDeleteNoContent creates a RegistryDeleteNoContent with default headers values
func NewRegistryDeleteNoContent() *RegistryDeleteNoContent {
	return &RegistryDeleteNoContent{}
}

/*RegistryDeleteNoContent handles this case with default header values.

Success
*/
type RegistryDeleteNoContent struct {
}

func (o *RegistryDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] registryDeleteNoContent ", 204)
}

func (o *RegistryDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryDeleteBadRequest creates a RegistryDeleteBadRequest with default headers values
func NewRegistryDeleteBadRequest() *RegistryDeleteBadRequest {
	return &RegistryDeleteBadRequest{}
}

/*RegistryDeleteBadRequest handles this case with default header values.

Invalid request
*/
type RegistryDeleteBadRequest struct {
	Payload *models.GenericError
}

func (o *RegistryDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] registryDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *RegistryDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryDeleteNotFound creates a RegistryDeleteNotFound with default headers values
func NewRegistryDeleteNotFound() *RegistryDeleteNotFound {
	return &RegistryDeleteNotFound{}
}

/*RegistryDeleteNotFound handles this case with default header values.

Registry not found
*/
type RegistryDeleteNotFound struct {
	Payload *models.GenericError
}

func (o *RegistryDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] registryDeleteNotFound  %+v", 404, o.Payload)
}

func (o *RegistryDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryDeleteInternalServerError creates a RegistryDeleteInternalServerError with default headers values
func NewRegistryDeleteInternalServerError() *RegistryDeleteInternalServerError {
	return &RegistryDeleteInternalServerError{}
}

/*RegistryDeleteInternalServerError handles this case with default header values.

Server error
*/
type RegistryDeleteInternalServerError struct {
	Payload *models.GenericError
}

func (o *RegistryDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] registryDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *RegistryDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
