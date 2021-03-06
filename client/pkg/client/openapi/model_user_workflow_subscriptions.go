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

// UserWorkflowSubscriptions User subscription options for a workflow
type UserWorkflowSubscriptions struct {
	// Custom workflow subscriptions
	Custom []string `json:"custom,omitempty"`
	// Subscribe to workflow notifications
	Subscribe *bool `json:"subscribe,omitempty"`
}

// NewUserWorkflowSubscriptions instantiates a new UserWorkflowSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWorkflowSubscriptions() *UserWorkflowSubscriptions {
	this := UserWorkflowSubscriptions{}
	return &this
}

// NewUserWorkflowSubscriptionsWithDefaults instantiates a new UserWorkflowSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWorkflowSubscriptionsWithDefaults() *UserWorkflowSubscriptions {
	this := UserWorkflowSubscriptions{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *UserWorkflowSubscriptions) GetCustom() []string {
	if o == nil || o.Custom == nil {
		var ret []string
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWorkflowSubscriptions) GetCustomOk() ([]string, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *UserWorkflowSubscriptions) HasCustom() bool {
	if o != nil && o.Custom != nil {
		return true
	}

	return false
}

// SetCustom gets a reference to the given []string and assigns it to the Custom field.
func (o *UserWorkflowSubscriptions) SetCustom(v []string) {
	o.Custom = v
}

// GetSubscribe returns the Subscribe field value if set, zero value otherwise.
func (o *UserWorkflowSubscriptions) GetSubscribe() bool {
	if o == nil || o.Subscribe == nil {
		var ret bool
		return ret
	}
	return *o.Subscribe
}

// GetSubscribeOk returns a tuple with the Subscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWorkflowSubscriptions) GetSubscribeOk() (*bool, bool) {
	if o == nil || o.Subscribe == nil {
		return nil, false
	}
	return o.Subscribe, true
}

// HasSubscribe returns a boolean if a field has been set.
func (o *UserWorkflowSubscriptions) HasSubscribe() bool {
	if o != nil && o.Subscribe != nil {
		return true
	}

	return false
}

// SetSubscribe gets a reference to the given bool and assigns it to the Subscribe field.
func (o *UserWorkflowSubscriptions) SetSubscribe(v bool) {
	o.Subscribe = &v
}

func (o UserWorkflowSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.Subscribe != nil {
		toSerialize["subscribe"] = o.Subscribe
	}
	return json.Marshal(toSerialize)
}

type NullableUserWorkflowSubscriptions struct {
	value *UserWorkflowSubscriptions
	isSet bool
}

func (v NullableUserWorkflowSubscriptions) Get() *UserWorkflowSubscriptions {
	return v.value
}

func (v *NullableUserWorkflowSubscriptions) Set(val *UserWorkflowSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWorkflowSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWorkflowSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWorkflowSubscriptions(val *UserWorkflowSubscriptions) *NullableUserWorkflowSubscriptions {
	return &NullableUserWorkflowSubscriptions{value: val, isSet: true}
}

func (v NullableUserWorkflowSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWorkflowSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
