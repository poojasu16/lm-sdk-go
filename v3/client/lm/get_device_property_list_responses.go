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

// GetDevicePropertyListReader is a Reader for the GetDevicePropertyList structure.
type GetDevicePropertyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicePropertyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicePropertyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDevicePropertyListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDevicePropertyListOK creates a GetDevicePropertyListOK with default headers values
func NewGetDevicePropertyListOK() *GetDevicePropertyListOK {
	return &GetDevicePropertyListOK{}
}

/* GetDevicePropertyListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDevicePropertyListOK struct {
	Payload *models.PropertyPaginationResponse
}

func (o *GetDevicePropertyListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/properties][%d] getDevicePropertyListOK  %+v", 200, o.Payload)
}
func (o *GetDevicePropertyListOK) GetPayload() *models.PropertyPaginationResponse {
	return o.Payload
}

func (o *GetDevicePropertyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicePropertyListDefault creates a GetDevicePropertyListDefault with default headers values
func NewGetDevicePropertyListDefault(code int) *GetDevicePropertyListDefault {
	return &GetDevicePropertyListDefault{
		_statusCode: code,
	}
}

/* GetDevicePropertyListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDevicePropertyListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device property list default response
func (o *GetDevicePropertyListDefault) Code() int {
	return o._statusCode
}

func (o *GetDevicePropertyListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/properties][%d] getDevicePropertyList default  %+v", o._statusCode, o.Payload)
}
func (o *GetDevicePropertyListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDevicePropertyListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
