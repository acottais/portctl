// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/portainerapi/models"
)

// EndpointGroupUpdateReader is a Reader for the EndpointGroupUpdate structure.
type EndpointGroupUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointGroupUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEndpointGroupUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEndpointGroupUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEndpointGroupUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEndpointGroupUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewEndpointGroupUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEndpointGroupUpdateOK creates a EndpointGroupUpdateOK with default headers values
func NewEndpointGroupUpdateOK() *EndpointGroupUpdateOK {
	return &EndpointGroupUpdateOK{}
}

/*EndpointGroupUpdateOK handles this case with default header values.

Success
*/
type EndpointGroupUpdateOK struct {
	Payload *models.EndpointGroup
}

func (o *EndpointGroupUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointGroupUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointGroupUpdateBadRequest creates a EndpointGroupUpdateBadRequest with default headers values
func NewEndpointGroupUpdateBadRequest() *EndpointGroupUpdateBadRequest {
	return &EndpointGroupUpdateBadRequest{}
}

/*EndpointGroupUpdateBadRequest handles this case with default header values.

Invalid request
*/
type EndpointGroupUpdateBadRequest struct {
	Payload *models.GenericError
}

func (o *EndpointGroupUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *EndpointGroupUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointGroupUpdateNotFound creates a EndpointGroupUpdateNotFound with default headers values
func NewEndpointGroupUpdateNotFound() *EndpointGroupUpdateNotFound {
	return &EndpointGroupUpdateNotFound{}
}

/*EndpointGroupUpdateNotFound handles this case with default header values.

EndpointGroup not found
*/
type EndpointGroupUpdateNotFound struct {
	Payload *models.GenericError
}

func (o *EndpointGroupUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateNotFound  %+v", 404, o.Payload)
}

func (o *EndpointGroupUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointGroupUpdateInternalServerError creates a EndpointGroupUpdateInternalServerError with default headers values
func NewEndpointGroupUpdateInternalServerError() *EndpointGroupUpdateInternalServerError {
	return &EndpointGroupUpdateInternalServerError{}
}

/*EndpointGroupUpdateInternalServerError handles this case with default header values.

Server error
*/
type EndpointGroupUpdateInternalServerError struct {
	Payload *models.GenericError
}

func (o *EndpointGroupUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *EndpointGroupUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointGroupUpdateServiceUnavailable creates a EndpointGroupUpdateServiceUnavailable with default headers values
func NewEndpointGroupUpdateServiceUnavailable() *EndpointGroupUpdateServiceUnavailable {
	return &EndpointGroupUpdateServiceUnavailable{}
}

/*EndpointGroupUpdateServiceUnavailable handles this case with default header values.

EndpointGroup management disabled
*/
type EndpointGroupUpdateServiceUnavailable struct {
	Payload *models.GenericError
}

func (o *EndpointGroupUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *EndpointGroupUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
