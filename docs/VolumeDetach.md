# VolumeDetach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with a virtual machine | 
**VolumeId** | **int32** | Volume ID to detach from the virtual machine | 

## Methods

### NewVolumeDetach

`func NewVolumeDetach(action string, volumeId int32, ) *VolumeDetach`

NewVolumeDetach instantiates a new VolumeDetach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeDetachWithDefaults

`func NewVolumeDetachWithDefaults() *VolumeDetach`

NewVolumeDetachWithDefaults instantiates a new VolumeDetach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VolumeDetach) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VolumeDetach) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VolumeDetach) SetAction(v string)`

SetAction sets Action field to given value.


### GetVolumeId

`func (o *VolumeDetach) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeDetach) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeDetach) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


