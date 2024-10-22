# ProductStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | emma&#39;s internal service type | [optional] 
**VmId** | Pointer to **int32** | ID of VM | [optional] 
**VmName** | Pointer to **string** | VM name | [optional] 
**HeadProductId** | Pointer to **int32** | ID of a head product | [optional] 
**HeadProductName** | Pointer to **string** | Head product name | [optional] 
**Currency** | Pointer to **string** | User currency | [optional] 
**Cost** | Pointer to **float32** | Product monthly cost | [optional] 
**ProviderName** | Pointer to **string** | Provider name | [optional] 
**Country** | Pointer to **string** | Location country | [optional] 
**Location** | Pointer to **string** | Location name | [optional] 
**Latitude** | Pointer to **float32** | Location latitude | [optional] 
**Longitude** | Pointer to **float32** | Location longitude | [optional] 
**StatusNormalized** | Pointer to **string** | Product UI status | [optional] 
**CpuTotal** | Pointer to **float32** | Total CPU, vCPUs | [optional] 
**RamTotal** | Pointer to **float32** | Total memory, MB | [optional] 
**DiskUsageTotal** | Pointer to **float32** | Total disk size, GB | [optional] 
**CpuUsage** | Pointer to **float32** | Average CPU utilization for the last hour | [optional] 
**RamUsage** | Pointer to **float32** | Average memory utilization for the last hour | [optional] 
**DiskUsage** | Pointer to **float32** | Average disk utilization for the last hour | [optional] 
**EmptyValue** | Pointer to **int32** | Internal service parameter | [optional] 

## Methods

### NewProductStatisticsResponse

`func NewProductStatisticsResponse() *ProductStatisticsResponse`

NewProductStatisticsResponse instantiates a new ProductStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductStatisticsResponseWithDefaults

`func NewProductStatisticsResponseWithDefaults() *ProductStatisticsResponse`

NewProductStatisticsResponseWithDefaults instantiates a new ProductStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ProductStatisticsResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProductStatisticsResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProductStatisticsResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ProductStatisticsResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetVmId

`func (o *ProductStatisticsResponse) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ProductStatisticsResponse) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ProductStatisticsResponse) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *ProductStatisticsResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetVmName

`func (o *ProductStatisticsResponse) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *ProductStatisticsResponse) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *ProductStatisticsResponse) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *ProductStatisticsResponse) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### GetHeadProductId

`func (o *ProductStatisticsResponse) GetHeadProductId() int32`

GetHeadProductId returns the HeadProductId field if non-nil, zero value otherwise.

### GetHeadProductIdOk

`func (o *ProductStatisticsResponse) GetHeadProductIdOk() (*int32, bool)`

GetHeadProductIdOk returns a tuple with the HeadProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadProductId

`func (o *ProductStatisticsResponse) SetHeadProductId(v int32)`

SetHeadProductId sets HeadProductId field to given value.

### HasHeadProductId

`func (o *ProductStatisticsResponse) HasHeadProductId() bool`

HasHeadProductId returns a boolean if a field has been set.

### GetHeadProductName

`func (o *ProductStatisticsResponse) GetHeadProductName() string`

GetHeadProductName returns the HeadProductName field if non-nil, zero value otherwise.

### GetHeadProductNameOk

`func (o *ProductStatisticsResponse) GetHeadProductNameOk() (*string, bool)`

GetHeadProductNameOk returns a tuple with the HeadProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadProductName

`func (o *ProductStatisticsResponse) SetHeadProductName(v string)`

SetHeadProductName sets HeadProductName field to given value.

### HasHeadProductName

`func (o *ProductStatisticsResponse) HasHeadProductName() bool`

HasHeadProductName returns a boolean if a field has been set.

### GetCurrency

`func (o *ProductStatisticsResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductStatisticsResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductStatisticsResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductStatisticsResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCost

`func (o *ProductStatisticsResponse) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProductStatisticsResponse) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProductStatisticsResponse) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProductStatisticsResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetProviderName

`func (o *ProductStatisticsResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProductStatisticsResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProductStatisticsResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProductStatisticsResponse) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetCountry

`func (o *ProductStatisticsResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ProductStatisticsResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ProductStatisticsResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ProductStatisticsResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLocation

`func (o *ProductStatisticsResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProductStatisticsResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProductStatisticsResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProductStatisticsResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLatitude

`func (o *ProductStatisticsResponse) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ProductStatisticsResponse) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ProductStatisticsResponse) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ProductStatisticsResponse) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *ProductStatisticsResponse) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ProductStatisticsResponse) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ProductStatisticsResponse) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ProductStatisticsResponse) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetStatusNormalized

`func (o *ProductStatisticsResponse) GetStatusNormalized() string`

GetStatusNormalized returns the StatusNormalized field if non-nil, zero value otherwise.

### GetStatusNormalizedOk

`func (o *ProductStatisticsResponse) GetStatusNormalizedOk() (*string, bool)`

GetStatusNormalizedOk returns a tuple with the StatusNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNormalized

`func (o *ProductStatisticsResponse) SetStatusNormalized(v string)`

SetStatusNormalized sets StatusNormalized field to given value.

### HasStatusNormalized

`func (o *ProductStatisticsResponse) HasStatusNormalized() bool`

HasStatusNormalized returns a boolean if a field has been set.

### GetCpuTotal

`func (o *ProductStatisticsResponse) GetCpuTotal() float32`

GetCpuTotal returns the CpuTotal field if non-nil, zero value otherwise.

### GetCpuTotalOk

`func (o *ProductStatisticsResponse) GetCpuTotalOk() (*float32, bool)`

GetCpuTotalOk returns a tuple with the CpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotal

`func (o *ProductStatisticsResponse) SetCpuTotal(v float32)`

SetCpuTotal sets CpuTotal field to given value.

### HasCpuTotal

`func (o *ProductStatisticsResponse) HasCpuTotal() bool`

HasCpuTotal returns a boolean if a field has been set.

### GetRamTotal

`func (o *ProductStatisticsResponse) GetRamTotal() float32`

GetRamTotal returns the RamTotal field if non-nil, zero value otherwise.

### GetRamTotalOk

`func (o *ProductStatisticsResponse) GetRamTotalOk() (*float32, bool)`

GetRamTotalOk returns a tuple with the RamTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamTotal

`func (o *ProductStatisticsResponse) SetRamTotal(v float32)`

SetRamTotal sets RamTotal field to given value.

### HasRamTotal

`func (o *ProductStatisticsResponse) HasRamTotal() bool`

HasRamTotal returns a boolean if a field has been set.

### GetDiskUsageTotal

`func (o *ProductStatisticsResponse) GetDiskUsageTotal() float32`

GetDiskUsageTotal returns the DiskUsageTotal field if non-nil, zero value otherwise.

### GetDiskUsageTotalOk

`func (o *ProductStatisticsResponse) GetDiskUsageTotalOk() (*float32, bool)`

GetDiskUsageTotalOk returns a tuple with the DiskUsageTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsageTotal

`func (o *ProductStatisticsResponse) SetDiskUsageTotal(v float32)`

SetDiskUsageTotal sets DiskUsageTotal field to given value.

### HasDiskUsageTotal

`func (o *ProductStatisticsResponse) HasDiskUsageTotal() bool`

HasDiskUsageTotal returns a boolean if a field has been set.

### GetCpuUsage

`func (o *ProductStatisticsResponse) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ProductStatisticsResponse) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ProductStatisticsResponse) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ProductStatisticsResponse) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetRamUsage

`func (o *ProductStatisticsResponse) GetRamUsage() float32`

GetRamUsage returns the RamUsage field if non-nil, zero value otherwise.

### GetRamUsageOk

`func (o *ProductStatisticsResponse) GetRamUsageOk() (*float32, bool)`

GetRamUsageOk returns a tuple with the RamUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamUsage

`func (o *ProductStatisticsResponse) SetRamUsage(v float32)`

SetRamUsage sets RamUsage field to given value.

### HasRamUsage

`func (o *ProductStatisticsResponse) HasRamUsage() bool`

HasRamUsage returns a boolean if a field has been set.

### GetDiskUsage

`func (o *ProductStatisticsResponse) GetDiskUsage() float32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *ProductStatisticsResponse) GetDiskUsageOk() (*float32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *ProductStatisticsResponse) SetDiskUsage(v float32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *ProductStatisticsResponse) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### GetEmptyValue

`func (o *ProductStatisticsResponse) GetEmptyValue() int32`

GetEmptyValue returns the EmptyValue field if non-nil, zero value otherwise.

### GetEmptyValueOk

`func (o *ProductStatisticsResponse) GetEmptyValueOk() (*int32, bool)`

GetEmptyValueOk returns a tuple with the EmptyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyValue

`func (o *ProductStatisticsResponse) SetEmptyValue(v int32)`

SetEmptyValue sets EmptyValue field to given value.

### HasEmptyValue

`func (o *ProductStatisticsResponse) HasEmptyValue() bool`

HasEmptyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


