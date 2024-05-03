# PaginatedResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **interface{}** |  | [optional] 
**TotalElements** | Pointer to **interface{}** |  | [optional] 
**Size** | Pointer to **interface{}** |  | [optional] 
**Content** | Pointer to **[]interface{}** |  | [optional] 
**Number** | Pointer to **interface{}** |  | [optional] 
**Sort** | Pointer to [**SortObject**](SortObject.md) |  | [optional] 
**Last** | Pointer to **interface{}** |  | [optional] 
**First** | Pointer to **interface{}** |  | [optional] 
**NumberOfElements** | Pointer to **interface{}** |  | [optional] 
**Pageable** | Pointer to [**PageableObject**](PageableObject.md) |  | [optional] 
**Empty** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPaginatedResult

`func NewPaginatedResult() *PaginatedResult`

NewPaginatedResult instantiates a new PaginatedResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultWithDefaults

`func NewPaginatedResultWithDefaults() *PaginatedResult`

NewPaginatedResultWithDefaults instantiates a new PaginatedResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPages

`func (o *PaginatedResult) GetTotalPages() interface{}`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginatedResult) GetTotalPagesOk() (*interface{}, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginatedResult) SetTotalPages(v interface{})`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PaginatedResult) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### SetTotalPagesNil

`func (o *PaginatedResult) SetTotalPagesNil(b bool)`

 SetTotalPagesNil sets the value for TotalPages to be an explicit nil

### UnsetTotalPages
`func (o *PaginatedResult) UnsetTotalPages()`

UnsetTotalPages ensures that no value is present for TotalPages, not even an explicit nil
### GetTotalElements

`func (o *PaginatedResult) GetTotalElements() interface{}`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PaginatedResult) GetTotalElementsOk() (*interface{}, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PaginatedResult) SetTotalElements(v interface{})`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PaginatedResult) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### SetTotalElementsNil

`func (o *PaginatedResult) SetTotalElementsNil(b bool)`

 SetTotalElementsNil sets the value for TotalElements to be an explicit nil

### UnsetTotalElements
`func (o *PaginatedResult) UnsetTotalElements()`

UnsetTotalElements ensures that no value is present for TotalElements, not even an explicit nil
### GetSize

`func (o *PaginatedResult) GetSize() interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PaginatedResult) GetSizeOk() (*interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PaginatedResult) SetSize(v interface{})`

SetSize sets Size field to given value.

### HasSize

`func (o *PaginatedResult) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *PaginatedResult) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *PaginatedResult) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetContent

`func (o *PaginatedResult) GetContent() []interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PaginatedResult) GetContentOk() (*[]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PaginatedResult) SetContent(v []interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *PaginatedResult) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetNumber

`func (o *PaginatedResult) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PaginatedResult) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PaginatedResult) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PaginatedResult) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *PaginatedResult) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *PaginatedResult) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetSort

`func (o *PaginatedResult) GetSort() SortObject`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PaginatedResult) GetSortOk() (*SortObject, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PaginatedResult) SetSort(v SortObject)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PaginatedResult) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetLast

`func (o *PaginatedResult) GetLast() interface{}`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginatedResult) GetLastOk() (*interface{}, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginatedResult) SetLast(v interface{})`

SetLast sets Last field to given value.

### HasLast

`func (o *PaginatedResult) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLastNil

`func (o *PaginatedResult) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *PaginatedResult) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil
### GetFirst

`func (o *PaginatedResult) GetFirst() interface{}`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginatedResult) GetFirstOk() (*interface{}, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginatedResult) SetFirst(v interface{})`

SetFirst sets First field to given value.

### HasFirst

`func (o *PaginatedResult) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *PaginatedResult) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *PaginatedResult) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetNumberOfElements

`func (o *PaginatedResult) GetNumberOfElements() interface{}`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PaginatedResult) GetNumberOfElementsOk() (*interface{}, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PaginatedResult) SetNumberOfElements(v interface{})`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PaginatedResult) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### SetNumberOfElementsNil

`func (o *PaginatedResult) SetNumberOfElementsNil(b bool)`

 SetNumberOfElementsNil sets the value for NumberOfElements to be an explicit nil

### UnsetNumberOfElements
`func (o *PaginatedResult) UnsetNumberOfElements()`

UnsetNumberOfElements ensures that no value is present for NumberOfElements, not even an explicit nil
### GetPageable

`func (o *PaginatedResult) GetPageable() PageableObject`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *PaginatedResult) GetPageableOk() (*PageableObject, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *PaginatedResult) SetPageable(v PageableObject)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *PaginatedResult) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetEmpty

`func (o *PaginatedResult) GetEmpty() interface{}`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PaginatedResult) GetEmptyOk() (*interface{}, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PaginatedResult) SetEmpty(v interface{})`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PaginatedResult) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### SetEmptyNil

`func (o *PaginatedResult) SetEmptyNil(b bool)`

 SetEmptyNil sets the value for Empty to be an explicit nil

### UnsetEmpty
`func (o *PaginatedResult) UnsetEmpty()`

UnsetEmpty ensures that no value is present for Empty, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


