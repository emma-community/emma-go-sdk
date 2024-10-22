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


