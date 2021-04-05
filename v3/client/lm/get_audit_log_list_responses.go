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

// GetAuditLogListReader is a Reader for the GetAuditLogList structure.
type GetAuditLogListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditLogListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditLogListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditLogListOK creates a GetAuditLogListOK with default headers values
func NewGetAuditLogListOK() *GetAuditLogListOK {
	return &GetAuditLogListOK{}
}

/* GetAuditLogListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAuditLogListOK struct {
	Payload *models.AccessLogPaginationResponse
}

func (o *GetAuditLogListOK) Error() string {
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogListOK  %+v", 200, o.Payload)
}
func (o *GetAuditLogListOK) GetPayload() *models.AccessLogPaginationResponse {
	return o.Payload
}

func (o *GetAuditLogListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessLogPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogListDefault creates a GetAuditLogListDefault with default headers values
func NewGetAuditLogListDefault(code int) *GetAuditLogListDefault {
	return &GetAuditLogListDefault{
		_statusCode: code,
	}
}

/* GetAuditLogListDefault describes a response with status code -1, with default header values.

Error
*/
type GetAuditLogListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get audit log list default response
func (o *GetAuditLogListDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditLogListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogList default  %+v", o._statusCode, o.Payload)
}
func (o *GetAuditLogListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAuditLogListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
