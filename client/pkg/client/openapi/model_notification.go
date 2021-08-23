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
	"time"
)

// Notification struct for Notification
type Notification struct {
	// The unique identifier of this notification
	Id string `json:"id"`
	// Time of creation
	CreatedAt time.Time `json:"created_at"`
	// Time of last update
	UpdatedAt time.Time `json:"updated_at"`
	// The attributes of this notification
	Attributes *[]string `json:"attributes,omitempty"`
	// The type of event that created the event
	Type string `json:"type"`
	// The current user has marked this notification done
	Done *bool `json:"done,omitempty"`
	// The fields to use for linking out from the notification
	Fields *map[string]interface{} `json:"fields,omitempty"`
	// Whether the current user has read this notification
	Read bool `json:"read"`
	// The state of this notification
	State string `json:"state"`
}

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification(id string, createdAt time.Time, updatedAt time.Time, type_ string, read bool, state string) *Notification {
	this := Notification{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Type = type_
	this.Read = read
	this.State = state
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetId returns the Id field value
func (o *Notification) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Notification) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Notification) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Notification) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Notification) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Notification) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Notification) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Notification) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Notification) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Notification) GetAttributes() []string {
	if o == nil || o.Attributes == nil {
		var ret []string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetAttributesOk() (*[]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Notification) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *Notification) SetAttributes(v []string) {
	o.Attributes = &v
}

// GetType returns the Type field value
func (o *Notification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Notification) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Notification) SetType(v string) {
	o.Type = v
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *Notification) GetDone() bool {
	if o == nil || o.Done == nil {
		var ret bool
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetDoneOk() (*bool, bool) {
	if o == nil || o.Done == nil {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *Notification) HasDone() bool {
	if o != nil && o.Done != nil {
		return true
	}

	return false
}

// SetDone gets a reference to the given bool and assigns it to the Done field.
func (o *Notification) SetDone(v bool) {
	o.Done = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Notification) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Notification) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *Notification) SetFields(v map[string]interface{}) {
	o.Fields = &v
}

// GetRead returns the Read field value
func (o *Notification) GetRead() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Read
}

// GetReadOk returns a tuple with the Read field value
// and a boolean to check if the value has been set.
func (o *Notification) GetReadOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Read, true
}

// SetRead sets field value
func (o *Notification) SetRead(v bool) {
	o.Read = v
}

// GetState returns the State field value
func (o *Notification) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Notification) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Notification) SetState(v string) {
	o.State = v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Done != nil {
		toSerialize["done"] = o.Done
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if true {
		toSerialize["read"] = o.Read
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

