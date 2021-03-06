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

// PermissionSummaryAllOf A summary of a system permission
type PermissionSummaryAllOf struct {
	// A human-readable explanation of the access provided by this permission
	Description *string `json:"description,omitempty"`
}

// NewPermissionSummaryAllOf instantiates a new PermissionSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionSummaryAllOf() *PermissionSummaryAllOf {
	this := PermissionSummaryAllOf{}
	return &this
}

// NewPermissionSummaryAllOfWithDefaults instantiates a new PermissionSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionSummaryAllOfWithDefaults() *PermissionSummaryAllOf {
	this := PermissionSummaryAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PermissionSummaryAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionSummaryAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PermissionSummaryAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PermissionSummaryAllOf) SetDescription(v string) {
	o.Description = &v
}

func (o PermissionSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionSummaryAllOf struct {
	value *PermissionSummaryAllOf
	isSet bool
}

func (v NullablePermissionSummaryAllOf) Get() *PermissionSummaryAllOf {
	return v.value
}

func (v *NullablePermissionSummaryAllOf) Set(val *PermissionSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionSummaryAllOf(val *PermissionSummaryAllOf) *NullablePermissionSummaryAllOf {
	return &NullablePermissionSummaryAllOf{value: val, isSet: true}
}

func (v NullablePermissionSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
