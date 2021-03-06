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

// InviteIdentifier A unique identifier for an invite
type InviteIdentifier struct {
	// The unique identifier of this invite
	Id string `json:"id"`
}

// NewInviteIdentifier instantiates a new InviteIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteIdentifier(id string) *InviteIdentifier {
	this := InviteIdentifier{}
	this.Id = id
	return &this
}

// NewInviteIdentifierWithDefaults instantiates a new InviteIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteIdentifierWithDefaults() *InviteIdentifier {
	this := InviteIdentifier{}
	return &this
}

// GetId returns the Id field value
func (o *InviteIdentifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InviteIdentifier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InviteIdentifier) SetId(v string) {
	o.Id = v
}

func (o InviteIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInviteIdentifier struct {
	value *InviteIdentifier
	isSet bool
}

func (v NullableInviteIdentifier) Get() *InviteIdentifier {
	return v.value
}

func (v *NullableInviteIdentifier) Set(val *InviteIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteIdentifier(val *InviteIdentifier) *NullableInviteIdentifier {
	return &NullableInviteIdentifier{value: val, isSet: true}
}

func (v NullableInviteIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
