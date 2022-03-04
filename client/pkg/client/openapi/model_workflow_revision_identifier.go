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

// WorkflowRevisionIdentifier A unique identifier for a workflow revision
type WorkflowRevisionIdentifier struct {
	// An opaque revision identifier
	Id string `json:"id"`
}

// NewWorkflowRevisionIdentifier instantiates a new WorkflowRevisionIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRevisionIdentifier(id string) *WorkflowRevisionIdentifier {
	this := WorkflowRevisionIdentifier{}
	this.Id = id
	return &this
}

// NewWorkflowRevisionIdentifierWithDefaults instantiates a new WorkflowRevisionIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRevisionIdentifierWithDefaults() *WorkflowRevisionIdentifier {
	this := WorkflowRevisionIdentifier{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowRevisionIdentifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowRevisionIdentifier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowRevisionIdentifier) SetId(v string) {
	o.Id = v
}

func (o WorkflowRevisionIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRevisionIdentifier struct {
	value *WorkflowRevisionIdentifier
	isSet bool
}

func (v NullableWorkflowRevisionIdentifier) Get() *WorkflowRevisionIdentifier {
	return v.value
}

func (v *NullableWorkflowRevisionIdentifier) Set(val *WorkflowRevisionIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRevisionIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRevisionIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRevisionIdentifier(val *WorkflowRevisionIdentifier) *NullableWorkflowRevisionIdentifier {
	return &NullableWorkflowRevisionIdentifier{value: val, isSet: true}
}

func (v NullableWorkflowRevisionIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRevisionIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
