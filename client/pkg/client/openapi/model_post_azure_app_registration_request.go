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

// PostAzureAppRegistrationRequest The parameters for retrieving the Azure app registration payload
type PostAzureAppRegistrationRequest struct {
	// The name of the connection to creat
	ConnectionName *string `json:"connection_name,omitempty"`
	// The name of the role to assign to the service principle
	RoleName *string `json:"role_name,omitempty"`
	// The subscription id to create the connection in
	SubscriptionId *string `json:"subscription_id,omitempty"`
}

// NewPostAzureAppRegistrationRequest instantiates a new PostAzureAppRegistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAzureAppRegistrationRequest() *PostAzureAppRegistrationRequest {
	this := PostAzureAppRegistrationRequest{}
	return &this
}

// NewPostAzureAppRegistrationRequestWithDefaults instantiates a new PostAzureAppRegistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAzureAppRegistrationRequestWithDefaults() *PostAzureAppRegistrationRequest {
	this := PostAzureAppRegistrationRequest{}
	return &this
}

// GetConnectionName returns the ConnectionName field value if set, zero value otherwise.
func (o *PostAzureAppRegistrationRequest) GetConnectionName() string {
	if o == nil || o.ConnectionName == nil {
		var ret string
		return ret
	}
	return *o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAzureAppRegistrationRequest) GetConnectionNameOk() (*string, bool) {
	if o == nil || o.ConnectionName == nil {
		return nil, false
	}
	return o.ConnectionName, true
}

// HasConnectionName returns a boolean if a field has been set.
func (o *PostAzureAppRegistrationRequest) HasConnectionName() bool {
	if o != nil && o.ConnectionName != nil {
		return true
	}

	return false
}

// SetConnectionName gets a reference to the given string and assigns it to the ConnectionName field.
func (o *PostAzureAppRegistrationRequest) SetConnectionName(v string) {
	o.ConnectionName = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *PostAzureAppRegistrationRequest) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAzureAppRegistrationRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *PostAzureAppRegistrationRequest) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *PostAzureAppRegistrationRequest) SetRoleName(v string) {
	o.RoleName = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *PostAzureAppRegistrationRequest) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAzureAppRegistrationRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *PostAzureAppRegistrationRequest) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *PostAzureAppRegistrationRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o PostAzureAppRegistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionName != nil {
		toSerialize["connection_name"] = o.ConnectionName
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	if o.SubscriptionId != nil {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullablePostAzureAppRegistrationRequest struct {
	value *PostAzureAppRegistrationRequest
	isSet bool
}

func (v NullablePostAzureAppRegistrationRequest) Get() *PostAzureAppRegistrationRequest {
	return v.value
}

func (v *NullablePostAzureAppRegistrationRequest) Set(val *PostAzureAppRegistrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAzureAppRegistrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAzureAppRegistrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAzureAppRegistrationRequest(val *PostAzureAppRegistrationRequest) *NullablePostAzureAppRegistrationRequest {
	return &NullablePostAzureAppRegistrationRequest{value: val, isSet: true}
}

func (v NullablePostAzureAppRegistrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAzureAppRegistrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
