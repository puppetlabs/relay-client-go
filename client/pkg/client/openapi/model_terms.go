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

// Terms Account terms and conditions
type Terms struct {
	// The creation time of this version
	CreatedAt string `json:"created_at"`
	// The version of the terms and conditions
	Version string `json:"version"`
}

// NewTerms instantiates a new Terms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerms(createdAt string, version string) *Terms {
	this := Terms{}
	this.CreatedAt = createdAt
	this.Version = version
	return &this
}

// NewTermsWithDefaults instantiates a new Terms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermsWithDefaults() *Terms {
	this := Terms{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Terms) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Terms) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Terms) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetVersion returns the Version field value
func (o *Terms) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Terms) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Terms) SetVersion(v string) {
	o.Version = v
}

func (o Terms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableTerms struct {
	value *Terms
	isSet bool
}

func (v NullableTerms) Get() *Terms {
	return v.value
}

func (v *NullableTerms) Set(val *Terms) {
	v.value = val
	v.isSet = true
}

func (v NullableTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerms(val *Terms) *NullableTerms {
	return &NullableTerms{value: val, isSet: true}
}

func (v NullableTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


