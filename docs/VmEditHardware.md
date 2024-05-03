# VmEditHardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **interface{}** | Action with a virtual machine | 
**VCpu** | **interface{}** |  | 
**VCpuType** | Pointer to **interface{}** | vCPU type | [optional] 
**RamGb** | **interface{}** |  | 
**VolumeGb** | **interface{}** |  | 

## Methods

### NewVmEditHardware

`func NewVmEditHardware(action interface{}, vCpu interface{}, ramGb interface{}, volumeGb interface{}, ) *VmEditHardware`

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

`func (o *VmEditHardware) GetAction() interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmEditHardware) GetActionOk() (*interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmEditHardware) SetAction(v interface{})`

SetAction sets Action field to given value.


### SetActionNil

`func (o *VmEditHardware) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *VmEditHardware) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetVCpu

`func (o *VmEditHardware) GetVCpu() interface{}`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmEditHardware) GetVCpuOk() (*interface{}, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmEditHardware) SetVCpu(v interface{})`

SetVCpu sets VCpu field to given value.


### SetVCpuNil

`func (o *VmEditHardware) SetVCpuNil(b bool)`

 SetVCpuNil sets the value for VCpu to be an explicit nil

### UnsetVCpu
`func (o *VmEditHardware) UnsetVCpu()`

UnsetVCpu ensures that no value is present for VCpu, not even an explicit nil
### GetVCpuType

`func (o *VmEditHardware) GetVCpuType() interface{}`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmEditHardware) GetVCpuTypeOk() (*interface{}, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmEditHardware) SetVCpuType(v interface{})`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *VmEditHardware) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### SetVCpuTypeNil

`func (o *VmEditHardware) SetVCpuTypeNil(b bool)`

 SetVCpuTypeNil sets the value for VCpuType to be an explicit nil

### UnsetVCpuType
`func (o *VmEditHardware) UnsetVCpuType()`

UnsetVCpuType ensures that no value is present for VCpuType, not even an explicit nil
### GetRamGb

`func (o *VmEditHardware) GetRamGb() interface{}`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmEditHardware) GetRamGbOk() (*interface{}, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmEditHardware) SetRamGb(v interface{})`

SetRamGb sets RamGb field to given value.


### SetRamGbNil

`func (o *VmEditHardware) SetRamGbNil(b bool)`

 SetRamGbNil sets the value for RamGb to be an explicit nil

### UnsetRamGb
`func (o *VmEditHardware) UnsetRamGb()`

UnsetRamGb ensures that no value is present for RamGb, not even an explicit nil
### GetVolumeGb

`func (o *VmEditHardware) GetVolumeGb() interface{}`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VmEditHardware) GetVolumeGbOk() (*interface{}, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VmEditHardware) SetVolumeGb(v interface{})`

SetVolumeGb sets VolumeGb field to given value.


### SetVolumeGbNil

`func (o *VmEditHardware) SetVolumeGbNil(b bool)`

 SetVolumeGbNil sets the value for VolumeGb to be an explicit nil

### UnsetVolumeGb
`func (o *VmEditHardware) UnsetVolumeGb()`

UnsetVolumeGb ensures that no value is present for VolumeGb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


