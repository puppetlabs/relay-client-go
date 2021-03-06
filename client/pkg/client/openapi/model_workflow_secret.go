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

// WorkflowSecret A workflow secret
type WorkflowSecret struct {
	Name  string       `json:"name"`
	Value BinaryString `json:"value"`
}

// NewWorkflowSecret instantiates a new WorkflowSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSecret(name string, value BinaryString) *WorkflowSecret {
	this := WorkflowSecret{}
	this.Name = name
	this.Value = value
	return &this
}

// NewWorkflowSecretWithDefaults instantiates a new WorkflowSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSecretWithDefaults() *WorkflowSecret {
	this := WorkflowSecret{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowSecret) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *WorkflowSecret) GetValue() BinaryString {
	if o == nil {
		var ret BinaryString
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WorkflowSecret) GetValueOk() (*BinaryString, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *WorkflowSecret) SetValue(v BinaryString) {
	o.Value = v
}

func (o WorkflowSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSecret struct {
	value *WorkflowSecret
	isSet bool
}

func (v NullableWorkflowSecret) Get() *WorkflowSecret {
	return v.value
}

func (v *NullableWorkflowSecret) Set(val *WorkflowSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSecret(val *WorkflowSecret) *NullableWorkflowSecret {
	return &NullableWorkflowSecret{value: val, isSet: true}
}

func (v NullableWorkflowSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
