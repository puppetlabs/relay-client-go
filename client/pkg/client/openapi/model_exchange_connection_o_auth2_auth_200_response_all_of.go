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

// ExchangeConnectionOAuth2Auth200ResponseAllOf The response type for a newly created connection
type ExchangeConnectionOAuth2Auth200ResponseAllOf struct {
	Connection *ConnectionWithDisplayOnceFields `json:"connection,omitempty"`
	// The original state information from the authentication request
	State interface{} `json:"state,omitempty"`
}

// NewExchangeConnectionOAuth2Auth200ResponseAllOf instantiates a new ExchangeConnectionOAuth2Auth200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeConnectionOAuth2Auth200ResponseAllOf() *ExchangeConnectionOAuth2Auth200ResponseAllOf {
	this := ExchangeConnectionOAuth2Auth200ResponseAllOf{}
	return &this
}

// NewExchangeConnectionOAuth2Auth200ResponseAllOfWithDefaults instantiates a new ExchangeConnectionOAuth2Auth200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeConnectionOAuth2Auth200ResponseAllOfWithDefaults() *ExchangeConnectionOAuth2Auth200ResponseAllOf {
	this := ExchangeConnectionOAuth2Auth200ResponseAllOf{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) GetConnection() ConnectionWithDisplayOnceFields {
	if o == nil || o.Connection == nil {
		var ret ConnectionWithDisplayOnceFields
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) GetConnectionOk() (*ConnectionWithDisplayOnceFields, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionWithDisplayOnceFields and assigns it to the Connection field.
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) SetConnection(v ConnectionWithDisplayOnceFields) {
	o.Connection = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) GetStateOk() (*interface{}, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given interface{} and assigns it to the State field.
func (o *ExchangeConnectionOAuth2Auth200ResponseAllOf) SetState(v interface{}) {
	o.State = v
}

func (o ExchangeConnectionOAuth2Auth200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeConnectionOAuth2Auth200ResponseAllOf struct {
	value *ExchangeConnectionOAuth2Auth200ResponseAllOf
	isSet bool
}

func (v NullableExchangeConnectionOAuth2Auth200ResponseAllOf) Get() *ExchangeConnectionOAuth2Auth200ResponseAllOf {
	return v.value
}

func (v *NullableExchangeConnectionOAuth2Auth200ResponseAllOf) Set(val *ExchangeConnectionOAuth2Auth200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeConnectionOAuth2Auth200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeConnectionOAuth2Auth200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeConnectionOAuth2Auth200ResponseAllOf(val *ExchangeConnectionOAuth2Auth200ResponseAllOf) *NullableExchangeConnectionOAuth2Auth200ResponseAllOf {
	return &NullableExchangeConnectionOAuth2Auth200ResponseAllOf{value: val, isSet: true}
}

func (v NullableExchangeConnectionOAuth2Auth200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeConnectionOAuth2Auth200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
