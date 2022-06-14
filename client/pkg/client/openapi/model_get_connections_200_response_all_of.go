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

// GetConnections200ResponseAllOf The response type for listing external connections
type GetConnections200ResponseAllOf struct {
	// A list of external connections
	Connections []Connection `json:"connections,omitempty"`
}

// NewGetConnections200ResponseAllOf instantiates a new GetConnections200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnections200ResponseAllOf() *GetConnections200ResponseAllOf {
	this := GetConnections200ResponseAllOf{}
	return &this
}

// NewGetConnections200ResponseAllOfWithDefaults instantiates a new GetConnections200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnections200ResponseAllOfWithDefaults() *GetConnections200ResponseAllOf {
	this := GetConnections200ResponseAllOf{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *GetConnections200ResponseAllOf) GetConnections() []Connection {
	if o == nil || o.Connections == nil {
		var ret []Connection
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnections200ResponseAllOf) GetConnectionsOk() ([]Connection, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *GetConnections200ResponseAllOf) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []Connection and assigns it to the Connections field.
func (o *GetConnections200ResponseAllOf) SetConnections(v []Connection) {
	o.Connections = v
}

func (o GetConnections200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	return json.Marshal(toSerialize)
}

type NullableGetConnections200ResponseAllOf struct {
	value *GetConnections200ResponseAllOf
	isSet bool
}

func (v NullableGetConnections200ResponseAllOf) Get() *GetConnections200ResponseAllOf {
	return v.value
}

func (v *NullableGetConnections200ResponseAllOf) Set(val *GetConnections200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnections200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnections200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnections200ResponseAllOf(val *GetConnections200ResponseAllOf) *NullableGetConnections200ResponseAllOf {
	return &NullableGetConnections200ResponseAllOf{value: val, isSet: true}
}

func (v NullableGetConnections200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnections200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}