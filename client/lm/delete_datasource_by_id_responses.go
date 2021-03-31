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

// DeleteDatasourceByIDReader is a Reader for the DeleteDatasourceByID structure.
type DeleteDatasourceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatasourceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDatasourceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDatasourceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDatasourceByIDOK creates a DeleteDatasourceByIDOK with default headers values
func NewDeleteDatasourceByIDOK() *DeleteDatasourceByIDOK {
	return &DeleteDatasourceByIDOK{}
}

/* DeleteDatasourceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteDatasourceByIDOK struct {
	Payload interface{}
}

func (o *DeleteDatasourceByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/datasources/{id}][%d] deleteDatasourceByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteDatasourceByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDatasourceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatasourceByIDDefault creates a DeleteDatasourceByIDDefault with default headers values
func NewDeleteDatasourceByIDDefault(code int) *DeleteDatasourceByIDDefault {
	return &DeleteDatasourceByIDDefault{
		_statusCode: code,
	}
}

/* DeleteDatasourceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteDatasourceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete datasource by Id default response
func (o *DeleteDatasourceByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDatasourceByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/datasources/{id}][%d] deleteDatasourceById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteDatasourceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDatasourceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
