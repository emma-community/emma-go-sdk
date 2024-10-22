# KubernetesClusterChangingMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubobjectName** | Pointer to **string** | Name of sub-object | [optional] 
**MetricName** | Pointer to **string** | emma&#39;s internal metric name | [optional] 
**UiMetricGroup** | Pointer to **string** | UI group of metrics | [optional] 
**UiColorGroupId** | Pointer to **int32** | UI color for a group of metrics | [optional] 
**UiMetricName** | Pointer to **string** | UI metric name | [optional] 
**HumanLabel** | Pointer to **string** | Label | [optional] 
**Timecodes** | Pointer to **[]string** |  | [optional] 
**LinechartValues** | Pointer to **[]float32** | List of metric values for line-chart | [optional] 
**TreemapValues** | Pointer to **float32** | Average metric value for treemap-chart | [optional] 

## Methods

### NewKubernetesClusterChangingMetricsResponse

`func NewKubernetesClusterChangingMetricsResponse() *KubernetesClusterChangingMetricsResponse`

NewKubernetesClusterChangingMetricsResponse instantiates a new KubernetesClusterChangingMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterChangingMetricsResponseWithDefaults

`func NewKubernetesClusterChangingMetricsResponseWithDefaults() *KubernetesClusterChangingMetricsResponse`

NewKubernetesClusterChangingMetricsResponseWithDefaults instantiates a new KubernetesClusterChangingMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubobjectName

`func (o *KubernetesClusterChangingMetricsResponse) GetSubobjectName() string`

GetSubobjectName returns the SubobjectName field if non-nil, zero value otherwise.

### GetSubobjectNameOk

`func (o *KubernetesClusterChangingMetricsResponse) GetSubobjectNameOk() (*string, bool)`

GetSubobjectNameOk returns a tuple with the SubobjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubobjectName

`func (o *KubernetesClusterChangingMetricsResponse) SetSubobjectName(v string)`

SetSubobjectName sets SubobjectName field to given value.

### HasSubobjectName

`func (o *KubernetesClusterChangingMetricsResponse) HasSubobjectName() bool`

HasSubobjectName returns a boolean if a field has been set.

### GetMetricName

`func (o *KubernetesClusterChangingMetricsResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *KubernetesClusterChangingMetricsResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *KubernetesClusterChangingMetricsResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *KubernetesClusterChangingMetricsResponse) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetUiMetricGroup

`func (o *KubernetesClusterChangingMetricsResponse) GetUiMetricGroup() string`

GetUiMetricGroup returns the UiMetricGroup field if non-nil, zero value otherwise.

### GetUiMetricGroupOk

`func (o *KubernetesClusterChangingMetricsResponse) GetUiMetricGroupOk() (*string, bool)`

GetUiMetricGroupOk returns a tuple with the UiMetricGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMetricGroup

`func (o *KubernetesClusterChangingMetricsResponse) SetUiMetricGroup(v string)`

SetUiMetricGroup sets UiMetricGroup field to given value.

### HasUiMetricGroup

`func (o *KubernetesClusterChangingMetricsResponse) HasUiMetricGroup() bool`

HasUiMetricGroup returns a boolean if a field has been set.

### GetUiColorGroupId

`func (o *KubernetesClusterChangingMetricsResponse) GetUiColorGroupId() int32`

GetUiColorGroupId returns the UiColorGroupId field if non-nil, zero value otherwise.

### GetUiColorGroupIdOk

`func (o *KubernetesClusterChangingMetricsResponse) GetUiColorGroupIdOk() (*int32, bool)`

GetUiColorGroupIdOk returns a tuple with the UiColorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiColorGroupId

`func (o *KubernetesClusterChangingMetricsResponse) SetUiColorGroupId(v int32)`

SetUiColorGroupId sets UiColorGroupId field to given value.

### HasUiColorGroupId

`func (o *KubernetesClusterChangingMetricsResponse) HasUiColorGroupId() bool`

HasUiColorGroupId returns a boolean if a field has been set.

### GetUiMetricName

`func (o *KubernetesClusterChangingMetricsResponse) GetUiMetricName() string`

GetUiMetricName returns the UiMetricName field if non-nil, zero value otherwise.

### GetUiMetricNameOk

`func (o *KubernetesClusterChangingMetricsResponse) GetUiMetricNameOk() (*string, bool)`

GetUiMetricNameOk returns a tuple with the UiMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMetricName

`func (o *KubernetesClusterChangingMetricsResponse) SetUiMetricName(v string)`

SetUiMetricName sets UiMetricName field to given value.

### HasUiMetricName

`func (o *KubernetesClusterChangingMetricsResponse) HasUiMetricName() bool`

HasUiMetricName returns a boolean if a field has been set.

### GetHumanLabel

`func (o *KubernetesClusterChangingMetricsResponse) GetHumanLabel() string`

GetHumanLabel returns the HumanLabel field if non-nil, zero value otherwise.

### GetHumanLabelOk

`func (o *KubernetesClusterChangingMetricsResponse) GetHumanLabelOk() (*string, bool)`

GetHumanLabelOk returns a tuple with the HumanLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanLabel

`func (o *KubernetesClusterChangingMetricsResponse) SetHumanLabel(v string)`

SetHumanLabel sets HumanLabel field to given value.

### HasHumanLabel

`func (o *KubernetesClusterChangingMetricsResponse) HasHumanLabel() bool`

HasHumanLabel returns a boolean if a field has been set.

### GetTimecodes

`func (o *KubernetesClusterChangingMetricsResponse) GetTimecodes() []string`

GetTimecodes returns the Timecodes field if non-nil, zero value otherwise.

### GetTimecodesOk

`func (o *KubernetesClusterChangingMetricsResponse) GetTimecodesOk() (*[]string, bool)`

GetTimecodesOk returns a tuple with the Timecodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecodes

`func (o *KubernetesClusterChangingMetricsResponse) SetTimecodes(v []string)`

SetTimecodes sets Timecodes field to given value.

### HasTimecodes

`func (o *KubernetesClusterChangingMetricsResponse) HasTimecodes() bool`

HasTimecodes returns a boolean if a field has been set.

### GetLinechartValues

`func (o *KubernetesClusterChangingMetricsResponse) GetLinechartValues() []float32`

GetLinechartValues returns the LinechartValues field if non-nil, zero value otherwise.

### GetLinechartValuesOk

`func (o *KubernetesClusterChangingMetricsResponse) GetLinechartValuesOk() (*[]float32, bool)`

GetLinechartValuesOk returns a tuple with the LinechartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinechartValues

`func (o *KubernetesClusterChangingMetricsResponse) SetLinechartValues(v []float32)`

SetLinechartValues sets LinechartValues field to given value.

### HasLinechartValues

`func (o *KubernetesClusterChangingMetricsResponse) HasLinechartValues() bool`

HasLinechartValues returns a boolean if a field has been set.

### GetTreemapValues

`func (o *KubernetesClusterChangingMetricsResponse) GetTreemapValues() float32`

GetTreemapValues returns the TreemapValues field if non-nil, zero value otherwise.

### GetTreemapValuesOk

`func (o *KubernetesClusterChangingMetricsResponse) GetTreemapValuesOk() (*float32, bool)`

GetTreemapValuesOk returns a tuple with the TreemapValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreemapValues

`func (o *KubernetesClusterChangingMetricsResponse) SetTreemapValues(v float32)`

SetTreemapValues sets TreemapValues field to given value.

### HasTreemapValues

`func (o *KubernetesClusterChangingMetricsResponse) HasTreemapValues() bool`

HasTreemapValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


