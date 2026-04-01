# KubernetesDeleteAutoscalingConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** | Group name | [optional] 
**DataCenterId** | Pointer to **string** | ID of the data center | [optional] 
**MinimumNodes** | Pointer to **int32** | Minimum number of nodes | [optional] 
**MaximumNodes** | Pointer to **int32** | Maximum number of nodes | [optional] 
**TargetNodes** | Pointer to **int32** | Target number of nodes | [optional] 
**MinimumVCpus** | Pointer to **int32** | Minimum number of vCPUs | [optional] 
**MaximumVCpus** | Pointer to **int32** | Maximum number of vCPUs | [optional] 
**TargetVCpus** | Pointer to **int32** | Target number of vCPUs | [optional] 
**NodeGroupPriceLimit** | Pointer to **float64** | Price limit of the autoscaling group | [optional] 
**UseOnDemandInstancesInsteadOfSpots** | Pointer to **bool** | Use on-demand compute instances instead of spot compute instances | [optional] 
**SpotPercent** | Pointer to **int32** | Percentage of spot instances in the autoscaling group | [optional] 
**SpotMarkup** | Pointer to **float64** | Extra markup to the spot instance price | [optional] 
**ConfigurationPriorities** | Pointer to [**[]KubernetesDeleteAutoscalingConfigsInnerConfigurationPrioritiesInner**](KubernetesDeleteAutoscalingConfigsInnerConfigurationPrioritiesInner.md) | Compute instance configurations with priorities | [optional] 
**LimitsNotice** | Pointer to **string** | Limits notice message | [optional] 

## Methods

### NewKubernetesDeleteAutoscalingConfigsInner

`func NewKubernetesDeleteAutoscalingConfigsInner() *KubernetesDeleteAutoscalingConfigsInner`

NewKubernetesDeleteAutoscalingConfigsInner instantiates a new KubernetesDeleteAutoscalingConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDeleteAutoscalingConfigsInnerWithDefaults

`func NewKubernetesDeleteAutoscalingConfigsInnerWithDefaults() *KubernetesDeleteAutoscalingConfigsInner`

NewKubernetesDeleteAutoscalingConfigsInnerWithDefaults instantiates a new KubernetesDeleteAutoscalingConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetMinimumNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetMaximumNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.

### HasMaximumNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasMaximumNodes() bool`

HasMaximumNodes returns a boolean if a field has been set.

### GetTargetNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetTargetNodes() int32`

GetTargetNodes returns the TargetNodes field if non-nil, zero value otherwise.

### GetTargetNodesOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetTargetNodesOk() (*int32, bool)`

GetTargetNodesOk returns a tuple with the TargetNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetTargetNodes(v int32)`

SetTargetNodes sets TargetNodes field to given value.

### HasTargetNodes

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasTargetNodes() bool`

HasTargetNodes returns a boolean if a field has been set.

### GetMinimumVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMinimumVCpus() int32`

GetMinimumVCpus returns the MinimumVCpus field if non-nil, zero value otherwise.

### GetMinimumVCpusOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMinimumVCpusOk() (*int32, bool)`

GetMinimumVCpusOk returns a tuple with the MinimumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetMinimumVCpus(v int32)`

SetMinimumVCpus sets MinimumVCpus field to given value.

### HasMinimumVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasMinimumVCpus() bool`

HasMinimumVCpus returns a boolean if a field has been set.

### GetMaximumVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMaximumVCpus() int32`

GetMaximumVCpus returns the MaximumVCpus field if non-nil, zero value otherwise.

### GetMaximumVCpusOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetMaximumVCpusOk() (*int32, bool)`

GetMaximumVCpusOk returns a tuple with the MaximumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetMaximumVCpus(v int32)`

SetMaximumVCpus sets MaximumVCpus field to given value.

### HasMaximumVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasMaximumVCpus() bool`

HasMaximumVCpus returns a boolean if a field has been set.

### GetTargetVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetTargetVCpus() int32`

GetTargetVCpus returns the TargetVCpus field if non-nil, zero value otherwise.

### GetTargetVCpusOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetTargetVCpusOk() (*int32, bool)`

GetTargetVCpusOk returns a tuple with the TargetVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetTargetVCpus(v int32)`

SetTargetVCpus sets TargetVCpus field to given value.

### HasTargetVCpus

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasTargetVCpus() bool`

HasTargetVCpus returns a boolean if a field has been set.

### GetNodeGroupPriceLimit

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetNodeGroupPriceLimit() float64`

GetNodeGroupPriceLimit returns the NodeGroupPriceLimit field if non-nil, zero value otherwise.

### GetNodeGroupPriceLimitOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetNodeGroupPriceLimitOk() (*float64, bool)`

GetNodeGroupPriceLimitOk returns a tuple with the NodeGroupPriceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroupPriceLimit

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetNodeGroupPriceLimit(v float64)`

SetNodeGroupPriceLimit sets NodeGroupPriceLimit field to given value.

### HasNodeGroupPriceLimit

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasNodeGroupPriceLimit() bool`

HasNodeGroupPriceLimit returns a boolean if a field has been set.

### GetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpots() bool`

GetUseOnDemandInstancesInsteadOfSpots returns the UseOnDemandInstancesInsteadOfSpots field if non-nil, zero value otherwise.

### GetUseOnDemandInstancesInsteadOfSpotsOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpotsOk() (*bool, bool)`

GetUseOnDemandInstancesInsteadOfSpotsOk returns a tuple with the UseOnDemandInstancesInsteadOfSpots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetUseOnDemandInstancesInsteadOfSpots(v bool)`

SetUseOnDemandInstancesInsteadOfSpots sets UseOnDemandInstancesInsteadOfSpots field to given value.

### HasUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasUseOnDemandInstancesInsteadOfSpots() bool`

HasUseOnDemandInstancesInsteadOfSpots returns a boolean if a field has been set.

### GetSpotPercent

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetSpotPercent() int32`

GetSpotPercent returns the SpotPercent field if non-nil, zero value otherwise.

### GetSpotPercentOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetSpotPercentOk() (*int32, bool)`

GetSpotPercentOk returns a tuple with the SpotPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPercent

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetSpotPercent(v int32)`

SetSpotPercent sets SpotPercent field to given value.

### HasSpotPercent

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasSpotPercent() bool`

HasSpotPercent returns a boolean if a field has been set.

### GetSpotMarkup

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetSpotMarkup() float64`

GetSpotMarkup returns the SpotMarkup field if non-nil, zero value otherwise.

### GetSpotMarkupOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetSpotMarkupOk() (*float64, bool)`

GetSpotMarkupOk returns a tuple with the SpotMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotMarkup

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetSpotMarkup(v float64)`

SetSpotMarkup sets SpotMarkup field to given value.

### HasSpotMarkup

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasSpotMarkup() bool`

HasSpotMarkup returns a boolean if a field has been set.

### GetConfigurationPriorities

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetConfigurationPriorities() []KubernetesDeleteAutoscalingConfigsInnerConfigurationPrioritiesInner`

GetConfigurationPriorities returns the ConfigurationPriorities field if non-nil, zero value otherwise.

### GetConfigurationPrioritiesOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetConfigurationPrioritiesOk() (*[]KubernetesDeleteAutoscalingConfigsInnerConfigurationPrioritiesInner, bool)`

GetConfigurationPrioritiesOk returns a tuple with the ConfigurationPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPriorities

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetConfigurationPriorities(v []KubernetesDeleteAutoscalingConfigsInnerConfigurationPrioritiesInner)`

SetConfigurationPriorities sets ConfigurationPriorities field to given value.

### HasConfigurationPriorities

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasConfigurationPriorities() bool`

HasConfigurationPriorities returns a boolean if a field has been set.

### GetLimitsNotice

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetLimitsNotice() string`

GetLimitsNotice returns the LimitsNotice field if non-nil, zero value otherwise.

### GetLimitsNoticeOk

`func (o *KubernetesDeleteAutoscalingConfigsInner) GetLimitsNoticeOk() (*string, bool)`

GetLimitsNoticeOk returns a tuple with the LimitsNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitsNotice

`func (o *KubernetesDeleteAutoscalingConfigsInner) SetLimitsNotice(v string)`

SetLimitsNotice sets LimitsNotice field to given value.

### HasLimitsNotice

`func (o *KubernetesDeleteAutoscalingConfigsInner) HasLimitsNotice() bool`

HasLimitsNotice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


