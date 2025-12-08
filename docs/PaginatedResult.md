# PaginatedResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **int32** | Total number of pages available | [optional] 
**TotalElements** | Pointer to **int64** | Total number of elements across all pages | [optional] 
**Size** | Pointer to **int32** | Number of elements per page | [optional] 
**Content** | Pointer to **[]interface{}** | List of elements on the current page | [optional] 
**Number** | Pointer to **int32** | Current page number (0-based index) | [optional] 
**Sort** | Pointer to [**SortObject**](SortObject.md) |  | [optional] 
**Last** | Pointer to **bool** | Indicates if this is the last page | [optional] 
**First** | Pointer to **bool** | Indicates if this is the first page | [optional] 
**NumberOfElements** | Pointer to **int32** | Number of elements on the current page | [optional] 
**Pageable** | Pointer to [**PageableObject**](PageableObject.md) |  | [optional] 
**Empty** | Pointer to **bool** | Indicates if the page is empty | [optional] 

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

`func (o *PaginatedResult) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginatedResult) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginatedResult) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PaginatedResult) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalElements

`func (o *PaginatedResult) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PaginatedResult) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PaginatedResult) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PaginatedResult) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetSize

`func (o *PaginatedResult) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PaginatedResult) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PaginatedResult) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PaginatedResult) HasSize() bool`

HasSize returns a boolean if a field has been set.

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

`func (o *PaginatedResult) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PaginatedResult) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PaginatedResult) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PaginatedResult) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

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

`func (o *PaginatedResult) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginatedResult) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginatedResult) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PaginatedResult) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetFirst

`func (o *PaginatedResult) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginatedResult) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginatedResult) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PaginatedResult) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PaginatedResult) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PaginatedResult) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PaginatedResult) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PaginatedResult) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

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

`func (o *PaginatedResult) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PaginatedResult) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PaginatedResult) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PaginatedResult) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


