// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/portainerapi/models"
)

// EndpointAccessUpdateReader is a Reader for the EndpointAccessUpdate structure.
type EndpointAccessUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointAccessUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEndpointAccessUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEndpointAccessUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEndpointAccessUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEndpointAccessUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEndpointAccessUpdateOK creates a EndpointAccessUpdateOK with default headers values
func NewEndpointAccessUpdateOK() *EndpointAccessUpdateOK {
	return &EndpointAccessUpdateOK{}
}

/*EndpointAccessUpdateOK handles this case with default header values.

Success
*/
type EndpointAccessUpdateOK struct {
	Payload *models.Endpoint
}

func (o *EndpointAccessUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/access][%d] endpointAccessUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointAccessUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Endpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointAccessUpdateBadRequest creates a EndpointAccessUpdateBadRequest with default headers values
func NewEndpointAccessUpdateBadRequest() *EndpointAccessUpdateBadRequest {
	return &EndpointAccessUpdateBadRequest{}
}

/*EndpointAccessUpdateBadRequest handles this case with default header values.

Invalid request
*/
type EndpointAccessUpdateBadRequest struct {
	Payload *models.GenericError
}

func (o *EndpointAccessUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/access][%d] endpointAccessUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *EndpointAccessUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointAccessUpdateNotFound creates a EndpointAccessUpdateNotFound with default headers values
func NewEndpointAccessUpdateNotFound() *EndpointAccessUpdateNotFound {
	return &EndpointAccessUpdateNotFound{}
}

/*EndpointAccessUpdateNotFound handles this case with default header values.

Endpoint not found
*/
type EndpointAccessUpdateNotFound struct {
	Payload *models.GenericError
}

func (o *EndpointAccessUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/access][%d] endpointAccessUpdateNotFound  %+v", 404, o.Payload)
}

func (o *EndpointAccessUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointAccessUpdateInternalServerError creates a EndpointAccessUpdateInternalServerError with default headers values
func NewEndpointAccessUpdateInternalServerError() *EndpointAccessUpdateInternalServerError {
	return &EndpointAccessUpdateInternalServerError{}
}

/*EndpointAccessUpdateInternalServerError handles this case with default header values.

Server error
*/
type EndpointAccessUpdateInternalServerError struct {
	Payload *models.GenericError
}

func (o *EndpointAccessUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/access][%d] endpointAccessUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *EndpointAccessUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
