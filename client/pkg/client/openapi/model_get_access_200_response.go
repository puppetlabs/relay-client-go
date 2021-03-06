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

// GetAccess200Response The response type for access to global entities
type GetAccess200Response struct {
	// The global grants associated with this user
	PermissionGrants []PermissionGrant `json:"permission_grants,omitempty"`
}

// NewGetAccess200Response instantiates a new GetAccess200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccess200Response() *GetAccess200Response {
	this := GetAccess200Response{}
	return &this
}

// NewGetAccess200ResponseWithDefaults instantiates a new GetAccess200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccess200ResponseWithDefaults() *GetAccess200Response {
	this := GetAccess200Response{}
	return &this
}

// GetPermissionGrants returns the PermissionGrants field value if set, zero value otherwise.
func (o *GetAccess200Response) GetPermissionGrants() []PermissionGrant {
	if o == nil || o.PermissionGrants == nil {
		var ret []PermissionGrant
		return ret
	}
	return o.PermissionGrants
}

// GetPermissionGrantsOk returns a tuple with the PermissionGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccess200Response) GetPermissionGrantsOk() ([]PermissionGrant, bool) {
	if o == nil || o.PermissionGrants == nil {
		return nil, false
	}
	return o.PermissionGrants, true
}

// HasPermissionGrants returns a boolean if a field has been set.
func (o *GetAccess200Response) HasPermissionGrants() bool {
	if o != nil && o.PermissionGrants != nil {
		return true
	}

	return false
}

// SetPermissionGrants gets a reference to the given []PermissionGrant and assigns it to the PermissionGrants field.
func (o *GetAccess200Response) SetPermissionGrants(v []PermissionGrant) {
	o.PermissionGrants = v
}

func (o GetAccess200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PermissionGrants != nil {
		toSerialize["permission_grants"] = o.PermissionGrants
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccess200Response struct {
	value *GetAccess200Response
	isSet bool
}

func (v NullableGetAccess200Response) Get() *GetAccess200Response {
	return v.value
}

func (v *NullableGetAccess200Response) Set(val *GetAccess200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccess200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccess200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccess200Response(val *GetAccess200Response) *NullableGetAccess200Response {
	return &NullableGetAccess200Response{value: val, isSet: true}
}

func (v NullableGetAccess200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccess200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
