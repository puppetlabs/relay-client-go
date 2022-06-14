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

// GetFolder200Response struct for GetFolder200Response
type GetFolder200Response struct {
	Access *EntityAccess `json:"access,omitempty"`
	Folder *Folder       `json:"folder,omitempty"`
}

// NewGetFolder200Response instantiates a new GetFolder200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFolder200Response() *GetFolder200Response {
	this := GetFolder200Response{}
	return &this
}

// NewGetFolder200ResponseWithDefaults instantiates a new GetFolder200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFolder200ResponseWithDefaults() *GetFolder200Response {
	this := GetFolder200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetFolder200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFolder200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetFolder200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *GetFolder200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *GetFolder200Response) GetFolder() Folder {
	if o == nil || o.Folder == nil {
		var ret Folder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFolder200Response) GetFolderOk() (*Folder, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *GetFolder200Response) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given Folder and assigns it to the Folder field.
func (o *GetFolder200Response) SetFolder(v Folder) {
	o.Folder = &v
}

func (o GetFolder200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	return json.Marshal(toSerialize)
}

type NullableGetFolder200Response struct {
	value *GetFolder200Response
	isSet bool
}

func (v NullableGetFolder200Response) Get() *GetFolder200Response {
	return v.value
}

func (v *NullableGetFolder200Response) Set(val *GetFolder200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFolder200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFolder200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFolder200Response(val *GetFolder200Response) *NullableGetFolder200Response {
	return &NullableGetFolder200Response{value: val, isSet: true}
}

func (v NullableGetFolder200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFolder200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}