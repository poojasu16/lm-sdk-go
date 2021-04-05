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

// AddCollectorReader is a Reader for the AddCollector structure.
type AddCollectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCollectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddCollectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddCollectorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddCollectorOK creates a AddCollectorOK with default headers values
func NewAddCollectorOK() *AddCollectorOK {
	return &AddCollectorOK{}
}

/* AddCollectorOK describes a response with status code 200, with default header values.

successful operation
*/
type AddCollectorOK struct {
	Payload *models.Collector
}

func (o *AddCollectorOK) Error() string {
	return fmt.Sprintf("[POST /setting/collector/collectors][%d] addCollectorOK  %+v", 200, o.Payload)
}
func (o *AddCollectorOK) GetPayload() *models.Collector {
	return o.Payload
}

func (o *AddCollectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCollectorDefault creates a AddCollectorDefault with default headers values
func NewAddCollectorDefault(code int) *AddCollectorDefault {
	return &AddCollectorDefault{
		_statusCode: code,
	}
}

/* AddCollectorDefault describes a response with status code -1, with default header values.

Error
*/
type AddCollectorDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add collector default response
func (o *AddCollectorDefault) Code() int {
	return o._statusCode
}

func (o *AddCollectorDefault) Error() string {
	return fmt.Sprintf("[POST /setting/collector/collectors][%d] addCollector default  %+v", o._statusCode, o.Payload)
}
func (o *AddCollectorDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddCollectorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
