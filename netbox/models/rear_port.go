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

// RearPort rear port
//
// swagger:model RearPort
type RearPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

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
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Positions
	// Maximum: 64
	// Minimum: 1
	Positions int64 `json:"positions,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// type
	// Required: true
	Type *RearPortType `json:"type"`
}

// Validate validates this rear port
func (m *RearPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
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

	if err := m.validatePositions(formats); err != nil {
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

func (m *RearPort) validateCable(formats strfmt.Registry) error {

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

func (m *RearPort) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateDevice(formats strfmt.Registry) error {

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

func (m *RearPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validatePositions(formats strfmt.Registry) error {

	if swag.IsZero(m.Positions) { // not required
		return nil
	}

	if err := validate.MinimumInt("positions", "body", int64(m.Positions), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("positions", "body", int64(m.Positions), 64, false); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateTags(formats strfmt.Registry) error {

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

func (m *RearPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
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

// ContextValidate validate this rear port based on the context it is used
func (m *RearPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCable(ctx, formats); err != nil {
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

func (m *RearPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RearPort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RearPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RearPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPort) UnmarshalBinary(b []byte) error {
	var res RearPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RearPortType Type
//
// swagger:model RearPortType
type RearPortType struct {

	// label
	// Required: true
	// Enum: [8P8C 110 Punch BNC MRJ21 FC LC LC/APC LSH LSH/APC MPO MTRJ SC SC/APC ST]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st]
	Value *string `json:"value"`
}

// Validate validates this rear port type
func (m *RearPortType) Validate(formats strfmt.Registry) error {
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

var rearPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","110 Punch","BNC","MRJ21","FC","LC","LC/APC","LSH","LSH/APC","MPO","MTRJ","SC","SC/APC","ST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTypeTypeLabelPropEnum = append(rearPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// RearPortTypeLabelNr8P8C captures enum value "8P8C"
	RearPortTypeLabelNr8P8C string = "8P8C"

	// RearPortTypeLabelNr110Punch captures enum value "110 Punch"
	RearPortTypeLabelNr110Punch string = "110 Punch"

	// RearPortTypeLabelBNC captures enum value "BNC"
	RearPortTypeLabelBNC string = "BNC"

	// RearPortTypeLabelMRJ21 captures enum value "MRJ21"
	RearPortTypeLabelMRJ21 string = "MRJ21"

	// RearPortTypeLabelFC captures enum value "FC"
	RearPortTypeLabelFC string = "FC"

	// RearPortTypeLabelLC captures enum value "LC"
	RearPortTypeLabelLC string = "LC"

	// RearPortTypeLabelLCAPC captures enum value "LC/APC"
	RearPortTypeLabelLCAPC string = "LC/APC"

	// RearPortTypeLabelLSH captures enum value "LSH"
	RearPortTypeLabelLSH string = "LSH"

	// RearPortTypeLabelLSHAPC captures enum value "LSH/APC"
	RearPortTypeLabelLSHAPC string = "LSH/APC"

	// RearPortTypeLabelMPO captures enum value "MPO"
	RearPortTypeLabelMPO string = "MPO"

	// RearPortTypeLabelMTRJ captures enum value "MTRJ"
	RearPortTypeLabelMTRJ string = "MTRJ"

	// RearPortTypeLabelSC captures enum value "SC"
	RearPortTypeLabelSC string = "SC"

	// RearPortTypeLabelSCAPC captures enum value "SC/APC"
	RearPortTypeLabelSCAPC string = "SC/APC"

	// RearPortTypeLabelST captures enum value "ST"
	RearPortTypeLabelST string = "ST"
)

// prop value enum
func (m *RearPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rearPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTypeTypeValuePropEnum = append(rearPortTypeTypeValuePropEnum, v)
	}
}

const (

	// RearPortTypeValueNr8p8c captures enum value "8p8c"
	RearPortTypeValueNr8p8c string = "8p8c"

	// RearPortTypeValueNr110Punch captures enum value "110-punch"
	RearPortTypeValueNr110Punch string = "110-punch"

	// RearPortTypeValueBnc captures enum value "bnc"
	RearPortTypeValueBnc string = "bnc"

	// RearPortTypeValueMrj21 captures enum value "mrj21"
	RearPortTypeValueMrj21 string = "mrj21"

	// RearPortTypeValueFc captures enum value "fc"
	RearPortTypeValueFc string = "fc"

	// RearPortTypeValueLc captures enum value "lc"
	RearPortTypeValueLc string = "lc"

	// RearPortTypeValueLcApc captures enum value "lc-apc"
	RearPortTypeValueLcApc string = "lc-apc"

	// RearPortTypeValueLsh captures enum value "lsh"
	RearPortTypeValueLsh string = "lsh"

	// RearPortTypeValueLshApc captures enum value "lsh-apc"
	RearPortTypeValueLshApc string = "lsh-apc"

	// RearPortTypeValueMpo captures enum value "mpo"
	RearPortTypeValueMpo string = "mpo"

	// RearPortTypeValueMtrj captures enum value "mtrj"
	RearPortTypeValueMtrj string = "mtrj"

	// RearPortTypeValueSc captures enum value "sc"
	RearPortTypeValueSc string = "sc"

	// RearPortTypeValueScApc captures enum value "sc-apc"
	RearPortTypeValueScApc string = "sc-apc"

	// RearPortTypeValueSt captures enum value "st"
	RearPortTypeValueSt string = "st"
)

// prop value enum
func (m *RearPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rear port type based on context it is used
func (m *RearPortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RearPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPortType) UnmarshalBinary(b []byte) error {
	var res RearPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
