/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserDigestSettings User digest settings
type UserDigestSettings struct {
	// Receive digest
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUserDigestSettings instantiates a new UserDigestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDigestSettings() *UserDigestSettings {
	this := UserDigestSettings{}
	return &this
}

// NewUserDigestSettingsWithDefaults instantiates a new UserDigestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDigestSettingsWithDefaults() *UserDigestSettings {
	this := UserDigestSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UserDigestSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDigestSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UserDigestSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UserDigestSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UserDigestSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableUserDigestSettings struct {
	value *UserDigestSettings
	isSet bool
}

func (v NullableUserDigestSettings) Get() *UserDigestSettings {
	return v.value
}

func (v *NullableUserDigestSettings) Set(val *UserDigestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDigestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDigestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDigestSettings(val *UserDigestSettings) *NullableUserDigestSettings {
	return &NullableUserDigestSettings{value: val, isSet: true}
}

func (v NullableUserDigestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDigestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
