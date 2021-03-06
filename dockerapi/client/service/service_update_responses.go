// Code generated by go-swagger; DO NOT EDIT.

package service

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

// ServiceUpdateReader is a Reader for the ServiceUpdate structure.
type ServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewServiceUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewServiceUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewServiceUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceUpdateOK creates a ServiceUpdateOK with default headers values
func NewServiceUpdateOK() *ServiceUpdateOK {
	return &ServiceUpdateOK{}
}

/*ServiceUpdateOK handles this case with default header values.

no error
*/
type ServiceUpdateOK struct {
	Payload *models.ServiceUpdateResponse
}

func (o *ServiceUpdateOK) Error() string {
	return fmt.Sprintf("[POST /services/{id}/update][%d] serviceUpdateOK  %+v", 200, o.Payload)
}

func (o *ServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceUpdateBadRequest creates a ServiceUpdateBadRequest with default headers values
func NewServiceUpdateBadRequest() *ServiceUpdateBadRequest {
	return &ServiceUpdateBadRequest{}
}

/*ServiceUpdateBadRequest handles this case with default header values.

bad parameter
*/
type ServiceUpdateBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ServiceUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/{id}/update][%d] serviceUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceUpdateNotFound creates a ServiceUpdateNotFound with default headers values
func NewServiceUpdateNotFound() *ServiceUpdateNotFound {
	return &ServiceUpdateNotFound{}
}

/*ServiceUpdateNotFound handles this case with default header values.

no such service
*/
type ServiceUpdateNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ServiceUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /services/{id}/update][%d] serviceUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ServiceUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceUpdateInternalServerError creates a ServiceUpdateInternalServerError with default headers values
func NewServiceUpdateInternalServerError() *ServiceUpdateInternalServerError {
	return &ServiceUpdateInternalServerError{}
}

/*ServiceUpdateInternalServerError handles this case with default header values.

server error
*/
type ServiceUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ServiceUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /services/{id}/update][%d] serviceUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceUpdateServiceUnavailable creates a ServiceUpdateServiceUnavailable with default headers values
func NewServiceUpdateServiceUnavailable() *ServiceUpdateServiceUnavailable {
	return &ServiceUpdateServiceUnavailable{}
}

/*ServiceUpdateServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type ServiceUpdateServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ServiceUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /services/{id}/update][%d] serviceUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ServiceUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ServiceUpdateBody service update body
swagger:model ServiceUpdateBody
*/
type ServiceUpdateBody struct {
	models.ServiceSpec

	ServiceUpdateParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ServiceUpdateBody) UnmarshalJSON(raw []byte) error {
	// ServiceUpdateParamsBodyAO0
	var serviceUpdateParamsBodyAO0 models.ServiceSpec
	if err := swag.ReadJSON(raw, &serviceUpdateParamsBodyAO0); err != nil {
		return err
	}
	o.ServiceSpec = serviceUpdateParamsBodyAO0

	// ServiceUpdateParamsBodyAO1
	var serviceUpdateParamsBodyAO1 ServiceUpdateParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &serviceUpdateParamsBodyAO1); err != nil {
		return err
	}
	o.ServiceUpdateParamsBodyAllOf1 = serviceUpdateParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ServiceUpdateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	serviceUpdateParamsBodyAO0, err := swag.WriteJSON(o.ServiceSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, serviceUpdateParamsBodyAO0)

	serviceUpdateParamsBodyAO1, err := swag.WriteJSON(o.ServiceUpdateParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, serviceUpdateParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this service update body
func (o *ServiceUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ServiceSpec
	if err := o.ServiceSpec.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ServiceUpdateParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceUpdateBody) UnmarshalBinary(b []byte) error {
	var res ServiceUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ServiceUpdateParamsBodyAllOf1 service update params body all of1
swagger:model ServiceUpdateParamsBodyAllOf1
*/
type ServiceUpdateParamsBodyAllOf1 interface{}
