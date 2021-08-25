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

// EventWorkflowRunCreator A summary of the event that created a workflow run
type EventWorkflowRunCreator struct {
	Event *Event `json:"event,omitempty"`
	// The event key used for this workflow run
	Key interface{} `json:"key,omitempty"`
	// The type of creator
	Type string `json:"type"`
}

// NewEventWorkflowRunCreator instantiates a new EventWorkflowRunCreator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventWorkflowRunCreator(type_ string) *EventWorkflowRunCreator {
	this := EventWorkflowRunCreator{}
	this.Type = type_
	return &this
}

// NewEventWorkflowRunCreatorWithDefaults instantiates a new EventWorkflowRunCreator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWorkflowRunCreatorWithDefaults() *EventWorkflowRunCreator {
	this := EventWorkflowRunCreator{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *EventWorkflowRunCreator) GetEvent() Event {
	if o == nil || o.Event == nil {
		var ret Event
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventWorkflowRunCreator) GetEventOk() (*Event, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EventWorkflowRunCreator) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given Event and assigns it to the Event field.
func (o *EventWorkflowRunCreator) SetEvent(v Event) {
	o.Event = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventWorkflowRunCreator) GetKey() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventWorkflowRunCreator) GetKeyOk() (*interface{}, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return &o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EventWorkflowRunCreator) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given interface{} and assigns it to the Key field.
func (o *EventWorkflowRunCreator) SetKey(v interface{}) {
	o.Key = v
}

// GetType returns the Type field value
func (o *EventWorkflowRunCreator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventWorkflowRunCreator) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventWorkflowRunCreator) SetType(v string) {
	o.Type = v
}

func (o EventWorkflowRunCreator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEventWorkflowRunCreator struct {
	value *EventWorkflowRunCreator
	isSet bool
}

func (v NullableEventWorkflowRunCreator) Get() *EventWorkflowRunCreator {
	return v.value
}

func (v *NullableEventWorkflowRunCreator) Set(val *EventWorkflowRunCreator) {
	v.value = val
	v.isSet = true
}

func (v NullableEventWorkflowRunCreator) IsSet() bool {
	return v.isSet
}

func (v *NullableEventWorkflowRunCreator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventWorkflowRunCreator(val *EventWorkflowRunCreator) *NullableEventWorkflowRunCreator {
	return &NullableEventWorkflowRunCreator{value: val, isSet: true}
}

func (v NullableEventWorkflowRunCreator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventWorkflowRunCreator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

