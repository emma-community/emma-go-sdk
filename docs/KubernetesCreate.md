# KubernetesCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Kubernetes cluster name | 
**DeploymentLocation** | **string** | Deployment region of the Kubernetes cluster. Currently, only Europe (eu) is available. | 
**WorkerNodes** | [**[]KubernetesCreateWorkerNodesInner**](KubernetesCreateWorkerNodesInner.md) | List of the worker nodes | 
**AutoscalingConfigs** | Pointer to [**[]KubernetesCreateAutoscalingConfigsInner**](KubernetesCreateAutoscalingConfigsInner.md) | Configurations of the autoscaling group | [optional] 

## Methods

### NewKubernetesCreate

`func NewKubernetesCreate(name string, deploymentLocation string, workerNodes []KubernetesCreateWorkerNodesInner, ) *KubernetesCreate`

NewKubernetesCreate instantiates a new KubernetesCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCreateWithDefaults

`func NewKubernetesCreateWithDefaults() *KubernetesCreate`

NewKubernetesCreateWithDefaults instantiates a new KubernetesCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDeploymentLocation

`func (o *KubernetesCreate) GetDeploymentLocation() string`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *KubernetesCreate) GetDeploymentLocationOk() (*string, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *KubernetesCreate) SetDeploymentLocation(v string)`

SetDeploymentLocation sets DeploymentLocation field to given value.


### GetWorkerNodes

`func (o *KubernetesCreate) GetWorkerNodes() []KubernetesCreateWorkerNodesInner`

GetWorkerNodes returns the WorkerNodes field if non-nil, zero value otherwise.

### GetWorkerNodesOk

`func (o *KubernetesCreate) GetWorkerNodesOk() (*[]KubernetesCreateWorkerNodesInner, bool)`

GetWorkerNodesOk returns a tuple with the WorkerNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNodes

`func (o *KubernetesCreate) SetWorkerNodes(v []KubernetesCreateWorkerNodesInner)`

SetWorkerNodes sets WorkerNodes field to given value.


### GetAutoscalingConfigs

`func (o *KubernetesCreate) GetAutoscalingConfigs() []KubernetesCreateAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *KubernetesCreate) GetAutoscalingConfigsOk() (*[]KubernetesCreateAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *KubernetesCreate) SetAutoscalingConfigs(v []KubernetesCreateAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.

### HasAutoscalingConfigs

`func (o *KubernetesCreate) HasAutoscalingConfigs() bool`

HasAutoscalingConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


