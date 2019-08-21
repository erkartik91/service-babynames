// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutListIDRemoveBabyNameHandlerFunc turns a function with the right signature into a put list ID remove baby name handler
type PutListIDRemoveBabyNameHandlerFunc func(PutListIDRemoveBabyNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutListIDRemoveBabyNameHandlerFunc) Handle(params PutListIDRemoveBabyNameParams) middleware.Responder {
	return fn(params)
}

// PutListIDRemoveBabyNameHandler interface for that can handle valid put list ID remove baby name params
type PutListIDRemoveBabyNameHandler interface {
	Handle(PutListIDRemoveBabyNameParams) middleware.Responder
}

// NewPutListIDRemoveBabyName creates a new http.Handler for the put list ID remove baby name operation
func NewPutListIDRemoveBabyName(ctx *middleware.Context, handler PutListIDRemoveBabyNameHandler) *PutListIDRemoveBabyName {
	return &PutListIDRemoveBabyName{Context: ctx, Handler: handler}
}

/*PutListIDRemoveBabyName swagger:route PUT /list/{id}/remove-baby-name putListIdRemoveBabyName

Remove given name from list

*/
type PutListIDRemoveBabyName struct {
	Context *middleware.Context
	Handler PutListIDRemoveBabyNameHandler
}

func (o *PutListIDRemoveBabyName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutListIDRemoveBabyNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
