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

// OldUpdateWorkflow200Response struct for OldUpdateWorkflow200Response
type OldUpdateWorkflow200Response struct {
	Access   *EntityAccess `json:"access,omitempty"`
	Workflow *Workflow     `json:"workflow,omitempty"`
}

// NewOldUpdateWorkflow200Response instantiates a new OldUpdateWorkflow200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOldUpdateWorkflow200Response() *OldUpdateWorkflow200Response {
	this := OldUpdateWorkflow200Response{}
	return &this
}

// NewOldUpdateWorkflow200ResponseWithDefaults instantiates a new OldUpdateWorkflow200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOldUpdateWorkflow200ResponseWithDefaults() *OldUpdateWorkflow200Response {
	this := OldUpdateWorkflow200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *OldUpdateWorkflow200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldUpdateWorkflow200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *OldUpdateWorkflow200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *OldUpdateWorkflow200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *OldUpdateWorkflow200Response) GetWorkflow() Workflow {
	if o == nil || o.Workflow == nil {
		var ret Workflow
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldUpdateWorkflow200Response) GetWorkflowOk() (*Workflow, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *OldUpdateWorkflow200Response) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given Workflow and assigns it to the Workflow field.
func (o *OldUpdateWorkflow200Response) SetWorkflow(v Workflow) {
	o.Workflow = &v
}

func (o OldUpdateWorkflow200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Workflow != nil {
		toSerialize["workflow"] = o.Workflow
	}
	return json.Marshal(toSerialize)
}

type NullableOldUpdateWorkflow200Response struct {
	value *OldUpdateWorkflow200Response
	isSet bool
}

func (v NullableOldUpdateWorkflow200Response) Get() *OldUpdateWorkflow200Response {
	return v.value
}

func (v *NullableOldUpdateWorkflow200Response) Set(val *OldUpdateWorkflow200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOldUpdateWorkflow200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOldUpdateWorkflow200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOldUpdateWorkflow200Response(val *OldUpdateWorkflow200Response) *NullableOldUpdateWorkflow200Response {
	return &NullableOldUpdateWorkflow200Response{value: val, isSet: true}
}

func (v NullableOldUpdateWorkflow200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOldUpdateWorkflow200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
