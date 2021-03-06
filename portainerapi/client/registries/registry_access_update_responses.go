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

// RegistryAccessUpdateReader is a Reader for the RegistryAccessUpdate structure.
type RegistryAccessUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryAccessUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRegistryAccessUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRegistryAccessUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRegistryAccessUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRegistryAccessUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegistryAccessUpdateOK creates a RegistryAccessUpdateOK with default headers values
func NewRegistryAccessUpdateOK() *RegistryAccessUpdateOK {
	return &RegistryAccessUpdateOK{}
}

/*RegistryAccessUpdateOK handles this case with default header values.

Success
*/
type RegistryAccessUpdateOK struct {
	Payload *models.Registry
}

func (o *RegistryAccessUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}/access][%d] registryAccessUpdateOK  %+v", 200, o.Payload)
}

func (o *RegistryAccessUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Registry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryAccessUpdateBadRequest creates a RegistryAccessUpdateBadRequest with default headers values
func NewRegistryAccessUpdateBadRequest() *RegistryAccessUpdateBadRequest {
	return &RegistryAccessUpdateBadRequest{}
}

/*RegistryAccessUpdateBadRequest handles this case with default header values.

Invalid request
*/
type RegistryAccessUpdateBadRequest struct {
	Payload *models.GenericError
}

func (o *RegistryAccessUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}/access][%d] registryAccessUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *RegistryAccessUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryAccessUpdateNotFound creates a RegistryAccessUpdateNotFound with default headers values
func NewRegistryAccessUpdateNotFound() *RegistryAccessUpdateNotFound {
	return &RegistryAccessUpdateNotFound{}
}

/*RegistryAccessUpdateNotFound handles this case with default header values.

Registry not found
*/
type RegistryAccessUpdateNotFound struct {
	Payload *models.GenericError
}

func (o *RegistryAccessUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}/access][%d] registryAccessUpdateNotFound  %+v", 404, o.Payload)
}

func (o *RegistryAccessUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryAccessUpdateInternalServerError creates a RegistryAccessUpdateInternalServerError with default headers values
func NewRegistryAccessUpdateInternalServerError() *RegistryAccessUpdateInternalServerError {
	return &RegistryAccessUpdateInternalServerError{}
}

/*RegistryAccessUpdateInternalServerError handles this case with default header values.

Server error
*/
type RegistryAccessUpdateInternalServerError struct {
	Payload *models.GenericError
}

func (o *RegistryAccessUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}/access][%d] registryAccessUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *RegistryAccessUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
