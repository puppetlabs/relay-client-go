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

// ConnectionAuthStatusAllOf Information about the authentication to this connection
type ConnectionAuthStatusAllOf struct {
	// An opaque identifier that identifies the owner of the connection
	Subject *string `json:"subject,omitempty"`
}

// NewConnectionAuthStatusAllOf instantiates a new ConnectionAuthStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionAuthStatusAllOf() *ConnectionAuthStatusAllOf {
	this := ConnectionAuthStatusAllOf{}
	return &this
}

// NewConnectionAuthStatusAllOfWithDefaults instantiates a new ConnectionAuthStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionAuthStatusAllOfWithDefaults() *ConnectionAuthStatusAllOf {
	this := ConnectionAuthStatusAllOf{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ConnectionAuthStatusAllOf) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionAuthStatusAllOf) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ConnectionAuthStatusAllOf) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ConnectionAuthStatusAllOf) SetSubject(v string) {
	o.Subject = &v
}

func (o ConnectionAuthStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionAuthStatusAllOf struct {
	value *ConnectionAuthStatusAllOf
	isSet bool
}

func (v NullableConnectionAuthStatusAllOf) Get() *ConnectionAuthStatusAllOf {
	return v.value
}

func (v *NullableConnectionAuthStatusAllOf) Set(val *ConnectionAuthStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionAuthStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionAuthStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionAuthStatusAllOf(val *ConnectionAuthStatusAllOf) *NullableConnectionAuthStatusAllOf {
	return &NullableConnectionAuthStatusAllOf{value: val, isSet: true}
}

func (v NullableConnectionAuthStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionAuthStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

