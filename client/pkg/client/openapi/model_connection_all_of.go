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

// ConnectionAllOf An external connection
type ConnectionAllOf struct {
	Auth ConnectionAuthStatus `json:"auth"`
	Availability ConnectionAvailability `json:"availability"`
	// The workflows being used by this connection
	Workflows *[]ConnectionWorkflowSummary `json:"workflows,omitempty"`
}

// NewConnectionAllOf instantiates a new ConnectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionAllOf(auth ConnectionAuthStatus, availability ConnectionAvailability) *ConnectionAllOf {
	this := ConnectionAllOf{}
	this.Auth = auth
	this.Availability = availability
	return &this
}

// NewConnectionAllOfWithDefaults instantiates a new ConnectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionAllOfWithDefaults() *ConnectionAllOf {
	this := ConnectionAllOf{}
	return &this
}

// GetAuth returns the Auth field value
func (o *ConnectionAllOf) GetAuth() ConnectionAuthStatus {
	if o == nil {
		var ret ConnectionAuthStatus
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ConnectionAllOf) GetAuthOk() (*ConnectionAuthStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *ConnectionAllOf) SetAuth(v ConnectionAuthStatus) {
	o.Auth = v
}

// GetAvailability returns the Availability field value
func (o *ConnectionAllOf) GetAvailability() ConnectionAvailability {
	if o == nil {
		var ret ConnectionAvailability
		return ret
	}

	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value
// and a boolean to check if the value has been set.
func (o *ConnectionAllOf) GetAvailabilityOk() (*ConnectionAvailability, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Availability, true
}

// SetAvailability sets field value
func (o *ConnectionAllOf) SetAvailability(v ConnectionAvailability) {
	o.Availability = v
}

// GetWorkflows returns the Workflows field value if set, zero value otherwise.
func (o *ConnectionAllOf) GetWorkflows() []ConnectionWorkflowSummary {
	if o == nil || o.Workflows == nil {
		var ret []ConnectionWorkflowSummary
		return ret
	}
	return *o.Workflows
}

// GetWorkflowsOk returns a tuple with the Workflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionAllOf) GetWorkflowsOk() (*[]ConnectionWorkflowSummary, bool) {
	if o == nil || o.Workflows == nil {
		return nil, false
	}
	return o.Workflows, true
}

// HasWorkflows returns a boolean if a field has been set.
func (o *ConnectionAllOf) HasWorkflows() bool {
	if o != nil && o.Workflows != nil {
		return true
	}

	return false
}

// SetWorkflows gets a reference to the given []ConnectionWorkflowSummary and assigns it to the Workflows field.
func (o *ConnectionAllOf) SetWorkflows(v []ConnectionWorkflowSummary) {
	o.Workflows = &v
}

func (o ConnectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auth"] = o.Auth
	}
	if true {
		toSerialize["availability"] = o.Availability
	}
	if o.Workflows != nil {
		toSerialize["workflows"] = o.Workflows
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionAllOf struct {
	value *ConnectionAllOf
	isSet bool
}

func (v NullableConnectionAllOf) Get() *ConnectionAllOf {
	return v.value
}

func (v *NullableConnectionAllOf) Set(val *ConnectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionAllOf(val *ConnectionAllOf) *NullableConnectionAllOf {
	return &NullableConnectionAllOf{value: val, isSet: true}
}

func (v NullableConnectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


