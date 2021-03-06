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

// GetConnectionProvider200Response struct for GetConnectionProvider200Response
type GetConnectionProvider200Response struct {
	Access     *EntityAccess       `json:"access,omitempty"`
	Connection *ConnectionProvider `json:"connection,omitempty"`
}

// NewGetConnectionProvider200Response instantiates a new GetConnectionProvider200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectionProvider200Response() *GetConnectionProvider200Response {
	this := GetConnectionProvider200Response{}
	return &this
}

// NewGetConnectionProvider200ResponseWithDefaults instantiates a new GetConnectionProvider200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectionProvider200ResponseWithDefaults() *GetConnectionProvider200Response {
	this := GetConnectionProvider200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetConnectionProvider200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionProvider200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetConnectionProvider200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *GetConnectionProvider200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *GetConnectionProvider200Response) GetConnection() ConnectionProvider {
	if o == nil || o.Connection == nil {
		var ret ConnectionProvider
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionProvider200Response) GetConnectionOk() (*ConnectionProvider, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *GetConnectionProvider200Response) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionProvider and assigns it to the Connection field.
func (o *GetConnectionProvider200Response) SetConnection(v ConnectionProvider) {
	o.Connection = &v
}

func (o GetConnectionProvider200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	return json.Marshal(toSerialize)
}

type NullableGetConnectionProvider200Response struct {
	value *GetConnectionProvider200Response
	isSet bool
}

func (v NullableGetConnectionProvider200Response) Get() *GetConnectionProvider200Response {
	return v.value
}

func (v *NullableGetConnectionProvider200Response) Set(val *GetConnectionProvider200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectionProvider200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectionProvider200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectionProvider200Response(val *GetConnectionProvider200Response) *NullableGetConnectionProvider200Response {
	return &NullableGetConnectionProvider200Response{value: val, isSet: true}
}

func (v NullableGetConnectionProvider200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectionProvider200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
