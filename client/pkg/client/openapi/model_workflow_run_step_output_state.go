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

// WorkflowRunStepOutputState Output state for a workflow run step
type WorkflowRunStepOutputState struct {
	// A list of workflow run step outputs
	Outputs *[]WorkflowRunStepOutput `json:"outputs,omitempty"`
}

// NewWorkflowRunStepOutputState instantiates a new WorkflowRunStepOutputState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunStepOutputState() *WorkflowRunStepOutputState {
	this := WorkflowRunStepOutputState{}
	return &this
}

// NewWorkflowRunStepOutputStateWithDefaults instantiates a new WorkflowRunStepOutputState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunStepOutputStateWithDefaults() *WorkflowRunStepOutputState {
	this := WorkflowRunStepOutputState{}
	return &this
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *WorkflowRunStepOutputState) GetOutputs() []WorkflowRunStepOutput {
	if o == nil || o.Outputs == nil {
		var ret []WorkflowRunStepOutput
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunStepOutputState) GetOutputsOk() (*[]WorkflowRunStepOutput, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *WorkflowRunStepOutputState) HasOutputs() bool {
	if o != nil && o.Outputs != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []WorkflowRunStepOutput and assigns it to the Outputs field.
func (o *WorkflowRunStepOutputState) SetOutputs(v []WorkflowRunStepOutput) {
	o.Outputs = &v
}

func (o WorkflowRunStepOutputState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunStepOutputState struct {
	value *WorkflowRunStepOutputState
	isSet bool
}

func (v NullableWorkflowRunStepOutputState) Get() *WorkflowRunStepOutputState {
	return v.value
}

func (v *NullableWorkflowRunStepOutputState) Set(val *WorkflowRunStepOutputState) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunStepOutputState) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunStepOutputState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunStepOutputState(val *WorkflowRunStepOutputState) *NullableWorkflowRunStepOutputState {
	return &NullableWorkflowRunStepOutputState{value: val, isSet: true}
}

func (v NullableWorkflowRunStepOutputState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunStepOutputState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
