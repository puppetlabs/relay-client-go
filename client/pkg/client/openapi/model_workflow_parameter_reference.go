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

// WorkflowParameterReference A workflow parameter reference in yaml
type WorkflowParameterReference struct {
	Name string `json:"name"`
}

// NewWorkflowParameterReference instantiates a new WorkflowParameterReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowParameterReference(name string) *WorkflowParameterReference {
	this := WorkflowParameterReference{}
	this.Name = name
	return &this
}

// NewWorkflowParameterReferenceWithDefaults instantiates a new WorkflowParameterReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowParameterReferenceWithDefaults() *WorkflowParameterReference {
	this := WorkflowParameterReference{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowParameterReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowParameterReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowParameterReference) SetName(v string) {
	o.Name = v
}

func (o WorkflowParameterReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowParameterReference struct {
	value *WorkflowParameterReference
	isSet bool
}

func (v NullableWorkflowParameterReference) Get() *WorkflowParameterReference {
	return v.value
}

func (v *NullableWorkflowParameterReference) Set(val *WorkflowParameterReference) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowParameterReference) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowParameterReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowParameterReference(val *WorkflowParameterReference) *NullableWorkflowParameterReference {
	return &NullableWorkflowParameterReference{value: val, isSet: true}
}

func (v NullableWorkflowParameterReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowParameterReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
