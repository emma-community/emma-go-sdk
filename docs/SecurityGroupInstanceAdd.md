# SecurityGroupInstanceAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **int32** | ID of the compute instance to be added to the Security group | [optional] 

## Methods

### NewSecurityGroupInstanceAdd

`func NewSecurityGroupInstanceAdd() *SecurityGroupInstanceAdd`

NewSecurityGroupInstanceAdd instantiates a new SecurityGroupInstanceAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupInstanceAddWithDefaults

`func NewSecurityGroupInstanceAddWithDefaults() *SecurityGroupInstanceAdd`

NewSecurityGroupInstanceAddWithDefaults instantiates a new SecurityGroupInstanceAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *SecurityGroupInstanceAdd) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *SecurityGroupInstanceAdd) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *SecurityGroupInstanceAdd) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *SecurityGroupInstanceAdd) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


