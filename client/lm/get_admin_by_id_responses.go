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

// GetAdminByIDReader is a Reader for the GetAdminByID structure.
type GetAdminByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAdminByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdminByIDOK creates a GetAdminByIDOK with default headers values
func NewGetAdminByIDOK() *GetAdminByIDOK {
	return &GetAdminByIDOK{}
}

/* GetAdminByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAdminByIDOK struct {
	Payload *models.Admin
}

func (o *GetAdminByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminByIdOK  %+v", 200, o.Payload)
}
func (o *GetAdminByIDOK) GetPayload() *models.Admin {
	return o.Payload
}

func (o *GetAdminByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Admin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminByIDDefault creates a GetAdminByIDDefault with default headers values
func NewGetAdminByIDDefault(code int) *GetAdminByIDDefault {
	return &GetAdminByIDDefault{
		_statusCode: code,
	}
}

/* GetAdminByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAdminByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get admin by Id default response
func (o *GetAdminByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAdminByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminById default  %+v", o._statusCode, o.Payload)
}
func (o *GetAdminByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAdminByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
