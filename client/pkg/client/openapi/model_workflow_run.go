/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WorkflowRun struct for WorkflowRun
type WorkflowRun struct {
	// The sequential identifier for a single workflow run
	RunNumber int32              `json:"run_number"`
	Workflow  WorkflowIdentifier `json:"workflow"`
	// The API URL to this object
	ApiUrl *string `json:"api_url,omitempty"`
	// The web/HTML URL to this object
	AppUrl    *string                 `json:"app_url,omitempty"`
	CreatedBy *WorkflowRunCreator     `json:"created_by,omitempty"`
	Revision  WorkflowRevisionSummary `json:"revision"`
	State     WorkflowRunState        `json:"state"`
	// Time of creation
	CreatedAt time.Time `json:"created_at"`
	// Time of last update
	UpdatedAt time.Time `json:"updated_at"`
	// Connection type and names used by the run
	Connections *[]ConnectionReference           `json:"connections,omitempty"`
	Error       *Error                           `json:"error,omitempty"`
	Parameters  *map[string]WorkflowRunParameter `json:"parameters,omitempty"`
	// Secret names provided to the run, both used and unused
	Secrets *[]WorkflowSecretSummary `json:"secrets,omitempty"`
}

// NewWorkflowRun instantiates a new WorkflowRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRun(runNumber int32, workflow WorkflowIdentifier, revision WorkflowRevisionSummary, state WorkflowRunState, createdAt time.Time, updatedAt time.Time) *WorkflowRun {
	this := WorkflowRun{}
	this.RunNumber = runNumber
	this.Workflow = workflow
	this.Revision = revision
	this.State = state
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewWorkflowRunWithDefaults instantiates a new WorkflowRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunWithDefaults() *WorkflowRun {
	this := WorkflowRun{}
	return &this
}

// GetRunNumber returns the RunNumber field value
func (o *WorkflowRun) GetRunNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RunNumber
}

// GetRunNumberOk returns a tuple with the RunNumber field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetRunNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunNumber, true
}

// SetRunNumber sets field value
func (o *WorkflowRun) SetRunNumber(v int32) {
	o.RunNumber = v
}

// GetWorkflow returns the Workflow field value
func (o *WorkflowRun) GetWorkflow() WorkflowIdentifier {
	if o == nil {
		var ret WorkflowIdentifier
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetWorkflowOk() (*WorkflowIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *WorkflowRun) SetWorkflow(v WorkflowIdentifier) {
	o.Workflow = v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *WorkflowRun) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *WorkflowRun) HasApiUrl() bool {
	if o != nil && o.ApiUrl != nil {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *WorkflowRun) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetAppUrl returns the AppUrl field value if set, zero value otherwise.
func (o *WorkflowRun) GetAppUrl() string {
	if o == nil || o.AppUrl == nil {
		var ret string
		return ret
	}
	return *o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetAppUrlOk() (*string, bool) {
	if o == nil || o.AppUrl == nil {
		return nil, false
	}
	return o.AppUrl, true
}

// HasAppUrl returns a boolean if a field has been set.
func (o *WorkflowRun) HasAppUrl() bool {
	if o != nil && o.AppUrl != nil {
		return true
	}

	return false
}

// SetAppUrl gets a reference to the given string and assigns it to the AppUrl field.
func (o *WorkflowRun) SetAppUrl(v string) {
	o.AppUrl = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkflowRun) GetCreatedBy() WorkflowRunCreator {
	if o == nil || o.CreatedBy == nil {
		var ret WorkflowRunCreator
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCreatedByOk() (*WorkflowRunCreator, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkflowRun) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given WorkflowRunCreator and assigns it to the CreatedBy field.
func (o *WorkflowRun) SetCreatedBy(v WorkflowRunCreator) {
	o.CreatedBy = &v
}

// GetRevision returns the Revision field value
func (o *WorkflowRun) GetRevision() WorkflowRevisionSummary {
	if o == nil {
		var ret WorkflowRevisionSummary
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetRevisionOk() (*WorkflowRevisionSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *WorkflowRun) SetRevision(v WorkflowRevisionSummary) {
	o.Revision = v
}

// GetState returns the State field value
func (o *WorkflowRun) GetState() WorkflowRunState {
	if o == nil {
		var ret WorkflowRunState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetStateOk() (*WorkflowRunState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *WorkflowRun) SetState(v WorkflowRunState) {
	o.State = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkflowRun) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkflowRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WorkflowRun) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WorkflowRun) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *WorkflowRun) GetConnections() []ConnectionReference {
	if o == nil || o.Connections == nil {
		var ret []ConnectionReference
		return ret
	}
	return *o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetConnectionsOk() (*[]ConnectionReference, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *WorkflowRun) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []ConnectionReference and assigns it to the Connections field.
func (o *WorkflowRun) SetConnections(v []ConnectionReference) {
	o.Connections = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WorkflowRun) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WorkflowRun) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *WorkflowRun) SetError(v Error) {
	o.Error = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WorkflowRun) GetParameters() map[string]WorkflowRunParameter {
	if o == nil || o.Parameters == nil {
		var ret map[string]WorkflowRunParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetParametersOk() (*map[string]WorkflowRunParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowRun) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]WorkflowRunParameter and assigns it to the Parameters field.
func (o *WorkflowRun) SetParameters(v map[string]WorkflowRunParameter) {
	o.Parameters = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *WorkflowRun) GetSecrets() []WorkflowSecretSummary {
	if o == nil || o.Secrets == nil {
		var ret []WorkflowSecretSummary
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetSecretsOk() (*[]WorkflowSecretSummary, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *WorkflowRun) HasSecrets() bool {
	if o != nil && o.Secrets != nil {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []WorkflowSecretSummary and assigns it to the Secrets field.
func (o *WorkflowRun) SetSecrets(v []WorkflowSecretSummary) {
	o.Secrets = &v
}

func (o WorkflowRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["run_number"] = o.RunNumber
	}
	if true {
		toSerialize["workflow"] = o.Workflow
	}
	if o.ApiUrl != nil {
		toSerialize["api_url"] = o.ApiUrl
	}
	if o.AppUrl != nil {
		toSerialize["app_url"] = o.AppUrl
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if true {
		toSerialize["revision"] = o.Revision
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRun struct {
	value *WorkflowRun
	isSet bool
}

func (v NullableWorkflowRun) Get() *WorkflowRun {
	return v.value
}

func (v *NullableWorkflowRun) Set(val *WorkflowRun) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRun) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRun(val *WorkflowRun) *NullableWorkflowRun {
	return &NullableWorkflowRun{value: val, isSet: true}
}

func (v NullableWorkflowRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
