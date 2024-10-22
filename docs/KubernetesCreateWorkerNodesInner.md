# KubernetesCreateWorkerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the worker node | 
**DataCenterId** | **string** | ID of the data center | 
**VCpuType** | **string** | Type of the virtual Central Processing Unit (vCPU) | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 
**VolumeType** | **string** | Volume type | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 

## Methods

### NewKubernetesCreateWorkerNodesInner

`func NewKubernetesCreateWorkerNodesInner(name string, dataCenterId string, vCpuType string, vCpu int32, ramGb int32, volumeType string, volumeGb int32, ) *KubernetesCreateWorkerNodesInner`

NewKubernetesCreateWorkerNodesInner instantiates a new KubernetesCreateWorkerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCreateWorkerNodesInnerWithDefaults

`func NewKubernetesCreateWorkerNodesInnerWithDefaults() *KubernetesCreateWorkerNodesInner`

NewKubernetesCreateWorkerNodesInnerWithDefaults instantiates a new KubernetesCreateWorkerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesCreateWorkerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCreateWorkerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCreateWorkerNodesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDataCenterId

`func (o *KubernetesCreateWorkerNodesInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesCreateWorkerNodesInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesCreateWorkerNodesInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetVCpuType

`func (o *KubernetesCreateWorkerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesCreateWorkerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesCreateWorkerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetVCpu

`func (o *KubernetesCreateWorkerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesCreateWorkerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesCreateWorkerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetRamGb

`func (o *KubernetesCreateWorkerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesCreateWorkerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesCreateWorkerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeType

`func (o *KubernetesCreateWorkerNodesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *KubernetesCreateWorkerNodesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *KubernetesCreateWorkerNodesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetVolumeGb

`func (o *KubernetesCreateWorkerNodesInner) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *KubernetesCreateWorkerNodesInner) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *KubernetesCreateWorkerNodesInner) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


