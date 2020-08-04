// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// GetClusterReader is a Reader for the GetCluster structure.
type GetClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterOK creates a GetClusterOK with default headers values
func NewGetClusterOK() *GetClusterOK {
	return &GetClusterOK{}
}

/*GetClusterOK handles this case with default header values.

Success.
*/
type GetClusterOK struct {
	Payload *models.Cluster
}

func (o *GetClusterOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GetClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *GetClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterNotFound creates a GetClusterNotFound with default headers values
func NewGetClusterNotFound() *GetClusterNotFound {
	return &GetClusterNotFound{}
}

/*GetClusterNotFound handles this case with default header values.

Error.
*/
type GetClusterNotFound struct {
	Payload *models.Error
}

func (o *GetClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}][%d] getClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterInternalServerError creates a GetClusterInternalServerError with default headers values
func NewGetClusterInternalServerError() *GetClusterInternalServerError {
	return &GetClusterInternalServerError{}
}

/*GetClusterInternalServerError handles this case with default header values.

Error.
*/
type GetClusterInternalServerError struct {
	Payload *models.Error
}

func (o *GetClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}][%d] getClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
