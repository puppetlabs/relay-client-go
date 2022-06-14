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

// StartConnectionProviderOAuth2AuthRequest Fields to start an OAuth 2.0 authentication flow
type StartConnectionProviderOAuth2AuthRequest struct {
	// The set of capabilities to enable for a connection
	Capabilities []ConnectionProviderCapability `json:"capabilities"`
	// The future name of the connection
	Name string `json:"name"`
	// Additional scopes to explicitly request
	Scopes []string `json:"scopes,omitempty"`
	// State to return to the requestor if the authentication succeeds
	State interface{} `json:"state,omitempty"`
}

// NewStartConnectionProviderOAuth2AuthRequest instantiates a new StartConnectionProviderOAuth2AuthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartConnectionProviderOAuth2AuthRequest(capabilities []ConnectionProviderCapability, name string) *StartConnectionProviderOAuth2AuthRequest {
	this := StartConnectionProviderOAuth2AuthRequest{}
	this.Capabilities = capabilities
	this.Name = name
	return &this
}

// NewStartConnectionProviderOAuth2AuthRequestWithDefaults instantiates a new StartConnectionProviderOAuth2AuthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartConnectionProviderOAuth2AuthRequestWithDefaults() *StartConnectionProviderOAuth2AuthRequest {
	this := StartConnectionProviderOAuth2AuthRequest{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *StartConnectionProviderOAuth2AuthRequest) GetCapabilities() []ConnectionProviderCapability {
	if o == nil {
		var ret []ConnectionProviderCapability
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *StartConnectionProviderOAuth2AuthRequest) GetCapabilitiesOk() ([]ConnectionProviderCapability, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *StartConnectionProviderOAuth2AuthRequest) SetCapabilities(v []ConnectionProviderCapability) {
	o.Capabilities = v
}

// GetName returns the Name field value
func (o *StartConnectionProviderOAuth2AuthRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StartConnectionProviderOAuth2AuthRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StartConnectionProviderOAuth2AuthRequest) SetName(v string) {
	o.Name = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *StartConnectionProviderOAuth2AuthRequest) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartConnectionProviderOAuth2AuthRequest) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *StartConnectionProviderOAuth2AuthRequest) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *StartConnectionProviderOAuth2AuthRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartConnectionProviderOAuth2AuthRequest) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartConnectionProviderOAuth2AuthRequest) GetStateOk() (*interface{}, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StartConnectionProviderOAuth2AuthRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given interface{} and assigns it to the State field.
func (o *StartConnectionProviderOAuth2AuthRequest) SetState(v interface{}) {
	o.State = v
}

func (o StartConnectionProviderOAuth2AuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["capabilities"] = o.Capabilities
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableStartConnectionProviderOAuth2AuthRequest struct {
	value *StartConnectionProviderOAuth2AuthRequest
	isSet bool
}

func (v NullableStartConnectionProviderOAuth2AuthRequest) Get() *StartConnectionProviderOAuth2AuthRequest {
	return v.value
}

func (v *NullableStartConnectionProviderOAuth2AuthRequest) Set(val *StartConnectionProviderOAuth2AuthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartConnectionProviderOAuth2AuthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartConnectionProviderOAuth2AuthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartConnectionProviderOAuth2AuthRequest(val *StartConnectionProviderOAuth2AuthRequest) *NullableStartConnectionProviderOAuth2AuthRequest {
	return &NullableStartConnectionProviderOAuth2AuthRequest{value: val, isSet: true}
}

func (v NullableStartConnectionProviderOAuth2AuthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartConnectionProviderOAuth2AuthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}