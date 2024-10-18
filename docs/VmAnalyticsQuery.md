# VmAnalyticsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**VmId** | **int32** | ID of VM | 

## Methods

### NewVmAnalyticsQuery

`func NewVmAnalyticsQuery(datasetName string, vmId int32, ) *VmAnalyticsQuery`

NewVmAnalyticsQuery instantiates a new VmAnalyticsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmAnalyticsQueryWithDefaults

`func NewVmAnalyticsQueryWithDefaults() *VmAnalyticsQuery`

NewVmAnalyticsQueryWithDefaults instantiates a new VmAnalyticsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *VmAnalyticsQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *VmAnalyticsQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *VmAnalyticsQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetVmId

`func (o *VmAnalyticsQuery) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VmAnalyticsQuery) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VmAnalyticsQuery) SetVmId(v int32)`

SetVmId sets VmId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


