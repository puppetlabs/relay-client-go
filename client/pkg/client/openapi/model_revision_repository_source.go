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

// RevisionRepositorySource struct for RevisionRepositorySource
type RevisionRepositorySource struct {
	Connection ConnectionIdentifier `json:"connection"`
	// The path to the workflow content file
	Path string `json:"path"`
	// The branch or tag
	Ref string `json:"ref"`
	// The repository name
	Repository string `json:"repository"`
	// The type discriminator for this repository source
	Type *string `json:"type,omitempty"`
	// The URL used for the repository source
	Url *string `json:"url,omitempty"`
	// The sha of the revision content
	Sha string `json:"sha"`
}

// NewRevisionRepositorySource instantiates a new RevisionRepositorySource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionRepositorySource(connection ConnectionIdentifier, path string, ref string, repository string, sha string) *RevisionRepositorySource {
	this := RevisionRepositorySource{}
	this.Connection = connection
	this.Path = path
	this.Ref = ref
	this.Repository = repository
	this.Sha = sha
	return &this
}

// NewRevisionRepositorySourceWithDefaults instantiates a new RevisionRepositorySource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionRepositorySourceWithDefaults() *RevisionRepositorySource {
	this := RevisionRepositorySource{}
	return &this
}

// GetConnection returns the Connection field value
func (o *RevisionRepositorySource) GetConnection() ConnectionIdentifier {
	if o == nil {
		var ret ConnectionIdentifier
		return ret
	}

	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySource) GetConnectionOk() (*ConnectionIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Connection, true
}

// SetConnection sets field value
func (o *RevisionRepositorySource) SetConnection(v ConnectionIdentifier) {
	o.Connection = v
}

// GetPath returns the Path field value
func (o *RevisionRepositorySource) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySource) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RevisionRepositorySource) SetPath(v string) {
	o.Path = v
}

// GetRef returns the Ref field value
func (o *RevisionRepositorySource) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySource) GetRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *RevisionRepositorySource) SetRef(v string) {
	o.Ref = v
}

// GetRepository returns the Repository field value
func (o *RevisionRepositorySource) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySource) GetRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *RevisionRepositorySource) SetRepository(v string) {
	o.Repository = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RevisionRepositorySource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RevisionRepositorySource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RevisionRepositorySource) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RevisionRepositorySource) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySource) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RevisionRepositorySource) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RevisionRepositorySource) SetUrl(v string) {
	o.Url = &v
}

// GetSha returns the Sha field value
func (o *RevisionRepositorySource) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *RevisionRepositorySource) GetShaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *RevisionRepositorySource) SetSha(v string) {
	o.Sha = v
}

func (o RevisionRepositorySource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connection"] = o.Connection
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["ref"] = o.Ref
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["sha"] = o.Sha
	}
	return json.Marshal(toSerialize)
}

type NullableRevisionRepositorySource struct {
	value *RevisionRepositorySource
	isSet bool
}

func (v NullableRevisionRepositorySource) Get() *RevisionRepositorySource {
	return v.value
}

func (v *NullableRevisionRepositorySource) Set(val *RevisionRepositorySource) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionRepositorySource) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionRepositorySource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionRepositorySource(val *RevisionRepositorySource) *NullableRevisionRepositorySource {
	return &NullableRevisionRepositorySource{value: val, isSet: true}
}

func (v NullableRevisionRepositorySource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionRepositorySource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

