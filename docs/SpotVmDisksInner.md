# SpotVmDisksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Volume ID | [optional] 
**Name** | Pointer to **string** | Volume name | [optional] 
**SizeGb** | Pointer to **int32** | Volume size in gigabytes | [optional] 
**TypeId** | Pointer to **int32** | ID of the volume type | [optional] 
**Type** | Pointer to **string** | Volume type | [optional] 
**IsBootable** | Pointer to **bool** | Indicates whether the volume is bootable or not | [optional] 

## Methods

### NewSpotVmDisksInner

`func NewSpotVmDisksInner() *SpotVmDisksInner`

NewSpotVmDisksInner instantiates a new SpotVmDisksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotVmDisksInnerWithDefaults

`func NewSpotVmDisksInnerWithDefaults() *SpotVmDisksInner`

NewSpotVmDisksInnerWithDefaults instantiates a new SpotVmDisksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpotVmDisksInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpotVmDisksInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpotVmDisksInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SpotVmDisksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SpotVmDisksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpotVmDisksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpotVmDisksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpotVmDisksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSizeGb

`func (o *SpotVmDisksInner) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *SpotVmDisksInner) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *SpotVmDisksInner) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *SpotVmDisksInner) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetTypeId

`func (o *SpotVmDisksInner) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *SpotVmDisksInner) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *SpotVmDisksInner) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *SpotVmDisksInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetType

`func (o *SpotVmDisksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpotVmDisksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpotVmDisksInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpotVmDisksInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsBootable

`func (o *SpotVmDisksInner) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *SpotVmDisksInner) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *SpotVmDisksInner) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.

### HasIsBootable

`func (o *SpotVmDisksInner) HasIsBootable() bool`

HasIsBootable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


