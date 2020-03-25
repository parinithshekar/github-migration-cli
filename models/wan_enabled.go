// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// WanEnabled WanEnabled
//
// Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
//
// swagger:model WanEnabled
type WanEnabled string

const (

	// WanEnabledEnabled captures enum value "enabled"
	WanEnabledEnabled WanEnabled = "enabled"

	// WanEnabledDisabled captures enum value "disabled"
	WanEnabledDisabled WanEnabled = "disabled"

	// WanEnabledNotConfigured captures enum value "not configured"
	WanEnabledNotConfigured WanEnabled = "not configured"
)

// for schema
var wanEnabledEnum []interface{}

func init() {
	var res []WanEnabled
	if err := json.Unmarshal([]byte(`["enabled","disabled","not configured"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wanEnabledEnum = append(wanEnabledEnum, v)
	}
}

func (m WanEnabled) validateWanEnabledEnum(path, location string, value WanEnabled) error {
	if err := validate.Enum(path, location, value, wanEnabledEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this wan enabled
func (m WanEnabled) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateWanEnabledEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}