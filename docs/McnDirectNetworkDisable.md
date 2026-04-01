# McnDirectNetworkDisable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | **string** | Network connectivity type for this request. Only Direct Connect is supported here. | 
**MacroRegion** | **string** | Macro region that identifies the target network. | 
**Action** | **string** | Action to apply to the target resource. | 

## Methods

### NewMcnDirectNetworkDisable

`func NewMcnDirectNetworkDisable(networkType string, macroRegion string, action string, ) *McnDirectNetworkDisable`

NewMcnDirectNetworkDisable instantiates a new McnDirectNetworkDisable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcnDirectNetworkDisableWithDefaults

`func NewMcnDirectNetworkDisableWithDefaults() *McnDirectNetworkDisable`

NewMcnDirectNetworkDisableWithDefaults instantiates a new McnDirectNetworkDisable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *McnDirectNetworkDisable) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *McnDirectNetworkDisable) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *McnDirectNetworkDisable) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetMacroRegion

`func (o *McnDirectNetworkDisable) GetMacroRegion() string`

GetMacroRegion returns the MacroRegion field if non-nil, zero value otherwise.

### GetMacroRegionOk

`func (o *McnDirectNetworkDisable) GetMacroRegionOk() (*string, bool)`

GetMacroRegionOk returns a tuple with the MacroRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegion

`func (o *McnDirectNetworkDisable) SetMacroRegion(v string)`

SetMacroRegion sets MacroRegion field to given value.


### GetAction

`func (o *McnDirectNetworkDisable) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *McnDirectNetworkDisable) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *McnDirectNetworkDisable) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


