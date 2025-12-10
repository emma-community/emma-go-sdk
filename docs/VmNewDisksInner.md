# VmNewDisksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Volume ID | [optional] 
**SizeGb** | Pointer to **int32** | Volume size in gigabytes | [optional] 
**TypeId** | Pointer to **int32** | ID of the volume type | [optional] 
**Type** | Pointer to **string** | Volume type | [optional] 
**IsBootable** | Pointer to **bool** | Indicates whether the volume is bootable or not | [optional] 

## Methods

### NewVmNewDisksInner

`func NewVmNewDisksInner() *VmNewDisksInner`

NewVmNewDisksInner instantiates a new VmNewDisksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmNewDisksInnerWithDefaults

`func NewVmNewDisksInnerWithDefaults() *VmNewDisksInner`

NewVmNewDisksInnerWithDefaults instantiates a new VmNewDisksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmNewDisksInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmNewDisksInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmNewDisksInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VmNewDisksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSizeGb

`func (o *VmNewDisksInner) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *VmNewDisksInner) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *VmNewDisksInner) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *VmNewDisksInner) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetTypeId

`func (o *VmNewDisksInner) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *VmNewDisksInner) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *VmNewDisksInner) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *VmNewDisksInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetType

`func (o *VmNewDisksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmNewDisksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmNewDisksInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VmNewDisksInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsBootable

`func (o *VmNewDisksInner) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *VmNewDisksInner) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *VmNewDisksInner) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.

### HasIsBootable

`func (o *VmNewDisksInner) HasIsBootable() bool`

HasIsBootable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


