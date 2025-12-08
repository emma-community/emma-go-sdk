# VmActionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with a virtual machine | 
**DataCenterId** | **string** | ID of the provider&#39;s data center | 
**IsKeepOriginalInstance** | Pointer to **bool** | Indicate if it is necessary to keep the original instance | [optional] [default to true]
**Name** | **string** | A new name of the virtual machine | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**VCpuType** | **string** | Type of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 
**VolumeId** | **int32** | Volume ID to detach from the virtual machine | 

## Methods

### NewVmActionsRequest

`func NewVmActionsRequest(action string, dataCenterId string, name string, vCpu int32, vCpuType string, ramGb int32, volumeGb int32, volumeId int32, ) *VmActionsRequest`

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

`func (o *VmActionsRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmActionsRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmActionsRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetDataCenterId

`func (o *VmActionsRequest) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VmActionsRequest) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VmActionsRequest) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetIsKeepOriginalInstance

`func (o *VmActionsRequest) GetIsKeepOriginalInstance() bool`

GetIsKeepOriginalInstance returns the IsKeepOriginalInstance field if non-nil, zero value otherwise.

### GetIsKeepOriginalInstanceOk

`func (o *VmActionsRequest) GetIsKeepOriginalInstanceOk() (*bool, bool)`

GetIsKeepOriginalInstanceOk returns a tuple with the IsKeepOriginalInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepOriginalInstance

`func (o *VmActionsRequest) SetIsKeepOriginalInstance(v bool)`

SetIsKeepOriginalInstance sets IsKeepOriginalInstance field to given value.

### HasIsKeepOriginalInstance

`func (o *VmActionsRequest) HasIsKeepOriginalInstance() bool`

HasIsKeepOriginalInstance returns a boolean if a field has been set.

### GetName

`func (o *VmActionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmActionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmActionsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVCpu

`func (o *VmActionsRequest) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmActionsRequest) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmActionsRequest) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetVCpuType

`func (o *VmActionsRequest) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmActionsRequest) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmActionsRequest) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetRamGb

`func (o *VmActionsRequest) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmActionsRequest) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmActionsRequest) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeGb

`func (o *VmActionsRequest) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VmActionsRequest) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VmActionsRequest) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.


### GetVolumeId

`func (o *VmActionsRequest) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VmActionsRequest) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VmActionsRequest) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


