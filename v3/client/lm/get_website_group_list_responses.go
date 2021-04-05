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

// GetWebsiteGroupListReader is a Reader for the GetWebsiteGroupList structure.
type GetWebsiteGroupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteGroupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteGroupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWebsiteGroupListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteGroupListOK creates a GetWebsiteGroupListOK with default headers values
func NewGetWebsiteGroupListOK() *GetWebsiteGroupListOK {
	return &GetWebsiteGroupListOK{}
}

/* GetWebsiteGroupListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteGroupListOK struct {
	Payload *models.WebsiteGroupPaginationResponse
}

func (o *GetWebsiteGroupListOK) Error() string {
	return fmt.Sprintf("[GET /website/groups][%d] getWebsiteGroupListOK  %+v", 200, o.Payload)
}
func (o *GetWebsiteGroupListOK) GetPayload() *models.WebsiteGroupPaginationResponse {
	return o.Payload
}

func (o *GetWebsiteGroupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteGroupPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteGroupListDefault creates a GetWebsiteGroupListDefault with default headers values
func NewGetWebsiteGroupListDefault(code int) *GetWebsiteGroupListDefault {
	return &GetWebsiteGroupListDefault{
		_statusCode: code,
	}
}

/* GetWebsiteGroupListDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteGroupListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website group list default response
func (o *GetWebsiteGroupListDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteGroupListDefault) Error() string {
	return fmt.Sprintf("[GET /website/groups][%d] getWebsiteGroupList default  %+v", o._statusCode, o.Payload)
}
func (o *GetWebsiteGroupListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteGroupListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
