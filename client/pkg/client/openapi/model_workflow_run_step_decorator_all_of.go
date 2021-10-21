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

// WorkflowRunStepDecoratorAllOf A decorator for a workflow step
type WorkflowRunStepDecoratorAllOf struct {
	// Name of the decorator
	Name string `json:"name"`
}

// NewWorkflowRunStepDecoratorAllOf instantiates a new WorkflowRunStepDecoratorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunStepDecoratorAllOf(name string) *WorkflowRunStepDecoratorAllOf {
	this := WorkflowRunStepDecoratorAllOf{}
	this.Name = name
	return &this
}

// NewWorkflowRunStepDecoratorAllOfWithDefaults instantiates a new WorkflowRunStepDecoratorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunStepDecoratorAllOfWithDefaults() *WorkflowRunStepDecoratorAllOf {
	this := WorkflowRunStepDecoratorAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowRunStepDecoratorAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunStepDecoratorAllOf) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowRunStepDecoratorAllOf) SetName(v string) {
	o.Name = v
}

func (o WorkflowRunStepDecoratorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunStepDecoratorAllOf struct {
	value *WorkflowRunStepDecoratorAllOf
	isSet bool
}

func (v NullableWorkflowRunStepDecoratorAllOf) Get() *WorkflowRunStepDecoratorAllOf {
	return v.value
}

func (v *NullableWorkflowRunStepDecoratorAllOf) Set(val *WorkflowRunStepDecoratorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunStepDecoratorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunStepDecoratorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunStepDecoratorAllOf(val *WorkflowRunStepDecoratorAllOf) *NullableWorkflowRunStepDecoratorAllOf {
	return &NullableWorkflowRunStepDecoratorAllOf{value: val, isSet: true}
}

func (v NullableWorkflowRunStepDecoratorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunStepDecoratorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

