# KubernetesCreateRequestAutoscalingConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** | Name of the autoscaling group | [optional] 
**DataCenterId** | Pointer to **string** | ID of the data center | [optional] 
**MinimumNodes** | Pointer to **int32** | Minimum number of nodes | [optional] 
**MaximumNodes** | Pointer to **int32** | Maximum number of nodes | [optional] 
**TargetNodes** | Pointer to **int32** | Target number of worker nodes in the autoscaling group | [optional] 
**MinimumVCpus** | Pointer to **int32** | Minimum number of vCPUs | [optional] 
**MaximumVCpus** | Pointer to **int32** | Maximum number of vCPUs | [optional] 
**TargetVCpus** | Pointer to **int32** | Target number of vCPUs in the autoscaling group | [optional] 
**UseOnDemandInstancesInsteadOfSpots** | Pointer to **bool** | Use on-demand compute instances instead of spot compute instances | [optional] 
**SpotPercent** | Pointer to **int32** | Percentage of spot instances in the autoscaling group | [optional] 
**SpotMarkup** | Pointer to **float32** | Extra markup to the spot instance price | [optional] 
**ConfigurationPriorities** | Pointer to [**[]KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner**](KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner.md) | Compute instance configurations with priorities | [optional] 

## Methods

### NewKubernetesCreateRequestAutoscalingConfigsInner

`func NewKubernetesCreateRequestAutoscalingConfigsInner() *KubernetesCreateRequestAutoscalingConfigsInner`

NewKubernetesCreateRequestAutoscalingConfigsInner instantiates a new KubernetesCreateRequestAutoscalingConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCreateRequestAutoscalingConfigsInnerWithDefaults

`func NewKubernetesCreateRequestAutoscalingConfigsInnerWithDefaults() *KubernetesCreateRequestAutoscalingConfigsInner`

NewKubernetesCreateRequestAutoscalingConfigsInnerWithDefaults instantiates a new KubernetesCreateRequestAutoscalingConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetMinimumNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetMaximumNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.

### HasMaximumNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasMaximumNodes() bool`

HasMaximumNodes returns a boolean if a field has been set.

### GetTargetNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetTargetNodes() int32`

GetTargetNodes returns the TargetNodes field if non-nil, zero value otherwise.

### GetTargetNodesOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetTargetNodesOk() (*int32, bool)`

GetTargetNodesOk returns a tuple with the TargetNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetTargetNodes(v int32)`

SetTargetNodes sets TargetNodes field to given value.

### HasTargetNodes

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasTargetNodes() bool`

HasTargetNodes returns a boolean if a field has been set.

### GetMinimumVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMinimumVCpus() int32`

GetMinimumVCpus returns the MinimumVCpus field if non-nil, zero value otherwise.

### GetMinimumVCpusOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMinimumVCpusOk() (*int32, bool)`

GetMinimumVCpusOk returns a tuple with the MinimumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetMinimumVCpus(v int32)`

SetMinimumVCpus sets MinimumVCpus field to given value.

### HasMinimumVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasMinimumVCpus() bool`

HasMinimumVCpus returns a boolean if a field has been set.

### GetMaximumVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMaximumVCpus() int32`

GetMaximumVCpus returns the MaximumVCpus field if non-nil, zero value otherwise.

### GetMaximumVCpusOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetMaximumVCpusOk() (*int32, bool)`

GetMaximumVCpusOk returns a tuple with the MaximumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetMaximumVCpus(v int32)`

SetMaximumVCpus sets MaximumVCpus field to given value.

### HasMaximumVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasMaximumVCpus() bool`

HasMaximumVCpus returns a boolean if a field has been set.

### GetTargetVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetTargetVCpus() int32`

GetTargetVCpus returns the TargetVCpus field if non-nil, zero value otherwise.

### GetTargetVCpusOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetTargetVCpusOk() (*int32, bool)`

GetTargetVCpusOk returns a tuple with the TargetVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetTargetVCpus(v int32)`

SetTargetVCpus sets TargetVCpus field to given value.

### HasTargetVCpus

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasTargetVCpus() bool`

HasTargetVCpus returns a boolean if a field has been set.

### GetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpots() bool`

GetUseOnDemandInstancesInsteadOfSpots returns the UseOnDemandInstancesInsteadOfSpots field if non-nil, zero value otherwise.

### GetUseOnDemandInstancesInsteadOfSpotsOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpotsOk() (*bool, bool)`

GetUseOnDemandInstancesInsteadOfSpotsOk returns a tuple with the UseOnDemandInstancesInsteadOfSpots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetUseOnDemandInstancesInsteadOfSpots(v bool)`

SetUseOnDemandInstancesInsteadOfSpots sets UseOnDemandInstancesInsteadOfSpots field to given value.

### HasUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasUseOnDemandInstancesInsteadOfSpots() bool`

HasUseOnDemandInstancesInsteadOfSpots returns a boolean if a field has been set.

### GetSpotPercent

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetSpotPercent() int32`

GetSpotPercent returns the SpotPercent field if non-nil, zero value otherwise.

### GetSpotPercentOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetSpotPercentOk() (*int32, bool)`

GetSpotPercentOk returns a tuple with the SpotPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPercent

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetSpotPercent(v int32)`

SetSpotPercent sets SpotPercent field to given value.

### HasSpotPercent

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasSpotPercent() bool`

HasSpotPercent returns a boolean if a field has been set.

### GetSpotMarkup

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetSpotMarkup() float32`

GetSpotMarkup returns the SpotMarkup field if non-nil, zero value otherwise.

### GetSpotMarkupOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetSpotMarkupOk() (*float32, bool)`

GetSpotMarkupOk returns a tuple with the SpotMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotMarkup

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetSpotMarkup(v float32)`

SetSpotMarkup sets SpotMarkup field to given value.

### HasSpotMarkup

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasSpotMarkup() bool`

HasSpotMarkup returns a boolean if a field has been set.

### GetConfigurationPriorities

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetConfigurationPriorities() []KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner`

GetConfigurationPriorities returns the ConfigurationPriorities field if non-nil, zero value otherwise.

### GetConfigurationPrioritiesOk

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) GetConfigurationPrioritiesOk() (*[]KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner, bool)`

GetConfigurationPrioritiesOk returns a tuple with the ConfigurationPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPriorities

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) SetConfigurationPriorities(v []KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner)`

SetConfigurationPriorities sets ConfigurationPriorities field to given value.

### HasConfigurationPriorities

`func (o *KubernetesCreateRequestAutoscalingConfigsInner) HasConfigurationPriorities() bool`

HasConfigurationPriorities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


