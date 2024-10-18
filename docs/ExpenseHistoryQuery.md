# ExpenseHistoryQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**GroupBy** | **[]string** | List of grouping categories | 
**Filters** | [**ExpenseHistoryQueryFilters**](ExpenseHistoryQueryFilters.md) |  | 

## Methods

### NewExpenseHistoryQuery

`func NewExpenseHistoryQuery(datasetName string, groupBy []string, filters ExpenseHistoryQueryFilters, ) *ExpenseHistoryQuery`

NewExpenseHistoryQuery instantiates a new ExpenseHistoryQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseHistoryQueryWithDefaults

`func NewExpenseHistoryQueryWithDefaults() *ExpenseHistoryQuery`

NewExpenseHistoryQueryWithDefaults instantiates a new ExpenseHistoryQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *ExpenseHistoryQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *ExpenseHistoryQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *ExpenseHistoryQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetGroupBy

`func (o *ExpenseHistoryQuery) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ExpenseHistoryQuery) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ExpenseHistoryQuery) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.


### GetFilters

`func (o *ExpenseHistoryQuery) GetFilters() ExpenseHistoryQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ExpenseHistoryQuery) GetFiltersOk() (*ExpenseHistoryQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ExpenseHistoryQuery) SetFilters(v ExpenseHistoryQueryFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


