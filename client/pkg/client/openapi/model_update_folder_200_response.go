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

// UpdateFolder200Response struct for UpdateFolder200Response
type UpdateFolder200Response struct {
	Access *EntityAccess `json:"access,omitempty"`
	Folder *Folder       `json:"folder,omitempty"`
}

// NewUpdateFolder200Response instantiates a new UpdateFolder200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFolder200Response() *UpdateFolder200Response {
	this := UpdateFolder200Response{}
	return &this
}

// NewUpdateFolder200ResponseWithDefaults instantiates a new UpdateFolder200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFolder200ResponseWithDefaults() *UpdateFolder200Response {
	this := UpdateFolder200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *UpdateFolder200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolder200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *UpdateFolder200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *UpdateFolder200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *UpdateFolder200Response) GetFolder() Folder {
	if o == nil || o.Folder == nil {
		var ret Folder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolder200Response) GetFolderOk() (*Folder, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *UpdateFolder200Response) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given Folder and assigns it to the Folder field.
func (o *UpdateFolder200Response) SetFolder(v Folder) {
	o.Folder = &v
}

func (o UpdateFolder200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFolder200Response struct {
	value *UpdateFolder200Response
	isSet bool
}

func (v NullableUpdateFolder200Response) Get() *UpdateFolder200Response {
	return v.value
}

func (v *NullableUpdateFolder200Response) Set(val *UpdateFolder200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFolder200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFolder200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFolder200Response(val *UpdateFolder200Response) *NullableUpdateFolder200Response {
	return &NullableUpdateFolder200Response{value: val, isSet: true}
}

func (v NullableUpdateFolder200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFolder200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
