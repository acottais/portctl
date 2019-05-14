// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/portainerapi/models"
)

// UserMembershipsInspectReader is a Reader for the UserMembershipsInspect structure.
type UserMembershipsInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserMembershipsInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserMembershipsInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserMembershipsInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUserMembershipsInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUserMembershipsInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserMembershipsInspectOK creates a UserMembershipsInspectOK with default headers values
func NewUserMembershipsInspectOK() *UserMembershipsInspectOK {
	return &UserMembershipsInspectOK{}
}

/*UserMembershipsInspectOK handles this case with default header values.

Success
*/
type UserMembershipsInspectOK struct {
	Payload models.UserMembershipsResponse
}

func (o *UserMembershipsInspectOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/memberships][%d] userMembershipsInspectOK  %+v", 200, o.Payload)
}

func (o *UserMembershipsInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMembershipsInspectBadRequest creates a UserMembershipsInspectBadRequest with default headers values
func NewUserMembershipsInspectBadRequest() *UserMembershipsInspectBadRequest {
	return &UserMembershipsInspectBadRequest{}
}

/*UserMembershipsInspectBadRequest handles this case with default header values.

Invalid request
*/
type UserMembershipsInspectBadRequest struct {
	Payload *models.GenericError
}

func (o *UserMembershipsInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}/memberships][%d] userMembershipsInspectBadRequest  %+v", 400, o.Payload)
}

func (o *UserMembershipsInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMembershipsInspectForbidden creates a UserMembershipsInspectForbidden with default headers values
func NewUserMembershipsInspectForbidden() *UserMembershipsInspectForbidden {
	return &UserMembershipsInspectForbidden{}
}

/*UserMembershipsInspectForbidden handles this case with default header values.

Unauthorized
*/
type UserMembershipsInspectForbidden struct {
	Payload *models.GenericError
}

func (o *UserMembershipsInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}/memberships][%d] userMembershipsInspectForbidden  %+v", 403, o.Payload)
}

func (o *UserMembershipsInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMembershipsInspectInternalServerError creates a UserMembershipsInspectInternalServerError with default headers values
func NewUserMembershipsInspectInternalServerError() *UserMembershipsInspectInternalServerError {
	return &UserMembershipsInspectInternalServerError{}
}

/*UserMembershipsInspectInternalServerError handles this case with default header values.

Server error
*/
type UserMembershipsInspectInternalServerError struct {
	Payload *models.GenericError
}

func (o *UserMembershipsInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{id}/memberships][%d] userMembershipsInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *UserMembershipsInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}