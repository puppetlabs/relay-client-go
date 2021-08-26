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

// Connection struct for Connection
type Connection struct {
	// The unique ID of this connection
	Id string `json:"id"`
	// A descriptive connection name
	Name string `json:"name"`
	// This connection's type identifier
	Type string `json:"type"`
	// The set of capabilities to enable for a connection
	Capabilities []ConnectionProviderCapability `json:"capabilities"`
	// Time of creation
	CreatedAt time.Time `json:"created_at"`
	// Time of last update
	UpdatedAt time.Time `json:"updated_at"`
	Auth ConnectionAuthStatus `json:"auth"`
	Availability ConnectionAvailability `json:"availability"`
	// The workflows being used by this connection
	Workflows *[]ConnectionWorkflowSummary `json:"workflows,omitempty"`
}

// NewConnection instantiates a new Connection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnection(id string, name string, type_ string, capabilities []ConnectionProviderCapability, createdAt time.Time, updatedAt time.Time, auth ConnectionAuthStatus, availability ConnectionAvailability) *Connection {
	this := Connection{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Capabilities = capabilities
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Auth = auth
	this.Availability = availability
	return &this
}

// NewConnectionWithDefaults instantiates a new Connection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDefaults() *Connection {
	this := Connection{}
	return &this
}

// GetId returns the Id field value
func (o *Connection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Connection) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Connection) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Connection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Connection) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Connection) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Connection) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Connection) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Connection) SetType(v string) {
	o.Type = v
}

// GetCapabilities returns the Capabilities field value
func (o *Connection) GetCapabilities() []ConnectionProviderCapability {
	if o == nil {
		var ret []ConnectionProviderCapability
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *Connection) GetCapabilitiesOk() (*[]ConnectionProviderCapability, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *Connection) SetCapabilities(v []ConnectionProviderCapability) {
	o.Capabilities = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Connection) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Connection) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Connection) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Connection) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Connection) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Connection) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAuth returns the Auth field value
func (o *Connection) GetAuth() ConnectionAuthStatus {
	if o == nil {
		var ret ConnectionAuthStatus
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *Connection) GetAuthOk() (*ConnectionAuthStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *Connection) SetAuth(v ConnectionAuthStatus) {
	o.Auth = v
}

// GetAvailability returns the Availability field value
func (o *Connection) GetAvailability() ConnectionAvailability {
	if o == nil {
		var ret ConnectionAvailability
		return ret
	}

	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value
// and a boolean to check if the value has been set.
func (o *Connection) GetAvailabilityOk() (*ConnectionAvailability, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Availability, true
}

// SetAvailability sets field value
func (o *Connection) SetAvailability(v ConnectionAvailability) {
	o.Availability = v
}

// GetWorkflows returns the Workflows field value if set, zero value otherwise.
func (o *Connection) GetWorkflows() []ConnectionWorkflowSummary {
	if o == nil || o.Workflows == nil {
		var ret []ConnectionWorkflowSummary
		return ret
	}
	return *o.Workflows
}

// GetWorkflowsOk returns a tuple with the Workflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetWorkflowsOk() (*[]ConnectionWorkflowSummary, bool) {
	if o == nil || o.Workflows == nil {
		return nil, false
	}
	return o.Workflows, true
}

// HasWorkflows returns a boolean if a field has been set.
func (o *Connection) HasWorkflows() bool {
	if o != nil && o.Workflows != nil {
		return true
	}

	return false
}

// SetWorkflows gets a reference to the given []ConnectionWorkflowSummary and assigns it to the Workflows field.
func (o *Connection) SetWorkflows(v []ConnectionWorkflowSummary) {
	o.Workflows = &v
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["capabilities"] = o.Capabilities
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["auth"] = o.Auth
	}
	if true {
		toSerialize["availability"] = o.Availability
	}
	if o.Workflows != nil {
		toSerialize["workflows"] = o.Workflows
	}
	return json.Marshal(toSerialize)
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


