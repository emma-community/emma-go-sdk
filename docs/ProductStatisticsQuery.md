# ProductStatisticsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**Filters** | [**ProductStatisticsQueryFilters**](ProductStatisticsQueryFilters.md) |  | 

## Methods

### NewProductStatisticsQuery

`func NewProductStatisticsQuery(datasetName string, filters ProductStatisticsQueryFilters, ) *ProductStatisticsQuery`

NewProductStatisticsQuery instantiates a new ProductStatisticsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductStatisticsQueryWithDefaults

`func NewProductStatisticsQueryWithDefaults() *ProductStatisticsQuery`

NewProductStatisticsQueryWithDefaults instantiates a new ProductStatisticsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *ProductStatisticsQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *ProductStatisticsQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *ProductStatisticsQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetFilters

`func (o *ProductStatisticsQuery) GetFilters() ProductStatisticsQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ProductStatisticsQuery) GetFiltersOk() (*ProductStatisticsQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ProductStatisticsQuery) SetFilters(v ProductStatisticsQueryFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


