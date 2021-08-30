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

// WorkflowDataReferences An index of data references in a section of yaml
type WorkflowDataReferences struct {
	Connections *[]ConnectionReference `json:"connections,omitempty"`
	Outputs *[]WorkflowOutputReference `json:"outputs,omitempty"`
	Parameters *[]WorkflowParameterReference `json:"parameters,omitempty"`
	Secrets *[]WorkflowSecretSummary `json:"secrets,omitempty"`
}

// NewWorkflowDataReferences instantiates a new WorkflowDataReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDataReferences() *WorkflowDataReferences {
	this := WorkflowDataReferences{}
	return &this
}

// NewWorkflowDataReferencesWithDefaults instantiates a new WorkflowDataReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDataReferencesWithDefaults() *WorkflowDataReferences {
	this := WorkflowDataReferences{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *WorkflowDataReferences) GetConnections() []ConnectionReference {
	if o == nil || o.Connections == nil {
		var ret []ConnectionReference
		return ret
	}
	return *o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataReferences) GetConnectionsOk() (*[]ConnectionReference, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *WorkflowDataReferences) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []ConnectionReference and assigns it to the Connections field.
func (o *WorkflowDataReferences) SetConnections(v []ConnectionReference) {
	o.Connections = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *WorkflowDataReferences) GetOutputs() []WorkflowOutputReference {
	if o == nil || o.Outputs == nil {
		var ret []WorkflowOutputReference
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataReferences) GetOutputsOk() (*[]WorkflowOutputReference, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *WorkflowDataReferences) HasOutputs() bool {
	if o != nil && o.Outputs != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []WorkflowOutputReference and assigns it to the Outputs field.
func (o *WorkflowDataReferences) SetOutputs(v []WorkflowOutputReference) {
	o.Outputs = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WorkflowDataReferences) GetParameters() []WorkflowParameterReference {
	if o == nil || o.Parameters == nil {
		var ret []WorkflowParameterReference
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataReferences) GetParametersOk() (*[]WorkflowParameterReference, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowDataReferences) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []WorkflowParameterReference and assigns it to the Parameters field.
func (o *WorkflowDataReferences) SetParameters(v []WorkflowParameterReference) {
	o.Parameters = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *WorkflowDataReferences) GetSecrets() []WorkflowSecretSummary {
	if o == nil || o.Secrets == nil {
		var ret []WorkflowSecretSummary
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataReferences) GetSecretsOk() (*[]WorkflowSecretSummary, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *WorkflowDataReferences) HasSecrets() bool {
	if o != nil && o.Secrets != nil {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []WorkflowSecretSummary and assigns it to the Secrets field.
func (o *WorkflowDataReferences) SetSecrets(v []WorkflowSecretSummary) {
	o.Secrets = &v
}

func (o WorkflowDataReferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowDataReferences struct {
	value *WorkflowDataReferences
	isSet bool
}

func (v NullableWorkflowDataReferences) Get() *WorkflowDataReferences {
	return v.value
}

func (v *NullableWorkflowDataReferences) Set(val *WorkflowDataReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDataReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDataReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDataReferences(val *WorkflowDataReferences) *NullableWorkflowDataReferences {
	return &NullableWorkflowDataReferences{value: val, isSet: true}
}

func (v NullableWorkflowDataReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDataReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

