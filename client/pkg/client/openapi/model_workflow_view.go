/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WorkflowView struct for WorkflowView
type WorkflowView struct {
	Name string `json:"name"`
	// User provided friendly workflow description
	Description *string `json:"description,omitempty"`
	// Time of creation
	CreatedAt time.Time `json:"created_at"`
	// Time of last update
	UpdatedAt time.Time `json:"updated_at"`
	MostRecentRun *WorkflowRunView `json:"most_recent_run,omitempty"`
}

// NewWorkflowView instantiates a new WorkflowView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowView(name string, createdAt time.Time, updatedAt time.Time) *WorkflowView {
	this := WorkflowView{}
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewWorkflowViewWithDefaults instantiates a new WorkflowView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowViewWithDefaults() *WorkflowView {
	this := WorkflowView{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowView) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowView) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowView) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowView) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowView) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowView) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkflowView) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowView) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkflowView) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WorkflowView) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowView) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WorkflowView) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetMostRecentRun returns the MostRecentRun field value if set, zero value otherwise.
func (o *WorkflowView) GetMostRecentRun() WorkflowRunView {
	if o == nil || o.MostRecentRun == nil {
		var ret WorkflowRunView
		return ret
	}
	return *o.MostRecentRun
}

// GetMostRecentRunOk returns a tuple with the MostRecentRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowView) GetMostRecentRunOk() (*WorkflowRunView, bool) {
	if o == nil || o.MostRecentRun == nil {
		return nil, false
	}
	return o.MostRecentRun, true
}

// HasMostRecentRun returns a boolean if a field has been set.
func (o *WorkflowView) HasMostRecentRun() bool {
	if o != nil && o.MostRecentRun != nil {
		return true
	}

	return false
}

// SetMostRecentRun gets a reference to the given WorkflowRunView and assigns it to the MostRecentRun field.
func (o *WorkflowView) SetMostRecentRun(v WorkflowRunView) {
	o.MostRecentRun = &v
}

func (o WorkflowView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.MostRecentRun != nil {
		toSerialize["most_recent_run"] = o.MostRecentRun
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowView struct {
	value *WorkflowView
	isSet bool
}

func (v NullableWorkflowView) Get() *WorkflowView {
	return v.value
}

func (v *NullableWorkflowView) Set(val *WorkflowView) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowView) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowView(val *WorkflowView) *NullableWorkflowView {
	return &NullableWorkflowView{value: val, isSet: true}
}

func (v NullableWorkflowView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

