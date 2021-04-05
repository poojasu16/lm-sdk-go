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

// GetAlertListByDeviceIDReader is a Reader for the GetAlertListByDeviceID structure.
type GetAlertListByDeviceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertListByDeviceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertListByDeviceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAlertListByDeviceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertListByDeviceIDOK creates a GetAlertListByDeviceIDOK with default headers values
func NewGetAlertListByDeviceIDOK() *GetAlertListByDeviceIDOK {
	return &GetAlertListByDeviceIDOK{}
}

/* GetAlertListByDeviceIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertListByDeviceIDOK struct {
	Payload *models.AlertPaginationResponse
}

func (o *GetAlertListByDeviceIDOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/alerts][%d] getAlertListByDeviceIdOK  %+v", 200, o.Payload)
}
func (o *GetAlertListByDeviceIDOK) GetPayload() *models.AlertPaginationResponse {
	return o.Payload
}

func (o *GetAlertListByDeviceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertListByDeviceIDDefault creates a GetAlertListByDeviceIDDefault with default headers values
func NewGetAlertListByDeviceIDDefault(code int) *GetAlertListByDeviceIDDefault {
	return &GetAlertListByDeviceIDDefault{
		_statusCode: code,
	}
}

/* GetAlertListByDeviceIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertListByDeviceIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alert list by device Id default response
func (o *GetAlertListByDeviceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertListByDeviceIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/alerts][%d] getAlertListByDeviceId default  %+v", o._statusCode, o.Payload)
}
func (o *GetAlertListByDeviceIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertListByDeviceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
