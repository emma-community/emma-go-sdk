# KubernetesUpdateWorkerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the kubernetes worker node | [optional] 
**Name** | **string** | Name of the worker node | 
**DataCenterId** | **string** | ID of the data center | 
**VCpuType** | **string** | Type of the virtual Central Processing Unit (vCPU) | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 
**VolumeType** | **string** | Volume type | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 

## Methods

### NewKubernetesUpdateWorkerNodesInner

`func NewKubernetesUpdateWorkerNodesInner(name string, dataCenterId string, vCpuType string, vCpu int32, ramGb int32, volumeType string, volumeGb int32, ) *KubernetesUpdateWorkerNodesInner`

NewKubernetesUpdateWorkerNodesInner instantiates a new KubernetesUpdateWorkerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateWorkerNodesInnerWithDefaults

`func NewKubernetesUpdateWorkerNodesInnerWithDefaults() *KubernetesUpdateWorkerNodesInner`

NewKubernetesUpdateWorkerNodesInnerWithDefaults instantiates a new KubernetesUpdateWorkerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesUpdateWorkerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesUpdateWorkerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesUpdateWorkerNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesUpdateWorkerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KubernetesUpdateWorkerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesUpdateWorkerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesUpdateWorkerNodesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDataCenterId

`func (o *KubernetesUpdateWorkerNodesInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesUpdateWorkerNodesInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesUpdateWorkerNodesInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetVCpuType

`func (o *KubernetesUpdateWorkerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesUpdateWorkerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesUpdateWorkerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetVCpu

`func (o *KubernetesUpdateWorkerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesUpdateWorkerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesUpdateWorkerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetRamGb

`func (o *KubernetesUpdateWorkerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesUpdateWorkerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesUpdateWorkerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeType

`func (o *KubernetesUpdateWorkerNodesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *KubernetesUpdateWorkerNodesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *KubernetesUpdateWorkerNodesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetVolumeGb

`func (o *KubernetesUpdateWorkerNodesInner) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *KubernetesUpdateWorkerNodesInner) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *KubernetesUpdateWorkerNodesInner) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


