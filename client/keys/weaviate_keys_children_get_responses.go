// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateKeysChildrenGetReader is a Reader for the WeaviateKeysChildrenGet structure.
type WeaviateKeysChildrenGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateKeysChildrenGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateKeysChildrenGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateKeysChildrenGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateKeysChildrenGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateKeysChildrenGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewWeaviateKeysChildrenGetNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateKeysChildrenGetOK creates a WeaviateKeysChildrenGetOK with default headers values
func NewWeaviateKeysChildrenGetOK() *WeaviateKeysChildrenGetOK {
	return &WeaviateKeysChildrenGetOK{}
}

/*WeaviateKeysChildrenGetOK handles this case with default header values.

Successful response.
*/
type WeaviateKeysChildrenGetOK struct {
	Payload *models.KeyChildrenGetResponse
}

func (o *WeaviateKeysChildrenGetOK) Error() string {
	return fmt.Sprintf("[GET /keys/{keyId}/children][%d] weaviateKeysChildrenGetOK  %+v", 200, o.Payload)
}

func (o *WeaviateKeysChildrenGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyChildrenGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateKeysChildrenGetUnauthorized creates a WeaviateKeysChildrenGetUnauthorized with default headers values
func NewWeaviateKeysChildrenGetUnauthorized() *WeaviateKeysChildrenGetUnauthorized {
	return &WeaviateKeysChildrenGetUnauthorized{}
}

/*WeaviateKeysChildrenGetUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateKeysChildrenGetUnauthorized struct {
}

func (o *WeaviateKeysChildrenGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /keys/{keyId}/children][%d] weaviateKeysChildrenGetUnauthorized ", 401)
}

func (o *WeaviateKeysChildrenGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateKeysChildrenGetForbidden creates a WeaviateKeysChildrenGetForbidden with default headers values
func NewWeaviateKeysChildrenGetForbidden() *WeaviateKeysChildrenGetForbidden {
	return &WeaviateKeysChildrenGetForbidden{}
}

/*WeaviateKeysChildrenGetForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateKeysChildrenGetForbidden struct {
}

func (o *WeaviateKeysChildrenGetForbidden) Error() string {
	return fmt.Sprintf("[GET /keys/{keyId}/children][%d] weaviateKeysChildrenGetForbidden ", 403)
}

func (o *WeaviateKeysChildrenGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateKeysChildrenGetNotFound creates a WeaviateKeysChildrenGetNotFound with default headers values
func NewWeaviateKeysChildrenGetNotFound() *WeaviateKeysChildrenGetNotFound {
	return &WeaviateKeysChildrenGetNotFound{}
}

/*WeaviateKeysChildrenGetNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateKeysChildrenGetNotFound struct {
}

func (o *WeaviateKeysChildrenGetNotFound) Error() string {
	return fmt.Sprintf("[GET /keys/{keyId}/children][%d] weaviateKeysChildrenGetNotFound ", 404)
}

func (o *WeaviateKeysChildrenGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateKeysChildrenGetNotImplemented creates a WeaviateKeysChildrenGetNotImplemented with default headers values
func NewWeaviateKeysChildrenGetNotImplemented() *WeaviateKeysChildrenGetNotImplemented {
	return &WeaviateKeysChildrenGetNotImplemented{}
}

/*WeaviateKeysChildrenGetNotImplemented handles this case with default header values.

Not (yet) implemented
*/
type WeaviateKeysChildrenGetNotImplemented struct {
}

func (o *WeaviateKeysChildrenGetNotImplemented) Error() string {
	return fmt.Sprintf("[GET /keys/{keyId}/children][%d] weaviateKeysChildrenGetNotImplemented ", 501)
}

func (o *WeaviateKeysChildrenGetNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
