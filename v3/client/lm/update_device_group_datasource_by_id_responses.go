// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// UpdateDeviceGroupDatasourceByIDReader is a Reader for the UpdateDeviceGroupDatasourceByID structure.
type UpdateDeviceGroupDatasourceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceGroupDatasourceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceGroupDatasourceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDeviceGroupDatasourceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceGroupDatasourceByIDOK creates a UpdateDeviceGroupDatasourceByIDOK with default headers values
func NewUpdateDeviceGroupDatasourceByIDOK() *UpdateDeviceGroupDatasourceByIDOK {
	return &UpdateDeviceGroupDatasourceByIDOK{}
}

/* UpdateDeviceGroupDatasourceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDeviceGroupDatasourceByIDOK struct {
	Payload *models.DeviceGroupDataSource
}

func (o *UpdateDeviceGroupDatasourceByIDOK) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{deviceGroupId}/datasources/{id}][%d] updateDeviceGroupDatasourceByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateDeviceGroupDatasourceByIDOK) GetPayload() *models.DeviceGroupDataSource {
	return o.Payload
}

func (o *UpdateDeviceGroupDatasourceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupDataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceGroupDatasourceByIDDefault creates a UpdateDeviceGroupDatasourceByIDDefault with default headers values
func NewUpdateDeviceGroupDatasourceByIDDefault(code int) *UpdateDeviceGroupDatasourceByIDDefault {
	return &UpdateDeviceGroupDatasourceByIDDefault{
		_statusCode: code,
	}
}

/* UpdateDeviceGroupDatasourceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDeviceGroupDatasourceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device group datasource by Id default response
func (o *UpdateDeviceGroupDatasourceByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceGroupDatasourceByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{deviceGroupId}/datasources/{id}][%d] updateDeviceGroupDatasourceById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDeviceGroupDatasourceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceGroupDatasourceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
