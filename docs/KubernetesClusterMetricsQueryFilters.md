# KubernetesClusterMetricsQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Type of kubernetes cluster object | 
**ObjectName** | **string** | Name of the object for which the user requests a list of metrics | 
**BreakdownLevel** | **string** | Kubernetes object breakdown | 

## Methods

### NewKubernetesClusterMetricsQueryFilters

`func NewKubernetesClusterMetricsQueryFilters(objectType string, objectName string, breakdownLevel string, ) *KubernetesClusterMetricsQueryFilters`

NewKubernetesClusterMetricsQueryFilters instantiates a new KubernetesClusterMetricsQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterMetricsQueryFiltersWithDefaults

`func NewKubernetesClusterMetricsQueryFiltersWithDefaults() *KubernetesClusterMetricsQueryFilters`

NewKubernetesClusterMetricsQueryFiltersWithDefaults instantiates a new KubernetesClusterMetricsQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *KubernetesClusterMetricsQueryFilters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterMetricsQueryFilters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterMetricsQueryFilters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectName

`func (o *KubernetesClusterMetricsQueryFilters) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *KubernetesClusterMetricsQueryFilters) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *KubernetesClusterMetricsQueryFilters) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetBreakdownLevel

`func (o *KubernetesClusterMetricsQueryFilters) GetBreakdownLevel() string`

GetBreakdownLevel returns the BreakdownLevel field if non-nil, zero value otherwise.

### GetBreakdownLevelOk

`func (o *KubernetesClusterMetricsQueryFilters) GetBreakdownLevelOk() (*string, bool)`

GetBreakdownLevelOk returns a tuple with the BreakdownLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownLevel

`func (o *KubernetesClusterMetricsQueryFilters) SetBreakdownLevel(v string)`

SetBreakdownLevel sets BreakdownLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


