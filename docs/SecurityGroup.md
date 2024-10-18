# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Security group id | [optional] 
**Name** | Pointer to **string** | Security group name | [optional] 
**CreatedById** | Pointer to **int32** | ID of the user who created the security group | [optional] 
**CreatedByName** | Pointer to **string** | Name of the user who created the security group | [optional] 
**ModifiedById** | Pointer to **int32** | ID of the user who last edited the security group | [optional] 
**ModifiedByName** | Pointer to **string** | Name of the user who last edited the security group | [optional] 
**CreatedAt** | Pointer to **string** | Date and time of the security group&#39;s creation | [optional] 
**ModifiedAt** | Pointer to **string** | Date and time of the security group&#39;s last update | [optional] 
**SynchronizationStatus** | Pointer to **string** | Synchronization status of the security group. When you make changes in the rules the changes are propagated to the respective provider’s security groups. While this is happening the security groups have the status Synchronizing. After it is done the status changes to Synchronized. When another VM is added to the security group it will not be synchronized at first with the other VMs, therefore the status will be Desynchronized. | [optional] 
**RecomposingStatus** | Pointer to **string** | Recomposing status of the security group. When a new Virtual machine is added to the Security group it starts a synchronization process. During this process the Security group will have a Recomposing status. | [optional] 
**LastModificationErrorDescription** | Pointer to **string** | Text of the error when the Security group was last edited | [optional] 
**Rules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | List of the inbound and outbound rules in the Security group | [optional] 

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

### GetCreatedById

`func (o *SecurityGroup) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SecurityGroup) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SecurityGroup) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SecurityGroup) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

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

### GetModifiedById

`func (o *SecurityGroup) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *SecurityGroup) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *SecurityGroup) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *SecurityGroup) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

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


