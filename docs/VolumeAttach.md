# VolumeAttach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with a virtual machine | 
**VolumeId** | **int32** | Volume ID to attach to the virtual machine | 

## Methods

### NewVolumeAttach

`func NewVolumeAttach(action string, volumeId int32, ) *VolumeAttach`

NewVolumeAttach instantiates a new VolumeAttach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeAttachWithDefaults

`func NewVolumeAttachWithDefaults() *VolumeAttach`

NewVolumeAttachWithDefaults instantiates a new VolumeAttach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VolumeAttach) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VolumeAttach) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VolumeAttach) SetAction(v string)`

SetAction sets Action field to given value.


### GetVolumeId

`func (o *VolumeAttach) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeAttach) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeAttach) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


