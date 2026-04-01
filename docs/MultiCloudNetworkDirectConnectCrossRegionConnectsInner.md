# MultiCloudNetworkDirectConnectCrossRegionConnectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the cross-region connection. | [optional] 
**Networks** | Pointer to [**[]MultiCloudNetworkDirectConnectCrossRegionConnectsInnerNetworksInner**](MultiCloudNetworkDirectConnectCrossRegionConnectsInnerNetworksInner.md) | Network endpoints participating in the cross-region connection. | [optional] 
**Bandwidth** | Pointer to [**MultiCloudNetworkDirectConnectCrossRegionConnectsInnerBandwidth**](MultiCloudNetworkDirectConnectCrossRegionConnectsInnerBandwidth.md) |  | [optional] 
**Cost** | Pointer to [**MultiCloudNetworkDirectConnectCrossRegionConnectsInnerCost**](MultiCloudNetworkDirectConnectCrossRegionConnectsInnerCost.md) |  | [optional] 
**Status** | Pointer to **string** | Operational status of the cross-region connection. | [optional] 

## Methods

### NewMultiCloudNetworkDirectConnectCrossRegionConnectsInner

`func NewMultiCloudNetworkDirectConnectCrossRegionConnectsInner() *MultiCloudNetworkDirectConnectCrossRegionConnectsInner`

NewMultiCloudNetworkDirectConnectCrossRegionConnectsInner instantiates a new MultiCloudNetworkDirectConnectCrossRegionConnectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudNetworkDirectConnectCrossRegionConnectsInnerWithDefaults

`func NewMultiCloudNetworkDirectConnectCrossRegionConnectsInnerWithDefaults() *MultiCloudNetworkDirectConnectCrossRegionConnectsInner`

NewMultiCloudNetworkDirectConnectCrossRegionConnectsInnerWithDefaults instantiates a new MultiCloudNetworkDirectConnectCrossRegionConnectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworks

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetNetworks() []MultiCloudNetworkDirectConnectCrossRegionConnectsInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetNetworksOk() (*[]MultiCloudNetworkDirectConnectCrossRegionConnectsInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) SetNetworks(v []MultiCloudNetworkDirectConnectCrossRegionConnectsInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetBandwidth

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetBandwidth() MultiCloudNetworkDirectConnectCrossRegionConnectsInnerBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetBandwidthOk() (*MultiCloudNetworkDirectConnectCrossRegionConnectsInnerBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) SetBandwidth(v MultiCloudNetworkDirectConnectCrossRegionConnectsInnerBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetCost

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetCost() MultiCloudNetworkDirectConnectCrossRegionConnectsInnerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetCostOk() (*MultiCloudNetworkDirectConnectCrossRegionConnectsInnerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) SetCost(v MultiCloudNetworkDirectConnectCrossRegionConnectsInnerCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetStatus

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MultiCloudNetworkDirectConnectCrossRegionConnectsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


