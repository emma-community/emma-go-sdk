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

// checks if the ConflictError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConflictError{}

// ConflictError struct for ConflictError
type ConflictError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type _ConflictError ConflictError

// NewConflictError instantiates a new ConflictError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictError(code string, message string) *ConflictError {
	this := ConflictError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewConflictErrorWithDefaults instantiates a new ConflictError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictErrorWithDefaults() *ConflictError {
	this := ConflictError{}
	return &this
}

// GetCode returns the Code field value
func (o *ConflictError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ConflictError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ConflictError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ConflictError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ConflictError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ConflictError) SetMessage(v string) {
	o.Message = v
}

func (o ConflictError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConflictError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ConflictError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varConflictError := _ConflictError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConflictError)

	if err != nil {
		return err
	}

	*o = ConflictError(varConflictError)

	return err
}

type NullableConflictError struct {
	value *ConflictError
	isSet bool
}

func (v NullableConflictError) Get() *ConflictError {
	return v.value
}

func (v *NullableConflictError) Set(val *ConflictError) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictError) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictError(val *ConflictError) *NullableConflictError {
	return &NullableConflictError{value: val, isSet: true}
}

func (v NullableConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
