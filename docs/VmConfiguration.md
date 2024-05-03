# VmConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProviderId** | Pointer to **int32** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **int32** |  | [optional] 
**LocationName** | Pointer to **string** |  | [optional] 
**DataCenterId** | Pointer to **string** |  | [optional] 
**DataCenterName** | Pointer to **string** |  | [optional] 
**OsId** | Pointer to **int32** |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**CloudNetworkTypes** | Pointer to **[]string** |  | [optional] 
**VCpuType** | Pointer to **string** |  | [optional] 
**VCpu** | Pointer to **int32** |  | [optional] 
**RamGb** | Pointer to **int32** |  | [optional] 
**VolumeGb** | Pointer to **int32** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to [**VmConfigurationCost**](VmConfigurationCost.md) |  | [optional] 

## Methods

### NewVmConfiguration

`func NewVmConfiguration() *VmConfiguration`

NewVmConfiguration instantiates a new VmConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmConfigurationWithDefaults

`func NewVmConfigurationWithDefaults() *VmConfiguration`

NewVmConfigurationWithDefaults instantiates a new VmConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VmConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderId

`func (o *VmConfiguration) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *VmConfiguration) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *VmConfiguration) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *VmConfiguration) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderName

`func (o *VmConfiguration) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *VmConfiguration) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *VmConfiguration) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *VmConfiguration) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetLocationId

`func (o *VmConfiguration) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *VmConfiguration) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *VmConfiguration) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *VmConfiguration) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetLocationName

`func (o *VmConfiguration) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *VmConfiguration) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *VmConfiguration) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *VmConfiguration) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *VmConfiguration) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VmConfiguration) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VmConfiguration) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *VmConfiguration) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetDataCenterName

`func (o *VmConfiguration) GetDataCenterName() string`

GetDataCenterName returns the DataCenterName field if non-nil, zero value otherwise.

### GetDataCenterNameOk

`func (o *VmConfiguration) GetDataCenterNameOk() (*string, bool)`

GetDataCenterNameOk returns a tuple with the DataCenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterName

`func (o *VmConfiguration) SetDataCenterName(v string)`

SetDataCenterName sets DataCenterName field to given value.

### HasDataCenterName

`func (o *VmConfiguration) HasDataCenterName() bool`

HasDataCenterName returns a boolean if a field has been set.

### GetOsId

`func (o *VmConfiguration) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *VmConfiguration) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *VmConfiguration) SetOsId(v int32)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *VmConfiguration) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### GetOsType

`func (o *VmConfiguration) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *VmConfiguration) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *VmConfiguration) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *VmConfiguration) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetOsVersion

`func (o *VmConfiguration) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *VmConfiguration) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *VmConfiguration) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *VmConfiguration) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetCloudNetworkTypes

`func (o *VmConfiguration) GetCloudNetworkTypes() []string`

GetCloudNetworkTypes returns the CloudNetworkTypes field if non-nil, zero value otherwise.

### GetCloudNetworkTypesOk

`func (o *VmConfiguration) GetCloudNetworkTypesOk() (*[]string, bool)`

GetCloudNetworkTypesOk returns a tuple with the CloudNetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkTypes

`func (o *VmConfiguration) SetCloudNetworkTypes(v []string)`

SetCloudNetworkTypes sets CloudNetworkTypes field to given value.

### HasCloudNetworkTypes

`func (o *VmConfiguration) HasCloudNetworkTypes() bool`

HasCloudNetworkTypes returns a boolean if a field has been set.

### GetVCpuType

`func (o *VmConfiguration) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmConfiguration) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmConfiguration) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *VmConfiguration) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *VmConfiguration) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmConfiguration) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmConfiguration) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *VmConfiguration) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *VmConfiguration) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmConfiguration) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmConfiguration) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *VmConfiguration) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetVolumeGb

`func (o *VmConfiguration) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VmConfiguration) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VmConfiguration) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.

### HasVolumeGb

`func (o *VmConfiguration) HasVolumeGb() bool`

HasVolumeGb returns a boolean if a field has been set.

### GetVolumeType

`func (o *VmConfiguration) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VmConfiguration) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VmConfiguration) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *VmConfiguration) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetCost

`func (o *VmConfiguration) GetCost() VmConfigurationCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *VmConfiguration) GetCostOk() (*VmConfigurationCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *VmConfiguration) SetCost(v VmConfigurationCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *VmConfiguration) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


