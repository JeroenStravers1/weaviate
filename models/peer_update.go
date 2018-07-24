/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PeerUpdate A single peer in the network
// swagger:model PeerUpdate
type PeerUpdate struct {

	// The session ID of the peer
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Human readable name
	Name string `json:"name,omitempty"`

	// The location where the peer is exposed to the internet
	// Format: uri
	URI strfmt.URI `json:"uri,omitempty"`
}

// Validate validates this peer update
func (m *PeerUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerUpdate) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PeerUpdate) validateURI(formats strfmt.Registry) error {

	if swag.IsZero(m.URI) { // not required
		return nil
	}

	if err := validate.FormatOf("uri", "body", "uri", m.URI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerUpdate) UnmarshalBinary(b []byte) error {
	var res PeerUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}