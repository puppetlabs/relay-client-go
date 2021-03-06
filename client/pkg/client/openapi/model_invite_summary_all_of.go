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

// InviteSummaryAllOf A summary of an invite
type InviteSummaryAllOf struct {
	// The email address of the invitee
	Email string `json:"email"`
	// The full name of the invitee
	Name string `json:"name"`
}

// NewInviteSummaryAllOf instantiates a new InviteSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteSummaryAllOf(email string, name string) *InviteSummaryAllOf {
	this := InviteSummaryAllOf{}
	this.Email = email
	this.Name = name
	return &this
}

// NewInviteSummaryAllOfWithDefaults instantiates a new InviteSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteSummaryAllOfWithDefaults() *InviteSummaryAllOf {
	this := InviteSummaryAllOf{}
	return &this
}

// GetEmail returns the Email field value
func (o *InviteSummaryAllOf) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InviteSummaryAllOf) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InviteSummaryAllOf) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *InviteSummaryAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InviteSummaryAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InviteSummaryAllOf) SetName(v string) {
	o.Name = v
}

func (o InviteSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInviteSummaryAllOf struct {
	value *InviteSummaryAllOf
	isSet bool
}

func (v NullableInviteSummaryAllOf) Get() *InviteSummaryAllOf {
	return v.value
}

func (v *NullableInviteSummaryAllOf) Set(val *InviteSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteSummaryAllOf(val *InviteSummaryAllOf) *NullableInviteSummaryAllOf {
	return &NullableInviteSummaryAllOf{value: val, isSet: true}
}

func (v NullableInviteSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
