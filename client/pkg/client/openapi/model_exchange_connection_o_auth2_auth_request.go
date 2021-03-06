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

// ExchangeConnectionOAuth2AuthRequest Fields to exchange for an OAuth 2.0 authenticated connection
type ExchangeConnectionOAuth2AuthRequest struct {
	Auth OAuth2ConnectionAuthInput `json:"auth"`
}

// NewExchangeConnectionOAuth2AuthRequest instantiates a new ExchangeConnectionOAuth2AuthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeConnectionOAuth2AuthRequest(auth OAuth2ConnectionAuthInput) *ExchangeConnectionOAuth2AuthRequest {
	this := ExchangeConnectionOAuth2AuthRequest{}
	this.Auth = auth
	return &this
}

// NewExchangeConnectionOAuth2AuthRequestWithDefaults instantiates a new ExchangeConnectionOAuth2AuthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeConnectionOAuth2AuthRequestWithDefaults() *ExchangeConnectionOAuth2AuthRequest {
	this := ExchangeConnectionOAuth2AuthRequest{}
	return &this
}

// GetAuth returns the Auth field value
func (o *ExchangeConnectionOAuth2AuthRequest) GetAuth() OAuth2ConnectionAuthInput {
	if o == nil {
		var ret OAuth2ConnectionAuthInput
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ExchangeConnectionOAuth2AuthRequest) GetAuthOk() (*OAuth2ConnectionAuthInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *ExchangeConnectionOAuth2AuthRequest) SetAuth(v OAuth2ConnectionAuthInput) {
	o.Auth = v
}

func (o ExchangeConnectionOAuth2AuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auth"] = o.Auth
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeConnectionOAuth2AuthRequest struct {
	value *ExchangeConnectionOAuth2AuthRequest
	isSet bool
}

func (v NullableExchangeConnectionOAuth2AuthRequest) Get() *ExchangeConnectionOAuth2AuthRequest {
	return v.value
}

func (v *NullableExchangeConnectionOAuth2AuthRequest) Set(val *ExchangeConnectionOAuth2AuthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeConnectionOAuth2AuthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeConnectionOAuth2AuthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeConnectionOAuth2AuthRequest(val *ExchangeConnectionOAuth2AuthRequest) *NullableExchangeConnectionOAuth2AuthRequest {
	return &NullableExchangeConnectionOAuth2AuthRequest{value: val, isSet: true}
}

func (v NullableExchangeConnectionOAuth2AuthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeConnectionOAuth2AuthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
