# KubernetesDeleteNodeGroupsInnerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Node ID | [optional] 
**CreatedAt** | Pointer to **time.Time** | Node creation date and time | [optional] 
**Name** | Pointer to **string** | Node name | [optional] 
**Status** | Pointer to **string** | Node status | [optional] 
**Provider** | Pointer to [**KubernetesDeleteNodeGroupsInnerNodesInnerProvider**](KubernetesDeleteNodeGroupsInnerNodesInnerProvider.md) |  | [optional] 
**Location** | Pointer to [**KubernetesDeleteNodeGroupsInnerNodesInnerLocation**](KubernetesDeleteNodeGroupsInnerNodesInnerLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**KubernetesDeleteNodeGroupsInnerNodesInnerDataCenter**](KubernetesDeleteNodeGroupsInnerNodesInnerDataCenter.md) |  | [optional] 
**Os** | Pointer to [**KubernetesDeleteNodeGroupsInnerNodesInnerOs**](KubernetesDeleteNodeGroupsInnerNodesInnerOs.md) |  | [optional] 
**CloudNetworkType** | Pointer to **string** | Cloud network type | [optional] 
**VCpuType** | Pointer to **string** | Type of the virtual Central Processing Unit (vCPU) | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Disks** | Pointer to [**[]KubernetesDeleteNodeGroupsInnerNodesInnerDisksInner**](KubernetesDeleteNodeGroupsInnerNodesInnerDisksInner.md) | List of attached disks | [optional] 
**Networks** | Pointer to [**[]KubernetesDeleteNodeGroupsInnerNodesInnerNetworksInner**](KubernetesDeleteNodeGroupsInnerNodesInnerNetworksInner.md) | List of attached networks | [optional] 
**Cost** | Pointer to [**KubernetesDeleteCost**](KubernetesDeleteCost.md) |  | [optional] 
**SecurityGroup** | Pointer to [**KubernetesDeleteNodeGroupsInnerNodesInnerSecurityGroup**](KubernetesDeleteNodeGroupsInnerNodesInnerSecurityGroup.md) |  | [optional] 
**Accelerator** | Pointer to **map[string]interface{}** | Accelerator attached to the node | [optional] 
**LocalDisks** | Pointer to **[]map[string]interface{}** | Local disks attached to the node | [optional] 

## Methods

### NewKubernetesDeleteNodeGroupsInnerNodesInner

`func NewKubernetesDeleteNodeGroupsInnerNodesInner() *KubernetesDeleteNodeGroupsInnerNodesInner`

NewKubernetesDeleteNodeGroupsInnerNodesInner instantiates a new KubernetesDeleteNodeGroupsInnerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDeleteNodeGroupsInnerNodesInnerWithDefaults

`func NewKubernetesDeleteNodeGroupsInnerNodesInnerWithDefaults() *KubernetesDeleteNodeGroupsInnerNodesInner`

NewKubernetesDeleteNodeGroupsInnerNodesInnerWithDefaults instantiates a new KubernetesDeleteNodeGroupsInnerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetProvider() KubernetesDeleteNodeGroupsInnerNodesInnerProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetProviderOk() (*KubernetesDeleteNodeGroupsInnerNodesInnerProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetProvider(v KubernetesDeleteNodeGroupsInnerNodesInnerProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetLocation() KubernetesDeleteNodeGroupsInnerNodesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetLocationOk() (*KubernetesDeleteNodeGroupsInnerNodesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetLocation(v KubernetesDeleteNodeGroupsInnerNodesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetDataCenter() KubernetesDeleteNodeGroupsInnerNodesInnerDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetDataCenterOk() (*KubernetesDeleteNodeGroupsInnerNodesInnerDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetDataCenter(v KubernetesDeleteNodeGroupsInnerNodesInnerDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetOs

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetOs() KubernetesDeleteNodeGroupsInnerNodesInnerOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetOsOk() (*KubernetesDeleteNodeGroupsInnerNodesInnerOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetOs(v KubernetesDeleteNodeGroupsInnerNodesInnerOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCloudNetworkType

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.

### HasCloudNetworkType

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasCloudNetworkType() bool`

HasCloudNetworkType returns a boolean if a field has been set.

### GetVCpuType

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetDisks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetDisks() []KubernetesDeleteNodeGroupsInnerNodesInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetDisksOk() (*[]KubernetesDeleteNodeGroupsInnerNodesInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetDisks(v []KubernetesDeleteNodeGroupsInnerNodesInnerDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetworks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetNetworks() []KubernetesDeleteNodeGroupsInnerNodesInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetNetworksOk() (*[]KubernetesDeleteNodeGroupsInnerNodesInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetNetworks(v []KubernetesDeleteNodeGroupsInnerNodesInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetCost() KubernetesDeleteCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetCostOk() (*KubernetesDeleteCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetCost(v KubernetesDeleteCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetSecurityGroup() KubernetesDeleteNodeGroupsInnerNodesInnerSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetSecurityGroupOk() (*KubernetesDeleteNodeGroupsInnerNodesInnerSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetSecurityGroup(v KubernetesDeleteNodeGroupsInnerNodesInnerSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetAccelerator

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetAccelerator() map[string]interface{}`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetAcceleratorOk() (*map[string]interface{}, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetAccelerator(v map[string]interface{})`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetLocalDisks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetLocalDisks() []map[string]interface{}`

GetLocalDisks returns the LocalDisks field if non-nil, zero value otherwise.

### GetLocalDisksOk

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) GetLocalDisksOk() (*[]map[string]interface{}, bool)`

GetLocalDisksOk returns a tuple with the LocalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDisks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) SetLocalDisks(v []map[string]interface{})`

SetLocalDisks sets LocalDisks field to given value.

### HasLocalDisks

`func (o *KubernetesDeleteNodeGroupsInnerNodesInner) HasLocalDisks() bool`

HasLocalDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


