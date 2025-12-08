# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Volume ID | [optional] 
**Name** | Pointer to **string** | Volume name | [optional] 
**AttachedToId** | Pointer to **int32** | Compute instance ID that the volume is currently attached to | [optional] 
**ProjectId** | Pointer to **int32** | Project ID | [optional] 
**SizeGb** | Pointer to **int32** | Volume size in gigabytes | [optional] 
**Type** | Pointer to **string** | Volume type | [optional] 
**IsSystem** | Pointer to **bool** | Indicates whether the volume contains OS or not | [optional] 
**Status** | Pointer to **string** | Status of the volume | [optional] 
**Provider** | Pointer to [**VolumeProvider**](VolumeProvider.md) |  | [optional] 
**Location** | Pointer to [**VolumeLocation**](VolumeLocation.md) |  | [optional] 
**DataCenter** | Pointer to [**VolumeDataCenter**](VolumeDataCenter.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Date and time when the volume was created | [optional] 
**CreatedByName** | Pointer to **string** | Name of the user who created the volume | [optional] 
**CreatedById** | Pointer to **int32** | ID of the user who created the volume | [optional] 
**ModifiedAt** | Pointer to **string** | Date and time when the volume was last edited | [optional] 
**ModifiedByName** | Pointer to **string** | Name of the user who last edited the volume | [optional] 
**ModifiedById** | Pointer to **int32** | ID of the user who last edited the volume | [optional] 
**Cost** | Pointer to [**VolumeCost**](VolumeCost.md) |  | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Volume) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Volume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttachedToId

`func (o *Volume) GetAttachedToId() int32`

GetAttachedToId returns the AttachedToId field if non-nil, zero value otherwise.

### GetAttachedToIdOk

`func (o *Volume) GetAttachedToIdOk() (*int32, bool)`

GetAttachedToIdOk returns a tuple with the AttachedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedToId

`func (o *Volume) SetAttachedToId(v int32)`

SetAttachedToId sets AttachedToId field to given value.

### HasAttachedToId

`func (o *Volume) HasAttachedToId() bool`

HasAttachedToId returns a boolean if a field has been set.

### GetProjectId

`func (o *Volume) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Volume) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Volume) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Volume) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSizeGb

`func (o *Volume) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *Volume) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *Volume) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *Volume) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetType

`func (o *Volume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Volume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Volume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Volume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsSystem

`func (o *Volume) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *Volume) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *Volume) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *Volume) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetStatus

`func (o *Volume) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Volume) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Volume) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Volume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *Volume) GetProvider() VolumeProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Volume) GetProviderOk() (*VolumeProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Volume) SetProvider(v VolumeProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Volume) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *Volume) GetLocation() VolumeLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Volume) GetLocationOk() (*VolumeLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Volume) SetLocation(v VolumeLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Volume) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDataCenter

`func (o *Volume) GetDataCenter() VolumeDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *Volume) GetDataCenterOk() (*VolumeDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *Volume) SetDataCenter(v VolumeDataCenter)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *Volume) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Volume) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Volume) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Volume) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Volume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *Volume) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *Volume) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *Volume) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *Volume) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedById

`func (o *Volume) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Volume) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Volume) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Volume) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Volume) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Volume) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Volume) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Volume) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByName

`func (o *Volume) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *Volume) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *Volume) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *Volume) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedById

`func (o *Volume) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *Volume) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *Volume) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *Volume) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetCost

`func (o *Volume) GetCost() VolumeCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Volume) GetCostOk() (*VolumeCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Volume) SetCost(v VolumeCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Volume) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


