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
)

// EventResponseAllOf The response type for creating an event
type EventResponseAllOf struct {
	Event *Event `json:"event,omitempty"`
}

// NewEventResponseAllOf instantiates a new EventResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseAllOf() *EventResponseAllOf {
	this := EventResponseAllOf{}
	return &this
}

// NewEventResponseAllOfWithDefaults instantiates a new EventResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseAllOfWithDefaults() *EventResponseAllOf {
	this := EventResponseAllOf{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *EventResponseAllOf) GetEvent() Event {
	if o == nil || o.Event == nil {
		var ret Event
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAllOf) GetEventOk() (*Event, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EventResponseAllOf) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given Event and assigns it to the Event field.
func (o *EventResponseAllOf) SetEvent(v Event) {
	o.Event = &v
}

func (o EventResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponseAllOf struct {
	value *EventResponseAllOf
	isSet bool
}

func (v NullableEventResponseAllOf) Get() *EventResponseAllOf {
	return v.value
}

func (v *NullableEventResponseAllOf) Set(val *EventResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseAllOf(val *EventResponseAllOf) *NullableEventResponseAllOf {
	return &NullableEventResponseAllOf{value: val, isSet: true}
}

func (v NullableEventResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


