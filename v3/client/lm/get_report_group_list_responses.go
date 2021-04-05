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

// GetReportGroupListReader is a Reader for the GetReportGroupList structure.
type GetReportGroupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportGroupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportGroupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetReportGroupListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportGroupListOK creates a GetReportGroupListOK with default headers values
func NewGetReportGroupListOK() *GetReportGroupListOK {
	return &GetReportGroupListOK{}
}

/* GetReportGroupListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetReportGroupListOK struct {
	Payload *models.ReportGroupPaginationResponse
}

func (o *GetReportGroupListOK) Error() string {
	return fmt.Sprintf("[GET /report/groups][%d] getReportGroupListOK  %+v", 200, o.Payload)
}
func (o *GetReportGroupListOK) GetPayload() *models.ReportGroupPaginationResponse {
	return o.Payload
}

func (o *GetReportGroupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportGroupPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportGroupListDefault creates a GetReportGroupListDefault with default headers values
func NewGetReportGroupListDefault(code int) *GetReportGroupListDefault {
	return &GetReportGroupListDefault{
		_statusCode: code,
	}
}

/* GetReportGroupListDefault describes a response with status code -1, with default header values.

Error
*/
type GetReportGroupListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get report group list default response
func (o *GetReportGroupListDefault) Code() int {
	return o._statusCode
}

func (o *GetReportGroupListDefault) Error() string {
	return fmt.Sprintf("[GET /report/groups][%d] getReportGroupList default  %+v", o._statusCode, o.Payload)
}
func (o *GetReportGroupListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetReportGroupListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
