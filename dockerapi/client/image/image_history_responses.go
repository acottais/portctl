// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acottais/portctl/dockerapi/models"
)

// ImageHistoryReader is a Reader for the ImageHistory structure.
type ImageHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewImageHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewImageHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageHistoryOK creates a ImageHistoryOK with default headers values
func NewImageHistoryOK() *ImageHistoryOK {
	return &ImageHistoryOK{}
}

/*ImageHistoryOK handles this case with default header values.

List of image layers
*/
type ImageHistoryOK struct {
	Payload []*HistoryResponseItem
}

func (o *ImageHistoryOK) Error() string {
	return fmt.Sprintf("[GET /images/{name}/history][%d] imageHistoryOK  %+v", 200, o.Payload)
}

func (o *ImageHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageHistoryNotFound creates a ImageHistoryNotFound with default headers values
func NewImageHistoryNotFound() *ImageHistoryNotFound {
	return &ImageHistoryNotFound{}
}

/*ImageHistoryNotFound handles this case with default header values.

No such image
*/
type ImageHistoryNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ImageHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /images/{name}/history][%d] imageHistoryNotFound  %+v", 404, o.Payload)
}

func (o *ImageHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageHistoryInternalServerError creates a ImageHistoryInternalServerError with default headers values
func NewImageHistoryInternalServerError() *ImageHistoryInternalServerError {
	return &ImageHistoryInternalServerError{}
}

/*ImageHistoryInternalServerError handles this case with default header values.

Server error
*/
type ImageHistoryInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImageHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{name}/history][%d] imageHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*HistoryResponseItem HistoryResponseItem
//
// individual image layer information in response to ImageHistory operation
swagger:model HistoryResponseItem
*/
type HistoryResponseItem struct {

	// comment
	// Required: true
	Comment string `json:"Comment"`

	// created
	// Required: true
	Created int64 `json:"Created"`

	// created by
	// Required: true
	CreatedBy string `json:"CreatedBy"`

	// Id
	// Required: true
	ID string `json:"Id"`

	// size
	// Required: true
	Size int64 `json:"Size"`

	// tags
	// Required: true
	Tags []string `json:"Tags"`
}

// Validate validates this history response item
func (o *HistoryResponseItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HistoryResponseItem) validateComment(formats strfmt.Registry) error {

	if err := validate.RequiredString("Comment", "body", string(o.Comment)); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("Created", "body", int64(o.Created)); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.RequiredString("CreatedBy", "body", string(o.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("Id", "body", string(o.ID)); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", int64(o.Size)); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("Tags", "body", o.Tags); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *HistoryResponseItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HistoryResponseItem) UnmarshalBinary(b []byte) error {
	var res HistoryResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
