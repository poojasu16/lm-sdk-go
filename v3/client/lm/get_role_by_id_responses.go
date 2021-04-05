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

// GetRoleByIDReader is a Reader for the GetRoleByID structure.
type GetRoleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRoleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleByIDOK creates a GetRoleByIDOK with default headers values
func NewGetRoleByIDOK() *GetRoleByIDOK {
	return &GetRoleByIDOK{}
}

/* GetRoleByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoleByIDOK struct {
	Payload *models.Role
}

func (o *GetRoleByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/roles/{id}][%d] getRoleByIdOK  %+v", 200, o.Payload)
}
func (o *GetRoleByIDOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *GetRoleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleByIDDefault creates a GetRoleByIDDefault with default headers values
func NewGetRoleByIDDefault(code int) *GetRoleByIDDefault {
	return &GetRoleByIDDefault{
		_statusCode: code,
	}
}

/* GetRoleByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetRoleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get role by Id default response
func (o *GetRoleByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRoleByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/roles/{id}][%d] getRoleById default  %+v", o._statusCode, o.Payload)
}
func (o *GetRoleByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRoleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}