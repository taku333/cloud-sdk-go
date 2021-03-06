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

// ActiveDirectorySecurityRealmLoadBalance Elasticsearch Security Active Directory load balancing configuration
// swagger:model ActiveDirectorySecurityRealmLoadBalance
type ActiveDirectorySecurityRealmLoadBalance struct {

	// When using dns_failover or dns_round_robin as the load balancing type, this setting controls the amount of time to cache DNS lookups. Defaults to 1h.
	CacheTTL string `json:"cache_ttl,omitempty"`

	// The behavior to use when there are multiple Active Directory URLs defined
	// Required: true
	// Enum: [failover dns_failover round_robin dns_round_robin]
	Type *string `json:"type"`
}

// Validate validates this active directory security realm load balance
func (m *ActiveDirectorySecurityRealmLoadBalance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var activeDirectorySecurityRealmLoadBalanceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["failover","dns_failover","round_robin","dns_round_robin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activeDirectorySecurityRealmLoadBalanceTypeTypePropEnum = append(activeDirectorySecurityRealmLoadBalanceTypeTypePropEnum, v)
	}
}

const (

	// ActiveDirectorySecurityRealmLoadBalanceTypeFailover captures enum value "failover"
	ActiveDirectorySecurityRealmLoadBalanceTypeFailover string = "failover"

	// ActiveDirectorySecurityRealmLoadBalanceTypeDNSFailover captures enum value "dns_failover"
	ActiveDirectorySecurityRealmLoadBalanceTypeDNSFailover string = "dns_failover"

	// ActiveDirectorySecurityRealmLoadBalanceTypeRoundRobin captures enum value "round_robin"
	ActiveDirectorySecurityRealmLoadBalanceTypeRoundRobin string = "round_robin"

	// ActiveDirectorySecurityRealmLoadBalanceTypeDNSRoundRobin captures enum value "dns_round_robin"
	ActiveDirectorySecurityRealmLoadBalanceTypeDNSRoundRobin string = "dns_round_robin"
)

// prop value enum
func (m *ActiveDirectorySecurityRealmLoadBalance) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, activeDirectorySecurityRealmLoadBalanceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ActiveDirectorySecurityRealmLoadBalance) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectorySecurityRealmLoadBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectorySecurityRealmLoadBalance) UnmarshalBinary(b []byte) error {
	var res ActiveDirectorySecurityRealmLoadBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
