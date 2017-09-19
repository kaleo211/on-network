package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// UpdateSwitch update switch
// swagger:model UpdateSwitch
type UpdateSwitch struct {

	// image URL
	// Required: true
	ImageURL *string `json:"imageURL"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// password
	// Required: true
	Password *string `json:"password"`

	// switch type
	// Required: true
	SwitchType *string `json:"switchType"`

	// update type
	// Required: true
	UpdateType *string `json:"updateType"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this update switch
func (m *UpdateSwitch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSwitchType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSwitch) validateImageURL(formats strfmt.Registry) error {

	if err := validate.Required("imageURL", "body", m.ImageURL); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSwitch) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSwitch) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSwitch) validateSwitchType(formats strfmt.Registry) error {

	if err := validate.Required("switchType", "body", m.SwitchType); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSwitch) validateUpdateType(formats strfmt.Registry) error {

	if err := validate.Required("updateType", "body", m.UpdateType); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSwitch) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}
