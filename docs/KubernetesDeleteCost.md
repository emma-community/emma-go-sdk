# KubernetesDeleteCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Billing unit | [optional] 
**Currency** | Pointer to **string** | Currency code | [optional] 
**Price** | Pointer to **float64** | Cost amount | [optional] 

## Methods

### NewKubernetesDeleteCost

`func NewKubernetesDeleteCost() *KubernetesDeleteCost`

NewKubernetesDeleteCost instantiates a new KubernetesDeleteCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDeleteCostWithDefaults

`func NewKubernetesDeleteCostWithDefaults() *KubernetesDeleteCost`

NewKubernetesDeleteCostWithDefaults instantiates a new KubernetesDeleteCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *KubernetesDeleteCost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *KubernetesDeleteCost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *KubernetesDeleteCost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *KubernetesDeleteCost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *KubernetesDeleteCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *KubernetesDeleteCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *KubernetesDeleteCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *KubernetesDeleteCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrice

`func (o *KubernetesDeleteCost) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *KubernetesDeleteCost) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *KubernetesDeleteCost) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *KubernetesDeleteCost) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


