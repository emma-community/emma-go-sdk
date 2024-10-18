# VmEditHardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with a virtual machine | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**VCpuType** | Pointer to **string** | Type of virtual Central Processing Units (vCPUs) | [optional] 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 

## Methods

### NewVmEditHardware

`func NewVmEditHardware(action string, vCpu int32, ramGb int32, volumeGb int32, ) *VmEditHardware`

NewVmEditHardware instantiates a new VmEditHardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmEditHardwareWithDefaults

`func NewVmEditHardwareWithDefaults() *VmEditHardware`

NewVmEditHardwareWithDefaults instantiates a new VmEditHardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VmEditHardware) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmEditHardware) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmEditHardware) SetAction(v string)`

SetAction sets Action field to given value.


### GetVCpu

`func (o *VmEditHardware) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmEditHardware) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmEditHardware) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetVCpuType

`func (o *VmEditHardware) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmEditHardware) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmEditHardware) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *VmEditHardware) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetRamGb

`func (o *VmEditHardware) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmEditHardware) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmEditHardware) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeGb

`func (o *VmEditHardware) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VmEditHardware) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VmEditHardware) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


