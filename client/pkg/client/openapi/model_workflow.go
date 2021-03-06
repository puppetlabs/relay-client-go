/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Workflow struct for Workflow
type Workflow struct {
	Folder string `json:"folder"`
	Name   string `json:"name"`
	// User provided friendly workflow description
	Description *string `json:"description,omitempty"`
	// If true, this workflow cannot be edited or run directly by the user.
	Immutable *bool `json:"immutable,omitempty"`
	// Time of creation
	CreatedAt time.Time `json:"created_at"`
	// Time of last update
	UpdatedAt      time.Time                `json:"updated_at"`
	Error          *WorkflowError           `json:"error,omitempty"`
	LatestRevision *WorkflowRevisionSummary `json:"latest_revision,omitempty"`
	Source         *WorkflowSource          `json:"source,omitempty"`
}

// NewWorkflow instantiates a new Workflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflow(folder string, name string, createdAt time.Time, updatedAt time.Time) *Workflow {
	this := Workflow{}
	this.Folder = folder
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewWorkflowWithDefaults instantiates a new Workflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWithDefaults() *Workflow {
	this := Workflow{}
	return &this
}

// GetFolder returns the Folder field value
func (o *Workflow) GetFolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Folder
}

// GetFolderOk returns a tuple with the Folder field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Folder, true
}

// SetFolder sets field value
func (o *Workflow) SetFolder(v string) {
	o.Folder = v
}

// GetName returns the Name field value
func (o *Workflow) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Workflow) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Workflow) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Workflow) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Workflow) SetDescription(v string) {
	o.Description = &v
}

// GetImmutable returns the Immutable field value if set, zero value otherwise.
func (o *Workflow) GetImmutable() bool {
	if o == nil || o.Immutable == nil {
		var ret bool
		return ret
	}
	return *o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetImmutableOk() (*bool, bool) {
	if o == nil || o.Immutable == nil {
		return nil, false
	}
	return o.Immutable, true
}

// HasImmutable returns a boolean if a field has been set.
func (o *Workflow) HasImmutable() bool {
	if o != nil && o.Immutable != nil {
		return true
	}

	return false
}

// SetImmutable gets a reference to the given bool and assigns it to the Immutable field.
func (o *Workflow) SetImmutable(v bool) {
	o.Immutable = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Workflow) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Workflow) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Workflow) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Workflow) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Workflow) GetError() WorkflowError {
	if o == nil || o.Error == nil {
		var ret WorkflowError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetErrorOk() (*WorkflowError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Workflow) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given WorkflowError and assigns it to the Error field.
func (o *Workflow) SetError(v WorkflowError) {
	o.Error = &v
}

// GetLatestRevision returns the LatestRevision field value if set, zero value otherwise.
func (o *Workflow) GetLatestRevision() WorkflowRevisionSummary {
	if o == nil || o.LatestRevision == nil {
		var ret WorkflowRevisionSummary
		return ret
	}
	return *o.LatestRevision
}

// GetLatestRevisionOk returns a tuple with the LatestRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetLatestRevisionOk() (*WorkflowRevisionSummary, bool) {
	if o == nil || o.LatestRevision == nil {
		return nil, false
	}
	return o.LatestRevision, true
}

// HasLatestRevision returns a boolean if a field has been set.
func (o *Workflow) HasLatestRevision() bool {
	if o != nil && o.LatestRevision != nil {
		return true
	}

	return false
}

// SetLatestRevision gets a reference to the given WorkflowRevisionSummary and assigns it to the LatestRevision field.
func (o *Workflow) SetLatestRevision(v WorkflowRevisionSummary) {
	o.LatestRevision = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Workflow) GetSource() WorkflowSource {
	if o == nil || o.Source == nil {
		var ret WorkflowSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetSourceOk() (*WorkflowSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Workflow) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given WorkflowSource and assigns it to the Source field.
func (o *Workflow) SetSource(v WorkflowSource) {
	o.Source = &v
}

func (o Workflow) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.LatestRevision != nil {
		toSerialize["latest_revision"] = o.LatestRevision
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflow struct {
	value *Workflow
	isSet bool
}

func (v NullableWorkflow) Get() *Workflow {
	return v.value
}

func (v *NullableWorkflow) Set(val *Workflow) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflow(val *Workflow) *NullableWorkflow {
	return &NullableWorkflow{value: val, isSet: true}
}

func (v NullableWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
