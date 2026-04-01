# KubernetesCreateRequestWorkerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Worker node name | 
**ProviderId** | Pointer to **int32** | Provider ID | [optional] 
**LocationId** | Pointer to **int32** | Location ID | [optional] 
**DataCenterId** | **string** | Data center ID | 
**OsId** | Pointer to **int32** | Operating system ID | [optional] 
**OsType** | Pointer to **string** | Operating system type | [optional] 
**OsVersion** | Pointer to **string** | Operating system version | [optional] 
**CloudNetworkTypes** | Pointer to **[]string** | Cloud network types | [optional] 
**VCpuType** | **string** | Type of the virtual Central Processing Unit (vCPU) | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 
**VolumeType** | **string** | Volume type | 
**AcceleratorTypeId** | Pointer to **string** | GPU accelerator type identifier | [optional] 
**Accelerators** | Pointer to **int32** | Number of GPUs for the node configuration | [optional] 

## Methods

### NewKubernetesCreateRequestWorkerNodesInner

`func NewKubernetesCreateRequestWorkerNodesInner(name string, dataCenterId string, vCpuType string, vCpu int32, ramGb int32, volumeGb int32, volumeType string, ) *KubernetesCreateRequestWorkerNodesInner`

NewKubernetesCreateRequestWorkerNodesInner instantiates a new KubernetesCreateRequestWorkerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCreateRequestWorkerNodesInnerWithDefaults

`func NewKubernetesCreateRequestWorkerNodesInnerWithDefaults() *KubernetesCreateRequestWorkerNodesInner`

NewKubernetesCreateRequestWorkerNodesInnerWithDefaults instantiates a new KubernetesCreateRequestWorkerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesCreateRequestWorkerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCreateRequestWorkerNodesInner) SetName(v string)`

SetName sets Name field to given value.


### GetProviderId

`func (o *KubernetesCreateRequestWorkerNodesInner) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *KubernetesCreateRequestWorkerNodesInner) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *KubernetesCreateRequestWorkerNodesInner) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetLocationId

`func (o *KubernetesCreateRequestWorkerNodesInner) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *KubernetesCreateRequestWorkerNodesInner) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *KubernetesCreateRequestWorkerNodesInner) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetDataCenterId

`func (o *KubernetesCreateRequestWorkerNodesInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesCreateRequestWorkerNodesInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetOsId

`func (o *KubernetesCreateRequestWorkerNodesInner) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *KubernetesCreateRequestWorkerNodesInner) SetOsId(v int32)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *KubernetesCreateRequestWorkerNodesInner) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### GetOsType

`func (o *KubernetesCreateRequestWorkerNodesInner) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *KubernetesCreateRequestWorkerNodesInner) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *KubernetesCreateRequestWorkerNodesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetOsVersion

`func (o *KubernetesCreateRequestWorkerNodesInner) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *KubernetesCreateRequestWorkerNodesInner) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *KubernetesCreateRequestWorkerNodesInner) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetCloudNetworkTypes

`func (o *KubernetesCreateRequestWorkerNodesInner) GetCloudNetworkTypes() []string`

GetCloudNetworkTypes returns the CloudNetworkTypes field if non-nil, zero value otherwise.

### GetCloudNetworkTypesOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetCloudNetworkTypesOk() (*[]string, bool)`

GetCloudNetworkTypesOk returns a tuple with the CloudNetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkTypes

`func (o *KubernetesCreateRequestWorkerNodesInner) SetCloudNetworkTypes(v []string)`

SetCloudNetworkTypes sets CloudNetworkTypes field to given value.

### HasCloudNetworkTypes

`func (o *KubernetesCreateRequestWorkerNodesInner) HasCloudNetworkTypes() bool`

HasCloudNetworkTypes returns a boolean if a field has been set.

### GetVCpuType

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesCreateRequestWorkerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetVCpu

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesCreateRequestWorkerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetRamGb

`func (o *KubernetesCreateRequestWorkerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesCreateRequestWorkerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeGb

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *KubernetesCreateRequestWorkerNodesInner) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.


### GetVolumeType

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *KubernetesCreateRequestWorkerNodesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetAcceleratorTypeId

`func (o *KubernetesCreateRequestWorkerNodesInner) GetAcceleratorTypeId() string`

GetAcceleratorTypeId returns the AcceleratorTypeId field if non-nil, zero value otherwise.

### GetAcceleratorTypeIdOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetAcceleratorTypeIdOk() (*string, bool)`

GetAcceleratorTypeIdOk returns a tuple with the AcceleratorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorTypeId

`func (o *KubernetesCreateRequestWorkerNodesInner) SetAcceleratorTypeId(v string)`

SetAcceleratorTypeId sets AcceleratorTypeId field to given value.

### HasAcceleratorTypeId

`func (o *KubernetesCreateRequestWorkerNodesInner) HasAcceleratorTypeId() bool`

HasAcceleratorTypeId returns a boolean if a field has been set.

### GetAccelerators

`func (o *KubernetesCreateRequestWorkerNodesInner) GetAccelerators() int32`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *KubernetesCreateRequestWorkerNodesInner) GetAcceleratorsOk() (*int32, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *KubernetesCreateRequestWorkerNodesInner) SetAccelerators(v int32)`

SetAccelerators sets Accelerators field to given value.

### HasAccelerators

`func (o *KubernetesCreateRequestWorkerNodesInner) HasAccelerators() bool`

HasAccelerators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


