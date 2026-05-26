# KubernetesUpdateRequestAutoscalingConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | Group name | 
**DataCenterId** | **string** | ID of the data center | 
**MinimumNodes** | Pointer to **int32** | Minimum number of nodes. Required when using node-count based autoscaling. Must be null if &#x60;minimumVCpus&#x60; is set. | [optional] 
**MaximumNodes** | Pointer to **int32** | Maximum number of nodes. Required when using node-count based autoscaling. Must be null if &#x60;maximumVCpus&#x60; is set. | [optional] 
**TargetNodes** | Pointer to **int32** | Target number of worker nodes in the autoscaling group. Required when using node-count based autoscaling. Must be null if &#x60;targetVCpus&#x60; is set. | [optional] 
**MinimumVCpus** | Pointer to **int32** | Minimum number of vCPUs. Required when using vCPU-count based autoscaling. Must be null if &#x60;minimumNodes&#x60; is set. | [optional] 
**MaximumVCpus** | Pointer to **int32** | Maximum number of vCPUs. Required when using vCPU-count based autoscaling. Must be null if &#x60;maximumNodes&#x60; is set. | [optional] 
**TargetVCpus** | Pointer to **int32** | Target number of vCPUs in the autoscaling group. Required when using vCPU-count based autoscaling. Must be null if &#x60;targetNodes&#x60; is set. | [optional] 
**NodeGroupPriceLimit** | Pointer to **float32** | Price limit of the autoscaling group | [optional] 
**UseOnDemandInstancesInsteadOfSpots** | **bool** | Use on-demand compute instances instead of spot compute instances | 
**SpotPercent** | Pointer to **int32** | Percentage of spot instances in the autoscaling group | [optional] 
**SpotMarkup** | Pointer to **float32** | Extra markup to the spot instance price | [optional] 
**UserLabels** | Pointer to [**[]KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner**](KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner.md) | ASG label applied to nodes created by this autoscaling group | [optional] 
**ConfigurationPriorities** | [**[]KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner**](KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner.md) | Compute instance configurations with priorities | 

## Methods

### NewKubernetesUpdateRequestAutoscalingConfigsInner

`func NewKubernetesUpdateRequestAutoscalingConfigsInner(groupName string, dataCenterId string, useOnDemandInstancesInsteadOfSpots bool, configurationPriorities []KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner, ) *KubernetesUpdateRequestAutoscalingConfigsInner`

NewKubernetesUpdateRequestAutoscalingConfigsInner instantiates a new KubernetesUpdateRequestAutoscalingConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateRequestAutoscalingConfigsInnerWithDefaults

`func NewKubernetesUpdateRequestAutoscalingConfigsInnerWithDefaults() *KubernetesUpdateRequestAutoscalingConfigsInner`

NewKubernetesUpdateRequestAutoscalingConfigsInnerWithDefaults instantiates a new KubernetesUpdateRequestAutoscalingConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetDataCenterId

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetMinimumNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetMaximumNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.

### HasMaximumNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasMaximumNodes() bool`

HasMaximumNodes returns a boolean if a field has been set.

### GetTargetNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetTargetNodes() int32`

GetTargetNodes returns the TargetNodes field if non-nil, zero value otherwise.

### GetTargetNodesOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetTargetNodesOk() (*int32, bool)`

GetTargetNodesOk returns a tuple with the TargetNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetTargetNodes(v int32)`

SetTargetNodes sets TargetNodes field to given value.

### HasTargetNodes

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasTargetNodes() bool`

HasTargetNodes returns a boolean if a field has been set.

### GetMinimumVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMinimumVCpus() int32`

GetMinimumVCpus returns the MinimumVCpus field if non-nil, zero value otherwise.

### GetMinimumVCpusOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMinimumVCpusOk() (*int32, bool)`

GetMinimumVCpusOk returns a tuple with the MinimumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetMinimumVCpus(v int32)`

SetMinimumVCpus sets MinimumVCpus field to given value.

### HasMinimumVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasMinimumVCpus() bool`

HasMinimumVCpus returns a boolean if a field has been set.

### GetMaximumVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMaximumVCpus() int32`

GetMaximumVCpus returns the MaximumVCpus field if non-nil, zero value otherwise.

### GetMaximumVCpusOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetMaximumVCpusOk() (*int32, bool)`

GetMaximumVCpusOk returns a tuple with the MaximumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetMaximumVCpus(v int32)`

SetMaximumVCpus sets MaximumVCpus field to given value.

### HasMaximumVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasMaximumVCpus() bool`

HasMaximumVCpus returns a boolean if a field has been set.

### GetTargetVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetTargetVCpus() int32`

GetTargetVCpus returns the TargetVCpus field if non-nil, zero value otherwise.

### GetTargetVCpusOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetTargetVCpusOk() (*int32, bool)`

GetTargetVCpusOk returns a tuple with the TargetVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetTargetVCpus(v int32)`

SetTargetVCpus sets TargetVCpus field to given value.

### HasTargetVCpus

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasTargetVCpus() bool`

HasTargetVCpus returns a boolean if a field has been set.

### GetNodeGroupPriceLimit

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetNodeGroupPriceLimit() float32`

GetNodeGroupPriceLimit returns the NodeGroupPriceLimit field if non-nil, zero value otherwise.

### GetNodeGroupPriceLimitOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetNodeGroupPriceLimitOk() (*float32, bool)`

GetNodeGroupPriceLimitOk returns a tuple with the NodeGroupPriceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroupPriceLimit

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetNodeGroupPriceLimit(v float32)`

SetNodeGroupPriceLimit sets NodeGroupPriceLimit field to given value.

### HasNodeGroupPriceLimit

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasNodeGroupPriceLimit() bool`

HasNodeGroupPriceLimit returns a boolean if a field has been set.

### GetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpots() bool`

GetUseOnDemandInstancesInsteadOfSpots returns the UseOnDemandInstancesInsteadOfSpots field if non-nil, zero value otherwise.

### GetUseOnDemandInstancesInsteadOfSpotsOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpotsOk() (*bool, bool)`

GetUseOnDemandInstancesInsteadOfSpotsOk returns a tuple with the UseOnDemandInstancesInsteadOfSpots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetUseOnDemandInstancesInsteadOfSpots(v bool)`

SetUseOnDemandInstancesInsteadOfSpots sets UseOnDemandInstancesInsteadOfSpots field to given value.


### GetSpotPercent

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetSpotPercent() int32`

GetSpotPercent returns the SpotPercent field if non-nil, zero value otherwise.

### GetSpotPercentOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetSpotPercentOk() (*int32, bool)`

GetSpotPercentOk returns a tuple with the SpotPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPercent

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetSpotPercent(v int32)`

SetSpotPercent sets SpotPercent field to given value.

### HasSpotPercent

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasSpotPercent() bool`

HasSpotPercent returns a boolean if a field has been set.

### GetSpotMarkup

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetSpotMarkup() float32`

GetSpotMarkup returns the SpotMarkup field if non-nil, zero value otherwise.

### GetSpotMarkupOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetSpotMarkupOk() (*float32, bool)`

GetSpotMarkupOk returns a tuple with the SpotMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotMarkup

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetSpotMarkup(v float32)`

SetSpotMarkup sets SpotMarkup field to given value.

### HasSpotMarkup

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasSpotMarkup() bool`

HasSpotMarkup returns a boolean if a field has been set.

### GetUserLabels

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetUserLabels() []KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner`

GetUserLabels returns the UserLabels field if non-nil, zero value otherwise.

### GetUserLabelsOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetUserLabelsOk() (*[]KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner, bool)`

GetUserLabelsOk returns a tuple with the UserLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabels

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetUserLabels(v []KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner)`

SetUserLabels sets UserLabels field to given value.

### HasUserLabels

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) HasUserLabels() bool`

HasUserLabels returns a boolean if a field has been set.

### GetConfigurationPriorities

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetConfigurationPriorities() []KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner`

GetConfigurationPriorities returns the ConfigurationPriorities field if non-nil, zero value otherwise.

### GetConfigurationPrioritiesOk

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) GetConfigurationPrioritiesOk() (*[]KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner, bool)`

GetConfigurationPrioritiesOk returns a tuple with the ConfigurationPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPriorities

`func (o *KubernetesUpdateRequestAutoscalingConfigsInner) SetConfigurationPriorities(v []KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner)`

SetConfigurationPriorities sets ConfigurationPriorities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


