# SecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | Direction of the network traffic | 
**Protocol** | **string** | Network protocol | 
**Ports** | **string** | Allowed port or port range | 
**IpRange** | **string** | Allowed IP or IP range | 

## Methods

### NewSecurityGroupRuleRequest

`func NewSecurityGroupRuleRequest(direction string, protocol string, ports string, ipRange string, ) *SecurityGroupRuleRequest`

NewSecurityGroupRuleRequest instantiates a new SecurityGroupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleRequestWithDefaults

`func NewSecurityGroupRuleRequestWithDefaults() *SecurityGroupRuleRequest`

NewSecurityGroupRuleRequestWithDefaults instantiates a new SecurityGroupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SecurityGroupRuleRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRuleRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRuleRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetProtocol

`func (o *SecurityGroupRuleRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRuleRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRuleRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetPorts

`func (o *SecurityGroupRuleRequest) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SecurityGroupRuleRequest) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SecurityGroupRuleRequest) SetPorts(v string)`

SetPorts sets Ports field to given value.


### GetIpRange

`func (o *SecurityGroupRuleRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SecurityGroupRuleRequest) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SecurityGroupRuleRequest) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


