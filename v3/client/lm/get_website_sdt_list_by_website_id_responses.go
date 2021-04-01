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

// GetWebsiteSDTListByWebsiteIDReader is a Reader for the GetWebsiteSDTListByWebsiteID structure.
type GetWebsiteSDTListByWebsiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteSDTListByWebsiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteSDTListByWebsiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWebsiteSDTListByWebsiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteSDTListByWebsiteIDOK creates a GetWebsiteSDTListByWebsiteIDOK with default headers values
func NewGetWebsiteSDTListByWebsiteIDOK() *GetWebsiteSDTListByWebsiteIDOK {
	return &GetWebsiteSDTListByWebsiteIDOK{}
}

/* GetWebsiteSDTListByWebsiteIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteSDTListByWebsiteIDOK struct {
	Payload *models.SDTPaginationResponse
}

func (o *GetWebsiteSDTListByWebsiteIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSdtListByWebsiteIdOK  %+v", 200, o.Payload)
}
func (o *GetWebsiteSDTListByWebsiteIDOK) GetPayload() *models.SDTPaginationResponse {
	return o.Payload
}

func (o *GetWebsiteSDTListByWebsiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDTPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteSDTListByWebsiteIDDefault creates a GetWebsiteSDTListByWebsiteIDDefault with default headers values
func NewGetWebsiteSDTListByWebsiteIDDefault(code int) *GetWebsiteSDTListByWebsiteIDDefault {
	return &GetWebsiteSDTListByWebsiteIDDefault{
		_statusCode: code,
	}
}

/* GetWebsiteSDTListByWebsiteIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteSDTListByWebsiteIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website SDT list by website Id default response
func (o *GetWebsiteSDTListByWebsiteIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteSDTListByWebsiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSDTListByWebsiteId default  %+v", o._statusCode, o.Payload)
}
func (o *GetWebsiteSDTListByWebsiteIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteSDTListByWebsiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
