# GetVmConfigs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **interface{}** |  | [optional] 
**TotalElements** | Pointer to **interface{}** |  | [optional] 
**Size** | Pointer to **interface{}** |  | [optional] 
**Content** | Pointer to [**[]VmConfiguration**](VmConfiguration.md) |  | [optional] 
**Number** | Pointer to **interface{}** |  | [optional] 
**Sort** | Pointer to [**SortObject**](SortObject.md) |  | [optional] 
**Last** | Pointer to **interface{}** |  | [optional] 
**First** | Pointer to **interface{}** |  | [optional] 
**NumberOfElements** | Pointer to **interface{}** |  | [optional] 
**Pageable** | Pointer to [**PageableObject**](PageableObject.md) |  | [optional] 
**Empty** | Pointer to **interface{}** |  | [optional] 

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

`func (o *GetVmConfigs200Response) GetTotalPages() interface{}`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetVmConfigs200Response) GetTotalPagesOk() (*interface{}, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetVmConfigs200Response) SetTotalPages(v interface{})`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetVmConfigs200Response) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### SetTotalPagesNil

`func (o *GetVmConfigs200Response) SetTotalPagesNil(b bool)`

 SetTotalPagesNil sets the value for TotalPages to be an explicit nil

### UnsetTotalPages
`func (o *GetVmConfigs200Response) UnsetTotalPages()`

UnsetTotalPages ensures that no value is present for TotalPages, not even an explicit nil
### GetTotalElements

`func (o *GetVmConfigs200Response) GetTotalElements() interface{}`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *GetVmConfigs200Response) GetTotalElementsOk() (*interface{}, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *GetVmConfigs200Response) SetTotalElements(v interface{})`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *GetVmConfigs200Response) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### SetTotalElementsNil

`func (o *GetVmConfigs200Response) SetTotalElementsNil(b bool)`

 SetTotalElementsNil sets the value for TotalElements to be an explicit nil

### UnsetTotalElements
`func (o *GetVmConfigs200Response) UnsetTotalElements()`

UnsetTotalElements ensures that no value is present for TotalElements, not even an explicit nil
### GetSize

`func (o *GetVmConfigs200Response) GetSize() interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetVmConfigs200Response) GetSizeOk() (*interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetVmConfigs200Response) SetSize(v interface{})`

SetSize sets Size field to given value.

### HasSize

`func (o *GetVmConfigs200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *GetVmConfigs200Response) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GetVmConfigs200Response) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
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

`func (o *GetVmConfigs200Response) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetVmConfigs200Response) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetVmConfigs200Response) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GetVmConfigs200Response) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GetVmConfigs200Response) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GetVmConfigs200Response) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
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

`func (o *GetVmConfigs200Response) GetLast() interface{}`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetVmConfigs200Response) GetLastOk() (*interface{}, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetVmConfigs200Response) SetLast(v interface{})`

SetLast sets Last field to given value.

### HasLast

`func (o *GetVmConfigs200Response) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLastNil

`func (o *GetVmConfigs200Response) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *GetVmConfigs200Response) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil
### GetFirst

`func (o *GetVmConfigs200Response) GetFirst() interface{}`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *GetVmConfigs200Response) GetFirstOk() (*interface{}, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *GetVmConfigs200Response) SetFirst(v interface{})`

SetFirst sets First field to given value.

### HasFirst

`func (o *GetVmConfigs200Response) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *GetVmConfigs200Response) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *GetVmConfigs200Response) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetNumberOfElements

`func (o *GetVmConfigs200Response) GetNumberOfElements() interface{}`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *GetVmConfigs200Response) GetNumberOfElementsOk() (*interface{}, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *GetVmConfigs200Response) SetNumberOfElements(v interface{})`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *GetVmConfigs200Response) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### SetNumberOfElementsNil

`func (o *GetVmConfigs200Response) SetNumberOfElementsNil(b bool)`

 SetNumberOfElementsNil sets the value for NumberOfElements to be an explicit nil

### UnsetNumberOfElements
`func (o *GetVmConfigs200Response) UnsetNumberOfElements()`

UnsetNumberOfElements ensures that no value is present for NumberOfElements, not even an explicit nil
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

`func (o *GetVmConfigs200Response) GetEmpty() interface{}`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *GetVmConfigs200Response) GetEmptyOk() (*interface{}, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *GetVmConfigs200Response) SetEmpty(v interface{})`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *GetVmConfigs200Response) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### SetEmptyNil

`func (o *GetVmConfigs200Response) SetEmptyNil(b bool)`

 SetEmptyNil sets the value for Empty to be an explicit nil

### UnsetEmpty
`func (o *GetVmConfigs200Response) UnsetEmpty()`

UnsetEmpty ensures that no value is present for Empty, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


