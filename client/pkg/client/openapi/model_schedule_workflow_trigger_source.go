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

// ScheduleWorkflowTriggerSource A trigger provided on a defined schedule
type ScheduleWorkflowTriggerSource struct {
	// The frequency to invoke this trigger expressed in the cron syntax
	Schedule string `json:"schedule"`
	// The type discriminator for this trigger source
	Type string `json:"type"`
}

// NewScheduleWorkflowTriggerSource instantiates a new ScheduleWorkflowTriggerSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleWorkflowTriggerSource(schedule string, type_ string) *ScheduleWorkflowTriggerSource {
	this := ScheduleWorkflowTriggerSource{}
	this.Schedule = schedule
	this.Type = type_
	return &this
}

// NewScheduleWorkflowTriggerSourceWithDefaults instantiates a new ScheduleWorkflowTriggerSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWorkflowTriggerSourceWithDefaults() *ScheduleWorkflowTriggerSource {
	this := ScheduleWorkflowTriggerSource{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *ScheduleWorkflowTriggerSource) GetSchedule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ScheduleWorkflowTriggerSource) GetScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *ScheduleWorkflowTriggerSource) SetSchedule(v string) {
	o.Schedule = v
}

// GetType returns the Type field value
func (o *ScheduleWorkflowTriggerSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleWorkflowTriggerSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleWorkflowTriggerSource) SetType(v string) {
	o.Type = v
}

func (o ScheduleWorkflowTriggerSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleWorkflowTriggerSource struct {
	value *ScheduleWorkflowTriggerSource
	isSet bool
}

func (v NullableScheduleWorkflowTriggerSource) Get() *ScheduleWorkflowTriggerSource {
	return v.value
}

func (v *NullableScheduleWorkflowTriggerSource) Set(val *ScheduleWorkflowTriggerSource) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleWorkflowTriggerSource) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleWorkflowTriggerSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleWorkflowTriggerSource(val *ScheduleWorkflowTriggerSource) *NullableScheduleWorkflowTriggerSource {
	return &NullableScheduleWorkflowTriggerSource{value: val, isSet: true}
}

func (v NullableScheduleWorkflowTriggerSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleWorkflowTriggerSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
