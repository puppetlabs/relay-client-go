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

// WorkflowAllOf A workflow
type WorkflowAllOf struct {
	Error          *WorkflowError           `json:"error,omitempty"`
	LatestRevision *WorkflowRevisionSummary `json:"latest_revision,omitempty"`
	Source         *WorkflowSource          `json:"source,omitempty"`
}

// NewWorkflowAllOf instantiates a new WorkflowAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAllOf() *WorkflowAllOf {
	this := WorkflowAllOf{}
	return &this
}

// NewWorkflowAllOfWithDefaults instantiates a new WorkflowAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAllOfWithDefaults() *WorkflowAllOf {
	this := WorkflowAllOf{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetError() WorkflowError {
	if o == nil || o.Error == nil {
		var ret WorkflowError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetErrorOk() (*WorkflowError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given WorkflowError and assigns it to the Error field.
func (o *WorkflowAllOf) SetError(v WorkflowError) {
	o.Error = &v
}

// GetLatestRevision returns the LatestRevision field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetLatestRevision() WorkflowRevisionSummary {
	if o == nil || o.LatestRevision == nil {
		var ret WorkflowRevisionSummary
		return ret
	}
	return *o.LatestRevision
}

// GetLatestRevisionOk returns a tuple with the LatestRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetLatestRevisionOk() (*WorkflowRevisionSummary, bool) {
	if o == nil || o.LatestRevision == nil {
		return nil, false
	}
	return o.LatestRevision, true
}

// HasLatestRevision returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasLatestRevision() bool {
	if o != nil && o.LatestRevision != nil {
		return true
	}

	return false
}

// SetLatestRevision gets a reference to the given WorkflowRevisionSummary and assigns it to the LatestRevision field.
func (o *WorkflowAllOf) SetLatestRevision(v WorkflowRevisionSummary) {
	o.LatestRevision = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetSource() WorkflowSource {
	if o == nil || o.Source == nil {
		var ret WorkflowSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetSourceOk() (*WorkflowSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given WorkflowSource and assigns it to the Source field.
func (o *WorkflowAllOf) SetSource(v WorkflowSource) {
	o.Source = &v
}

func (o WorkflowAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableWorkflowAllOf struct {
	value *WorkflowAllOf
	isSet bool
}

func (v NullableWorkflowAllOf) Get() *WorkflowAllOf {
	return v.value
}

func (v *NullableWorkflowAllOf) Set(val *WorkflowAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAllOf(val *WorkflowAllOf) *NullableWorkflowAllOf {
	return &NullableWorkflowAllOf{value: val, isSet: true}
}

func (v NullableWorkflowAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
