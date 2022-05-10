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

// CreateConnectionRequest Required fields to create a Relay external connection
type CreateConnectionRequest struct {
	Auth *ConnectionAuthInput `json:"auth,omitempty"`
	// A descriptive connection name
	Name string `json:"name"`
	// This connection's type identifier
	Type string `json:"type"`
}

// NewCreateConnectionRequest instantiates a new CreateConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnectionRequest(name string, type_ string) *CreateConnectionRequest {
	this := CreateConnectionRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewCreateConnectionRequestWithDefaults instantiates a new CreateConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnectionRequestWithDefaults() *CreateConnectionRequest {
	this := CreateConnectionRequest{}
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *CreateConnectionRequest) GetAuth() ConnectionAuthInput {
	if o == nil || o.Auth == nil {
		var ret ConnectionAuthInput
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequest) GetAuthOk() (*ConnectionAuthInput, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *CreateConnectionRequest) HasAuth() bool {
	if o != nil && o.Auth != nil {
		return true
	}

	return false
}

// SetAuth gets a reference to the given ConnectionAuthInput and assigns it to the Auth field.
func (o *CreateConnectionRequest) SetAuth(v ConnectionAuthInput) {
	o.Auth = &v
}

// GetName returns the Name field value
func (o *CreateConnectionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateConnectionRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CreateConnectionRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateConnectionRequest) SetType(v string) {
	o.Type = v
}

func (o CreateConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCreateConnectionRequest struct {
	value *CreateConnectionRequest
	isSet bool
}

func (v NullableCreateConnectionRequest) Get() *CreateConnectionRequest {
	return v.value
}

func (v *NullableCreateConnectionRequest) Set(val *CreateConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnectionRequest(val *CreateConnectionRequest) *NullableCreateConnectionRequest {
	return &NullableCreateConnectionRequest{value: val, isSet: true}
}

func (v NullableCreateConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
