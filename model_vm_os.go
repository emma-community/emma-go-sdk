/*
Public EMMA API

This <b>Infrastructure</b> API is for managing the cloud infrastructure within a project.  To access the API, enter your project, navigate to <b>Settings</b> > <b>Service Apps</b>, and create a service application. Select the access level: <b>Read</b>, <b>Operate</b>, or <b>Manage</b>.  After creating the service application, copy the <b>Client ID</b> and <b>Client Secret</b>. Send an API request to the endpoint <b>/issue-token</b> as specified in the <b>Authentication</b> section of the API documentation. You will receive access and refresh tokens in the response.  The Bearer access token is a text string, included in the request header:  -H \"Authorization: Bearer {token}\"  Use this token for API requests.  The access token will expire in 5 minutes, after which it must be refreshed using the refresh token.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
)

// checks if the VmOs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmOs{}

// VmOs struct for VmOs
type VmOs struct {
	Id           interface{} `json:"id,omitempty"`
	Family       interface{} `json:"family,omitempty"`
	Architecture interface{} `json:"architecture,omitempty"`
	Type         interface{} `json:"type,omitempty"`
	Version      interface{} `json:"version,omitempty"`
}

// NewVmOs instantiates a new VmOs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmOs() *VmOs {
	this := VmOs{}
	return &this
}

// NewVmOsWithDefaults instantiates a new VmOs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmOsWithDefaults() *VmOs {
	this := VmOs{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmOs) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmOs) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VmOs) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *VmOs) SetId(v interface{}) {
	o.Id = v
}

// GetFamily returns the Family field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmOs) GetFamily() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmOs) GetFamilyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Family) {
		return nil, false
	}
	return &o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *VmOs) HasFamily() bool {
	if o != nil && !IsNil(o.Family) {
		return true
	}

	return false
}

// SetFamily gets a reference to the given interface{} and assigns it to the Family field.
func (o *VmOs) SetFamily(v interface{}) {
	o.Family = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmOs) GetArchitecture() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmOs) GetArchitectureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return &o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *VmOs) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given interface{} and assigns it to the Architecture field.
func (o *VmOs) SetArchitecture(v interface{}) {
	o.Architecture = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmOs) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmOs) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VmOs) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *VmOs) SetType(v interface{}) {
	o.Type = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmOs) GetVersion() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmOs) GetVersionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return &o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VmOs) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given interface{} and assigns it to the Version field.
func (o *VmOs) SetVersion(v interface{}) {
	o.Version = v
}

func (o VmOs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmOs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Architecture != nil {
		toSerialize["architecture"] = o.Architecture
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableVmOs struct {
	value *VmOs
	isSet bool
}

func (v NullableVmOs) Get() *VmOs {
	return v.value
}

func (v *NullableVmOs) Set(val *VmOs) {
	v.value = val
	v.isSet = true
}

func (v NullableVmOs) IsSet() bool {
	return v.isSet
}

func (v *NullableVmOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmOs(val *VmOs) *NullableVmOs {
	return &NullableVmOs{value: val, isSet: true}
}

func (v NullableVmOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
