# GetVmResizeConfigs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **int32** | Total number of pages available | [optional] 
**TotalElements** | Pointer to **int64** | Total number of elements across all pages | [optional] 
**Size** | Pointer to **int32** | Number of elements per page | [optional] 
**Content** | Pointer to [**[]VmResizeConfiguration**](VmResizeConfiguration.md) |  | [optional] 
**Number** | Pointer to **int32** | Current page number (0-based index) | [optional] 
**Sort** | Pointer to [**SortObject**](SortObject.md) |  | [optional] 
**Last** | Pointer to **bool** | Indicates if this is the last page | [optional] 
**First** | Pointer to **bool** | Indicates if this is the first page | [optional] 
**NumberOfElements** | Pointer to **int32** | Number of elements on the current page | [optional] 
**Pageable** | Pointer to [**PageableObject**](PageableObject.md) |  | [optional] 
**Empty** | Pointer to **bool** | Indicates if the page is empty | [optional] 

## Methods

### NewGetVmResizeConfigs200Response

`func NewGetVmResizeConfigs200Response() *GetVmResizeConfigs200Response`

NewGetVmResizeConfigs200Response instantiates a new GetVmResizeConfigs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVmResizeConfigs200ResponseWithDefaults

`func NewGetVmResizeConfigs200ResponseWithDefaults() *GetVmResizeConfigs200Response`

NewGetVmResizeConfigs200ResponseWithDefaults instantiates a new GetVmResizeConfigs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPages

`func (o *GetVmResizeConfigs200Response) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetVmResizeConfigs200Response) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetVmResizeConfigs200Response) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetVmResizeConfigs200Response) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalElements

`func (o *GetVmResizeConfigs200Response) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *GetVmResizeConfigs200Response) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *GetVmResizeConfigs200Response) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *GetVmResizeConfigs200Response) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetSize

`func (o *GetVmResizeConfigs200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetVmResizeConfigs200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetVmResizeConfigs200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetVmResizeConfigs200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetContent

`func (o *GetVmResizeConfigs200Response) GetContent() []VmResizeConfiguration`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetVmResizeConfigs200Response) GetContentOk() (*[]VmResizeConfiguration, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetVmResizeConfigs200Response) SetContent(v []VmResizeConfiguration)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetVmResizeConfigs200Response) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetNumber

`func (o *GetVmResizeConfigs200Response) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetVmResizeConfigs200Response) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetVmResizeConfigs200Response) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GetVmResizeConfigs200Response) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSort

`func (o *GetVmResizeConfigs200Response) GetSort() SortObject`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GetVmResizeConfigs200Response) GetSortOk() (*SortObject, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GetVmResizeConfigs200Response) SetSort(v SortObject)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GetVmResizeConfigs200Response) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetLast

`func (o *GetVmResizeConfigs200Response) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetVmResizeConfigs200Response) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetVmResizeConfigs200Response) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetVmResizeConfigs200Response) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetFirst

`func (o *GetVmResizeConfigs200Response) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *GetVmResizeConfigs200Response) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *GetVmResizeConfigs200Response) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *GetVmResizeConfigs200Response) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *GetVmResizeConfigs200Response) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *GetVmResizeConfigs200Response) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *GetVmResizeConfigs200Response) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *GetVmResizeConfigs200Response) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetPageable

`func (o *GetVmResizeConfigs200Response) GetPageable() PageableObject`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *GetVmResizeConfigs200Response) GetPageableOk() (*PageableObject, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *GetVmResizeConfigs200Response) SetPageable(v PageableObject)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *GetVmResizeConfigs200Response) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetEmpty

`func (o *GetVmResizeConfigs200Response) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *GetVmResizeConfigs200Response) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *GetVmResizeConfigs200Response) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *GetVmResizeConfigs200Response) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


