# GetVmConfigs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **int32** |  | [optional] 
**TotalElements** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Content** | Pointer to [**[]VmConfiguration**](VmConfiguration.md) |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**SortObject**](SortObject.md) |  | [optional] 
**Last** | Pointer to **bool** |  | [optional] 
**First** | Pointer to **bool** |  | [optional] 
**NumberOfElements** | Pointer to **int32** |  | [optional] 
**Pageable** | Pointer to [**PageableObject**](PageableObject.md) |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetVmConfigs200Response

`func NewGetVmConfigs200Response() *GetVmConfigs200Response`

NewGetVmConfigs200Response instantiates a new GetVmConfigs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVmConfigs200ResponseWithDefaults

`func NewGetVmConfigs200ResponseWithDefaults() *GetVmConfigs200Response`

NewGetVmConfigs200ResponseWithDefaults instantiates a new GetVmConfigs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPages

`func (o *GetVmConfigs200Response) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetVmConfigs200Response) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetVmConfigs200Response) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetVmConfigs200Response) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalElements

`func (o *GetVmConfigs200Response) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *GetVmConfigs200Response) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *GetVmConfigs200Response) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *GetVmConfigs200Response) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetSize

`func (o *GetVmConfigs200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetVmConfigs200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetVmConfigs200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetVmConfigs200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetContent

`func (o *GetVmConfigs200Response) GetContent() []VmConfiguration`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetVmConfigs200Response) GetContentOk() (*[]VmConfiguration, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetVmConfigs200Response) SetContent(v []VmConfiguration)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetVmConfigs200Response) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetNumber

`func (o *GetVmConfigs200Response) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetVmConfigs200Response) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetVmConfigs200Response) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GetVmConfigs200Response) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSort

`func (o *GetVmConfigs200Response) GetSort() SortObject`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GetVmConfigs200Response) GetSortOk() (*SortObject, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GetVmConfigs200Response) SetSort(v SortObject)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GetVmConfigs200Response) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetLast

`func (o *GetVmConfigs200Response) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetVmConfigs200Response) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetVmConfigs200Response) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetVmConfigs200Response) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetFirst

`func (o *GetVmConfigs200Response) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *GetVmConfigs200Response) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *GetVmConfigs200Response) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *GetVmConfigs200Response) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *GetVmConfigs200Response) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *GetVmConfigs200Response) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *GetVmConfigs200Response) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *GetVmConfigs200Response) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetPageable

`func (o *GetVmConfigs200Response) GetPageable() PageableObject`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *GetVmConfigs200Response) GetPageableOk() (*PageableObject, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *GetVmConfigs200Response) SetPageable(v PageableObject)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *GetVmConfigs200Response) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetEmpty

`func (o *GetVmConfigs200Response) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *GetVmConfigs200Response) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *GetVmConfigs200Response) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *GetVmConfigs200Response) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


