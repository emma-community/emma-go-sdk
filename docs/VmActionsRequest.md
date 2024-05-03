# VmActionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **interface{}** | Action with a virtual machine | 
**DataCenterId** | **interface{}** | Provider&#39;s data center ID | 
**IsKeepOriginalInstance** | Pointer to **interface{}** | Keep original instance | [optional] [default to true]
**Name** | **interface{}** | Virtual machine name | 
**VCpu** | **interface{}** |  | 
**VCpuType** | Pointer to **interface{}** | vCPU type | [optional] 
**RamGb** | **interface{}** |  | 
**VolumeGb** | **interface{}** |  | 

## Methods

### NewVmActionsRequest

`func NewVmActionsRequest(action interface{}, dataCenterId interface{}, name interface{}, vCpu interface{}, ramGb interface{}, volumeGb interface{}, ) *VmActionsRequest`

NewVmActionsRequest instantiates a new VmActionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmActionsRequestWithDefaults

`func NewVmActionsRequestWithDefaults() *VmActionsRequest`

NewVmActionsRequestWithDefaults instantiates a new VmActionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VmActionsRequest) GetAction() interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmActionsRequest) GetActionOk() (*interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmActionsRequest) SetAction(v interface{})`

SetAction sets Action field to given value.


### SetActionNil

`func (o *VmActionsRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *VmActionsRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDataCenterId

`func (o *VmActionsRequest) GetDataCenterId() interface{}`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VmActionsRequest) GetDataCenterIdOk() (*interface{}, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VmActionsRequest) SetDataCenterId(v interface{})`

SetDataCenterId sets DataCenterId field to given value.


### SetDataCenterIdNil

`func (o *VmActionsRequest) SetDataCenterIdNil(b bool)`

 SetDataCenterIdNil sets the value for DataCenterId to be an explicit nil

### UnsetDataCenterId
`func (o *VmActionsRequest) UnsetDataCenterId()`

UnsetDataCenterId ensures that no value is present for DataCenterId, not even an explicit nil
### GetIsKeepOriginalInstance

`func (o *VmActionsRequest) GetIsKeepOriginalInstance() interface{}`

GetIsKeepOriginalInstance returns the IsKeepOriginalInstance field if non-nil, zero value otherwise.

### GetIsKeepOriginalInstanceOk

`func (o *VmActionsRequest) GetIsKeepOriginalInstanceOk() (*interface{}, bool)`

GetIsKeepOriginalInstanceOk returns a tuple with the IsKeepOriginalInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepOriginalInstance

`func (o *VmActionsRequest) SetIsKeepOriginalInstance(v interface{})`

SetIsKeepOriginalInstance sets IsKeepOriginalInstance field to given value.

### HasIsKeepOriginalInstance

`func (o *VmActionsRequest) HasIsKeepOriginalInstance() bool`

HasIsKeepOriginalInstance returns a boolean if a field has been set.

### SetIsKeepOriginalInstanceNil

`func (o *VmActionsRequest) SetIsKeepOriginalInstanceNil(b bool)`

 SetIsKeepOriginalInstanceNil sets the value for IsKeepOriginalInstance to be an explicit nil

### UnsetIsKeepOriginalInstance
`func (o *VmActionsRequest) UnsetIsKeepOriginalInstance()`

UnsetIsKeepOriginalInstance ensures that no value is present for IsKeepOriginalInstance, not even an explicit nil
### GetName

`func (o *VmActionsRequest) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmActionsRequest) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmActionsRequest) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *VmActionsRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmActionsRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVCpu

`func (o *VmActionsRequest) GetVCpu() interface{}`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmActionsRequest) GetVCpuOk() (*interface{}, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmActionsRequest) SetVCpu(v interface{})`

SetVCpu sets VCpu field to given value.


### SetVCpuNil

`func (o *VmActionsRequest) SetVCpuNil(b bool)`

 SetVCpuNil sets the value for VCpu to be an explicit nil

### UnsetVCpu
`func (o *VmActionsRequest) UnsetVCpu()`

UnsetVCpu ensures that no value is present for VCpu, not even an explicit nil
### GetVCpuType

`func (o *VmActionsRequest) GetVCpuType() interface{}`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmActionsRequest) GetVCpuTypeOk() (*interface{}, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmActionsRequest) SetVCpuType(v interface{})`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *VmActionsRequest) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### SetVCpuTypeNil

`func (o *VmActionsRequest) SetVCpuTypeNil(b bool)`

 SetVCpuTypeNil sets the value for VCpuType to be an explicit nil

### UnsetVCpuType
`func (o *VmActionsRequest) UnsetVCpuType()`

UnsetVCpuType ensures that no value is present for VCpuType, not even an explicit nil
### GetRamGb

`func (o *VmActionsRequest) GetRamGb() interface{}`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmActionsRequest) GetRamGbOk() (*interface{}, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmActionsRequest) SetRamGb(v interface{})`

SetRamGb sets RamGb field to given value.


### SetRamGbNil

`func (o *VmActionsRequest) SetRamGbNil(b bool)`

 SetRamGbNil sets the value for RamGb to be an explicit nil

### UnsetRamGb
`func (o *VmActionsRequest) UnsetRamGb()`

UnsetRamGb ensures that no value is present for RamGb, not even an explicit nil
### GetVolumeGb

`func (o *VmActionsRequest) GetVolumeGb() interface{}`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VmActionsRequest) GetVolumeGbOk() (*interface{}, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VmActionsRequest) SetVolumeGb(v interface{})`

SetVolumeGb sets VolumeGb field to given value.


### SetVolumeGbNil

`func (o *VmActionsRequest) SetVolumeGbNil(b bool)`

 SetVolumeGbNil sets the value for VolumeGb to be an explicit nil

### UnsetVolumeGb
`func (o *VmActionsRequest) UnsetVolumeGb()`

UnsetVolumeGb ensures that no value is present for VolumeGb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


