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

// PutAccountContactsRequest struct for PutAccountContactsRequest
type PutAccountContactsRequest struct {
	// The contact's email
	Email string `json:"email"`
	// The contact's name
	Name string `json:"name"`
	// The contact's phone
	Phone string `json:"phone"`
}

// NewPutAccountContactsRequest instantiates a new PutAccountContactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutAccountContactsRequest(email string, name string, phone string) *PutAccountContactsRequest {
	this := PutAccountContactsRequest{}
	this.Email = email
	this.Name = name
	this.Phone = phone
	return &this
}

// NewPutAccountContactsRequestWithDefaults instantiates a new PutAccountContactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutAccountContactsRequestWithDefaults() *PutAccountContactsRequest {
	this := PutAccountContactsRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *PutAccountContactsRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PutAccountContactsRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PutAccountContactsRequest) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *PutAccountContactsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PutAccountContactsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PutAccountContactsRequest) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value
func (o *PutAccountContactsRequest) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *PutAccountContactsRequest) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *PutAccountContactsRequest) SetPhone(v string) {
	o.Phone = v
}

func (o PutAccountContactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullablePutAccountContactsRequest struct {
	value *PutAccountContactsRequest
	isSet bool
}

func (v NullablePutAccountContactsRequest) Get() *PutAccountContactsRequest {
	return v.value
}

func (v *NullablePutAccountContactsRequest) Set(val *PutAccountContactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutAccountContactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutAccountContactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutAccountContactsRequest(val *PutAccountContactsRequest) *NullablePutAccountContactsRequest {
	return &NullablePutAccountContactsRequest{value: val, isSet: true}
}

func (v NullablePutAccountContactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutAccountContactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
