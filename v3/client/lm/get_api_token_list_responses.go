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

// GetAPITokenListReader is a Reader for the GetAPITokenList structure.
type GetAPITokenListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITokenListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPITokenListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAPITokenListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPITokenListOK creates a GetAPITokenListOK with default headers values
func NewGetAPITokenListOK() *GetAPITokenListOK {
	return &GetAPITokenListOK{}
}

/* GetAPITokenListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAPITokenListOK struct {
	Payload *models.APITokenPaginationResponse
}

func (o *GetAPITokenListOK) Error() string {
	return fmt.Sprintf("[GET /setting/admins/apitokens][%d] getApiTokenListOK  %+v", 200, o.Payload)
}
func (o *GetAPITokenListOK) GetPayload() *models.APITokenPaginationResponse {
	return o.Payload
}

func (o *GetAPITokenListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APITokenPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPITokenListDefault creates a GetAPITokenListDefault with default headers values
func NewGetAPITokenListDefault(code int) *GetAPITokenListDefault {
	return &GetAPITokenListDefault{
		_statusCode: code,
	}
}

/* GetAPITokenListDefault describes a response with status code -1, with default header values.

Error
*/
type GetAPITokenListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get Api token list default response
func (o *GetAPITokenListDefault) Code() int {
	return o._statusCode
}

func (o *GetAPITokenListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/admins/apitokens][%d] getApiTokenList default  %+v", o._statusCode, o.Payload)
}
func (o *GetAPITokenListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAPITokenListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}