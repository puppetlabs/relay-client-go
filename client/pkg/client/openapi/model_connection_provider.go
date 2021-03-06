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

// ConnectionProvider A description of a connection provider
type ConnectionProvider struct {
	// The supported authentication mechanisms for this connection provider
	Auth []ConnectionProviderAuth `json:"auth,omitempty"`
	// The set of capabilities to enable for a connection
	Capabilities []ConnectionProviderCapability `json:"capabilities,omitempty"`
}

// NewConnectionProvider instantiates a new ConnectionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionProvider() *ConnectionProvider {
	this := ConnectionProvider{}
	return &this
}

// NewConnectionProviderWithDefaults instantiates a new ConnectionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionProviderWithDefaults() *ConnectionProvider {
	this := ConnectionProvider{}
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ConnectionProvider) GetAuth() []ConnectionProviderAuth {
	if o == nil || o.Auth == nil {
		var ret []ConnectionProviderAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionProvider) GetAuthOk() ([]ConnectionProviderAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ConnectionProvider) HasAuth() bool {
	if o != nil && o.Auth != nil {
		return true
	}

	return false
}

// SetAuth gets a reference to the given []ConnectionProviderAuth and assigns it to the Auth field.
func (o *ConnectionProvider) SetAuth(v []ConnectionProviderAuth) {
	o.Auth = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ConnectionProvider) GetCapabilities() []ConnectionProviderCapability {
	if o == nil || o.Capabilities == nil {
		var ret []ConnectionProviderCapability
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionProvider) GetCapabilitiesOk() ([]ConnectionProviderCapability, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ConnectionProvider) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []ConnectionProviderCapability and assigns it to the Capabilities field.
func (o *ConnectionProvider) SetCapabilities(v []ConnectionProviderCapability) {
	o.Capabilities = v
}

func (o ConnectionProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionProvider struct {
	value *ConnectionProvider
	isSet bool
}

func (v NullableConnectionProvider) Get() *ConnectionProvider {
	return v.value
}

func (v *NullableConnectionProvider) Set(val *ConnectionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionProvider(val *ConnectionProvider) *NullableConnectionProvider {
	return &NullableConnectionProvider{value: val, isSet: true}
}

func (v NullableConnectionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
