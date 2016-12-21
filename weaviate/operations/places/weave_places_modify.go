package places


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeavePlacesModifyHandlerFunc turns a function with the right signature into a weave places modify handler
type WeavePlacesModifyHandlerFunc func(WeavePlacesModifyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeavePlacesModifyHandlerFunc) Handle(params WeavePlacesModifyParams) middleware.Responder {
	return fn(params)
}

// WeavePlacesModifyHandler interface for that can handle valid weave places modify params
type WeavePlacesModifyHandler interface {
	Handle(WeavePlacesModifyParams) middleware.Responder
}

// NewWeavePlacesModify creates a new http.Handler for the weave places modify operation
func NewWeavePlacesModify(ctx *middleware.Context, handler WeavePlacesModifyHandler) *WeavePlacesModify {
	return &WeavePlacesModify{Context: ctx, Handler: handler}
}

/*WeavePlacesModify swagger:route POST /places/{placeId}/modify places weavePlacesModify

Modifies a place.

*/
type WeavePlacesModify struct {
	Context *middleware.Context
	Handler WeavePlacesModifyHandler
}

func (o *WeavePlacesModify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeavePlacesModifyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}