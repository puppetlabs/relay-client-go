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

// UpdateUserRequest The request type for updating an account user
type UpdateUserRequest struct {
	// User name
	Name *string `json:"name,omitempty"`
	// The roles to replace in this account
	Roles []RoleIdentifier `json:"roles,omitempty"`
}

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateUserRequest) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetRoles() []RoleIdentifier {
	if o == nil || o.Roles == nil {
		var ret []RoleIdentifier
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetRolesOk() ([]RoleIdentifier, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []RoleIdentifier and assigns it to the Roles field.
func (o *UpdateUserRequest) SetRoles(v []RoleIdentifier) {
	o.Roles = v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}