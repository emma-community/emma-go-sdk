# MultiCloudNetworkDirectConnectNetworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network ID | [optional] 
**MacroRegion** | Pointer to **string** | Macro region of the network | [optional] 
**ConnectivityCenter** | Pointer to [**MultiCloudNetworkDirectConnectNetworksInnerConnectivityCenter**](MultiCloudNetworkDirectConnectNetworksInnerConnectivityCenter.md) |  | [optional] 
**Bandwidth** | Pointer to [**MultiCloudNetworkDirectConnectNetworksInnerBandwidth**](MultiCloudNetworkDirectConnectNetworksInnerBandwidth.md) |  | [optional] 
**Cost** | Pointer to [**MultiCloudNetworkDirectConnectNetworksInnerCost**](MultiCloudNetworkDirectConnectNetworksInnerCost.md) |  | [optional] 
**Status** | Pointer to **string** | Operational status of the network. | [optional] 
**CloudConnects** | Pointer to [**[]MultiCloudNetworkDirectConnectNetworksInnerCloudConnectsInner**](MultiCloudNetworkDirectConnectNetworksInnerCloudConnectsInner.md) | Cloud provider connections available from this Direct Connect network. | [optional] 

## Methods

### NewMultiCloudNetworkDirectConnectNetworksInner

`func NewMultiCloudNetworkDirectConnectNetworksInner() *MultiCloudNetworkDirectConnectNetworksInner`

NewMultiCloudNetworkDirectConnectNetworksInner instantiates a new MultiCloudNetworkDirectConnectNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudNetworkDirectConnectNetworksInnerWithDefaults

`func NewMultiCloudNetworkDirectConnectNetworksInnerWithDefaults() *MultiCloudNetworkDirectConnectNetworksInner`

NewMultiCloudNetworkDirectConnectNetworksInnerWithDefaults instantiates a new MultiCloudNetworkDirectConnectNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultiCloudNetworkDirectConnectNetworksInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MultiCloudNetworkDirectConnectNetworksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMacroRegion

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetMacroRegion() string`

GetMacroRegion returns the MacroRegion field if non-nil, zero value otherwise.

### GetMacroRegionOk

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetMacroRegionOk() (*string, bool)`

GetMacroRegionOk returns a tuple with the MacroRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegion

`func (o *MultiCloudNetworkDirectConnectNetworksInner) SetMacroRegion(v string)`

SetMacroRegion sets MacroRegion field to given value.

### HasMacroRegion

`func (o *MultiCloudNetworkDirectConnectNetworksInner) HasMacroRegion() bool`

HasMacroRegion returns a boolean if a field has been set.

### GetConnectivityCenter

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetConnectivityCenter() MultiCloudNetworkDirectConnectNetworksInnerConnectivityCenter`

GetConnectivityCenter returns the ConnectivityCenter field if non-nil, zero value otherwise.

### GetConnectivityCenterOk

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetConnectivityCenterOk() (*MultiCloudNetworkDirectConnectNetworksInnerConnectivityCenter, bool)`

GetConnectivityCenterOk returns a tuple with the ConnectivityCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityCenter

`func (o *MultiCloudNetworkDirectConnectNetworksInner) SetConnectivityCenter(v MultiCloudNetworkDirectConnectNetworksInnerConnectivityCenter)`

SetConnectivityCenter sets ConnectivityCenter field to given value.

### HasConnectivityCenter

`func (o *MultiCloudNetworkDirectConnectNetworksInner) HasConnectivityCenter() bool`

HasConnectivityCenter returns a boolean if a field has been set.

### GetBandwidth

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetBandwidth() MultiCloudNetworkDirectConnectNetworksInnerBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetBandwidthOk() (*MultiCloudNetworkDirectConnectNetworksInnerBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *MultiCloudNetworkDirectConnectNetworksInner) SetBandwidth(v MultiCloudNetworkDirectConnectNetworksInnerBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *MultiCloudNetworkDirectConnectNetworksInner) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetCost

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetCost() MultiCloudNetworkDirectConnectNetworksInnerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetCostOk() (*MultiCloudNetworkDirectConnectNetworksInnerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MultiCloudNetworkDirectConnectNetworksInner) SetCost(v MultiCloudNetworkDirectConnectNetworksInnerCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MultiCloudNetworkDirectConnectNetworksInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetStatus

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MultiCloudNetworkDirectConnectNetworksInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MultiCloudNetworkDirectConnectNetworksInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudConnects

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetCloudConnects() []MultiCloudNetworkDirectConnectNetworksInnerCloudConnectsInner`

GetCloudConnects returns the CloudConnects field if non-nil, zero value otherwise.

### GetCloudConnectsOk

`func (o *MultiCloudNetworkDirectConnectNetworksInner) GetCloudConnectsOk() (*[]MultiCloudNetworkDirectConnectNetworksInnerCloudConnectsInner, bool)`

GetCloudConnectsOk returns a tuple with the CloudConnects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudConnects

`func (o *MultiCloudNetworkDirectConnectNetworksInner) SetCloudConnects(v []MultiCloudNetworkDirectConnectNetworksInnerCloudConnectsInner)`

SetCloudConnects sets CloudConnects field to given value.

### HasCloudConnects

`func (o *MultiCloudNetworkDirectConnectNetworksInner) HasCloudConnects() bool`

HasCloudConnects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


