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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedPowerPort Power port
//
// swagger:model NestedPowerPort
type NestedPowerPort struct {

	// Cable
	Cable *int64 `json:"cable,omitempty"`

	// connection status
	ConnectionStatus *NestedPowerPortConnectionStatus `json:"connection_status,omitempty"`

	// device
	Device *NestedDevice `json:"device,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this nested power port
func (m *NestedPowerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedPowerPort) validateConnectionStatus(formats strfmt.Registry) error {

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

func (m *NestedPowerPort) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
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

func (m *NestedPowerPort) validateName(formats strfmt.Registry) error {

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

func (m *NestedPowerPort) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedPowerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedPowerPort) UnmarshalBinary(b []byte) error {
	var res NestedPowerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NestedPowerPortConnectionStatus Connection status
//
// swagger:model NestedPowerPortConnectionStatus
type NestedPowerPortConnectionStatus struct {

	// label
	// Required: true
	// Enum: [Not Connected Connected]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [false true]
	Value *bool `json:"value"`
}

// Validate validates this nested power port connection status
func (m *NestedPowerPortConnectionStatus) Validate(formats strfmt.Registry) error {
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

var nestedPowerPortConnectionStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Not Connected","Connected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nestedPowerPortConnectionStatusTypeLabelPropEnum = append(nestedPowerPortConnectionStatusTypeLabelPropEnum, v)
	}
}

const (

	// NestedPowerPortConnectionStatusLabelNotConnected captures enum value "Not Connected"
	NestedPowerPortConnectionStatusLabelNotConnected string = "Not Connected"

	// NestedPowerPortConnectionStatusLabelConnected captures enum value "Connected"
	NestedPowerPortConnectionStatusLabelConnected string = "Connected"
)

// prop value enum
func (m *NestedPowerPortConnectionStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nestedPowerPortConnectionStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NestedPowerPortConnectionStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("connection_status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var nestedPowerPortConnectionStatusTypeValuePropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nestedPowerPortConnectionStatusTypeValuePropEnum = append(nestedPowerPortConnectionStatusTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *NestedPowerPortConnectionStatus) validateValueEnum(path, location string, value bool) error {
	if err := validate.EnumCase(path, location, value, nestedPowerPortConnectionStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NestedPowerPortConnectionStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("connection_status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedPowerPortConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedPowerPortConnectionStatus) UnmarshalBinary(b []byte) error {
	var res NestedPowerPortConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
