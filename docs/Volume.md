# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Volume ID | [optional] 
**AttachedToId** | Pointer to **int32** | Compute instance ID that the volume is currently attached to | [optional] 
**ProjectId** | Pointer to **int32** | Project ID | [optional] 
**SizeGb** | Pointer to **int32** | Volume size in gigabytes | [optional] 
**Type** | Pointer to **string** | Volume type | [optional] 
**IsSystem** | Pointer to **bool** | Indicates whether the volume contains OS or not | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Volume) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttachedToId

`func (o *Volume) GetAttachedToId() int32`

GetAttachedToId returns the AttachedToId field if non-nil, zero value otherwise.

### GetAttachedToIdOk

`func (o *Volume) GetAttachedToIdOk() (*int32, bool)`

GetAttachedToIdOk returns a tuple with the AttachedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedToId

`func (o *Volume) SetAttachedToId(v int32)`

SetAttachedToId sets AttachedToId field to given value.

### HasAttachedToId

`func (o *Volume) HasAttachedToId() bool`

HasAttachedToId returns a boolean if a field has been set.

### GetProjectId

`func (o *Volume) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Volume) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Volume) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Volume) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSizeGb

`func (o *Volume) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *Volume) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *Volume) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *Volume) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetType

`func (o *Volume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Volume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Volume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Volume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsSystem

`func (o *Volume) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *Volume) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *Volume) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *Volume) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


