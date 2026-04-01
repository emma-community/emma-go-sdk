# MultiCloudNetworkDirectConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]MultiCloudNetworkDirectConnectNetworksInner**](MultiCloudNetworkDirectConnectNetworksInner.md) | Network within macro region (AMER, EMEA, or APAC). | [optional] 
**CrossRegionConnects** | Pointer to [**[]MultiCloudNetworkDirectConnectCrossRegionConnectsInner**](MultiCloudNetworkDirectConnectCrossRegionConnectsInner.md) | Cross-region links between networks in different macro regions. | [optional] 

## Methods

### NewMultiCloudNetworkDirectConnect

`func NewMultiCloudNetworkDirectConnect() *MultiCloudNetworkDirectConnect`

NewMultiCloudNetworkDirectConnect instantiates a new MultiCloudNetworkDirectConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudNetworkDirectConnectWithDefaults

`func NewMultiCloudNetworkDirectConnectWithDefaults() *MultiCloudNetworkDirectConnect`

NewMultiCloudNetworkDirectConnectWithDefaults instantiates a new MultiCloudNetworkDirectConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *MultiCloudNetworkDirectConnect) GetNetworks() []MultiCloudNetworkDirectConnectNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *MultiCloudNetworkDirectConnect) GetNetworksOk() (*[]MultiCloudNetworkDirectConnectNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *MultiCloudNetworkDirectConnect) SetNetworks(v []MultiCloudNetworkDirectConnectNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *MultiCloudNetworkDirectConnect) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetCrossRegionConnects

`func (o *MultiCloudNetworkDirectConnect) GetCrossRegionConnects() []MultiCloudNetworkDirectConnectCrossRegionConnectsInner`

GetCrossRegionConnects returns the CrossRegionConnects field if non-nil, zero value otherwise.

### GetCrossRegionConnectsOk

`func (o *MultiCloudNetworkDirectConnect) GetCrossRegionConnectsOk() (*[]MultiCloudNetworkDirectConnectCrossRegionConnectsInner, bool)`

GetCrossRegionConnectsOk returns a tuple with the CrossRegionConnects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossRegionConnects

`func (o *MultiCloudNetworkDirectConnect) SetCrossRegionConnects(v []MultiCloudNetworkDirectConnectCrossRegionConnectsInner)`

SetCrossRegionConnects sets CrossRegionConnects field to given value.

### HasCrossRegionConnects

`func (o *MultiCloudNetworkDirectConnect) HasCrossRegionConnects() bool`

HasCrossRegionConnects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


