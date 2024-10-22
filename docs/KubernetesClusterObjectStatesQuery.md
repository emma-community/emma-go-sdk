# KubernetesClusterObjectStatesQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**CoreClusterId** | **int32** | ID of Kubernetes cluster | 
**Filters** | [**KubernetesClusterObjectStatesQueryFilters**](KubernetesClusterObjectStatesQueryFilters.md) |  | 

## Methods

### NewKubernetesClusterObjectStatesQuery

`func NewKubernetesClusterObjectStatesQuery(datasetName string, coreClusterId int32, filters KubernetesClusterObjectStatesQueryFilters, ) *KubernetesClusterObjectStatesQuery`

NewKubernetesClusterObjectStatesQuery instantiates a new KubernetesClusterObjectStatesQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterObjectStatesQueryWithDefaults

`func NewKubernetesClusterObjectStatesQueryWithDefaults() *KubernetesClusterObjectStatesQuery`

NewKubernetesClusterObjectStatesQueryWithDefaults instantiates a new KubernetesClusterObjectStatesQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *KubernetesClusterObjectStatesQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *KubernetesClusterObjectStatesQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *KubernetesClusterObjectStatesQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetCoreClusterId

`func (o *KubernetesClusterObjectStatesQuery) GetCoreClusterId() int32`

GetCoreClusterId returns the CoreClusterId field if non-nil, zero value otherwise.

### GetCoreClusterIdOk

`func (o *KubernetesClusterObjectStatesQuery) GetCoreClusterIdOk() (*int32, bool)`

GetCoreClusterIdOk returns a tuple with the CoreClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreClusterId

`func (o *KubernetesClusterObjectStatesQuery) SetCoreClusterId(v int32)`

SetCoreClusterId sets CoreClusterId field to given value.


### GetFilters

`func (o *KubernetesClusterObjectStatesQuery) GetFilters() KubernetesClusterObjectStatesQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *KubernetesClusterObjectStatesQuery) GetFiltersOk() (*KubernetesClusterObjectStatesQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *KubernetesClusterObjectStatesQuery) SetFilters(v KubernetesClusterObjectStatesQueryFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


