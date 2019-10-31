// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InstanceResponse instance response
// swagger:model InstanceResponse
type InstanceResponse struct {

	// build cause
	BuildCause *BuildCause `json:"build_cause,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this instance response
func (m *InstanceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildCause(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceResponse) validateBuildCause(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildCause) { // not required
		return nil
	}

	if m.BuildCause != nil {
		if err := m.BuildCause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build_cause")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceResponse) UnmarshalBinary(b []byte) error {
	var res InstanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}