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

// UserUpdateReader is a Reader for the UserUpdate structure.
type UserUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUserUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUserUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserUpdateOK creates a UserUpdateOK with default headers values
func NewUserUpdateOK() *UserUpdateOK {
	return &UserUpdateOK{}
}

/*UserUpdateOK handles this case with default header values.

Success
*/
type UserUpdateOK struct {
	Payload *models.User
}

func (o *UserUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateOK  %+v", 200, o.Payload)
}

func (o *UserUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateBadRequest creates a UserUpdateBadRequest with default headers values
func NewUserUpdateBadRequest() *UserUpdateBadRequest {
	return &UserUpdateBadRequest{}
}

/*UserUpdateBadRequest handles this case with default header values.

Invalid request
*/
type UserUpdateBadRequest struct {
	Payload *models.GenericError
}

func (o *UserUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UserUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateForbidden creates a UserUpdateForbidden with default headers values
func NewUserUpdateForbidden() *UserUpdateForbidden {
	return &UserUpdateForbidden{}
}

/*UserUpdateForbidden handles this case with default header values.

Unauthorized
*/
type UserUpdateForbidden struct {
	Payload *models.GenericError
}

func (o *UserUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateForbidden  %+v", 403, o.Payload)
}

func (o *UserUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateNotFound creates a UserUpdateNotFound with default headers values
func NewUserUpdateNotFound() *UserUpdateNotFound {
	return &UserUpdateNotFound{}
}

/*UserUpdateNotFound handles this case with default header values.

User not found
*/
type UserUpdateNotFound struct {
	Payload *models.GenericError
}

func (o *UserUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UserUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateInternalServerError creates a UserUpdateInternalServerError with default headers values
func NewUserUpdateInternalServerError() *UserUpdateInternalServerError {
	return &UserUpdateInternalServerError{}
}

/*UserUpdateInternalServerError handles this case with default header values.

Server error
*/
type UserUpdateInternalServerError struct {
	Payload *models.GenericError
}

func (o *UserUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *UserUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
