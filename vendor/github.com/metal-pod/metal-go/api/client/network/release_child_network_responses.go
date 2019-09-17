// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// ReleaseChildNetworkReader is a Reader for the ReleaseChildNetwork structure.
type ReleaseChildNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReleaseChildNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReleaseChildNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewReleaseChildNetworkConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReleaseChildNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReleaseChildNetworkOK creates a ReleaseChildNetworkOK with default headers values
func NewReleaseChildNetworkOK() *ReleaseChildNetworkOK {
	return &ReleaseChildNetworkOK{}
}

/*ReleaseChildNetworkOK handles this case with default header values.

OK
*/
type ReleaseChildNetworkOK struct {
	Payload *models.V1NetworkResponse
}

func (o *ReleaseChildNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/network/release/{id}][%d] releaseChildNetworkOK  %+v", 200, o.Payload)
}

func (o *ReleaseChildNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseChildNetworkConflict creates a ReleaseChildNetworkConflict with default headers values
func NewReleaseChildNetworkConflict() *ReleaseChildNetworkConflict {
	return &ReleaseChildNetworkConflict{}
}

/*ReleaseChildNetworkConflict handles this case with default header values.

Conflict
*/
type ReleaseChildNetworkConflict struct {
	Payload *models.HttperrorsHTTPErrorResponse
}

func (o *ReleaseChildNetworkConflict) Error() string {
	return fmt.Sprintf("[POST /v1/network/release/{id}][%d] releaseChildNetworkConflict  %+v", 409, o.Payload)
}

func (o *ReleaseChildNetworkConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseChildNetworkDefault creates a ReleaseChildNetworkDefault with default headers values
func NewReleaseChildNetworkDefault(code int) *ReleaseChildNetworkDefault {
	return &ReleaseChildNetworkDefault{
		_statusCode: code,
	}
}

/*ReleaseChildNetworkDefault handles this case with default header values.

Error
*/
type ReleaseChildNetworkDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the release child network default response
func (o *ReleaseChildNetworkDefault) Code() int {
	return o._statusCode
}

func (o *ReleaseChildNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/network/release/{id}][%d] releaseChildNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *ReleaseChildNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}