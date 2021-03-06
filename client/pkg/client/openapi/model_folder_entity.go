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

// FolderEntity struct for FolderEntity
type FolderEntity struct {
	Access *EntityAccess `json:"access,omitempty"`
}

// NewFolderEntity instantiates a new FolderEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderEntity() *FolderEntity {
	this := FolderEntity{}
	return &this
}

// NewFolderEntityWithDefaults instantiates a new FolderEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderEntityWithDefaults() *FolderEntity {
	this := FolderEntity{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FolderEntity) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderEntity) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FolderEntity) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *FolderEntity) SetAccess(v EntityAccess) {
	o.Access = &v
}

func (o FolderEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableFolderEntity struct {
	value *FolderEntity
	isSet bool
}

func (v NullableFolderEntity) Get() *FolderEntity {
	return v.value
}

func (v *NullableFolderEntity) Set(val *FolderEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderEntity(val *FolderEntity) *NullableFolderEntity {
	return &NullableFolderEntity{value: val, isSet: true}
}

func (v NullableFolderEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
