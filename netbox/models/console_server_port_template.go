// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsoleServerPortTemplate console server port template
//
// swagger:model ConsoleServerPortTemplate
type ConsoleServerPortTemplate struct {

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// type
	Type *ConsoleServerPortTemplateType `json:"type,omitempty"`
}

// Validate validates this console server port template
func (m *ConsoleServerPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsoleServerPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this console server port template based on the context it is used
func (m *ConsoleServerPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsoleServerPortTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceType != nil {
		if err := m.DeviceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortTemplate) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsoleServerPortTemplateType Type
//
// swagger:model ConsoleServerPortTemplateType
type ConsoleServerPortTemplateType struct {

	// label
	// Required: true
	// Enum: [DE-9 DB-25 RJ-11 RJ-12 RJ-45 USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Value *string `json:"value"`
}

// Validate validates this console server port template type
func (m *ConsoleServerPortTemplateType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consoleServerPortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE-9","DB-25","RJ-11","RJ-12","RJ-45","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTemplateTypeTypeLabelPropEnum = append(consoleServerPortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// ConsoleServerPortTemplateTypeLabelDE9 captures enum value "DE-9"
	ConsoleServerPortTemplateTypeLabelDE9 string = "DE-9"

	// ConsoleServerPortTemplateTypeLabelDB25 captures enum value "DB-25"
	ConsoleServerPortTemplateTypeLabelDB25 string = "DB-25"

	// ConsoleServerPortTemplateTypeLabelRJ11 captures enum value "RJ-11"
	ConsoleServerPortTemplateTypeLabelRJ11 string = "RJ-11"

	// ConsoleServerPortTemplateTypeLabelRJ12 captures enum value "RJ-12"
	ConsoleServerPortTemplateTypeLabelRJ12 string = "RJ-12"

	// ConsoleServerPortTemplateTypeLabelRJ45 captures enum value "RJ-45"
	ConsoleServerPortTemplateTypeLabelRJ45 string = "RJ-45"

	// ConsoleServerPortTemplateTypeLabelUSBTypeA captures enum value "USB Type A"
	ConsoleServerPortTemplateTypeLabelUSBTypeA string = "USB Type A"

	// ConsoleServerPortTemplateTypeLabelUSBTypeB captures enum value "USB Type B"
	ConsoleServerPortTemplateTypeLabelUSBTypeB string = "USB Type B"

	// ConsoleServerPortTemplateTypeLabelUSBTypeC captures enum value "USB Type C"
	ConsoleServerPortTemplateTypeLabelUSBTypeC string = "USB Type C"

	// ConsoleServerPortTemplateTypeLabelUSBMiniA captures enum value "USB Mini A"
	ConsoleServerPortTemplateTypeLabelUSBMiniA string = "USB Mini A"

	// ConsoleServerPortTemplateTypeLabelUSBMiniB captures enum value "USB Mini B"
	ConsoleServerPortTemplateTypeLabelUSBMiniB string = "USB Mini B"

	// ConsoleServerPortTemplateTypeLabelUSBMicroA captures enum value "USB Micro A"
	ConsoleServerPortTemplateTypeLabelUSBMicroA string = "USB Micro A"

	// ConsoleServerPortTemplateTypeLabelUSBMicroB captures enum value "USB Micro B"
	ConsoleServerPortTemplateTypeLabelUSBMicroB string = "USB Micro B"

	// ConsoleServerPortTemplateTypeLabelOther captures enum value "Other"
	ConsoleServerPortTemplateTypeLabelOther string = "Other"
)

// prop value enum
func (m *ConsoleServerPortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consoleServerPortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTemplateTypeTypeValuePropEnum = append(consoleServerPortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// ConsoleServerPortTemplateTypeValueDe9 captures enum value "de-9"
	ConsoleServerPortTemplateTypeValueDe9 string = "de-9"

	// ConsoleServerPortTemplateTypeValueDb25 captures enum value "db-25"
	ConsoleServerPortTemplateTypeValueDb25 string = "db-25"

	// ConsoleServerPortTemplateTypeValueRj11 captures enum value "rj-11"
	ConsoleServerPortTemplateTypeValueRj11 string = "rj-11"

	// ConsoleServerPortTemplateTypeValueRj12 captures enum value "rj-12"
	ConsoleServerPortTemplateTypeValueRj12 string = "rj-12"

	// ConsoleServerPortTemplateTypeValueRj45 captures enum value "rj-45"
	ConsoleServerPortTemplateTypeValueRj45 string = "rj-45"

	// ConsoleServerPortTemplateTypeValueUsba captures enum value "usb-a"
	ConsoleServerPortTemplateTypeValueUsba string = "usb-a"

	// ConsoleServerPortTemplateTypeValueUsbb captures enum value "usb-b"
	ConsoleServerPortTemplateTypeValueUsbb string = "usb-b"

	// ConsoleServerPortTemplateTypeValueUsbc captures enum value "usb-c"
	ConsoleServerPortTemplateTypeValueUsbc string = "usb-c"

	// ConsoleServerPortTemplateTypeValueUsbMinia captures enum value "usb-mini-a"
	ConsoleServerPortTemplateTypeValueUsbMinia string = "usb-mini-a"

	// ConsoleServerPortTemplateTypeValueUsbMinib captures enum value "usb-mini-b"
	ConsoleServerPortTemplateTypeValueUsbMinib string = "usb-mini-b"

	// ConsoleServerPortTemplateTypeValueUsbMicroa captures enum value "usb-micro-a"
	ConsoleServerPortTemplateTypeValueUsbMicroa string = "usb-micro-a"

	// ConsoleServerPortTemplateTypeValueUsbMicrob captures enum value "usb-micro-b"
	ConsoleServerPortTemplateTypeValueUsbMicrob string = "usb-micro-b"

	// ConsoleServerPortTemplateTypeValueOther captures enum value "other"
	ConsoleServerPortTemplateTypeValueOther string = "other"
)

// prop value enum
func (m *ConsoleServerPortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console server port template type based on context it is used
func (m *ConsoleServerPortTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortTemplateType) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
