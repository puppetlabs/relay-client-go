/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ScheduleWorkflowTriggerSourceState The state of a schedule trigger
type ScheduleWorkflowTriggerSourceState struct {
	// The time at which this trigger source will emit an event next
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	// The lifecycle status of this trigger
	Status *string `json:"status,omitempty"`
	// The type discriminator for this trigger source
	Type string `json:"type"`
}

// NewScheduleWorkflowTriggerSourceState instantiates a new ScheduleWorkflowTriggerSourceState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleWorkflowTriggerSourceState(type_ string) *ScheduleWorkflowTriggerSourceState {
	this := ScheduleWorkflowTriggerSourceState{}
	this.Type = type_
	return &this
}

// NewScheduleWorkflowTriggerSourceStateWithDefaults instantiates a new ScheduleWorkflowTriggerSourceState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWorkflowTriggerSourceStateWithDefaults() *ScheduleWorkflowTriggerSourceState {
	this := ScheduleWorkflowTriggerSourceState{}
	return &this
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *ScheduleWorkflowTriggerSourceState) GetScheduledAt() time.Time {
	if o == nil || o.ScheduledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleWorkflowTriggerSourceState) GetScheduledAtOk() (*time.Time, bool) {
	if o == nil || o.ScheduledAt == nil {
		return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *ScheduleWorkflowTriggerSourceState) HasScheduledAt() bool {
	if o != nil && o.ScheduledAt != nil {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given time.Time and assigns it to the ScheduledAt field.
func (o *ScheduleWorkflowTriggerSourceState) SetScheduledAt(v time.Time) {
	o.ScheduledAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ScheduleWorkflowTriggerSourceState) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleWorkflowTriggerSourceState) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ScheduleWorkflowTriggerSourceState) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ScheduleWorkflowTriggerSourceState) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value
func (o *ScheduleWorkflowTriggerSourceState) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleWorkflowTriggerSourceState) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleWorkflowTriggerSourceState) SetType(v string) {
	o.Type = v
}

func (o ScheduleWorkflowTriggerSourceState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScheduledAt != nil {
		toSerialize["scheduled_at"] = o.ScheduledAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleWorkflowTriggerSourceState struct {
	value *ScheduleWorkflowTriggerSourceState
	isSet bool
}

func (v NullableScheduleWorkflowTriggerSourceState) Get() *ScheduleWorkflowTriggerSourceState {
	return v.value
}

func (v *NullableScheduleWorkflowTriggerSourceState) Set(val *ScheduleWorkflowTriggerSourceState) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleWorkflowTriggerSourceState) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleWorkflowTriggerSourceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleWorkflowTriggerSourceState(val *ScheduleWorkflowTriggerSourceState) *NullableScheduleWorkflowTriggerSourceState {
	return &NullableScheduleWorkflowTriggerSourceState{value: val, isSet: true}
}

func (v NullableScheduleWorkflowTriggerSourceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleWorkflowTriggerSourceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


