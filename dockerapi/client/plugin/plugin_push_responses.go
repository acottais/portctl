// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// PluginPushReader is a Reader for the PluginPush structure.
type PluginPushReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginPushReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPluginPushOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPluginPushNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPluginPushInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPluginPushOK creates a PluginPushOK with default headers values
func NewPluginPushOK() *PluginPushOK {
	return &PluginPushOK{}
}

/*PluginPushOK handles this case with default header values.

no error
*/
type PluginPushOK struct {
}

func (o *PluginPushOK) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/push][%d] pluginPushOK ", 200)
}

func (o *PluginPushOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginPushNotFound creates a PluginPushNotFound with default headers values
func NewPluginPushNotFound() *PluginPushNotFound {
	return &PluginPushNotFound{}
}

/*PluginPushNotFound handles this case with default header values.

plugin not installed
*/
type PluginPushNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PluginPushNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/push][%d] pluginPushNotFound  %+v", 404, o.Payload)
}

func (o *PluginPushNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginPushInternalServerError creates a PluginPushInternalServerError with default headers values
func NewPluginPushInternalServerError() *PluginPushInternalServerError {
	return &PluginPushInternalServerError{}
}

/*PluginPushInternalServerError handles this case with default header values.

server error
*/
type PluginPushInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginPushInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/push][%d] pluginPushInternalServerError  %+v", 500, o.Payload)
}

func (o *PluginPushInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
