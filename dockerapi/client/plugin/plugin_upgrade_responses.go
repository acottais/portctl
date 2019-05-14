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

// PluginUpgradeReader is a Reader for the PluginUpgrade structure.
type PluginUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPluginUpgradeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPluginUpgradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPluginUpgradeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPluginUpgradeNoContent creates a PluginUpgradeNoContent with default headers values
func NewPluginUpgradeNoContent() *PluginUpgradeNoContent {
	return &PluginUpgradeNoContent{}
}

/*PluginUpgradeNoContent handles this case with default header values.

no error
*/
type PluginUpgradeNoContent struct {
}

func (o *PluginUpgradeNoContent) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/upgrade][%d] pluginUpgradeNoContent ", 204)
}

func (o *PluginUpgradeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginUpgradeNotFound creates a PluginUpgradeNotFound with default headers values
func NewPluginUpgradeNotFound() *PluginUpgradeNotFound {
	return &PluginUpgradeNotFound{}
}

/*PluginUpgradeNotFound handles this case with default header values.

plugin not installed
*/
type PluginUpgradeNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PluginUpgradeNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/upgrade][%d] pluginUpgradeNotFound  %+v", 404, o.Payload)
}

func (o *PluginUpgradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginUpgradeInternalServerError creates a PluginUpgradeInternalServerError with default headers values
func NewPluginUpgradeInternalServerError() *PluginUpgradeInternalServerError {
	return &PluginUpgradeInternalServerError{}
}

/*PluginUpgradeInternalServerError handles this case with default header values.

server error
*/
type PluginUpgradeInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginUpgradeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/upgrade][%d] pluginUpgradeInternalServerError  %+v", 500, o.Payload)
}

func (o *PluginUpgradeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PluginUpgradeParamsBodyItems0 Describes a permission accepted by the user upon installing the plugin.
swagger:model PluginUpgradeParamsBodyItems0
*/
type PluginUpgradeParamsBodyItems0 struct {

	// description
	Description string `json:"Description,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// value
	Value []string `json:"Value"`
}

// Validate validates this plugin upgrade params body items0
func (o *PluginUpgradeParamsBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PluginUpgradeParamsBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginUpgradeParamsBodyItems0) UnmarshalBinary(b []byte) error {
	var res PluginUpgradeParamsBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
