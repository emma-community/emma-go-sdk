# VmMonitoringQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**VmId** | **int32** | ID of VM | 
**Filters** | [**VmMonitoringQueryFilters**](VmMonitoringQueryFilters.md) |  | 

## Methods

### NewVmMonitoringQuery

`func NewVmMonitoringQuery(datasetName string, vmId int32, filters VmMonitoringQueryFilters, ) *VmMonitoringQuery`

NewVmMonitoringQuery instantiates a new VmMonitoringQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmMonitoringQueryWithDefaults

`func NewVmMonitoringQueryWithDefaults() *VmMonitoringQuery`

NewVmMonitoringQueryWithDefaults instantiates a new VmMonitoringQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *VmMonitoringQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *VmMonitoringQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *VmMonitoringQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetVmId

`func (o *VmMonitoringQuery) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VmMonitoringQuery) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VmMonitoringQuery) SetVmId(v int32)`

SetVmId sets VmId field to given value.


### GetFilters

`func (o *VmMonitoringQuery) GetFilters() VmMonitoringQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *VmMonitoringQuery) GetFiltersOk() (*VmMonitoringQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *VmMonitoringQuery) SetFilters(v VmMonitoringQueryFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


