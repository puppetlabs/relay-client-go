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

// OAuth2ConnectionAuthInput Connection authentication values using the OAuth 2.0 protocol
type OAuth2ConnectionAuthInput struct {
	// The code generated by the successful authentication code request
	Code string `json:"code"`
	// The state information returned in the callback URL
	State string `json:"state"`
	// The type of authentication provided by the connection
	Type string `json:"type"`
}

// NewOAuth2ConnectionAuthInput instantiates a new OAuth2ConnectionAuthInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ConnectionAuthInput(code string, state string, type_ string) *OAuth2ConnectionAuthInput {
	this := OAuth2ConnectionAuthInput{}
	this.Code = code
	this.State = state
	this.Type = type_
	return &this
}

// NewOAuth2ConnectionAuthInputWithDefaults instantiates a new OAuth2ConnectionAuthInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ConnectionAuthInputWithDefaults() *OAuth2ConnectionAuthInput {
	this := OAuth2ConnectionAuthInput{}
	return &this
}

// GetCode returns the Code field value
func (o *OAuth2ConnectionAuthInput) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionAuthInput) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OAuth2ConnectionAuthInput) SetCode(v string) {
	o.Code = v
}

// GetState returns the State field value
func (o *OAuth2ConnectionAuthInput) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionAuthInput) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *OAuth2ConnectionAuthInput) SetState(v string) {
	o.State = v
}

// GetType returns the Type field value
func (o *OAuth2ConnectionAuthInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionAuthInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OAuth2ConnectionAuthInput) SetType(v string) {
	o.Type = v
}

func (o OAuth2ConnectionAuthInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2ConnectionAuthInput struct {
	value *OAuth2ConnectionAuthInput
	isSet bool
}

func (v NullableOAuth2ConnectionAuthInput) Get() *OAuth2ConnectionAuthInput {
	return v.value
}

func (v *NullableOAuth2ConnectionAuthInput) Set(val *OAuth2ConnectionAuthInput) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ConnectionAuthInput) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ConnectionAuthInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ConnectionAuthInput(val *OAuth2ConnectionAuthInput) *NullableOAuth2ConnectionAuthInput {
	return &NullableOAuth2ConnectionAuthInput{value: val, isSet: true}
}

func (v NullableOAuth2ConnectionAuthInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ConnectionAuthInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
