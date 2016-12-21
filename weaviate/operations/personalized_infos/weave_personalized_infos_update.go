package personalized_infos


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeavePersonalizedInfosUpdateHandlerFunc turns a function with the right signature into a weave personalized infos update handler
type WeavePersonalizedInfosUpdateHandlerFunc func(WeavePersonalizedInfosUpdateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeavePersonalizedInfosUpdateHandlerFunc) Handle(params WeavePersonalizedInfosUpdateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeavePersonalizedInfosUpdateHandler interface for that can handle valid weave personalized infos update params
type WeavePersonalizedInfosUpdateHandler interface {
	Handle(WeavePersonalizedInfosUpdateParams, interface{}) middleware.Responder
}

// NewWeavePersonalizedInfosUpdate creates a new http.Handler for the weave personalized infos update operation
func NewWeavePersonalizedInfosUpdate(ctx *middleware.Context, handler WeavePersonalizedInfosUpdateHandler) *WeavePersonalizedInfosUpdate {
	return &WeavePersonalizedInfosUpdate{Context: ctx, Handler: handler}
}

/*WeavePersonalizedInfosUpdate swagger:route PUT /devices/{deviceId}/personalizedInfos/{personalizedInfoId} personalizedInfos weavePersonalizedInfosUpdate

Update the personalized info for device.

*/
type WeavePersonalizedInfosUpdate struct {
	Context *middleware.Context
	Handler WeavePersonalizedInfosUpdateHandler
}

func (o *WeavePersonalizedInfosUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeavePersonalizedInfosUpdateParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}