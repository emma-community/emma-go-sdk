# McnDirectNetworkEnable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | **string** | Network connectivity type. Only Direct Connect is supported here. | 
**MacroRegion** | **string** | Macro region that identifies the target network. | 
**ConnectivityCenterId** | **int32** | Connectivity Center identifier (city) selected when the network is created. Required for enabling a network. | 
**Action** | **string** | Action to apply to the target resource. | 

## Methods

### NewMcnDirectNetworkEnable

`func NewMcnDirectNetworkEnable(networkType string, macroRegion string, connectivityCenterId int32, action string, ) *McnDirectNetworkEnable`

NewMcnDirectNetworkEnable instantiates a new McnDirectNetworkEnable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcnDirectNetworkEnableWithDefaults

`func NewMcnDirectNetworkEnableWithDefaults() *McnDirectNetworkEnable`

NewMcnDirectNetworkEnableWithDefaults instantiates a new McnDirectNetworkEnable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *McnDirectNetworkEnable) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *McnDirectNetworkEnable) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *McnDirectNetworkEnable) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetMacroRegion

`func (o *McnDirectNetworkEnable) GetMacroRegion() string`

GetMacroRegion returns the MacroRegion field if non-nil, zero value otherwise.

### GetMacroRegionOk

`func (o *McnDirectNetworkEnable) GetMacroRegionOk() (*string, bool)`

GetMacroRegionOk returns a tuple with the MacroRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegion

`func (o *McnDirectNetworkEnable) SetMacroRegion(v string)`

SetMacroRegion sets MacroRegion field to given value.


### GetConnectivityCenterId

`func (o *McnDirectNetworkEnable) GetConnectivityCenterId() int32`

GetConnectivityCenterId returns the ConnectivityCenterId field if non-nil, zero value otherwise.

### GetConnectivityCenterIdOk

`func (o *McnDirectNetworkEnable) GetConnectivityCenterIdOk() (*int32, bool)`

GetConnectivityCenterIdOk returns a tuple with the ConnectivityCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityCenterId

`func (o *McnDirectNetworkEnable) SetConnectivityCenterId(v int32)`

SetConnectivityCenterId sets ConnectivityCenterId field to given value.


### GetAction

`func (o *McnDirectNetworkEnable) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *McnDirectNetworkEnable) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *McnDirectNetworkEnable) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


