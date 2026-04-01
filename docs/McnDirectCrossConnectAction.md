# McnDirectCrossConnectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | **string** | Network connectivity type for this request. Only Direct Connect is supported here. | 
**MacroRegions** | **[]string** | Exactly two macro regions that form the cross-region connection to be enabled or disabled. | 
**Action** | **string** | Action to apply to the target resource. | 

## Methods

### NewMcnDirectCrossConnectAction

`func NewMcnDirectCrossConnectAction(networkType string, macroRegions []string, action string, ) *McnDirectCrossConnectAction`

NewMcnDirectCrossConnectAction instantiates a new McnDirectCrossConnectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcnDirectCrossConnectActionWithDefaults

`func NewMcnDirectCrossConnectActionWithDefaults() *McnDirectCrossConnectAction`

NewMcnDirectCrossConnectActionWithDefaults instantiates a new McnDirectCrossConnectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *McnDirectCrossConnectAction) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *McnDirectCrossConnectAction) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *McnDirectCrossConnectAction) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetMacroRegions

`func (o *McnDirectCrossConnectAction) GetMacroRegions() []string`

GetMacroRegions returns the MacroRegions field if non-nil, zero value otherwise.

### GetMacroRegionsOk

`func (o *McnDirectCrossConnectAction) GetMacroRegionsOk() (*[]string, bool)`

GetMacroRegionsOk returns a tuple with the MacroRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegions

`func (o *McnDirectCrossConnectAction) SetMacroRegions(v []string)`

SetMacroRegions sets MacroRegions field to given value.


### GetAction

`func (o *McnDirectCrossConnectAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *McnDirectCrossConnectAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *McnDirectCrossConnectAction) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


