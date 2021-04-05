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

// GetAlertListByDeviceGroupIDReader is a Reader for the GetAlertListByDeviceGroupID structure.
type GetAlertListByDeviceGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertListByDeviceGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertListByDeviceGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAlertListByDeviceGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertListByDeviceGroupIDOK creates a GetAlertListByDeviceGroupIDOK with default headers values
func NewGetAlertListByDeviceGroupIDOK() *GetAlertListByDeviceGroupIDOK {
	return &GetAlertListByDeviceGroupIDOK{}
}

/* GetAlertListByDeviceGroupIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertListByDeviceGroupIDOK struct {
	Payload *models.AlertPaginationResponse
}

func (o *GetAlertListByDeviceGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{id}/alerts][%d] getAlertListByDeviceGroupIdOK  %+v", 200, o.Payload)
}
func (o *GetAlertListByDeviceGroupIDOK) GetPayload() *models.AlertPaginationResponse {
	return o.Payload
}

func (o *GetAlertListByDeviceGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertListByDeviceGroupIDDefault creates a GetAlertListByDeviceGroupIDDefault with default headers values
func NewGetAlertListByDeviceGroupIDDefault(code int) *GetAlertListByDeviceGroupIDDefault {
	return &GetAlertListByDeviceGroupIDDefault{
		_statusCode: code,
	}
}

/* GetAlertListByDeviceGroupIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertListByDeviceGroupIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alert list by device group Id default response
func (o *GetAlertListByDeviceGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertListByDeviceGroupIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{id}/alerts][%d] getAlertListByDeviceGroupId default  %+v", o._statusCode, o.Payload)
}
func (o *GetAlertListByDeviceGroupIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertListByDeviceGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}