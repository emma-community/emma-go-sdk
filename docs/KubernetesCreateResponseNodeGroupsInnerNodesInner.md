# KubernetesCreateResponseNodeGroupsInnerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the worker node | [optional] 
**CreatedAt** | Pointer to **string** | Date and time when the worker node was created | [optional] 
**Name** | Pointer to **string** | Name of the worker node | [optional] 
**Status** | Pointer to **string** | Status of the worker node | [optional] 
**Provider** | Pointer to [**KubernetesCreateResponseNodeGroupsInnerNodesInnerProvider**](KubernetesCreateResponseNodeGroupsInnerNodesInnerProvider.md) |  | [optional] 
**Location** | Pointer to [**KubernetesCreateResponseNodeGroupsInnerNodesInnerLocation**](KubernetesCreateResponseNodeGroupsInnerNodesInnerLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**KubernetesCreateResponseNodeGroupsInnerNodesInnerDataCenter**](KubernetesCreateResponseNodeGroupsInnerNodesInnerDataCenter.md) |  | [optional] 
**Os** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerOs**](KubernetesGetResponseNodeGroupsInnerNodesInnerOs.md) |  | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**VCpuType** | Pointer to **string** | Type of the virtual Central Processing Unit | [optional] 
**CloudNetworkType** | Pointer to **string** | Type of the cloud network (multi-cloud, isolated) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Disks** | Pointer to [**[]KubernetesCreateResponseNodeGroupsInnerNodesInnerDisksInner**](KubernetesCreateResponseNodeGroupsInnerNodesInnerDisksInner.md) | Volumes attached | [optional] 
**Networks** | Pointer to [**[]KubernetesCreateResponseNodeGroupsInnerNodesInnerNetworksInner**](KubernetesCreateResponseNodeGroupsInnerNodesInnerNetworksInner.md) | Public and private networks | [optional] 
**SecurityGroup** | Pointer to [**KubernetesCreateResponseNodeGroupsInnerNodesInnerSecurityGroup**](KubernetesCreateResponseNodeGroupsInnerNodesInnerSecurityGroup.md) |  | [optional] 
**LocalDisks** | Pointer to [**[]KubernetesCreateResponseNodeGroupsInnerNodesInnerLocalDisksInner**](KubernetesCreateResponseNodeGroupsInnerNodesInnerLocalDisksInner.md) | Local SSD storage directly attached to the host. Data is ephemeral and will be lost when the VM is stopped or deleted. | [optional] 
**Accelerator** | Pointer to [**KubernetesCreateResponseNodeGroupsInnerNodesInnerAccelerator**](KubernetesCreateResponseNodeGroupsInnerNodesInnerAccelerator.md) |  | [optional] 
**Cost** | Pointer to [**KubernetesCreateResponseNodeGroupsInnerNodesInnerCost**](KubernetesCreateResponseNodeGroupsInnerNodesInnerCost.md) |  | [optional] 

## Methods

### NewKubernetesCreateResponseNodeGroupsInnerNodesInner

`func NewKubernetesCreateResponseNodeGroupsInnerNodesInner() *KubernetesCreateResponseNodeGroupsInnerNodesInner`

NewKubernetesCreateResponseNodeGroupsInnerNodesInner instantiates a new KubernetesCreateResponseNodeGroupsInnerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCreateResponseNodeGroupsInnerNodesInnerWithDefaults

`func NewKubernetesCreateResponseNodeGroupsInnerNodesInnerWithDefaults() *KubernetesCreateResponseNodeGroupsInnerNodesInner`

NewKubernetesCreateResponseNodeGroupsInnerNodesInnerWithDefaults instantiates a new KubernetesCreateResponseNodeGroupsInnerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetProvider() KubernetesCreateResponseNodeGroupsInnerNodesInnerProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetProviderOk() (*KubernetesCreateResponseNodeGroupsInnerNodesInnerProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetProvider(v KubernetesCreateResponseNodeGroupsInnerNodesInnerProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetLocation() KubernetesCreateResponseNodeGroupsInnerNodesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetLocationOk() (*KubernetesCreateResponseNodeGroupsInnerNodesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetLocation(v KubernetesCreateResponseNodeGroupsInnerNodesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetDataCenter() KubernetesCreateResponseNodeGroupsInnerNodesInnerDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetDataCenterOk() (*KubernetesCreateResponseNodeGroupsInnerNodesInnerDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetDataCenter(v KubernetesCreateResponseNodeGroupsInnerNodesInnerDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetOs() KubernetesGetResponseNodeGroupsInnerNodesInnerOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetOsOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetOs(v KubernetesGetResponseNodeGroupsInnerNodesInnerOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetVCpu

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetVCpuType

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetRamGb

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetDisks() []KubernetesCreateResponseNodeGroupsInnerNodesInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetDisksOk() (*[]KubernetesCreateResponseNodeGroupsInnerNodesInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetDisks(v []KubernetesCreateResponseNodeGroupsInnerNodesInnerDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetNetworks() []KubernetesCreateResponseNodeGroupsInnerNodesInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetNetworksOk() (*[]KubernetesCreateResponseNodeGroupsInnerNodesInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetNetworks(v []KubernetesCreateResponseNodeGroupsInnerNodesInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetSecurityGroup() KubernetesCreateResponseNodeGroupsInnerNodesInnerSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetSecurityGroupOk() (*KubernetesCreateResponseNodeGroupsInnerNodesInnerSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetSecurityGroup(v KubernetesCreateResponseNodeGroupsInnerNodesInnerSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetLocalDisks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetLocalDisks() []KubernetesCreateResponseNodeGroupsInnerNodesInnerLocalDisksInner`

GetLocalDisks returns the LocalDisks field if non-nil, zero value otherwise.

### GetLocalDisksOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetLocalDisksOk() (*[]KubernetesCreateResponseNodeGroupsInnerNodesInnerLocalDisksInner, bool)`

GetLocalDisksOk returns a tuple with the LocalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDisks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetLocalDisks(v []KubernetesCreateResponseNodeGroupsInnerNodesInnerLocalDisksInner)`

SetLocalDisks sets LocalDisks field to given value.

### HasLocalDisks

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasLocalDisks() bool`

HasLocalDisks returns a boolean if a field has been set.

### GetAccelerator

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetAccelerator() KubernetesCreateResponseNodeGroupsInnerNodesInnerAccelerator`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetAcceleratorOk() (*KubernetesCreateResponseNodeGroupsInnerNodesInnerAccelerator, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetAccelerator(v KubernetesCreateResponseNodeGroupsInnerNodesInnerAccelerator)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetCost() KubernetesCreateResponseNodeGroupsInnerNodesInnerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) GetCostOk() (*KubernetesCreateResponseNodeGroupsInnerNodesInnerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) SetCost(v KubernetesCreateResponseNodeGroupsInnerNodesInnerCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *KubernetesCreateResponseNodeGroupsInnerNodesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


