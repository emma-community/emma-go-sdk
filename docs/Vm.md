# Vm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the virtual machine | [optional] 
**CreatedAt** | Pointer to **string** | Date and time when the virtual machine was created | [optional] 
**CreatedByName** | Pointer to **string** | Name of the user who created the virtual machine | [optional] 
**CreatedById** | Pointer to **int32** | ID of the user who created the virtual machine | [optional] 
**ModifiedAt** | Pointer to **string** | Date and time when the virtual machine was last edited | [optional] 
**ModifiedByName** | Pointer to **string** | Name of the user who last edited the virtual machine | [optional] 
**ModifiedById** | Pointer to **int32** | ID of the user who last edited the virtual machine | [optional] 
**Name** | Pointer to **string** | Name of the virtual machine | [optional] 
**ProjectId** | Pointer to **int32** | Project ID | [optional] 
**Status** | Pointer to **string** | Status of the virtual machine | [optional] 
**Provider** | Pointer to [**VmProvider**](VmProvider.md) |  | [optional] 
**Location** | Pointer to [**VmLocation**](VmLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**VmDataCenter**](VmDataCenter.md) |  | [optional] 
**Os** | Pointer to [**VmOs**](VmOs.md) |  | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**VCpuType** | Pointer to **string** | Type of virtual Central Processing Units (vCPUs) | [optional] 
**CloudNetworkType** | Pointer to **string** | Cloud network type | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Disks** | Pointer to [**[]VmDisksInner**](VmDisksInner.md) | List of volumes | [optional] 
**Networks** | Pointer to [**[]VmNetworksInner**](VmNetworksInner.md) |  | [optional] 
**SecurityGroup** | Pointer to [**VmSecurityGroup**](VmSecurityGroup.md) |  | [optional] 
**Subnetwork** | Pointer to [**VmSubnetwork**](VmSubnetwork.md) |  | [optional] 
**SshKeyId** | Pointer to **int32** | SSH key ID | [optional] 
**UserName** | Pointer to **string** | User name | [optional] 
**UserPassword** | Pointer to **string** | User password | [optional] 
**Cost** | Pointer to [**VmCost**](VmCost.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

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

`func (o *Vm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Vm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Vm) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vm) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vm) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Vm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *Vm) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *Vm) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *Vm) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *Vm) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedById

`func (o *Vm) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Vm) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Vm) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Vm) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Vm) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Vm) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Vm) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Vm) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *Vm) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *Vm) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *Vm) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *Vm) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *Vm) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *Vm) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *Vm) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *Vm) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *Vm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *Vm) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Vm) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Vm) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Vm) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatus

`func (o *Vm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Vm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Vm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Vm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### GetVCpu

`func (o *Vm) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *Vm) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *Vm) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *Vm) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetVCpuType

`func (o *Vm) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *Vm) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *Vm) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *Vm) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *Vm) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *Vm) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *Vm) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *Vm) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetRamGb

`func (o *Vm) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *Vm) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *Vm) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *Vm) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

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

### GetSubnetwork

`func (o *Vm) GetSubnetwork() VmSubnetwork`

GetSubnetwork returns the Subnetwork field if non-nil, zero value otherwise.

### GetSubnetworkOk

`func (o *Vm) GetSubnetworkOk() (*VmSubnetwork, bool)`

GetSubnetworkOk returns a tuple with the Subnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetwork

`func (o *Vm) SetSubnetwork(v VmSubnetwork)`

SetSubnetwork sets Subnetwork field to given value.

### HasSubnetwork

`func (o *Vm) HasSubnetwork() bool`

HasSubnetwork returns a boolean if a field has been set.

### GetSshKeyId

`func (o *Vm) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *Vm) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *Vm) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *Vm) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetUserName

`func (o *Vm) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Vm) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Vm) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Vm) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPassword

`func (o *Vm) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *Vm) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *Vm) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *Vm) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

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

### GetTags

`func (o *Vm) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Vm) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Vm) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Vm) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


