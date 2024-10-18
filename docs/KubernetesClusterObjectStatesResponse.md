# KubernetesClusterObjectStatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **string** | emma&#39;s internal metric name | [optional] 
**UiMetricGroup** | Pointer to **string** | UI group of metrics | [optional] 
**UiColorGroupId** | Pointer to **[]int32** | UI color for a status | [optional] 
**UiMetricName** | Pointer to **string** | UI metric name | [optional] 
**SubobjectName** | Pointer to **string** | Name of sub-object | [optional] 
**Value** | Pointer to **[]string** | Sub-object status | [optional] 

## Methods

### NewKubernetesClusterObjectStatesResponse

`func NewKubernetesClusterObjectStatesResponse() *KubernetesClusterObjectStatesResponse`

NewKubernetesClusterObjectStatesResponse instantiates a new KubernetesClusterObjectStatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterObjectStatesResponseWithDefaults

`func NewKubernetesClusterObjectStatesResponseWithDefaults() *KubernetesClusterObjectStatesResponse`

NewKubernetesClusterObjectStatesResponseWithDefaults instantiates a new KubernetesClusterObjectStatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *KubernetesClusterObjectStatesResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *KubernetesClusterObjectStatesResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *KubernetesClusterObjectStatesResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *KubernetesClusterObjectStatesResponse) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetUiMetricGroup

`func (o *KubernetesClusterObjectStatesResponse) GetUiMetricGroup() string`

GetUiMetricGroup returns the UiMetricGroup field if non-nil, zero value otherwise.

### GetUiMetricGroupOk

`func (o *KubernetesClusterObjectStatesResponse) GetUiMetricGroupOk() (*string, bool)`

GetUiMetricGroupOk returns a tuple with the UiMetricGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMetricGroup

`func (o *KubernetesClusterObjectStatesResponse) SetUiMetricGroup(v string)`

SetUiMetricGroup sets UiMetricGroup field to given value.

### HasUiMetricGroup

`func (o *KubernetesClusterObjectStatesResponse) HasUiMetricGroup() bool`

HasUiMetricGroup returns a boolean if a field has been set.

### GetUiColorGroupId

`func (o *KubernetesClusterObjectStatesResponse) GetUiColorGroupId() []int32`

GetUiColorGroupId returns the UiColorGroupId field if non-nil, zero value otherwise.

### GetUiColorGroupIdOk

`func (o *KubernetesClusterObjectStatesResponse) GetUiColorGroupIdOk() (*[]int32, bool)`

GetUiColorGroupIdOk returns a tuple with the UiColorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiColorGroupId

`func (o *KubernetesClusterObjectStatesResponse) SetUiColorGroupId(v []int32)`

SetUiColorGroupId sets UiColorGroupId field to given value.

### HasUiColorGroupId

`func (o *KubernetesClusterObjectStatesResponse) HasUiColorGroupId() bool`

HasUiColorGroupId returns a boolean if a field has been set.

### GetUiMetricName

`func (o *KubernetesClusterObjectStatesResponse) GetUiMetricName() string`

GetUiMetricName returns the UiMetricName field if non-nil, zero value otherwise.

### GetUiMetricNameOk

`func (o *KubernetesClusterObjectStatesResponse) GetUiMetricNameOk() (*string, bool)`

GetUiMetricNameOk returns a tuple with the UiMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMetricName

`func (o *KubernetesClusterObjectStatesResponse) SetUiMetricName(v string)`

SetUiMetricName sets UiMetricName field to given value.

### HasUiMetricName

`func (o *KubernetesClusterObjectStatesResponse) HasUiMetricName() bool`

HasUiMetricName returns a boolean if a field has been set.

### GetSubobjectName

`func (o *KubernetesClusterObjectStatesResponse) GetSubobjectName() string`

GetSubobjectName returns the SubobjectName field if non-nil, zero value otherwise.

### GetSubobjectNameOk

`func (o *KubernetesClusterObjectStatesResponse) GetSubobjectNameOk() (*string, bool)`

GetSubobjectNameOk returns a tuple with the SubobjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubobjectName

`func (o *KubernetesClusterObjectStatesResponse) SetSubobjectName(v string)`

SetSubobjectName sets SubobjectName field to given value.

### HasSubobjectName

`func (o *KubernetesClusterObjectStatesResponse) HasSubobjectName() bool`

HasSubobjectName returns a boolean if a field has been set.

### GetValue

`func (o *KubernetesClusterObjectStatesResponse) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesClusterObjectStatesResponse) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesClusterObjectStatesResponse) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesClusterObjectStatesResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


