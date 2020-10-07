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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsoleServerPort console server port
//
// swagger:model ConsoleServerPort
type ConsoleServerPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// connection status
	ConnectionStatus *ConsoleServerPortConnectionStatus `json:"connection_status,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// tags
	Tags []string `json:"tags"`

	// type
	Type *ConsoleServerPortType `json:"type,omitempty"`
}

// Validate validates this console server port
func (m *ConsoleServerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *ConsoleServerPort) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateName(formats strfmt.Registry) error {

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

func (m *ConsoleServerPort) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConsoleServerPort) validateType(formats strfmt.Registry) error {

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

// ContextValidate validate this console server port based on the context it is used
func (m *ConsoleServerPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectionStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
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

func (m *ConsoleServerPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ConsoleServerPort) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateConnectionStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConsoleServerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPort) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsoleServerPortConnectionStatus Connection status
//
// swagger:model ConsoleServerPortConnectionStatus
type ConsoleServerPortConnectionStatus struct {

	// label
	// Required: true
	// Enum: [Not Connected Connected]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [false true]
	Value *bool `json:"value"`
}

// Validate validates this console server port connection status
func (m *ConsoleServerPortConnectionStatus) Validate(formats strfmt.Registry) error {
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

var consoleServerPortConnectionStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Not Connected","Connected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortConnectionStatusTypeLabelPropEnum = append(consoleServerPortConnectionStatusTypeLabelPropEnum, v)
	}
}

const (

	// ConsoleServerPortConnectionStatusLabelNotConnected captures enum value "Not Connected"
	ConsoleServerPortConnectionStatusLabelNotConnected string = "Not Connected"

	// ConsoleServerPortConnectionStatusLabelConnected captures enum value "Connected"
	ConsoleServerPortConnectionStatusLabelConnected string = "Connected"
)

// prop value enum
func (m *ConsoleServerPortConnectionStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortConnectionStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortConnectionStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("connection_status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consoleServerPortConnectionStatusTypeValuePropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortConnectionStatusTypeValuePropEnum = append(consoleServerPortConnectionStatusTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *ConsoleServerPortConnectionStatus) validateValueEnum(path, location string, value bool) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortConnectionStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortConnectionStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("connection_status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this console server port connection status based on the context it is used
func (m *ConsoleServerPortConnectionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortConnectionStatus) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsoleServerPortType Type
//
// swagger:model ConsoleServerPortType
type ConsoleServerPortType struct {

	// label
	// Required: true
	// Enum: [DE-9 DB-25 RJ-11 RJ-12 RJ-45 USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Value *string `json:"value"`
}

// Validate validates this console server port type
func (m *ConsoleServerPortType) Validate(formats strfmt.Registry) error {
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

var consoleServerPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE-9","DB-25","RJ-11","RJ-12","RJ-45","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTypeTypeLabelPropEnum = append(consoleServerPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// ConsoleServerPortTypeLabelDE9 captures enum value "DE-9"
	ConsoleServerPortTypeLabelDE9 string = "DE-9"

	// ConsoleServerPortTypeLabelDB25 captures enum value "DB-25"
	ConsoleServerPortTypeLabelDB25 string = "DB-25"

	// ConsoleServerPortTypeLabelRJ11 captures enum value "RJ-11"
	ConsoleServerPortTypeLabelRJ11 string = "RJ-11"

	// ConsoleServerPortTypeLabelRJ12 captures enum value "RJ-12"
	ConsoleServerPortTypeLabelRJ12 string = "RJ-12"

	// ConsoleServerPortTypeLabelRJ45 captures enum value "RJ-45"
	ConsoleServerPortTypeLabelRJ45 string = "RJ-45"

	// ConsoleServerPortTypeLabelUSBTypeA captures enum value "USB Type A"
	ConsoleServerPortTypeLabelUSBTypeA string = "USB Type A"

	// ConsoleServerPortTypeLabelUSBTypeB captures enum value "USB Type B"
	ConsoleServerPortTypeLabelUSBTypeB string = "USB Type B"

	// ConsoleServerPortTypeLabelUSBTypeC captures enum value "USB Type C"
	ConsoleServerPortTypeLabelUSBTypeC string = "USB Type C"

	// ConsoleServerPortTypeLabelUSBMiniA captures enum value "USB Mini A"
	ConsoleServerPortTypeLabelUSBMiniA string = "USB Mini A"

	// ConsoleServerPortTypeLabelUSBMiniB captures enum value "USB Mini B"
	ConsoleServerPortTypeLabelUSBMiniB string = "USB Mini B"

	// ConsoleServerPortTypeLabelUSBMicroA captures enum value "USB Micro A"
	ConsoleServerPortTypeLabelUSBMicroA string = "USB Micro A"

	// ConsoleServerPortTypeLabelUSBMicroB captures enum value "USB Micro B"
	ConsoleServerPortTypeLabelUSBMicroB string = "USB Micro B"

	// ConsoleServerPortTypeLabelOther captures enum value "Other"
	ConsoleServerPortTypeLabelOther string = "Other"
)

// prop value enum
func (m *ConsoleServerPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consoleServerPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTypeTypeValuePropEnum = append(consoleServerPortTypeTypeValuePropEnum, v)
	}
}

const (

	// ConsoleServerPortTypeValueDe9 captures enum value "de-9"
	ConsoleServerPortTypeValueDe9 string = "de-9"

	// ConsoleServerPortTypeValueDb25 captures enum value "db-25"
	ConsoleServerPortTypeValueDb25 string = "db-25"

	// ConsoleServerPortTypeValueRj11 captures enum value "rj-11"
	ConsoleServerPortTypeValueRj11 string = "rj-11"

	// ConsoleServerPortTypeValueRj12 captures enum value "rj-12"
	ConsoleServerPortTypeValueRj12 string = "rj-12"

	// ConsoleServerPortTypeValueRj45 captures enum value "rj-45"
	ConsoleServerPortTypeValueRj45 string = "rj-45"

	// ConsoleServerPortTypeValueUsba captures enum value "usb-a"
	ConsoleServerPortTypeValueUsba string = "usb-a"

	// ConsoleServerPortTypeValueUsbb captures enum value "usb-b"
	ConsoleServerPortTypeValueUsbb string = "usb-b"

	// ConsoleServerPortTypeValueUsbc captures enum value "usb-c"
	ConsoleServerPortTypeValueUsbc string = "usb-c"

	// ConsoleServerPortTypeValueUsbMinia captures enum value "usb-mini-a"
	ConsoleServerPortTypeValueUsbMinia string = "usb-mini-a"

	// ConsoleServerPortTypeValueUsbMinib captures enum value "usb-mini-b"
	ConsoleServerPortTypeValueUsbMinib string = "usb-mini-b"

	// ConsoleServerPortTypeValueUsbMicroa captures enum value "usb-micro-a"
	ConsoleServerPortTypeValueUsbMicroa string = "usb-micro-a"

	// ConsoleServerPortTypeValueUsbMicrob captures enum value "usb-micro-b"
	ConsoleServerPortTypeValueUsbMicrob string = "usb-micro-b"

	// ConsoleServerPortTypeValueOther captures enum value "other"
	ConsoleServerPortTypeValueOther string = "other"
)

// prop value enum
func (m *ConsoleServerPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console server port type based on context it is used
func (m *ConsoleServerPortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortType) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
