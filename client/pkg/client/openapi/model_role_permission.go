/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RolePermission A permission as assigned to a role
type RolePermission struct {
	Permission PermissionSummary `json:"permission"`
}

// NewRolePermission instantiates a new RolePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermission(permission PermissionSummary) *RolePermission {
	this := RolePermission{}
	this.Permission = permission
	return &this
}

// NewRolePermissionWithDefaults instantiates a new RolePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionWithDefaults() *RolePermission {
	this := RolePermission{}
	return &this
}

// GetPermission returns the Permission field value
func (o *RolePermission) GetPermission() PermissionSummary {
	if o == nil {
		var ret PermissionSummary
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *RolePermission) GetPermissionOk() (*PermissionSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *RolePermission) SetPermission(v PermissionSummary) {
	o.Permission = v
}

func (o RolePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermission struct {
	value *RolePermission
	isSet bool
}

func (v NullableRolePermission) Get() *RolePermission {
	return v.value
}

func (v *NullableRolePermission) Set(val *RolePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermission(val *RolePermission) *NullableRolePermission {
	return &NullableRolePermission{value: val, isSet: true}
}

func (v NullableRolePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

