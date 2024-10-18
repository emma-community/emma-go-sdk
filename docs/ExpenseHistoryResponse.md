# ExpenseHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroup** | Pointer to **string** | Resource group name | [optional] 
**Timecode** | Pointer to **string** | Timecode | [optional] 
**ExpensesPerDay** | Pointer to **float32** | Daily group expenses | [optional] 

## Methods

### NewExpenseHistoryResponse

`func NewExpenseHistoryResponse() *ExpenseHistoryResponse`

NewExpenseHistoryResponse instantiates a new ExpenseHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseHistoryResponseWithDefaults

`func NewExpenseHistoryResponseWithDefaults() *ExpenseHistoryResponse`

NewExpenseHistoryResponseWithDefaults instantiates a new ExpenseHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroup

`func (o *ExpenseHistoryResponse) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ExpenseHistoryResponse) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ExpenseHistoryResponse) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *ExpenseHistoryResponse) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimecode

`func (o *ExpenseHistoryResponse) GetTimecode() string`

GetTimecode returns the Timecode field if non-nil, zero value otherwise.

### GetTimecodeOk

`func (o *ExpenseHistoryResponse) GetTimecodeOk() (*string, bool)`

GetTimecodeOk returns a tuple with the Timecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecode

`func (o *ExpenseHistoryResponse) SetTimecode(v string)`

SetTimecode sets Timecode field to given value.

### HasTimecode

`func (o *ExpenseHistoryResponse) HasTimecode() bool`

HasTimecode returns a boolean if a field has been set.

### GetExpensesPerDay

`func (o *ExpenseHistoryResponse) GetExpensesPerDay() float32`

GetExpensesPerDay returns the ExpensesPerDay field if non-nil, zero value otherwise.

### GetExpensesPerDayOk

`func (o *ExpenseHistoryResponse) GetExpensesPerDayOk() (*float32, bool)`

GetExpensesPerDayOk returns a tuple with the ExpensesPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpensesPerDay

`func (o *ExpenseHistoryResponse) SetExpensesPerDay(v float32)`

SetExpensesPerDay sets ExpensesPerDay field to given value.

### HasExpensesPerDay

`func (o *ExpenseHistoryResponse) HasExpensesPerDay() bool`

HasExpensesPerDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


