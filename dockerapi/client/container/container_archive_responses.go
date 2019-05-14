// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// ContainerArchiveReader is a Reader for the ContainerArchive structure.
type ContainerArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewContainerArchiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContainerArchiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerArchiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerArchiveOK creates a ContainerArchiveOK with default headers values
func NewContainerArchiveOK() *ContainerArchiveOK {
	return &ContainerArchiveOK{}
}

/*ContainerArchiveOK handles this case with default header values.

no error
*/
type ContainerArchiveOK struct {
}

func (o *ContainerArchiveOK) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/archive][%d] containerArchiveOK ", 200)
}

func (o *ContainerArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerArchiveBadRequest creates a ContainerArchiveBadRequest with default headers values
func NewContainerArchiveBadRequest() *ContainerArchiveBadRequest {
	return &ContainerArchiveBadRequest{}
}

/*ContainerArchiveBadRequest handles this case with default header values.

Bad parameter
*/
type ContainerArchiveBadRequest struct {
	Payload *ContainerArchiveBadRequestBody
}

func (o *ContainerArchiveBadRequest) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/archive][%d] containerArchiveBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerArchiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerArchiveBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerArchiveNotFound creates a ContainerArchiveNotFound with default headers values
func NewContainerArchiveNotFound() *ContainerArchiveNotFound {
	return &ContainerArchiveNotFound{}
}

/*ContainerArchiveNotFound handles this case with default header values.

Container or path does not exist
*/
type ContainerArchiveNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerArchiveNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/archive][%d] containerArchiveNotFound  %+v", 404, o.Payload)
}

func (o *ContainerArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerArchiveInternalServerError creates a ContainerArchiveInternalServerError with default headers values
func NewContainerArchiveInternalServerError() *ContainerArchiveInternalServerError {
	return &ContainerArchiveInternalServerError{}
}

/*ContainerArchiveInternalServerError handles this case with default header values.

server error
*/
type ContainerArchiveInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerArchiveInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/archive][%d] containerArchiveInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerArchiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerArchiveBadRequestBody container archive bad request body
swagger:model ContainerArchiveBadRequestBody
*/
type ContainerArchiveBadRequestBody struct {
	models.ErrorResponse

	// The error message. Either "must specify path parameter" (path cannot be empty) or "not a directory" (path was asserted to be a directory but exists as a file).
	Message string `json:"message,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ContainerArchiveBadRequestBody) UnmarshalJSON(raw []byte) error {
	// ContainerArchiveBadRequestBodyAO0
	var containerArchiveBadRequestBodyAO0 models.ErrorResponse
	if err := swag.ReadJSON(raw, &containerArchiveBadRequestBodyAO0); err != nil {
		return err
	}
	o.ErrorResponse = containerArchiveBadRequestBodyAO0

	// ContainerArchiveBadRequestBodyAO1
	var dataContainerArchiveBadRequestBodyAO1 struct {
		Message string `json:"message,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataContainerArchiveBadRequestBodyAO1); err != nil {
		return err
	}

	o.Message = dataContainerArchiveBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ContainerArchiveBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	containerArchiveBadRequestBodyAO0, err := swag.WriteJSON(o.ErrorResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, containerArchiveBadRequestBodyAO0)

	var dataContainerArchiveBadRequestBodyAO1 struct {
		Message string `json:"message,omitempty"`
	}

	dataContainerArchiveBadRequestBodyAO1.Message = o.Message

	jsonDataContainerArchiveBadRequestBodyAO1, errContainerArchiveBadRequestBodyAO1 := swag.WriteJSON(dataContainerArchiveBadRequestBodyAO1)
	if errContainerArchiveBadRequestBodyAO1 != nil {
		return nil, errContainerArchiveBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataContainerArchiveBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this container archive bad request body
func (o *ContainerArchiveBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerArchiveBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerArchiveBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ContainerArchiveBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
