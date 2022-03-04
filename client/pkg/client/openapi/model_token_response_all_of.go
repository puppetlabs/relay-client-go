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

// TokenResponseAllOf The response for a newly created token
type TokenResponseAllOf struct {
	Token *TokenWithSecret `json:"token,omitempty"`
}

// NewTokenResponseAllOf instantiates a new TokenResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponseAllOf() *TokenResponseAllOf {
	this := TokenResponseAllOf{}
	return &this
}

// NewTokenResponseAllOfWithDefaults instantiates a new TokenResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseAllOfWithDefaults() *TokenResponseAllOf {
	this := TokenResponseAllOf{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetToken() TokenWithSecret {
	if o == nil || o.Token == nil {
		var ret TokenWithSecret
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetTokenOk() (*TokenWithSecret, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given TokenWithSecret and assigns it to the Token field.
func (o *TokenResponseAllOf) SetToken(v TokenWithSecret) {
	o.Token = &v
}

func (o TokenResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableTokenResponseAllOf struct {
	value *TokenResponseAllOf
	isSet bool
}

func (v NullableTokenResponseAllOf) Get() *TokenResponseAllOf {
	return v.value
}

func (v *NullableTokenResponseAllOf) Set(val *TokenResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponseAllOf(val *TokenResponseAllOf) *NullableTokenResponseAllOf {
	return &NullableTokenResponseAllOf{value: val, isSet: true}
}

func (v NullableTokenResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
