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

// EventSummary struct for EventSummary
type EventSummary struct {
	// The unique identifier of this event
	Id string `json:"id"`
}

// NewEventSummary instantiates a new EventSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSummary(id string) *EventSummary {
	this := EventSummary{}
	this.Id = id
	return &this
}

// NewEventSummaryWithDefaults instantiates a new EventSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSummaryWithDefaults() *EventSummary {
	this := EventSummary{}
	return &this
}

// GetId returns the Id field value
func (o *EventSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventSummary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventSummary) SetId(v string) {
	o.Id = v
}

func (o EventSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableEventSummary struct {
	value *EventSummary
	isSet bool
}

func (v NullableEventSummary) Get() *EventSummary {
	return v.value
}

func (v *NullableEventSummary) Set(val *EventSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSummary(val *EventSummary) *NullableEventSummary {
	return &NullableEventSummary{value: val, isSet: true}
}

func (v NullableEventSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
