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

// GetAccount200Response struct for GetAccount200Response
type GetAccount200Response struct {
	Access  *EntityAccess `json:"access,omitempty"`
	Account *Account      `json:"account,omitempty"`
}

// NewGetAccount200Response instantiates a new GetAccount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccount200Response() *GetAccount200Response {
	this := GetAccount200Response{}
	return &this
}

// NewGetAccount200ResponseWithDefaults instantiates a new GetAccount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccount200ResponseWithDefaults() *GetAccount200Response {
	this := GetAccount200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetAccount200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetAccount200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *GetAccount200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *GetAccount200Response) GetAccount() Account {
	if o == nil || o.Account == nil {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetAccountOk() (*Account, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *GetAccount200Response) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *GetAccount200Response) SetAccount(v Account) {
	o.Account = &v
}

func (o GetAccount200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccount200Response struct {
	value *GetAccount200Response
	isSet bool
}

func (v NullableGetAccount200Response) Get() *GetAccount200Response {
	return v.value
}

func (v *NullableGetAccount200Response) Set(val *GetAccount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccount200Response(val *GetAccount200Response) *NullableGetAccount200Response {
	return &NullableGetAccount200Response{value: val, isSet: true}
}

func (v NullableGetAccount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
