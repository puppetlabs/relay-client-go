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

// CreateConnection201Response struct for CreateConnection201Response
type CreateConnection201Response struct {
	Access     *EntityAccess                    `json:"access,omitempty"`
	Connection *ConnectionWithDisplayOnceFields `json:"connection,omitempty"`
}

// NewCreateConnection201Response instantiates a new CreateConnection201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnection201Response() *CreateConnection201Response {
	this := CreateConnection201Response{}
	return &this
}

// NewCreateConnection201ResponseWithDefaults instantiates a new CreateConnection201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnection201ResponseWithDefaults() *CreateConnection201Response {
	this := CreateConnection201Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *CreateConnection201Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnection201Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *CreateConnection201Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *CreateConnection201Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *CreateConnection201Response) GetConnection() ConnectionWithDisplayOnceFields {
	if o == nil || o.Connection == nil {
		var ret ConnectionWithDisplayOnceFields
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnection201Response) GetConnectionOk() (*ConnectionWithDisplayOnceFields, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *CreateConnection201Response) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionWithDisplayOnceFields and assigns it to the Connection field.
func (o *CreateConnection201Response) SetConnection(v ConnectionWithDisplayOnceFields) {
	o.Connection = &v
}

func (o CreateConnection201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	return json.Marshal(toSerialize)
}

type NullableCreateConnection201Response struct {
	value *CreateConnection201Response
	isSet bool
}

func (v NullableCreateConnection201Response) Get() *CreateConnection201Response {
	return v.value
}

func (v *NullableCreateConnection201Response) Set(val *CreateConnection201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnection201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnection201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnection201Response(val *CreateConnection201Response) *NullableCreateConnection201Response {
	return &NullableCreateConnection201Response{value: val, isSet: true}
}

func (v NullableCreateConnection201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnection201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}