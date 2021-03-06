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

// OAuth2ConnectionAuthStatus struct for OAuth2ConnectionAuthStatus
type OAuth2ConnectionAuthStatus struct {
	// The scopes available to this connection
	Scopes []string `json:"scopes,omitempty"`
	// The authentication mechanism for this connection
	Type string `json:"type"`
}

// NewOAuth2ConnectionAuthStatus instantiates a new OAuth2ConnectionAuthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ConnectionAuthStatus(type_ string) *OAuth2ConnectionAuthStatus {
	this := OAuth2ConnectionAuthStatus{}
	this.Type = type_
	return &this
}

// NewOAuth2ConnectionAuthStatusWithDefaults instantiates a new OAuth2ConnectionAuthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ConnectionAuthStatusWithDefaults() *OAuth2ConnectionAuthStatus {
	this := OAuth2ConnectionAuthStatus{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuth2ConnectionAuthStatus) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionAuthStatus) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuth2ConnectionAuthStatus) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAuth2ConnectionAuthStatus) SetScopes(v []string) {
	o.Scopes = v
}

// GetType returns the Type field value
func (o *OAuth2ConnectionAuthStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionAuthStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OAuth2ConnectionAuthStatus) SetType(v string) {
	o.Type = v
}

func (o OAuth2ConnectionAuthStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2ConnectionAuthStatus struct {
	value *OAuth2ConnectionAuthStatus
	isSet bool
}

func (v NullableOAuth2ConnectionAuthStatus) Get() *OAuth2ConnectionAuthStatus {
	return v.value
}

func (v *NullableOAuth2ConnectionAuthStatus) Set(val *OAuth2ConnectionAuthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ConnectionAuthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ConnectionAuthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ConnectionAuthStatus(val *OAuth2ConnectionAuthStatus) *NullableOAuth2ConnectionAuthStatus {
	return &NullableOAuth2ConnectionAuthStatus{value: val, isSet: true}
}

func (v NullableOAuth2ConnectionAuthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ConnectionAuthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
