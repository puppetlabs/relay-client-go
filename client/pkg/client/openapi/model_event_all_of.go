/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// EventAllOf An external event
type EventAllOf struct {
	// The time this event was received
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The attribute data for this event
	Data map[string]interface{} `json:"data,omitempty"`
	// An optional key for this event
	Key    *string     `json:"key,omitempty"`
	Source EventSource `json:"source"`
}

// NewEventAllOf instantiates a new EventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventAllOf(source EventSource) *EventAllOf {
	this := EventAllOf{}
	this.Source = source
	return &this
}

// NewEventAllOfWithDefaults instantiates a new EventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventAllOfWithDefaults() *EventAllOf {
	this := EventAllOf{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EventAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EventAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventAllOf) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *EventAllOf) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EventAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EventAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EventAllOf) SetKey(v string) {
	o.Key = &v
}

// GetSource returns the Source field value
func (o *EventAllOf) GetSource() EventSource {
	if o == nil {
		var ret EventSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetSourceOk() (*EventSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EventAllOf) SetSource(v EventSource) {
	o.Source = v
}

func (o EventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableEventAllOf struct {
	value *EventAllOf
	isSet bool
}

func (v NullableEventAllOf) Get() *EventAllOf {
	return v.value
}

func (v *NullableEventAllOf) Set(val *EventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventAllOf(val *EventAllOf) *NullableEventAllOf {
	return &NullableEventAllOf{value: val, isSet: true}
}

func (v NullableEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
