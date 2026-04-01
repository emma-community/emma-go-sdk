# KubernetesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerNodes** | [**[]KubernetesUpdateRequestWorkerNodesInner**](KubernetesUpdateRequestWorkerNodesInner.md) | List of the worker nodes | 
**AutoscalingConfigs** | Pointer to [**[]KubernetesUpdateRequestAutoscalingConfigsInner**](KubernetesUpdateRequestAutoscalingConfigsInner.md) | Configurations of the autoscaling group | [optional] 

## Methods

### NewKubernetesUpdateRequest

`func NewKubernetesUpdateRequest(workerNodes []KubernetesUpdateRequestWorkerNodesInner, ) *KubernetesUpdateRequest`

NewKubernetesUpdateRequest instantiates a new KubernetesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateRequestWithDefaults

`func NewKubernetesUpdateRequestWithDefaults() *KubernetesUpdateRequest`

NewKubernetesUpdateRequestWithDefaults instantiates a new KubernetesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerNodes

`func (o *KubernetesUpdateRequest) GetWorkerNodes() []KubernetesUpdateRequestWorkerNodesInner`

GetWorkerNodes returns the WorkerNodes field if non-nil, zero value otherwise.

### GetWorkerNodesOk

`func (o *KubernetesUpdateRequest) GetWorkerNodesOk() (*[]KubernetesUpdateRequestWorkerNodesInner, bool)`

GetWorkerNodesOk returns a tuple with the WorkerNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNodes

`func (o *KubernetesUpdateRequest) SetWorkerNodes(v []KubernetesUpdateRequestWorkerNodesInner)`

SetWorkerNodes sets WorkerNodes field to given value.


### GetAutoscalingConfigs

`func (o *KubernetesUpdateRequest) GetAutoscalingConfigs() []KubernetesUpdateRequestAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *KubernetesUpdateRequest) GetAutoscalingConfigsOk() (*[]KubernetesUpdateRequestAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *KubernetesUpdateRequest) SetAutoscalingConfigs(v []KubernetesUpdateRequestAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.

### HasAutoscalingConfigs

`func (o *KubernetesUpdateRequest) HasAutoscalingConfigs() bool`

HasAutoscalingConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


