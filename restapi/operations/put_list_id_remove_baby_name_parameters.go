// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutListIDRemoveBabyNameParams creates a new PutListIDRemoveBabyNameParams object
// no default values defined in spec.
func NewPutListIDRemoveBabyNameParams() PutListIDRemoveBabyNameParams {

	return PutListIDRemoveBabyNameParams{}
}

// PutListIDRemoveBabyNameParams contains all the bound params for the put list ID remove baby name operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutListIDRemoveBabyName
type PutListIDRemoveBabyNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	BabyNames string
	/*
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutListIDRemoveBabyNameParams() beforehand.
func (o *PutListIDRemoveBabyNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBabyNames, qhkBabyNames, _ := qs.GetOK("babyNames")
	if err := o.bindBabyNames(qBabyNames, qhkBabyNames, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBabyNames binds and validates parameter BabyNames from query.
func (o *PutListIDRemoveBabyNameParams) bindBabyNames(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("babyNames", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("babyNames", "query", raw); err != nil {
		return err
	}

	o.BabyNames = raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PutListIDRemoveBabyNameParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}
