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

// WorkflowRunOperation struct for WorkflowRunOperation
type WorkflowRunOperation struct {
	// Cancel a workflow run
	Cancel *bool `json:"cancel,omitempty"`
}

// NewWorkflowRunOperation instantiates a new WorkflowRunOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunOperation() *WorkflowRunOperation {
	this := WorkflowRunOperation{}
	return &this
}

// NewWorkflowRunOperationWithDefaults instantiates a new WorkflowRunOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunOperationWithDefaults() *WorkflowRunOperation {
	this := WorkflowRunOperation{}
	return &this
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *WorkflowRunOperation) GetCancel() bool {
	if o == nil || o.Cancel == nil {
		var ret bool
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunOperation) GetCancelOk() (*bool, bool) {
	if o == nil || o.Cancel == nil {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *WorkflowRunOperation) HasCancel() bool {
	if o != nil && o.Cancel != nil {
		return true
	}

	return false
}

// SetCancel gets a reference to the given bool and assigns it to the Cancel field.
func (o *WorkflowRunOperation) SetCancel(v bool) {
	o.Cancel = &v
}

func (o WorkflowRunOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cancel != nil {
		toSerialize["cancel"] = o.Cancel
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunOperation struct {
	value *WorkflowRunOperation
	isSet bool
}

func (v NullableWorkflowRunOperation) Get() *WorkflowRunOperation {
	return v.value
}

func (v *NullableWorkflowRunOperation) Set(val *WorkflowRunOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunOperation(val *WorkflowRunOperation) *NullableWorkflowRunOperation {
	return &NullableWorkflowRunOperation{value: val, isSet: true}
}

func (v NullableWorkflowRunOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
