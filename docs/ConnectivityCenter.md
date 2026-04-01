# ConnectivityCenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Connectivity center ID | 
**Location** | **string** | Connectivity center location | 
**MacroRegion** | **string** | Macro region of the connectivity center | 
**NetworkType** | **string** | Network type supported by the connectivity center | 

## Methods

### NewConnectivityCenter

`func NewConnectivityCenter(id int32, location string, macroRegion string, networkType string, ) *ConnectivityCenter`

NewConnectivityCenter instantiates a new ConnectivityCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityCenterWithDefaults

`func NewConnectivityCenterWithDefaults() *ConnectivityCenter`

NewConnectivityCenterWithDefaults instantiates a new ConnectivityCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectivityCenter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectivityCenter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectivityCenter) SetId(v int32)`

SetId sets Id field to given value.


### GetLocation

`func (o *ConnectivityCenter) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ConnectivityCenter) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ConnectivityCenter) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetMacroRegion

`func (o *ConnectivityCenter) GetMacroRegion() string`

GetMacroRegion returns the MacroRegion field if non-nil, zero value otherwise.

### GetMacroRegionOk

`func (o *ConnectivityCenter) GetMacroRegionOk() (*string, bool)`

GetMacroRegionOk returns a tuple with the MacroRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroRegion

`func (o *ConnectivityCenter) SetMacroRegion(v string)`

SetMacroRegion sets MacroRegion field to given value.


### GetNetworkType

`func (o *ConnectivityCenter) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *ConnectivityCenter) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *ConnectivityCenter) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


