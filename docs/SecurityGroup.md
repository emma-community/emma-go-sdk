# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByName** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **int32** |  | [optional] 
**ModifiedByName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**SynchronizationStatus** | Pointer to **string** |  | [optional] 
**RecomposingStatus** | Pointer to **string** |  | [optional] 
**LastModificationErrorDescription** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) |  | [optional] 

## Methods

### NewSecurityGroup

`func NewSecurityGroup() *SecurityGroup`

NewSecurityGroup instantiates a new SecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupWithDefaults

`func NewSecurityGroupWithDefaults() *SecurityGroup`

NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SecurityGroup) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SecurityGroup) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SecurityGroup) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SecurityGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByName

`func (o *SecurityGroup) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *SecurityGroup) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *SecurityGroup) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *SecurityGroup) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetModifiedBy

`func (o *SecurityGroup) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SecurityGroup) GetModifiedByOk() (*int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SecurityGroup) SetModifiedBy(v int32)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *SecurityGroup) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedByName

`func (o *SecurityGroup) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *SecurityGroup) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *SecurityGroup) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *SecurityGroup) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecurityGroup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityGroup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityGroup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SecurityGroup) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SecurityGroup) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SecurityGroup) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SecurityGroup) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetSynchronizationStatus

`func (o *SecurityGroup) GetSynchronizationStatus() string`

GetSynchronizationStatus returns the SynchronizationStatus field if non-nil, zero value otherwise.

### GetSynchronizationStatusOk

`func (o *SecurityGroup) GetSynchronizationStatusOk() (*string, bool)`

GetSynchronizationStatusOk returns a tuple with the SynchronizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationStatus

`func (o *SecurityGroup) SetSynchronizationStatus(v string)`

SetSynchronizationStatus sets SynchronizationStatus field to given value.

### HasSynchronizationStatus

`func (o *SecurityGroup) HasSynchronizationStatus() bool`

HasSynchronizationStatus returns a boolean if a field has been set.

### GetRecomposingStatus

`func (o *SecurityGroup) GetRecomposingStatus() string`

GetRecomposingStatus returns the RecomposingStatus field if non-nil, zero value otherwise.

### GetRecomposingStatusOk

`func (o *SecurityGroup) GetRecomposingStatusOk() (*string, bool)`

GetRecomposingStatusOk returns a tuple with the RecomposingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecomposingStatus

`func (o *SecurityGroup) SetRecomposingStatus(v string)`

SetRecomposingStatus sets RecomposingStatus field to given value.

### HasRecomposingStatus

`func (o *SecurityGroup) HasRecomposingStatus() bool`

HasRecomposingStatus returns a boolean if a field has been set.

### GetLastModificationErrorDescription

`func (o *SecurityGroup) GetLastModificationErrorDescription() string`

GetLastModificationErrorDescription returns the LastModificationErrorDescription field if non-nil, zero value otherwise.

### GetLastModificationErrorDescriptionOk

`func (o *SecurityGroup) GetLastModificationErrorDescriptionOk() (*string, bool)`

GetLastModificationErrorDescriptionOk returns a tuple with the LastModificationErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationErrorDescription

`func (o *SecurityGroup) SetLastModificationErrorDescription(v string)`

SetLastModificationErrorDescription sets LastModificationErrorDescription field to given value.

### HasLastModificationErrorDescription

`func (o *SecurityGroup) HasLastModificationErrorDescription() bool`

HasLastModificationErrorDescription returns a boolean if a field has been set.

### GetRules

`func (o *SecurityGroup) GetRules() []SecurityGroupRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroup) GetRulesOk() (*[]SecurityGroupRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroup) SetRules(v []SecurityGroupRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


