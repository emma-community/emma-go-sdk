# SshKeyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | SSH key name | 
**KeyType** | **string** | SSH key type (RSA or ED25519) | 

## Methods

### NewSshKeyCreate

`func NewSshKeyCreate(name string, keyType string, ) *SshKeyCreate`

NewSshKeyCreate instantiates a new SshKeyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyCreateWithDefaults

`func NewSshKeyCreateWithDefaults() *SshKeyCreate`

NewSshKeyCreateWithDefaults instantiates a new SshKeyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshKeyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetKeyType

`func (o *SshKeyCreate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshKeyCreate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshKeyCreate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


