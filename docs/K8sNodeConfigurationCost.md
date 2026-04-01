# K8sNodeConfigurationCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Cost period | [optional] 
**Currency** | Pointer to **string** | Currency of the cost | [optional] 
**PricePerUnit** | Pointer to **float32** | Cost for the period | [optional] 

## Methods

### NewK8sNodeConfigurationCost

`func NewK8sNodeConfigurationCost() *K8sNodeConfigurationCost`

NewK8sNodeConfigurationCost instantiates a new K8sNodeConfigurationCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sNodeConfigurationCostWithDefaults

`func NewK8sNodeConfigurationCostWithDefaults() *K8sNodeConfigurationCost`

NewK8sNodeConfigurationCostWithDefaults instantiates a new K8sNodeConfigurationCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *K8sNodeConfigurationCost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *K8sNodeConfigurationCost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *K8sNodeConfigurationCost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *K8sNodeConfigurationCost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *K8sNodeConfigurationCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *K8sNodeConfigurationCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *K8sNodeConfigurationCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *K8sNodeConfigurationCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *K8sNodeConfigurationCost) GetPricePerUnit() float32`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *K8sNodeConfigurationCost) GetPricePerUnitOk() (*float32, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *K8sNodeConfigurationCost) SetPricePerUnit(v float32)`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *K8sNodeConfigurationCost) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


