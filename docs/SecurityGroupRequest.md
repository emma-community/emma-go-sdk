# SecurityGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the security group | 
**Rules** | [**[]SecurityGroupRuleRequest**](SecurityGroupRuleRequest.md) |  | 

## Methods

### NewSecurityGroupRequest

`func NewSecurityGroupRequest(name string, rules []SecurityGroupRuleRequest, ) *SecurityGroupRequest`

NewSecurityGroupRequest instantiates a new SecurityGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRequestWithDefaults

`func NewSecurityGroupRequestWithDefaults() *SecurityGroupRequest`

NewSecurityGroupRequestWithDefaults instantiates a new SecurityGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *SecurityGroupRequest) GetRules() []SecurityGroupRuleRequest`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroupRequest) GetRulesOk() (*[]SecurityGroupRuleRequest, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroupRequest) SetRules(v []SecurityGroupRuleRequest)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


