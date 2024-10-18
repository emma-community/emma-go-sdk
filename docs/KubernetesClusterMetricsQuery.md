# KubernetesClusterMetricsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**CoreClusterId** | **int32** | ID of Kubernetes cluster | 
**Filters** | [**KubernetesClusterMetricsQueryFilters**](KubernetesClusterMetricsQueryFilters.md) |  | 

## Methods

### NewKubernetesClusterMetricsQuery

`func NewKubernetesClusterMetricsQuery(datasetName string, coreClusterId int32, filters KubernetesClusterMetricsQueryFilters, ) *KubernetesClusterMetricsQuery`

NewKubernetesClusterMetricsQuery instantiates a new KubernetesClusterMetricsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterMetricsQueryWithDefaults

`func NewKubernetesClusterMetricsQueryWithDefaults() *KubernetesClusterMetricsQuery`

NewKubernetesClusterMetricsQueryWithDefaults instantiates a new KubernetesClusterMetricsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *KubernetesClusterMetricsQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *KubernetesClusterMetricsQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *KubernetesClusterMetricsQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetCoreClusterId

`func (o *KubernetesClusterMetricsQuery) GetCoreClusterId() int32`

GetCoreClusterId returns the CoreClusterId field if non-nil, zero value otherwise.

### GetCoreClusterIdOk

`func (o *KubernetesClusterMetricsQuery) GetCoreClusterIdOk() (*int32, bool)`

GetCoreClusterIdOk returns a tuple with the CoreClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreClusterId

`func (o *KubernetesClusterMetricsQuery) SetCoreClusterId(v int32)`

SetCoreClusterId sets CoreClusterId field to given value.


### GetFilters

`func (o *KubernetesClusterMetricsQuery) GetFilters() KubernetesClusterMetricsQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *KubernetesClusterMetricsQuery) GetFiltersOk() (*KubernetesClusterMetricsQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *KubernetesClusterMetricsQuery) SetFilters(v KubernetesClusterMetricsQueryFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


