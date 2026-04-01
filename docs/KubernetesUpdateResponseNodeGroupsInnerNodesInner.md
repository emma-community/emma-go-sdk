# KubernetesUpdateResponseNodeGroupsInnerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Node ID | 
**CreatedAt** | **time.Time** | Node creation date and time | 
**Name** | **string** | Node name | 
**Status** | **string** | Node status | 
**Provider** | [**KubernetesUpdateResponseNodeGroupsInnerNodesInnerProvider**](KubernetesUpdateResponseNodeGroupsInnerNodesInnerProvider.md) |  | 
**Location** | [**KubernetesUpdateResponseNodeGroupsInnerNodesInnerLocation**](KubernetesUpdateResponseNodeGroupsInnerNodesInnerLocation.md) |  | 
**DataCenter** | [**KubernetesUpdateResponseNodeGroupsInnerNodesInnerDataCenter**](KubernetesUpdateResponseNodeGroupsInnerNodesInnerDataCenter.md) |  | 
**Os** | [**KubernetesUpdateResponseNodeGroupsInnerNodesInnerOs**](KubernetesUpdateResponseNodeGroupsInnerNodesInnerOs.md) |  | 
**CloudNetworkType** | **string** | Cloud network type | 
**VCpuType** | **string** | Type of the virtual Central Processing Unit (vCPU) | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 
**Disks** | [**[]KubernetesUpdateResponseNodeGroupsInnerNodesInnerDisksInner**](KubernetesUpdateResponseNodeGroupsInnerNodesInnerDisksInner.md) | List of attached disks | 
**Networks** | Pointer to **[]map[string]interface{}** | List of attached networks | [optional] 
**Cost** | [**KubernetesUpdateResponseCost**](KubernetesUpdateResponseCost.md) |  | 
**SecurityGroup** | Pointer to **map[string]interface{}** | Security group attached to the node | [optional] 
**Accelerator** | Pointer to **map[string]interface{}** | Accelerator attached to the node | [optional] 
**LocalDisks** | Pointer to **[]map[string]interface{}** | Local disks attached to the node | [optional] 

## Methods

### NewKubernetesUpdateResponseNodeGroupsInnerNodesInner

`func NewKubernetesUpdateResponseNodeGroupsInnerNodesInner(id int32, createdAt time.Time, name string, status string, provider KubernetesUpdateResponseNodeGroupsInnerNodesInnerProvider, location KubernetesUpdateResponseNodeGroupsInnerNodesInnerLocation, dataCenter KubernetesUpdateResponseNodeGroupsInnerNodesInnerDataCenter, os KubernetesUpdateResponseNodeGroupsInnerNodesInnerOs, cloudNetworkType string, vCpuType string, vCpu int32, ramGb int32, disks []KubernetesUpdateResponseNodeGroupsInnerNodesInnerDisksInner, cost KubernetesUpdateResponseCost, ) *KubernetesUpdateResponseNodeGroupsInnerNodesInner`

NewKubernetesUpdateResponseNodeGroupsInnerNodesInner instantiates a new KubernetesUpdateResponseNodeGroupsInnerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateResponseNodeGroupsInnerNodesInnerWithDefaults

`func NewKubernetesUpdateResponseNodeGroupsInnerNodesInnerWithDefaults() *KubernetesUpdateResponseNodeGroupsInnerNodesInner`

NewKubernetesUpdateResponseNodeGroupsInnerNodesInnerWithDefaults instantiates a new KubernetesUpdateResponseNodeGroupsInnerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProvider

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetProvider() KubernetesUpdateResponseNodeGroupsInnerNodesInnerProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetProviderOk() (*KubernetesUpdateResponseNodeGroupsInnerNodesInnerProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetProvider(v KubernetesUpdateResponseNodeGroupsInnerNodesInnerProvider)`

SetProvider sets Provider field to given value.


### GetLocation

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetLocation() KubernetesUpdateResponseNodeGroupsInnerNodesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetLocationOk() (*KubernetesUpdateResponseNodeGroupsInnerNodesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetLocation(v KubernetesUpdateResponseNodeGroupsInnerNodesInnerLocation)`

SetLocation sets Location field to given value.


### GetDataCenter

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetDataCenter() KubernetesUpdateResponseNodeGroupsInnerNodesInnerDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetDataCenterOk() (*KubernetesUpdateResponseNodeGroupsInnerNodesInnerDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetDataCenter(v KubernetesUpdateResponseNodeGroupsInnerNodesInnerDataCenter)`

SetDataCenter sets DataCenter field to given value.


### GetOs

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetOs() KubernetesUpdateResponseNodeGroupsInnerNodesInnerOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetOsOk() (*KubernetesUpdateResponseNodeGroupsInnerNodesInnerOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetOs(v KubernetesUpdateResponseNodeGroupsInnerNodesInnerOs)`

SetOs sets Os field to given value.


### GetCloudNetworkType

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.


### GetVCpuType

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetVCpu

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetRamGb

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetDisks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetDisks() []KubernetesUpdateResponseNodeGroupsInnerNodesInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetDisksOk() (*[]KubernetesUpdateResponseNodeGroupsInnerNodesInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetDisks(v []KubernetesUpdateResponseNodeGroupsInnerNodesInnerDisksInner)`

SetDisks sets Disks field to given value.


### GetNetworks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetNetworks() []map[string]interface{}`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetNetworksOk() (*[]map[string]interface{}, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetNetworks(v []map[string]interface{})`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetCost() KubernetesUpdateResponseCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetCostOk() (*KubernetesUpdateResponseCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetCost(v KubernetesUpdateResponseCost)`

SetCost sets Cost field to given value.


### GetSecurityGroup

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetSecurityGroup() map[string]interface{}`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetSecurityGroupOk() (*map[string]interface{}, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetSecurityGroup(v map[string]interface{})`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetAccelerator

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetAccelerator() map[string]interface{}`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetAcceleratorOk() (*map[string]interface{}, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetAccelerator(v map[string]interface{})`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetLocalDisks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetLocalDisks() []map[string]interface{}`

GetLocalDisks returns the LocalDisks field if non-nil, zero value otherwise.

### GetLocalDisksOk

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) GetLocalDisksOk() (*[]map[string]interface{}, bool)`

GetLocalDisksOk returns a tuple with the LocalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDisks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) SetLocalDisks(v []map[string]interface{})`

SetLocalDisks sets LocalDisks field to given value.

### HasLocalDisks

`func (o *KubernetesUpdateResponseNodeGroupsInnerNodesInner) HasLocalDisks() bool`

HasLocalDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


