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

// GetSDTByIDReader is a Reader for the GetSDTByID structure.
type GetSDTByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDTByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSDTByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSDTByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSDTByIDOK creates a GetSDTByIDOK with default headers values
func NewGetSDTByIDOK() *GetSDTByIDOK {
	return &GetSDTByIDOK{}
}

/* GetSDTByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSDTByIDOK struct {
	Payload models.SDT
}

func (o *GetSDTByIDOK) Error() string {
	return fmt.Sprintf("[GET /sdt/sdts/{id}][%d] getSdtByIdOK  %+v", 200, o.Payload)
}
func (o *GetSDTByIDOK) GetPayload() models.SDT {
	return o.Payload
}

func (o *GetSDTByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalSDT(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetSDTByIDDefault creates a GetSDTByIDDefault with default headers values
func NewGetSDTByIDDefault(code int) *GetSDTByIDDefault {
	return &GetSDTByIDDefault{
		_statusCode: code,
	}
}

/* GetSDTByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetSDTByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get SDT by Id default response
func (o *GetSDTByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSDTByIDDefault) Error() string {
	return fmt.Sprintf("[GET /sdt/sdts/{id}][%d] getSDTById default  %+v", o._statusCode, o.Payload)
}
func (o *GetSDTByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSDTByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
