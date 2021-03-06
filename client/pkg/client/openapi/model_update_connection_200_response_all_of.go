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

// UpdateConnection200ResponseAllOf The response type for updating a connection
type UpdateConnection200ResponseAllOf struct {
	Connection *ConnectionWithDisplayOnceFields `json:"connection,omitempty"`
}

// NewUpdateConnection200ResponseAllOf instantiates a new UpdateConnection200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConnection200ResponseAllOf() *UpdateConnection200ResponseAllOf {
	this := UpdateConnection200ResponseAllOf{}
	return &this
}

// NewUpdateConnection200ResponseAllOfWithDefaults instantiates a new UpdateConnection200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConnection200ResponseAllOfWithDefaults() *UpdateConnection200ResponseAllOf {
	this := UpdateConnection200ResponseAllOf{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *UpdateConnection200ResponseAllOf) GetConnection() ConnectionWithDisplayOnceFields {
	if o == nil || o.Connection == nil {
		var ret ConnectionWithDisplayOnceFields
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnection200ResponseAllOf) GetConnectionOk() (*ConnectionWithDisplayOnceFields, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *UpdateConnection200ResponseAllOf) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionWithDisplayOnceFields and assigns it to the Connection field.
func (o *UpdateConnection200ResponseAllOf) SetConnection(v ConnectionWithDisplayOnceFields) {
	o.Connection = &v
}

func (o UpdateConnection200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateConnection200ResponseAllOf struct {
	value *UpdateConnection200ResponseAllOf
	isSet bool
}

func (v NullableUpdateConnection200ResponseAllOf) Get() *UpdateConnection200ResponseAllOf {
	return v.value
}

func (v *NullableUpdateConnection200ResponseAllOf) Set(val *UpdateConnection200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConnection200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConnection200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConnection200ResponseAllOf(val *UpdateConnection200ResponseAllOf) *NullableUpdateConnection200ResponseAllOf {
	return &NullableUpdateConnection200ResponseAllOf{value: val, isSet: true}
}

func (v NullableUpdateConnection200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConnection200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
