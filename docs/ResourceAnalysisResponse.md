# ResourceAnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | emma&#39;s internal service type | [optional] 
**Timecode** | Pointer to **string** |  | [optional] 
**CpuCoresNumber** | Pointer to **float32** | Total CPU per service, vCPUs | [optional] 
**RamTotalAmountGb** | Pointer to **float32** | Total memory per service, GB | [optional] 
**DiskSpaceTotalGb** | Pointer to **float32** | Total disk size per service, GB | [optional] 
**Type** | Pointer to **string** | Dataset type | [optional] 

## Methods

### NewResourceAnalysisResponse

`func NewResourceAnalysisResponse() *ResourceAnalysisResponse`

NewResourceAnalysisResponse instantiates a new ResourceAnalysisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAnalysisResponseWithDefaults

`func NewResourceAnalysisResponseWithDefaults() *ResourceAnalysisResponse`

NewResourceAnalysisResponseWithDefaults instantiates a new ResourceAnalysisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ResourceAnalysisResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ResourceAnalysisResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ResourceAnalysisResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ResourceAnalysisResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTimecode

`func (o *ResourceAnalysisResponse) GetTimecode() string`

GetTimecode returns the Timecode field if non-nil, zero value otherwise.

### GetTimecodeOk

`func (o *ResourceAnalysisResponse) GetTimecodeOk() (*string, bool)`

GetTimecodeOk returns a tuple with the Timecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecode

`func (o *ResourceAnalysisResponse) SetTimecode(v string)`

SetTimecode sets Timecode field to given value.

### HasTimecode

`func (o *ResourceAnalysisResponse) HasTimecode() bool`

HasTimecode returns a boolean if a field has been set.

### GetCpuCoresNumber

`func (o *ResourceAnalysisResponse) GetCpuCoresNumber() float32`

GetCpuCoresNumber returns the CpuCoresNumber field if non-nil, zero value otherwise.

### GetCpuCoresNumberOk

`func (o *ResourceAnalysisResponse) GetCpuCoresNumberOk() (*float32, bool)`

GetCpuCoresNumberOk returns a tuple with the CpuCoresNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoresNumber

`func (o *ResourceAnalysisResponse) SetCpuCoresNumber(v float32)`

SetCpuCoresNumber sets CpuCoresNumber field to given value.

### HasCpuCoresNumber

`func (o *ResourceAnalysisResponse) HasCpuCoresNumber() bool`

HasCpuCoresNumber returns a boolean if a field has been set.

### GetRamTotalAmountGb

`func (o *ResourceAnalysisResponse) GetRamTotalAmountGb() float32`

GetRamTotalAmountGb returns the RamTotalAmountGb field if non-nil, zero value otherwise.

### GetRamTotalAmountGbOk

`func (o *ResourceAnalysisResponse) GetRamTotalAmountGbOk() (*float32, bool)`

GetRamTotalAmountGbOk returns a tuple with the RamTotalAmountGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamTotalAmountGb

`func (o *ResourceAnalysisResponse) SetRamTotalAmountGb(v float32)`

SetRamTotalAmountGb sets RamTotalAmountGb field to given value.

### HasRamTotalAmountGb

`func (o *ResourceAnalysisResponse) HasRamTotalAmountGb() bool`

HasRamTotalAmountGb returns a boolean if a field has been set.

### GetDiskSpaceTotalGb

`func (o *ResourceAnalysisResponse) GetDiskSpaceTotalGb() float32`

GetDiskSpaceTotalGb returns the DiskSpaceTotalGb field if non-nil, zero value otherwise.

### GetDiskSpaceTotalGbOk

`func (o *ResourceAnalysisResponse) GetDiskSpaceTotalGbOk() (*float32, bool)`

GetDiskSpaceTotalGbOk returns a tuple with the DiskSpaceTotalGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceTotalGb

`func (o *ResourceAnalysisResponse) SetDiskSpaceTotalGb(v float32)`

SetDiskSpaceTotalGb sets DiskSpaceTotalGb field to given value.

### HasDiskSpaceTotalGb

`func (o *ResourceAnalysisResponse) HasDiskSpaceTotalGb() bool`

HasDiskSpaceTotalGb returns a boolean if a field has been set.

### GetType

`func (o *ResourceAnalysisResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceAnalysisResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceAnalysisResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceAnalysisResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


