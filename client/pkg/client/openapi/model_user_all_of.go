/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// UserAllOf An account user
type UserAllOf struct {
	// Timestamp of the last activity of this user
	LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
	// The roles that this user has been assigned
	Roles []RoleSummary `json:"roles"`
}

// NewUserAllOf instantiates a new UserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAllOf(roles []RoleSummary) *UserAllOf {
	this := UserAllOf{}
	this.Roles = roles
	return &this
}

// NewUserAllOfWithDefaults instantiates a new UserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAllOfWithDefaults() *UserAllOf {
	this := UserAllOf{}
	return &this
}

// GetLastActivityAt returns the LastActivityAt field value if set, zero value otherwise.
func (o *UserAllOf) GetLastActivityAt() time.Time {
	if o == nil || o.LastActivityAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivityAt
}

// GetLastActivityAtOk returns a tuple with the LastActivityAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAllOf) GetLastActivityAtOk() (*time.Time, bool) {
	if o == nil || o.LastActivityAt == nil {
		return nil, false
	}
	return o.LastActivityAt, true
}

// HasLastActivityAt returns a boolean if a field has been set.
func (o *UserAllOf) HasLastActivityAt() bool {
	if o != nil && o.LastActivityAt != nil {
		return true
	}

	return false
}

// SetLastActivityAt gets a reference to the given time.Time and assigns it to the LastActivityAt field.
func (o *UserAllOf) SetLastActivityAt(v time.Time) {
	o.LastActivityAt = &v
}

// GetRoles returns the Roles field value
func (o *UserAllOf) GetRoles() []RoleSummary {
	if o == nil {
		var ret []RoleSummary
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserAllOf) GetRolesOk() (*[]RoleSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *UserAllOf) SetRoles(v []RoleSummary) {
	o.Roles = v
}

func (o UserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastActivityAt != nil {
		toSerialize["last_activity_at"] = o.LastActivityAt
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableUserAllOf struct {
	value *UserAllOf
	isSet bool
}

func (v NullableUserAllOf) Get() *UserAllOf {
	return v.value
}

func (v *NullableUserAllOf) Set(val *UserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAllOf(val *UserAllOf) *NullableUserAllOf {
	return &NullableUserAllOf{value: val, isSet: true}
}

func (v NullableUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
