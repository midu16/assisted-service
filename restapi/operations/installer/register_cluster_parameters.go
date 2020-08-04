// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/openshift/assisted-service/models"
)

// NewRegisterClusterParams creates a new RegisterClusterParams object
// no default values defined in spec.
func NewRegisterClusterParams() RegisterClusterParams {

	return RegisterClusterParams{}
}

// RegisterClusterParams contains all the bound params for the register cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters RegisterCluster
type RegisterClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	NewClusterParams *models.ClusterCreateParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRegisterClusterParams() beforehand.
func (o *RegisterClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ClusterCreateParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("newClusterParams", "body", ""))
			} else {
				res = append(res, errors.NewParseError("newClusterParams", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.NewClusterParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("newClusterParams", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
