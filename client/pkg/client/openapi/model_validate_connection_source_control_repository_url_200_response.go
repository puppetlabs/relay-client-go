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

// ValidateConnectionSourceControlRepositoryURL200Response The response type for reauthorization an OAuth 2.0 connection
type ValidateConnectionSourceControlRepositoryURL200Response struct {
	// Whether the connection is able to access the file
	Accessible              *bool           `json:"accessible,omitempty"`
	Source                  *WorkflowSource `json:"source,omitempty"`
	WorkflowValidationError *Error          `json:"workflow_validation_error,omitempty"`
}

// NewValidateConnectionSourceControlRepositoryURL200Response instantiates a new ValidateConnectionSourceControlRepositoryURL200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateConnectionSourceControlRepositoryURL200Response() *ValidateConnectionSourceControlRepositoryURL200Response {
	this := ValidateConnectionSourceControlRepositoryURL200Response{}
	return &this
}

// NewValidateConnectionSourceControlRepositoryURL200ResponseWithDefaults instantiates a new ValidateConnectionSourceControlRepositoryURL200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateConnectionSourceControlRepositoryURL200ResponseWithDefaults() *ValidateConnectionSourceControlRepositoryURL200Response {
	this := ValidateConnectionSourceControlRepositoryURL200Response{}
	return &this
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) GetAccessible() bool {
	if o == nil || o.Accessible == nil {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) GetAccessibleOk() (*bool, bool) {
	if o == nil || o.Accessible == nil {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) HasAccessible() bool {
	if o != nil && o.Accessible != nil {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) GetSource() WorkflowSource {
	if o == nil || o.Source == nil {
		var ret WorkflowSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) GetSourceOk() (*WorkflowSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given WorkflowSource and assigns it to the Source field.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) SetSource(v WorkflowSource) {
	o.Source = &v
}

// GetWorkflowValidationError returns the WorkflowValidationError field value if set, zero value otherwise.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) GetWorkflowValidationError() Error {
	if o == nil || o.WorkflowValidationError == nil {
		var ret Error
		return ret
	}
	return *o.WorkflowValidationError
}

// GetWorkflowValidationErrorOk returns a tuple with the WorkflowValidationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) GetWorkflowValidationErrorOk() (*Error, bool) {
	if o == nil || o.WorkflowValidationError == nil {
		return nil, false
	}
	return o.WorkflowValidationError, true
}

// HasWorkflowValidationError returns a boolean if a field has been set.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) HasWorkflowValidationError() bool {
	if o != nil && o.WorkflowValidationError != nil {
		return true
	}

	return false
}

// SetWorkflowValidationError gets a reference to the given Error and assigns it to the WorkflowValidationError field.
func (o *ValidateConnectionSourceControlRepositoryURL200Response) SetWorkflowValidationError(v Error) {
	o.WorkflowValidationError = &v
}

func (o ValidateConnectionSourceControlRepositoryURL200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessible != nil {
		toSerialize["accessible"] = o.Accessible
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.WorkflowValidationError != nil {
		toSerialize["workflow_validation_error"] = o.WorkflowValidationError
	}
	return json.Marshal(toSerialize)
}

type NullableValidateConnectionSourceControlRepositoryURL200Response struct {
	value *ValidateConnectionSourceControlRepositoryURL200Response
	isSet bool
}

func (v NullableValidateConnectionSourceControlRepositoryURL200Response) Get() *ValidateConnectionSourceControlRepositoryURL200Response {
	return v.value
}

func (v *NullableValidateConnectionSourceControlRepositoryURL200Response) Set(val *ValidateConnectionSourceControlRepositoryURL200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateConnectionSourceControlRepositoryURL200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateConnectionSourceControlRepositoryURL200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateConnectionSourceControlRepositoryURL200Response(val *ValidateConnectionSourceControlRepositoryURL200Response) *NullableValidateConnectionSourceControlRepositoryURL200Response {
	return &NullableValidateConnectionSourceControlRepositoryURL200Response{value: val, isSet: true}
}

func (v NullableValidateConnectionSourceControlRepositoryURL200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateConnectionSourceControlRepositoryURL200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}