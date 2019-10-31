// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Modification modification
// swagger:model Modification
type Modification struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// email address
	EmailAddress string `json:"email_address,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// modified time
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// revision
	Revision string `json:"revision,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this modification
func (m *Modification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Modification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Modification) UnmarshalBinary(b []byte) error {
	var res Modification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}