# VmNew

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
**Disks** | Pointer to [**[]VmNewDisksInner**](VmNewDisksInner.md) | List of volumes | [optional] 
**Networks** | Pointer to [**[]VmNetworksInner**](VmNetworksInner.md) |  | [optional] 
**SecurityGroup** | Pointer to [**VmSecurityGroup**](VmSecurityGroup.md) |  | [optional] 
**Subnetwork** | Pointer to [**VmSubnetwork**](VmSubnetwork.md) |  | [optional] 
**SshKeyId** | Pointer to **int32** | SSH key ID | [optional] 
**UserName** | Pointer to **string** | User name | [optional] 
**UserPassword** | Pointer to **string** | User password | [optional] 
**Cost** | Pointer to [**VmCost**](VmCost.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewVmNew

`func NewVmNew() *VmNew`

NewVmNew instantiates a new VmNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmNewWithDefaults

`func NewVmNewWithDefaults() *VmNew`

NewVmNewWithDefaults instantiates a new VmNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmNew) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmNew) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmNew) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VmNew) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VmNew) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VmNew) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VmNew) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VmNew) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *VmNew) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *VmNew) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *VmNew) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *VmNew) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedById

`func (o *VmNew) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *VmNew) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *VmNew) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *VmNew) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedAt

`func (o *VmNew) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VmNew) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VmNew) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *VmNew) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *VmNew) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *VmNew) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *VmNew) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *VmNew) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *VmNew) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *VmNew) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *VmNew) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *VmNew) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *VmNew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmNew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmNew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmNew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *VmNew) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *VmNew) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *VmNew) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *VmNew) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatus

`func (o *VmNew) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VmNew) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VmNew) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VmNew) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *VmNew) GetProvider() VmProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VmNew) GetProviderOk() (*VmProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VmNew) SetProvider(v VmProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VmNew) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *VmNew) GetLocation() VmLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VmNew) GetLocationOk() (*VmLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VmNew) SetLocation(v VmLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VmNew) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *VmNew) GetDataCenter() VmDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *VmNew) GetDataCenterOk() (*VmDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *VmNew) SetDataCenter(v VmDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *VmNew) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *VmNew) GetOs() VmOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *VmNew) GetOsOk() (*VmOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *VmNew) SetOs(v VmOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *VmNew) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetVCpu

`func (o *VmNew) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmNew) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmNew) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *VmNew) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetVCpuType

`func (o *VmNew) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmNew) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmNew) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *VmNew) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *VmNew) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *VmNew) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *VmNew) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *VmNew) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetRamGb

`func (o *VmNew) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmNew) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmNew) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *VmNew) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *VmNew) GetDisks() []VmNewDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VmNew) GetDisksOk() (*[]VmNewDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VmNew) SetDisks(v []VmNewDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VmNew) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *VmNew) GetNetworks() []VmNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VmNew) GetNetworksOk() (*[]VmNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VmNew) SetNetworks(v []VmNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VmNew) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *VmNew) GetSecurityGroup() VmSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *VmNew) GetSecurityGroupOk() (*VmSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *VmNew) SetSecurityGroup(v VmSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *VmNew) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetSubnetwork

`func (o *VmNew) GetSubnetwork() VmSubnetwork`

GetSubnetwork returns the Subnetwork field if non-nil, zero value otherwise.

### GetSubnetworkOk

`func (o *VmNew) GetSubnetworkOk() (*VmSubnetwork, bool)`

GetSubnetworkOk returns a tuple with the Subnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetwork

`func (o *VmNew) SetSubnetwork(v VmSubnetwork)`

SetSubnetwork sets Subnetwork field to given value.

### HasSubnetwork

`func (o *VmNew) HasSubnetwork() bool`

HasSubnetwork returns a boolean if a field has been set.

### GetSshKeyId

`func (o *VmNew) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *VmNew) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *VmNew) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *VmNew) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetUserName

`func (o *VmNew) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VmNew) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VmNew) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VmNew) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPassword

`func (o *VmNew) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *VmNew) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *VmNew) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *VmNew) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetCost

`func (o *VmNew) GetCost() VmCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *VmNew) GetCostOk() (*VmCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *VmNew) SetCost(v VmCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *VmNew) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetTags

`func (o *VmNew) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VmNew) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VmNew) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VmNew) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


