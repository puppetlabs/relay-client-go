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

// WorkflowRunStep struct for WorkflowRunStep
type WorkflowRunStep struct {
	WorkflowRunStepState
}

// NewWorkflowRunStep instantiates a new WorkflowRunStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunStep(status string, type_ string) *WorkflowRunStep {
	this := WorkflowRunStep{}
	this.Status = status
	var approval string = "waiting"
	this.Approval = &approval
	this.Type = type_
	return &this
}

// NewWorkflowRunStepWithDefaults instantiates a new WorkflowRunStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunStepWithDefaults() *WorkflowRunStep {
	this := WorkflowRunStep{}
	return &this
}

func (o WorkflowRunStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowRunStepState, errWorkflowRunStepState := json.Marshal(o.WorkflowRunStepState)
	if errWorkflowRunStepState != nil {
		return []byte{}, errWorkflowRunStepState
	}
	errWorkflowRunStepState = json.Unmarshal([]byte(serializedWorkflowRunStepState), &toSerialize)
	if errWorkflowRunStepState != nil {
		return []byte{}, errWorkflowRunStepState
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunStep struct {
	value *WorkflowRunStep
	isSet bool
}

func (v NullableWorkflowRunStep) Get() *WorkflowRunStep {
	return v.value
}

func (v *NullableWorkflowRunStep) Set(val *WorkflowRunStep) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunStep) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunStep(val *WorkflowRunStep) *NullableWorkflowRunStep {
	return &NullableWorkflowRunStep{value: val, isSet: true}
}

func (v NullableWorkflowRunStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
