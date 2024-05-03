# SshKeyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** |  | 
**KeyType** | **interface{}** |  | 

## Methods

### NewSshKeyCreate

`func NewSshKeyCreate(name interface{}, keyType interface{}, ) *SshKeyCreate`

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

`func (o *SshKeyCreate) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeyCreate) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeyCreate) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *SshKeyCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SshKeyCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKeyType

`func (o *SshKeyCreate) GetKeyType() interface{}`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshKeyCreate) GetKeyTypeOk() (*interface{}, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshKeyCreate) SetKeyType(v interface{})`

SetKeyType sets KeyType field to given value.


### SetKeyTypeNil

`func (o *SshKeyCreate) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *SshKeyCreate) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


