# KubernetesListResponseInnerNodeGroupsInnerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the worker node | [optional] 
**CreatedAt** | Pointer to **string** | Date and time when the worker node was created | [optional] 
**Name** | Pointer to **string** | Name of the worker node | [optional] 
**Status** | Pointer to **string** | Status of the worker node | [optional] 
**Provider** | Pointer to [**KubernetesListResponseInnerNodeGroupsInnerNodesInnerProvider**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerProvider.md) |  | [optional] 
**Location** | Pointer to [**KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocation**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**KubernetesListResponseInnerNodeGroupsInnerNodesInnerDataCenter**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerDataCenter.md) |  | [optional] 
**Os** | Pointer to [**KubernetesListResponseInnerNodeGroupsInnerNodesInnerOs**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerOs.md) |  | [optional] 
**CloudNetworkType** | Pointer to **string** | Type of the cloud network | [optional] 
**VCpuType** | Pointer to **string** | Type of the virtual Central Processing Unit | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Disks** | Pointer to [**[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner.md) | Volumes attached | [optional] 
**Networks** | Pointer to [**[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner.md) | Public and private networks | [optional] 
**Cost** | Pointer to [**KubernetesListResponseInnerNodeGroupsInnerNodesInnerCost**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerCost.md) |  | [optional] 
**SecurityGroup** | Pointer to [**KubernetesListResponseInnerNodeGroupsInnerNodesInnerSecurityGroup**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerSecurityGroup.md) |  | [optional] 
**Accelerator** | Pointer to [**KubernetesListResponseInnerNodeGroupsInnerNodesInnerAccelerator**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerAccelerator.md) |  | [optional] 
**LocalDisks** | Pointer to [**[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner.md) | Local SSD storage directly attached to the host | [optional] 

## Methods

### NewKubernetesListResponseInnerNodeGroupsInnerNodesInner

`func NewKubernetesListResponseInnerNodeGroupsInnerNodesInner() *KubernetesListResponseInnerNodeGroupsInnerNodesInner`

NewKubernetesListResponseInnerNodeGroupsInnerNodesInner instantiates a new KubernetesListResponseInnerNodeGroupsInnerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesListResponseInnerNodeGroupsInnerNodesInnerWithDefaults

`func NewKubernetesListResponseInnerNodeGroupsInnerNodesInnerWithDefaults() *KubernetesListResponseInnerNodeGroupsInnerNodesInner`

NewKubernetesListResponseInnerNodeGroupsInnerNodesInnerWithDefaults instantiates a new KubernetesListResponseInnerNodeGroupsInnerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetProvider() KubernetesListResponseInnerNodeGroupsInnerNodesInnerProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetProviderOk() (*KubernetesListResponseInnerNodeGroupsInnerNodesInnerProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetProvider(v KubernetesListResponseInnerNodeGroupsInnerNodesInnerProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetLocation() KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetLocationOk() (*KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetLocation(v KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetDataCenter() KubernetesListResponseInnerNodeGroupsInnerNodesInnerDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetDataCenterOk() (*KubernetesListResponseInnerNodeGroupsInnerNodesInnerDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetDataCenter(v KubernetesListResponseInnerNodeGroupsInnerNodesInnerDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetOs() KubernetesListResponseInnerNodeGroupsInnerNodesInnerOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetOsOk() (*KubernetesListResponseInnerNodeGroupsInnerNodesInnerOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetOs(v KubernetesListResponseInnerNodeGroupsInnerNodesInnerOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetVCpuType

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetDisks() []KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetDisksOk() (*[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetDisks(v []KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetNetworks() []KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetNetworksOk() (*[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetNetworks(v []KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetCost() KubernetesListResponseInnerNodeGroupsInnerNodesInnerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetCostOk() (*KubernetesListResponseInnerNodeGroupsInnerNodesInnerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetCost(v KubernetesListResponseInnerNodeGroupsInnerNodesInnerCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetSecurityGroup() KubernetesListResponseInnerNodeGroupsInnerNodesInnerSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetSecurityGroupOk() (*KubernetesListResponseInnerNodeGroupsInnerNodesInnerSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetSecurityGroup(v KubernetesListResponseInnerNodeGroupsInnerNodesInnerSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetAccelerator

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetAccelerator() KubernetesListResponseInnerNodeGroupsInnerNodesInnerAccelerator`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetAcceleratorOk() (*KubernetesListResponseInnerNodeGroupsInnerNodesInnerAccelerator, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetAccelerator(v KubernetesListResponseInnerNodeGroupsInnerNodesInnerAccelerator)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetLocalDisks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetLocalDisks() []KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner`

GetLocalDisks returns the LocalDisks field if non-nil, zero value otherwise.

### GetLocalDisksOk

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) GetLocalDisksOk() (*[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner, bool)`

GetLocalDisksOk returns a tuple with the LocalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDisks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) SetLocalDisks(v []KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner)`

SetLocalDisks sets LocalDisks field to given value.

### HasLocalDisks

`func (o *KubernetesListResponseInnerNodeGroupsInnerNodesInner) HasLocalDisks() bool`

HasLocalDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


