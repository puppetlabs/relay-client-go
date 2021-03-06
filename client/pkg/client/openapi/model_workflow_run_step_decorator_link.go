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

// WorkflowRunStepDecoratorLink A link-type decorator
type WorkflowRunStepDecoratorLink struct {
	Link WorkflowRunStepDecoratorLinkLink `json:"link"`
	Type string                           `json:"type"`
}

// NewWorkflowRunStepDecoratorLink instantiates a new WorkflowRunStepDecoratorLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunStepDecoratorLink(link WorkflowRunStepDecoratorLinkLink, type_ string) *WorkflowRunStepDecoratorLink {
	this := WorkflowRunStepDecoratorLink{}
	this.Link = link
	this.Type = type_
	return &this
}

// NewWorkflowRunStepDecoratorLinkWithDefaults instantiates a new WorkflowRunStepDecoratorLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunStepDecoratorLinkWithDefaults() *WorkflowRunStepDecoratorLink {
	this := WorkflowRunStepDecoratorLink{}
	return &this
}

// GetLink returns the Link field value
func (o *WorkflowRunStepDecoratorLink) GetLink() WorkflowRunStepDecoratorLinkLink {
	if o == nil {
		var ret WorkflowRunStepDecoratorLinkLink
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunStepDecoratorLink) GetLinkOk() (*WorkflowRunStepDecoratorLinkLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *WorkflowRunStepDecoratorLink) SetLink(v WorkflowRunStepDecoratorLinkLink) {
	o.Link = v
}

// GetType returns the Type field value
func (o *WorkflowRunStepDecoratorLink) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunStepDecoratorLink) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WorkflowRunStepDecoratorLink) SetType(v string) {
	o.Type = v
}

func (o WorkflowRunStepDecoratorLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["link"] = o.Link
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunStepDecoratorLink struct {
	value *WorkflowRunStepDecoratorLink
	isSet bool
}

func (v NullableWorkflowRunStepDecoratorLink) Get() *WorkflowRunStepDecoratorLink {
	return v.value
}

func (v *NullableWorkflowRunStepDecoratorLink) Set(val *WorkflowRunStepDecoratorLink) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunStepDecoratorLink) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunStepDecoratorLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunStepDecoratorLink(val *WorkflowRunStepDecoratorLink) *NullableWorkflowRunStepDecoratorLink {
	return &NullableWorkflowRunStepDecoratorLink{value: val, isSet: true}
}

func (v NullableWorkflowRunStepDecoratorLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunStepDecoratorLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
