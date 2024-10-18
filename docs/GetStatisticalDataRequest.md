# GetStatisticalDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**GroupBy** | **[]string** | List of grouping categories | 
**Filters** | [**VmMonitoringQueryFilters**](VmMonitoringQueryFilters.md) |  | 
**CoreClusterId** | **int32** | ID of Kubernetes cluster | 
**VmId** | **int32** | ID of VM | 

## Methods

### NewGetStatisticalDataRequest

`func NewGetStatisticalDataRequest(datasetName string, groupBy []string, filters VmMonitoringQueryFilters, coreClusterId int32, vmId int32, ) *GetStatisticalDataRequest`

NewGetStatisticalDataRequest instantiates a new GetStatisticalDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatisticalDataRequestWithDefaults

`func NewGetStatisticalDataRequestWithDefaults() *GetStatisticalDataRequest`

NewGetStatisticalDataRequestWithDefaults instantiates a new GetStatisticalDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *GetStatisticalDataRequest) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *GetStatisticalDataRequest) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *GetStatisticalDataRequest) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetGroupBy

`func (o *GetStatisticalDataRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GetStatisticalDataRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GetStatisticalDataRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.


### GetFilters

`func (o *GetStatisticalDataRequest) GetFilters() VmMonitoringQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetStatisticalDataRequest) GetFiltersOk() (*VmMonitoringQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetStatisticalDataRequest) SetFilters(v VmMonitoringQueryFilters)`

SetFilters sets Filters field to given value.


### GetCoreClusterId

`func (o *GetStatisticalDataRequest) GetCoreClusterId() int32`

GetCoreClusterId returns the CoreClusterId field if non-nil, zero value otherwise.

### GetCoreClusterIdOk

`func (o *GetStatisticalDataRequest) GetCoreClusterIdOk() (*int32, bool)`

GetCoreClusterIdOk returns a tuple with the CoreClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreClusterId

`func (o *GetStatisticalDataRequest) SetCoreClusterId(v int32)`

SetCoreClusterId sets CoreClusterId field to given value.


### GetVmId

`func (o *GetStatisticalDataRequest) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *GetStatisticalDataRequest) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *GetStatisticalDataRequest) SetVmId(v int32)`

SetVmId sets VmId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


