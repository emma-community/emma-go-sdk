# SystemVolumeConfiguration

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

### NewSystemVolumeConfiguration

`func NewSystemVolumeConfiguration() *SystemVolumeConfiguration`

NewSystemVolumeConfiguration instantiates a new SystemVolumeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemVolumeConfigurationWithDefaults

`func NewSystemVolumeConfigurationWithDefaults() *SystemVolumeConfiguration`

NewSystemVolumeConfigurationWithDefaults instantiates a new SystemVolumeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *SystemVolumeConfiguration) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *SystemVolumeConfiguration) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *SystemVolumeConfiguration) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *SystemVolumeConfiguration) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderName

`func (o *SystemVolumeConfiguration) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *SystemVolumeConfiguration) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *SystemVolumeConfiguration) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *SystemVolumeConfiguration) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetLocationId

`func (o *SystemVolumeConfiguration) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *SystemVolumeConfiguration) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *SystemVolumeConfiguration) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *SystemVolumeConfiguration) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetLocationName

`func (o *SystemVolumeConfiguration) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *SystemVolumeConfiguration) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *SystemVolumeConfiguration) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *SystemVolumeConfiguration) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *SystemVolumeConfiguration) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *SystemVolumeConfiguration) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *SystemVolumeConfiguration) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *SystemVolumeConfiguration) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetDataCenterName

`func (o *SystemVolumeConfiguration) GetDataCenterName() string`

GetDataCenterName returns the DataCenterName field if non-nil, zero value otherwise.

### GetDataCenterNameOk

`func (o *SystemVolumeConfiguration) GetDataCenterNameOk() (*string, bool)`

GetDataCenterNameOk returns a tuple with the DataCenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterName

`func (o *SystemVolumeConfiguration) SetDataCenterName(v string)`

SetDataCenterName sets DataCenterName field to given value.

### HasDataCenterName

`func (o *SystemVolumeConfiguration) HasDataCenterName() bool`

HasDataCenterName returns a boolean if a field has been set.

### GetVolumeGb

`func (o *SystemVolumeConfiguration) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *SystemVolumeConfiguration) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *SystemVolumeConfiguration) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.

### HasVolumeGb

`func (o *SystemVolumeConfiguration) HasVolumeGb() bool`

HasVolumeGb returns a boolean if a field has been set.

### GetVolumeType

`func (o *SystemVolumeConfiguration) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *SystemVolumeConfiguration) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *SystemVolumeConfiguration) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *SystemVolumeConfiguration) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetCost

`func (o *SystemVolumeConfiguration) GetCost() VmConfigurationCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SystemVolumeConfiguration) GetCostOk() (*VmConfigurationCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SystemVolumeConfiguration) SetCost(v VmConfigurationCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SystemVolumeConfiguration) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


