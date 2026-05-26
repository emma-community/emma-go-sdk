# SpotVmAccelerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceleratorTypeId** | Pointer to **string** | GPU accelerator type ID | [optional] 
**AcceleratorType** | Pointer to **string** | GPU accelerator type | [optional] 
**Accelerators** | Pointer to **float32** | Quantity of GPU accelerators | [optional] 
**TotalGpuMemoryGb** | Pointer to **int32** | Total GPU memory | [optional] 

## Methods

### NewSpotVmAccelerator

`func NewSpotVmAccelerator() *SpotVmAccelerator`

NewSpotVmAccelerator instantiates a new SpotVmAccelerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotVmAcceleratorWithDefaults

`func NewSpotVmAcceleratorWithDefaults() *SpotVmAccelerator`

NewSpotVmAcceleratorWithDefaults instantiates a new SpotVmAccelerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceleratorTypeId

`func (o *SpotVmAccelerator) GetAcceleratorTypeId() string`

GetAcceleratorTypeId returns the AcceleratorTypeId field if non-nil, zero value otherwise.

### GetAcceleratorTypeIdOk

`func (o *SpotVmAccelerator) GetAcceleratorTypeIdOk() (*string, bool)`

GetAcceleratorTypeIdOk returns a tuple with the AcceleratorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorTypeId

`func (o *SpotVmAccelerator) SetAcceleratorTypeId(v string)`

SetAcceleratorTypeId sets AcceleratorTypeId field to given value.

### HasAcceleratorTypeId

`func (o *SpotVmAccelerator) HasAcceleratorTypeId() bool`

HasAcceleratorTypeId returns a boolean if a field has been set.

### GetAcceleratorType

`func (o *SpotVmAccelerator) GetAcceleratorType() string`

GetAcceleratorType returns the AcceleratorType field if non-nil, zero value otherwise.

### GetAcceleratorTypeOk

`func (o *SpotVmAccelerator) GetAcceleratorTypeOk() (*string, bool)`

GetAcceleratorTypeOk returns a tuple with the AcceleratorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorType

`func (o *SpotVmAccelerator) SetAcceleratorType(v string)`

SetAcceleratorType sets AcceleratorType field to given value.

### HasAcceleratorType

`func (o *SpotVmAccelerator) HasAcceleratorType() bool`

HasAcceleratorType returns a boolean if a field has been set.

### GetAccelerators

`func (o *SpotVmAccelerator) GetAccelerators() float32`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *SpotVmAccelerator) GetAcceleratorsOk() (*float32, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *SpotVmAccelerator) SetAccelerators(v float32)`

SetAccelerators sets Accelerators field to given value.

### HasAccelerators

`func (o *SpotVmAccelerator) HasAccelerators() bool`

HasAccelerators returns a boolean if a field has been set.

### GetTotalGpuMemoryGb

`func (o *SpotVmAccelerator) GetTotalGpuMemoryGb() int32`

GetTotalGpuMemoryGb returns the TotalGpuMemoryGb field if non-nil, zero value otherwise.

### GetTotalGpuMemoryGbOk

`func (o *SpotVmAccelerator) GetTotalGpuMemoryGbOk() (*int32, bool)`

GetTotalGpuMemoryGbOk returns a tuple with the TotalGpuMemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGpuMemoryGb

`func (o *SpotVmAccelerator) SetTotalGpuMemoryGb(v int32)`

SetTotalGpuMemoryGb sets TotalGpuMemoryGb field to given value.

### HasTotalGpuMemoryGb

`func (o *SpotVmAccelerator) HasTotalGpuMemoryGb() bool`

HasTotalGpuMemoryGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


