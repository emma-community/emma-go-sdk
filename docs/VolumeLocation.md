# VolumeLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the data center location | [optional] 
**Name** | Pointer to **string** | Name of the data center location (city or state) | [optional] 
**Continent** | Pointer to **string** | Name of the geographical continent | [optional] 
**Region** | Pointer to **string** | Name of the geographical region | [optional] 
**Latitude** | Pointer to **float64** | Approximate latitude of the geographical location | [optional] 
**Longitude** | Pointer to **float64** | Approximate longitude of the geographical location | [optional] 

## Methods

### NewVolumeLocation

`func NewVolumeLocation() *VolumeLocation`

NewVolumeLocation instantiates a new VolumeLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeLocationWithDefaults

`func NewVolumeLocationWithDefaults() *VolumeLocation`

NewVolumeLocationWithDefaults instantiates a new VolumeLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumeLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeLocation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VolumeLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContinent

`func (o *VolumeLocation) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *VolumeLocation) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *VolumeLocation) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *VolumeLocation) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetRegion

`func (o *VolumeLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VolumeLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VolumeLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VolumeLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetLatitude

`func (o *VolumeLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VolumeLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VolumeLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VolumeLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *VolumeLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VolumeLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VolumeLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VolumeLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


