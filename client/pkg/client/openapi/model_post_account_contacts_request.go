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

// PostAccountContactsRequest struct for PostAccountContactsRequest
type PostAccountContactsRequest struct {
	// The contact's email
	Email string `json:"email"`
	// The contact's name
	Name string `json:"name"`
	// The contact's phone
	Phone string `json:"phone"`
	Type  string `json:"type"`
}

// NewPostAccountContactsRequest instantiates a new PostAccountContactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAccountContactsRequest(email string, name string, phone string, type_ string) *PostAccountContactsRequest {
	this := PostAccountContactsRequest{}
	this.Email = email
	this.Name = name
	this.Phone = phone
	this.Type = type_
	return &this
}

// NewPostAccountContactsRequestWithDefaults instantiates a new PostAccountContactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAccountContactsRequestWithDefaults() *PostAccountContactsRequest {
	this := PostAccountContactsRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *PostAccountContactsRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PostAccountContactsRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PostAccountContactsRequest) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *PostAccountContactsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostAccountContactsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostAccountContactsRequest) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value
func (o *PostAccountContactsRequest) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *PostAccountContactsRequest) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *PostAccountContactsRequest) SetPhone(v string) {
	o.Phone = v
}

// GetType returns the Type field value
func (o *PostAccountContactsRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostAccountContactsRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostAccountContactsRequest) SetType(v string) {
	o.Type = v
}

func (o PostAccountContactsRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePostAccountContactsRequest struct {
	value *PostAccountContactsRequest
	isSet bool
}

func (v NullablePostAccountContactsRequest) Get() *PostAccountContactsRequest {
	return v.value
}

func (v *NullablePostAccountContactsRequest) Set(val *PostAccountContactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAccountContactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAccountContactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAccountContactsRequest(val *PostAccountContactsRequest) *NullablePostAccountContactsRequest {
	return &NullablePostAccountContactsRequest{value: val, isSet: true}
}

func (v NullablePostAccountContactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAccountContactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
