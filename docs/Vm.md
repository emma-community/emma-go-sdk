# Vm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **interface{}** |  | [optional] 
**CreatedByName** | Pointer to **interface{}** |  | [optional] 
**CreatedById** | Pointer to **interface{}** |  | [optional] 
**ModifiedAt** | Pointer to **interface{}** |  | [optional] 
**ModifiedByName** | Pointer to **interface{}** |  | [optional] 
**ModifiedById** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **interface{}** |  | [optional] 
**ProjectId** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **interface{}** |  | [optional] 
**Provider** | Pointer to [**VmProvider**](VmProvider.md) |  | [optional] 
**Location** | Pointer to [**VmLocation**](VmLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**VmDataCenter**](VmDataCenter.md) |  | [optional] 
**Os** | Pointer to [**VmOs**](VmOs.md) |  | [optional] 
**Cpu** | Pointer to **interface{}** |  | [optional] 
**RamGb** | Pointer to **interface{}** |  | [optional] 
**Disks** | Pointer to [**[]VmDisksInner**](VmDisksInner.md) |  | [optional] 
**Networks** | Pointer to [**[]VmNetworksInner**](VmNetworksInner.md) |  | [optional] 
**SecurityGroup** | Pointer to [**VmSecurityGroup**](VmSecurityGroup.md) |  | [optional] 
**SshKeyId** | Pointer to **interface{}** |  | [optional] 
**UserName** | Pointer to **interface{}** |  | [optional] 
**Cost** | Pointer to [**VmCost**](VmCost.md) |  | [optional] 

## Methods

### NewVm

`func NewVm() *Vm`

NewVm instantiates a new Vm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmWithDefaults

`func NewVmWithDefaults() *Vm`

NewVmWithDefaults instantiates a new Vm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vm) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vm) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vm) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *Vm) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Vm) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Vm) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedAt

`func (o *Vm) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vm) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vm) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Vm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Vm) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Vm) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedByName

`func (o *Vm) GetCreatedByName() interface{}`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *Vm) GetCreatedByNameOk() (*interface{}, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *Vm) SetCreatedByName(v interface{})`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *Vm) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### SetCreatedByNameNil

`func (o *Vm) SetCreatedByNameNil(b bool)`

 SetCreatedByNameNil sets the value for CreatedByName to be an explicit nil

### UnsetCreatedByName
`func (o *Vm) UnsetCreatedByName()`

UnsetCreatedByName ensures that no value is present for CreatedByName, not even an explicit nil
### GetCreatedById

`func (o *Vm) GetCreatedById() interface{}`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Vm) GetCreatedByIdOk() (*interface{}, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Vm) SetCreatedById(v interface{})`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Vm) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### SetCreatedByIdNil

`func (o *Vm) SetCreatedByIdNil(b bool)`

 SetCreatedByIdNil sets the value for CreatedById to be an explicit nil

### UnsetCreatedById
`func (o *Vm) UnsetCreatedById()`

UnsetCreatedById ensures that no value is present for CreatedById, not even an explicit nil
### GetModifiedAt

`func (o *Vm) GetModifiedAt() interface{}`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Vm) GetModifiedAtOk() (*interface{}, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Vm) SetModifiedAt(v interface{})`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Vm) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *Vm) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Vm) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetModifiedByName

`func (o *Vm) GetModifiedByName() interface{}`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *Vm) GetModifiedByNameOk() (*interface{}, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *Vm) SetModifiedByName(v interface{})`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *Vm) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### SetModifiedByNameNil

`func (o *Vm) SetModifiedByNameNil(b bool)`

 SetModifiedByNameNil sets the value for ModifiedByName to be an explicit nil

### UnsetModifiedByName
`func (o *Vm) UnsetModifiedByName()`

UnsetModifiedByName ensures that no value is present for ModifiedByName, not even an explicit nil
### GetModifiedById

`func (o *Vm) GetModifiedById() interface{}`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *Vm) GetModifiedByIdOk() (*interface{}, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *Vm) SetModifiedById(v interface{})`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *Vm) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *Vm) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *Vm) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetName

`func (o *Vm) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vm) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vm) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *Vm) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Vm) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Vm) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *Vm) GetProjectId() interface{}`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Vm) GetProjectIdOk() (*interface{}, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Vm) SetProjectId(v interface{})`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Vm) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Vm) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Vm) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetStatus

`func (o *Vm) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Vm) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Vm) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Vm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Vm) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Vm) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetProvider

`func (o *Vm) GetProvider() VmProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Vm) GetProviderOk() (*VmProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Vm) SetProvider(v VmProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Vm) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *Vm) GetLocation() VmLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Vm) GetLocationOk() (*VmLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Vm) SetLocation(v VmLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Vm) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *Vm) GetDataCenter() VmDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *Vm) GetDataCenterOk() (*VmDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *Vm) SetDataCenter(v VmDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *Vm) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *Vm) GetOs() VmOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Vm) GetOsOk() (*VmOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Vm) SetOs(v VmOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *Vm) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCpu

`func (o *Vm) GetCpu() interface{}`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Vm) GetCpuOk() (*interface{}, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Vm) SetCpu(v interface{})`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Vm) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *Vm) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *Vm) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetRamGb

`func (o *Vm) GetRamGb() interface{}`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *Vm) GetRamGbOk() (*interface{}, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *Vm) SetRamGb(v interface{})`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *Vm) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### SetRamGbNil

`func (o *Vm) SetRamGbNil(b bool)`

 SetRamGbNil sets the value for RamGb to be an explicit nil

### UnsetRamGb
`func (o *Vm) UnsetRamGb()`

UnsetRamGb ensures that no value is present for RamGb, not even an explicit nil
### GetDisks

`func (o *Vm) GetDisks() []VmDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Vm) GetDisksOk() (*[]VmDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Vm) SetDisks(v []VmDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Vm) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *Vm) GetNetworks() []VmNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Vm) GetNetworksOk() (*[]VmNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Vm) SetNetworks(v []VmNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *Vm) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *Vm) GetSecurityGroup() VmSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *Vm) GetSecurityGroupOk() (*VmSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *Vm) SetSecurityGroup(v VmSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *Vm) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetSshKeyId

`func (o *Vm) GetSshKeyId() interface{}`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *Vm) GetSshKeyIdOk() (*interface{}, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *Vm) SetSshKeyId(v interface{})`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *Vm) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### SetSshKeyIdNil

`func (o *Vm) SetSshKeyIdNil(b bool)`

 SetSshKeyIdNil sets the value for SshKeyId to be an explicit nil

### UnsetSshKeyId
`func (o *Vm) UnsetSshKeyId()`

UnsetSshKeyId ensures that no value is present for SshKeyId, not even an explicit nil
### GetUserName

`func (o *Vm) GetUserName() interface{}`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Vm) GetUserNameOk() (*interface{}, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Vm) SetUserName(v interface{})`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Vm) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *Vm) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *Vm) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetCost

`func (o *Vm) GetCost() VmCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Vm) GetCostOk() (*VmCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Vm) SetCost(v VmCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Vm) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


