# KubernetesClusterChangingMetricsQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Type of kubernetes cluster object | 
**ObjectName** | **string** | Name of the object for which the user requests a list of metrics | 
**BreakdownLevel** | **string** | Kubernetes object breakdown | 
**ChangingMetrics** | **[]string** | List of metrics describing changes in the cluster | 
**Timespan** | **string** | Requested period of statistics | 
**CustomFilterState** | **NullableString** | State of objects in the custom filter | 
**CustomFilterAvgCpuRule** | **NullableString** | Rule for object filtering by CPU utilization | 
**CustomFilterAvgCpuValue** | **NullableFloat32** | CPU utilization in the custom filter | 
**CustomFilterAvgMemoryRule** | **NullableString** | Rule for object filtering by memory utilization | 
**CustomFilterAvgMemoryValue** | **NullableFloat32** | Memory utilization in the custom filter | 
**CustomFilterAvgStorageRule** | **NullableString** | Rule for object filtering by disk utilization | 
**CustomFilterAvgStorageValue** | **NullableFloat32** | Disk utilization in the custom filter | 
**CustomFilterSubobjects** | **[]string** | List of subobject to explore | 

## Methods

### NewKubernetesClusterChangingMetricsQueryFilters

`func NewKubernetesClusterChangingMetricsQueryFilters(objectType string, objectName string, breakdownLevel string, changingMetrics []string, timespan string, customFilterState NullableString, customFilterAvgCpuRule NullableString, customFilterAvgCpuValue NullableFloat32, customFilterAvgMemoryRule NullableString, customFilterAvgMemoryValue NullableFloat32, customFilterAvgStorageRule NullableString, customFilterAvgStorageValue NullableFloat32, customFilterSubobjects []string, ) *KubernetesClusterChangingMetricsQueryFilters`

NewKubernetesClusterChangingMetricsQueryFilters instantiates a new KubernetesClusterChangingMetricsQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterChangingMetricsQueryFiltersWithDefaults

`func NewKubernetesClusterChangingMetricsQueryFiltersWithDefaults() *KubernetesClusterChangingMetricsQueryFilters`

NewKubernetesClusterChangingMetricsQueryFiltersWithDefaults instantiates a new KubernetesClusterChangingMetricsQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectName

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetBreakdownLevel

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetBreakdownLevel() string`

GetBreakdownLevel returns the BreakdownLevel field if non-nil, zero value otherwise.

### GetBreakdownLevelOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetBreakdownLevelOk() (*string, bool)`

GetBreakdownLevelOk returns a tuple with the BreakdownLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownLevel

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetBreakdownLevel(v string)`

SetBreakdownLevel sets BreakdownLevel field to given value.


### GetChangingMetrics

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetChangingMetrics() []string`

GetChangingMetrics returns the ChangingMetrics field if non-nil, zero value otherwise.

### GetChangingMetricsOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetChangingMetricsOk() (*[]string, bool)`

GetChangingMetricsOk returns a tuple with the ChangingMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangingMetrics

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetChangingMetrics(v []string)`

SetChangingMetrics sets ChangingMetrics field to given value.


### GetTimespan

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetTimespan() string`

GetTimespan returns the Timespan field if non-nil, zero value otherwise.

### GetTimespanOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetTimespanOk() (*string, bool)`

GetTimespanOk returns a tuple with the Timespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimespan

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetTimespan(v string)`

SetTimespan sets Timespan field to given value.


### GetCustomFilterState

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterState() string`

GetCustomFilterState returns the CustomFilterState field if non-nil, zero value otherwise.

### GetCustomFilterStateOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterStateOk() (*string, bool)`

GetCustomFilterStateOk returns a tuple with the CustomFilterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterState

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterState(v string)`

SetCustomFilterState sets CustomFilterState field to given value.


### SetCustomFilterStateNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterStateNil(b bool)`

 SetCustomFilterStateNil sets the value for CustomFilterState to be an explicit nil

### UnsetCustomFilterState
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterState()`

UnsetCustomFilterState ensures that no value is present for CustomFilterState, not even an explicit nil
### GetCustomFilterAvgCpuRule

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgCpuRule() string`

GetCustomFilterAvgCpuRule returns the CustomFilterAvgCpuRule field if non-nil, zero value otherwise.

### GetCustomFilterAvgCpuRuleOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgCpuRuleOk() (*string, bool)`

GetCustomFilterAvgCpuRuleOk returns a tuple with the CustomFilterAvgCpuRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgCpuRule

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgCpuRule(v string)`

SetCustomFilterAvgCpuRule sets CustomFilterAvgCpuRule field to given value.


### SetCustomFilterAvgCpuRuleNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgCpuRuleNil(b bool)`

 SetCustomFilterAvgCpuRuleNil sets the value for CustomFilterAvgCpuRule to be an explicit nil

### UnsetCustomFilterAvgCpuRule
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterAvgCpuRule()`

UnsetCustomFilterAvgCpuRule ensures that no value is present for CustomFilterAvgCpuRule, not even an explicit nil
### GetCustomFilterAvgCpuValue

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgCpuValue() float32`

GetCustomFilterAvgCpuValue returns the CustomFilterAvgCpuValue field if non-nil, zero value otherwise.

### GetCustomFilterAvgCpuValueOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgCpuValueOk() (*float32, bool)`

GetCustomFilterAvgCpuValueOk returns a tuple with the CustomFilterAvgCpuValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgCpuValue

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgCpuValue(v float32)`

SetCustomFilterAvgCpuValue sets CustomFilterAvgCpuValue field to given value.


### SetCustomFilterAvgCpuValueNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgCpuValueNil(b bool)`

 SetCustomFilterAvgCpuValueNil sets the value for CustomFilterAvgCpuValue to be an explicit nil

### UnsetCustomFilterAvgCpuValue
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterAvgCpuValue()`

UnsetCustomFilterAvgCpuValue ensures that no value is present for CustomFilterAvgCpuValue, not even an explicit nil
### GetCustomFilterAvgMemoryRule

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgMemoryRule() string`

GetCustomFilterAvgMemoryRule returns the CustomFilterAvgMemoryRule field if non-nil, zero value otherwise.

### GetCustomFilterAvgMemoryRuleOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgMemoryRuleOk() (*string, bool)`

GetCustomFilterAvgMemoryRuleOk returns a tuple with the CustomFilterAvgMemoryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgMemoryRule

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgMemoryRule(v string)`

SetCustomFilterAvgMemoryRule sets CustomFilterAvgMemoryRule field to given value.


### SetCustomFilterAvgMemoryRuleNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgMemoryRuleNil(b bool)`

 SetCustomFilterAvgMemoryRuleNil sets the value for CustomFilterAvgMemoryRule to be an explicit nil

### UnsetCustomFilterAvgMemoryRule
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterAvgMemoryRule()`

UnsetCustomFilterAvgMemoryRule ensures that no value is present for CustomFilterAvgMemoryRule, not even an explicit nil
### GetCustomFilterAvgMemoryValue

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgMemoryValue() float32`

GetCustomFilterAvgMemoryValue returns the CustomFilterAvgMemoryValue field if non-nil, zero value otherwise.

### GetCustomFilterAvgMemoryValueOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgMemoryValueOk() (*float32, bool)`

GetCustomFilterAvgMemoryValueOk returns a tuple with the CustomFilterAvgMemoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgMemoryValue

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgMemoryValue(v float32)`

SetCustomFilterAvgMemoryValue sets CustomFilterAvgMemoryValue field to given value.


### SetCustomFilterAvgMemoryValueNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgMemoryValueNil(b bool)`

 SetCustomFilterAvgMemoryValueNil sets the value for CustomFilterAvgMemoryValue to be an explicit nil

### UnsetCustomFilterAvgMemoryValue
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterAvgMemoryValue()`

UnsetCustomFilterAvgMemoryValue ensures that no value is present for CustomFilterAvgMemoryValue, not even an explicit nil
### GetCustomFilterAvgStorageRule

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgStorageRule() string`

GetCustomFilterAvgStorageRule returns the CustomFilterAvgStorageRule field if non-nil, zero value otherwise.

### GetCustomFilterAvgStorageRuleOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgStorageRuleOk() (*string, bool)`

GetCustomFilterAvgStorageRuleOk returns a tuple with the CustomFilterAvgStorageRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgStorageRule

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgStorageRule(v string)`

SetCustomFilterAvgStorageRule sets CustomFilterAvgStorageRule field to given value.


### SetCustomFilterAvgStorageRuleNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgStorageRuleNil(b bool)`

 SetCustomFilterAvgStorageRuleNil sets the value for CustomFilterAvgStorageRule to be an explicit nil

### UnsetCustomFilterAvgStorageRule
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterAvgStorageRule()`

UnsetCustomFilterAvgStorageRule ensures that no value is present for CustomFilterAvgStorageRule, not even an explicit nil
### GetCustomFilterAvgStorageValue

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgStorageValue() float32`

GetCustomFilterAvgStorageValue returns the CustomFilterAvgStorageValue field if non-nil, zero value otherwise.

### GetCustomFilterAvgStorageValueOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterAvgStorageValueOk() (*float32, bool)`

GetCustomFilterAvgStorageValueOk returns a tuple with the CustomFilterAvgStorageValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgStorageValue

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgStorageValue(v float32)`

SetCustomFilterAvgStorageValue sets CustomFilterAvgStorageValue field to given value.


### SetCustomFilterAvgStorageValueNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterAvgStorageValueNil(b bool)`

 SetCustomFilterAvgStorageValueNil sets the value for CustomFilterAvgStorageValue to be an explicit nil

### UnsetCustomFilterAvgStorageValue
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterAvgStorageValue()`

UnsetCustomFilterAvgStorageValue ensures that no value is present for CustomFilterAvgStorageValue, not even an explicit nil
### GetCustomFilterSubobjects

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterSubobjects() []string`

GetCustomFilterSubobjects returns the CustomFilterSubobjects field if non-nil, zero value otherwise.

### GetCustomFilterSubobjectsOk

`func (o *KubernetesClusterChangingMetricsQueryFilters) GetCustomFilterSubobjectsOk() (*[]string, bool)`

GetCustomFilterSubobjectsOk returns a tuple with the CustomFilterSubobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterSubobjects

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterSubobjects(v []string)`

SetCustomFilterSubobjects sets CustomFilterSubobjects field to given value.


### SetCustomFilterSubobjectsNil

`func (o *KubernetesClusterChangingMetricsQueryFilters) SetCustomFilterSubobjectsNil(b bool)`

 SetCustomFilterSubobjectsNil sets the value for CustomFilterSubobjects to be an explicit nil

### UnsetCustomFilterSubobjects
`func (o *KubernetesClusterChangingMetricsQueryFilters) UnsetCustomFilterSubobjects()`

UnsetCustomFilterSubobjects ensures that no value is present for CustomFilterSubobjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


