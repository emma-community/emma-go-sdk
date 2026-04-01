# KubernetesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the Kubernetes cluster | [optional] 
**CreatedAt** | Pointer to **string** | Date and time of the Kubernetes cluster creation | [optional] 
**CreatedById** | Pointer to **int32** | ID of the user who created the Kubernetes cluster | [optional] 
**CreatedByName** | Pointer to **string** | Name of the user who created the Kubernetes cluster | [optional] 
**ModifiedAt** | Pointer to **string** | Date and time when the Kubernetes cluster was last edited | [optional] 
**ModifiedByName** | Pointer to **string** | Name of the user who last edited the Kubernetes cluster | [optional] 
**ModifiedById** | Pointer to **int32** | ID of the user who last edited the Kubernetes cluster | [optional] 
**Name** | Pointer to **string** | Name of the Kubernetes cluster | [optional] 
**Version** | Pointer to **string** | Version of the Kubernetes cluster | [optional] 
**DeploymentLocation** | Pointer to **string** | Deployment region of the Kubernetes cluster | [optional] 
**K8sConnectionType** | Pointer to **string** | Kubernetes connection type | [optional] 
**DomainName** | Pointer to **string** | Domain attached to the Kubernetes cluster | [optional] 
**Status** | Pointer to **string** | Status of the Kubernetes cluster | [optional] 
**ControlPlaneStatus** | Pointer to **string** | Control plane status | [optional] 
**Cost** | Pointer to [**KubernetesListResponseInnerCost**](KubernetesListResponseInnerCost.md) |  | [optional] 
**NodeGroups** | Pointer to [**[]KubernetesGetResponseNodeGroupsInner**](KubernetesGetResponseNodeGroupsInner.md) | List of the worker node groups | [optional] 
**AutoscalingConfigs** | Pointer to [**[]KubernetesListResponseInnerAutoscalingConfigsInner**](KubernetesListResponseInnerAutoscalingConfigsInner.md) | Configurations of the autoscaling group | [optional] 

## Methods

### NewKubernetesGetResponse

`func NewKubernetesGetResponse() *KubernetesGetResponse`

NewKubernetesGetResponse instantiates a new KubernetesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesGetResponseWithDefaults

`func NewKubernetesGetResponseWithDefaults() *KubernetesGetResponse`

NewKubernetesGetResponseWithDefaults instantiates a new KubernetesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesGetResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesGetResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesGetResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubernetesGetResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesGetResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesGetResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesGetResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *KubernetesGetResponse) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *KubernetesGetResponse) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *KubernetesGetResponse) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *KubernetesGetResponse) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *KubernetesGetResponse) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *KubernetesGetResponse) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *KubernetesGetResponse) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *KubernetesGetResponse) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetModifiedAt

`func (o *KubernetesGetResponse) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *KubernetesGetResponse) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *KubernetesGetResponse) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *KubernetesGetResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *KubernetesGetResponse) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *KubernetesGetResponse) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *KubernetesGetResponse) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *KubernetesGetResponse) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *KubernetesGetResponse) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *KubernetesGetResponse) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *KubernetesGetResponse) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *KubernetesGetResponse) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *KubernetesGetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesGetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesGetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesGetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *KubernetesGetResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesGetResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesGetResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesGetResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentLocation

`func (o *KubernetesGetResponse) GetDeploymentLocation() string`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *KubernetesGetResponse) GetDeploymentLocationOk() (*string, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *KubernetesGetResponse) SetDeploymentLocation(v string)`

SetDeploymentLocation sets DeploymentLocation field to given value.

### HasDeploymentLocation

`func (o *KubernetesGetResponse) HasDeploymentLocation() bool`

HasDeploymentLocation returns a boolean if a field has been set.

### GetK8sConnectionType

`func (o *KubernetesGetResponse) GetK8sConnectionType() string`

GetK8sConnectionType returns the K8sConnectionType field if non-nil, zero value otherwise.

### GetK8sConnectionTypeOk

`func (o *KubernetesGetResponse) GetK8sConnectionTypeOk() (*string, bool)`

GetK8sConnectionTypeOk returns a tuple with the K8sConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sConnectionType

`func (o *KubernetesGetResponse) SetK8sConnectionType(v string)`

SetK8sConnectionType sets K8sConnectionType field to given value.

### HasK8sConnectionType

`func (o *KubernetesGetResponse) HasK8sConnectionType() bool`

HasK8sConnectionType returns a boolean if a field has been set.

### GetDomainName

`func (o *KubernetesGetResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *KubernetesGetResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *KubernetesGetResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *KubernetesGetResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesGetResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesGetResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesGetResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesGetResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetControlPlaneStatus

`func (o *KubernetesGetResponse) GetControlPlaneStatus() string`

GetControlPlaneStatus returns the ControlPlaneStatus field if non-nil, zero value otherwise.

### GetControlPlaneStatusOk

`func (o *KubernetesGetResponse) GetControlPlaneStatusOk() (*string, bool)`

GetControlPlaneStatusOk returns a tuple with the ControlPlaneStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneStatus

`func (o *KubernetesGetResponse) SetControlPlaneStatus(v string)`

SetControlPlaneStatus sets ControlPlaneStatus field to given value.

### HasControlPlaneStatus

`func (o *KubernetesGetResponse) HasControlPlaneStatus() bool`

HasControlPlaneStatus returns a boolean if a field has been set.

### GetCost

`func (o *KubernetesGetResponse) GetCost() KubernetesListResponseInnerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *KubernetesGetResponse) GetCostOk() (*KubernetesListResponseInnerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *KubernetesGetResponse) SetCost(v KubernetesListResponseInnerCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *KubernetesGetResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetNodeGroups

`func (o *KubernetesGetResponse) GetNodeGroups() []KubernetesGetResponseNodeGroupsInner`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *KubernetesGetResponse) GetNodeGroupsOk() (*[]KubernetesGetResponseNodeGroupsInner, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *KubernetesGetResponse) SetNodeGroups(v []KubernetesGetResponseNodeGroupsInner)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *KubernetesGetResponse) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### GetAutoscalingConfigs

`func (o *KubernetesGetResponse) GetAutoscalingConfigs() []KubernetesListResponseInnerAutoscalingConfigsInner`

GetAutoscalingConfigs returns the AutoscalingConfigs field if non-nil, zero value otherwise.

### GetAutoscalingConfigsOk

`func (o *KubernetesGetResponse) GetAutoscalingConfigsOk() (*[]KubernetesListResponseInnerAutoscalingConfigsInner, bool)`

GetAutoscalingConfigsOk returns a tuple with the AutoscalingConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingConfigs

`func (o *KubernetesGetResponse) SetAutoscalingConfigs(v []KubernetesListResponseInnerAutoscalingConfigsInner)`

SetAutoscalingConfigs sets AutoscalingConfigs field to given value.

### HasAutoscalingConfigs

`func (o *KubernetesGetResponse) HasAutoscalingConfigs() bool`

HasAutoscalingConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


