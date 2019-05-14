// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/portainerapi/models"
)

// TeamMembershipUpdateReader is a Reader for the TeamMembershipUpdate structure.
type TeamMembershipUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamMembershipUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamMembershipUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamMembershipUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewTeamMembershipUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTeamMembershipUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTeamMembershipUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamMembershipUpdateOK creates a TeamMembershipUpdateOK with default headers values
func NewTeamMembershipUpdateOK() *TeamMembershipUpdateOK {
	return &TeamMembershipUpdateOK{}
}

/*TeamMembershipUpdateOK handles this case with default header values.

Success
*/
type TeamMembershipUpdateOK struct {
	Payload *models.TeamMembership
}

func (o *TeamMembershipUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateOK  %+v", 200, o.Payload)
}

func (o *TeamMembershipUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamMembershipUpdateBadRequest creates a TeamMembershipUpdateBadRequest with default headers values
func NewTeamMembershipUpdateBadRequest() *TeamMembershipUpdateBadRequest {
	return &TeamMembershipUpdateBadRequest{}
}

/*TeamMembershipUpdateBadRequest handles this case with default header values.

Invalid request
*/
type TeamMembershipUpdateBadRequest struct {
	Payload *models.GenericError
}

func (o *TeamMembershipUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TeamMembershipUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamMembershipUpdateForbidden creates a TeamMembershipUpdateForbidden with default headers values
func NewTeamMembershipUpdateForbidden() *TeamMembershipUpdateForbidden {
	return &TeamMembershipUpdateForbidden{}
}

/*TeamMembershipUpdateForbidden handles this case with default header values.

Unauthorized
*/
type TeamMembershipUpdateForbidden struct {
	Payload *models.GenericError
}

func (o *TeamMembershipUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateForbidden  %+v", 403, o.Payload)
}

func (o *TeamMembershipUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamMembershipUpdateNotFound creates a TeamMembershipUpdateNotFound with default headers values
func NewTeamMembershipUpdateNotFound() *TeamMembershipUpdateNotFound {
	return &TeamMembershipUpdateNotFound{}
}

/*TeamMembershipUpdateNotFound handles this case with default header values.

Team membership not found
*/
type TeamMembershipUpdateNotFound struct {
	Payload *models.GenericError
}

func (o *TeamMembershipUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateNotFound  %+v", 404, o.Payload)
}

func (o *TeamMembershipUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamMembershipUpdateInternalServerError creates a TeamMembershipUpdateInternalServerError with default headers values
func NewTeamMembershipUpdateInternalServerError() *TeamMembershipUpdateInternalServerError {
	return &TeamMembershipUpdateInternalServerError{}
}

/*TeamMembershipUpdateInternalServerError handles this case with default header values.

Server error
*/
type TeamMembershipUpdateInternalServerError struct {
	Payload *models.GenericError
}

func (o *TeamMembershipUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *TeamMembershipUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
