// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActiveDirectoryGroupSearch Elasticsearch Security Active Directory realm group search configuration
// swagger:model ActiveDirectoryGroupSearch
type ActiveDirectoryGroupSearch struct {

	// Specifies a container DN to search for groups in which the user has membership
	BaseDn string `json:"base_dn,omitempty"`

	// Specifies whether the group search should be sub_tree, one_level or base. one_level only searches objects directly contained within the base_dn. The default sub_tree searches all objects contained under base_dn. base specifies that the base_dn is a group object, and that it is the only group considered.
	// Enum: [sub_tree one_level base]
	Scope string `json:"scope,omitempty"`
}

// Validate validates this active directory group search
func (m *ActiveDirectoryGroupSearch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var activeDirectoryGroupSearchTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sub_tree","one_level","base"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activeDirectoryGroupSearchTypeScopePropEnum = append(activeDirectoryGroupSearchTypeScopePropEnum, v)
	}
}

const (

	// ActiveDirectoryGroupSearchScopeSubTree captures enum value "sub_tree"
	ActiveDirectoryGroupSearchScopeSubTree string = "sub_tree"

	// ActiveDirectoryGroupSearchScopeOneLevel captures enum value "one_level"
	ActiveDirectoryGroupSearchScopeOneLevel string = "one_level"

	// ActiveDirectoryGroupSearchScopeBase captures enum value "base"
	ActiveDirectoryGroupSearchScopeBase string = "base"
)

// prop value enum
func (m *ActiveDirectoryGroupSearch) validateScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, activeDirectoryGroupSearchTypeScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ActiveDirectoryGroupSearch) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectoryGroupSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryGroupSearch) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryGroupSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
