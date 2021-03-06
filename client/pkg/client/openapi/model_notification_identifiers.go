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

// NotificationIdentifiers A list of notifications
type NotificationIdentifiers struct {
	// The list of notification IDs
	Ids []string `json:"ids,omitempty"`
}

// NewNotificationIdentifiers instantiates a new NotificationIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationIdentifiers() *NotificationIdentifiers {
	this := NotificationIdentifiers{}
	return &this
}

// NewNotificationIdentifiersWithDefaults instantiates a new NotificationIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationIdentifiersWithDefaults() *NotificationIdentifiers {
	this := NotificationIdentifiers{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *NotificationIdentifiers) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationIdentifiers) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *NotificationIdentifiers) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *NotificationIdentifiers) SetIds(v []string) {
	o.Ids = v
}

func (o NotificationIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationIdentifiers struct {
	value *NotificationIdentifiers
	isSet bool
}

func (v NullableNotificationIdentifiers) Get() *NotificationIdentifiers {
	return v.value
}

func (v *NullableNotificationIdentifiers) Set(val *NotificationIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationIdentifiers(val *NotificationIdentifiers) *NullableNotificationIdentifiers {
	return &NullableNotificationIdentifiers{value: val, isSet: true}
}

func (v NullableNotificationIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
