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

// RegistryListReader is a Reader for the RegistryList structure.
type RegistryListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRegistryListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewRegistryListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegistryListOK creates a RegistryListOK with default headers values
func NewRegistryListOK() *RegistryListOK {
	return &RegistryListOK{}
}

/*RegistryListOK handles this case with default header values.

Success
*/
type RegistryListOK struct {
	Payload models.RegistryListResponse
}

func (o *RegistryListOK) Error() string {
	return fmt.Sprintf("[GET /registries][%d] registryListOK  %+v", 200, o.Payload)
}

func (o *RegistryListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryListInternalServerError creates a RegistryListInternalServerError with default headers values
func NewRegistryListInternalServerError() *RegistryListInternalServerError {
	return &RegistryListInternalServerError{}
}

/*RegistryListInternalServerError handles this case with default header values.

Server error
*/
type RegistryListInternalServerError struct {
	Payload *models.GenericError
}

func (o *RegistryListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries][%d] registryListInternalServerError  %+v", 500, o.Payload)
}

func (o *RegistryListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}