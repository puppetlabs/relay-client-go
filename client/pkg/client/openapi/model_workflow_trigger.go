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

// WorkflowTrigger struct for WorkflowTrigger
type WorkflowTrigger struct {
	// The name of this workflow trigger
	Name       string                  `json:"name"`
	Binding    *WorkflowTriggerBinding `json:"binding,omitempty"`
	References *WorkflowDataReferences `json:"references,omitempty"`
	Source     WorkflowTriggerSource   `json:"source"`
	// An expression evaluated by the backend
	When interface{} `json:"when,omitempty"`
}

// NewWorkflowTrigger instantiates a new WorkflowTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTrigger(name string, source WorkflowTriggerSource) *WorkflowTrigger {
	this := WorkflowTrigger{}
	this.Name = name
	this.Source = source
	return &this
}

// NewWorkflowTriggerWithDefaults instantiates a new WorkflowTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTriggerWithDefaults() *WorkflowTrigger {
	this := WorkflowTrigger{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowTrigger) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowTrigger) SetName(v string) {
	o.Name = v
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *WorkflowTrigger) GetBinding() WorkflowTriggerBinding {
	if o == nil || o.Binding == nil {
		var ret WorkflowTriggerBinding
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetBindingOk() (*WorkflowTriggerBinding, bool) {
	if o == nil || o.Binding == nil {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasBinding() bool {
	if o != nil && o.Binding != nil {
		return true
	}

	return false
}

// SetBinding gets a reference to the given WorkflowTriggerBinding and assigns it to the Binding field.
func (o *WorkflowTrigger) SetBinding(v WorkflowTriggerBinding) {
	o.Binding = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *WorkflowTrigger) GetReferences() WorkflowDataReferences {
	if o == nil || o.References == nil {
		var ret WorkflowDataReferences
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetReferencesOk() (*WorkflowDataReferences, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given WorkflowDataReferences and assigns it to the References field.
func (o *WorkflowTrigger) SetReferences(v WorkflowDataReferences) {
	o.References = &v
}

// GetSource returns the Source field value
func (o *WorkflowTrigger) GetSource() WorkflowTriggerSource {
	if o == nil {
		var ret WorkflowTriggerSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetSourceOk() (*WorkflowTriggerSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *WorkflowTrigger) SetSource(v WorkflowTriggerSource) {
	o.Source = v
}

// GetWhen returns the When field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTrigger) GetWhen() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTrigger) GetWhenOk() (*interface{}, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return &o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given interface{} and assigns it to the When field.
func (o *WorkflowTrigger) SetWhen(v interface{}) {
	o.When = v
}

func (o WorkflowTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Binding != nil {
		toSerialize["binding"] = o.Binding
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTrigger struct {
	value *WorkflowTrigger
	isSet bool
}

func (v NullableWorkflowTrigger) Get() *WorkflowTrigger {
	return v.value
}

func (v *NullableWorkflowTrigger) Set(val *WorkflowTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTrigger(val *WorkflowTrigger) *NullableWorkflowTrigger {
	return &NullableWorkflowTrigger{value: val, isSet: true}
}

func (v NullableWorkflowTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
