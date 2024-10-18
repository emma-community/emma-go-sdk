# ProjectSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | emma&#39;s internal service type | [optional] 
**AllInstallations** | Pointer to **int32** | All products in the service | [optional] 
**Cost** | Pointer to **float32** | Product monthly cost | [optional] 

## Methods

### NewProjectSummaryResponse

`func NewProjectSummaryResponse() *ProjectSummaryResponse`

NewProjectSummaryResponse instantiates a new ProjectSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSummaryResponseWithDefaults

`func NewProjectSummaryResponseWithDefaults() *ProjectSummaryResponse`

NewProjectSummaryResponseWithDefaults instantiates a new ProjectSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ProjectSummaryResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProjectSummaryResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProjectSummaryResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ProjectSummaryResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetAllInstallations

`func (o *ProjectSummaryResponse) GetAllInstallations() int32`

GetAllInstallations returns the AllInstallations field if non-nil, zero value otherwise.

### GetAllInstallationsOk

`func (o *ProjectSummaryResponse) GetAllInstallationsOk() (*int32, bool)`

GetAllInstallationsOk returns a tuple with the AllInstallations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllInstallations

`func (o *ProjectSummaryResponse) SetAllInstallations(v int32)`

SetAllInstallations sets AllInstallations field to given value.

### HasAllInstallations

`func (o *ProjectSummaryResponse) HasAllInstallations() bool`

HasAllInstallations returns a boolean if a field has been set.

### GetCost

`func (o *ProjectSummaryResponse) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProjectSummaryResponse) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProjectSummaryResponse) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProjectSummaryResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


