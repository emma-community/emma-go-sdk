# KubernetesClusterCurrentStateQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** | Query name | 
**CoreClusterId** | **int32** | ID of Kubernetes cluster | 
**Filters** | [**KubernetesClusterCurrentStateQueryFilters**](KubernetesClusterCurrentStateQueryFilters.md) |  | 

## Methods

### NewKubernetesClusterCurrentStateQuery

`func NewKubernetesClusterCurrentStateQuery(datasetName string, coreClusterId int32, filters KubernetesClusterCurrentStateQueryFilters, ) *KubernetesClusterCurrentStateQuery`

NewKubernetesClusterCurrentStateQuery instantiates a new KubernetesClusterCurrentStateQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterCurrentStateQueryWithDefaults

`func NewKubernetesClusterCurrentStateQueryWithDefaults() *KubernetesClusterCurrentStateQuery`

NewKubernetesClusterCurrentStateQueryWithDefaults instantiates a new KubernetesClusterCurrentStateQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *KubernetesClusterCurrentStateQuery) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *KubernetesClusterCurrentStateQuery) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *KubernetesClusterCurrentStateQuery) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetCoreClusterId

`func (o *KubernetesClusterCurrentStateQuery) GetCoreClusterId() int32`

GetCoreClusterId returns the CoreClusterId field if non-nil, zero value otherwise.

### GetCoreClusterIdOk

`func (o *KubernetesClusterCurrentStateQuery) GetCoreClusterIdOk() (*int32, bool)`

GetCoreClusterIdOk returns a tuple with the CoreClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreClusterId

`func (o *KubernetesClusterCurrentStateQuery) SetCoreClusterId(v int32)`

SetCoreClusterId sets CoreClusterId field to given value.


### GetFilters

`func (o *KubernetesClusterCurrentStateQuery) GetFilters() KubernetesClusterCurrentStateQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *KubernetesClusterCurrentStateQuery) GetFiltersOk() (*KubernetesClusterCurrentStateQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *KubernetesClusterCurrentStateQuery) SetFilters(v KubernetesClusterCurrentStateQueryFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


