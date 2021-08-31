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

// WorkflowViewAllOf A workflow view used to show aggregated workflow data
type WorkflowViewAllOf struct {
	MostRecentRun *WorkflowRunView `json:"most_recent_run,omitempty"`
}

// NewWorkflowViewAllOf instantiates a new WorkflowViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowViewAllOf() *WorkflowViewAllOf {
	this := WorkflowViewAllOf{}
	return &this
}

// NewWorkflowViewAllOfWithDefaults instantiates a new WorkflowViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowViewAllOfWithDefaults() *WorkflowViewAllOf {
	this := WorkflowViewAllOf{}
	return &this
}

// GetMostRecentRun returns the MostRecentRun field value if set, zero value otherwise.
func (o *WorkflowViewAllOf) GetMostRecentRun() WorkflowRunView {
	if o == nil || o.MostRecentRun == nil {
		var ret WorkflowRunView
		return ret
	}
	return *o.MostRecentRun
}

// GetMostRecentRunOk returns a tuple with the MostRecentRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowViewAllOf) GetMostRecentRunOk() (*WorkflowRunView, bool) {
	if o == nil || o.MostRecentRun == nil {
		return nil, false
	}
	return o.MostRecentRun, true
}

// HasMostRecentRun returns a boolean if a field has been set.
func (o *WorkflowViewAllOf) HasMostRecentRun() bool {
	if o != nil && o.MostRecentRun != nil {
		return true
	}

	return false
}

// SetMostRecentRun gets a reference to the given WorkflowRunView and assigns it to the MostRecentRun field.
func (o *WorkflowViewAllOf) SetMostRecentRun(v WorkflowRunView) {
	o.MostRecentRun = &v
}

func (o WorkflowViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MostRecentRun != nil {
		toSerialize["most_recent_run"] = o.MostRecentRun
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowViewAllOf struct {
	value *WorkflowViewAllOf
	isSet bool
}

func (v NullableWorkflowViewAllOf) Get() *WorkflowViewAllOf {
	return v.value
}

func (v *NullableWorkflowViewAllOf) Set(val *WorkflowViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowViewAllOf(val *WorkflowViewAllOf) *NullableWorkflowViewAllOf {
	return &NullableWorkflowViewAllOf{value: val, isSet: true}
}

func (v NullableWorkflowViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


