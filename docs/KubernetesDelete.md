# KubernetesDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Kubernetes cluster ID | [optional] 
**CreatedAt** | Pointer to **time.Time** | Cluster creation date and time | [optional] 
**CreatedById** | Pointer to **int32** | ID of the user who created the cluster | [optional] 
**CreatedByName** | Pointer to **string** | Name of the user who created the cluster | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Cluster last modification date and time | [optional] 
**ModifiedByName** | Pointer to **string** | Name of the user who last modified the cluster | [optional] 
**ModifiedById** | Pointer to **int32** | ID of the user who last modified the cluster | [optional] 
**Name** | Pointer to **string** | Cluster name | [optional] 
**Version** | Pointer to **string** | Kubernetes version | [optional] 
**DeploymentLocation** | Pointer to **string** | Deployment location | [optional] 
**K8sConnectionType** | Pointer to **string** | Kubernetes connection type | [optional] 
**DomainName** | Pointer to **string** | Cluster domain name | [optional] 
**Status** | Pointer to **string** | Cluster status | [optional] 
**ControlPlaneStatus** | Pointer to **string** | Control plane status | [optional] 
**Cost** | Pointer to [**KubernetesDeleteCost**](KubernetesDeleteCost.md) |  | [optional] 
**NodeGroups** | Pointer to [**[]KubernetesDeleteNodeGroupsInner**](KubernetesDeleteNodeGroupsInner.md) | List of cluster node groups | [optional] 
**AutoscalingConfigs** | Pointer to [**[]KubernetesDeleteAutoscalingConfigsInner**](KubernetesDeleteAutoscalingConfigsInner.md) | Configurations of autoscaling groups | [optional] 

## Methods

### NewKubernetesDelete

`func NewKubernetesDelete() *KubernetesDelete`

NewKubernetesDelete instantiates a new KubernetesDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDeleteWithDefaults

`func NewKubernetesDeleteWithDefaults() *KubernetesDelete`

NewKubernetesDeleteWithDefaults instantiates a new KubernetesDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesDelete) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesDelete) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesDelete) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesDelete) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubernetesDelete) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesDelete) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesDelete) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesDelete) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *KubernetesDelete) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *KubernetesDelete) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *KubernetesDelete) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *KubernetesDelete) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *KubernetesDelete) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *KubernetesDelete) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *KubernetesDelete) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *KubernetesDelete) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetModifiedAt

`func (o *KubernetesDelete) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *KubernetesDelete) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *KubernetesDelete) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *KubernetesDelete) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *KubernetesDelete) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *KubernetesDelete) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *KubernetesDelete) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *KubernetesDelete) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *KubernetesDelete) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *KubernetesDelete) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *KubernetesDelete) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *KubernetesDelete) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *KubernetesDelete) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesDelete) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesDelete) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesDelete) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *KubernetesDelete) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesDelete) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesDelete) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesDelete) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentLocation

`func (o *KubernetesDelete) GetDeploymentLocation() string`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *KubernetesDelete) GetDeploymentLocationOk() (*string, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *KubernetesDelete) SetDeploymentLocation(v string)`

SetDeploymentLocation sets DeploymentLocation field to given value.

### HasDeploymentLocation

`func (o *KubernetesDelete) HasDeploymentLocation() bool`

HasDeploymentLocation returns a boolean if a field has been set.

### GetK8sConnectionType

`func (o *KubernetesDelete) GetK8sConnectionType() string`

GetK8sConnectionType returns the K8sConnectionType field if non-nil, zero value otherwise.

### GetK8sConnectionTypeOk

`func (o *KubernetesDelete) GetK8sConnectionTypeOk() (*string, bool)`

GetK8sConnectionTypeOk returns a tuple with the K8sConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sConnectionType

`func (o *KubernetesDelete) SetK8sConnectionType(v string)`

SetK8sConnectionType sets K8sConnectionType field to given value.

### HasK8sConnectionType

`func (o *KubernetesDelete) HasK8sConnectionType() bool`

HasK8sConnectionType returns a boolean if a field has been set.

### GetDomainName

`func (o *KubernetesDelete) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *KubernetesDelete) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *KubernetesDelete) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *KubernetesDelete) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesDelete) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesDelete) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesDelete) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesDelete) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetControlPlaneStatus

`func (o *KubernetesDelete) GetControlPlaneStatus() string`

GetControlPlaneStatus returns the ControlPlaneStatus field if non-nil, zero value otherwise.

### GetControlPlaneStatusOk

`func (o *KubernetesDelete) GetControlPlaneStatusOk() (*string, bool)`

GetControlPlaneStatusOk returns a tuple with the ControlPlaneStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneStatus

`func (o *KubernetesDelete) SetControlPlaneStatus(v string)`

SetControlPlaneStatus sets ControlPlaneStatus field to given value.

### HasControlPlaneStatus

`func (o *KubernetesDelete) HasControlPlaneStatus() bool`

HasControlPlaneStatus returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesDelete) GetCost() KubernetesDeleteCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesDelete) GetCostOk() (*KubernetesDeleteCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesDelete) SetCost(v KubernetesDeleteCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *KubernetesDelete) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetNodeGroups

`func (o *KubernetesDelete) GetNodeGroups() []KubernetesDeleteNodeGroupsInner`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *KubernetesDelete) GetNodeGroupsOk() (*[]KubernetesDeleteNodeGroupsInner, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *KubernetesDelete) SetNodeGroups(v []KubernetesDeleteNodeGroupsInner)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *KubernetesDelete) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### GetAutoscalingConfigs

`func (o *KubernetesDelete) GetAutoscalingConfigs() []KubernetesDeleteAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *KubernetesDelete) GetAutoscalingConfigsOk() (*[]KubernetesDeleteAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *KubernetesDelete) SetAutoscalingConfigs(v []KubernetesDeleteAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.

### HasAutoscalingConfigs

`func (o *KubernetesDelete) HasAutoscalingConfigs() bool`

HasAutoscalingConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


