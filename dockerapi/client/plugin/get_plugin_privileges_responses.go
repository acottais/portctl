// Code generated by go-swagger; DO NOT EDIT.

package plugin

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

// GetPluginPrivilegesReader is a Reader for the GetPluginPrivileges structure.
type GetPluginPrivilegesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginPrivilegesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPluginPrivilegesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetPluginPrivilegesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPluginPrivilegesOK creates a GetPluginPrivilegesOK with default headers values
func NewGetPluginPrivilegesOK() *GetPluginPrivilegesOK {
	return &GetPluginPrivilegesOK{}
}

/*GetPluginPrivilegesOK handles this case with default header values.

no error
*/
type GetPluginPrivilegesOK struct {
	Payload []*GetPluginPrivilegesOKBodyItems0
}

func (o *GetPluginPrivilegesOK) Error() string {
	return fmt.Sprintf("[GET /plugins/privileges][%d] getPluginPrivilegesOK  %+v", 200, o.Payload)
}

func (o *GetPluginPrivilegesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginPrivilegesInternalServerError creates a GetPluginPrivilegesInternalServerError with default headers values
func NewGetPluginPrivilegesInternalServerError() *GetPluginPrivilegesInternalServerError {
	return &GetPluginPrivilegesInternalServerError{}
}

/*GetPluginPrivilegesInternalServerError handles this case with default header values.

server error
*/
type GetPluginPrivilegesInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetPluginPrivilegesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /plugins/privileges][%d] getPluginPrivilegesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPluginPrivilegesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPluginPrivilegesOKBodyItems0 PluginPrivilegeItem
//
// Describes a permission the user has to accept upon installing the plugin.
swagger:model GetPluginPrivilegesOKBodyItems0
*/
type GetPluginPrivilegesOKBodyItems0 struct {

	// description
	Description string `json:"Description,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// value
	Value []string `json:"Value"`
}

// Validate validates this get plugin privileges o k body items0
func (o *GetPluginPrivilegesOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPluginPrivilegesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPluginPrivilegesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetPluginPrivilegesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
