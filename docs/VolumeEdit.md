# VolumeEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with a volume  **Note:** To retrieve available configurations for DigitalOcean system volumes, use the VM &#x60;edithardware&#x60; action  | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 

## Methods

### NewVolumeEdit

`func NewVolumeEdit(action string, volumeGb int32, ) *VolumeEdit`

NewVolumeEdit instantiates a new VolumeEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeEditWithDefaults

`func NewVolumeEditWithDefaults() *VolumeEdit`

NewVolumeEditWithDefaults instantiates a new VolumeEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VolumeEdit) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VolumeEdit) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VolumeEdit) SetAction(v string)`

SetAction sets Action field to given value.


### GetVolumeGb

`func (o *VolumeEdit) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VolumeEdit) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VolumeEdit) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


