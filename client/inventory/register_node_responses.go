// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// RegisterNodeReader is a Reader for the RegisterNode structure.
type RegisterNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRegisterNodeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 405:
		result := NewRegisterNodeMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegisterNodeCreated creates a RegisterNodeCreated with default headers values
func NewRegisterNodeCreated() *RegisterNodeCreated {
	return &RegisterNodeCreated{}
}

/*RegisterNodeCreated handles this case with default header values.

Created
*/
type RegisterNodeCreated struct {
	Payload *models.RegisteredNode
}

func (o *RegisterNodeCreated) Error() string {
	return fmt.Sprintf("[POST /node/register][%d] registerNodeCreated  %+v", 201, o.Payload)
}

func (o *RegisterNodeCreated) GetPayload() *models.RegisteredNode {
	return o.Payload
}

func (o *RegisterNodeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisteredNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterNodeMethodNotAllowed creates a RegisterNodeMethodNotAllowed with default headers values
func NewRegisterNodeMethodNotAllowed() *RegisterNodeMethodNotAllowed {
	return &RegisterNodeMethodNotAllowed{}
}

/*RegisterNodeMethodNotAllowed handles this case with default header values.

Invalid input
*/
type RegisterNodeMethodNotAllowed struct {
}

func (o *RegisterNodeMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /node/register][%d] registerNodeMethodNotAllowed ", 405)
}

func (o *RegisterNodeMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
