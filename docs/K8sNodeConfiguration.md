# K8sNodeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Configuration ID | [optional] 
**ProviderId** | Pointer to **int32** | ID of the cloud provider | [optional] 
**ProviderName** | Pointer to **string** | Name of the cloud provider | [optional] 
**LocationId** | Pointer to **int32** | Location ID | [optional] 
**LocationName** | Pointer to **string** | Location name (city or state) | [optional] 
**DataCenterId** | Pointer to **string** | ID of the data center | [optional] 
**DataCenterName** | Pointer to **string** | Name of the data center | [optional] 
**OsId** | Pointer to **int32** | ID of the operating system | [optional] 
**OsType** | Pointer to **string** | Type of the operating system | [optional] 
**OsVersion** | Pointer to **string** | Version of the operating system | [optional] 
**CloudNetworkTypes** | Pointer to **[]string** | List of the cloud network types | [optional] 
**VCpuType** | Pointer to **string** | Type of virtual Central Processing Units (vCPUs) | [optional] 
**VCpu** | Pointer to **int32** | Number of virtual Central Processing Units (vCPUs) | [optional] 
**RamGb** | Pointer to **int32** | Capacity of the RAM in gigabytes | [optional] 
**VolumeGb** | Pointer to **int32** | Capacity of the volume in gigabytes | [optional] 
**VolumeType** | Pointer to **string** | Volume type | [optional] 
**Accelerator** | Pointer to [**K8sNodeConfigurationAccelerator**](K8sNodeConfigurationAccelerator.md) |  | [optional] 
**LocalDisks** | Pointer to [**K8sNodeConfigurationLocalDisks**](K8sNodeConfigurationLocalDisks.md) |  | [optional] 
**Cost** | Pointer to [**K8sNodeConfigurationCost**](K8sNodeConfigurationCost.md) |  | [optional] 

## Methods

### NewK8sNodeConfiguration

`func NewK8sNodeConfiguration() *K8sNodeConfiguration`

NewK8sNodeConfiguration instantiates a new K8sNodeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sNodeConfigurationWithDefaults

`func NewK8sNodeConfigurationWithDefaults() *K8sNodeConfiguration`

NewK8sNodeConfigurationWithDefaults instantiates a new K8sNodeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *K8sNodeConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *K8sNodeConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *K8sNodeConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *K8sNodeConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderId

`func (o *K8sNodeConfiguration) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *K8sNodeConfiguration) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *K8sNodeConfiguration) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *K8sNodeConfiguration) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderName

`func (o *K8sNodeConfiguration) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *K8sNodeConfiguration) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *K8sNodeConfiguration) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *K8sNodeConfiguration) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetLocationId

`func (o *K8sNodeConfiguration) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *K8sNodeConfiguration) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *K8sNodeConfiguration) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *K8sNodeConfiguration) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetLocationName

`func (o *K8sNodeConfiguration) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *K8sNodeConfiguration) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *K8sNodeConfiguration) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *K8sNodeConfiguration) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *K8sNodeConfiguration) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *K8sNodeConfiguration) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *K8sNodeConfiguration) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *K8sNodeConfiguration) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetDataCenterName

`func (o *K8sNodeConfiguration) GetDataCenterName() string`

GetDataCenterName returns the DataCenterName field if non-nil, zero value otherwise.

### GetDataCenterNameOk

`func (o *K8sNodeConfiguration) GetDataCenterNameOk() (*string, bool)`

GetDataCenterNameOk returns a tuple with the DataCenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterName

`func (o *K8sNodeConfiguration) SetDataCenterName(v string)`

SetDataCenterName sets DataCenterName field to given value.

### HasDataCenterName

`func (o *K8sNodeConfiguration) HasDataCenterName() bool`

HasDataCenterName returns a boolean if a field has been set.

### GetOsId

`func (o *K8sNodeConfiguration) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *K8sNodeConfiguration) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *K8sNodeConfiguration) SetOsId(v int32)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *K8sNodeConfiguration) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### GetOsType

`func (o *K8sNodeConfiguration) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *K8sNodeConfiguration) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *K8sNodeConfiguration) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *K8sNodeConfiguration) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetOsVersion

`func (o *K8sNodeConfiguration) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *K8sNodeConfiguration) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *K8sNodeConfiguration) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *K8sNodeConfiguration) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetCloudNetworkTypes

`func (o *K8sNodeConfiguration) GetCloudNetworkTypes() []string`

GetCloudNetworkTypes returns the CloudNetworkTypes field if non-nil, zero value otherwise.

### GetCloudNetworkTypesOk

`func (o *K8sNodeConfiguration) GetCloudNetworkTypesOk() (*[]string, bool)`

GetCloudNetworkTypesOk returns a tuple with the CloudNetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkTypes

`func (o *K8sNodeConfiguration) SetCloudNetworkTypes(v []string)`

SetCloudNetworkTypes sets CloudNetworkTypes field to given value.

### HasCloudNetworkTypes

`func (o *K8sNodeConfiguration) HasCloudNetworkTypes() bool`

HasCloudNetworkTypes returns a boolean if a field has been set.

### GetVCpuType

`func (o *K8sNodeConfiguration) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *K8sNodeConfiguration) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *K8sNodeConfiguration) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.

### HasVCpuType

`func (o *K8sNodeConfiguration) HasVCpuType() bool`

HasVCpuType returns a boolean if a field has been set.

### GetVCpu

`func (o *K8sNodeConfiguration) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *K8sNodeConfiguration) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *K8sNodeConfiguration) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *K8sNodeConfiguration) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetRamGb

`func (o *K8sNodeConfiguration) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *K8sNodeConfiguration) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *K8sNodeConfiguration) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *K8sNodeConfiguration) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetVolumeGb

`func (o *K8sNodeConfiguration) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *K8sNodeConfiguration) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *K8sNodeConfiguration) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.

### HasVolumeGb

`func (o *K8sNodeConfiguration) HasVolumeGb() bool`

HasVolumeGb returns a boolean if a field has been set.

### GetVolumeType

`func (o *K8sNodeConfiguration) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *K8sNodeConfiguration) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *K8sNodeConfiguration) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *K8sNodeConfiguration) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetAccelerator

`func (o *K8sNodeConfiguration) GetAccelerator() K8sNodeConfigurationAccelerator`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *K8sNodeConfiguration) GetAcceleratorOk() (*K8sNodeConfigurationAccelerator, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *K8sNodeConfiguration) SetAccelerator(v K8sNodeConfigurationAccelerator)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *K8sNodeConfiguration) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetLocalDisks

`func (o *K8sNodeConfiguration) GetLocalDisks() K8sNodeConfigurationLocalDisks`

GetLocalDisks returns the LocalDisks field if non-nil, zero value otherwise.

### GetLocalDisksOk

`func (o *K8sNodeConfiguration) GetLocalDisksOk() (*K8sNodeConfigurationLocalDisks, bool)`

GetLocalDisksOk returns a tuple with the LocalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDisks

`func (o *K8sNodeConfiguration) SetLocalDisks(v K8sNodeConfigurationLocalDisks)`

SetLocalDisks sets LocalDisks field to given value.

### HasLocalDisks

`func (o *K8sNodeConfiguration) HasLocalDisks() bool`

HasLocalDisks returns a boolean if a field has been set.

### GetCost

`func (o *K8sNodeConfiguration) GetCost() K8sNodeConfigurationCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *K8sNodeConfiguration) GetCostOk() (*K8sNodeConfigurationCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *K8sNodeConfiguration) SetCost(v K8sNodeConfigurationCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *K8sNodeConfiguration) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


