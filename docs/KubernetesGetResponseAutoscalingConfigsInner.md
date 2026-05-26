# KubernetesGetResponseAutoscalingConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** | Name of the autoscaling group | [optional] 
**DataCenterId** | Pointer to **string** | ID of the data center | [optional] 
**MinimumNodes** | Pointer to **int32** | Minimum number of nodes. Required when using node-count based autoscaling. Must be null if &#x60;minimumVCpus&#x60; is set. | [optional] 
**MaximumNodes** | Pointer to **int32** | Maximum number of nodes.Required when using node-count based autoscaling. Must be null if &#x60;maximumVCpus&#x60; is set. | [optional] 
**TargetNodes** | Pointer to **int32** | Target number of nodes. Required when using node-count based autoscaling. Must be null if &#x60;targetVCpus&#x60; is set. | [optional] 
**MinimumVCpus** | Pointer to **int32** | Minimum number of vCPUs. Required when using vCPU-count based autoscaling. Must be null if &#x60;minimumNodes&#x60; is set. | [optional] 
**MaximumVCpus** | Pointer to **int32** | Maximum number of vCPUs. Required when using vCPU-count based autoscaling. Must be null if &#x60;maximumNodes&#x60; is set. | [optional] 
**TargetVCpus** | Pointer to **int32** | Target number of vCPUs. Required when using vCPU-count based autoscaling. Must be null if &#x60;targetNodes&#x60; is set. | [optional] 
**NodeGroupPriceLimit** | Pointer to **float32** | Price limit of the autoscaling group | [optional] 
**UseOnDemandInstancesInsteadOfSpots** | Pointer to **bool** | Use on-demand compute instances instead of spot compute instances | [optional] 
**SpotPercent** | Pointer to **int32** | Percentage of spot instances in the autoscaling group | [optional] 
**SpotMarkup** | Pointer to **float32** | Extra markup to the spot instance price | [optional] 
**UserLabels** | Pointer to [**[]KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner**](KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner.md) | ASG labels applied to nodes created by this autoscaling group | [optional] 
**ConfigurationPriorities** | Pointer to [**[]KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner**](KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner.md) | Compute instance configurations with priorities | [optional] 
**LimitsNotice** | Pointer to **string** | Limits notice message | [optional] 

## Methods

### NewKubernetesGetResponseAutoscalingConfigsInner

`func NewKubernetesGetResponseAutoscalingConfigsInner() *KubernetesGetResponseAutoscalingConfigsInner`

NewKubernetesGetResponseAutoscalingConfigsInner instantiates a new KubernetesGetResponseAutoscalingConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesGetResponseAutoscalingConfigsInnerWithDefaults

`func NewKubernetesGetResponseAutoscalingConfigsInnerWithDefaults() *KubernetesGetResponseAutoscalingConfigsInner`

NewKubernetesGetResponseAutoscalingConfigsInnerWithDefaults instantiates a new KubernetesGetResponseAutoscalingConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetMinimumNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetMaximumNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.

### HasMaximumNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasMaximumNodes() bool`

HasMaximumNodes returns a boolean if a field has been set.

### GetTargetNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetTargetNodes() int32`

GetTargetNodes returns the TargetNodes field if non-nil, zero value otherwise.

### GetTargetNodesOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetTargetNodesOk() (*int32, bool)`

GetTargetNodesOk returns a tuple with the TargetNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetTargetNodes(v int32)`

SetTargetNodes sets TargetNodes field to given value.

### HasTargetNodes

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasTargetNodes() bool`

HasTargetNodes returns a boolean if a field has been set.

### GetMinimumVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMinimumVCpus() int32`

GetMinimumVCpus returns the MinimumVCpus field if non-nil, zero value otherwise.

### GetMinimumVCpusOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMinimumVCpusOk() (*int32, bool)`

GetMinimumVCpusOk returns a tuple with the MinimumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetMinimumVCpus(v int32)`

SetMinimumVCpus sets MinimumVCpus field to given value.

### HasMinimumVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasMinimumVCpus() bool`

HasMinimumVCpus returns a boolean if a field has been set.

### GetMaximumVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMaximumVCpus() int32`

GetMaximumVCpus returns the MaximumVCpus field if non-nil, zero value otherwise.

### GetMaximumVCpusOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetMaximumVCpusOk() (*int32, bool)`

GetMaximumVCpusOk returns a tuple with the MaximumVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetMaximumVCpus(v int32)`

SetMaximumVCpus sets MaximumVCpus field to given value.

### HasMaximumVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasMaximumVCpus() bool`

HasMaximumVCpus returns a boolean if a field has been set.

### GetTargetVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetTargetVCpus() int32`

GetTargetVCpus returns the TargetVCpus field if non-nil, zero value otherwise.

### GetTargetVCpusOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetTargetVCpusOk() (*int32, bool)`

GetTargetVCpusOk returns a tuple with the TargetVCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetTargetVCpus(v int32)`

SetTargetVCpus sets TargetVCpus field to given value.

### HasTargetVCpus

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasTargetVCpus() bool`

HasTargetVCpus returns a boolean if a field has been set.

### GetNodeGroupPriceLimit

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetNodeGroupPriceLimit() float32`

GetNodeGroupPriceLimit returns the NodeGroupPriceLimit field if non-nil, zero value otherwise.

### GetNodeGroupPriceLimitOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetNodeGroupPriceLimitOk() (*float32, bool)`

GetNodeGroupPriceLimitOk returns a tuple with the NodeGroupPriceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroupPriceLimit

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetNodeGroupPriceLimit(v float32)`

SetNodeGroupPriceLimit sets NodeGroupPriceLimit field to given value.

### HasNodeGroupPriceLimit

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasNodeGroupPriceLimit() bool`

HasNodeGroupPriceLimit returns a boolean if a field has been set.

### GetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpots() bool`

GetUseOnDemandInstancesInsteadOfSpots returns the UseOnDemandInstancesInsteadOfSpots field if non-nil, zero value otherwise.

### GetUseOnDemandInstancesInsteadOfSpotsOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetUseOnDemandInstancesInsteadOfSpotsOk() (*bool, bool)`

GetUseOnDemandInstancesInsteadOfSpotsOk returns a tuple with the UseOnDemandInstancesInsteadOfSpots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetUseOnDemandInstancesInsteadOfSpots(v bool)`

SetUseOnDemandInstancesInsteadOfSpots sets UseOnDemandInstancesInsteadOfSpots field to given value.

### HasUseOnDemandInstancesInsteadOfSpots

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasUseOnDemandInstancesInsteadOfSpots() bool`

HasUseOnDemandInstancesInsteadOfSpots returns a boolean if a field has been set.

### GetSpotPercent

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetSpotPercent() int32`

GetSpotPercent returns the SpotPercent field if non-nil, zero value otherwise.

### GetSpotPercentOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetSpotPercentOk() (*int32, bool)`

GetSpotPercentOk returns a tuple with the SpotPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPercent

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetSpotPercent(v int32)`

SetSpotPercent sets SpotPercent field to given value.

### HasSpotPercent

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasSpotPercent() bool`

HasSpotPercent returns a boolean if a field has been set.

### GetSpotMarkup

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetSpotMarkup() float32`

GetSpotMarkup returns the SpotMarkup field if non-nil, zero value otherwise.

### GetSpotMarkupOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetSpotMarkupOk() (*float32, bool)`

GetSpotMarkupOk returns a tuple with the SpotMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotMarkup

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetSpotMarkup(v float32)`

SetSpotMarkup sets SpotMarkup field to given value.

### HasSpotMarkup

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasSpotMarkup() bool`

HasSpotMarkup returns a boolean if a field has been set.

### GetUserLabels

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetUserLabels() []KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner`

GetUserLabels returns the UserLabels field if non-nil, zero value otherwise.

### GetUserLabelsOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetUserLabelsOk() (*[]KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner, bool)`

GetUserLabelsOk returns a tuple with the UserLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabels

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetUserLabels(v []KubernetesListResponseInnerAutoscalingConfigsInnerUserLabelsInner)`

SetUserLabels sets UserLabels field to given value.

### HasUserLabels

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasUserLabels() bool`

HasUserLabels returns a boolean if a field has been set.

### GetConfigurationPriorities

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetConfigurationPriorities() []KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner`

GetConfigurationPriorities returns the ConfigurationPriorities field if non-nil, zero value otherwise.

### GetConfigurationPrioritiesOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetConfigurationPrioritiesOk() (*[]KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner, bool)`

GetConfigurationPrioritiesOk returns a tuple with the ConfigurationPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPriorities

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetConfigurationPriorities(v []KubernetesListResponseInnerAutoscalingConfigsInnerConfigurationPrioritiesInner)`

SetConfigurationPriorities sets ConfigurationPriorities field to given value.

### HasConfigurationPriorities

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasConfigurationPriorities() bool`

HasConfigurationPriorities returns a boolean if a field has been set.

### GetLimitsNotice

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetLimitsNotice() string`

GetLimitsNotice returns the LimitsNotice field if non-nil, zero value otherwise.

### GetLimitsNoticeOk

`func (o *KubernetesGetResponseAutoscalingConfigsInner) GetLimitsNoticeOk() (*string, bool)`

GetLimitsNoticeOk returns a tuple with the LimitsNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitsNotice

`func (o *KubernetesGetResponseAutoscalingConfigsInner) SetLimitsNotice(v string)`

SetLimitsNotice sets LimitsNotice field to given value.

### HasLimitsNotice

`func (o *KubernetesGetResponseAutoscalingConfigsInner) HasLimitsNotice() bool`

HasLimitsNotice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


