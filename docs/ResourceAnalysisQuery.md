# ResourceAnalysisQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**Filters** | [**ResourceAnalysisQueryFilters**](ResourceAnalysisQueryFilters.md) |  | 

## Methods

### NewResourceAnalysisQuery

`func NewResourceAnalysisQuery(datasetName string, filters ResourceAnalysisQueryFilters, ) *ResourceAnalysisQuery`

NewResourceAnalysisQuery instantiates a new ResourceAnalysisQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAnalysisQueryWithDefaults

`func NewResourceAnalysisQueryWithDefaults() *ResourceAnalysisQuery`

NewResourceAnalysisQueryWithDefaults instantiates a new ResourceAnalysisQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *ResourceAnalysisQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *ResourceAnalysisQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *ResourceAnalysisQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetFilters

`func (o *ResourceAnalysisQuery) GetFilters() ResourceAnalysisQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ResourceAnalysisQuery) GetFiltersOk() (*ResourceAnalysisQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ResourceAnalysisQuery) SetFilters(v ResourceAnalysisQueryFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


