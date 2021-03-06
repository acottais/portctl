// Code generated by go-swagger; DO NOT EDIT.

package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/portainerapi/models"
)

// TemplateCreateReader is a Reader for the TemplateCreate structure.
type TemplateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTemplateCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTemplateCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewTemplateCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTemplateCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTemplateCreateOK creates a TemplateCreateOK with default headers values
func NewTemplateCreateOK() *TemplateCreateOK {
	return &TemplateCreateOK{}
}

/*TemplateCreateOK handles this case with default header values.

Success
*/
type TemplateCreateOK struct {
	Payload *models.Template
}

func (o *TemplateCreateOK) Error() string {
	return fmt.Sprintf("[POST /templates][%d] templateCreateOK  %+v", 200, o.Payload)
}

func (o *TemplateCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Template)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplateCreateBadRequest creates a TemplateCreateBadRequest with default headers values
func NewTemplateCreateBadRequest() *TemplateCreateBadRequest {
	return &TemplateCreateBadRequest{}
}

/*TemplateCreateBadRequest handles this case with default header values.

Invalid request
*/
type TemplateCreateBadRequest struct {
	Payload *models.GenericError
}

func (o *TemplateCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /templates][%d] templateCreateBadRequest  %+v", 400, o.Payload)
}

func (o *TemplateCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplateCreateForbidden creates a TemplateCreateForbidden with default headers values
func NewTemplateCreateForbidden() *TemplateCreateForbidden {
	return &TemplateCreateForbidden{}
}

/*TemplateCreateForbidden handles this case with default header values.

Unauthorized
*/
type TemplateCreateForbidden struct {
	Payload *models.GenericError
}

func (o *TemplateCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /templates][%d] templateCreateForbidden  %+v", 403, o.Payload)
}

func (o *TemplateCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplateCreateInternalServerError creates a TemplateCreateInternalServerError with default headers values
func NewTemplateCreateInternalServerError() *TemplateCreateInternalServerError {
	return &TemplateCreateInternalServerError{}
}

/*TemplateCreateInternalServerError handles this case with default header values.

Server error
*/
type TemplateCreateInternalServerError struct {
	Payload *models.GenericError
}

func (o *TemplateCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /templates][%d] templateCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *TemplateCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
