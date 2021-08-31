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

// RevisionRepositorySourceAllOf The repository source of a revision
type RevisionRepositorySourceAllOf struct {
	// The sha of the revision content
	Sha string `json:"sha"`
}

// NewRevisionRepositorySourceAllOf instantiates a new RevisionRepositorySourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionRepositorySourceAllOf(sha string) *RevisionRepositorySourceAllOf {
	this := RevisionRepositorySourceAllOf{}
	this.Sha = sha
	return &this
}

// NewRevisionRepositorySourceAllOfWithDefaults instantiates a new RevisionRepositorySourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionRepositorySourceAllOfWithDefaults() *RevisionRepositorySourceAllOf {
	this := RevisionRepositorySourceAllOf{}
	return &this
}

// GetSha returns the Sha field value
func (o *RevisionRepositorySourceAllOf) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySourceAllOf) GetShaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *RevisionRepositorySourceAllOf) SetSha(v string) {
	o.Sha = v
}

func (o RevisionRepositorySourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sha"] = o.Sha
	}
	return json.Marshal(toSerialize)
}

type NullableRevisionRepositorySourceAllOf struct {
	value *RevisionRepositorySourceAllOf
	isSet bool
}

func (v NullableRevisionRepositorySourceAllOf) Get() *RevisionRepositorySourceAllOf {
	return v.value
}

func (v *NullableRevisionRepositorySourceAllOf) Set(val *RevisionRepositorySourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionRepositorySourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionRepositorySourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionRepositorySourceAllOf(val *RevisionRepositorySourceAllOf) *NullableRevisionRepositorySourceAllOf {
	return &NullableRevisionRepositorySourceAllOf{value: val, isSet: true}
}

func (v NullableRevisionRepositorySourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionRepositorySourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


