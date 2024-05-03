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

// checks if the SshKeyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshKeyUpdate{}

// SshKeyUpdate struct for SshKeyUpdate
type SshKeyUpdate struct {
	Name string `json:"name"`
}

type _SshKeyUpdate SshKeyUpdate

// NewSshKeyUpdate instantiates a new SshKeyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyUpdate(name string) *SshKeyUpdate {
	this := SshKeyUpdate{}
	this.Name = name
	return &this
}

// NewSshKeyUpdateWithDefaults instantiates a new SshKeyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyUpdateWithDefaults() *SshKeyUpdate {
	this := SshKeyUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *SshKeyUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SshKeyUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SshKeyUpdate) SetName(v string) {
	o.Name = v
}

func (o SshKeyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshKeyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *SshKeyUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varSshKeyUpdate := _SshKeyUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSshKeyUpdate)

	if err != nil {
		return err
	}

	*o = SshKeyUpdate(varSshKeyUpdate)

	return err
}

type NullableSshKeyUpdate struct {
	value *SshKeyUpdate
	isSet bool
}

func (v NullableSshKeyUpdate) Get() *SshKeyUpdate {
	return v.value
}

func (v *NullableSshKeyUpdate) Set(val *SshKeyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyUpdate(val *SshKeyUpdate) *NullableSshKeyUpdate {
	return &NullableSshKeyUpdate{value: val, isSet: true}
}

func (v NullableSshKeyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
