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
)

// ApprovalWorkflowStep struct for ApprovalWorkflowStep
type ApprovalWorkflowStep struct {
	// Step names that must complete before this one starts
	DependsOn *[]string `json:"depends_on,omitempty"`
	// A user provided step name. Must be unique within the workflow definition
	Name string `json:"name"`
	References *WorkflowDataReferences `json:"references,omitempty"`
	// An expression evaluated by the backend
	When interface{} `json:"when,omitempty"`
	// Type of step
	Type string `json:"type"`
}

// NewApprovalWorkflowStep instantiates a new ApprovalWorkflowStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalWorkflowStep(name string, type_ string) *ApprovalWorkflowStep {
	this := ApprovalWorkflowStep{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewApprovalWorkflowStepWithDefaults instantiates a new ApprovalWorkflowStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalWorkflowStepWithDefaults() *ApprovalWorkflowStep {
	this := ApprovalWorkflowStep{}
	return &this
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise.
func (o *ApprovalWorkflowStep) GetDependsOn() []string {
	if o == nil || o.DependsOn == nil {
		var ret []string
		return ret
	}
	return *o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowStep) GetDependsOnOk() (*[]string, bool) {
	if o == nil || o.DependsOn == nil {
		return nil, false
	}
	return o.DependsOn, true
}

// HasDependsOn returns a boolean if a field has been set.
func (o *ApprovalWorkflowStep) HasDependsOn() bool {
	if o != nil && o.DependsOn != nil {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given []string and assigns it to the DependsOn field.
func (o *ApprovalWorkflowStep) SetDependsOn(v []string) {
	o.DependsOn = &v
}

// GetName returns the Name field value
func (o *ApprovalWorkflowStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApprovalWorkflowStep) SetName(v string) {
	o.Name = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ApprovalWorkflowStep) GetReferences() WorkflowDataReferences {
	if o == nil || o.References == nil {
		var ret WorkflowDataReferences
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowStep) GetReferencesOk() (*WorkflowDataReferences, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ApprovalWorkflowStep) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given WorkflowDataReferences and assigns it to the References field.
func (o *ApprovalWorkflowStep) SetReferences(v WorkflowDataReferences) {
	o.References = &v
}

// GetWhen returns the When field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalWorkflowStep) GetWhen() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalWorkflowStep) GetWhenOk() (*interface{}, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return &o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *ApprovalWorkflowStep) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given interface{} and assigns it to the When field.
func (o *ApprovalWorkflowStep) SetWhen(v interface{}) {
	o.When = v
}

// GetType returns the Type field value
func (o *ApprovalWorkflowStep) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowStep) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApprovalWorkflowStep) SetType(v string) {
	o.Type = v
}

func (o ApprovalWorkflowStep) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalWorkflowStep struct {
	value *ApprovalWorkflowStep
	isSet bool
}

func (v NullableApprovalWorkflowStep) Get() *ApprovalWorkflowStep {
	return v.value
}

func (v *NullableApprovalWorkflowStep) Set(val *ApprovalWorkflowStep) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalWorkflowStep) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalWorkflowStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalWorkflowStep(val *ApprovalWorkflowStep) *NullableApprovalWorkflowStep {
	return &NullableApprovalWorkflowStep{value: val, isSet: true}
}

func (v NullableApprovalWorkflowStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalWorkflowStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


