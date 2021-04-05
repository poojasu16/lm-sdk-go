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

// ImportDNSMappingReader is a Reader for the ImportDNSMapping structure.
type ImportDNSMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportDNSMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportDNSMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewImportDNSMappingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImportDNSMappingOK creates a ImportDNSMappingOK with default headers values
func NewImportDNSMappingOK() *ImportDNSMappingOK {
	return &ImportDNSMappingOK{}
}

/* ImportDNSMappingOK describes a response with status code 200, with default header values.

successful operation
*/
type ImportDNSMappingOK struct {
	Payload interface{}
}

func (o *ImportDNSMappingOK) Error() string {
	return fmt.Sprintf("[POST /setting/dnsmappings][%d] importDnsMappingOK  %+v", 200, o.Payload)
}
func (o *ImportDNSMappingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ImportDNSMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDNSMappingDefault creates a ImportDNSMappingDefault with default headers values
func NewImportDNSMappingDefault(code int) *ImportDNSMappingDefault {
	return &ImportDNSMappingDefault{
		_statusCode: code,
	}
}

/* ImportDNSMappingDefault describes a response with status code -1, with default header values.

Error
*/
type ImportDNSMappingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the import DNS mapping default response
func (o *ImportDNSMappingDefault) Code() int {
	return o._statusCode
}

func (o *ImportDNSMappingDefault) Error() string {
	return fmt.Sprintf("[POST /setting/dnsmappings][%d] importDNSMapping default  %+v", o._statusCode, o.Payload)
}
func (o *ImportDNSMappingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImportDNSMappingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
