# VmMonitoringResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timecode** | Pointer to **string** |  | [optional] 
**CpuDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**CpuUtilizationAverageCores** | Pointer to **float32** | CPU utilization with values in range [0, 100*vCPUs] | [optional] 
**AvgCpuUtilizationAverageCores** | Pointer to **float32** | Average CPU utilization for the requested period of statistics | [optional] 
**MaxCpuUtilizationAverageCores** | Pointer to **float32** | Maximum CPU utilization for the requested period of statistics | [optional] 
**CpuTotalPercent** | Pointer to **int32** | Total CPU, % | [optional] 
**CpuHumanLabel** | Pointer to **string** | Label | [optional] 
**RamDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**RamUsageAverageGb** | Pointer to **float32** | Memory utilization | [optional] 
**AvgRamUsageAverageGb** | Pointer to **float32** | Average memory utilization for the requested period of statistics | [optional] 
**MaxRamUsageAverageGb** | Pointer to **float32** | Maximum memory utilization for the requested period of statistics | [optional] 
**RamTotalAmountGb** | Pointer to **float32** | Total memory, GB | [optional] 
**RamHumanLabel** | Pointer to **string** | Label | [optional] 
**DiskUsedDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**DiskSpaceUsedGb** | Pointer to **float32** | Disk utilization | [optional] 
**AvgDiskSpaceUsedGb** | Pointer to **float32** | Average disk utilization for the requested period of statistics | [optional] 
**MaxDiskSpaceUsedGb** | Pointer to **float32** | Maximum disk utilization for the requested period of statistics | [optional] 
**DiskSpaceTotalGb** | Pointer to **int32** | Total disk size, GB | [optional] 
**DiskSpaceHumanLabel** | Pointer to **string** | Label | [optional] 
**DiskWriteDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**DiskWriteCoef** | Pointer to **float32** | Internal service parameter | [optional] 
**DiskWriteHuman** | Pointer to **float32** | Disk write | [optional] 
**AvgDiskWriteHuman** | Pointer to **float32** | Average disk write for the requested period of statistics | [optional] 
**MaxDiskWriteHuman** | Pointer to **float32** | Maximum disk write for the requested period of statistics | [optional] 
**DiskWriteHumanLabel** | Pointer to **string** | Label | [optional] 
**DiskReadDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**DiskReadCoef** | Pointer to **float32** | Internal service parameter | [optional] 
**DiskReadHuman** | Pointer to **float32** | Disk read | [optional] 
**AvgDiskReadHuman** | Pointer to **float32** | Average disk read for the requested period of statistics | [optional] 
**MaxDiskReadHuman** | Pointer to **float32** | Maximum disk read for the requested period of statistics | [optional] 
**DiskReadHumanLabel** | Pointer to **string** | Label | [optional] 
**NetworkOutDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**NetworkOutCoef** | Pointer to **float32** | Internal service parameter | [optional] 
**NetworkOutHuman** | Pointer to **float32** | Network out | [optional] 
**AvgNetworkOutHuman** | Pointer to **float32** | Average network out for the requested period of statistics | [optional] 
**MaxNetworkOutHuman** | Pointer to **float32** | Maximum network out for the requested period of statistics | [optional] 
**NetworkOutHumanLabel** | Pointer to **string** | Label | [optional] 
**NetworkInDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**NetworkInCoef** | Pointer to **float32** | Internal service parameter | [optional] 
**NetworkInHuman** | Pointer to **float32** | Network in | [optional] 
**AvgNetworkInHuman** | Pointer to **float32** | Average network in for the requested period of statistics | [optional] 
**MaxNetworkInHuman** | Pointer to **float32** | Maximum network in for the requested period of statistics | [optional] 
**NetworkInHumanLabel** | Pointer to **string** | Label | [optional] 

## Methods

### NewVmMonitoringResponse

`func NewVmMonitoringResponse() *VmMonitoringResponse`

NewVmMonitoringResponse instantiates a new VmMonitoringResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmMonitoringResponseWithDefaults

`func NewVmMonitoringResponseWithDefaults() *VmMonitoringResponse`

NewVmMonitoringResponseWithDefaults instantiates a new VmMonitoringResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimecode

`func (o *VmMonitoringResponse) GetTimecode() string`

GetTimecode returns the Timecode field if non-nil, zero value otherwise.

### GetTimecodeOk

`func (o *VmMonitoringResponse) GetTimecodeOk() (*string, bool)`

GetTimecodeOk returns a tuple with the Timecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecode

`func (o *VmMonitoringResponse) SetTimecode(v string)`

SetTimecode sets Timecode field to given value.

### HasTimecode

`func (o *VmMonitoringResponse) HasTimecode() bool`

HasTimecode returns a boolean if a field has been set.

### GetCpuDataPresent

`func (o *VmMonitoringResponse) GetCpuDataPresent() int32`

GetCpuDataPresent returns the CpuDataPresent field if non-nil, zero value otherwise.

### GetCpuDataPresentOk

`func (o *VmMonitoringResponse) GetCpuDataPresentOk() (*int32, bool)`

GetCpuDataPresentOk returns a tuple with the CpuDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuDataPresent

`func (o *VmMonitoringResponse) SetCpuDataPresent(v int32)`

SetCpuDataPresent sets CpuDataPresent field to given value.

### HasCpuDataPresent

`func (o *VmMonitoringResponse) HasCpuDataPresent() bool`

HasCpuDataPresent returns a boolean if a field has been set.

### GetCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) GetCpuUtilizationAverageCores() float32`

GetCpuUtilizationAverageCores returns the CpuUtilizationAverageCores field if non-nil, zero value otherwise.

### GetCpuUtilizationAverageCoresOk

`func (o *VmMonitoringResponse) GetCpuUtilizationAverageCoresOk() (*float32, bool)`

GetCpuUtilizationAverageCoresOk returns a tuple with the CpuUtilizationAverageCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) SetCpuUtilizationAverageCores(v float32)`

SetCpuUtilizationAverageCores sets CpuUtilizationAverageCores field to given value.

### HasCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) HasCpuUtilizationAverageCores() bool`

HasCpuUtilizationAverageCores returns a boolean if a field has been set.

### GetAvgCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) GetAvgCpuUtilizationAverageCores() float32`

GetAvgCpuUtilizationAverageCores returns the AvgCpuUtilizationAverageCores field if non-nil, zero value otherwise.

### GetAvgCpuUtilizationAverageCoresOk

`func (o *VmMonitoringResponse) GetAvgCpuUtilizationAverageCoresOk() (*float32, bool)`

GetAvgCpuUtilizationAverageCoresOk returns a tuple with the AvgCpuUtilizationAverageCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) SetAvgCpuUtilizationAverageCores(v float32)`

SetAvgCpuUtilizationAverageCores sets AvgCpuUtilizationAverageCores field to given value.

### HasAvgCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) HasAvgCpuUtilizationAverageCores() bool`

HasAvgCpuUtilizationAverageCores returns a boolean if a field has been set.

### GetMaxCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) GetMaxCpuUtilizationAverageCores() float32`

GetMaxCpuUtilizationAverageCores returns the MaxCpuUtilizationAverageCores field if non-nil, zero value otherwise.

### GetMaxCpuUtilizationAverageCoresOk

`func (o *VmMonitoringResponse) GetMaxCpuUtilizationAverageCoresOk() (*float32, bool)`

GetMaxCpuUtilizationAverageCoresOk returns a tuple with the MaxCpuUtilizationAverageCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) SetMaxCpuUtilizationAverageCores(v float32)`

SetMaxCpuUtilizationAverageCores sets MaxCpuUtilizationAverageCores field to given value.

### HasMaxCpuUtilizationAverageCores

`func (o *VmMonitoringResponse) HasMaxCpuUtilizationAverageCores() bool`

HasMaxCpuUtilizationAverageCores returns a boolean if a field has been set.

### GetCpuTotalPercent

`func (o *VmMonitoringResponse) GetCpuTotalPercent() int32`

GetCpuTotalPercent returns the CpuTotalPercent field if non-nil, zero value otherwise.

### GetCpuTotalPercentOk

`func (o *VmMonitoringResponse) GetCpuTotalPercentOk() (*int32, bool)`

GetCpuTotalPercentOk returns a tuple with the CpuTotalPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalPercent

`func (o *VmMonitoringResponse) SetCpuTotalPercent(v int32)`

SetCpuTotalPercent sets CpuTotalPercent field to given value.

### HasCpuTotalPercent

`func (o *VmMonitoringResponse) HasCpuTotalPercent() bool`

HasCpuTotalPercent returns a boolean if a field has been set.

### GetCpuHumanLabel

`func (o *VmMonitoringResponse) GetCpuHumanLabel() string`

GetCpuHumanLabel returns the CpuHumanLabel field if non-nil, zero value otherwise.

### GetCpuHumanLabelOk

`func (o *VmMonitoringResponse) GetCpuHumanLabelOk() (*string, bool)`

GetCpuHumanLabelOk returns a tuple with the CpuHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHumanLabel

`func (o *VmMonitoringResponse) SetCpuHumanLabel(v string)`

SetCpuHumanLabel sets CpuHumanLabel field to given value.

### HasCpuHumanLabel

`func (o *VmMonitoringResponse) HasCpuHumanLabel() bool`

HasCpuHumanLabel returns a boolean if a field has been set.

### GetRamDataPresent

`func (o *VmMonitoringResponse) GetRamDataPresent() int32`

GetRamDataPresent returns the RamDataPresent field if non-nil, zero value otherwise.

### GetRamDataPresentOk

`func (o *VmMonitoringResponse) GetRamDataPresentOk() (*int32, bool)`

GetRamDataPresentOk returns a tuple with the RamDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamDataPresent

`func (o *VmMonitoringResponse) SetRamDataPresent(v int32)`

SetRamDataPresent sets RamDataPresent field to given value.

### HasRamDataPresent

`func (o *VmMonitoringResponse) HasRamDataPresent() bool`

HasRamDataPresent returns a boolean if a field has been set.

### GetRamUsageAverageGb

`func (o *VmMonitoringResponse) GetRamUsageAverageGb() float32`

GetRamUsageAverageGb returns the RamUsageAverageGb field if non-nil, zero value otherwise.

### GetRamUsageAverageGbOk

`func (o *VmMonitoringResponse) GetRamUsageAverageGbOk() (*float32, bool)`

GetRamUsageAverageGbOk returns a tuple with the RamUsageAverageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamUsageAverageGb

`func (o *VmMonitoringResponse) SetRamUsageAverageGb(v float32)`

SetRamUsageAverageGb sets RamUsageAverageGb field to given value.

### HasRamUsageAverageGb

`func (o *VmMonitoringResponse) HasRamUsageAverageGb() bool`

HasRamUsageAverageGb returns a boolean if a field has been set.

### GetAvgRamUsageAverageGb

`func (o *VmMonitoringResponse) GetAvgRamUsageAverageGb() float32`

GetAvgRamUsageAverageGb returns the AvgRamUsageAverageGb field if non-nil, zero value otherwise.

### GetAvgRamUsageAverageGbOk

`func (o *VmMonitoringResponse) GetAvgRamUsageAverageGbOk() (*float32, bool)`

GetAvgRamUsageAverageGbOk returns a tuple with the AvgRamUsageAverageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRamUsageAverageGb

`func (o *VmMonitoringResponse) SetAvgRamUsageAverageGb(v float32)`

SetAvgRamUsageAverageGb sets AvgRamUsageAverageGb field to given value.

### HasAvgRamUsageAverageGb

`func (o *VmMonitoringResponse) HasAvgRamUsageAverageGb() bool`

HasAvgRamUsageAverageGb returns a boolean if a field has been set.

### GetMaxRamUsageAverageGb

`func (o *VmMonitoringResponse) GetMaxRamUsageAverageGb() float32`

GetMaxRamUsageAverageGb returns the MaxRamUsageAverageGb field if non-nil, zero value otherwise.

### GetMaxRamUsageAverageGbOk

`func (o *VmMonitoringResponse) GetMaxRamUsageAverageGbOk() (*float32, bool)`

GetMaxRamUsageAverageGbOk returns a tuple with the MaxRamUsageAverageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamUsageAverageGb

`func (o *VmMonitoringResponse) SetMaxRamUsageAverageGb(v float32)`

SetMaxRamUsageAverageGb sets MaxRamUsageAverageGb field to given value.

### HasMaxRamUsageAverageGb

`func (o *VmMonitoringResponse) HasMaxRamUsageAverageGb() bool`

HasMaxRamUsageAverageGb returns a boolean if a field has been set.

### GetRamTotalAmountGb

`func (o *VmMonitoringResponse) GetRamTotalAmountGb() float32`

GetRamTotalAmountGb returns the RamTotalAmountGb field if non-nil, zero value otherwise.

### GetRamTotalAmountGbOk

`func (o *VmMonitoringResponse) GetRamTotalAmountGbOk() (*float32, bool)`

GetRamTotalAmountGbOk returns a tuple with the RamTotalAmountGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamTotalAmountGb

`func (o *VmMonitoringResponse) SetRamTotalAmountGb(v float32)`

SetRamTotalAmountGb sets RamTotalAmountGb field to given value.

### HasRamTotalAmountGb

`func (o *VmMonitoringResponse) HasRamTotalAmountGb() bool`

HasRamTotalAmountGb returns a boolean if a field has been set.

### GetRamHumanLabel

`func (o *VmMonitoringResponse) GetRamHumanLabel() string`

GetRamHumanLabel returns the RamHumanLabel field if non-nil, zero value otherwise.

### GetRamHumanLabelOk

`func (o *VmMonitoringResponse) GetRamHumanLabelOk() (*string, bool)`

GetRamHumanLabelOk returns a tuple with the RamHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamHumanLabel

`func (o *VmMonitoringResponse) SetRamHumanLabel(v string)`

SetRamHumanLabel sets RamHumanLabel field to given value.

### HasRamHumanLabel

`func (o *VmMonitoringResponse) HasRamHumanLabel() bool`

HasRamHumanLabel returns a boolean if a field has been set.

### GetDiskUsedDataPresent

`func (o *VmMonitoringResponse) GetDiskUsedDataPresent() int32`

GetDiskUsedDataPresent returns the DiskUsedDataPresent field if non-nil, zero value otherwise.

### GetDiskUsedDataPresentOk

`func (o *VmMonitoringResponse) GetDiskUsedDataPresentOk() (*int32, bool)`

GetDiskUsedDataPresentOk returns a tuple with the DiskUsedDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsedDataPresent

`func (o *VmMonitoringResponse) SetDiskUsedDataPresent(v int32)`

SetDiskUsedDataPresent sets DiskUsedDataPresent field to given value.

### HasDiskUsedDataPresent

`func (o *VmMonitoringResponse) HasDiskUsedDataPresent() bool`

HasDiskUsedDataPresent returns a boolean if a field has been set.

### GetDiskSpaceUsedGb

`func (o *VmMonitoringResponse) GetDiskSpaceUsedGb() float32`

GetDiskSpaceUsedGb returns the DiskSpaceUsedGb field if non-nil, zero value otherwise.

### GetDiskSpaceUsedGbOk

`func (o *VmMonitoringResponse) GetDiskSpaceUsedGbOk() (*float32, bool)`

GetDiskSpaceUsedGbOk returns a tuple with the DiskSpaceUsedGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsedGb

`func (o *VmMonitoringResponse) SetDiskSpaceUsedGb(v float32)`

SetDiskSpaceUsedGb sets DiskSpaceUsedGb field to given value.

### HasDiskSpaceUsedGb

`func (o *VmMonitoringResponse) HasDiskSpaceUsedGb() bool`

HasDiskSpaceUsedGb returns a boolean if a field has been set.

### GetAvgDiskSpaceUsedGb

`func (o *VmMonitoringResponse) GetAvgDiskSpaceUsedGb() float32`

GetAvgDiskSpaceUsedGb returns the AvgDiskSpaceUsedGb field if non-nil, zero value otherwise.

### GetAvgDiskSpaceUsedGbOk

`func (o *VmMonitoringResponse) GetAvgDiskSpaceUsedGbOk() (*float32, bool)`

GetAvgDiskSpaceUsedGbOk returns a tuple with the AvgDiskSpaceUsedGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDiskSpaceUsedGb

`func (o *VmMonitoringResponse) SetAvgDiskSpaceUsedGb(v float32)`

SetAvgDiskSpaceUsedGb sets AvgDiskSpaceUsedGb field to given value.

### HasAvgDiskSpaceUsedGb

`func (o *VmMonitoringResponse) HasAvgDiskSpaceUsedGb() bool`

HasAvgDiskSpaceUsedGb returns a boolean if a field has been set.

### GetMaxDiskSpaceUsedGb

`func (o *VmMonitoringResponse) GetMaxDiskSpaceUsedGb() float32`

GetMaxDiskSpaceUsedGb returns the MaxDiskSpaceUsedGb field if non-nil, zero value otherwise.

### GetMaxDiskSpaceUsedGbOk

`func (o *VmMonitoringResponse) GetMaxDiskSpaceUsedGbOk() (*float32, bool)`

GetMaxDiskSpaceUsedGbOk returns a tuple with the MaxDiskSpaceUsedGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiskSpaceUsedGb

`func (o *VmMonitoringResponse) SetMaxDiskSpaceUsedGb(v float32)`

SetMaxDiskSpaceUsedGb sets MaxDiskSpaceUsedGb field to given value.

### HasMaxDiskSpaceUsedGb

`func (o *VmMonitoringResponse) HasMaxDiskSpaceUsedGb() bool`

HasMaxDiskSpaceUsedGb returns a boolean if a field has been set.

### GetDiskSpaceTotalGb

`func (o *VmMonitoringResponse) GetDiskSpaceTotalGb() int32`

GetDiskSpaceTotalGb returns the DiskSpaceTotalGb field if non-nil, zero value otherwise.

### GetDiskSpaceTotalGbOk

`func (o *VmMonitoringResponse) GetDiskSpaceTotalGbOk() (*int32, bool)`

GetDiskSpaceTotalGbOk returns a tuple with the DiskSpaceTotalGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceTotalGb

`func (o *VmMonitoringResponse) SetDiskSpaceTotalGb(v int32)`

SetDiskSpaceTotalGb sets DiskSpaceTotalGb field to given value.

### HasDiskSpaceTotalGb

`func (o *VmMonitoringResponse) HasDiskSpaceTotalGb() bool`

HasDiskSpaceTotalGb returns a boolean if a field has been set.

### GetDiskSpaceHumanLabel

`func (o *VmMonitoringResponse) GetDiskSpaceHumanLabel() string`

GetDiskSpaceHumanLabel returns the DiskSpaceHumanLabel field if non-nil, zero value otherwise.

### GetDiskSpaceHumanLabelOk

`func (o *VmMonitoringResponse) GetDiskSpaceHumanLabelOk() (*string, bool)`

GetDiskSpaceHumanLabelOk returns a tuple with the DiskSpaceHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceHumanLabel

`func (o *VmMonitoringResponse) SetDiskSpaceHumanLabel(v string)`

SetDiskSpaceHumanLabel sets DiskSpaceHumanLabel field to given value.

### HasDiskSpaceHumanLabel

`func (o *VmMonitoringResponse) HasDiskSpaceHumanLabel() bool`

HasDiskSpaceHumanLabel returns a boolean if a field has been set.

### GetDiskWriteDataPresent

`func (o *VmMonitoringResponse) GetDiskWriteDataPresent() int32`

GetDiskWriteDataPresent returns the DiskWriteDataPresent field if non-nil, zero value otherwise.

### GetDiskWriteDataPresentOk

`func (o *VmMonitoringResponse) GetDiskWriteDataPresentOk() (*int32, bool)`

GetDiskWriteDataPresentOk returns a tuple with the DiskWriteDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteDataPresent

`func (o *VmMonitoringResponse) SetDiskWriteDataPresent(v int32)`

SetDiskWriteDataPresent sets DiskWriteDataPresent field to given value.

### HasDiskWriteDataPresent

`func (o *VmMonitoringResponse) HasDiskWriteDataPresent() bool`

HasDiskWriteDataPresent returns a boolean if a field has been set.

### GetDiskWriteCoef

`func (o *VmMonitoringResponse) GetDiskWriteCoef() float32`

GetDiskWriteCoef returns the DiskWriteCoef field if non-nil, zero value otherwise.

### GetDiskWriteCoefOk

`func (o *VmMonitoringResponse) GetDiskWriteCoefOk() (*float32, bool)`

GetDiskWriteCoefOk returns a tuple with the DiskWriteCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteCoef

`func (o *VmMonitoringResponse) SetDiskWriteCoef(v float32)`

SetDiskWriteCoef sets DiskWriteCoef field to given value.

### HasDiskWriteCoef

`func (o *VmMonitoringResponse) HasDiskWriteCoef() bool`

HasDiskWriteCoef returns a boolean if a field has been set.

### GetDiskWriteHuman

`func (o *VmMonitoringResponse) GetDiskWriteHuman() float32`

GetDiskWriteHuman returns the DiskWriteHuman field if non-nil, zero value otherwise.

### GetDiskWriteHumanOk

`func (o *VmMonitoringResponse) GetDiskWriteHumanOk() (*float32, bool)`

GetDiskWriteHumanOk returns a tuple with the DiskWriteHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteHuman

`func (o *VmMonitoringResponse) SetDiskWriteHuman(v float32)`

SetDiskWriteHuman sets DiskWriteHuman field to given value.

### HasDiskWriteHuman

`func (o *VmMonitoringResponse) HasDiskWriteHuman() bool`

HasDiskWriteHuman returns a boolean if a field has been set.

### GetAvgDiskWriteHuman

`func (o *VmMonitoringResponse) GetAvgDiskWriteHuman() float32`

GetAvgDiskWriteHuman returns the AvgDiskWriteHuman field if non-nil, zero value otherwise.

### GetAvgDiskWriteHumanOk

`func (o *VmMonitoringResponse) GetAvgDiskWriteHumanOk() (*float32, bool)`

GetAvgDiskWriteHumanOk returns a tuple with the AvgDiskWriteHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDiskWriteHuman

`func (o *VmMonitoringResponse) SetAvgDiskWriteHuman(v float32)`

SetAvgDiskWriteHuman sets AvgDiskWriteHuman field to given value.

### HasAvgDiskWriteHuman

`func (o *VmMonitoringResponse) HasAvgDiskWriteHuman() bool`

HasAvgDiskWriteHuman returns a boolean if a field has been set.

### GetMaxDiskWriteHuman

`func (o *VmMonitoringResponse) GetMaxDiskWriteHuman() float32`

GetMaxDiskWriteHuman returns the MaxDiskWriteHuman field if non-nil, zero value otherwise.

### GetMaxDiskWriteHumanOk

`func (o *VmMonitoringResponse) GetMaxDiskWriteHumanOk() (*float32, bool)`

GetMaxDiskWriteHumanOk returns a tuple with the MaxDiskWriteHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiskWriteHuman

`func (o *VmMonitoringResponse) SetMaxDiskWriteHuman(v float32)`

SetMaxDiskWriteHuman sets MaxDiskWriteHuman field to given value.

### HasMaxDiskWriteHuman

`func (o *VmMonitoringResponse) HasMaxDiskWriteHuman() bool`

HasMaxDiskWriteHuman returns a boolean if a field has been set.

### GetDiskWriteHumanLabel

`func (o *VmMonitoringResponse) GetDiskWriteHumanLabel() string`

GetDiskWriteHumanLabel returns the DiskWriteHumanLabel field if non-nil, zero value otherwise.

### GetDiskWriteHumanLabelOk

`func (o *VmMonitoringResponse) GetDiskWriteHumanLabelOk() (*string, bool)`

GetDiskWriteHumanLabelOk returns a tuple with the DiskWriteHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteHumanLabel

`func (o *VmMonitoringResponse) SetDiskWriteHumanLabel(v string)`

SetDiskWriteHumanLabel sets DiskWriteHumanLabel field to given value.

### HasDiskWriteHumanLabel

`func (o *VmMonitoringResponse) HasDiskWriteHumanLabel() bool`

HasDiskWriteHumanLabel returns a boolean if a field has been set.

### GetDiskReadDataPresent

`func (o *VmMonitoringResponse) GetDiskReadDataPresent() int32`

GetDiskReadDataPresent returns the DiskReadDataPresent field if non-nil, zero value otherwise.

### GetDiskReadDataPresentOk

`func (o *VmMonitoringResponse) GetDiskReadDataPresentOk() (*int32, bool)`

GetDiskReadDataPresentOk returns a tuple with the DiskReadDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadDataPresent

`func (o *VmMonitoringResponse) SetDiskReadDataPresent(v int32)`

SetDiskReadDataPresent sets DiskReadDataPresent field to given value.

### HasDiskReadDataPresent

`func (o *VmMonitoringResponse) HasDiskReadDataPresent() bool`

HasDiskReadDataPresent returns a boolean if a field has been set.

### GetDiskReadCoef

`func (o *VmMonitoringResponse) GetDiskReadCoef() float32`

GetDiskReadCoef returns the DiskReadCoef field if non-nil, zero value otherwise.

### GetDiskReadCoefOk

`func (o *VmMonitoringResponse) GetDiskReadCoefOk() (*float32, bool)`

GetDiskReadCoefOk returns a tuple with the DiskReadCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadCoef

`func (o *VmMonitoringResponse) SetDiskReadCoef(v float32)`

SetDiskReadCoef sets DiskReadCoef field to given value.

### HasDiskReadCoef

`func (o *VmMonitoringResponse) HasDiskReadCoef() bool`

HasDiskReadCoef returns a boolean if a field has been set.

### GetDiskReadHuman

`func (o *VmMonitoringResponse) GetDiskReadHuman() float32`

GetDiskReadHuman returns the DiskReadHuman field if non-nil, zero value otherwise.

### GetDiskReadHumanOk

`func (o *VmMonitoringResponse) GetDiskReadHumanOk() (*float32, bool)`

GetDiskReadHumanOk returns a tuple with the DiskReadHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadHuman

`func (o *VmMonitoringResponse) SetDiskReadHuman(v float32)`

SetDiskReadHuman sets DiskReadHuman field to given value.

### HasDiskReadHuman

`func (o *VmMonitoringResponse) HasDiskReadHuman() bool`

HasDiskReadHuman returns a boolean if a field has been set.

### GetAvgDiskReadHuman

`func (o *VmMonitoringResponse) GetAvgDiskReadHuman() float32`

GetAvgDiskReadHuman returns the AvgDiskReadHuman field if non-nil, zero value otherwise.

### GetAvgDiskReadHumanOk

`func (o *VmMonitoringResponse) GetAvgDiskReadHumanOk() (*float32, bool)`

GetAvgDiskReadHumanOk returns a tuple with the AvgDiskReadHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDiskReadHuman

`func (o *VmMonitoringResponse) SetAvgDiskReadHuman(v float32)`

SetAvgDiskReadHuman sets AvgDiskReadHuman field to given value.

### HasAvgDiskReadHuman

`func (o *VmMonitoringResponse) HasAvgDiskReadHuman() bool`

HasAvgDiskReadHuman returns a boolean if a field has been set.

### GetMaxDiskReadHuman

`func (o *VmMonitoringResponse) GetMaxDiskReadHuman() float32`

GetMaxDiskReadHuman returns the MaxDiskReadHuman field if non-nil, zero value otherwise.

### GetMaxDiskReadHumanOk

`func (o *VmMonitoringResponse) GetMaxDiskReadHumanOk() (*float32, bool)`

GetMaxDiskReadHumanOk returns a tuple with the MaxDiskReadHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiskReadHuman

`func (o *VmMonitoringResponse) SetMaxDiskReadHuman(v float32)`

SetMaxDiskReadHuman sets MaxDiskReadHuman field to given value.

### HasMaxDiskReadHuman

`func (o *VmMonitoringResponse) HasMaxDiskReadHuman() bool`

HasMaxDiskReadHuman returns a boolean if a field has been set.

### GetDiskReadHumanLabel

`func (o *VmMonitoringResponse) GetDiskReadHumanLabel() string`

GetDiskReadHumanLabel returns the DiskReadHumanLabel field if non-nil, zero value otherwise.

### GetDiskReadHumanLabelOk

`func (o *VmMonitoringResponse) GetDiskReadHumanLabelOk() (*string, bool)`

GetDiskReadHumanLabelOk returns a tuple with the DiskReadHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadHumanLabel

`func (o *VmMonitoringResponse) SetDiskReadHumanLabel(v string)`

SetDiskReadHumanLabel sets DiskReadHumanLabel field to given value.

### HasDiskReadHumanLabel

`func (o *VmMonitoringResponse) HasDiskReadHumanLabel() bool`

HasDiskReadHumanLabel returns a boolean if a field has been set.

### GetNetworkOutDataPresent

`func (o *VmMonitoringResponse) GetNetworkOutDataPresent() int32`

GetNetworkOutDataPresent returns the NetworkOutDataPresent field if non-nil, zero value otherwise.

### GetNetworkOutDataPresentOk

`func (o *VmMonitoringResponse) GetNetworkOutDataPresentOk() (*int32, bool)`

GetNetworkOutDataPresentOk returns a tuple with the NetworkOutDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutDataPresent

`func (o *VmMonitoringResponse) SetNetworkOutDataPresent(v int32)`

SetNetworkOutDataPresent sets NetworkOutDataPresent field to given value.

### HasNetworkOutDataPresent

`func (o *VmMonitoringResponse) HasNetworkOutDataPresent() bool`

HasNetworkOutDataPresent returns a boolean if a field has been set.

### GetNetworkOutCoef

`func (o *VmMonitoringResponse) GetNetworkOutCoef() float32`

GetNetworkOutCoef returns the NetworkOutCoef field if non-nil, zero value otherwise.

### GetNetworkOutCoefOk

`func (o *VmMonitoringResponse) GetNetworkOutCoefOk() (*float32, bool)`

GetNetworkOutCoefOk returns a tuple with the NetworkOutCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutCoef

`func (o *VmMonitoringResponse) SetNetworkOutCoef(v float32)`

SetNetworkOutCoef sets NetworkOutCoef field to given value.

### HasNetworkOutCoef

`func (o *VmMonitoringResponse) HasNetworkOutCoef() bool`

HasNetworkOutCoef returns a boolean if a field has been set.

### GetNetworkOutHuman

`func (o *VmMonitoringResponse) GetNetworkOutHuman() float32`

GetNetworkOutHuman returns the NetworkOutHuman field if non-nil, zero value otherwise.

### GetNetworkOutHumanOk

`func (o *VmMonitoringResponse) GetNetworkOutHumanOk() (*float32, bool)`

GetNetworkOutHumanOk returns a tuple with the NetworkOutHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutHuman

`func (o *VmMonitoringResponse) SetNetworkOutHuman(v float32)`

SetNetworkOutHuman sets NetworkOutHuman field to given value.

### HasNetworkOutHuman

`func (o *VmMonitoringResponse) HasNetworkOutHuman() bool`

HasNetworkOutHuman returns a boolean if a field has been set.

### GetAvgNetworkOutHuman

`func (o *VmMonitoringResponse) GetAvgNetworkOutHuman() float32`

GetAvgNetworkOutHuman returns the AvgNetworkOutHuman field if non-nil, zero value otherwise.

### GetAvgNetworkOutHumanOk

`func (o *VmMonitoringResponse) GetAvgNetworkOutHumanOk() (*float32, bool)`

GetAvgNetworkOutHumanOk returns a tuple with the AvgNetworkOutHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgNetworkOutHuman

`func (o *VmMonitoringResponse) SetAvgNetworkOutHuman(v float32)`

SetAvgNetworkOutHuman sets AvgNetworkOutHuman field to given value.

### HasAvgNetworkOutHuman

`func (o *VmMonitoringResponse) HasAvgNetworkOutHuman() bool`

HasAvgNetworkOutHuman returns a boolean if a field has been set.

### GetMaxNetworkOutHuman

`func (o *VmMonitoringResponse) GetMaxNetworkOutHuman() float32`

GetMaxNetworkOutHuman returns the MaxNetworkOutHuman field if non-nil, zero value otherwise.

### GetMaxNetworkOutHumanOk

`func (o *VmMonitoringResponse) GetMaxNetworkOutHumanOk() (*float32, bool)`

GetMaxNetworkOutHumanOk returns a tuple with the MaxNetworkOutHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworkOutHuman

`func (o *VmMonitoringResponse) SetMaxNetworkOutHuman(v float32)`

SetMaxNetworkOutHuman sets MaxNetworkOutHuman field to given value.

### HasMaxNetworkOutHuman

`func (o *VmMonitoringResponse) HasMaxNetworkOutHuman() bool`

HasMaxNetworkOutHuman returns a boolean if a field has been set.

### GetNetworkOutHumanLabel

`func (o *VmMonitoringResponse) GetNetworkOutHumanLabel() string`

GetNetworkOutHumanLabel returns the NetworkOutHumanLabel field if non-nil, zero value otherwise.

### GetNetworkOutHumanLabelOk

`func (o *VmMonitoringResponse) GetNetworkOutHumanLabelOk() (*string, bool)`

GetNetworkOutHumanLabelOk returns a tuple with the NetworkOutHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutHumanLabel

`func (o *VmMonitoringResponse) SetNetworkOutHumanLabel(v string)`

SetNetworkOutHumanLabel sets NetworkOutHumanLabel field to given value.

### HasNetworkOutHumanLabel

`func (o *VmMonitoringResponse) HasNetworkOutHumanLabel() bool`

HasNetworkOutHumanLabel returns a boolean if a field has been set.

### GetNetworkInDataPresent

`func (o *VmMonitoringResponse) GetNetworkInDataPresent() int32`

GetNetworkInDataPresent returns the NetworkInDataPresent field if non-nil, zero value otherwise.

### GetNetworkInDataPresentOk

`func (o *VmMonitoringResponse) GetNetworkInDataPresentOk() (*int32, bool)`

GetNetworkInDataPresentOk returns a tuple with the NetworkInDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInDataPresent

`func (o *VmMonitoringResponse) SetNetworkInDataPresent(v int32)`

SetNetworkInDataPresent sets NetworkInDataPresent field to given value.

### HasNetworkInDataPresent

`func (o *VmMonitoringResponse) HasNetworkInDataPresent() bool`

HasNetworkInDataPresent returns a boolean if a field has been set.

### GetNetworkInCoef

`func (o *VmMonitoringResponse) GetNetworkInCoef() float32`

GetNetworkInCoef returns the NetworkInCoef field if non-nil, zero value otherwise.

### GetNetworkInCoefOk

`func (o *VmMonitoringResponse) GetNetworkInCoefOk() (*float32, bool)`

GetNetworkInCoefOk returns a tuple with the NetworkInCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInCoef

`func (o *VmMonitoringResponse) SetNetworkInCoef(v float32)`

SetNetworkInCoef sets NetworkInCoef field to given value.

### HasNetworkInCoef

`func (o *VmMonitoringResponse) HasNetworkInCoef() bool`

HasNetworkInCoef returns a boolean if a field has been set.

### GetNetworkInHuman

`func (o *VmMonitoringResponse) GetNetworkInHuman() float32`

GetNetworkInHuman returns the NetworkInHuman field if non-nil, zero value otherwise.

### GetNetworkInHumanOk

`func (o *VmMonitoringResponse) GetNetworkInHumanOk() (*float32, bool)`

GetNetworkInHumanOk returns a tuple with the NetworkInHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInHuman

`func (o *VmMonitoringResponse) SetNetworkInHuman(v float32)`

SetNetworkInHuman sets NetworkInHuman field to given value.

### HasNetworkInHuman

`func (o *VmMonitoringResponse) HasNetworkInHuman() bool`

HasNetworkInHuman returns a boolean if a field has been set.

### GetAvgNetworkInHuman

`func (o *VmMonitoringResponse) GetAvgNetworkInHuman() float32`

GetAvgNetworkInHuman returns the AvgNetworkInHuman field if non-nil, zero value otherwise.

### GetAvgNetworkInHumanOk

`func (o *VmMonitoringResponse) GetAvgNetworkInHumanOk() (*float32, bool)`

GetAvgNetworkInHumanOk returns a tuple with the AvgNetworkInHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgNetworkInHuman

`func (o *VmMonitoringResponse) SetAvgNetworkInHuman(v float32)`

SetAvgNetworkInHuman sets AvgNetworkInHuman field to given value.

### HasAvgNetworkInHuman

`func (o *VmMonitoringResponse) HasAvgNetworkInHuman() bool`

HasAvgNetworkInHuman returns a boolean if a field has been set.

### GetMaxNetworkInHuman

`func (o *VmMonitoringResponse) GetMaxNetworkInHuman() float32`

GetMaxNetworkInHuman returns the MaxNetworkInHuman field if non-nil, zero value otherwise.

### GetMaxNetworkInHumanOk

`func (o *VmMonitoringResponse) GetMaxNetworkInHumanOk() (*float32, bool)`

GetMaxNetworkInHumanOk returns a tuple with the MaxNetworkInHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworkInHuman

`func (o *VmMonitoringResponse) SetMaxNetworkInHuman(v float32)`

SetMaxNetworkInHuman sets MaxNetworkInHuman field to given value.

### HasMaxNetworkInHuman

`func (o *VmMonitoringResponse) HasMaxNetworkInHuman() bool`

HasMaxNetworkInHuman returns a boolean if a field has been set.

### GetNetworkInHumanLabel

`func (o *VmMonitoringResponse) GetNetworkInHumanLabel() string`

GetNetworkInHumanLabel returns the NetworkInHumanLabel field if non-nil, zero value otherwise.

### GetNetworkInHumanLabelOk

`func (o *VmMonitoringResponse) GetNetworkInHumanLabelOk() (*string, bool)`

GetNetworkInHumanLabelOk returns a tuple with the NetworkInHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInHumanLabel

`func (o *VmMonitoringResponse) SetNetworkInHumanLabel(v string)`

SetNetworkInHumanLabel sets NetworkInHumanLabel field to given value.

### HasNetworkInHumanLabel

`func (o *VmMonitoringResponse) HasNetworkInHumanLabel() bool`

HasNetworkInHumanLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


