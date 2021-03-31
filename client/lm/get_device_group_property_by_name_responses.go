// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// GetDeviceGroupPropertyByNameReader is a Reader for the GetDeviceGroupPropertyByName structure.
type GetDeviceGroupPropertyByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceGroupPropertyByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceGroupPropertyByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceGroupPropertyByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceGroupPropertyByNameOK creates a GetDeviceGroupPropertyByNameOK with default headers values
func NewGetDeviceGroupPropertyByNameOK() *GetDeviceGroupPropertyByNameOK {
	return &GetDeviceGroupPropertyByNameOK{}
}

/* GetDeviceGroupPropertyByNameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceGroupPropertyByNameOK struct {
	Payload *models.EntityProperty
}

func (o *GetDeviceGroupPropertyByNameOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{gid}/properties/{name}][%d] getDeviceGroupPropertyByNameOK  %+v", 200, o.Payload)
}
func (o *GetDeviceGroupPropertyByNameOK) GetPayload() *models.EntityProperty {
	return o.Payload
}

func (o *GetDeviceGroupPropertyByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceGroupPropertyByNameDefault creates a GetDeviceGroupPropertyByNameDefault with default headers values
func NewGetDeviceGroupPropertyByNameDefault(code int) *GetDeviceGroupPropertyByNameDefault {
	return &GetDeviceGroupPropertyByNameDefault{
		_statusCode: code,
	}
}

/* GetDeviceGroupPropertyByNameDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceGroupPropertyByNameDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device group property by name default response
func (o *GetDeviceGroupPropertyByNameDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceGroupPropertyByNameDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{gid}/properties/{name}][%d] getDeviceGroupPropertyByName default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceGroupPropertyByNameDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceGroupPropertyByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
