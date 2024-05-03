/*
Public EMMA API

**Base URL:** *<u>https://api.emma.ms/external</u>*  This **Infrastructure API** is for managing the cloud infrastructure within a project.  To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.  After creating the service application, copy the **Client ID** and **Client Secret**. Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.  The Bearer access token is a text string, included in the request header, example:  *-H Authorization: Bearer {token}*  Use this token for API requests. The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SshKeyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshKeyCreate{}

// SshKeyCreate struct for SshKeyCreate
type SshKeyCreate struct {
	Name    string `json:"name"`
	KeyType string `json:"keyType"`
}

type _SshKeyCreate SshKeyCreate

// NewSshKeyCreate instantiates a new SshKeyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyCreate(name string, keyType string) *SshKeyCreate {
	this := SshKeyCreate{}
	this.Name = name
	this.KeyType = keyType
	return &this
}

// NewSshKeyCreateWithDefaults instantiates a new SshKeyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyCreateWithDefaults() *SshKeyCreate {
	this := SshKeyCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SshKeyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SshKeyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SshKeyCreate) SetName(v string) {
	o.Name = v
}

// GetKeyType returns the KeyType field value
func (o *SshKeyCreate) GetKeyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *SshKeyCreate) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *SshKeyCreate) SetKeyType(v string) {
	o.KeyType = v
}

func (o SshKeyCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshKeyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["keyType"] = o.KeyType
	return toSerialize, nil
}

func (o *SshKeyCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"keyType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSshKeyCreate := _SshKeyCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSshKeyCreate)

	if err != nil {
		return err
	}

	*o = SshKeyCreate(varSshKeyCreate)

	return err
}

type NullableSshKeyCreate struct {
	value *SshKeyCreate
	isSet bool
}

func (v NullableSshKeyCreate) Get() *SshKeyCreate {
	return v.value
}

func (v *NullableSshKeyCreate) Set(val *SshKeyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyCreate(val *SshKeyCreate) *NullableSshKeyCreate {
	return &NullableSshKeyCreate{value: val, isSet: true}
}

func (v NullableSshKeyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
