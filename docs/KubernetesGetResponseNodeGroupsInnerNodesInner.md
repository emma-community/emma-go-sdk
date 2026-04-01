# KubernetesGetResponseNodeGroupsInnerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the worker node | [optional] 
**CreatedAt** | Pointer to **string** | Date and time when the worker node was created | [optional] 
**Name** | Pointer to **string** | Name of the worker node | [optional] 
**Status** | Pointer to **string** | Status of the worker node | [optional] 
**Provider** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerProvider**](KubernetesGetResponseNodeGroupsInnerNodesInnerProvider.md) |  | [optional] 
**Location** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerLocation**](KubernetesGetResponseNodeGroupsInnerNodesInnerLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerDataCenter**](KubernetesGetResponseNodeGroupsInnerNodesInnerDataCenter.md) |  | [optional] 
**Os** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerOs**](KubernetesGetResponseNodeGroupsInnerNodesInnerOs.md) |  | [optional] 
**CloudNetworkType** | Pointer to **string** | Type of the cloud network | [optional] 
**VCpuType** | Pointer to **string** | Type of the virtual Central Processing Unit | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Disks** | Pointer to [**[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner.md) | Volumes attached | [optional] 
**Networks** | Pointer to [**[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner.md) | Public and private networks | [optional] 
**Cost** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerCost**](KubernetesGetResponseNodeGroupsInnerNodesInnerCost.md) |  | [optional] 
**SecurityGroup** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerSecurityGroup**](KubernetesGetResponseNodeGroupsInnerNodesInnerSecurityGroup.md) |  | [optional] 
**Accelerator** | Pointer to [**KubernetesGetResponseNodeGroupsInnerNodesInnerAccelerator**](KubernetesGetResponseNodeGroupsInnerNodesInnerAccelerator.md) |  | [optional] 
**LocalDisks** | Pointer to [**[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner**](KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner.md) | Local SSD storage directly attached to the host | [optional] 

## Methods

### NewKubernetesGetResponseNodeGroupsInnerNodesInner

`func NewKubernetesGetResponseNodeGroupsInnerNodesInner() *KubernetesGetResponseNodeGroupsInnerNodesInner`

NewKubernetesGetResponseNodeGroupsInnerNodesInner instantiates a new KubernetesGetResponseNodeGroupsInnerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesGetResponseNodeGroupsInnerNodesInnerWithDefaults

`func NewKubernetesGetResponseNodeGroupsInnerNodesInnerWithDefaults() *KubernetesGetResponseNodeGroupsInnerNodesInner`

NewKubernetesGetResponseNodeGroupsInnerNodesInnerWithDefaults instantiates a new KubernetesGetResponseNodeGroupsInnerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetProvider() KubernetesGetResponseNodeGroupsInnerNodesInnerProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetProviderOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetProvider(v KubernetesGetResponseNodeGroupsInnerNodesInnerProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetLocation() KubernetesGetResponseNodeGroupsInnerNodesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetLocationOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetLocation(v KubernetesGetResponseNodeGroupsInnerNodesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetDataCenter() KubernetesGetResponseNodeGroupsInnerNodesInnerDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetDataCenterOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetDataCenter(v KubernetesGetResponseNodeGroupsInnerNodesInnerDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetOs() KubernetesGetResponseNodeGroupsInnerNodesInnerOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetOsOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetOs(v KubernetesGetResponseNodeGroupsInnerNodesInnerOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetVCpuType

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetDisks() []KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetDisksOk() (*[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetDisks(v []KubernetesListResponseInnerNodeGroupsInnerNodesInnerDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetNetworks() []KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetNetworksOk() (*[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetNetworks(v []KubernetesListResponseInnerNodeGroupsInnerNodesInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetCost() KubernetesGetResponseNodeGroupsInnerNodesInnerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetCostOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetCost(v KubernetesGetResponseNodeGroupsInnerNodesInnerCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetSecurityGroup() KubernetesGetResponseNodeGroupsInnerNodesInnerSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetSecurityGroupOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetSecurityGroup(v KubernetesGetResponseNodeGroupsInnerNodesInnerSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetAccelerator

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetAccelerator() KubernetesGetResponseNodeGroupsInnerNodesInnerAccelerator`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetAcceleratorOk() (*KubernetesGetResponseNodeGroupsInnerNodesInnerAccelerator, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetAccelerator(v KubernetesGetResponseNodeGroupsInnerNodesInnerAccelerator)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetLocalDisks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetLocalDisks() []KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner`

GetLocalDisks returns the LocalDisks field if non-nil, zero value otherwise.

### GetLocalDisksOk

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) GetLocalDisksOk() (*[]KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner, bool)`

GetLocalDisksOk returns a tuple with the LocalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDisks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) SetLocalDisks(v []KubernetesListResponseInnerNodeGroupsInnerNodesInnerLocalDisksInner)`

SetLocalDisks sets LocalDisks field to given value.

### HasLocalDisks

`func (o *KubernetesGetResponseNodeGroupsInnerNodesInner) HasLocalDisks() bool`

HasLocalDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


