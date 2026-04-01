# KubernetesUpdateResponseAutoscalingConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | Group name | 
**DataCenterId** | **string** | ID of the data center | 
**MinimumNodes** | **int32** | Minimum number of nodes | 
**MaximumNodes** | **int32** | Maximum number of nodes | 
**TargetNodes** | **int32** | Target number of nodes | 
**MinimumVCpus** | Pointer to **int32** | Minimum number of vCPUs | [optional] 
**MaximumVCpus** | Pointer to **int32** | Maximum number of vCPUs | [optional] 
**TargetVCpus** | Pointer to **int32** | Target number of vCPUs | [optional] 
**NodeGroupPriceLimit** | **float64** | Price limit of the autoscaling group | 
**UseOnDemandInstancesInsteadOfSpots** | **bool** | Use on-demand compute instances instead of spot compute instances | 
**SpotPercent** | **int32** | Percentage of spot instances in the autoscaling group | 
**SpotMarkup** | **float64** | Extra markup to the spot instance price | 
**ConfigurationPriorities** | [**[]KubernetesUpdateResponseAutoscalingConfigsInnerConfigurationPrioritiesInner**](KubernetesUpdateResponseAutoscalingConfigsInnerConfigurationPrioritiesInner.md) | Compute instance configurations with priorities | 
**LimitsNotice** | Pointer to **string** | Limits notice message | [optional] 

## Methods

### NewKubernetesUpdateResponseAutoscalingConfigsInner

`func NewKubernetesUpdateResponseAutoscalingConfigsInner(groupName string, dataCenterId string, minimumNodes int32, maximumNodes int32, targetNodes int32, nodeGroupPriceLimit float64, useOnDemandInstancesInsteadOfSpots bool, spotPercent int32, spotMarkup float64, configurationPriorities []KubernetesUpdateResponseAutoscalingConfigsInnerConfigurationPrioritiesInner, ) *KubernetesUpdateResponseAutoscalingConfigsInner`

NewKubernetesUpdateResponseAutoscalingConfigsInner instantiates a new KubernetesUpdateResponseAutoscalingConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateResponseAutoscalingConfigsInnerWithDefaults

`func NewKubernetesUpdateResponseAutoscalingConfigsInnerWithDefaults() *KubernetesUpdateResponseAutoscalingConfigsInner`

NewKubernetesUpdateResponseAutoscalingConfigsInnerWithDefaults instantiates a new KubernetesUpdateResponseAutoscalingConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetDataCenterId

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetMinimumNodes

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.


### GetMaximumNodes

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.


### GetTargetNodes

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetTargetNodes() int32`

GetTargetNodes returns the TargetNodes field if non-nil, zero value otherwise.

### GetTargetNodesOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetTargetNodesOk() (*int32, bool)`

GetTargetNodesOk returns a tuple with the TargetNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodes

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetTargetNodes(v int32)`

SetTargetNodes sets TargetNodes field to given value.


### GetMinimumVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMinimumVCpus() int32`

GetMinimumVCpus returns the MinimumVCpus field if non-nil, zero value otherwise.

### GetMinimumVCpusOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMinimumVCpusOk() (*int32, bool)`

GetMinimumVCpusOk returns a tuple with the MinimumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetMinimumVCpus(v int32)`

SetMinimumVCpus sets MinimumVCpus field to given value.

### HasMinimumVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) HasMinimumVCpus() bool`

HasMinimumVCpus returns a boolean if a field has been set.

### GetMaximumVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMaximumVCpus() int32`

GetMaximumVCpus returns the MaximumVCpus field if non-nil, zero value otherwise.

### GetMaximumVCpusOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetMaximumVCpusOk() (*int32, bool)`

GetMaximumVCpusOk returns a tuple with the MaximumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetMaximumVCpus(v int32)`

SetMaximumVCpus sets MaximumVCpus field to given value.

### HasMaximumVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) HasMaximumVCpus() bool`

HasMaximumVCpus returns a boolean if a field has been set.

### GetTargetVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetTargetVCpus() int32`

GetTargetVCpus returns the TargetVCpus field if non-nil, zero value otherwise.

### GetTargetVCpusOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetTargetVCpusOk() (*int32, bool)`

GetTargetVCpusOk returns a tuple with the TargetVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetTargetVCpus(v int32)`

SetTargetVCpus sets TargetVCpus field to given value.

### HasTargetVCpus

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) HasTargetVCpus() bool`

HasTargetVCpus returns a boolean if a field has been set.

### GetNodeGroupPriceLimit

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetNodeGroupPriceLimit() float64`

GetNodeGroupPriceLimit returns the NodeGroupPriceLimit field if non-nil, zero value otherwise.

### GetNodeGroupPriceLimitOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetNodeGroupPriceLimitOk() (*float64, bool)`

GetNodeGroupPriceLimitOk returns a tuple with the NodeGroupPriceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroupPriceLimit

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetNodeGroupPriceLimit(v float64)`

SetNodeGroupPriceLimit sets NodeGroupPriceLimit field to given value.


### GetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpots() bool`

GetUseOnDemandInstancesInsteadOfSpots returns the UseOnDemandInstancesInsteadOfSpots field if non-nil, zero value otherwise.

### GetUseOnDemandInstancesInsteadOfSpotsOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpotsOk() (*bool, bool)`

GetUseOnDemandInstancesInsteadOfSpotsOk returns a tuple with the UseOnDemandInstancesInsteadOfSpots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetUseOnDemandInstancesInsteadOfSpots(v bool)`

SetUseOnDemandInstancesInsteadOfSpots sets UseOnDemandInstancesInsteadOfSpots field to given value.


### GetSpotPercent

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetSpotPercent() int32`

GetSpotPercent returns the SpotPercent field if non-nil, zero value otherwise.

### GetSpotPercentOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetSpotPercentOk() (*int32, bool)`

GetSpotPercentOk returns a tuple with the SpotPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPercent

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetSpotPercent(v int32)`

SetSpotPercent sets SpotPercent field to given value.


### GetSpotMarkup

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetSpotMarkup() float64`

GetSpotMarkup returns the SpotMarkup field if non-nil, zero value otherwise.

### GetSpotMarkupOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetSpotMarkupOk() (*float64, bool)`

GetSpotMarkupOk returns a tuple with the SpotMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotMarkup

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetSpotMarkup(v float64)`

SetSpotMarkup sets SpotMarkup field to given value.


### GetConfigurationPriorities

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetConfigurationPriorities() []KubernetesUpdateResponseAutoscalingConfigsInnerConfigurationPrioritiesInner`

GetConfigurationPriorities returns the ConfigurationPriorities field if non-nil, zero value otherwise.

### GetConfigurationPrioritiesOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetConfigurationPrioritiesOk() (*[]KubernetesUpdateResponseAutoscalingConfigsInnerConfigurationPrioritiesInner, bool)`

GetConfigurationPrioritiesOk returns a tuple with the ConfigurationPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPriorities

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetConfigurationPriorities(v []KubernetesUpdateResponseAutoscalingConfigsInnerConfigurationPrioritiesInner)`

SetConfigurationPriorities sets ConfigurationPriorities field to given value.


### GetLimitsNotice

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetLimitsNotice() string`

GetLimitsNotice returns the LimitsNotice field if non-nil, zero value otherwise.

### GetLimitsNoticeOk

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) GetLimitsNoticeOk() (*string, bool)`

GetLimitsNoticeOk returns a tuple with the LimitsNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitsNotice

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) SetLimitsNotice(v string)`

SetLimitsNotice sets LimitsNotice field to given value.

### HasLimitsNotice

`func (o *KubernetesUpdateResponseAutoscalingConfigsInner) HasLimitsNotice() bool`

HasLimitsNotice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


