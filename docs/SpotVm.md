# SpotVm

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
**Provider** | Pointer to [**SpotVmProvider**](SpotVmProvider.md) |  | [optional] 
**Location** | Pointer to [**VmLocation**](VmLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**SpotVmDataCenter**](SpotVmDataCenter.md) |  | [optional] 
**Os** | Pointer to [**SpotVmOs**](SpotVmOs.md) |  | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**VCpuType** | Pointer to **string** | Type of virtual Central Processing Units (vCPUs) | [optional] 
**CloudNetworkType** | Pointer to **string** | Cloud network type | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Disks** | Pointer to [**[]SpotVmDisksInner**](SpotVmDisksInner.md) | List of volumes | [optional] 
**Networks** | Pointer to [**[]KubernetesNodeGroupsInnerNodesInnerNetworksInner**](KubernetesNodeGroupsInnerNodesInnerNetworksInner.md) |  | [optional] 
**SecurityGroup** | Pointer to [**VmSecurityGroup**](VmSecurityGroup.md) |  | [optional] 
**SshKeyId** | Pointer to **int32** | SSH key ID | [optional] 
**UserName** | Pointer to **string** | User name | [optional] 
**UserPassword** | Pointer to **string** | Administrator password in the Windows operating system | [optional] 
**Cost** | Pointer to [**SpotVmCost**](SpotVmCost.md) |  | [optional] 

## Methods

### NewSpotVm

`func NewSpotVm() *SpotVm`

NewSpotVm instantiates a new SpotVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotVmWithDefaults

`func NewSpotVmWithDefaults() *SpotVm`

NewSpotVmWithDefaults instantiates a new SpotVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpotVm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpotVm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpotVm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SpotVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SpotVm) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpotVm) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpotVm) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpotVm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *SpotVm) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *SpotVm) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *SpotVm) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *SpotVm) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedById

`func (o *SpotVm) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SpotVm) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SpotVm) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SpotVm) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SpotVm) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SpotVm) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SpotVm) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SpotVm) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *SpotVm) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *SpotVm) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *SpotVm) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *SpotVm) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *SpotVm) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *SpotVm) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *SpotVm) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *SpotVm) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *SpotVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpotVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpotVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpotVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *SpotVm) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SpotVm) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SpotVm) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SpotVm) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatus

`func (o *SpotVm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpotVm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpotVm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpotVm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *SpotVm) GetProvider() SpotVmProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SpotVm) GetProviderOk() (*SpotVmProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SpotVm) SetProvider(v SpotVmProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SpotVm) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *SpotVm) GetLocation() VmLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SpotVm) GetLocationOk() (*VmLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SpotVm) SetLocation(v VmLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SpotVm) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *SpotVm) GetDataCenter() SpotVmDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *SpotVm) GetDataCenterOk() (*SpotVmDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *SpotVm) SetDataCenter(v SpotVmDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *SpotVm) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *SpotVm) GetOs() SpotVmOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *SpotVm) GetOsOk() (*SpotVmOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *SpotVm) SetOs(v SpotVmOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *SpotVm) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetVCpu

`func (o *SpotVm) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *SpotVm) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *SpotVm) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *SpotVm) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetVCpuType

`func (o *SpotVm) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *SpotVm) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *SpotVm) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *SpotVm) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *SpotVm) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *SpotVm) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *SpotVm) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *SpotVm) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetRamGb

`func (o *SpotVm) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *SpotVm) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *SpotVm) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *SpotVm) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *SpotVm) GetDisks() []SpotVmDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *SpotVm) GetDisksOk() (*[]SpotVmDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *SpotVm) SetDisks(v []SpotVmDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *SpotVm) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *SpotVm) GetNetworks() []KubernetesNodeGroupsInnerNodesInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *SpotVm) GetNetworksOk() (*[]KubernetesNodeGroupsInnerNodesInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *SpotVm) SetNetworks(v []KubernetesNodeGroupsInnerNodesInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *SpotVm) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *SpotVm) GetSecurityGroup() VmSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *SpotVm) GetSecurityGroupOk() (*VmSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *SpotVm) SetSecurityGroup(v VmSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *SpotVm) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetSshKeyId

`func (o *SpotVm) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *SpotVm) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *SpotVm) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *SpotVm) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetUserName

`func (o *SpotVm) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SpotVm) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SpotVm) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SpotVm) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPassword

`func (o *SpotVm) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *SpotVm) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *SpotVm) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *SpotVm) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetCost

`func (o *SpotVm) GetCost() SpotVmCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SpotVm) GetCostOk() (*SpotVmCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SpotVm) SetCost(v SpotVmCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SpotVm) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


