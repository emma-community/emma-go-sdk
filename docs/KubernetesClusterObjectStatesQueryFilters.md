# KubernetesClusterObjectStatesQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Type of kubernetes cluster object | 
**ObjectName** | **string** | Name of the object for which the user requests a list of metrics | 
**BreakdownLevel** | **string** | Kubernetes object breakdown | 
**ObjectStatesMetrics** | **[]string** | List of metrics describing state of objects in the cluster | 

## Methods

### NewKubernetesClusterObjectStatesQueryFilters

`func NewKubernetesClusterObjectStatesQueryFilters(objectType string, objectName string, breakdownLevel string, objectStatesMetrics []string, ) *KubernetesClusterObjectStatesQueryFilters`

NewKubernetesClusterObjectStatesQueryFilters instantiates a new KubernetesClusterObjectStatesQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterObjectStatesQueryFiltersWithDefaults

`func NewKubernetesClusterObjectStatesQueryFiltersWithDefaults() *KubernetesClusterObjectStatesQueryFilters`

NewKubernetesClusterObjectStatesQueryFiltersWithDefaults instantiates a new KubernetesClusterObjectStatesQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *KubernetesClusterObjectStatesQueryFilters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterObjectStatesQueryFilters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterObjectStatesQueryFilters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectName

`func (o *KubernetesClusterObjectStatesQueryFilters) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *KubernetesClusterObjectStatesQueryFilters) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *KubernetesClusterObjectStatesQueryFilters) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetBreakdownLevel

`func (o *KubernetesClusterObjectStatesQueryFilters) GetBreakdownLevel() string`

GetBreakdownLevel returns the BreakdownLevel field if non-nil, zero value otherwise.

### GetBreakdownLevelOk

`func (o *KubernetesClusterObjectStatesQueryFilters) GetBreakdownLevelOk() (*string, bool)`

GetBreakdownLevelOk returns a tuple with the BreakdownLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownLevel

`func (o *KubernetesClusterObjectStatesQueryFilters) SetBreakdownLevel(v string)`

SetBreakdownLevel sets BreakdownLevel field to given value.


### GetObjectStatesMetrics

`func (o *KubernetesClusterObjectStatesQueryFilters) GetObjectStatesMetrics() []string`

GetObjectStatesMetrics returns the ObjectStatesMetrics field if non-nil, zero value otherwise.

### GetObjectStatesMetricsOk

`func (o *KubernetesClusterObjectStatesQueryFilters) GetObjectStatesMetricsOk() (*[]string, bool)`

GetObjectStatesMetricsOk returns a tuple with the ObjectStatesMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStatesMetrics

`func (o *KubernetesClusterObjectStatesQueryFilters) SetObjectStatesMetrics(v []string)`

SetObjectStatesMetrics sets ObjectStatesMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


