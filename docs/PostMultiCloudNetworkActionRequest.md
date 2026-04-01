# PostMultiCloudNetworkActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | **string** | Network connectivity type for this request. Only Direct Connect is supported here. | 
**MacroRegion** | **string** | Macro region that identifies the parent network containing the cloud connection. | 
**ConnectivityCenterId** | **int32** | Connectivity Center identifier (city) selected when the network is created. Required for enabling a network. | 
**Action** | **string** | Action to apply to the target resource. | 
**CloudConnectionProvider** | **string** | Cloud provider of the cloud connection to be enabled or disabled. | 
**MacroRegions** | **[]string** | Exactly two macro regions that form the cross-region connection to be enabled or disabled. | 

## Methods

### NewPostMultiCloudNetworkActionRequest

`func NewPostMultiCloudNetworkActionRequest(networkType string, macroRegion string, connectivityCenterId int32, action string, cloudConnectionProvider string, macroRegions []string, ) *PostMultiCloudNetworkActionRequest`

NewPostMultiCloudNetworkActionRequest instantiates a new PostMultiCloudNetworkActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMultiCloudNetworkActionRequestWithDefaults

`func NewPostMultiCloudNetworkActionRequestWithDefaults() *PostMultiCloudNetworkActionRequest`

NewPostMultiCloudNetworkActionRequestWithDefaults instantiates a new PostMultiCloudNetworkActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *PostMultiCloudNetworkActionRequest) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *PostMultiCloudNetworkActionRequest) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *PostMultiCloudNetworkActionRequest) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetMacroRegion

`func (o *PostMultiCloudNetworkActionRequest) GetMacroRegion() string`

GetMacroRegion returns the MacroRegion field if non-nil, zero value otherwise.

### GetMacroRegionOk

`func (o *PostMultiCloudNetworkActionRequest) GetMacroRegionOk() (*string, bool)`

GetMacroRegionOk returns a tuple with the MacroRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegion

`func (o *PostMultiCloudNetworkActionRequest) SetMacroRegion(v string)`

SetMacroRegion sets MacroRegion field to given value.


### GetConnectivityCenterId

`func (o *PostMultiCloudNetworkActionRequest) GetConnectivityCenterId() int32`

GetConnectivityCenterId returns the ConnectivityCenterId field if non-nil, zero value otherwise.

### GetConnectivityCenterIdOk

`func (o *PostMultiCloudNetworkActionRequest) GetConnectivityCenterIdOk() (*int32, bool)`

GetConnectivityCenterIdOk returns a tuple with the ConnectivityCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityCenterId

`func (o *PostMultiCloudNetworkActionRequest) SetConnectivityCenterId(v int32)`

SetConnectivityCenterId sets ConnectivityCenterId field to given value.


### GetAction

`func (o *PostMultiCloudNetworkActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PostMultiCloudNetworkActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PostMultiCloudNetworkActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetCloudConnectionProvider

`func (o *PostMultiCloudNetworkActionRequest) GetCloudConnectionProvider() string`

GetCloudConnectionProvider returns the CloudConnectionProvider field if non-nil, zero value otherwise.

### GetCloudConnectionProviderOk

`func (o *PostMultiCloudNetworkActionRequest) GetCloudConnectionProviderOk() (*string, bool)`

GetCloudConnectionProviderOk returns a tuple with the CloudConnectionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudConnectionProvider

`func (o *PostMultiCloudNetworkActionRequest) SetCloudConnectionProvider(v string)`

SetCloudConnectionProvider sets CloudConnectionProvider field to given value.


### GetMacroRegions

`func (o *PostMultiCloudNetworkActionRequest) GetMacroRegions() []string`

GetMacroRegions returns the MacroRegions field if non-nil, zero value otherwise.

### GetMacroRegionsOk

`func (o *PostMultiCloudNetworkActionRequest) GetMacroRegionsOk() (*[]string, bool)`

GetMacroRegionsOk returns a tuple with the MacroRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegions

`func (o *PostMultiCloudNetworkActionRequest) SetMacroRegions(v []string)`

SetMacroRegions sets MacroRegions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


