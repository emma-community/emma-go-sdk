# KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner

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

### NewKubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner

`func NewKubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner() *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner`

NewKubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner instantiates a new KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInnerWithDefaults

`func NewKubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInnerWithDefaults() *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner`

NewKubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInnerWithDefaults instantiates a new KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVCpuType

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetVolumeGb

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.

### HasVolumeGb

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVolumeGb() bool`

HasVolumeGb returns a boolean if a field has been set.

### GetVolumeType

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetPriority

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


