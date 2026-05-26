# VmAnalyticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmId** | Pointer to **int32** | ID of VM | [optional] 
**Timecode** | Pointer to **string** |  | [optional] 
**AvgDateStart** | Pointer to **string** | Start of the period for average value calculation | [optional] 
**AvgDateEnd** | Pointer to **string** | End of the period for average value calculation | [optional] 
**QuantilesDateStart** | Pointer to **string** | Start of the period for percentile value calculation | [optional] 
**QuantilesDateEnd** | Pointer to **string** | End of the period for percentile value calculation | [optional] 
**CpuDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**CpuUtilizationAverageCores** | Pointer to **float32** | CPU utilization with values in range [0, 100*vCPUs] | [optional] 
**CpuUtilizationAverageCoresQ10** | Pointer to **float32** | 10th percentile of CPU utilization | [optional] 
**CpuUtilizationAverageCoresQ90** | Pointer to **float32** | 90th percentile of CPU utilization | [optional] 
**CpuCoresNumber** | Pointer to **int32** | Total CPU, vCPUs | [optional] 
**CpuTotalPercent** | Pointer to **int32** | Total CPU, % | [optional] 
**CpuHumanLabel** | Pointer to **string** | Label | [optional] 
**RamDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**RamUsageAverageMb** | Pointer to **float32** | Memory utilization | [optional] 
**RamUsageAverageMbQ10** | Pointer to **float32** | 10th percentile of memory utilization | [optional] 
**RamUsageAverageMbQ90** | Pointer to **float32** | 90th percentile of memory utilization | [optional] 
**RamTotalAmountMb** | Pointer to **int32** | Total memory, MB | [optional] 
**RamHumanLabel** | Pointer to **string** | Label | [optional] 
**DiskUsedDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**DiskSpaceUsedGb** | Pointer to **float32** | Disk utilization | [optional] 
**DiskSpaceUsedGbQ10** | Pointer to **float32** | 10th percentile of disk utilization | [optional] 
**DiskSpaceUsedGbQ90** | Pointer to **float32** | 90th percentile of disk utilization | [optional] 
**DiskSpaceTotalGb** | Pointer to **float32** | Total disk size, GB | [optional] 
**DiskSpaceHumanLabel** | Pointer to **string** | Label | [optional] 
**DiskWriteDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**DiskWriteBps** | Pointer to **float32** | Disk write, bps | [optional] 
**DiskWriteHuman** | Pointer to **float32** | Disk write | [optional] 
**DiskWriteHumanLabel** | Pointer to **string** | Label | [optional] 
**DiskReadDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**DiskReadBps** | Pointer to **float32** | Disk read, bps | [optional] 
**DiskReadHuman** | Pointer to **float32** | Disk read | [optional] 
**DiskReadHumanLabel** | Pointer to **string** | Label | [optional] 
**NetworkOutDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**NetworkOutBps** | Pointer to **float32** | Network out, bps | [optional] 
**NetworkOutHuman** | Pointer to **float32** | Network out | [optional] 
**NetworkOutHumanLabel** | Pointer to **string** | Label | [optional] 
**NetworkInDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**NetworkInBps** | Pointer to **float32** | Network in, bps | [optional] 
**NetworkInHuman** | Pointer to **float32** | Network in | [optional] 
**NetworkInHumanLabel** | Pointer to **string** | Label | [optional] 
**GpuDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**GpuUtilizationAvgPercent** | Pointer to **float32** | GPU utilization | [optional] 
**GpuUtilizationAvgPercentQ10** | Pointer to **float32** | 10th percentile of GPU utilization | [optional] 
**GpuUtilizationAvgPercentQ90** | Pointer to **float32** | 90th percentile of GPU utilization | [optional] 
**GpuTotalPercent** | Pointer to **int32** | Total GPU, % | [optional] 
**GpuHumanLabel** | Pointer to **string** | Label | [optional] 
**GpuRamDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**GpuRamUsageAvgGb** | Pointer to **float32** | GPU vRAM usage | [optional] 
**GpuRamUsageAvgGbQ10** | Pointer to **float32** | 10th percentile of GPU vRAM usage | [optional] 
**GpuRamUsageAvgGbQ90** | Pointer to **float32** | 90th percentile of GPU vRAM usage | [optional] 
**GpuRamTotalGb** | Pointer to **float32** | Total GPU vRAM, GB | [optional] 
**GpuRamHumanLabel** | Pointer to **string** | Label | [optional] 
**GpuRamUtilizationDataPresent** | Pointer to **int32** | Internal service parameter | [optional] 
**GpuRamUtilizationAvgPercent** | Pointer to **float32** | GPU vRAM utilization | [optional] 
**GpuRamUtilizationAvgPercentQ10** | Pointer to **float32** | 10th percentile of GPU vRAM utilization | [optional] 
**GpuRamUtilizationAvgPercentQ90** | Pointer to **float32** | 90th percentile of GPU vRAM utilization | [optional] 
**GpuRamUtilizationTotalPercent** | Pointer to **int32** | Total GPU utilization, % | [optional] 
**GpuRamUtilizationHumanLabel** | Pointer to **string** | Label | [optional] 
**IsShownShort** | Pointer to **int32** | Internal service parameter | [optional] 
**Type** | Pointer to **string** | Dataset type | [optional] 

## Methods

### NewVmAnalyticsResponse

`func NewVmAnalyticsResponse() *VmAnalyticsResponse`

NewVmAnalyticsResponse instantiates a new VmAnalyticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmAnalyticsResponseWithDefaults

`func NewVmAnalyticsResponseWithDefaults() *VmAnalyticsResponse`

NewVmAnalyticsResponseWithDefaults instantiates a new VmAnalyticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmId

`func (o *VmAnalyticsResponse) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VmAnalyticsResponse) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VmAnalyticsResponse) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *VmAnalyticsResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetTimecode

`func (o *VmAnalyticsResponse) GetTimecode() string`

GetTimecode returns the Timecode field if non-nil, zero value otherwise.

### GetTimecodeOk

`func (o *VmAnalyticsResponse) GetTimecodeOk() (*string, bool)`

GetTimecodeOk returns a tuple with the Timecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecode

`func (o *VmAnalyticsResponse) SetTimecode(v string)`

SetTimecode sets Timecode field to given value.

### HasTimecode

`func (o *VmAnalyticsResponse) HasTimecode() bool`

HasTimecode returns a boolean if a field has been set.

### GetAvgDateStart

`func (o *VmAnalyticsResponse) GetAvgDateStart() string`

GetAvgDateStart returns the AvgDateStart field if non-nil, zero value otherwise.

### GetAvgDateStartOk

`func (o *VmAnalyticsResponse) GetAvgDateStartOk() (*string, bool)`

GetAvgDateStartOk returns a tuple with the AvgDateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDateStart

`func (o *VmAnalyticsResponse) SetAvgDateStart(v string)`

SetAvgDateStart sets AvgDateStart field to given value.

### HasAvgDateStart

`func (o *VmAnalyticsResponse) HasAvgDateStart() bool`

HasAvgDateStart returns a boolean if a field has been set.

### GetAvgDateEnd

`func (o *VmAnalyticsResponse) GetAvgDateEnd() string`

GetAvgDateEnd returns the AvgDateEnd field if non-nil, zero value otherwise.

### GetAvgDateEndOk

`func (o *VmAnalyticsResponse) GetAvgDateEndOk() (*string, bool)`

GetAvgDateEndOk returns a tuple with the AvgDateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDateEnd

`func (o *VmAnalyticsResponse) SetAvgDateEnd(v string)`

SetAvgDateEnd sets AvgDateEnd field to given value.

### HasAvgDateEnd

`func (o *VmAnalyticsResponse) HasAvgDateEnd() bool`

HasAvgDateEnd returns a boolean if a field has been set.

### GetQuantilesDateStart

`func (o *VmAnalyticsResponse) GetQuantilesDateStart() string`

GetQuantilesDateStart returns the QuantilesDateStart field if non-nil, zero value otherwise.

### GetQuantilesDateStartOk

`func (o *VmAnalyticsResponse) GetQuantilesDateStartOk() (*string, bool)`

GetQuantilesDateStartOk returns a tuple with the QuantilesDateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantilesDateStart

`func (o *VmAnalyticsResponse) SetQuantilesDateStart(v string)`

SetQuantilesDateStart sets QuantilesDateStart field to given value.

### HasQuantilesDateStart

`func (o *VmAnalyticsResponse) HasQuantilesDateStart() bool`

HasQuantilesDateStart returns a boolean if a field has been set.

### GetQuantilesDateEnd

`func (o *VmAnalyticsResponse) GetQuantilesDateEnd() string`

GetQuantilesDateEnd returns the QuantilesDateEnd field if non-nil, zero value otherwise.

### GetQuantilesDateEndOk

`func (o *VmAnalyticsResponse) GetQuantilesDateEndOk() (*string, bool)`

GetQuantilesDateEndOk returns a tuple with the QuantilesDateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantilesDateEnd

`func (o *VmAnalyticsResponse) SetQuantilesDateEnd(v string)`

SetQuantilesDateEnd sets QuantilesDateEnd field to given value.

### HasQuantilesDateEnd

`func (o *VmAnalyticsResponse) HasQuantilesDateEnd() bool`

HasQuantilesDateEnd returns a boolean if a field has been set.

### GetCpuDataPresent

`func (o *VmAnalyticsResponse) GetCpuDataPresent() int32`

GetCpuDataPresent returns the CpuDataPresent field if non-nil, zero value otherwise.

### GetCpuDataPresentOk

`func (o *VmAnalyticsResponse) GetCpuDataPresentOk() (*int32, bool)`

GetCpuDataPresentOk returns a tuple with the CpuDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuDataPresent

`func (o *VmAnalyticsResponse) SetCpuDataPresent(v int32)`

SetCpuDataPresent sets CpuDataPresent field to given value.

### HasCpuDataPresent

`func (o *VmAnalyticsResponse) HasCpuDataPresent() bool`

HasCpuDataPresent returns a boolean if a field has been set.

### GetCpuUtilizationAverageCores

`func (o *VmAnalyticsResponse) GetCpuUtilizationAverageCores() float32`

GetCpuUtilizationAverageCores returns the CpuUtilizationAverageCores field if non-nil, zero value otherwise.

### GetCpuUtilizationAverageCoresOk

`func (o *VmAnalyticsResponse) GetCpuUtilizationAverageCoresOk() (*float32, bool)`

GetCpuUtilizationAverageCoresOk returns a tuple with the CpuUtilizationAverageCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilizationAverageCores

`func (o *VmAnalyticsResponse) SetCpuUtilizationAverageCores(v float32)`

SetCpuUtilizationAverageCores sets CpuUtilizationAverageCores field to given value.

### HasCpuUtilizationAverageCores

`func (o *VmAnalyticsResponse) HasCpuUtilizationAverageCores() bool`

HasCpuUtilizationAverageCores returns a boolean if a field has been set.

### GetCpuUtilizationAverageCoresQ10

`func (o *VmAnalyticsResponse) GetCpuUtilizationAverageCoresQ10() float32`

GetCpuUtilizationAverageCoresQ10 returns the CpuUtilizationAverageCoresQ10 field if non-nil, zero value otherwise.

### GetCpuUtilizationAverageCoresQ10Ok

`func (o *VmAnalyticsResponse) GetCpuUtilizationAverageCoresQ10Ok() (*float32, bool)`

GetCpuUtilizationAverageCoresQ10Ok returns a tuple with the CpuUtilizationAverageCoresQ10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilizationAverageCoresQ10

`func (o *VmAnalyticsResponse) SetCpuUtilizationAverageCoresQ10(v float32)`

SetCpuUtilizationAverageCoresQ10 sets CpuUtilizationAverageCoresQ10 field to given value.

### HasCpuUtilizationAverageCoresQ10

`func (o *VmAnalyticsResponse) HasCpuUtilizationAverageCoresQ10() bool`

HasCpuUtilizationAverageCoresQ10 returns a boolean if a field has been set.

### GetCpuUtilizationAverageCoresQ90

`func (o *VmAnalyticsResponse) GetCpuUtilizationAverageCoresQ90() float32`

GetCpuUtilizationAverageCoresQ90 returns the CpuUtilizationAverageCoresQ90 field if non-nil, zero value otherwise.

### GetCpuUtilizationAverageCoresQ90Ok

`func (o *VmAnalyticsResponse) GetCpuUtilizationAverageCoresQ90Ok() (*float32, bool)`

GetCpuUtilizationAverageCoresQ90Ok returns a tuple with the CpuUtilizationAverageCoresQ90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilizationAverageCoresQ90

`func (o *VmAnalyticsResponse) SetCpuUtilizationAverageCoresQ90(v float32)`

SetCpuUtilizationAverageCoresQ90 sets CpuUtilizationAverageCoresQ90 field to given value.

### HasCpuUtilizationAverageCoresQ90

`func (o *VmAnalyticsResponse) HasCpuUtilizationAverageCoresQ90() bool`

HasCpuUtilizationAverageCoresQ90 returns a boolean if a field has been set.

### GetCpuCoresNumber

`func (o *VmAnalyticsResponse) GetCpuCoresNumber() int32`

GetCpuCoresNumber returns the CpuCoresNumber field if non-nil, zero value otherwise.

### GetCpuCoresNumberOk

`func (o *VmAnalyticsResponse) GetCpuCoresNumberOk() (*int32, bool)`

GetCpuCoresNumberOk returns a tuple with the CpuCoresNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoresNumber

`func (o *VmAnalyticsResponse) SetCpuCoresNumber(v int32)`

SetCpuCoresNumber sets CpuCoresNumber field to given value.

### HasCpuCoresNumber

`func (o *VmAnalyticsResponse) HasCpuCoresNumber() bool`

HasCpuCoresNumber returns a boolean if a field has been set.

### GetCpuTotalPercent

`func (o *VmAnalyticsResponse) GetCpuTotalPercent() int32`

GetCpuTotalPercent returns the CpuTotalPercent field if non-nil, zero value otherwise.

### GetCpuTotalPercentOk

`func (o *VmAnalyticsResponse) GetCpuTotalPercentOk() (*int32, bool)`

GetCpuTotalPercentOk returns a tuple with the CpuTotalPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalPercent

`func (o *VmAnalyticsResponse) SetCpuTotalPercent(v int32)`

SetCpuTotalPercent sets CpuTotalPercent field to given value.

### HasCpuTotalPercent

`func (o *VmAnalyticsResponse) HasCpuTotalPercent() bool`

HasCpuTotalPercent returns a boolean if a field has been set.

### GetCpuHumanLabel

`func (o *VmAnalyticsResponse) GetCpuHumanLabel() string`

GetCpuHumanLabel returns the CpuHumanLabel field if non-nil, zero value otherwise.

### GetCpuHumanLabelOk

`func (o *VmAnalyticsResponse) GetCpuHumanLabelOk() (*string, bool)`

GetCpuHumanLabelOk returns a tuple with the CpuHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHumanLabel

`func (o *VmAnalyticsResponse) SetCpuHumanLabel(v string)`

SetCpuHumanLabel sets CpuHumanLabel field to given value.

### HasCpuHumanLabel

`func (o *VmAnalyticsResponse) HasCpuHumanLabel() bool`

HasCpuHumanLabel returns a boolean if a field has been set.

### GetRamDataPresent

`func (o *VmAnalyticsResponse) GetRamDataPresent() int32`

GetRamDataPresent returns the RamDataPresent field if non-nil, zero value otherwise.

### GetRamDataPresentOk

`func (o *VmAnalyticsResponse) GetRamDataPresentOk() (*int32, bool)`

GetRamDataPresentOk returns a tuple with the RamDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamDataPresent

`func (o *VmAnalyticsResponse) SetRamDataPresent(v int32)`

SetRamDataPresent sets RamDataPresent field to given value.

### HasRamDataPresent

`func (o *VmAnalyticsResponse) HasRamDataPresent() bool`

HasRamDataPresent returns a boolean if a field has been set.

### GetRamUsageAverageMb

`func (o *VmAnalyticsResponse) GetRamUsageAverageMb() float32`

GetRamUsageAverageMb returns the RamUsageAverageMb field if non-nil, zero value otherwise.

### GetRamUsageAverageMbOk

`func (o *VmAnalyticsResponse) GetRamUsageAverageMbOk() (*float32, bool)`

GetRamUsageAverageMbOk returns a tuple with the RamUsageAverageMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamUsageAverageMb

`func (o *VmAnalyticsResponse) SetRamUsageAverageMb(v float32)`

SetRamUsageAverageMb sets RamUsageAverageMb field to given value.

### HasRamUsageAverageMb

`func (o *VmAnalyticsResponse) HasRamUsageAverageMb() bool`

HasRamUsageAverageMb returns a boolean if a field has been set.

### GetRamUsageAverageMbQ10

`func (o *VmAnalyticsResponse) GetRamUsageAverageMbQ10() float32`

GetRamUsageAverageMbQ10 returns the RamUsageAverageMbQ10 field if non-nil, zero value otherwise.

### GetRamUsageAverageMbQ10Ok

`func (o *VmAnalyticsResponse) GetRamUsageAverageMbQ10Ok() (*float32, bool)`

GetRamUsageAverageMbQ10Ok returns a tuple with the RamUsageAverageMbQ10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamUsageAverageMbQ10

`func (o *VmAnalyticsResponse) SetRamUsageAverageMbQ10(v float32)`

SetRamUsageAverageMbQ10 sets RamUsageAverageMbQ10 field to given value.

### HasRamUsageAverageMbQ10

`func (o *VmAnalyticsResponse) HasRamUsageAverageMbQ10() bool`

HasRamUsageAverageMbQ10 returns a boolean if a field has been set.

### GetRamUsageAverageMbQ90

`func (o *VmAnalyticsResponse) GetRamUsageAverageMbQ90() float32`

GetRamUsageAverageMbQ90 returns the RamUsageAverageMbQ90 field if non-nil, zero value otherwise.

### GetRamUsageAverageMbQ90Ok

`func (o *VmAnalyticsResponse) GetRamUsageAverageMbQ90Ok() (*float32, bool)`

GetRamUsageAverageMbQ90Ok returns a tuple with the RamUsageAverageMbQ90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamUsageAverageMbQ90

`func (o *VmAnalyticsResponse) SetRamUsageAverageMbQ90(v float32)`

SetRamUsageAverageMbQ90 sets RamUsageAverageMbQ90 field to given value.

### HasRamUsageAverageMbQ90

`func (o *VmAnalyticsResponse) HasRamUsageAverageMbQ90() bool`

HasRamUsageAverageMbQ90 returns a boolean if a field has been set.

### GetRamTotalAmountMb

`func (o *VmAnalyticsResponse) GetRamTotalAmountMb() int32`

GetRamTotalAmountMb returns the RamTotalAmountMb field if non-nil, zero value otherwise.

### GetRamTotalAmountMbOk

`func (o *VmAnalyticsResponse) GetRamTotalAmountMbOk() (*int32, bool)`

GetRamTotalAmountMbOk returns a tuple with the RamTotalAmountMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamTotalAmountMb

`func (o *VmAnalyticsResponse) SetRamTotalAmountMb(v int32)`

SetRamTotalAmountMb sets RamTotalAmountMb field to given value.

### HasRamTotalAmountMb

`func (o *VmAnalyticsResponse) HasRamTotalAmountMb() bool`

HasRamTotalAmountMb returns a boolean if a field has been set.

### GetRamHumanLabel

`func (o *VmAnalyticsResponse) GetRamHumanLabel() string`

GetRamHumanLabel returns the RamHumanLabel field if non-nil, zero value otherwise.

### GetRamHumanLabelOk

`func (o *VmAnalyticsResponse) GetRamHumanLabelOk() (*string, bool)`

GetRamHumanLabelOk returns a tuple with the RamHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamHumanLabel

`func (o *VmAnalyticsResponse) SetRamHumanLabel(v string)`

SetRamHumanLabel sets RamHumanLabel field to given value.

### HasRamHumanLabel

`func (o *VmAnalyticsResponse) HasRamHumanLabel() bool`

HasRamHumanLabel returns a boolean if a field has been set.

### GetDiskUsedDataPresent

`func (o *VmAnalyticsResponse) GetDiskUsedDataPresent() int32`

GetDiskUsedDataPresent returns the DiskUsedDataPresent field if non-nil, zero value otherwise.

### GetDiskUsedDataPresentOk

`func (o *VmAnalyticsResponse) GetDiskUsedDataPresentOk() (*int32, bool)`

GetDiskUsedDataPresentOk returns a tuple with the DiskUsedDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsedDataPresent

`func (o *VmAnalyticsResponse) SetDiskUsedDataPresent(v int32)`

SetDiskUsedDataPresent sets DiskUsedDataPresent field to given value.

### HasDiskUsedDataPresent

`func (o *VmAnalyticsResponse) HasDiskUsedDataPresent() bool`

HasDiskUsedDataPresent returns a boolean if a field has been set.

### GetDiskSpaceUsedGb

`func (o *VmAnalyticsResponse) GetDiskSpaceUsedGb() float32`

GetDiskSpaceUsedGb returns the DiskSpaceUsedGb field if non-nil, zero value otherwise.

### GetDiskSpaceUsedGbOk

`func (o *VmAnalyticsResponse) GetDiskSpaceUsedGbOk() (*float32, bool)`

GetDiskSpaceUsedGbOk returns a tuple with the DiskSpaceUsedGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsedGb

`func (o *VmAnalyticsResponse) SetDiskSpaceUsedGb(v float32)`

SetDiskSpaceUsedGb sets DiskSpaceUsedGb field to given value.

### HasDiskSpaceUsedGb

`func (o *VmAnalyticsResponse) HasDiskSpaceUsedGb() bool`

HasDiskSpaceUsedGb returns a boolean if a field has been set.

### GetDiskSpaceUsedGbQ10

`func (o *VmAnalyticsResponse) GetDiskSpaceUsedGbQ10() float32`

GetDiskSpaceUsedGbQ10 returns the DiskSpaceUsedGbQ10 field if non-nil, zero value otherwise.

### GetDiskSpaceUsedGbQ10Ok

`func (o *VmAnalyticsResponse) GetDiskSpaceUsedGbQ10Ok() (*float32, bool)`

GetDiskSpaceUsedGbQ10Ok returns a tuple with the DiskSpaceUsedGbQ10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsedGbQ10

`func (o *VmAnalyticsResponse) SetDiskSpaceUsedGbQ10(v float32)`

SetDiskSpaceUsedGbQ10 sets DiskSpaceUsedGbQ10 field to given value.

### HasDiskSpaceUsedGbQ10

`func (o *VmAnalyticsResponse) HasDiskSpaceUsedGbQ10() bool`

HasDiskSpaceUsedGbQ10 returns a boolean if a field has been set.

### GetDiskSpaceUsedGbQ90

`func (o *VmAnalyticsResponse) GetDiskSpaceUsedGbQ90() float32`

GetDiskSpaceUsedGbQ90 returns the DiskSpaceUsedGbQ90 field if non-nil, zero value otherwise.

### GetDiskSpaceUsedGbQ90Ok

`func (o *VmAnalyticsResponse) GetDiskSpaceUsedGbQ90Ok() (*float32, bool)`

GetDiskSpaceUsedGbQ90Ok returns a tuple with the DiskSpaceUsedGbQ90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsedGbQ90

`func (o *VmAnalyticsResponse) SetDiskSpaceUsedGbQ90(v float32)`

SetDiskSpaceUsedGbQ90 sets DiskSpaceUsedGbQ90 field to given value.

### HasDiskSpaceUsedGbQ90

`func (o *VmAnalyticsResponse) HasDiskSpaceUsedGbQ90() bool`

HasDiskSpaceUsedGbQ90 returns a boolean if a field has been set.

### GetDiskSpaceTotalGb

`func (o *VmAnalyticsResponse) GetDiskSpaceTotalGb() float32`

GetDiskSpaceTotalGb returns the DiskSpaceTotalGb field if non-nil, zero value otherwise.

### GetDiskSpaceTotalGbOk

`func (o *VmAnalyticsResponse) GetDiskSpaceTotalGbOk() (*float32, bool)`

GetDiskSpaceTotalGbOk returns a tuple with the DiskSpaceTotalGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceTotalGb

`func (o *VmAnalyticsResponse) SetDiskSpaceTotalGb(v float32)`

SetDiskSpaceTotalGb sets DiskSpaceTotalGb field to given value.

### HasDiskSpaceTotalGb

`func (o *VmAnalyticsResponse) HasDiskSpaceTotalGb() bool`

HasDiskSpaceTotalGb returns a boolean if a field has been set.

### GetDiskSpaceHumanLabel

`func (o *VmAnalyticsResponse) GetDiskSpaceHumanLabel() string`

GetDiskSpaceHumanLabel returns the DiskSpaceHumanLabel field if non-nil, zero value otherwise.

### GetDiskSpaceHumanLabelOk

`func (o *VmAnalyticsResponse) GetDiskSpaceHumanLabelOk() (*string, bool)`

GetDiskSpaceHumanLabelOk returns a tuple with the DiskSpaceHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceHumanLabel

`func (o *VmAnalyticsResponse) SetDiskSpaceHumanLabel(v string)`

SetDiskSpaceHumanLabel sets DiskSpaceHumanLabel field to given value.

### HasDiskSpaceHumanLabel

`func (o *VmAnalyticsResponse) HasDiskSpaceHumanLabel() bool`

HasDiskSpaceHumanLabel returns a boolean if a field has been set.

### GetDiskWriteDataPresent

`func (o *VmAnalyticsResponse) GetDiskWriteDataPresent() int32`

GetDiskWriteDataPresent returns the DiskWriteDataPresent field if non-nil, zero value otherwise.

### GetDiskWriteDataPresentOk

`func (o *VmAnalyticsResponse) GetDiskWriteDataPresentOk() (*int32, bool)`

GetDiskWriteDataPresentOk returns a tuple with the DiskWriteDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteDataPresent

`func (o *VmAnalyticsResponse) SetDiskWriteDataPresent(v int32)`

SetDiskWriteDataPresent sets DiskWriteDataPresent field to given value.

### HasDiskWriteDataPresent

`func (o *VmAnalyticsResponse) HasDiskWriteDataPresent() bool`

HasDiskWriteDataPresent returns a boolean if a field has been set.

### GetDiskWriteBps

`func (o *VmAnalyticsResponse) GetDiskWriteBps() float32`

GetDiskWriteBps returns the DiskWriteBps field if non-nil, zero value otherwise.

### GetDiskWriteBpsOk

`func (o *VmAnalyticsResponse) GetDiskWriteBpsOk() (*float32, bool)`

GetDiskWriteBpsOk returns a tuple with the DiskWriteBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteBps

`func (o *VmAnalyticsResponse) SetDiskWriteBps(v float32)`

SetDiskWriteBps sets DiskWriteBps field to given value.

### HasDiskWriteBps

`func (o *VmAnalyticsResponse) HasDiskWriteBps() bool`

HasDiskWriteBps returns a boolean if a field has been set.

### GetDiskWriteHuman

`func (o *VmAnalyticsResponse) GetDiskWriteHuman() float32`

GetDiskWriteHuman returns the DiskWriteHuman field if non-nil, zero value otherwise.

### GetDiskWriteHumanOk

`func (o *VmAnalyticsResponse) GetDiskWriteHumanOk() (*float32, bool)`

GetDiskWriteHumanOk returns a tuple with the DiskWriteHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteHuman

`func (o *VmAnalyticsResponse) SetDiskWriteHuman(v float32)`

SetDiskWriteHuman sets DiskWriteHuman field to given value.

### HasDiskWriteHuman

`func (o *VmAnalyticsResponse) HasDiskWriteHuman() bool`

HasDiskWriteHuman returns a boolean if a field has been set.

### GetDiskWriteHumanLabel

`func (o *VmAnalyticsResponse) GetDiskWriteHumanLabel() string`

GetDiskWriteHumanLabel returns the DiskWriteHumanLabel field if non-nil, zero value otherwise.

### GetDiskWriteHumanLabelOk

`func (o *VmAnalyticsResponse) GetDiskWriteHumanLabelOk() (*string, bool)`

GetDiskWriteHumanLabelOk returns a tuple with the DiskWriteHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteHumanLabel

`func (o *VmAnalyticsResponse) SetDiskWriteHumanLabel(v string)`

SetDiskWriteHumanLabel sets DiskWriteHumanLabel field to given value.

### HasDiskWriteHumanLabel

`func (o *VmAnalyticsResponse) HasDiskWriteHumanLabel() bool`

HasDiskWriteHumanLabel returns a boolean if a field has been set.

### GetDiskReadDataPresent

`func (o *VmAnalyticsResponse) GetDiskReadDataPresent() int32`

GetDiskReadDataPresent returns the DiskReadDataPresent field if non-nil, zero value otherwise.

### GetDiskReadDataPresentOk

`func (o *VmAnalyticsResponse) GetDiskReadDataPresentOk() (*int32, bool)`

GetDiskReadDataPresentOk returns a tuple with the DiskReadDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadDataPresent

`func (o *VmAnalyticsResponse) SetDiskReadDataPresent(v int32)`

SetDiskReadDataPresent sets DiskReadDataPresent field to given value.

### HasDiskReadDataPresent

`func (o *VmAnalyticsResponse) HasDiskReadDataPresent() bool`

HasDiskReadDataPresent returns a boolean if a field has been set.

### GetDiskReadBps

`func (o *VmAnalyticsResponse) GetDiskReadBps() float32`

GetDiskReadBps returns the DiskReadBps field if non-nil, zero value otherwise.

### GetDiskReadBpsOk

`func (o *VmAnalyticsResponse) GetDiskReadBpsOk() (*float32, bool)`

GetDiskReadBpsOk returns a tuple with the DiskReadBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadBps

`func (o *VmAnalyticsResponse) SetDiskReadBps(v float32)`

SetDiskReadBps sets DiskReadBps field to given value.

### HasDiskReadBps

`func (o *VmAnalyticsResponse) HasDiskReadBps() bool`

HasDiskReadBps returns a boolean if a field has been set.

### GetDiskReadHuman

`func (o *VmAnalyticsResponse) GetDiskReadHuman() float32`

GetDiskReadHuman returns the DiskReadHuman field if non-nil, zero value otherwise.

### GetDiskReadHumanOk

`func (o *VmAnalyticsResponse) GetDiskReadHumanOk() (*float32, bool)`

GetDiskReadHumanOk returns a tuple with the DiskReadHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadHuman

`func (o *VmAnalyticsResponse) SetDiskReadHuman(v float32)`

SetDiskReadHuman sets DiskReadHuman field to given value.

### HasDiskReadHuman

`func (o *VmAnalyticsResponse) HasDiskReadHuman() bool`

HasDiskReadHuman returns a boolean if a field has been set.

### GetDiskReadHumanLabel

`func (o *VmAnalyticsResponse) GetDiskReadHumanLabel() string`

GetDiskReadHumanLabel returns the DiskReadHumanLabel field if non-nil, zero value otherwise.

### GetDiskReadHumanLabelOk

`func (o *VmAnalyticsResponse) GetDiskReadHumanLabelOk() (*string, bool)`

GetDiskReadHumanLabelOk returns a tuple with the DiskReadHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadHumanLabel

`func (o *VmAnalyticsResponse) SetDiskReadHumanLabel(v string)`

SetDiskReadHumanLabel sets DiskReadHumanLabel field to given value.

### HasDiskReadHumanLabel

`func (o *VmAnalyticsResponse) HasDiskReadHumanLabel() bool`

HasDiskReadHumanLabel returns a boolean if a field has been set.

### GetNetworkOutDataPresent

`func (o *VmAnalyticsResponse) GetNetworkOutDataPresent() int32`

GetNetworkOutDataPresent returns the NetworkOutDataPresent field if non-nil, zero value otherwise.

### GetNetworkOutDataPresentOk

`func (o *VmAnalyticsResponse) GetNetworkOutDataPresentOk() (*int32, bool)`

GetNetworkOutDataPresentOk returns a tuple with the NetworkOutDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutDataPresent

`func (o *VmAnalyticsResponse) SetNetworkOutDataPresent(v int32)`

SetNetworkOutDataPresent sets NetworkOutDataPresent field to given value.

### HasNetworkOutDataPresent

`func (o *VmAnalyticsResponse) HasNetworkOutDataPresent() bool`

HasNetworkOutDataPresent returns a boolean if a field has been set.

### GetNetworkOutBps

`func (o *VmAnalyticsResponse) GetNetworkOutBps() float32`

GetNetworkOutBps returns the NetworkOutBps field if non-nil, zero value otherwise.

### GetNetworkOutBpsOk

`func (o *VmAnalyticsResponse) GetNetworkOutBpsOk() (*float32, bool)`

GetNetworkOutBpsOk returns a tuple with the NetworkOutBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutBps

`func (o *VmAnalyticsResponse) SetNetworkOutBps(v float32)`

SetNetworkOutBps sets NetworkOutBps field to given value.

### HasNetworkOutBps

`func (o *VmAnalyticsResponse) HasNetworkOutBps() bool`

HasNetworkOutBps returns a boolean if a field has been set.

### GetNetworkOutHuman

`func (o *VmAnalyticsResponse) GetNetworkOutHuman() float32`

GetNetworkOutHuman returns the NetworkOutHuman field if non-nil, zero value otherwise.

### GetNetworkOutHumanOk

`func (o *VmAnalyticsResponse) GetNetworkOutHumanOk() (*float32, bool)`

GetNetworkOutHumanOk returns a tuple with the NetworkOutHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutHuman

`func (o *VmAnalyticsResponse) SetNetworkOutHuman(v float32)`

SetNetworkOutHuman sets NetworkOutHuman field to given value.

### HasNetworkOutHuman

`func (o *VmAnalyticsResponse) HasNetworkOutHuman() bool`

HasNetworkOutHuman returns a boolean if a field has been set.

### GetNetworkOutHumanLabel

`func (o *VmAnalyticsResponse) GetNetworkOutHumanLabel() string`

GetNetworkOutHumanLabel returns the NetworkOutHumanLabel field if non-nil, zero value otherwise.

### GetNetworkOutHumanLabelOk

`func (o *VmAnalyticsResponse) GetNetworkOutHumanLabelOk() (*string, bool)`

GetNetworkOutHumanLabelOk returns a tuple with the NetworkOutHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOutHumanLabel

`func (o *VmAnalyticsResponse) SetNetworkOutHumanLabel(v string)`

SetNetworkOutHumanLabel sets NetworkOutHumanLabel field to given value.

### HasNetworkOutHumanLabel

`func (o *VmAnalyticsResponse) HasNetworkOutHumanLabel() bool`

HasNetworkOutHumanLabel returns a boolean if a field has been set.

### GetNetworkInDataPresent

`func (o *VmAnalyticsResponse) GetNetworkInDataPresent() int32`

GetNetworkInDataPresent returns the NetworkInDataPresent field if non-nil, zero value otherwise.

### GetNetworkInDataPresentOk

`func (o *VmAnalyticsResponse) GetNetworkInDataPresentOk() (*int32, bool)`

GetNetworkInDataPresentOk returns a tuple with the NetworkInDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInDataPresent

`func (o *VmAnalyticsResponse) SetNetworkInDataPresent(v int32)`

SetNetworkInDataPresent sets NetworkInDataPresent field to given value.

### HasNetworkInDataPresent

`func (o *VmAnalyticsResponse) HasNetworkInDataPresent() bool`

HasNetworkInDataPresent returns a boolean if a field has been set.

### GetNetworkInBps

`func (o *VmAnalyticsResponse) GetNetworkInBps() float32`

GetNetworkInBps returns the NetworkInBps field if non-nil, zero value otherwise.

### GetNetworkInBpsOk

`func (o *VmAnalyticsResponse) GetNetworkInBpsOk() (*float32, bool)`

GetNetworkInBpsOk returns a tuple with the NetworkInBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInBps

`func (o *VmAnalyticsResponse) SetNetworkInBps(v float32)`

SetNetworkInBps sets NetworkInBps field to given value.

### HasNetworkInBps

`func (o *VmAnalyticsResponse) HasNetworkInBps() bool`

HasNetworkInBps returns a boolean if a field has been set.

### GetNetworkInHuman

`func (o *VmAnalyticsResponse) GetNetworkInHuman() float32`

GetNetworkInHuman returns the NetworkInHuman field if non-nil, zero value otherwise.

### GetNetworkInHumanOk

`func (o *VmAnalyticsResponse) GetNetworkInHumanOk() (*float32, bool)`

GetNetworkInHumanOk returns a tuple with the NetworkInHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInHuman

`func (o *VmAnalyticsResponse) SetNetworkInHuman(v float32)`

SetNetworkInHuman sets NetworkInHuman field to given value.

### HasNetworkInHuman

`func (o *VmAnalyticsResponse) HasNetworkInHuman() bool`

HasNetworkInHuman returns a boolean if a field has been set.

### GetNetworkInHumanLabel

`func (o *VmAnalyticsResponse) GetNetworkInHumanLabel() string`

GetNetworkInHumanLabel returns the NetworkInHumanLabel field if non-nil, zero value otherwise.

### GetNetworkInHumanLabelOk

`func (o *VmAnalyticsResponse) GetNetworkInHumanLabelOk() (*string, bool)`

GetNetworkInHumanLabelOk returns a tuple with the NetworkInHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInHumanLabel

`func (o *VmAnalyticsResponse) SetNetworkInHumanLabel(v string)`

SetNetworkInHumanLabel sets NetworkInHumanLabel field to given value.

### HasNetworkInHumanLabel

`func (o *VmAnalyticsResponse) HasNetworkInHumanLabel() bool`

HasNetworkInHumanLabel returns a boolean if a field has been set.

### GetGpuDataPresent

`func (o *VmAnalyticsResponse) GetGpuDataPresent() int32`

GetGpuDataPresent returns the GpuDataPresent field if non-nil, zero value otherwise.

### GetGpuDataPresentOk

`func (o *VmAnalyticsResponse) GetGpuDataPresentOk() (*int32, bool)`

GetGpuDataPresentOk returns a tuple with the GpuDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDataPresent

`func (o *VmAnalyticsResponse) SetGpuDataPresent(v int32)`

SetGpuDataPresent sets GpuDataPresent field to given value.

### HasGpuDataPresent

`func (o *VmAnalyticsResponse) HasGpuDataPresent() bool`

HasGpuDataPresent returns a boolean if a field has been set.

### GetGpuUtilizationAvgPercent

`func (o *VmAnalyticsResponse) GetGpuUtilizationAvgPercent() float32`

GetGpuUtilizationAvgPercent returns the GpuUtilizationAvgPercent field if non-nil, zero value otherwise.

### GetGpuUtilizationAvgPercentOk

`func (o *VmAnalyticsResponse) GetGpuUtilizationAvgPercentOk() (*float32, bool)`

GetGpuUtilizationAvgPercentOk returns a tuple with the GpuUtilizationAvgPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtilizationAvgPercent

`func (o *VmAnalyticsResponse) SetGpuUtilizationAvgPercent(v float32)`

SetGpuUtilizationAvgPercent sets GpuUtilizationAvgPercent field to given value.

### HasGpuUtilizationAvgPercent

`func (o *VmAnalyticsResponse) HasGpuUtilizationAvgPercent() bool`

HasGpuUtilizationAvgPercent returns a boolean if a field has been set.

### GetGpuUtilizationAvgPercentQ10

`func (o *VmAnalyticsResponse) GetGpuUtilizationAvgPercentQ10() float32`

GetGpuUtilizationAvgPercentQ10 returns the GpuUtilizationAvgPercentQ10 field if non-nil, zero value otherwise.

### GetGpuUtilizationAvgPercentQ10Ok

`func (o *VmAnalyticsResponse) GetGpuUtilizationAvgPercentQ10Ok() (*float32, bool)`

GetGpuUtilizationAvgPercentQ10Ok returns a tuple with the GpuUtilizationAvgPercentQ10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtilizationAvgPercentQ10

`func (o *VmAnalyticsResponse) SetGpuUtilizationAvgPercentQ10(v float32)`

SetGpuUtilizationAvgPercentQ10 sets GpuUtilizationAvgPercentQ10 field to given value.

### HasGpuUtilizationAvgPercentQ10

`func (o *VmAnalyticsResponse) HasGpuUtilizationAvgPercentQ10() bool`

HasGpuUtilizationAvgPercentQ10 returns a boolean if a field has been set.

### GetGpuUtilizationAvgPercentQ90

`func (o *VmAnalyticsResponse) GetGpuUtilizationAvgPercentQ90() float32`

GetGpuUtilizationAvgPercentQ90 returns the GpuUtilizationAvgPercentQ90 field if non-nil, zero value otherwise.

### GetGpuUtilizationAvgPercentQ90Ok

`func (o *VmAnalyticsResponse) GetGpuUtilizationAvgPercentQ90Ok() (*float32, bool)`

GetGpuUtilizationAvgPercentQ90Ok returns a tuple with the GpuUtilizationAvgPercentQ90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtilizationAvgPercentQ90

`func (o *VmAnalyticsResponse) SetGpuUtilizationAvgPercentQ90(v float32)`

SetGpuUtilizationAvgPercentQ90 sets GpuUtilizationAvgPercentQ90 field to given value.

### HasGpuUtilizationAvgPercentQ90

`func (o *VmAnalyticsResponse) HasGpuUtilizationAvgPercentQ90() bool`

HasGpuUtilizationAvgPercentQ90 returns a boolean if a field has been set.

### GetGpuTotalPercent

`func (o *VmAnalyticsResponse) GetGpuTotalPercent() int32`

GetGpuTotalPercent returns the GpuTotalPercent field if non-nil, zero value otherwise.

### GetGpuTotalPercentOk

`func (o *VmAnalyticsResponse) GetGpuTotalPercentOk() (*int32, bool)`

GetGpuTotalPercentOk returns a tuple with the GpuTotalPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTotalPercent

`func (o *VmAnalyticsResponse) SetGpuTotalPercent(v int32)`

SetGpuTotalPercent sets GpuTotalPercent field to given value.

### HasGpuTotalPercent

`func (o *VmAnalyticsResponse) HasGpuTotalPercent() bool`

HasGpuTotalPercent returns a boolean if a field has been set.

### GetGpuHumanLabel

`func (o *VmAnalyticsResponse) GetGpuHumanLabel() string`

GetGpuHumanLabel returns the GpuHumanLabel field if non-nil, zero value otherwise.

### GetGpuHumanLabelOk

`func (o *VmAnalyticsResponse) GetGpuHumanLabelOk() (*string, bool)`

GetGpuHumanLabelOk returns a tuple with the GpuHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuHumanLabel

`func (o *VmAnalyticsResponse) SetGpuHumanLabel(v string)`

SetGpuHumanLabel sets GpuHumanLabel field to given value.

### HasGpuHumanLabel

`func (o *VmAnalyticsResponse) HasGpuHumanLabel() bool`

HasGpuHumanLabel returns a boolean if a field has been set.

### GetGpuRamDataPresent

`func (o *VmAnalyticsResponse) GetGpuRamDataPresent() int32`

GetGpuRamDataPresent returns the GpuRamDataPresent field if non-nil, zero value otherwise.

### GetGpuRamDataPresentOk

`func (o *VmAnalyticsResponse) GetGpuRamDataPresentOk() (*int32, bool)`

GetGpuRamDataPresentOk returns a tuple with the GpuRamDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamDataPresent

`func (o *VmAnalyticsResponse) SetGpuRamDataPresent(v int32)`

SetGpuRamDataPresent sets GpuRamDataPresent field to given value.

### HasGpuRamDataPresent

`func (o *VmAnalyticsResponse) HasGpuRamDataPresent() bool`

HasGpuRamDataPresent returns a boolean if a field has been set.

### GetGpuRamUsageAvgGb

`func (o *VmAnalyticsResponse) GetGpuRamUsageAvgGb() float32`

GetGpuRamUsageAvgGb returns the GpuRamUsageAvgGb field if non-nil, zero value otherwise.

### GetGpuRamUsageAvgGbOk

`func (o *VmAnalyticsResponse) GetGpuRamUsageAvgGbOk() (*float32, bool)`

GetGpuRamUsageAvgGbOk returns a tuple with the GpuRamUsageAvgGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUsageAvgGb

`func (o *VmAnalyticsResponse) SetGpuRamUsageAvgGb(v float32)`

SetGpuRamUsageAvgGb sets GpuRamUsageAvgGb field to given value.

### HasGpuRamUsageAvgGb

`func (o *VmAnalyticsResponse) HasGpuRamUsageAvgGb() bool`

HasGpuRamUsageAvgGb returns a boolean if a field has been set.

### GetGpuRamUsageAvgGbQ10

`func (o *VmAnalyticsResponse) GetGpuRamUsageAvgGbQ10() float32`

GetGpuRamUsageAvgGbQ10 returns the GpuRamUsageAvgGbQ10 field if non-nil, zero value otherwise.

### GetGpuRamUsageAvgGbQ10Ok

`func (o *VmAnalyticsResponse) GetGpuRamUsageAvgGbQ10Ok() (*float32, bool)`

GetGpuRamUsageAvgGbQ10Ok returns a tuple with the GpuRamUsageAvgGbQ10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUsageAvgGbQ10

`func (o *VmAnalyticsResponse) SetGpuRamUsageAvgGbQ10(v float32)`

SetGpuRamUsageAvgGbQ10 sets GpuRamUsageAvgGbQ10 field to given value.

### HasGpuRamUsageAvgGbQ10

`func (o *VmAnalyticsResponse) HasGpuRamUsageAvgGbQ10() bool`

HasGpuRamUsageAvgGbQ10 returns a boolean if a field has been set.

### GetGpuRamUsageAvgGbQ90

`func (o *VmAnalyticsResponse) GetGpuRamUsageAvgGbQ90() float32`

GetGpuRamUsageAvgGbQ90 returns the GpuRamUsageAvgGbQ90 field if non-nil, zero value otherwise.

### GetGpuRamUsageAvgGbQ90Ok

`func (o *VmAnalyticsResponse) GetGpuRamUsageAvgGbQ90Ok() (*float32, bool)`

GetGpuRamUsageAvgGbQ90Ok returns a tuple with the GpuRamUsageAvgGbQ90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUsageAvgGbQ90

`func (o *VmAnalyticsResponse) SetGpuRamUsageAvgGbQ90(v float32)`

SetGpuRamUsageAvgGbQ90 sets GpuRamUsageAvgGbQ90 field to given value.

### HasGpuRamUsageAvgGbQ90

`func (o *VmAnalyticsResponse) HasGpuRamUsageAvgGbQ90() bool`

HasGpuRamUsageAvgGbQ90 returns a boolean if a field has been set.

### GetGpuRamTotalGb

`func (o *VmAnalyticsResponse) GetGpuRamTotalGb() float32`

GetGpuRamTotalGb returns the GpuRamTotalGb field if non-nil, zero value otherwise.

### GetGpuRamTotalGbOk

`func (o *VmAnalyticsResponse) GetGpuRamTotalGbOk() (*float32, bool)`

GetGpuRamTotalGbOk returns a tuple with the GpuRamTotalGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamTotalGb

`func (o *VmAnalyticsResponse) SetGpuRamTotalGb(v float32)`

SetGpuRamTotalGb sets GpuRamTotalGb field to given value.

### HasGpuRamTotalGb

`func (o *VmAnalyticsResponse) HasGpuRamTotalGb() bool`

HasGpuRamTotalGb returns a boolean if a field has been set.

### GetGpuRamHumanLabel

`func (o *VmAnalyticsResponse) GetGpuRamHumanLabel() string`

GetGpuRamHumanLabel returns the GpuRamHumanLabel field if non-nil, zero value otherwise.

### GetGpuRamHumanLabelOk

`func (o *VmAnalyticsResponse) GetGpuRamHumanLabelOk() (*string, bool)`

GetGpuRamHumanLabelOk returns a tuple with the GpuRamHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamHumanLabel

`func (o *VmAnalyticsResponse) SetGpuRamHumanLabel(v string)`

SetGpuRamHumanLabel sets GpuRamHumanLabel field to given value.

### HasGpuRamHumanLabel

`func (o *VmAnalyticsResponse) HasGpuRamHumanLabel() bool`

HasGpuRamHumanLabel returns a boolean if a field has been set.

### GetGpuRamUtilizationDataPresent

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationDataPresent() int32`

GetGpuRamUtilizationDataPresent returns the GpuRamUtilizationDataPresent field if non-nil, zero value otherwise.

### GetGpuRamUtilizationDataPresentOk

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationDataPresentOk() (*int32, bool)`

GetGpuRamUtilizationDataPresentOk returns a tuple with the GpuRamUtilizationDataPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUtilizationDataPresent

`func (o *VmAnalyticsResponse) SetGpuRamUtilizationDataPresent(v int32)`

SetGpuRamUtilizationDataPresent sets GpuRamUtilizationDataPresent field to given value.

### HasGpuRamUtilizationDataPresent

`func (o *VmAnalyticsResponse) HasGpuRamUtilizationDataPresent() bool`

HasGpuRamUtilizationDataPresent returns a boolean if a field has been set.

### GetGpuRamUtilizationAvgPercent

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationAvgPercent() float32`

GetGpuRamUtilizationAvgPercent returns the GpuRamUtilizationAvgPercent field if non-nil, zero value otherwise.

### GetGpuRamUtilizationAvgPercentOk

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationAvgPercentOk() (*float32, bool)`

GetGpuRamUtilizationAvgPercentOk returns a tuple with the GpuRamUtilizationAvgPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUtilizationAvgPercent

`func (o *VmAnalyticsResponse) SetGpuRamUtilizationAvgPercent(v float32)`

SetGpuRamUtilizationAvgPercent sets GpuRamUtilizationAvgPercent field to given value.

### HasGpuRamUtilizationAvgPercent

`func (o *VmAnalyticsResponse) HasGpuRamUtilizationAvgPercent() bool`

HasGpuRamUtilizationAvgPercent returns a boolean if a field has been set.

### GetGpuRamUtilizationAvgPercentQ10

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationAvgPercentQ10() float32`

GetGpuRamUtilizationAvgPercentQ10 returns the GpuRamUtilizationAvgPercentQ10 field if non-nil, zero value otherwise.

### GetGpuRamUtilizationAvgPercentQ10Ok

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationAvgPercentQ10Ok() (*float32, bool)`

GetGpuRamUtilizationAvgPercentQ10Ok returns a tuple with the GpuRamUtilizationAvgPercentQ10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUtilizationAvgPercentQ10

`func (o *VmAnalyticsResponse) SetGpuRamUtilizationAvgPercentQ10(v float32)`

SetGpuRamUtilizationAvgPercentQ10 sets GpuRamUtilizationAvgPercentQ10 field to given value.

### HasGpuRamUtilizationAvgPercentQ10

`func (o *VmAnalyticsResponse) HasGpuRamUtilizationAvgPercentQ10() bool`

HasGpuRamUtilizationAvgPercentQ10 returns a boolean if a field has been set.

### GetGpuRamUtilizationAvgPercentQ90

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationAvgPercentQ90() float32`

GetGpuRamUtilizationAvgPercentQ90 returns the GpuRamUtilizationAvgPercentQ90 field if non-nil, zero value otherwise.

### GetGpuRamUtilizationAvgPercentQ90Ok

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationAvgPercentQ90Ok() (*float32, bool)`

GetGpuRamUtilizationAvgPercentQ90Ok returns a tuple with the GpuRamUtilizationAvgPercentQ90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUtilizationAvgPercentQ90

`func (o *VmAnalyticsResponse) SetGpuRamUtilizationAvgPercentQ90(v float32)`

SetGpuRamUtilizationAvgPercentQ90 sets GpuRamUtilizationAvgPercentQ90 field to given value.

### HasGpuRamUtilizationAvgPercentQ90

`func (o *VmAnalyticsResponse) HasGpuRamUtilizationAvgPercentQ90() bool`

HasGpuRamUtilizationAvgPercentQ90 returns a boolean if a field has been set.

### GetGpuRamUtilizationTotalPercent

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationTotalPercent() int32`

GetGpuRamUtilizationTotalPercent returns the GpuRamUtilizationTotalPercent field if non-nil, zero value otherwise.

### GetGpuRamUtilizationTotalPercentOk

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationTotalPercentOk() (*int32, bool)`

GetGpuRamUtilizationTotalPercentOk returns a tuple with the GpuRamUtilizationTotalPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUtilizationTotalPercent

`func (o *VmAnalyticsResponse) SetGpuRamUtilizationTotalPercent(v int32)`

SetGpuRamUtilizationTotalPercent sets GpuRamUtilizationTotalPercent field to given value.

### HasGpuRamUtilizationTotalPercent

`func (o *VmAnalyticsResponse) HasGpuRamUtilizationTotalPercent() bool`

HasGpuRamUtilizationTotalPercent returns a boolean if a field has been set.

### GetGpuRamUtilizationHumanLabel

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationHumanLabel() string`

GetGpuRamUtilizationHumanLabel returns the GpuRamUtilizationHumanLabel field if non-nil, zero value otherwise.

### GetGpuRamUtilizationHumanLabelOk

`func (o *VmAnalyticsResponse) GetGpuRamUtilizationHumanLabelOk() (*string, bool)`

GetGpuRamUtilizationHumanLabelOk returns a tuple with the GpuRamUtilizationHumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuRamUtilizationHumanLabel

`func (o *VmAnalyticsResponse) SetGpuRamUtilizationHumanLabel(v string)`

SetGpuRamUtilizationHumanLabel sets GpuRamUtilizationHumanLabel field to given value.

### HasGpuRamUtilizationHumanLabel

`func (o *VmAnalyticsResponse) HasGpuRamUtilizationHumanLabel() bool`

HasGpuRamUtilizationHumanLabel returns a boolean if a field has been set.

### GetIsShownShort

`func (o *VmAnalyticsResponse) GetIsShownShort() int32`

GetIsShownShort returns the IsShownShort field if non-nil, zero value otherwise.

### GetIsShownShortOk

`func (o *VmAnalyticsResponse) GetIsShownShortOk() (*int32, bool)`

GetIsShownShortOk returns a tuple with the IsShownShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShownShort

`func (o *VmAnalyticsResponse) SetIsShownShort(v int32)`

SetIsShownShort sets IsShownShort field to given value.

### HasIsShownShort

`func (o *VmAnalyticsResponse) HasIsShownShort() bool`

HasIsShownShort returns a boolean if a field has been set.

### GetType

`func (o *VmAnalyticsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmAnalyticsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmAnalyticsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VmAnalyticsResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


