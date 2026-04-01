# McnDirectCloudConnectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | **string** | Network connectivity type for this request. Only Direct Connect is supported here. | 
**MacroRegion** | **string** | Macro region that identifies the parent network containing the cloud connection. | 
**CloudConnectionProvider** | **string** | Cloud provider of the cloud connection to be enabled or disabled. | 
**Action** | **string** | Action to apply to the target resource. | 

## Methods

### NewMcnDirectCloudConnectAction

`func NewMcnDirectCloudConnectAction(networkType string, macroRegion string, cloudConnectionProvider string, action string, ) *McnDirectCloudConnectAction`

NewMcnDirectCloudConnectAction instantiates a new McnDirectCloudConnectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcnDirectCloudConnectActionWithDefaults

`func NewMcnDirectCloudConnectActionWithDefaults() *McnDirectCloudConnectAction`

NewMcnDirectCloudConnectActionWithDefaults instantiates a new McnDirectCloudConnectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *McnDirectCloudConnectAction) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *McnDirectCloudConnectAction) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *McnDirectCloudConnectAction) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetMacroRegion

`func (o *McnDirectCloudConnectAction) GetMacroRegion() string`

GetMacroRegion returns the MacroRegion field if non-nil, zero value otherwise.

### GetMacroRegionOk

`func (o *McnDirectCloudConnectAction) GetMacroRegionOk() (*string, bool)`

GetMacroRegionOk returns a tuple with the MacroRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegion

`func (o *McnDirectCloudConnectAction) SetMacroRegion(v string)`

SetMacroRegion sets MacroRegion field to given value.


### GetCloudConnectionProvider

`func (o *McnDirectCloudConnectAction) GetCloudConnectionProvider() string`

GetCloudConnectionProvider returns the CloudConnectionProvider field if non-nil, zero value otherwise.

### GetCloudConnectionProviderOk

`func (o *McnDirectCloudConnectAction) GetCloudConnectionProviderOk() (*string, bool)`

GetCloudConnectionProviderOk returns a tuple with the CloudConnectionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudConnectionProvider

`func (o *McnDirectCloudConnectAction) SetCloudConnectionProvider(v string)`

SetCloudConnectionProvider sets CloudConnectionProvider field to given value.


### GetAction

`func (o *McnDirectCloudConnectAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *McnDirectCloudConnectAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *McnDirectCloudConnectAction) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


