// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// GenerateClusterISOCreatedCode is the HTTP code returned for type GenerateClusterISOCreated
const GenerateClusterISOCreatedCode int = 201

/*GenerateClusterISOCreated Success.

swagger:response generateClusterISOCreated
*/
type GenerateClusterISOCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewGenerateClusterISOCreated creates GenerateClusterISOCreated with default headers values
func NewGenerateClusterISOCreated() *GenerateClusterISOCreated {

	return &GenerateClusterISOCreated{}
}

// WithPayload adds the payload to the generate cluster i s o created response
func (o *GenerateClusterISOCreated) WithPayload(payload *models.Cluster) *GenerateClusterISOCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate cluster i s o created response
func (o *GenerateClusterISOCreated) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateClusterISOCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateClusterISOBadRequestCode is the HTTP code returned for type GenerateClusterISOBadRequest
const GenerateClusterISOBadRequestCode int = 400

/*GenerateClusterISOBadRequest Error.

swagger:response generateClusterISOBadRequest
*/
type GenerateClusterISOBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateClusterISOBadRequest creates GenerateClusterISOBadRequest with default headers values
func NewGenerateClusterISOBadRequest() *GenerateClusterISOBadRequest {

	return &GenerateClusterISOBadRequest{}
}

// WithPayload adds the payload to the generate cluster i s o bad request response
func (o *GenerateClusterISOBadRequest) WithPayload(payload *models.Error) *GenerateClusterISOBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate cluster i s o bad request response
func (o *GenerateClusterISOBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateClusterISOBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateClusterISONotFoundCode is the HTTP code returned for type GenerateClusterISONotFound
const GenerateClusterISONotFoundCode int = 404

/*GenerateClusterISONotFound Error.

swagger:response generateClusterISONotFound
*/
type GenerateClusterISONotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateClusterISONotFound creates GenerateClusterISONotFound with default headers values
func NewGenerateClusterISONotFound() *GenerateClusterISONotFound {

	return &GenerateClusterISONotFound{}
}

// WithPayload adds the payload to the generate cluster i s o not found response
func (o *GenerateClusterISONotFound) WithPayload(payload *models.Error) *GenerateClusterISONotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate cluster i s o not found response
func (o *GenerateClusterISONotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateClusterISONotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateClusterISOConflictCode is the HTTP code returned for type GenerateClusterISOConflict
const GenerateClusterISOConflictCode int = 409

/*GenerateClusterISOConflict Error.

swagger:response generateClusterISOConflict
*/
type GenerateClusterISOConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateClusterISOConflict creates GenerateClusterISOConflict with default headers values
func NewGenerateClusterISOConflict() *GenerateClusterISOConflict {

	return &GenerateClusterISOConflict{}
}

// WithPayload adds the payload to the generate cluster i s o conflict response
func (o *GenerateClusterISOConflict) WithPayload(payload *models.Error) *GenerateClusterISOConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate cluster i s o conflict response
func (o *GenerateClusterISOConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateClusterISOConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateClusterISOInternalServerErrorCode is the HTTP code returned for type GenerateClusterISOInternalServerError
const GenerateClusterISOInternalServerErrorCode int = 500

/*GenerateClusterISOInternalServerError Error.

swagger:response generateClusterISOInternalServerError
*/
type GenerateClusterISOInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateClusterISOInternalServerError creates GenerateClusterISOInternalServerError with default headers values
func NewGenerateClusterISOInternalServerError() *GenerateClusterISOInternalServerError {

	return &GenerateClusterISOInternalServerError{}
}

// WithPayload adds the payload to the generate cluster i s o internal server error response
func (o *GenerateClusterISOInternalServerError) WithPayload(payload *models.Error) *GenerateClusterISOInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate cluster i s o internal server error response
func (o *GenerateClusterISOInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateClusterISOInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
