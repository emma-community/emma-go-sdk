# VolumeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataCenterId** | **string** | ID of the data center | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 
**VolumeType** | **string** | Volume type | 
**AttachedToId** | Pointer to **int32** | Target virtual machine for attaching the volume | [optional] 

## Methods

### NewVolumeCreate

`func NewVolumeCreate(dataCenterId string, volumeGb int32, volumeType string, ) *VolumeCreate`

NewVolumeCreate instantiates a new VolumeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCreateWithDefaults

`func NewVolumeCreateWithDefaults() *VolumeCreate`

NewVolumeCreateWithDefaults instantiates a new VolumeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataCenterId

`func (o *VolumeCreate) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VolumeCreate) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VolumeCreate) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetVolumeGb

`func (o *VolumeCreate) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VolumeCreate) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VolumeCreate) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.


### GetVolumeType

`func (o *VolumeCreate) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeCreate) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeCreate) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetAttachedToId

`func (o *VolumeCreate) GetAttachedToId() int32`

GetAttachedToId returns the AttachedToId field if non-nil, zero value otherwise.

### GetAttachedToIdOk

`func (o *VolumeCreate) GetAttachedToIdOk() (*int32, bool)`

GetAttachedToIdOk returns a tuple with the AttachedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedToId

`func (o *VolumeCreate) SetAttachedToId(v int32)`

SetAttachedToId sets AttachedToId field to given value.

### HasAttachedToId

`func (o *VolumeCreate) HasAttachedToId() bool`

HasAttachedToId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


