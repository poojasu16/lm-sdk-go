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

// UpdateDeviceDatasourceInstanceByIDReader is a Reader for the UpdateDeviceDatasourceInstanceByID structure.
type UpdateDeviceDatasourceInstanceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceDatasourceInstanceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceDatasourceInstanceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDeviceDatasourceInstanceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceDatasourceInstanceByIDOK creates a UpdateDeviceDatasourceInstanceByIDOK with default headers values
func NewUpdateDeviceDatasourceInstanceByIDOK() *UpdateDeviceDatasourceInstanceByIDOK {
	return &UpdateDeviceDatasourceInstanceByIDOK{}
}

/* UpdateDeviceDatasourceInstanceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDeviceDatasourceInstanceByIDOK struct {
	Payload *models.DeviceDataSourceInstance
}

func (o *UpdateDeviceDatasourceInstanceByIDOK) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}][%d] updateDeviceDatasourceInstanceByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateDeviceDatasourceInstanceByIDOK) GetPayload() *models.DeviceDataSourceInstance {
	return o.Payload
}

func (o *UpdateDeviceDatasourceInstanceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceDatasourceInstanceByIDDefault creates a UpdateDeviceDatasourceInstanceByIDDefault with default headers values
func NewUpdateDeviceDatasourceInstanceByIDDefault(code int) *UpdateDeviceDatasourceInstanceByIDDefault {
	return &UpdateDeviceDatasourceInstanceByIDDefault{
		_statusCode: code,
	}
}

/* UpdateDeviceDatasourceInstanceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDeviceDatasourceInstanceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device datasource instance by Id default response
func (o *UpdateDeviceDatasourceInstanceByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceDatasourceInstanceByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}][%d] updateDeviceDatasourceInstanceById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDeviceDatasourceInstanceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceDatasourceInstanceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
