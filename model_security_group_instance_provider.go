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

// checks if the SecurityGroupInstanceProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupInstanceProvider{}

// SecurityGroupInstanceProvider struct for SecurityGroupInstanceProvider
type SecurityGroupInstanceProvider struct {
	Id   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewSecurityGroupInstanceProvider instantiates a new SecurityGroupInstanceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupInstanceProvider() *SecurityGroupInstanceProvider {
	this := SecurityGroupInstanceProvider{}
	return &this
}

// NewSecurityGroupInstanceProviderWithDefaults instantiates a new SecurityGroupInstanceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupInstanceProviderWithDefaults() *SecurityGroupInstanceProvider {
	this := SecurityGroupInstanceProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityGroupInstanceProvider) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstanceProvider) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityGroupInstanceProvider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SecurityGroupInstanceProvider) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityGroupInstanceProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstanceProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityGroupInstanceProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityGroupInstanceProvider) SetName(v string) {
	o.Name = &v
}

func (o SecurityGroupInstanceProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupInstanceProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSecurityGroupInstanceProvider struct {
	value *SecurityGroupInstanceProvider
	isSet bool
}

func (v NullableSecurityGroupInstanceProvider) Get() *SecurityGroupInstanceProvider {
	return v.value
}

func (v *NullableSecurityGroupInstanceProvider) Set(val *SecurityGroupInstanceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupInstanceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupInstanceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupInstanceProvider(val *SecurityGroupInstanceProvider) *NullableSecurityGroupInstanceProvider {
	return &NullableSecurityGroupInstanceProvider{value: val, isSet: true}
}

func (v NullableSecurityGroupInstanceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupInstanceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
