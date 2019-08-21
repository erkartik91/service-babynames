// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/erkartik91/service-babynames/models"
)

// PutListIDRemoveBabyNameOKCode is the HTTP code returned for type PutListIDRemoveBabyNameOK
const PutListIDRemoveBabyNameOKCode int = 200

/*PutListIDRemoveBabyNameOK OK

swagger:response putListIdRemoveBabyNameOK
*/
type PutListIDRemoveBabyNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.List `json:"body,omitempty"`
}

// NewPutListIDRemoveBabyNameOK creates PutListIDRemoveBabyNameOK with default headers values
func NewPutListIDRemoveBabyNameOK() *PutListIDRemoveBabyNameOK {

	return &PutListIDRemoveBabyNameOK{}
}

// WithPayload adds the payload to the put list Id remove baby name o k response
func (o *PutListIDRemoveBabyNameOK) WithPayload(payload *models.List) *PutListIDRemoveBabyNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put list Id remove baby name o k response
func (o *PutListIDRemoveBabyNameOK) SetPayload(payload *models.List) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutListIDRemoveBabyNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutListIDRemoveBabyNameDefault Error

swagger:response putListIdRemoveBabyNameDefault
*/
type PutListIDRemoveBabyNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutListIDRemoveBabyNameDefault creates PutListIDRemoveBabyNameDefault with default headers values
func NewPutListIDRemoveBabyNameDefault(code int) *PutListIDRemoveBabyNameDefault {
	if code <= 0 {
		code = 500
	}

	return &PutListIDRemoveBabyNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put list ID remove baby name default response
func (o *PutListIDRemoveBabyNameDefault) WithStatusCode(code int) *PutListIDRemoveBabyNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put list ID remove baby name default response
func (o *PutListIDRemoveBabyNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put list ID remove baby name default response
func (o *PutListIDRemoveBabyNameDefault) WithPayload(payload *models.Error) *PutListIDRemoveBabyNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put list ID remove baby name default response
func (o *PutListIDRemoveBabyNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutListIDRemoveBabyNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
