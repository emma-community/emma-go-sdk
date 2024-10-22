# KubernetesCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Cost period | [optional] 
**Currency** | Pointer to **string** | Currency of the cost | [optional] 
**Price** | Pointer to **float32** | Cost for the period | [optional] 

## Methods

### NewKubernetesCost

`func NewKubernetesCost() *KubernetesCost`

NewKubernetesCost instantiates a new KubernetesCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCostWithDefaults

`func NewKubernetesCostWithDefaults() *KubernetesCost`

NewKubernetesCostWithDefaults instantiates a new KubernetesCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *KubernetesCost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *KubernetesCost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *KubernetesCost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *KubernetesCost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *KubernetesCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *KubernetesCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *KubernetesCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *KubernetesCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrice

`func (o *KubernetesCost) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *KubernetesCost) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *KubernetesCost) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *KubernetesCost) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


