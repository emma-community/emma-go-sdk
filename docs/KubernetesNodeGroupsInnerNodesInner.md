# KubernetesNodeGroupsInnerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the worker node | [optional] 
**CreatedAt** | Pointer to **string** | Date and time when the worker node was created | [optional] 
**Name** | Pointer to **string** | Name of the worker node | [optional] 
**Status** | Pointer to **string** | Status of the worker node | [optional] 
**Provider** | Pointer to [**KubernetesNodeGroupsInnerNodesInnerProvider**](KubernetesNodeGroupsInnerNodesInnerProvider.md) |  | [optional] 
**Location** | Pointer to [**KubernetesNodeGroupsInnerNodesInnerLocation**](KubernetesNodeGroupsInnerNodesInnerLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**KubernetesNodeGroupsInnerNodesInnerDataCenter**](KubernetesNodeGroupsInnerNodesInnerDataCenter.md) |  | [optional] 
**Os** | Pointer to [**KubernetesNodeGroupsInnerNodesInnerOs**](KubernetesNodeGroupsInnerNodesInnerOs.md) |  | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**VCpuType** | Pointer to **string** | Type of the virtual Central Processing Unit | [optional] 
**CloudNetworkType** | Pointer to **string** | Type of the cloud network (multi-cloud, isolated, or default) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Disks** | Pointer to [**[]KubernetesNodeGroupsInnerNodesInnerDisksInner**](KubernetesNodeGroupsInnerNodesInnerDisksInner.md) | Volumes attached | [optional] 
**Networks** | Pointer to [**[]KubernetesNodeGroupsInnerNodesInnerNetworksInner**](KubernetesNodeGroupsInnerNodesInnerNetworksInner.md) | Public and private networks | [optional] 
**SecurityGroup** | Pointer to [**KubernetesNodeGroupsInnerNodesInnerSecurityGroup**](KubernetesNodeGroupsInnerNodesInnerSecurityGroup.md) |  | [optional] 
**Cost** | Pointer to [**KubernetesNodeGroupsInnerNodesInnerCost**](KubernetesNodeGroupsInnerNodesInnerCost.md) |  | [optional] 

## Methods

### NewKubernetesNodeGroupsInnerNodesInner

`func NewKubernetesNodeGroupsInnerNodesInner() *KubernetesNodeGroupsInnerNodesInner`

NewKubernetesNodeGroupsInnerNodesInner instantiates a new KubernetesNodeGroupsInnerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeGroupsInnerNodesInnerWithDefaults

`func NewKubernetesNodeGroupsInnerNodesInnerWithDefaults() *KubernetesNodeGroupsInnerNodesInner`

NewKubernetesNodeGroupsInnerNodesInnerWithDefaults instantiates a new KubernetesNodeGroupsInnerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNodeGroupsInnerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNodeGroupsInnerNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesNodeGroupsInnerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubernetesNodeGroupsInnerNodesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesNodeGroupsInnerNodesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesNodeGroupsInnerNodesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *KubernetesNodeGroupsInnerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodeGroupsInnerNodesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesNodeGroupsInnerNodesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesNodeGroupsInnerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesNodeGroupsInnerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesNodeGroupsInnerNodesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *KubernetesNodeGroupsInnerNodesInner) GetProvider() KubernetesNodeGroupsInnerNodesInnerProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetProviderOk() (*KubernetesNodeGroupsInnerNodesInnerProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KubernetesNodeGroupsInnerNodesInner) SetProvider(v KubernetesNodeGroupsInnerNodesInnerProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KubernetesNodeGroupsInnerNodesInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *KubernetesNodeGroupsInnerNodesInner) GetLocation() KubernetesNodeGroupsInnerNodesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetLocationOk() (*KubernetesNodeGroupsInnerNodesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesNodeGroupsInnerNodesInner) SetLocation(v KubernetesNodeGroupsInnerNodesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KubernetesNodeGroupsInnerNodesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *KubernetesNodeGroupsInnerNodesInner) GetDataCenter() KubernetesNodeGroupsInnerNodesInnerDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetDataCenterOk() (*KubernetesNodeGroupsInnerNodesInnerDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *KubernetesNodeGroupsInnerNodesInner) SetDataCenter(v KubernetesNodeGroupsInnerNodesInnerDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *KubernetesNodeGroupsInnerNodesInner) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *KubernetesNodeGroupsInnerNodesInner) GetOs() KubernetesNodeGroupsInnerNodesInnerOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetOsOk() (*KubernetesNodeGroupsInnerNodesInnerOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *KubernetesNodeGroupsInnerNodesInner) SetOs(v KubernetesNodeGroupsInnerNodesInnerOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *KubernetesNodeGroupsInnerNodesInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetVCpu

`func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesNodeGroupsInnerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *KubernetesNodeGroupsInnerNodesInner) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetVCpuType

`func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesNodeGroupsInnerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *KubernetesNodeGroupsInnerNodesInner) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *KubernetesNodeGroupsInnerNodesInner) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *KubernetesNodeGroupsInnerNodesInner) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *KubernetesNodeGroupsInnerNodesInner) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetRamGb

`func (o *KubernetesNodeGroupsInnerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesNodeGroupsInnerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *KubernetesNodeGroupsInnerNodesInner) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *KubernetesNodeGroupsInnerNodesInner) GetDisks() []KubernetesNodeGroupsInnerNodesInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetDisksOk() (*[]KubernetesNodeGroupsInnerNodesInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *KubernetesNodeGroupsInnerNodesInner) SetDisks(v []KubernetesNodeGroupsInnerNodesInnerDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *KubernetesNodeGroupsInnerNodesInner) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *KubernetesNodeGroupsInnerNodesInner) GetNetworks() []KubernetesNodeGroupsInnerNodesInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetNetworksOk() (*[]KubernetesNodeGroupsInnerNodesInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *KubernetesNodeGroupsInnerNodesInner) SetNetworks(v []KubernetesNodeGroupsInnerNodesInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *KubernetesNodeGroupsInnerNodesInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *KubernetesNodeGroupsInnerNodesInner) GetSecurityGroup() KubernetesNodeGroupsInnerNodesInnerSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetSecurityGroupOk() (*KubernetesNodeGroupsInnerNodesInnerSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *KubernetesNodeGroupsInnerNodesInner) SetSecurityGroup(v KubernetesNodeGroupsInnerNodesInnerSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *KubernetesNodeGroupsInnerNodesInner) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesNodeGroupsInnerNodesInner) GetCost() KubernetesNodeGroupsInnerNodesInnerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesNodeGroupsInnerNodesInner) GetCostOk() (*KubernetesNodeGroupsInnerNodesInnerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesNodeGroupsInnerNodesInner) SetCost(v KubernetesNodeGroupsInnerNodesInnerCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *KubernetesNodeGroupsInnerNodesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


