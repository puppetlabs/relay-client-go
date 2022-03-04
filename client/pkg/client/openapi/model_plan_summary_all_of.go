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

// PlanSummaryAllOf struct for PlanSummaryAllOf
type PlanSummaryAllOf struct {
	// Visible plan name, should be unique across available plans
	Name string `json:"name"`
}

// NewPlanSummaryAllOf instantiates a new PlanSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanSummaryAllOf(name string) *PlanSummaryAllOf {
	this := PlanSummaryAllOf{}
	this.Name = name
	return &this
}

// NewPlanSummaryAllOfWithDefaults instantiates a new PlanSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanSummaryAllOfWithDefaults() *PlanSummaryAllOf {
	this := PlanSummaryAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *PlanSummaryAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlanSummaryAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlanSummaryAllOf) SetName(v string) {
	o.Name = v
}

func (o PlanSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePlanSummaryAllOf struct {
	value *PlanSummaryAllOf
	isSet bool
}

func (v NullablePlanSummaryAllOf) Get() *PlanSummaryAllOf {
	return v.value
}

func (v *NullablePlanSummaryAllOf) Set(val *PlanSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanSummaryAllOf(val *PlanSummaryAllOf) *NullablePlanSummaryAllOf {
	return &NullablePlanSummaryAllOf{value: val, isSet: true}
}

func (v NullablePlanSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
