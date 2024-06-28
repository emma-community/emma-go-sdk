# VmSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the security group | [optional] 
**Name** | Pointer to **string** | Name of the security group | [optional] 

## Methods

### NewVmSecurityGroup

`func NewVmSecurityGroup() *VmSecurityGroup`

NewVmSecurityGroup instantiates a new VmSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSecurityGroupWithDefaults

`func NewVmSecurityGroupWithDefaults() *VmSecurityGroup`

NewVmSecurityGroupWithDefaults instantiates a new VmSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmSecurityGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmSecurityGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmSecurityGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VmSecurityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VmSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmSecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmSecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


