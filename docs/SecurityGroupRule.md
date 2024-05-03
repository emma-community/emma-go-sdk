# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **string** |  | [optional] 
**IpRange** | Pointer to **string** |  | [optional] 

## Methods

### NewSecurityGroupRule

`func NewSecurityGroupRule() *SecurityGroupRule`

NewSecurityGroupRule instantiates a new SecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleWithDefaults

`func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule`

NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecurityGroupRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityGroupRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityGroupRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityGroupRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPolicy

`func (o *SecurityGroupRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SecurityGroupRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SecurityGroupRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *SecurityGroupRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetDirection

`func (o *SecurityGroupRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SecurityGroupRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProtocol

`func (o *SecurityGroupRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityGroupRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPorts

`func (o *SecurityGroupRule) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SecurityGroupRule) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SecurityGroupRule) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SecurityGroupRule) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetIpRange

`func (o *SecurityGroupRule) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SecurityGroupRule) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SecurityGroupRule) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *SecurityGroupRule) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


