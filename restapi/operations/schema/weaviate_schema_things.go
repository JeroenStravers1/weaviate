/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
   

package schema

 
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateSchemaThingsHandlerFunc turns a function with the right signature into a weaviate schema things handler
type WeaviateSchemaThingsHandlerFunc func(WeaviateSchemaThingsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateSchemaThingsHandlerFunc) Handle(params WeaviateSchemaThingsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateSchemaThingsHandler interface for that can handle valid weaviate schema things params
type WeaviateSchemaThingsHandler interface {
	Handle(WeaviateSchemaThingsParams, interface{}) middleware.Responder
}

// NewWeaviateSchemaThings creates a new http.Handler for the weaviate schema things operation
func NewWeaviateSchemaThings(ctx *middleware.Context, handler WeaviateSchemaThingsHandler) *WeaviateSchemaThings {
	return &WeaviateSchemaThings{Context: ctx, Handler: handler}
}

/*WeaviateSchemaThings swagger:route GET /schema/things schema weaviateSchemaThings

Download the schema file where all things are based on.

Download the schema where all things are based on.

*/
type WeaviateSchemaThings struct {
	Context *middleware.Context
	Handler WeaviateSchemaThingsHandler
}

func (o *WeaviateSchemaThings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateSchemaThingsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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