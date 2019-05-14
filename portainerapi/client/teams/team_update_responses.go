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

// TeamUpdateReader is a Reader for the TeamUpdate structure.
type TeamUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTeamUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTeamUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamUpdateOK creates a TeamUpdateOK with default headers values
func NewTeamUpdateOK() *TeamUpdateOK {
	return &TeamUpdateOK{}
}

/*TeamUpdateOK handles this case with default header values.

Success
*/
type TeamUpdateOK struct {
}

func (o *TeamUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateOK ", 200)
}

func (o *TeamUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamUpdateBadRequest creates a TeamUpdateBadRequest with default headers values
func NewTeamUpdateBadRequest() *TeamUpdateBadRequest {
	return &TeamUpdateBadRequest{}
}

/*TeamUpdateBadRequest handles this case with default header values.

Invalid request
*/
type TeamUpdateBadRequest struct {
	Payload *models.GenericError
}

func (o *TeamUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TeamUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamUpdateNotFound creates a TeamUpdateNotFound with default headers values
func NewTeamUpdateNotFound() *TeamUpdateNotFound {
	return &TeamUpdateNotFound{}
}

/*TeamUpdateNotFound handles this case with default header values.

Team not found
*/
type TeamUpdateNotFound struct {
	Payload *models.GenericError
}

func (o *TeamUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateNotFound  %+v", 404, o.Payload)
}

func (o *TeamUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamUpdateInternalServerError creates a TeamUpdateInternalServerError with default headers values
func NewTeamUpdateInternalServerError() *TeamUpdateInternalServerError {
	return &TeamUpdateInternalServerError{}
}

/*TeamUpdateInternalServerError handles this case with default header values.

Server error
*/
type TeamUpdateInternalServerError struct {
	Payload *models.GenericError
}

func (o *TeamUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *TeamUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}