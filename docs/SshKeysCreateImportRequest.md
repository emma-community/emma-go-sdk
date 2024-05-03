# SshKeysCreateImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**KeyType** | **string** |  | 
**Key** | **string** |  | 

## Methods

### NewSshKeysCreateImportRequest

`func NewSshKeysCreateImportRequest(name string, keyType string, key string, ) *SshKeysCreateImportRequest`

NewSshKeysCreateImportRequest instantiates a new SshKeysCreateImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeysCreateImportRequestWithDefaults

`func NewSshKeysCreateImportRequestWithDefaults() *SshKeysCreateImportRequest`

NewSshKeysCreateImportRequestWithDefaults instantiates a new SshKeysCreateImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshKeysCreateImportRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeysCreateImportRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeysCreateImportRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKeyType

`func (o *SshKeysCreateImportRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshKeysCreateImportRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshKeysCreateImportRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKey

`func (o *SshKeysCreateImportRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeysCreateImportRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeysCreateImportRequest) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


