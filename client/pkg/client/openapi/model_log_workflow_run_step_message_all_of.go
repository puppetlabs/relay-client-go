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

// LogWorkflowRunStepMessageAllOf An error or notice encountered during execution of a step
type LogWorkflowRunStepMessageAllOf struct {
	// The type of message
	Type string `json:"type"`
}

// NewLogWorkflowRunStepMessageAllOf instantiates a new LogWorkflowRunStepMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogWorkflowRunStepMessageAllOf(type_ string) *LogWorkflowRunStepMessageAllOf {
	this := LogWorkflowRunStepMessageAllOf{}
	this.Type = type_
	return &this
}

// NewLogWorkflowRunStepMessageAllOfWithDefaults instantiates a new LogWorkflowRunStepMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWorkflowRunStepMessageAllOfWithDefaults() *LogWorkflowRunStepMessageAllOf {
	this := LogWorkflowRunStepMessageAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *LogWorkflowRunStepMessageAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogWorkflowRunStepMessageAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogWorkflowRunStepMessageAllOf) SetType(v string) {
	o.Type = v
}

func (o LogWorkflowRunStepMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLogWorkflowRunStepMessageAllOf struct {
	value *LogWorkflowRunStepMessageAllOf
	isSet bool
}

func (v NullableLogWorkflowRunStepMessageAllOf) Get() *LogWorkflowRunStepMessageAllOf {
	return v.value
}

func (v *NullableLogWorkflowRunStepMessageAllOf) Set(val *LogWorkflowRunStepMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogWorkflowRunStepMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogWorkflowRunStepMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogWorkflowRunStepMessageAllOf(val *LogWorkflowRunStepMessageAllOf) *NullableLogWorkflowRunStepMessageAllOf {
	return &NullableLogWorkflowRunStepMessageAllOf{value: val, isSet: true}
}

func (v NullableLogWorkflowRunStepMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogWorkflowRunStepMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
