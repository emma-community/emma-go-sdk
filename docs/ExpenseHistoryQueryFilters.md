# ExpenseHistoryQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** | Project ID | [optional] 
**PeriodStart** | **string** | Start date of the requested period of statistics | 
**PeriodEnd** | **string** | End date of the requested period of statistics | 

## Methods

### NewExpenseHistoryQueryFilters

`func NewExpenseHistoryQueryFilters(periodStart string, periodEnd string, ) *ExpenseHistoryQueryFilters`

NewExpenseHistoryQueryFilters instantiates a new ExpenseHistoryQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseHistoryQueryFiltersWithDefaults

`func NewExpenseHistoryQueryFiltersWithDefaults() *ExpenseHistoryQueryFilters`

NewExpenseHistoryQueryFiltersWithDefaults instantiates a new ExpenseHistoryQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ExpenseHistoryQueryFilters) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ExpenseHistoryQueryFilters) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ExpenseHistoryQueryFilters) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ExpenseHistoryQueryFilters) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPeriodStart

`func (o *ExpenseHistoryQueryFilters) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ExpenseHistoryQueryFilters) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ExpenseHistoryQueryFilters) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *ExpenseHistoryQueryFilters) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ExpenseHistoryQueryFilters) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ExpenseHistoryQueryFilters) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


