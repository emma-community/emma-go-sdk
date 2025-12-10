# VolumeCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Cost period | [optional] 
**Currency** | Pointer to **string** | Currency of cost | [optional] 
**Price** | Pointer to **float32** | Cost of the volume for the period | [optional] 

## Methods

### NewVolumeCost

`func NewVolumeCost() *VolumeCost`

NewVolumeCost instantiates a new VolumeCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCostWithDefaults

`func NewVolumeCostWithDefaults() *VolumeCost`

NewVolumeCostWithDefaults instantiates a new VolumeCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *VolumeCost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *VolumeCost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *VolumeCost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *VolumeCost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *VolumeCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VolumeCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VolumeCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VolumeCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrice

`func (o *VolumeCost) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *VolumeCost) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *VolumeCost) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *VolumeCost) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


