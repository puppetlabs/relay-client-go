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

// NotificationIdentifier A unique identifier for a notification
type NotificationIdentifier struct {
	// The unique identifier of this notification
	Id string `json:"id"`
}

// NewNotificationIdentifier instantiates a new NotificationIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationIdentifier(id string) *NotificationIdentifier {
	this := NotificationIdentifier{}
	this.Id = id
	return &this
}

// NewNotificationIdentifierWithDefaults instantiates a new NotificationIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationIdentifierWithDefaults() *NotificationIdentifier {
	this := NotificationIdentifier{}
	return &this
}

// GetId returns the Id field value
func (o *NotificationIdentifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotificationIdentifier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotificationIdentifier) SetId(v string) {
	o.Id = v
}

func (o NotificationIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationIdentifier struct {
	value *NotificationIdentifier
	isSet bool
}

func (v NullableNotificationIdentifier) Get() *NotificationIdentifier {
	return v.value
}

func (v *NullableNotificationIdentifier) Set(val *NotificationIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationIdentifier(val *NotificationIdentifier) *NullableNotificationIdentifier {
	return &NullableNotificationIdentifier{value: val, isSet: true}
}

func (v NullableNotificationIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
