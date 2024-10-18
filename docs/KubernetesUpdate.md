# KubernetesUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerNodes** | [**[]KubernetesUpdateWorkerNodesInner**](KubernetesUpdateWorkerNodesInner.md) | List of the worker nodes | 
**AutoscalingConfigs** | Pointer to [**[]KubernetesCreateAutoscalingConfigsInner**](KubernetesCreateAutoscalingConfigsInner.md) | Configurations of the autoscaling group | [optional] 

## Methods

### NewKubernetesUpdate

`func NewKubernetesUpdate(workerNodes []KubernetesUpdateWorkerNodesInner, ) *KubernetesUpdate`

NewKubernetesUpdate instantiates a new KubernetesUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateWithDefaults

`func NewKubernetesUpdateWithDefaults() *KubernetesUpdate`

NewKubernetesUpdateWithDefaults instantiates a new KubernetesUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerNodes

`func (o *KubernetesUpdate) GetWorkerNodes() []KubernetesUpdateWorkerNodesInner`

GetWorkerNodes returns the WorkerNodes field if non-nil, zero value otherwise.

### GetWorkerNodesOk

`func (o *KubernetesUpdate) GetWorkerNodesOk() (*[]KubernetesUpdateWorkerNodesInner, bool)`

GetWorkerNodesOk returns a tuple with the WorkerNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNodes

`func (o *KubernetesUpdate) SetWorkerNodes(v []KubernetesUpdateWorkerNodesInner)`

SetWorkerNodes sets WorkerNodes field to given value.


### GetAutoscalingConfigs

`func (o *KubernetesUpdate) GetAutoscalingConfigs() []KubernetesCreateAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *KubernetesUpdate) GetAutoscalingConfigsOk() (*[]KubernetesCreateAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *KubernetesUpdate) SetAutoscalingConfigs(v []KubernetesCreateAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.

### HasAutoscalingConfigs

`func (o *KubernetesUpdate) HasAutoscalingConfigs() bool`

HasAutoscalingConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


