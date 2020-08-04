// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// UpdateClusterCreatedCode is the HTTP code returned for type UpdateClusterCreated
const UpdateClusterCreatedCode int = 201

/*UpdateClusterCreated Success.

swagger:response updateClusterCreated
*/
type UpdateClusterCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewUpdateClusterCreated creates UpdateClusterCreated with default headers values
func NewUpdateClusterCreated() *UpdateClusterCreated {

	return &UpdateClusterCreated{}
}

// WithPayload adds the payload to the update cluster created response
func (o *UpdateClusterCreated) WithPayload(payload *models.Cluster) *UpdateClusterCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster created response
func (o *UpdateClusterCreated) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterBadRequestCode is the HTTP code returned for type UpdateClusterBadRequest
const UpdateClusterBadRequestCode int = 400

/*UpdateClusterBadRequest Error.

swagger:response updateClusterBadRequest
*/
type UpdateClusterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterBadRequest creates UpdateClusterBadRequest with default headers values
func NewUpdateClusterBadRequest() *UpdateClusterBadRequest {

	return &UpdateClusterBadRequest{}
}

// WithPayload adds the payload to the update cluster bad request response
func (o *UpdateClusterBadRequest) WithPayload(payload *models.Error) *UpdateClusterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster bad request response
func (o *UpdateClusterBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterNotFoundCode is the HTTP code returned for type UpdateClusterNotFound
const UpdateClusterNotFoundCode int = 404

/*UpdateClusterNotFound Error.

swagger:response updateClusterNotFound
*/
type UpdateClusterNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterNotFound creates UpdateClusterNotFound with default headers values
func NewUpdateClusterNotFound() *UpdateClusterNotFound {

	return &UpdateClusterNotFound{}
}

// WithPayload adds the payload to the update cluster not found response
func (o *UpdateClusterNotFound) WithPayload(payload *models.Error) *UpdateClusterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster not found response
func (o *UpdateClusterNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterConflictCode is the HTTP code returned for type UpdateClusterConflict
const UpdateClusterConflictCode int = 409

/*UpdateClusterConflict Error.

swagger:response updateClusterConflict
*/
type UpdateClusterConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterConflict creates UpdateClusterConflict with default headers values
func NewUpdateClusterConflict() *UpdateClusterConflict {

	return &UpdateClusterConflict{}
}

// WithPayload adds the payload to the update cluster conflict response
func (o *UpdateClusterConflict) WithPayload(payload *models.Error) *UpdateClusterConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster conflict response
func (o *UpdateClusterConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInternalServerErrorCode is the HTTP code returned for type UpdateClusterInternalServerError
const UpdateClusterInternalServerErrorCode int = 500

/*UpdateClusterInternalServerError Error.

swagger:response updateClusterInternalServerError
*/
type UpdateClusterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInternalServerError creates UpdateClusterInternalServerError with default headers values
func NewUpdateClusterInternalServerError() *UpdateClusterInternalServerError {

	return &UpdateClusterInternalServerError{}
}

// WithPayload adds the payload to the update cluster internal server error response
func (o *UpdateClusterInternalServerError) WithPayload(payload *models.Error) *UpdateClusterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster internal server error response
func (o *UpdateClusterInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
