// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// ConfigCreateReader is a Reader for the ConfigCreate structure.
type ConfigCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewConfigCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewConfigCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewConfigCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewConfigCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConfigCreateCreated creates a ConfigCreateCreated with default headers values
func NewConfigCreateCreated() *ConfigCreateCreated {
	return &ConfigCreateCreated{}
}

/*ConfigCreateCreated handles this case with default header values.

no error
*/
type ConfigCreateCreated struct {
	Payload *models.IDResponse
}

func (o *ConfigCreateCreated) Error() string {
	return fmt.Sprintf("[POST /configs/create][%d] configCreateCreated  %+v", 201, o.Payload)
}

func (o *ConfigCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigCreateConflict creates a ConfigCreateConflict with default headers values
func NewConfigCreateConflict() *ConfigCreateConflict {
	return &ConfigCreateConflict{}
}

/*ConfigCreateConflict handles this case with default header values.

name conflicts with an existing object
*/
type ConfigCreateConflict struct {
	Payload *models.ErrorResponse
}

func (o *ConfigCreateConflict) Error() string {
	return fmt.Sprintf("[POST /configs/create][%d] configCreateConflict  %+v", 409, o.Payload)
}

func (o *ConfigCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigCreateInternalServerError creates a ConfigCreateInternalServerError with default headers values
func NewConfigCreateInternalServerError() *ConfigCreateInternalServerError {
	return &ConfigCreateInternalServerError{}
}

/*ConfigCreateInternalServerError handles this case with default header values.

server error
*/
type ConfigCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ConfigCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /configs/create][%d] configCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigCreateServiceUnavailable creates a ConfigCreateServiceUnavailable with default headers values
func NewConfigCreateServiceUnavailable() *ConfigCreateServiceUnavailable {
	return &ConfigCreateServiceUnavailable{}
}

/*ConfigCreateServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type ConfigCreateServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ConfigCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /configs/create][%d] configCreateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConfigCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ConfigCreateBody config create body
swagger:model ConfigCreateBody
*/
type ConfigCreateBody struct {
	models.ConfigSpec

	ConfigCreateParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ConfigCreateBody) UnmarshalJSON(raw []byte) error {
	// ConfigCreateParamsBodyAO0
	var configCreateParamsBodyAO0 models.ConfigSpec
	if err := swag.ReadJSON(raw, &configCreateParamsBodyAO0); err != nil {
		return err
	}
	o.ConfigSpec = configCreateParamsBodyAO0

	// ConfigCreateParamsBodyAO1
	var configCreateParamsBodyAO1 ConfigCreateParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &configCreateParamsBodyAO1); err != nil {
		return err
	}
	o.ConfigCreateParamsBodyAllOf1 = configCreateParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ConfigCreateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	configCreateParamsBodyAO0, err := swag.WriteJSON(o.ConfigSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, configCreateParamsBodyAO0)

	configCreateParamsBodyAO1, err := swag.WriteJSON(o.ConfigCreateParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, configCreateParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this config create body
func (o *ConfigCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ConfigSpec
	if err := o.ConfigSpec.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ConfigCreateParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ConfigCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigCreateBody) UnmarshalBinary(b []byte) error {
	var res ConfigCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConfigCreateParamsBodyAllOf1 config create params body all of1
swagger:model ConfigCreateParamsBodyAllOf1
*/
type ConfigCreateParamsBodyAllOf1 interface{}