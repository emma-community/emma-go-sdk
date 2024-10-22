# KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VCpuType** | Pointer to **string** | Type of the virtual Central Processing Unit (vCPU) | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**VolumeGb** | Pointer to **int32** | Capacity of the volume in gigabytes | [optional] 
**VolumeType** | Pointer to **string** | Volume type | [optional] 
**Priority** | Pointer to **string** | Configuration priority | [optional] 

## Methods

### NewKubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner

`func NewKubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner() *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner`

NewKubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner instantiates a new KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInnerWithDefaults

`func NewKubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInnerWithDefaults() *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner`

NewKubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInnerWithDefaults instantiates a new KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVCpuType

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetVolumeGb

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.

### HasVolumeGb

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVolumeGb() bool`

HasVolumeGb returns a boolean if a field has been set.

### GetVolumeType

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetPriority

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *KubernetesCreateAutoscalingConfigsInnerConfigurationPrioritiesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


