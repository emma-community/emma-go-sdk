# SpotActionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with the spot instance - rename | 
**NewSpotPrice** | **float32** | New price for the spot instance | 
**Name** | **string** | A new name of the spot instance | 

## Methods

### NewSpotActionsRequest

`func NewSpotActionsRequest(action string, newSpotPrice float32, name string, ) *SpotActionsRequest`

NewSpotActionsRequest instantiates a new SpotActionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotActionsRequestWithDefaults

`func NewSpotActionsRequestWithDefaults() *SpotActionsRequest`

NewSpotActionsRequestWithDefaults instantiates a new SpotActionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SpotActionsRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SpotActionsRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SpotActionsRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetNewSpotPrice

`func (o *SpotActionsRequest) GetNewSpotPrice() float32`

GetNewSpotPrice returns the NewSpotPrice field if non-nil, zero value otherwise.

### GetNewSpotPriceOk

`func (o *SpotActionsRequest) GetNewSpotPriceOk() (*float32, bool)`

GetNewSpotPriceOk returns a tuple with the NewSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSpotPrice

`func (o *SpotActionsRequest) SetNewSpotPrice(v float32)`

SetNewSpotPrice sets NewSpotPrice field to given value.


### GetName

`func (o *SpotActionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpotActionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpotActionsRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


