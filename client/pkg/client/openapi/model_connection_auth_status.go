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

// ConnectionAuthStatus struct for ConnectionAuthStatus
type ConnectionAuthStatus struct {
	// An opaque identifier that identifies the owner of the connection
	Subject *string `json:"subject,omitempty"`
	// The authentication mechanism for this connection
	Type string `json:"type"`
	// The scopes available to this connection
	Scopes []string `json:"scopes,omitempty"`
}

// NewConnectionAuthStatus instantiates a new ConnectionAuthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionAuthStatus(type_ string) *ConnectionAuthStatus {
	this := ConnectionAuthStatus{}
	this.Type = type_
	return &this
}

// NewConnectionAuthStatusWithDefaults instantiates a new ConnectionAuthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionAuthStatusWithDefaults() *ConnectionAuthStatus {
	this := ConnectionAuthStatus{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ConnectionAuthStatus) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionAuthStatus) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ConnectionAuthStatus) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ConnectionAuthStatus) SetSubject(v string) {
	o.Subject = &v
}

// GetType returns the Type field value
func (o *ConnectionAuthStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectionAuthStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectionAuthStatus) SetType(v string) {
	o.Type = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ConnectionAuthStatus) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionAuthStatus) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ConnectionAuthStatus) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ConnectionAuthStatus) SetScopes(v []string) {
	o.Scopes = v
}

func (o ConnectionAuthStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionAuthStatus struct {
	value *ConnectionAuthStatus
	isSet bool
}

func (v NullableConnectionAuthStatus) Get() *ConnectionAuthStatus {
	return v.value
}

func (v *NullableConnectionAuthStatus) Set(val *ConnectionAuthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionAuthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionAuthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionAuthStatus(val *ConnectionAuthStatus) *NullableConnectionAuthStatus {
	return &NullableConnectionAuthStatus{value: val, isSet: true}
}

func (v NullableConnectionAuthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionAuthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
