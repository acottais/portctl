// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/portainerapi/models"
)

// TeamDeleteReader is a Reader for the TeamDelete structure.
type TeamDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewTeamDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTeamDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTeamDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamDeleteNoContent creates a TeamDeleteNoContent with default headers values
func NewTeamDeleteNoContent() *TeamDeleteNoContent {
	return &TeamDeleteNoContent{}
}

/*TeamDeleteNoContent handles this case with default header values.

Success
*/
type TeamDeleteNoContent struct {
}

func (o *TeamDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteNoContent ", 204)
}

func (o *TeamDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamDeleteBadRequest creates a TeamDeleteBadRequest with default headers values
func NewTeamDeleteBadRequest() *TeamDeleteBadRequest {
	return &TeamDeleteBadRequest{}
}

/*TeamDeleteBadRequest handles this case with default header values.

Invalid request
*/
type TeamDeleteBadRequest struct {
	Payload *models.GenericError
}

func (o *TeamDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TeamDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamDeleteNotFound creates a TeamDeleteNotFound with default headers values
func NewTeamDeleteNotFound() *TeamDeleteNotFound {
	return &TeamDeleteNotFound{}
}

/*TeamDeleteNotFound handles this case with default header values.

Team not found
*/
type TeamDeleteNotFound struct {
	Payload *models.GenericError
}

func (o *TeamDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteNotFound  %+v", 404, o.Payload)
}

func (o *TeamDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamDeleteInternalServerError creates a TeamDeleteInternalServerError with default headers values
func NewTeamDeleteInternalServerError() *TeamDeleteInternalServerError {
	return &TeamDeleteInternalServerError{}
}

/*TeamDeleteInternalServerError handles this case with default header values.

Server error
*/
type TeamDeleteInternalServerError struct {
	Payload *models.GenericError
}

func (o *TeamDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *TeamDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
