# KubernetesUpdateRequestWorkerNodesInner

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
**AcceleratorTypeId** | Pointer to **string** | GPU accelerator type identifier (from the UI accelerator dictionary). | [optional] 
**Accelerators** | Pointer to **float64** | Number of GPUs for the node configuration. | [optional] 

## Methods

### NewKubernetesUpdateRequestWorkerNodesInner

`func NewKubernetesUpdateRequestWorkerNodesInner(name string, dataCenterId string, vCpuType string, vCpu int32, ramGb int32, volumeType string, volumeGb int32, ) *KubernetesUpdateRequestWorkerNodesInner`

NewKubernetesUpdateRequestWorkerNodesInner instantiates a new KubernetesUpdateRequestWorkerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateRequestWorkerNodesInnerWithDefaults

`func NewKubernetesUpdateRequestWorkerNodesInnerWithDefaults() *KubernetesUpdateRequestWorkerNodesInner`

NewKubernetesUpdateRequestWorkerNodesInnerWithDefaults instantiates a new KubernetesUpdateRequestWorkerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesUpdateRequestWorkerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDataCenterId

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetVCpuType

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetVCpu

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetRamGb

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeType

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetVolumeGb

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.


### GetAcceleratorTypeId

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetAcceleratorTypeId() string`

GetAcceleratorTypeId returns the AcceleratorTypeId field if non-nil, zero value otherwise.

### GetAcceleratorTypeIdOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetAcceleratorTypeIdOk() (*string, bool)`

GetAcceleratorTypeIdOk returns a tuple with the AcceleratorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorTypeId

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetAcceleratorTypeId(v string)`

SetAcceleratorTypeId sets AcceleratorTypeId field to given value.

### HasAcceleratorTypeId

`func (o *KubernetesUpdateRequestWorkerNodesInner) HasAcceleratorTypeId() bool`

HasAcceleratorTypeId returns a boolean if a field has been set.

### GetAccelerators

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetAccelerators() float64`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *KubernetesUpdateRequestWorkerNodesInner) GetAcceleratorsOk() (*float64, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *KubernetesUpdateRequestWorkerNodesInner) SetAccelerators(v float64)`

SetAccelerators sets Accelerators field to given value.

### HasAccelerators

`func (o *KubernetesUpdateRequestWorkerNodesInner) HasAccelerators() bool`

HasAccelerators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


