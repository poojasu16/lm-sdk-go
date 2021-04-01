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

// AddAdminReader is a Reader for the AddAdmin structure.
type AddAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAdminOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddAdminDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAdminOK creates a AddAdminOK with default headers values
func NewAddAdminOK() *AddAdminOK {
	return &AddAdminOK{}
}

/* AddAdminOK describes a response with status code 200, with default header values.

successful operation
*/
type AddAdminOK struct {
	Payload *models.Admin
}

func (o *AddAdminOK) Error() string {
	return fmt.Sprintf("[POST /setting/admins][%d] addAdminOK  %+v", 200, o.Payload)
}
func (o *AddAdminOK) GetPayload() *models.Admin {
	return o.Payload
}

func (o *AddAdminOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Admin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAdminDefault creates a AddAdminDefault with default headers values
func NewAddAdminDefault(code int) *AddAdminDefault {
	return &AddAdminDefault{
		_statusCode: code,
	}
}

/* AddAdminDefault describes a response with status code -1, with default header values.

Error
*/
type AddAdminDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add admin default response
func (o *AddAdminDefault) Code() int {
	return o._statusCode
}

func (o *AddAdminDefault) Error() string {
	return fmt.Sprintf("[POST /setting/admins][%d] addAdmin default  %+v", o._statusCode, o.Payload)
}
func (o *AddAdminDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddAdminDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
