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

// WorkflowSecretSummary Metadata about a workflow secret
type WorkflowSecretSummary struct {
	Name string `json:"name"`
}

// NewWorkflowSecretSummary instantiates a new WorkflowSecretSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSecretSummary(name string) *WorkflowSecretSummary {
	this := WorkflowSecretSummary{}
	this.Name = name
	return &this
}

// NewWorkflowSecretSummaryWithDefaults instantiates a new WorkflowSecretSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSecretSummaryWithDefaults() *WorkflowSecretSummary {
	this := WorkflowSecretSummary{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowSecretSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowSecretSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowSecretSummary) SetName(v string) {
	o.Name = v
}

func (o WorkflowSecretSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSecretSummary struct {
	value *WorkflowSecretSummary
	isSet bool
}

func (v NullableWorkflowSecretSummary) Get() *WorkflowSecretSummary {
	return v.value
}

func (v *NullableWorkflowSecretSummary) Set(val *WorkflowSecretSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSecretSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSecretSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSecretSummary(val *WorkflowSecretSummary) *NullableWorkflowSecretSummary {
	return &NullableWorkflowSecretSummary{value: val, isSet: true}
}

func (v NullableWorkflowSecretSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSecretSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
