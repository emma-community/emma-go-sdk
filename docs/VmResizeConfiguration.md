# VmResizeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VCpuType** | Pointer to **string** | Type of virtual Central Processing Units (vCPUs) | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**Cost** | Pointer to [**VmConfigurationCost**](VmConfigurationCost.md) |  | [optional] 

## Methods

### NewVmResizeConfiguration

`func NewVmResizeConfiguration() *VmResizeConfiguration`

NewVmResizeConfiguration instantiates a new VmResizeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmResizeConfigurationWithDefaults

`func NewVmResizeConfigurationWithDefaults() *VmResizeConfiguration`

NewVmResizeConfigurationWithDefaults instantiates a new VmResizeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVCpuType

`func (o *VmResizeConfiguration) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmResizeConfiguration) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmResizeConfiguration) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *VmResizeConfiguration) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *VmResizeConfiguration) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmResizeConfiguration) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmResizeConfiguration) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *VmResizeConfiguration) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *VmResizeConfiguration) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmResizeConfiguration) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmResizeConfiguration) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *VmResizeConfiguration) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetCost

`func (o *VmResizeConfiguration) GetCost() VmConfigurationCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *VmResizeConfiguration) GetCostOk() (*VmConfigurationCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *VmResizeConfiguration) SetCost(v VmConfigurationCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *VmResizeConfiguration) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


