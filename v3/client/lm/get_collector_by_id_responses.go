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

// GetCollectorByIDReader is a Reader for the GetCollectorByID structure.
type GetCollectorByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectorByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectorByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCollectorByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectorByIDOK creates a GetCollectorByIDOK with default headers values
func NewGetCollectorByIDOK() *GetCollectorByIDOK {
	return &GetCollectorByIDOK{}
}

/* GetCollectorByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCollectorByIDOK struct {
	Payload *models.Collector
}

func (o *GetCollectorByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{id}][%d] getCollectorByIdOK  %+v", 200, o.Payload)
}
func (o *GetCollectorByIDOK) GetPayload() *models.Collector {
	return o.Payload
}

func (o *GetCollectorByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectorByIDDefault creates a GetCollectorByIDDefault with default headers values
func NewGetCollectorByIDDefault(code int) *GetCollectorByIDDefault {
	return &GetCollectorByIDDefault{
		_statusCode: code,
	}
}

/* GetCollectorByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetCollectorByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get collector by Id default response
func (o *GetCollectorByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCollectorByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{id}][%d] getCollectorById default  %+v", o._statusCode, o.Payload)
}
func (o *GetCollectorByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCollectorByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
