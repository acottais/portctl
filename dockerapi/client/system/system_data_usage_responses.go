// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// SystemDataUsageReader is a Reader for the SystemDataUsage structure.
type SystemDataUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemDataUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSystemDataUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSystemDataUsageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemDataUsageOK creates a SystemDataUsageOK with default headers values
func NewSystemDataUsageOK() *SystemDataUsageOK {
	return &SystemDataUsageOK{}
}

/*SystemDataUsageOK handles this case with default header values.

no error
*/
type SystemDataUsageOK struct {
	Payload *SystemDataUsageOKBody
}

func (o *SystemDataUsageOK) Error() string {
	return fmt.Sprintf("[GET /system/df][%d] systemDataUsageOK  %+v", 200, o.Payload)
}

func (o *SystemDataUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SystemDataUsageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemDataUsageInternalServerError creates a SystemDataUsageInternalServerError with default headers values
func NewSystemDataUsageInternalServerError() *SystemDataUsageInternalServerError {
	return &SystemDataUsageInternalServerError{}
}

/*SystemDataUsageInternalServerError handles this case with default header values.

server error
*/
type SystemDataUsageInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SystemDataUsageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/df][%d] systemDataUsageInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemDataUsageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SystemDataUsageOKBody SystemDataUsageResponse
swagger:model SystemDataUsageOKBody
*/
type SystemDataUsageOKBody struct {

	// build cache
	BuildCache []*models.BuildCache `json:"BuildCache"`

	// containers
	Containers []models.ContainerSummary `json:"Containers"`

	// images
	Images []*models.ImageSummary `json:"Images"`

	// layers size
	LayersSize int64 `json:"LayersSize,omitempty"`

	// volumes
	Volumes []*models.Volume `json:"Volumes"`
}

// Validate validates this system data usage o k body
func (o *SystemDataUsageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBuildCache(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SystemDataUsageOKBody) validateBuildCache(formats strfmt.Registry) error {

	if swag.IsZero(o.BuildCache) { // not required
		return nil
	}

	for i := 0; i < len(o.BuildCache); i++ {
		if swag.IsZero(o.BuildCache[i]) { // not required
			continue
		}

		if o.BuildCache[i] != nil {
			if err := o.BuildCache[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemDataUsageOK" + "." + "BuildCache" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SystemDataUsageOKBody) validateContainers(formats strfmt.Registry) error {

	if swag.IsZero(o.Containers) { // not required
		return nil
	}

	for i := 0; i < len(o.Containers); i++ {

		if err := o.Containers[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemDataUsageOK" + "." + "Containers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (o *SystemDataUsageOKBody) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(o.Images) { // not required
		return nil
	}

	for i := 0; i < len(o.Images); i++ {
		if swag.IsZero(o.Images[i]) { // not required
			continue
		}

		if o.Images[i] != nil {
			if err := o.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemDataUsageOK" + "." + "Images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SystemDataUsageOKBody) validateVolumes(formats strfmt.Registry) error {

	if swag.IsZero(o.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(o.Volumes); i++ {
		if swag.IsZero(o.Volumes[i]) { // not required
			continue
		}

		if o.Volumes[i] != nil {
			if err := o.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemDataUsageOK" + "." + "Volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SystemDataUsageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemDataUsageOKBody) UnmarshalBinary(b []byte) error {
	var res SystemDataUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}