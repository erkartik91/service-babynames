// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/erkartik91/service-babynames/models"
)

// GetListIDOKCode is the HTTP code returned for type GetListIDOK
const GetListIDOKCode int = 200

/*GetListIDOK OK

swagger:response getListIdOK
*/
type GetListIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.List `json:"body,omitempty"`
}

// NewGetListIDOK creates GetListIDOK with default headers values
func NewGetListIDOK() *GetListIDOK {

	return &GetListIDOK{}
}

// WithPayload adds the payload to the get list Id o k response
func (o *GetListIDOK) WithPayload(payload *models.List) *GetListIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list Id o k response
func (o *GetListIDOK) SetPayload(payload *models.List) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetListIDDefault Error

swagger:response getListIdDefault
*/
type GetListIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetListIDDefault creates GetListIDDefault with default headers values
func NewGetListIDDefault(code int) *GetListIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetListIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get list ID default response
func (o *GetListIDDefault) WithStatusCode(code int) *GetListIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get list ID default response
func (o *GetListIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get list ID default response
func (o *GetListIDDefault) WithPayload(payload *models.Error) *GetListIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list ID default response
func (o *GetListIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
