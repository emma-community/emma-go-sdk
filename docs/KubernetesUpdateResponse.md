# KubernetesUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Kubernetes cluster ID | 
**CreatedAt** | **time.Time** | Cluster creation date and time | 
**CreatedById** | **int32** | ID of the user who created the cluster | 
**CreatedByName** | **string** | Name of the user who created the cluster | 
**ModifiedAt** | **time.Time** | Cluster last modification date and time | 
**ModifiedByName** | **string** | Name of the user who last modified the cluster | 
**ModifiedById** | **int32** | ID of the user who last modified the cluster | 
**Name** | **string** | Cluster name | 
**Version** | **string** | Kubernetes version | 
**DeploymentLocation** | **string** | Deployment location | 
**K8sConnectionType** | **string** | Kubernetes connection type | 
**DomainName** | Pointer to **string** | Cluster domain name | [optional] 
**Status** | **string** | Cluster status | 
**ControlPlaneStatus** | **string** | Control plane status | 
**Cost** | [**KubernetesUpdateResponseCost**](KubernetesUpdateResponseCost.md) |  | 
**NodeGroups** | [**[]KubernetesUpdateResponseNodeGroupsInner**](KubernetesUpdateResponseNodeGroupsInner.md) | List of cluster node groups | 
**AutoscalingConfigs** | [**[]KubernetesUpdateResponseAutoscalingConfigsInner**](KubernetesUpdateResponseAutoscalingConfigsInner.md) | Configurations of autoscaling groups | 

## Methods

### NewKubernetesUpdateResponse

`func NewKubernetesUpdateResponse(id int32, createdAt time.Time, createdById int32, createdByName string, modifiedAt time.Time, modifiedByName string, modifiedById int32, name string, version string, deploymentLocation string, k8sConnectionType string, status string, controlPlaneStatus string, cost KubernetesUpdateResponseCost, nodeGroups []KubernetesUpdateResponseNodeGroupsInner, autoscalingConfigs []KubernetesUpdateResponseAutoscalingConfigsInner, ) *KubernetesUpdateResponse`

NewKubernetesUpdateResponse instantiates a new KubernetesUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateResponseWithDefaults

`func NewKubernetesUpdateResponseWithDefaults() *KubernetesUpdateResponse`

NewKubernetesUpdateResponseWithDefaults instantiates a new KubernetesUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesUpdateResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesUpdateResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesUpdateResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *KubernetesUpdateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesUpdateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesUpdateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedById

`func (o *KubernetesUpdateResponse) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *KubernetesUpdateResponse) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *KubernetesUpdateResponse) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedByName

`func (o *KubernetesUpdateResponse) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *KubernetesUpdateResponse) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *KubernetesUpdateResponse) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.


### GetModifiedAt

`func (o *KubernetesUpdateResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *KubernetesUpdateResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *KubernetesUpdateResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedByName

`func (o *KubernetesUpdateResponse) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *KubernetesUpdateResponse) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *KubernetesUpdateResponse) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.


### GetModifiedById

`func (o *KubernetesUpdateResponse) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *KubernetesUpdateResponse) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *KubernetesUpdateResponse) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.


### GetName

`func (o *KubernetesUpdateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesUpdateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesUpdateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *KubernetesUpdateResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesUpdateResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesUpdateResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDeploymentLocation

`func (o *KubernetesUpdateResponse) GetDeploymentLocation() string`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *KubernetesUpdateResponse) GetDeploymentLocationOk() (*string, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *KubernetesUpdateResponse) SetDeploymentLocation(v string)`

SetDeploymentLocation sets DeploymentLocation field to given value.


### GetK8sConnectionType

`func (o *KubernetesUpdateResponse) GetK8sConnectionType() string`

GetK8sConnectionType returns the K8sConnectionType field if non-nil, zero value otherwise.

### GetK8sConnectionTypeOk

`func (o *KubernetesUpdateResponse) GetK8sConnectionTypeOk() (*string, bool)`

GetK8sConnectionTypeOk returns a tuple with the K8sConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sConnectionType

`func (o *KubernetesUpdateResponse) SetK8sConnectionType(v string)`

SetK8sConnectionType sets K8sConnectionType field to given value.


### GetDomainName

`func (o *KubernetesUpdateResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *KubernetesUpdateResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *KubernetesUpdateResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *KubernetesUpdateResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesUpdateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesUpdateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesUpdateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetControlPlaneStatus

`func (o *KubernetesUpdateResponse) GetControlPlaneStatus() string`

GetControlPlaneStatus returns the ControlPlaneStatus field if non-nil, zero value otherwise.

### GetControlPlaneStatusOk

`func (o *KubernetesUpdateResponse) GetControlPlaneStatusOk() (*string, bool)`

GetControlPlaneStatusOk returns a tuple with the ControlPlaneStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneStatus

`func (o *KubernetesUpdateResponse) SetControlPlaneStatus(v string)`

SetControlPlaneStatus sets ControlPlaneStatus field to given value.


### GetCost

`func (o *KubernetesUpdateResponse) GetCost() KubernetesUpdateResponseCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesUpdateResponse) GetCostOk() (*KubernetesUpdateResponseCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesUpdateResponse) SetCost(v KubernetesUpdateResponseCost)`

SetCost sets Cost field to given value.


### GetNodeGroups

`func (o *KubernetesUpdateResponse) GetNodeGroups() []KubernetesUpdateResponseNodeGroupsInner`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *KubernetesUpdateResponse) GetNodeGroupsOk() (*[]KubernetesUpdateResponseNodeGroupsInner, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *KubernetesUpdateResponse) SetNodeGroups(v []KubernetesUpdateResponseNodeGroupsInner)`

SetNodeGroups sets NodeGroups field to given value.


### GetAutoscalingConfigs

`func (o *KubernetesUpdateResponse) GetAutoscalingConfigs() []KubernetesUpdateResponseAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *KubernetesUpdateResponse) GetAutoscalingConfigsOk() (*[]KubernetesUpdateResponseAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *KubernetesUpdateResponse) SetAutoscalingConfigs(v []KubernetesUpdateResponseAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


