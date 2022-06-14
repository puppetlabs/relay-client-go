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

// PlanAllOf struct for PlanAllOf
type PlanAllOf struct {
	// Whether or not this billing plan should be available to users
	Available bool `json:"available"`
	// Available pricing models for this plan
	Prices []PlanPrice `json:"prices"`
}

// NewPlanAllOf instantiates a new PlanAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanAllOf(available bool, prices []PlanPrice) *PlanAllOf {
	this := PlanAllOf{}
	this.Available = available
	this.Prices = prices
	return &this
}

// NewPlanAllOfWithDefaults instantiates a new PlanAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanAllOfWithDefaults() *PlanAllOf {
	this := PlanAllOf{}
	return &this
}

// GetAvailable returns the Available field value
func (o *PlanAllOf) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *PlanAllOf) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *PlanAllOf) SetAvailable(v bool) {
	o.Available = v
}

// GetPrices returns the Prices field value
func (o *PlanAllOf) GetPrices() []PlanPrice {
	if o == nil {
		var ret []PlanPrice
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *PlanAllOf) GetPricesOk() ([]PlanPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *PlanAllOf) SetPrices(v []PlanPrice) {
	o.Prices = v
}

func (o PlanAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["available"] = o.Available
	}
	if true {
		toSerialize["prices"] = o.Prices
	}
	return json.Marshal(toSerialize)
}

type NullablePlanAllOf struct {
	value *PlanAllOf
	isSet bool
}

func (v NullablePlanAllOf) Get() *PlanAllOf {
	return v.value
}

func (v *NullablePlanAllOf) Set(val *PlanAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanAllOf(val *PlanAllOf) *NullablePlanAllOf {
	return &NullablePlanAllOf{value: val, isSet: true}
}

func (v NullablePlanAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}