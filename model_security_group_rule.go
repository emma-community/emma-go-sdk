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

// checks if the SecurityGroupRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupRule{}

// SecurityGroupRule struct for SecurityGroupRule
type SecurityGroupRule struct {
	Type      *string `json:"type,omitempty"`
	Policy    *string `json:"policy,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Protocol  *string `json:"protocol,omitempty"`
	Ports     *string `json:"ports,omitempty"`
	IpRange   *string `json:"ipRange,omitempty"`
	IsMutable *bool   `json:"isMutable,omitempty"`
}

// NewSecurityGroupRule instantiates a new SecurityGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupRule() *SecurityGroupRule {
	this := SecurityGroupRule{}
	return &this
}

// NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule {
	this := SecurityGroupRule{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityGroupRule) SetType(v string) {
	o.Type = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *SecurityGroupRule) SetPolicy(v string) {
	o.Policy = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *SecurityGroupRule) SetDirection(v string) {
	o.Direction = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SecurityGroupRule) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetPorts() string {
	if o == nil || IsNil(o.Ports) {
		var ret string
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetPortsOk() (*string, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given string and assigns it to the Ports field.
func (o *SecurityGroupRule) SetPorts(v string) {
	o.Ports = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIpRange() string {
	if o == nil || IsNil(o.IpRange) {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIpRangeOk() (*string, bool) {
	if o == nil || IsNil(o.IpRange) {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIpRange() bool {
	if o != nil && !IsNil(o.IpRange) {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *SecurityGroupRule) SetIpRange(v string) {
	o.IpRange = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIsMutable() bool {
	if o == nil || IsNil(o.IsMutable) {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIsMutableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMutable) {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIsMutable() bool {
	if o != nil && !IsNil(o.IsMutable) {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *SecurityGroupRule) SetIsMutable(v bool) {
	o.IsMutable = &v
}

func (o SecurityGroupRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.IpRange) {
		toSerialize["ipRange"] = o.IpRange
	}
	if !IsNil(o.IsMutable) {
		toSerialize["isMutable"] = o.IsMutable
	}
	return toSerialize, nil
}

type NullableSecurityGroupRule struct {
	value *SecurityGroupRule
	isSet bool
}

func (v NullableSecurityGroupRule) Get() *SecurityGroupRule {
	return v.value
}

func (v *NullableSecurityGroupRule) Set(val *SecurityGroupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupRule(val *SecurityGroupRule) *NullableSecurityGroupRule {
	return &NullableSecurityGroupRule{value: val, isSet: true}
}

func (v NullableSecurityGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
