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

// GetConnection200Response struct for GetConnection200Response
type GetConnection200Response struct {
	Access     *EntityAccess `json:"access,omitempty"`
	Connection *Connection   `json:"connection,omitempty"`
}

// NewGetConnection200Response instantiates a new GetConnection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnection200Response() *GetConnection200Response {
	this := GetConnection200Response{}
	return &this
}

// NewGetConnection200ResponseWithDefaults instantiates a new GetConnection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnection200ResponseWithDefaults() *GetConnection200Response {
	this := GetConnection200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetConnection200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnection200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetConnection200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *GetConnection200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *GetConnection200Response) GetConnection() Connection {
	if o == nil || o.Connection == nil {
		var ret Connection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnection200Response) GetConnectionOk() (*Connection, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *GetConnection200Response) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given Connection and assigns it to the Connection field.
func (o *GetConnection200Response) SetConnection(v Connection) {
	o.Connection = &v
}

func (o GetConnection200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	return json.Marshal(toSerialize)
}

type NullableGetConnection200Response struct {
	value *GetConnection200Response
	isSet bool
}

func (v NullableGetConnection200Response) Get() *GetConnection200Response {
	return v.value
}

func (v *NullableGetConnection200Response) Set(val *GetConnection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnection200Response(val *GetConnection200Response) *NullableGetConnection200Response {
	return &NullableGetConnection200Response{value: val, isSet: true}
}

func (v NullableGetConnection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}