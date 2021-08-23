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

// UserProfileAllOf A user's profile
type UserProfileAllOf struct {
	// User preferences
	Preferences map[string]interface{} `json:"preferences"`
}

// NewUserProfileAllOf instantiates a new UserProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfileAllOf(preferences map[string]interface{}) *UserProfileAllOf {
	this := UserProfileAllOf{}
	this.Preferences = preferences
	return &this
}

// NewUserProfileAllOfWithDefaults instantiates a new UserProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileAllOfWithDefaults() *UserProfileAllOf {
	this := UserProfileAllOf{}
	return &this
}

// GetPreferences returns the Preferences field value
func (o *UserProfileAllOf) GetPreferences() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value
// and a boolean to check if the value has been set.
func (o *UserProfileAllOf) GetPreferencesOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Preferences, true
}

// SetPreferences sets field value
func (o *UserProfileAllOf) SetPreferences(v map[string]interface{}) {
	o.Preferences = v
}

func (o UserProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["preferences"] = o.Preferences
	}
	return json.Marshal(toSerialize)
}

type NullableUserProfileAllOf struct {
	value *UserProfileAllOf
	isSet bool
}

func (v NullableUserProfileAllOf) Get() *UserProfileAllOf {
	return v.value
}

func (v *NullableUserProfileAllOf) Set(val *UserProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfileAllOf(val *UserProfileAllOf) *NullableUserProfileAllOf {
	return &NullableUserProfileAllOf{value: val, isSet: true}
}

func (v NullableUserProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

