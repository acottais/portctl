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

// RegistryUpdateReader is a Reader for the RegistryUpdate structure.
type RegistryUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRegistryUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRegistryUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRegistryUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRegistryUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRegistryUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegistryUpdateOK creates a RegistryUpdateOK with default headers values
func NewRegistryUpdateOK() *RegistryUpdateOK {
	return &RegistryUpdateOK{}
}

/*RegistryUpdateOK handles this case with default header values.

Success
*/
type RegistryUpdateOK struct {
	Payload *models.Registry
}

func (o *RegistryUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateOK  %+v", 200, o.Payload)
}

func (o *RegistryUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Registry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryUpdateBadRequest creates a RegistryUpdateBadRequest with default headers values
func NewRegistryUpdateBadRequest() *RegistryUpdateBadRequest {
	return &RegistryUpdateBadRequest{}
}

/*RegistryUpdateBadRequest handles this case with default header values.

Invalid request
*/
type RegistryUpdateBadRequest struct {
	Payload *models.GenericError
}

func (o *RegistryUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *RegistryUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryUpdateNotFound creates a RegistryUpdateNotFound with default headers values
func NewRegistryUpdateNotFound() *RegistryUpdateNotFound {
	return &RegistryUpdateNotFound{}
}

/*RegistryUpdateNotFound handles this case with default header values.

Registry not found
*/
type RegistryUpdateNotFound struct {
	Payload *models.GenericError
}

func (o *RegistryUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateNotFound  %+v", 404, o.Payload)
}

func (o *RegistryUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryUpdateConflict creates a RegistryUpdateConflict with default headers values
func NewRegistryUpdateConflict() *RegistryUpdateConflict {
	return &RegistryUpdateConflict{}
}

/*RegistryUpdateConflict handles this case with default header values.

Registry already exists
*/
type RegistryUpdateConflict struct {
	Payload *models.GenericError
}

func (o *RegistryUpdateConflict) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateConflict  %+v", 409, o.Payload)
}

func (o *RegistryUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryUpdateInternalServerError creates a RegistryUpdateInternalServerError with default headers values
func NewRegistryUpdateInternalServerError() *RegistryUpdateInternalServerError {
	return &RegistryUpdateInternalServerError{}
}

/*RegistryUpdateInternalServerError handles this case with default header values.

Server error
*/
type RegistryUpdateInternalServerError struct {
	Payload *models.GenericError
}

func (o *RegistryUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *RegistryUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
