// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// CancelInstallationAcceptedCode is the HTTP code returned for type CancelInstallationAccepted
const CancelInstallationAcceptedCode int = 202

/*CancelInstallationAccepted Success.

swagger:response cancelInstallationAccepted
*/
type CancelInstallationAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewCancelInstallationAccepted creates CancelInstallationAccepted with default headers values
func NewCancelInstallationAccepted() *CancelInstallationAccepted {

	return &CancelInstallationAccepted{}
}

// WithPayload adds the payload to the cancel installation accepted response
func (o *CancelInstallationAccepted) WithPayload(payload *models.Cluster) *CancelInstallationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cancel installation accepted response
func (o *CancelInstallationAccepted) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CancelInstallationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CancelInstallationNotFoundCode is the HTTP code returned for type CancelInstallationNotFound
const CancelInstallationNotFoundCode int = 404

/*CancelInstallationNotFound Error.

swagger:response cancelInstallationNotFound
*/
type CancelInstallationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCancelInstallationNotFound creates CancelInstallationNotFound with default headers values
func NewCancelInstallationNotFound() *CancelInstallationNotFound {

	return &CancelInstallationNotFound{}
}

// WithPayload adds the payload to the cancel installation not found response
func (o *CancelInstallationNotFound) WithPayload(payload *models.Error) *CancelInstallationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cancel installation not found response
func (o *CancelInstallationNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CancelInstallationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CancelInstallationConflictCode is the HTTP code returned for type CancelInstallationConflict
const CancelInstallationConflictCode int = 409

/*CancelInstallationConflict Error.

swagger:response cancelInstallationConflict
*/
type CancelInstallationConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCancelInstallationConflict creates CancelInstallationConflict with default headers values
func NewCancelInstallationConflict() *CancelInstallationConflict {

	return &CancelInstallationConflict{}
}

// WithPayload adds the payload to the cancel installation conflict response
func (o *CancelInstallationConflict) WithPayload(payload *models.Error) *CancelInstallationConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cancel installation conflict response
func (o *CancelInstallationConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CancelInstallationConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CancelInstallationInternalServerErrorCode is the HTTP code returned for type CancelInstallationInternalServerError
const CancelInstallationInternalServerErrorCode int = 500

/*CancelInstallationInternalServerError Error.

swagger:response cancelInstallationInternalServerError
*/
type CancelInstallationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCancelInstallationInternalServerError creates CancelInstallationInternalServerError with default headers values
func NewCancelInstallationInternalServerError() *CancelInstallationInternalServerError {

	return &CancelInstallationInternalServerError{}
}

// WithPayload adds the payload to the cancel installation internal server error response
func (o *CancelInstallationInternalServerError) WithPayload(payload *models.Error) *CancelInstallationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cancel installation internal server error response
func (o *CancelInstallationInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CancelInstallationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
