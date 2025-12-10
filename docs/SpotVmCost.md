# SpotVmCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Cost period | [optional] 
**Currency** | Pointer to **string** | Currency of cost | [optional] 
**Price** | Pointer to **float32** | Cost of the spot instance for the period | [optional] 
**MinPrice** | Pointer to **float32** | Minimal cost of the spot instance for the period | [optional] 
**MaxPrice** | Pointer to **float32** | Maximal cost of the spot instance for the period | [optional] 

## Methods

### NewSpotVmCost

`func NewSpotVmCost() *SpotVmCost`

NewSpotVmCost instantiates a new SpotVmCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotVmCostWithDefaults

`func NewSpotVmCostWithDefaults() *SpotVmCost`

NewSpotVmCostWithDefaults instantiates a new SpotVmCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *SpotVmCost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SpotVmCost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SpotVmCost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SpotVmCost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *SpotVmCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SpotVmCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SpotVmCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SpotVmCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrice

`func (o *SpotVmCost) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpotVmCost) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SpotVmCost) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SpotVmCost) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMinPrice

`func (o *SpotVmCost) GetMinPrice() float32`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *SpotVmCost) GetMinPriceOk() (*float32, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *SpotVmCost) SetMinPrice(v float32)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *SpotVmCost) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *SpotVmCost) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *SpotVmCost) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *SpotVmCost) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *SpotVmCost) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


