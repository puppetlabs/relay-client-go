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

// Folder A folder
type Folder struct {
	// The creation date of the folder
	CreatedAt *string `json:"created_at,omitempty"`
	Name      string  `json:"name"`
	// The last update date of the folder
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewFolder instantiates a new Folder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolder(name string) *Folder {
	this := Folder{}
	this.Name = name
	return &this
}

// NewFolderWithDefaults instantiates a new Folder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderWithDefaults() *Folder {
	this := Folder{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Folder) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Folder) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Folder) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetName returns the Name field value
func (o *Folder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Folder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Folder) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Folder) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Folder) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Folder) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Folder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableFolder struct {
	value *Folder
	isSet bool
}

func (v NullableFolder) Get() *Folder {
	return v.value
}

func (v *NullableFolder) Set(val *Folder) {
	v.value = val
	v.isSet = true
}

func (v NullableFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolder(val *Folder) *NullableFolder {
	return &NullableFolder{value: val, isSet: true}
}

func (v NullableFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
