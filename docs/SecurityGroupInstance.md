# SecurityGroupInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByName** | Pointer to **string** |  | [optional] 
**CreatedById** | Pointer to **int32** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**ModifiedByName** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**SecurityGroupInstanceProvider**](SecurityGroupInstanceProvider.md) |  | [optional] 
**Location** | Pointer to [**SecurityGroupInstanceLocation**](SecurityGroupInstanceLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**SecurityGroupInstanceDataCenter**](SecurityGroupInstanceDataCenter.md) |  | [optional] 
**Os** | Pointer to [**SecurityGroupInstanceOs**](SecurityGroupInstanceOs.md) |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**RamGb** | Pointer to **int32** |  | [optional] 
**Disks** | Pointer to [**[]SecurityGroupInstanceDisksInner**](SecurityGroupInstanceDisksInner.md) |  | [optional] 
**Networks** | Pointer to [**[]SecurityGroupInstanceNetworksInner**](SecurityGroupInstanceNetworksInner.md) |  | [optional] 
**SecurityGroup** | Pointer to [**SecurityGroupInstanceSecurityGroup**](SecurityGroupInstanceSecurityGroup.md) |  | [optional] 
**SshKeyId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to [**SecurityGroupInstanceCost**](SecurityGroupInstanceCost.md) |  | [optional] 

## Methods

### NewSecurityGroupInstance

`func NewSecurityGroupInstance() *SecurityGroupInstance`

NewSecurityGroupInstance instantiates a new SecurityGroupInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupInstanceWithDefaults

`func NewSecurityGroupInstanceWithDefaults() *SecurityGroupInstance`

NewSecurityGroupInstanceWithDefaults instantiates a new SecurityGroupInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecurityGroupInstance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityGroupInstance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityGroupInstance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityGroupInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *SecurityGroupInstance) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *SecurityGroupInstance) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *SecurityGroupInstance) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *SecurityGroupInstance) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedById

`func (o *SecurityGroupInstance) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SecurityGroupInstance) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SecurityGroupInstance) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SecurityGroupInstance) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SecurityGroupInstance) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SecurityGroupInstance) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SecurityGroupInstance) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SecurityGroupInstance) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *SecurityGroupInstance) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *SecurityGroupInstance) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *SecurityGroupInstance) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *SecurityGroupInstance) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *SecurityGroupInstance) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *SecurityGroupInstance) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *SecurityGroupInstance) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *SecurityGroupInstance) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *SecurityGroupInstance) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SecurityGroupInstance) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SecurityGroupInstance) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SecurityGroupInstance) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityGroupInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityGroupInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityGroupInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityGroupInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *SecurityGroupInstance) GetProvider() SecurityGroupInstanceProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SecurityGroupInstance) GetProviderOk() (*SecurityGroupInstanceProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SecurityGroupInstance) SetProvider(v SecurityGroupInstanceProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SecurityGroupInstance) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *SecurityGroupInstance) GetLocation() SecurityGroupInstanceLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SecurityGroupInstance) GetLocationOk() (*SecurityGroupInstanceLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SecurityGroupInstance) SetLocation(v SecurityGroupInstanceLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SecurityGroupInstance) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *SecurityGroupInstance) GetDataCenter() SecurityGroupInstanceDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *SecurityGroupInstance) GetDataCenterOk() (*SecurityGroupInstanceDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *SecurityGroupInstance) SetDataCenter(v SecurityGroupInstanceDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *SecurityGroupInstance) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *SecurityGroupInstance) GetOs() SecurityGroupInstanceOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *SecurityGroupInstance) GetOsOk() (*SecurityGroupInstanceOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *SecurityGroupInstance) SetOs(v SecurityGroupInstanceOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *SecurityGroupInstance) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCpu

`func (o *SecurityGroupInstance) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *SecurityGroupInstance) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *SecurityGroupInstance) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *SecurityGroupInstance) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *SecurityGroupInstance) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *SecurityGroupInstance) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *SecurityGroupInstance) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *SecurityGroupInstance) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *SecurityGroupInstance) GetDisks() []SecurityGroupInstanceDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *SecurityGroupInstance) GetDisksOk() (*[]SecurityGroupInstanceDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *SecurityGroupInstance) SetDisks(v []SecurityGroupInstanceDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *SecurityGroupInstance) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *SecurityGroupInstance) GetNetworks() []SecurityGroupInstanceNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *SecurityGroupInstance) GetNetworksOk() (*[]SecurityGroupInstanceNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *SecurityGroupInstance) SetNetworks(v []SecurityGroupInstanceNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *SecurityGroupInstance) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *SecurityGroupInstance) GetSecurityGroup() SecurityGroupInstanceSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *SecurityGroupInstance) GetSecurityGroupOk() (*SecurityGroupInstanceSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *SecurityGroupInstance) SetSecurityGroup(v SecurityGroupInstanceSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *SecurityGroupInstance) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetSshKeyId

`func (o *SecurityGroupInstance) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *SecurityGroupInstance) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *SecurityGroupInstance) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *SecurityGroupInstance) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetUserName

`func (o *SecurityGroupInstance) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SecurityGroupInstance) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SecurityGroupInstance) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SecurityGroupInstance) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetCost

`func (o *SecurityGroupInstance) GetCost() SecurityGroupInstanceCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SecurityGroupInstance) GetCostOk() (*SecurityGroupInstanceCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SecurityGroupInstance) SetCost(v SecurityGroupInstanceCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SecurityGroupInstance) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


