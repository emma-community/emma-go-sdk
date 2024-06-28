# SshKeysCreateImport201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | SSH key ID | [optional] 
**Name** | Pointer to **string** | SSH key name | [optional] 
**Key** | Pointer to **string** | SSH public key | [optional] 
**Fingerprint** | Pointer to **string** | SSH key fingerprint | [optional] 
**KeyType** | Pointer to **string** | SSH key type (RSA or ED25519) | [optional] 
**CreatedAt** | Pointer to **string** | Date and time when the SSH key was created | [optional] 
**CreatedByName** | Pointer to **string** | Name of the user who created the SSH key | [optional] 
**CreatedById** | Pointer to **int32** | ID of the user who created the SSH key | [optional] 
**PrivateKey** | Pointer to **string** | SSH private key | [optional] 

## Methods

### NewSshKeysCreateImport201Response

`func NewSshKeysCreateImport201Response() *SshKeysCreateImport201Response`

NewSshKeysCreateImport201Response instantiates a new SshKeysCreateImport201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeysCreateImport201ResponseWithDefaults

`func NewSshKeysCreateImport201ResponseWithDefaults() *SshKeysCreateImport201Response`

NewSshKeysCreateImport201ResponseWithDefaults instantiates a new SshKeysCreateImport201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SshKeysCreateImport201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshKeysCreateImport201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshKeysCreateImport201Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SshKeysCreateImport201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SshKeysCreateImport201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeysCreateImport201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeysCreateImport201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SshKeysCreateImport201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *SshKeysCreateImport201Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeysCreateImport201Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeysCreateImport201Response) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshKeysCreateImport201Response) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetFingerprint

`func (o *SshKeysCreateImport201Response) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SshKeysCreateImport201Response) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SshKeysCreateImport201Response) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SshKeysCreateImport201Response) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetKeyType

`func (o *SshKeysCreateImport201Response) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshKeysCreateImport201Response) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshKeysCreateImport201Response) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *SshKeysCreateImport201Response) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SshKeysCreateImport201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SshKeysCreateImport201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SshKeysCreateImport201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SshKeysCreateImport201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *SshKeysCreateImport201Response) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *SshKeysCreateImport201Response) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *SshKeysCreateImport201Response) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *SshKeysCreateImport201Response) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedById

`func (o *SshKeysCreateImport201Response) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SshKeysCreateImport201Response) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SshKeysCreateImport201Response) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SshKeysCreateImport201Response) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetPrivateKey

`func (o *SshKeysCreateImport201Response) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SshKeysCreateImport201Response) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SshKeysCreateImport201Response) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SshKeysCreateImport201Response) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


