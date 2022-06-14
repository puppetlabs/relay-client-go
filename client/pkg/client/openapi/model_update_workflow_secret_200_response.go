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

// UpdateWorkflowSecret200Response struct for UpdateWorkflowSecret200Response
type UpdateWorkflowSecret200Response struct {
	Access *EntityAccess          `json:"access,omitempty"`
	Secret *WorkflowSecretSummary `json:"secret,omitempty"`
}

// NewUpdateWorkflowSecret200Response instantiates a new UpdateWorkflowSecret200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWorkflowSecret200Response() *UpdateWorkflowSecret200Response {
	this := UpdateWorkflowSecret200Response{}
	return &this
}

// NewUpdateWorkflowSecret200ResponseWithDefaults instantiates a new UpdateWorkflowSecret200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWorkflowSecret200ResponseWithDefaults() *UpdateWorkflowSecret200Response {
	this := UpdateWorkflowSecret200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *UpdateWorkflowSecret200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorkflowSecret200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *UpdateWorkflowSecret200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *UpdateWorkflowSecret200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UpdateWorkflowSecret200Response) GetSecret() WorkflowSecretSummary {
	if o == nil || o.Secret == nil {
		var ret WorkflowSecretSummary
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorkflowSecret200Response) GetSecretOk() (*WorkflowSecretSummary, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UpdateWorkflowSecret200Response) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given WorkflowSecretSummary and assigns it to the Secret field.
func (o *UpdateWorkflowSecret200Response) SetSecret(v WorkflowSecretSummary) {
	o.Secret = &v
}

func (o UpdateWorkflowSecret200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWorkflowSecret200Response struct {
	value *UpdateWorkflowSecret200Response
	isSet bool
}

func (v NullableUpdateWorkflowSecret200Response) Get() *UpdateWorkflowSecret200Response {
	return v.value
}

func (v *NullableUpdateWorkflowSecret200Response) Set(val *UpdateWorkflowSecret200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWorkflowSecret200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWorkflowSecret200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWorkflowSecret200Response(val *UpdateWorkflowSecret200Response) *NullableUpdateWorkflowSecret200Response {
	return &NullableUpdateWorkflowSecret200Response{value: val, isSet: true}
}

func (v NullableUpdateWorkflowSecret200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWorkflowSecret200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}