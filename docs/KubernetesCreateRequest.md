# KubernetesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentLocation** | **string** | Deployment location | 
**K8sVersion** | Pointer to **string** | Kubernetes version | [optional] 
**ProjectName** | Pointer to **string** | Project name | [optional] 
**Name** | **string** | Kubernetes cluster name | 
**K8sConnectionType** | **string** | Kubernetes connection type | 
**WorkerNodes** | [**[]KubernetesCreateRequestWorkerNodesInner**](KubernetesCreateRequestWorkerNodesInner.md) | List of worker nodes | 
**AutoscalingConfigs** | Pointer to [**[]KubernetesCreateRequestAutoscalingConfigsInner**](KubernetesCreateRequestAutoscalingConfigsInner.md) | Configurations of the autoscaling group | [optional] 

## Methods

### NewKubernetesCreateRequest

`func NewKubernetesCreateRequest(deploymentLocation string, name string, k8sConnectionType string, workerNodes []KubernetesCreateRequestWorkerNodesInner, ) *KubernetesCreateRequest`

NewKubernetesCreateRequest instantiates a new KubernetesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCreateRequestWithDefaults

`func NewKubernetesCreateRequestWithDefaults() *KubernetesCreateRequest`

NewKubernetesCreateRequestWithDefaults instantiates a new KubernetesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentLocation

`func (o *KubernetesCreateRequest) GetDeploymentLocation() string`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *KubernetesCreateRequest) GetDeploymentLocationOk() (*string, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *KubernetesCreateRequest) SetDeploymentLocation(v string)`

SetDeploymentLocation sets DeploymentLocation field to given value.


### GetK8sVersion

`func (o *KubernetesCreateRequest) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesCreateRequest) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesCreateRequest) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *KubernetesCreateRequest) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetProjectName

`func (o *KubernetesCreateRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *KubernetesCreateRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *KubernetesCreateRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *KubernetesCreateRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetName

`func (o *KubernetesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetK8sConnectionType

`func (o *KubernetesCreateRequest) GetK8sConnectionType() string`

GetK8sConnectionType returns the K8sConnectionType field if non-nil, zero value otherwise.

### GetK8sConnectionTypeOk

`func (o *KubernetesCreateRequest) GetK8sConnectionTypeOk() (*string, bool)`

GetK8sConnectionTypeOk returns a tuple with the K8sConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sConnectionType

`func (o *KubernetesCreateRequest) SetK8sConnectionType(v string)`

SetK8sConnectionType sets K8sConnectionType field to given value.


### GetWorkerNodes

`func (o *KubernetesCreateRequest) GetWorkerNodes() []KubernetesCreateRequestWorkerNodesInner`

GetWorkerNodes returns the WorkerNodes field if non-nil, zero value otherwise.

### GetWorkerNodesOk

`func (o *KubernetesCreateRequest) GetWorkerNodesOk() (*[]KubernetesCreateRequestWorkerNodesInner, bool)`

GetWorkerNodesOk returns a tuple with the WorkerNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNodes

`func (o *KubernetesCreateRequest) SetWorkerNodes(v []KubernetesCreateRequestWorkerNodesInner)`

SetWorkerNodes sets WorkerNodes field to given value.


### GetAutoscalingConfigs

`func (o *KubernetesCreateRequest) GetAutoscalingConfigs() []KubernetesCreateRequestAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *KubernetesCreateRequest) GetAutoscalingConfigsOk() (*[]KubernetesCreateRequestAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *KubernetesCreateRequest) SetAutoscalingConfigs(v []KubernetesCreateRequestAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.

### HasAutoscalingConfigs

`func (o *KubernetesCreateRequest) HasAutoscalingConfigs() bool`

HasAutoscalingConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


