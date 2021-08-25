/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Auth0Auth Authentication code and redirect URI for authentication
type Auth0Auth struct {
	// OAuth authentication code
	Code string `json:"code"`
	// OAuth redirect URI
	RedirectUri string `json:"redirect_uri"`
}

// NewAuth0Auth instantiates a new Auth0Auth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuth0Auth(code string, redirectUri string) *Auth0Auth {
	this := Auth0Auth{}
	this.Code = code
	this.RedirectUri = redirectUri
	return &this
}

// NewAuth0AuthWithDefaults instantiates a new Auth0Auth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuth0AuthWithDefaults() *Auth0Auth {
	this := Auth0Auth{}
	return &this
}

// GetCode returns the Code field value
func (o *Auth0Auth) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Auth0Auth) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Auth0Auth) SetCode(v string) {
	o.Code = v
}

// GetRedirectUri returns the RedirectUri field value
func (o *Auth0Auth) GetRedirectUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value
// and a boolean to check if the value has been set.
func (o *Auth0Auth) GetRedirectUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RedirectUri, true
}

// SetRedirectUri sets field value
func (o *Auth0Auth) SetRedirectUri(v string) {
	o.RedirectUri = v
}

func (o Auth0Auth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	return json.Marshal(toSerialize)
}

type NullableAuth0Auth struct {
	value *Auth0Auth
	isSet bool
}

func (v NullableAuth0Auth) Get() *Auth0Auth {
	return v.value
}

func (v *NullableAuth0Auth) Set(val *Auth0Auth) {
	v.value = val
	v.isSet = true
}

func (v NullableAuth0Auth) IsSet() bool {
	return v.isSet
}

func (v *NullableAuth0Auth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuth0Auth(val *Auth0Auth) *NullableAuth0Auth {
	return &NullableAuth0Auth{value: val, isSet: true}
}

func (v NullableAuth0Auth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuth0Auth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

