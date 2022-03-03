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

// ConnectionWorkflowSummaryAllOf A summary of a workflow with information about its use in a connection
type ConnectionWorkflowSummaryAllOf struct {
	// The set of capabilities to enable for a connection
	Capabilities []ConnectionProviderCapability `json:"capabilities"`
}

// NewConnectionWorkflowSummaryAllOf instantiates a new ConnectionWorkflowSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionWorkflowSummaryAllOf(capabilities []ConnectionProviderCapability) *ConnectionWorkflowSummaryAllOf {
	this := ConnectionWorkflowSummaryAllOf{}
	this.Capabilities = capabilities
	return &this
}

// NewConnectionWorkflowSummaryAllOfWithDefaults instantiates a new ConnectionWorkflowSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWorkflowSummaryAllOfWithDefaults() *ConnectionWorkflowSummaryAllOf {
	this := ConnectionWorkflowSummaryAllOf{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *ConnectionWorkflowSummaryAllOf) GetCapabilities() []ConnectionProviderCapability {
	if o == nil {
		var ret []ConnectionProviderCapability
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *ConnectionWorkflowSummaryAllOf) GetCapabilitiesOk() (*[]ConnectionProviderCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *ConnectionWorkflowSummaryAllOf) SetCapabilities(v []ConnectionProviderCapability) {
	o.Capabilities = v
}

func (o ConnectionWorkflowSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["capabilities"] = o.Capabilities
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionWorkflowSummaryAllOf struct {
	value *ConnectionWorkflowSummaryAllOf
	isSet bool
}

func (v NullableConnectionWorkflowSummaryAllOf) Get() *ConnectionWorkflowSummaryAllOf {
	return v.value
}

func (v *NullableConnectionWorkflowSummaryAllOf) Set(val *ConnectionWorkflowSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionWorkflowSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionWorkflowSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionWorkflowSummaryAllOf(val *ConnectionWorkflowSummaryAllOf) *NullableConnectionWorkflowSummaryAllOf {
	return &NullableConnectionWorkflowSummaryAllOf{value: val, isSet: true}
}

func (v NullableConnectionWorkflowSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionWorkflowSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
