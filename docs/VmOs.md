# VmOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the operating system | [optional] 
**Family** | Pointer to **string** | Family of the operating system | [optional] 
**Architecture** | Pointer to **string** | Architecture of the operating system | [optional] 
**Type** | Pointer to **string** | Type of the operating system | [optional] 
**Version** | Pointer to **string** | Version of the operating system | [optional] 

## Methods

### NewVmOs

`func NewVmOs() *VmOs`

NewVmOs instantiates a new VmOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmOsWithDefaults

`func NewVmOsWithDefaults() *VmOs`

NewVmOsWithDefaults instantiates a new VmOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmOs) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmOs) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmOs) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VmOs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFamily

`func (o *VmOs) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *VmOs) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *VmOs) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *VmOs) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetArchitecture

`func (o *VmOs) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *VmOs) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *VmOs) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *VmOs) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetType

`func (o *VmOs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmOs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmOs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VmOs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *VmOs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VmOs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VmOs) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VmOs) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


