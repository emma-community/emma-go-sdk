# KubernetesClusterCurrentStateQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Type of kubernetes cluster object | 
**ObjectName** | **string** | Name of the object for which the user requests a list of metrics | 
**BreakdownLevel** | **string** | Kubernetes object breakdown | 
**CurrentStateMetrics** | **[]string** | List of metrics describing the current state | 
**CustomFilterState** | **string** | State of objects in the custom filter | 
**CustomFilterAvgCpuRule** | **string** | Rule for object filtering by CPU utilization | 
**CustomFilterAvgCpuValue** | **float32** | CPU utilization in the custom filter | 
**CustomFilterAvgMemoryRule** | **string** | Rule for object filtering by memory utilization | 
**CustomFilterAvgMemoryValue** | **float32** | Memory utilization in the custom filter | 
**CustomFilterAvgStorageRule** | **string** | Rule for object filtering by disk utilization | 
**CustomFilterAvgStorageValue** | **float32** | Disk utilization in the custom filter | 
**CustomFilterSubobjects** | **[]string** | List of subobject to explore | 

## Methods

### NewKubernetesClusterCurrentStateQueryFilters

`func NewKubernetesClusterCurrentStateQueryFilters(objectType string, objectName string, breakdownLevel string, currentStateMetrics []string, customFilterState string, customFilterAvgCpuRule string, customFilterAvgCpuValue float32, customFilterAvgMemoryRule string, customFilterAvgMemoryValue float32, customFilterAvgStorageRule string, customFilterAvgStorageValue float32, customFilterSubobjects []string, ) *KubernetesClusterCurrentStateQueryFilters`

NewKubernetesClusterCurrentStateQueryFilters instantiates a new KubernetesClusterCurrentStateQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterCurrentStateQueryFiltersWithDefaults

`func NewKubernetesClusterCurrentStateQueryFiltersWithDefaults() *KubernetesClusterCurrentStateQueryFilters`

NewKubernetesClusterCurrentStateQueryFiltersWithDefaults instantiates a new KubernetesClusterCurrentStateQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *KubernetesClusterCurrentStateQueryFilters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterCurrentStateQueryFilters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectName

`func (o *KubernetesClusterCurrentStateQueryFilters) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *KubernetesClusterCurrentStateQueryFilters) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetBreakdownLevel

`func (o *KubernetesClusterCurrentStateQueryFilters) GetBreakdownLevel() string`

GetBreakdownLevel returns the BreakdownLevel field if non-nil, zero value otherwise.

### GetBreakdownLevelOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetBreakdownLevelOk() (*string, bool)`

GetBreakdownLevelOk returns a tuple with the BreakdownLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownLevel

`func (o *KubernetesClusterCurrentStateQueryFilters) SetBreakdownLevel(v string)`

SetBreakdownLevel sets BreakdownLevel field to given value.


### GetCurrentStateMetrics

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCurrentStateMetrics() []string`

GetCurrentStateMetrics returns the CurrentStateMetrics field if non-nil, zero value otherwise.

### GetCurrentStateMetricsOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCurrentStateMetricsOk() (*[]string, bool)`

GetCurrentStateMetricsOk returns a tuple with the CurrentStateMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStateMetrics

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCurrentStateMetrics(v []string)`

SetCurrentStateMetrics sets CurrentStateMetrics field to given value.


### GetCustomFilterState

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterState() string`

GetCustomFilterState returns the CustomFilterState field if non-nil, zero value otherwise.

### GetCustomFilterStateOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterStateOk() (*string, bool)`

GetCustomFilterStateOk returns a tuple with the CustomFilterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterState

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterState(v string)`

SetCustomFilterState sets CustomFilterState field to given value.


### GetCustomFilterAvgCpuRule

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgCpuRule() string`

GetCustomFilterAvgCpuRule returns the CustomFilterAvgCpuRule field if non-nil, zero value otherwise.

### GetCustomFilterAvgCpuRuleOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgCpuRuleOk() (*string, bool)`

GetCustomFilterAvgCpuRuleOk returns a tuple with the CustomFilterAvgCpuRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgCpuRule

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterAvgCpuRule(v string)`

SetCustomFilterAvgCpuRule sets CustomFilterAvgCpuRule field to given value.


### GetCustomFilterAvgCpuValue

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgCpuValue() float32`

GetCustomFilterAvgCpuValue returns the CustomFilterAvgCpuValue field if non-nil, zero value otherwise.

### GetCustomFilterAvgCpuValueOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgCpuValueOk() (*float32, bool)`

GetCustomFilterAvgCpuValueOk returns a tuple with the CustomFilterAvgCpuValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgCpuValue

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterAvgCpuValue(v float32)`

SetCustomFilterAvgCpuValue sets CustomFilterAvgCpuValue field to given value.


### GetCustomFilterAvgMemoryRule

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgMemoryRule() string`

GetCustomFilterAvgMemoryRule returns the CustomFilterAvgMemoryRule field if non-nil, zero value otherwise.

### GetCustomFilterAvgMemoryRuleOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgMemoryRuleOk() (*string, bool)`

GetCustomFilterAvgMemoryRuleOk returns a tuple with the CustomFilterAvgMemoryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgMemoryRule

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterAvgMemoryRule(v string)`

SetCustomFilterAvgMemoryRule sets CustomFilterAvgMemoryRule field to given value.


### GetCustomFilterAvgMemoryValue

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgMemoryValue() float32`

GetCustomFilterAvgMemoryValue returns the CustomFilterAvgMemoryValue field if non-nil, zero value otherwise.

### GetCustomFilterAvgMemoryValueOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgMemoryValueOk() (*float32, bool)`

GetCustomFilterAvgMemoryValueOk returns a tuple with the CustomFilterAvgMemoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgMemoryValue

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterAvgMemoryValue(v float32)`

SetCustomFilterAvgMemoryValue sets CustomFilterAvgMemoryValue field to given value.


### GetCustomFilterAvgStorageRule

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgStorageRule() string`

GetCustomFilterAvgStorageRule returns the CustomFilterAvgStorageRule field if non-nil, zero value otherwise.

### GetCustomFilterAvgStorageRuleOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgStorageRuleOk() (*string, bool)`

GetCustomFilterAvgStorageRuleOk returns a tuple with the CustomFilterAvgStorageRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgStorageRule

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterAvgStorageRule(v string)`

SetCustomFilterAvgStorageRule sets CustomFilterAvgStorageRule field to given value.


### GetCustomFilterAvgStorageValue

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgStorageValue() float32`

GetCustomFilterAvgStorageValue returns the CustomFilterAvgStorageValue field if non-nil, zero value otherwise.

### GetCustomFilterAvgStorageValueOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterAvgStorageValueOk() (*float32, bool)`

GetCustomFilterAvgStorageValueOk returns a tuple with the CustomFilterAvgStorageValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterAvgStorageValue

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterAvgStorageValue(v float32)`

SetCustomFilterAvgStorageValue sets CustomFilterAvgStorageValue field to given value.


### GetCustomFilterSubobjects

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterSubobjects() []string`

GetCustomFilterSubobjects returns the CustomFilterSubobjects field if non-nil, zero value otherwise.

### GetCustomFilterSubobjectsOk

`func (o *KubernetesClusterCurrentStateQueryFilters) GetCustomFilterSubobjectsOk() (*[]string, bool)`

GetCustomFilterSubobjectsOk returns a tuple with the CustomFilterSubobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterSubobjects

`func (o *KubernetesClusterCurrentStateQueryFilters) SetCustomFilterSubobjects(v []string)`

SetCustomFilterSubobjects sets CustomFilterSubobjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


