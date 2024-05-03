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

// checks if the SecurityGroupInstanceCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupInstanceCost{}

// SecurityGroupInstanceCost struct for SecurityGroupInstanceCost
type SecurityGroupInstanceCost struct {
	Unit     *string  `json:"unit,omitempty"`
	Currency *string  `json:"currency,omitempty"`
	Price    *float32 `json:"price,omitempty"`
}

// NewSecurityGroupInstanceCost instantiates a new SecurityGroupInstanceCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupInstanceCost() *SecurityGroupInstanceCost {
	this := SecurityGroupInstanceCost{}
	return &this
}

// NewSecurityGroupInstanceCostWithDefaults instantiates a new SecurityGroupInstanceCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupInstanceCostWithDefaults() *SecurityGroupInstanceCost {
	this := SecurityGroupInstanceCost{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *SecurityGroupInstanceCost) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstanceCost) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *SecurityGroupInstanceCost) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *SecurityGroupInstanceCost) SetUnit(v string) {
	o.Unit = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SecurityGroupInstanceCost) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstanceCost) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SecurityGroupInstanceCost) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SecurityGroupInstanceCost) SetCurrency(v string) {
	o.Currency = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SecurityGroupInstanceCost) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstanceCost) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SecurityGroupInstanceCost) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *SecurityGroupInstanceCost) SetPrice(v float32) {
	o.Price = &v
}

func (o SecurityGroupInstanceCost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupInstanceCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableSecurityGroupInstanceCost struct {
	value *SecurityGroupInstanceCost
	isSet bool
}

func (v NullableSecurityGroupInstanceCost) Get() *SecurityGroupInstanceCost {
	return v.value
}

func (v *NullableSecurityGroupInstanceCost) Set(val *SecurityGroupInstanceCost) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupInstanceCost) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupInstanceCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupInstanceCost(val *SecurityGroupInstanceCost) *NullableSecurityGroupInstanceCost {
	return &NullableSecurityGroupInstanceCost{value: val, isSet: true}
}

func (v NullableSecurityGroupInstanceCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupInstanceCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
