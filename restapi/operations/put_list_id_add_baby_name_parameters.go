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

// NewPutListIDAddBabyNameParams creates a new PutListIDAddBabyNameParams object
// no default values defined in spec.
func NewPutListIDAddBabyNameParams() PutListIDAddBabyNameParams {

	return PutListIDAddBabyNameParams{}
}

// PutListIDAddBabyNameParams contains all the bound params for the put list ID add baby name operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutListIDAddBabyName
type PutListIDAddBabyNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	BabyName string
	/*
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutListIDAddBabyNameParams() beforehand.
func (o *PutListIDAddBabyNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBabyName, qhkBabyName, _ := qs.GetOK("babyName")
	if err := o.bindBabyName(qBabyName, qhkBabyName, route.Formats); err != nil {
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

// bindBabyName binds and validates parameter BabyName from query.
func (o *PutListIDAddBabyNameParams) bindBabyName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("babyName", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("babyName", "query", raw); err != nil {
		return err
	}

	o.BabyName = raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PutListIDAddBabyNameParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}
