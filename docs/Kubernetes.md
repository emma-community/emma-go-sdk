# Kubernetes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the Kubernetes cluster | [optional] 
**CreatedAt** | Pointer to **string** | Date and time of the Kubernetes cluster creation | [optional] 
**CreatedByName** | Pointer to **string** | Name of the user who created the Kubernetes cluster | [optional] 
**CreatedById** | Pointer to **int32** | ID of the user who created the Kubernetes cluster | [optional] 
**ModifiedAt** | Pointer to **string** | Date and time when the Kubernetes cluster was last edited | [optional] 
**ModifiedByName** | Pointer to **string** | Name of the user who last edited the Kubernetes cluster | [optional] 
**ModifiedById** | Pointer to **int32** | ID of the user who last edited the Kubernetes cluster | [optional] 
**Name** | Pointer to **string** | Name of the Kubernetes cluster | [optional] 
**Version** | Pointer to **string** | Version of the Kubernetes cluster | [optional] 
**DeploymentLocation** | Pointer to **string** | Deployment region of the Kubernetes cluster. Currently, Europe (eu) and North America (n_america) is available. | [optional] 
**K8sConnectionType** | Pointer to **string** | Specifies the network connectivity type for the cluster. Immutable: set at creation only. To change, delete and create a new cluster. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the Kubernetes cluster | [optional] 
**Cost** | Pointer to [**KubernetesCost**](KubernetesCost.md) |  | [optional] 
**ControlPlaneStatus** | Pointer to **string** | Control Plane Status. The Control Plane status provides information about the health and functionality of various components that are part of the Control Plane. | [optional] 
**DomainName** | Pointer to **string** | Domain attached to the Kubernetes cluster | [optional] 
**NodeGroups** | Pointer to [**[]KubernetesNodeGroupsInner**](KubernetesNodeGroupsInner.md) | List of the worker node groups | [optional] 
**AutoscalingConfigs** | Pointer to [**[]KubernetesAutoscalingConfigsInner**](KubernetesAutoscalingConfigsInner.md) | Configurations of the autoscaling group | [optional] 

## Methods

### NewKubernetes

`func NewKubernetes() *Kubernetes`

NewKubernetes instantiates a new Kubernetes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesWithDefaults

`func NewKubernetesWithDefaults() *Kubernetes`

NewKubernetesWithDefaults instantiates a new Kubernetes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kubernetes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kubernetes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kubernetes) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Kubernetes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Kubernetes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Kubernetes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Kubernetes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Kubernetes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *Kubernetes) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *Kubernetes) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *Kubernetes) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *Kubernetes) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedById

`func (o *Kubernetes) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Kubernetes) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Kubernetes) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Kubernetes) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Kubernetes) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Kubernetes) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Kubernetes) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Kubernetes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *Kubernetes) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *Kubernetes) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *Kubernetes) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *Kubernetes) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *Kubernetes) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *Kubernetes) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *Kubernetes) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *Kubernetes) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *Kubernetes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Kubernetes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Kubernetes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Kubernetes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Kubernetes) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Kubernetes) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Kubernetes) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Kubernetes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentLocation

`func (o *Kubernetes) GetDeploymentLocation() string`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *Kubernetes) GetDeploymentLocationOk() (*string, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *Kubernetes) SetDeploymentLocation(v string)`

SetDeploymentLocation sets DeploymentLocation field to given value.

### HasDeploymentLocation

`func (o *Kubernetes) HasDeploymentLocation() bool`

HasDeploymentLocation returns a boolean if a field has been set.

### GetK8sConnectionType

`func (o *Kubernetes) GetK8sConnectionType() string`

GetK8sConnectionType returns the K8sConnectionType field if non-nil, zero value otherwise.

### GetK8sConnectionTypeOk

`func (o *Kubernetes) GetK8sConnectionTypeOk() (*string, bool)`

GetK8sConnectionTypeOk returns a tuple with the K8sConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sConnectionType

`func (o *Kubernetes) SetK8sConnectionType(v string)`

SetK8sConnectionType sets K8sConnectionType field to given value.

### HasK8sConnectionType

`func (o *Kubernetes) HasK8sConnectionType() bool`

HasK8sConnectionType returns a boolean if a field has been set.

### GetStatus

`func (o *Kubernetes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Kubernetes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Kubernetes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Kubernetes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCost

`func (o *Kubernetes) GetCost() KubernetesCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Kubernetes) GetCostOk() (*KubernetesCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Kubernetes) SetCost(v KubernetesCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Kubernetes) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetControlPlaneStatus

`func (o *Kubernetes) GetControlPlaneStatus() string`

GetControlPlaneStatus returns the ControlPlaneStatus field if non-nil, zero value otherwise.

### GetControlPlaneStatusOk

`func (o *Kubernetes) GetControlPlaneStatusOk() (*string, bool)`

GetControlPlaneStatusOk returns a tuple with the ControlPlaneStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneStatus

`func (o *Kubernetes) SetControlPlaneStatus(v string)`

SetControlPlaneStatus sets ControlPlaneStatus field to given value.

### HasControlPlaneStatus

`func (o *Kubernetes) HasControlPlaneStatus() bool`

HasControlPlaneStatus returns a boolean if a field has been set.

### GetDomainName

`func (o *Kubernetes) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Kubernetes) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Kubernetes) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Kubernetes) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetNodeGroups

`func (o *Kubernetes) GetNodeGroups() []KubernetesNodeGroupsInner`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *Kubernetes) GetNodeGroupsOk() (*[]KubernetesNodeGroupsInner, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *Kubernetes) SetNodeGroups(v []KubernetesNodeGroupsInner)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *Kubernetes) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### GetAutoscalingConfigs

`func (o *Kubernetes) GetAutoscalingConfigs() []KubernetesAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *Kubernetes) GetAutoscalingConfigsOk() (*[]KubernetesAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *Kubernetes) SetAutoscalingConfigs(v []KubernetesAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.

### HasAutoscalingConfigs

`func (o *Kubernetes) HasAutoscalingConfigs() bool`

HasAutoscalingConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


