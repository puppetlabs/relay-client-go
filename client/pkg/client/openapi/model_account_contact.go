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

// AccountContact Account contact
type AccountContact struct {
	Email string `json:"email"`
	// The unique identifier for the account contact
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Type  string `json:"type"`
}

// NewAccountContact instantiates a new AccountContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountContact(email string, id string, name string, phone string, type_ string) *AccountContact {
	this := AccountContact{}
	this.Email = email
	this.Id = id
	this.Name = name
	this.Phone = phone
	this.Type = type_
	return &this
}

// NewAccountContactWithDefaults instantiates a new AccountContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountContactWithDefaults() *AccountContact {
	this := AccountContact{}
	return &this
}

// GetEmail returns the Email field value
func (o *AccountContact) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AccountContact) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AccountContact) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *AccountContact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountContact) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountContact) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccountContact) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountContact) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountContact) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value
func (o *AccountContact) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *AccountContact) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *AccountContact) SetPhone(v string) {
	o.Phone = v
}

// GetType returns the Type field value
func (o *AccountContact) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountContact) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountContact) SetType(v string) {
	o.Type = v
}

func (o AccountContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAccountContact struct {
	value *AccountContact
	isSet bool
}

func (v NullableAccountContact) Get() *AccountContact {
	return v.value
}

func (v *NullableAccountContact) Set(val *AccountContact) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountContact) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountContact(val *AccountContact) *NullableAccountContact {
	return &NullableAccountContact{value: val, isSet: true}
}

func (v NullableAccountContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
