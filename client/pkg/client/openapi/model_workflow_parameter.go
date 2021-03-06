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

// WorkflowParameter struct for WorkflowParameter
type WorkflowParameter struct {
	Type *SchemaType `json:"type,omitempty"`
	// The default value for the parameter
	Default interface{} `json:"default,omitempty"`
	// A description of the parameter
	Description *string `json:"description,omitempty"`
}

// NewWorkflowParameter instantiates a new WorkflowParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowParameter() *WorkflowParameter {
	this := WorkflowParameter{}
	var type_ SchemaType = STRING
	this.Type = &type_
	return &this
}

// NewWorkflowParameterWithDefaults instantiates a new WorkflowParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowParameterWithDefaults() *WorkflowParameter {
	this := WorkflowParameter{}
	var type_ SchemaType = STRING
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowParameter) GetType() SchemaType {
	if o == nil || o.Type == nil {
		var ret SchemaType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameter) GetTypeOk() (*SchemaType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowParameter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SchemaType and assigns it to the Type field.
func (o *WorkflowParameter) SetType(v SchemaType) {
	o.Type = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowParameter) GetDefault() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowParameter) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *WorkflowParameter) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *WorkflowParameter) SetDefault(v interface{}) {
	o.Default = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowParameter) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameter) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowParameter) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowParameter) SetDescription(v string) {
	o.Description = &v
}

func (o WorkflowParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowParameter struct {
	value *WorkflowParameter
	isSet bool
}

func (v NullableWorkflowParameter) Get() *WorkflowParameter {
	return v.value
}

func (v *NullableWorkflowParameter) Set(val *WorkflowParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowParameter(val *WorkflowParameter) *NullableWorkflowParameter {
	return &NullableWorkflowParameter{value: val, isSet: true}
}

func (v NullableWorkflowParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
