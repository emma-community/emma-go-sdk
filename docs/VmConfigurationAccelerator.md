# VmConfigurationAccelerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceleratorTypeId** | Pointer to **string** | Accelerator type ID | [optional] 
**AcceleratorType** | Pointer to **string** | GPU accelerator type | [optional] 
**Accelerators** | Pointer to **float32** | Quantity of GPU accelerators | [optional] 
**TotalGpuMemoryGb** | Pointer to **int32** | Total GPU memory (GB) | [optional] 

## Methods

### NewVmConfigurationAccelerator

`func NewVmConfigurationAccelerator() *VmConfigurationAccelerator`

NewVmConfigurationAccelerator instantiates a new VmConfigurationAccelerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmConfigurationAcceleratorWithDefaults

`func NewVmConfigurationAcceleratorWithDefaults() *VmConfigurationAccelerator`

NewVmConfigurationAcceleratorWithDefaults instantiates a new VmConfigurationAccelerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceleratorTypeId

`func (o *VmConfigurationAccelerator) GetAcceleratorTypeId() string`

GetAcceleratorTypeId returns the AcceleratorTypeId field if non-nil, zero value otherwise.

### GetAcceleratorTypeIdOk

`func (o *VmConfigurationAccelerator) GetAcceleratorTypeIdOk() (*string, bool)`

GetAcceleratorTypeIdOk returns a tuple with the AcceleratorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorTypeId

`func (o *VmConfigurationAccelerator) SetAcceleratorTypeId(v string)`

SetAcceleratorTypeId sets AcceleratorTypeId field to given value.

### HasAcceleratorTypeId

`func (o *VmConfigurationAccelerator) HasAcceleratorTypeId() bool`

HasAcceleratorTypeId returns a boolean if a field has been set.

### GetAcceleratorType

`func (o *VmConfigurationAccelerator) GetAcceleratorType() string`

GetAcceleratorType returns the AcceleratorType field if non-nil, zero value otherwise.

### GetAcceleratorTypeOk

`func (o *VmConfigurationAccelerator) GetAcceleratorTypeOk() (*string, bool)`

GetAcceleratorTypeOk returns a tuple with the AcceleratorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorType

`func (o *VmConfigurationAccelerator) SetAcceleratorType(v string)`

SetAcceleratorType sets AcceleratorType field to given value.

### HasAcceleratorType

`func (o *VmConfigurationAccelerator) HasAcceleratorType() bool`

HasAcceleratorType returns a boolean if a field has been set.

### GetAccelerators

`func (o *VmConfigurationAccelerator) GetAccelerators() float32`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *VmConfigurationAccelerator) GetAcceleratorsOk() (*float32, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *VmConfigurationAccelerator) SetAccelerators(v float32)`

SetAccelerators sets Accelerators field to given value.

### HasAccelerators

`func (o *VmConfigurationAccelerator) HasAccelerators() bool`

HasAccelerators returns a boolean if a field has been set.

### GetTotalGpuMemoryGb

`func (o *VmConfigurationAccelerator) GetTotalGpuMemoryGb() int32`

GetTotalGpuMemoryGb returns the TotalGpuMemoryGb field if non-nil, zero value otherwise.

### GetTotalGpuMemoryGbOk

`func (o *VmConfigurationAccelerator) GetTotalGpuMemoryGbOk() (*int32, bool)`

GetTotalGpuMemoryGbOk returns a tuple with the TotalGpuMemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGpuMemoryGb

`func (o *VmConfigurationAccelerator) SetTotalGpuMemoryGb(v int32)`

SetTotalGpuMemoryGb sets TotalGpuMemoryGb field to given value.

### HasTotalGpuMemoryGb

`func (o *VmConfigurationAccelerator) HasTotalGpuMemoryGb() bool`

HasTotalGpuMemoryGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


