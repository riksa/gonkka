// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MaterialRevision material revision
// swagger:model MaterialRevision
type MaterialRevision struct {

	// material
	Material *Material `json:"material,omitempty"`

	// modifications
	Modifications []*Modification `json:"modifications"`
}

// Validate validates this material revision
func (m *MaterialRevision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaterial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaterialRevision) validateMaterial(formats strfmt.Registry) error {

	if swag.IsZero(m.Material) { // not required
		return nil
	}

	if m.Material != nil {
		if err := m.Material.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("material")
			}
			return err
		}
	}

	return nil
}

func (m *MaterialRevision) validateModifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Modifications) { // not required
		return nil
	}

	for i := 0; i < len(m.Modifications); i++ {
		if swag.IsZero(m.Modifications[i]) { // not required
			continue
		}

		if m.Modifications[i] != nil {
			if err := m.Modifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaterialRevision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaterialRevision) UnmarshalBinary(b []byte) error {
	var res MaterialRevision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
