# VmResizeCompute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with a virtual machine.  **Note:** To retrieve available configurations for DigitalOcean system volumes, use the &#x60;edithardware&#x60; action  | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**VCpuType** | **string** | Type of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 

## Methods

### NewVmResizeCompute

`func NewVmResizeCompute(action string, vCpu int32, vCpuType string, ramGb int32, ) *VmResizeCompute`

NewVmResizeCompute instantiates a new VmResizeCompute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmResizeComputeWithDefaults

`func NewVmResizeComputeWithDefaults() *VmResizeCompute`

NewVmResizeComputeWithDefaults instantiates a new VmResizeCompute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VmResizeCompute) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmResizeCompute) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmResizeCompute) SetAction(v string)`

SetAction sets Action field to given value.


### GetVCpu

`func (o *VmResizeCompute) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmResizeCompute) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmResizeCompute) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetVCpuType

`func (o *VmResizeCompute) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmResizeCompute) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmResizeCompute) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetRamGb

`func (o *VmResizeCompute) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmResizeCompute) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmResizeCompute) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


