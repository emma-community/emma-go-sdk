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

// checks if the VmShutdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmShutdown{}

// VmShutdown struct for VmShutdown
type VmShutdown struct {
	// Action with a virtual machine
	Action string `json:"action"`
}

type _VmShutdown VmShutdown

// NewVmShutdown instantiates a new VmShutdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmShutdown(action string) *VmShutdown {
	this := VmShutdown{}
	this.Action = action
	return &this
}

// NewVmShutdownWithDefaults instantiates a new VmShutdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmShutdownWithDefaults() *VmShutdown {
	this := VmShutdown{}
	return &this
}

// GetAction returns the Action field value
func (o *VmShutdown) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *VmShutdown) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *VmShutdown) SetAction(v string) {
	o.Action = v
}

func (o VmShutdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmShutdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *VmShutdown) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varVmShutdown := _VmShutdown{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmShutdown)

	if err != nil {
		return err
	}

	*o = VmShutdown(varVmShutdown)

	return err
}

type NullableVmShutdown struct {
	value *VmShutdown
	isSet bool
}

func (v NullableVmShutdown) Get() *VmShutdown {
	return v.value
}

func (v *NullableVmShutdown) Set(val *VmShutdown) {
	v.value = val
	v.isSet = true
}

func (v NullableVmShutdown) IsSet() bool {
	return v.isSet
}

func (v *NullableVmShutdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmShutdown(val *VmShutdown) *NullableVmShutdown {
	return &NullableVmShutdown{value: val, isSet: true}
}

func (v NullableVmShutdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmShutdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
