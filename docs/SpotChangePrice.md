# SpotChangePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with the spot instance - changeprice | 
**NewSpotPrice** | **float32** | New price for the spot instance | 

## Methods

### NewSpotChangePrice

`func NewSpotChangePrice(action string, newSpotPrice float32, ) *SpotChangePrice`

NewSpotChangePrice instantiates a new SpotChangePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotChangePriceWithDefaults

`func NewSpotChangePriceWithDefaults() *SpotChangePrice`

NewSpotChangePriceWithDefaults instantiates a new SpotChangePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SpotChangePrice) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SpotChangePrice) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SpotChangePrice) SetAction(v string)`

SetAction sets Action field to given value.


### GetNewSpotPrice

`func (o *SpotChangePrice) GetNewSpotPrice() float32`

GetNewSpotPrice returns the NewSpotPrice field if non-nil, zero value otherwise.

### GetNewSpotPriceOk

`func (o *SpotChangePrice) GetNewSpotPriceOk() (*float32, bool)`

GetNewSpotPriceOk returns a tuple with the NewSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSpotPrice

`func (o *SpotChangePrice) SetNewSpotPrice(v float32)`

SetNewSpotPrice sets NewSpotPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


