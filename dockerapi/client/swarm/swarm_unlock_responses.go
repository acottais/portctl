// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// SwarmUnlockReader is a Reader for the SwarmUnlock structure.
type SwarmUnlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwarmUnlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSwarmUnlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSwarmUnlockInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewSwarmUnlockServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSwarmUnlockOK creates a SwarmUnlockOK with default headers values
func NewSwarmUnlockOK() *SwarmUnlockOK {
	return &SwarmUnlockOK{}
}

/*SwarmUnlockOK handles this case with default header values.

no error
*/
type SwarmUnlockOK struct {
}

func (o *SwarmUnlockOK) Error() string {
	return fmt.Sprintf("[POST /swarm/unlock][%d] swarmUnlockOK ", 200)
}

func (o *SwarmUnlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSwarmUnlockInternalServerError creates a SwarmUnlockInternalServerError with default headers values
func NewSwarmUnlockInternalServerError() *SwarmUnlockInternalServerError {
	return &SwarmUnlockInternalServerError{}
}

/*SwarmUnlockInternalServerError handles this case with default header values.

server error
*/
type SwarmUnlockInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SwarmUnlockInternalServerError) Error() string {
	return fmt.Sprintf("[POST /swarm/unlock][%d] swarmUnlockInternalServerError  %+v", 500, o.Payload)
}

func (o *SwarmUnlockInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwarmUnlockServiceUnavailable creates a SwarmUnlockServiceUnavailable with default headers values
func NewSwarmUnlockServiceUnavailable() *SwarmUnlockServiceUnavailable {
	return &SwarmUnlockServiceUnavailable{}
}

/*SwarmUnlockServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type SwarmUnlockServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *SwarmUnlockServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /swarm/unlock][%d] swarmUnlockServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SwarmUnlockServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SwarmUnlockBody swarm unlock body
swagger:model SwarmUnlockBody
*/
type SwarmUnlockBody struct {

	// The swarm's unlock key.
	UnlockKey string `json:"UnlockKey,omitempty"`
}

// Validate validates this swarm unlock body
func (o *SwarmUnlockBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SwarmUnlockBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwarmUnlockBody) UnmarshalBinary(b []byte) error {
	var res SwarmUnlockBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
