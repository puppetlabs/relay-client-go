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

// WorkflowSummary struct for WorkflowSummary
type WorkflowSummary struct {
	Folder string `json:"folder"`
	Name   string `json:"name"`
	// User provided friendly workflow description
	Description *string `json:"description,omitempty"`
	// If true, this workflow cannot be edited or run directly by the user.
	Immutable *bool `json:"immutable,omitempty"`
}

// NewWorkflowSummary instantiates a new WorkflowSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSummary(folder string, name string) *WorkflowSummary {
	this := WorkflowSummary{}
	this.Folder = folder
	this.Name = name
	return &this
}

// NewWorkflowSummaryWithDefaults instantiates a new WorkflowSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSummaryWithDefaults() *WorkflowSummary {
	this := WorkflowSummary{}
	return &this
}

// GetFolder returns the Folder field value
func (o *WorkflowSummary) GetFolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Folder
}

// GetFolderOk returns a tuple with the Folder field value
// and a boolean to check if the value has been set.
func (o *WorkflowSummary) GetFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Folder, true
}

// SetFolder sets field value
func (o *WorkflowSummary) SetFolder(v string) {
	o.Folder = v
}

// GetName returns the Name field value
func (o *WorkflowSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowSummary) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowSummary) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowSummary) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowSummary) SetDescription(v string) {
	o.Description = &v
}

// GetImmutable returns the Immutable field value if set, zero value otherwise.
func (o *WorkflowSummary) GetImmutable() bool {
	if o == nil || o.Immutable == nil {
		var ret bool
		return ret
	}
	return *o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSummary) GetImmutableOk() (*bool, bool) {
	if o == nil || o.Immutable == nil {
		return nil, false
	}
	return o.Immutable, true
}

// HasImmutable returns a boolean if a field has been set.
func (o *WorkflowSummary) HasImmutable() bool {
	if o != nil && o.Immutable != nil {
		return true
	}

	return false
}

// SetImmutable gets a reference to the given bool and assigns it to the Immutable field.
func (o *WorkflowSummary) SetImmutable(v bool) {
	o.Immutable = &v
}

func (o WorkflowSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["folder"] = o.Folder
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Immutable != nil {
		toSerialize["immutable"] = o.Immutable
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSummary struct {
	value *WorkflowSummary
	isSet bool
}

func (v NullableWorkflowSummary) Get() *WorkflowSummary {
	return v.value
}

func (v *NullableWorkflowSummary) Set(val *WorkflowSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSummary(val *WorkflowSummary) *NullableWorkflowSummary {
	return &NullableWorkflowSummary{value: val, isSet: true}
}

func (v NullableWorkflowSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
