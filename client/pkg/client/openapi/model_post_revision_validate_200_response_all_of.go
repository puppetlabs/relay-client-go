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

// PostRevisionValidate200ResponseAllOf The response type for retrieving a workflow revision
type PostRevisionValidate200ResponseAllOf struct {
	Revision *WorkflowRevision `json:"revision,omitempty"`
}

// NewPostRevisionValidate200ResponseAllOf instantiates a new PostRevisionValidate200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRevisionValidate200ResponseAllOf() *PostRevisionValidate200ResponseAllOf {
	this := PostRevisionValidate200ResponseAllOf{}
	return &this
}

// NewPostRevisionValidate200ResponseAllOfWithDefaults instantiates a new PostRevisionValidate200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRevisionValidate200ResponseAllOfWithDefaults() *PostRevisionValidate200ResponseAllOf {
	this := PostRevisionValidate200ResponseAllOf{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *PostRevisionValidate200ResponseAllOf) GetRevision() WorkflowRevision {
	if o == nil || o.Revision == nil {
		var ret WorkflowRevision
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRevisionValidate200ResponseAllOf) GetRevisionOk() (*WorkflowRevision, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *PostRevisionValidate200ResponseAllOf) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given WorkflowRevision and assigns it to the Revision field.
func (o *PostRevisionValidate200ResponseAllOf) SetRevision(v WorkflowRevision) {
	o.Revision = &v
}

func (o PostRevisionValidate200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullablePostRevisionValidate200ResponseAllOf struct {
	value *PostRevisionValidate200ResponseAllOf
	isSet bool
}

func (v NullablePostRevisionValidate200ResponseAllOf) Get() *PostRevisionValidate200ResponseAllOf {
	return v.value
}

func (v *NullablePostRevisionValidate200ResponseAllOf) Set(val *PostRevisionValidate200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRevisionValidate200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRevisionValidate200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRevisionValidate200ResponseAllOf(val *PostRevisionValidate200ResponseAllOf) *NullablePostRevisionValidate200ResponseAllOf {
	return &NullablePostRevisionValidate200ResponseAllOf{value: val, isSet: true}
}

func (v NullablePostRevisionValidate200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRevisionValidate200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
