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

// AnyWorkflowStep struct for AnyWorkflowStep
type AnyWorkflowStep struct {
	// Step names that must complete before this one starts
	DependsOn []string `json:"depends_on,omitempty"`
	// A user provided step name. Must be unique within the workflow definition
	Name       string                  `json:"name"`
	References *WorkflowDataReferences `json:"references,omitempty"`
	// An expression evaluated by the backend
	When interface{} `json:"when,omitempty"`
}

// NewAnyWorkflowStep instantiates a new AnyWorkflowStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnyWorkflowStep(name string) *AnyWorkflowStep {
	this := AnyWorkflowStep{}
	this.Name = name
	return &this
}

// NewAnyWorkflowStepWithDefaults instantiates a new AnyWorkflowStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnyWorkflowStepWithDefaults() *AnyWorkflowStep {
	this := AnyWorkflowStep{}
	return &this
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise.
func (o *AnyWorkflowStep) GetDependsOn() []string {
	if o == nil || o.DependsOn == nil {
		var ret []string
		return ret
	}
	return o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyWorkflowStep) GetDependsOnOk() ([]string, bool) {
	if o == nil || o.DependsOn == nil {
		return nil, false
	}
	return o.DependsOn, true
}

// HasDependsOn returns a boolean if a field has been set.
func (o *AnyWorkflowStep) HasDependsOn() bool {
	if o != nil && o.DependsOn != nil {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given []string and assigns it to the DependsOn field.
func (o *AnyWorkflowStep) SetDependsOn(v []string) {
	o.DependsOn = v
}

// GetName returns the Name field value
func (o *AnyWorkflowStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnyWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnyWorkflowStep) SetName(v string) {
	o.Name = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AnyWorkflowStep) GetReferences() WorkflowDataReferences {
	if o == nil || o.References == nil {
		var ret WorkflowDataReferences
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyWorkflowStep) GetReferencesOk() (*WorkflowDataReferences, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AnyWorkflowStep) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given WorkflowDataReferences and assigns it to the References field.
func (o *AnyWorkflowStep) SetReferences(v WorkflowDataReferences) {
	o.References = &v
}

// GetWhen returns the When field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnyWorkflowStep) GetWhen() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnyWorkflowStep) GetWhenOk() (*interface{}, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return &o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *AnyWorkflowStep) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given interface{} and assigns it to the When field.
func (o *AnyWorkflowStep) SetWhen(v interface{}) {
	o.When = v
}

func (o AnyWorkflowStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DependsOn != nil {
		toSerialize["depends_on"] = o.DependsOn
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	return json.Marshal(toSerialize)
}

type NullableAnyWorkflowStep struct {
	value *AnyWorkflowStep
	isSet bool
}

func (v NullableAnyWorkflowStep) Get() *AnyWorkflowStep {
	return v.value
}

func (v *NullableAnyWorkflowStep) Set(val *AnyWorkflowStep) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyWorkflowStep) IsSet() bool {
	return v.isSet
}

func (v *NullableAnyWorkflowStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnyWorkflowStep(val *AnyWorkflowStep) *NullableAnyWorkflowStep {
	return &NullableAnyWorkflowStep{value: val, isSet: true}
}

func (v NullableAnyWorkflowStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnyWorkflowStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
