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

// GetProfile200Response struct for GetProfile200Response
type GetProfile200Response struct {
	Access *EntityAccess `json:"access,omitempty"`
	User   *UserProfile  `json:"user,omitempty"`
}

// NewGetProfile200Response instantiates a new GetProfile200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfile200Response() *GetProfile200Response {
	this := GetProfile200Response{}
	return &this
}

// NewGetProfile200ResponseWithDefaults instantiates a new GetProfile200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfile200ResponseWithDefaults() *GetProfile200Response {
	this := GetProfile200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetProfile200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetProfile200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *GetProfile200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GetProfile200Response) GetUser() UserProfile {
	if o == nil || o.User == nil {
		var ret UserProfile
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetUserOk() (*UserProfile, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GetProfile200Response) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserProfile and assigns it to the User field.
func (o *GetProfile200Response) SetUser(v UserProfile) {
	o.User = &v
}

func (o GetProfile200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableGetProfile200Response struct {
	value *GetProfile200Response
	isSet bool
}

func (v NullableGetProfile200Response) Get() *GetProfile200Response {
	return v.value
}

func (v *NullableGetProfile200Response) Set(val *GetProfile200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfile200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfile200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfile200Response(val *GetProfile200Response) *NullableGetProfile200Response {
	return &NullableGetProfile200Response{value: val, isSet: true}
}

func (v NullableGetProfile200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfile200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}