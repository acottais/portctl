// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// VolumeInspectReader is a Reader for the VolumeInspect structure.
type VolumeInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVolumeInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewVolumeInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewVolumeInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVolumeInspectOK creates a VolumeInspectOK with default headers values
func NewVolumeInspectOK() *VolumeInspectOK {
	return &VolumeInspectOK{}
}

/*VolumeInspectOK handles this case with default header values.

No error
*/
type VolumeInspectOK struct {
	Payload *models.Volume
}

func (o *VolumeInspectOK) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectOK  %+v", 200, o.Payload)
}

func (o *VolumeInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectNotFound creates a VolumeInspectNotFound with default headers values
func NewVolumeInspectNotFound() *VolumeInspectNotFound {
	return &VolumeInspectNotFound{}
}

/*VolumeInspectNotFound handles this case with default header values.

No such volume
*/
type VolumeInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *VolumeInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectInternalServerError creates a VolumeInspectInternalServerError with default headers values
func NewVolumeInspectInternalServerError() *VolumeInspectInternalServerError {
	return &VolumeInspectInternalServerError{}
}

/*VolumeInspectInternalServerError handles this case with default header values.

Server error
*/
type VolumeInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *VolumeInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
