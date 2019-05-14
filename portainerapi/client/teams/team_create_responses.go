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

// TeamCreateReader is a Reader for the TeamCreate structure.
type TeamCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewTeamCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewTeamCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTeamCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamCreateOK creates a TeamCreateOK with default headers values
func NewTeamCreateOK() *TeamCreateOK {
	return &TeamCreateOK{}
}

/*TeamCreateOK handles this case with default header values.

Success
*/
type TeamCreateOK struct {
	Payload *models.Team
}

func (o *TeamCreateOK) Error() string {
	return fmt.Sprintf("[POST /teams][%d] teamCreateOK  %+v", 200, o.Payload)
}

func (o *TeamCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamCreateBadRequest creates a TeamCreateBadRequest with default headers values
func NewTeamCreateBadRequest() *TeamCreateBadRequest {
	return &TeamCreateBadRequest{}
}

/*TeamCreateBadRequest handles this case with default header values.

Invalid request
*/
type TeamCreateBadRequest struct {
	Payload *models.GenericError
}

func (o *TeamCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /teams][%d] teamCreateBadRequest  %+v", 400, o.Payload)
}

func (o *TeamCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamCreateForbidden creates a TeamCreateForbidden with default headers values
func NewTeamCreateForbidden() *TeamCreateForbidden {
	return &TeamCreateForbidden{}
}

/*TeamCreateForbidden handles this case with default header values.

Unauthorized
*/
type TeamCreateForbidden struct {
	Payload *models.GenericError
}

func (o *TeamCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /teams][%d] teamCreateForbidden  %+v", 403, o.Payload)
}

func (o *TeamCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamCreateConflict creates a TeamCreateConflict with default headers values
func NewTeamCreateConflict() *TeamCreateConflict {
	return &TeamCreateConflict{}
}

/*TeamCreateConflict handles this case with default header values.

Team already exists
*/
type TeamCreateConflict struct {
	Payload *models.GenericError
}

func (o *TeamCreateConflict) Error() string {
	return fmt.Sprintf("[POST /teams][%d] teamCreateConflict  %+v", 409, o.Payload)
}

func (o *TeamCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamCreateInternalServerError creates a TeamCreateInternalServerError with default headers values
func NewTeamCreateInternalServerError() *TeamCreateInternalServerError {
	return &TeamCreateInternalServerError{}
}

/*TeamCreateInternalServerError handles this case with default header values.

Server error
*/
type TeamCreateInternalServerError struct {
	Payload *models.GenericError
}

func (o *TeamCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /teams][%d] teamCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *TeamCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}