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

// WorkflowRunStepOutput An output for a workflow run step
type WorkflowRunStepOutput struct {
	// Step output name
	Name *string `json:"name,omitempty"`
	// Step output value
	Value interface{} `json:"value,omitempty"`
}

// NewWorkflowRunStepOutput instantiates a new WorkflowRunStepOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunStepOutput() *WorkflowRunStepOutput {
	this := WorkflowRunStepOutput{}
	return &this
}

// NewWorkflowRunStepOutputWithDefaults instantiates a new WorkflowRunStepOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunStepOutputWithDefaults() *WorkflowRunStepOutput {
	this := WorkflowRunStepOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRunStepOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunStepOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRunStepOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRunStepOutput) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRunStepOutput) GetValue() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRunStepOutput) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowRunStepOutput) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *WorkflowRunStepOutput) SetValue(v interface{}) {
	o.Value = v
}

func (o WorkflowRunStepOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunStepOutput struct {
	value *WorkflowRunStepOutput
	isSet bool
}

func (v NullableWorkflowRunStepOutput) Get() *WorkflowRunStepOutput {
	return v.value
}

func (v *NullableWorkflowRunStepOutput) Set(val *WorkflowRunStepOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunStepOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunStepOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunStepOutput(val *WorkflowRunStepOutput) *NullableWorkflowRunStepOutput {
	return &NullableWorkflowRunStepOutput{value: val, isSet: true}
}

func (v NullableWorkflowRunStepOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunStepOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


