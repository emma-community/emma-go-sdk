# VolumeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **int32** | ID of the cloud provider | [optional] 
**ProviderName** | Pointer to **string** | Name of the cloud provider | [optional] 
**LocationId** | Pointer to **int32** | Location ID | [optional] 
**LocationName** | Pointer to **string** | Location name (city or state) | [optional] 
**DataCenterId** | Pointer to **string** | ID of the data center | [optional] 
**DataCenterName** | Pointer to **string** | Name of the data center | [optional] 
**VolumeGb** | Pointer to **int32** | Capacity of the volume in gigabytes | [optional] 
**VolumeType** | Pointer to **string** | Volume type | [optional] 
**Cost** | Pointer to [**VmConfigurationCost**](VmConfigurationCost.md) |  | [optional] 

## Methods

### NewVolumeConfiguration

`func NewVolumeConfiguration() *VolumeConfiguration`

NewVolumeConfiguration instantiates a new VolumeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeConfigurationWithDefaults

`func NewVolumeConfigurationWithDefaults() *VolumeConfiguration`

NewVolumeConfigurationWithDefaults instantiates a new VolumeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *VolumeConfiguration) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *VolumeConfiguration) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *VolumeConfiguration) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *VolumeConfiguration) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderName

`func (o *VolumeConfiguration) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *VolumeConfiguration) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *VolumeConfiguration) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *VolumeConfiguration) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetLocationId

`func (o *VolumeConfiguration) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *VolumeConfiguration) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *VolumeConfiguration) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *VolumeConfiguration) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetLocationName

`func (o *VolumeConfiguration) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *VolumeConfiguration) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *VolumeConfiguration) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *VolumeConfiguration) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *VolumeConfiguration) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VolumeConfiguration) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VolumeConfiguration) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *VolumeConfiguration) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetDataCenterName

`func (o *VolumeConfiguration) GetDataCenterName() string`

GetDataCenterName returns the DataCenterName field if non-nil, zero value otherwise.

### GetDataCenterNameOk

`func (o *VolumeConfiguration) GetDataCenterNameOk() (*string, bool)`

GetDataCenterNameOk returns a tuple with the DataCenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterName

`func (o *VolumeConfiguration) SetDataCenterName(v string)`

SetDataCenterName sets DataCenterName field to given value.

### HasDataCenterName

`func (o *VolumeConfiguration) HasDataCenterName() bool`

HasDataCenterName returns a boolean if a field has been set.

### GetVolumeGb

`func (o *VolumeConfiguration) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VolumeConfiguration) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VolumeConfiguration) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.

### HasVolumeGb

`func (o *VolumeConfiguration) HasVolumeGb() bool`

HasVolumeGb returns a boolean if a field has been set.

### GetVolumeType

`func (o *VolumeConfiguration) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeConfiguration) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeConfiguration) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *VolumeConfiguration) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetCost

`func (o *VolumeConfiguration) GetCost() VmConfigurationCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *VolumeConfiguration) GetCostOk() (*VmConfigurationCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *VolumeConfiguration) SetCost(v VmConfigurationCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *VolumeConfiguration) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


