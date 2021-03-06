// Code generated by go-swagger; DO NOT EDIT.

package container

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

// ContainerUpdateReader is a Reader for the ContainerUpdate structure.
type ContainerUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerUpdateOK creates a ContainerUpdateOK with default headers values
func NewContainerUpdateOK() *ContainerUpdateOK {
	return &ContainerUpdateOK{}
}

/*ContainerUpdateOK handles this case with default header values.

The container has been updated.
*/
type ContainerUpdateOK struct {
	Payload *ContainerUpdateOKBody
}

func (o *ContainerUpdateOK) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/update][%d] containerUpdateOK  %+v", 200, o.Payload)
}

func (o *ContainerUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerUpdateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerUpdateNotFound creates a ContainerUpdateNotFound with default headers values
func NewContainerUpdateNotFound() *ContainerUpdateNotFound {
	return &ContainerUpdateNotFound{}
}

/*ContainerUpdateNotFound handles this case with default header values.

no such container
*/
type ContainerUpdateNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/update][%d] containerUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ContainerUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerUpdateInternalServerError creates a ContainerUpdateInternalServerError with default headers values
func NewContainerUpdateInternalServerError() *ContainerUpdateInternalServerError {
	return &ContainerUpdateInternalServerError{}
}

/*ContainerUpdateInternalServerError handles this case with default header values.

server error
*/
type ContainerUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/update][%d] containerUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerUpdateBody container update body
swagger:model ContainerUpdateBody
*/
type ContainerUpdateBody struct {
	models.Resources

	// restart policy
	RestartPolicy *models.RestartPolicy `json:"RestartPolicy,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ContainerUpdateBody) UnmarshalJSON(raw []byte) error {
	// ContainerUpdateParamsBodyAO0
	var containerUpdateParamsBodyAO0 models.Resources
	if err := swag.ReadJSON(raw, &containerUpdateParamsBodyAO0); err != nil {
		return err
	}
	o.Resources = containerUpdateParamsBodyAO0

	// ContainerUpdateParamsBodyAO1
	var dataContainerUpdateParamsBodyAO1 struct {
		RestartPolicy *models.RestartPolicy `json:"RestartPolicy,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataContainerUpdateParamsBodyAO1); err != nil {
		return err
	}

	o.RestartPolicy = dataContainerUpdateParamsBodyAO1.RestartPolicy

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ContainerUpdateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	containerUpdateParamsBodyAO0, err := swag.WriteJSON(o.Resources)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, containerUpdateParamsBodyAO0)

	var dataContainerUpdateParamsBodyAO1 struct {
		RestartPolicy *models.RestartPolicy `json:"RestartPolicy,omitempty"`
	}

	dataContainerUpdateParamsBodyAO1.RestartPolicy = o.RestartPolicy

	jsonDataContainerUpdateParamsBodyAO1, errContainerUpdateParamsBodyAO1 := swag.WriteJSON(dataContainerUpdateParamsBodyAO1)
	if errContainerUpdateParamsBodyAO1 != nil {
		return nil, errContainerUpdateParamsBodyAO1
	}
	_parts = append(_parts, jsonDataContainerUpdateParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this container update body
func (o *ContainerUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Resources
	if err := o.Resources.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRestartPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerUpdateBody) validateRestartPolicy(formats strfmt.Registry) error {

	if swag.IsZero(o.RestartPolicy) { // not required
		return nil
	}

	if o.RestartPolicy != nil {
		if err := o.RestartPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update" + "." + "RestartPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ContainerUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerUpdateBody) UnmarshalBinary(b []byte) error {
	var res ContainerUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerUpdateOKBody ContainerUpdateResponse
//
// OK response to ContainerUpdate operation
swagger:model ContainerUpdateOKBody
*/
type ContainerUpdateOKBody struct {

	// warnings
	Warnings []string `json:"Warnings"`
}

// Validate validates this container update o k body
func (o *ContainerUpdateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
