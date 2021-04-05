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

// GetSDTListReader is a Reader for the GetSDTList structure.
type GetSDTListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDTListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSDTListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSDTListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSDTListOK creates a GetSDTListOK with default headers values
func NewGetSDTListOK() *GetSDTListOK {
	return &GetSDTListOK{}
}

/* GetSDTListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSDTListOK struct {
	Payload *models.SDTPaginationResponse
}

func (o *GetSDTListOK) Error() string {
	return fmt.Sprintf("[GET /sdt/sdts][%d] getSdtListOK  %+v", 200, o.Payload)
}
func (o *GetSDTListOK) GetPayload() *models.SDTPaginationResponse {
	return o.Payload
}

func (o *GetSDTListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDTPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDTListDefault creates a GetSDTListDefault with default headers values
func NewGetSDTListDefault(code int) *GetSDTListDefault {
	return &GetSDTListDefault{
		_statusCode: code,
	}
}

/* GetSDTListDefault describes a response with status code -1, with default header values.

Error
*/
type GetSDTListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get SDT list default response
func (o *GetSDTListDefault) Code() int {
	return o._statusCode
}

func (o *GetSDTListDefault) Error() string {
	return fmt.Sprintf("[GET /sdt/sdts][%d] getSDTList default  %+v", o._statusCode, o.Payload)
}
func (o *GetSDTListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSDTListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
