// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateThingsPatchParams creates a new WeaviateThingsPatchParams object
// with the default values initialized.
func NewWeaviateThingsPatchParams() *WeaviateThingsPatchParams {
	var ()
	return &WeaviateThingsPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateThingsPatchParamsWithTimeout creates a new WeaviateThingsPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateThingsPatchParamsWithTimeout(timeout time.Duration) *WeaviateThingsPatchParams {
	var ()
	return &WeaviateThingsPatchParams{

		timeout: timeout,
	}
}

// NewWeaviateThingsPatchParamsWithContext creates a new WeaviateThingsPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateThingsPatchParamsWithContext(ctx context.Context) *WeaviateThingsPatchParams {
	var ()
	return &WeaviateThingsPatchParams{

		Context: ctx,
	}
}

// NewWeaviateThingsPatchParamsWithHTTPClient creates a new WeaviateThingsPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateThingsPatchParamsWithHTTPClient(client *http.Client) *WeaviateThingsPatchParams {
	var ()
	return &WeaviateThingsPatchParams{
		HTTPClient: client,
	}
}

/*WeaviateThingsPatchParams contains all the parameters to send to the API endpoint
for the weaviate things patch operation typically these are written to a http.Request
*/
type WeaviateThingsPatchParams struct {

	/*Body
	  JSONPatch document as defined by RFC 6902.

	*/
	Body []*models.PatchDocument
	/*ThingID
	  Unique ID of the thing.

	*/
	ThingID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate things patch params
func (o *WeaviateThingsPatchParams) WithTimeout(timeout time.Duration) *WeaviateThingsPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate things patch params
func (o *WeaviateThingsPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate things patch params
func (o *WeaviateThingsPatchParams) WithContext(ctx context.Context) *WeaviateThingsPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate things patch params
func (o *WeaviateThingsPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate things patch params
func (o *WeaviateThingsPatchParams) WithHTTPClient(client *http.Client) *WeaviateThingsPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate things patch params
func (o *WeaviateThingsPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate things patch params
func (o *WeaviateThingsPatchParams) WithBody(body []*models.PatchDocument) *WeaviateThingsPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate things patch params
func (o *WeaviateThingsPatchParams) SetBody(body []*models.PatchDocument) {
	o.Body = body
}

// WithThingID adds the thingID to the weaviate things patch params
func (o *WeaviateThingsPatchParams) WithThingID(thingID strfmt.UUID) *WeaviateThingsPatchParams {
	o.SetThingID(thingID)
	return o
}

// SetThingID adds the thingId to the weaviate things patch params
func (o *WeaviateThingsPatchParams) SetThingID(thingID strfmt.UUID) {
	o.ThingID = thingID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateThingsPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param thingId
	if err := r.SetPathParam("thingId", o.ThingID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
