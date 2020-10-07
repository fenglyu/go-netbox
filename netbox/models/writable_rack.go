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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableRack writable rack
//
// swagger:model WritableRack
type WritableRack struct {

	// Asset tag
	//
	// A unique tag used to identify this rack
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Descending units
	//
	// Units are numbered top-to-bottom
	DescUnits *bool `json:"desc_units,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// Facility ID
	//
	// Locally-assigned identifier
	// Max Length: 50
	FacilityID *string `json:"facility_id,omitempty"`

	// Group
	//
	// Assigned group
	Group *int64 `json:"group,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Outer depth
	//
	// Outer dimension of rack (depth)
	// Maximum: 32767
	// Minimum: 0
	OuterDepth *int64 `json:"outer_depth,omitempty"`

	// Outer unit
	// Enum: [mm in]
	OuterUnit string `json:"outer_unit,omitempty"`

	// Outer width
	//
	// Outer dimension of rack (width)
	// Maximum: 32767
	// Minimum: 0
	OuterWidth *int64 `json:"outer_width,omitempty"`

	// Powerfeed count
	// Read Only: true
	PowerfeedCount int64 `json:"powerfeed_count,omitempty"`

	// Role
	//
	// Functional role
	Role *int64 `json:"role,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// Site
	// Required: true
	Site *int64 `json:"site"`

	// Status
	// Enum: [reserved available planned active deprecated]
	Status string `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Type
	// Enum: [2-post-frame 4-post-frame 4-post-cabinet wall-frame wall-cabinet]
	Type string `json:"type,omitempty"`

	// Height (U)
	//
	// Height in rack units
	// Maximum: 100
	// Minimum: 1
	UHeight int64 `json:"u_height,omitempty"`

	// Width
	//
	// Rail-to-rail width
	// Enum: [10 19 21 23]
	Width int64 `json:"width,omitempty"`
}

// Validate validates this writable rack
func (m *WritableRack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRack) validateAssetTag(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", string(*m.AssetTag), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateFacilityID(formats strfmt.Registry) error {

	if swag.IsZero(m.FacilityID) { // not required
		return nil
	}

	if err := validate.MaxLength("facility_id", "body", string(*m.FacilityID), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateName(formats strfmt.Registry) error {

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

func (m *WritableRack) validateOuterDepth(formats strfmt.Registry) error {

	if swag.IsZero(m.OuterDepth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_depth", "body", int64(*m.OuterDepth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_depth", "body", int64(*m.OuterDepth), 32767, false); err != nil {
		return err
	}

	return nil
}

var writableRackTypeOuterUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mm","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeOuterUnitPropEnum = append(writableRackTypeOuterUnitPropEnum, v)
	}
}

const (

	// WritableRackOuterUnitMm captures enum value "mm"
	WritableRackOuterUnitMm string = "mm"

	// WritableRackOuterUnitIn captures enum value "in"
	WritableRackOuterUnitIn string = "in"
)

// prop value enum
func (m *WritableRack) validateOuterUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeOuterUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateOuterUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.OuterUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateOuterUnitEnum("outer_unit", "body", m.OuterUnit); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateOuterWidth(formats strfmt.Registry) error {

	if swag.IsZero(m.OuterWidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_width", "body", int64(*m.OuterWidth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_width", "body", int64(*m.OuterWidth), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateSerial(formats strfmt.Registry) error {

	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", string(m.Serial), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	return nil
}

var writableRackTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["reserved","available","planned","active","deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeStatusPropEnum = append(writableRackTypeStatusPropEnum, v)
	}
}

const (

	// WritableRackStatusReserved captures enum value "reserved"
	WritableRackStatusReserved string = "reserved"

	// WritableRackStatusAvailable captures enum value "available"
	WritableRackStatusAvailable string = "available"

	// WritableRackStatusPlanned captures enum value "planned"
	WritableRackStatusPlanned string = "planned"

	// WritableRackStatusActive captures enum value "active"
	WritableRackStatusActive string = "active"

	// WritableRackStatusDeprecated captures enum value "deprecated"
	WritableRackStatusDeprecated string = "deprecated"
)

// prop value enum
func (m *WritableRack) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateTags(formats strfmt.Registry) error {

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

var writableRackTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-cabinet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeTypePropEnum = append(writableRackTypeTypePropEnum, v)
	}
}

const (

	// WritableRackTypeNr2PostFrame captures enum value "2-post-frame"
	WritableRackTypeNr2PostFrame string = "2-post-frame"

	// WritableRackTypeNr4PostFrame captures enum value "4-post-frame"
	WritableRackTypeNr4PostFrame string = "4-post-frame"

	// WritableRackTypeNr4PostCabinet captures enum value "4-post-cabinet"
	WritableRackTypeNr4PostCabinet string = "4-post-cabinet"

	// WritableRackTypeWallFrame captures enum value "wall-frame"
	WritableRackTypeWallFrame string = "wall-frame"

	// WritableRackTypeWallCabinet captures enum value "wall-cabinet"
	WritableRackTypeWallCabinet string = "wall-cabinet"
)

// prop value enum
func (m *WritableRack) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateUHeight(formats strfmt.Registry) error {

	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", int64(m.UHeight), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", int64(m.UHeight), 100, false); err != nil {
		return err
	}

	return nil
}

var writableRackTypeWidthPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[10,19,21,23]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeWidthPropEnum = append(writableRackTypeWidthPropEnum, v)
	}
}

// prop value enum
func (m *WritableRack) validateWidthEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeWidthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateWidth(formats strfmt.Registry) error {

	if swag.IsZero(m.Width) { // not required
		return nil
	}

	// value enum
	if err := m.validateWidthEnum("width", "body", m.Width); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRack) UnmarshalBinary(b []byte) error {
	var res WritableRack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
