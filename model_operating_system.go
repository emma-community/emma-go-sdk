/*
Public EMMA API

**Base URL:** *<u>https://api.emma.ms/external</u>*  This **Infrastructure API** is for managing the cloud infrastructure within a project.  To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.  After creating the service application, copy the **Client ID** and **Client Secret**. Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.  The Bearer access token is a text string, included in the request header, example:  *-H Authorization: Bearer {token}*  Use this token for API requests. The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
)

// checks if the OperatingSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatingSystem{}

// OperatingSystem struct for OperatingSystem
type OperatingSystem struct {
	Id           *int32  `json:"id,omitempty"`
	Family       *string `json:"family,omitempty"`
	Type         *string `json:"type,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Version      *string `json:"version,omitempty"`
}

// NewOperatingSystem instantiates a new OperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingSystem() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingSystemWithDefaults() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OperatingSystem) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperatingSystem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OperatingSystem) SetId(v int32) {
	o.Id = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *OperatingSystem) GetFamily() string {
	if o == nil || IsNil(o.Family) {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.Family) {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *OperatingSystem) HasFamily() bool {
	if o != nil && !IsNil(o.Family) {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *OperatingSystem) SetFamily(v string) {
	o.Family = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OperatingSystem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OperatingSystem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OperatingSystem) SetType(v string) {
	o.Type = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *OperatingSystem) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *OperatingSystem) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *OperatingSystem) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OperatingSystem) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OperatingSystem) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OperatingSystem) SetVersion(v string) {
	o.Version = &v
}

func (o OperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatingSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Family) {
		toSerialize["family"] = o.Family
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Architecture) {
		toSerialize["architecture"] = o.Architecture
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableOperatingSystem struct {
	value *OperatingSystem
	isSet bool
}

func (v NullableOperatingSystem) Get() *OperatingSystem {
	return v.value
}

func (v *NullableOperatingSystem) Set(val *OperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingSystem(val *OperatingSystem) *NullableOperatingSystem {
	return &NullableOperatingSystem{value: val, isSet: true}
}

func (v NullableOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
