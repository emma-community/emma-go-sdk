# K8sNodeConfigurationLocalDisks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | Pointer to **int32** | Number of local disks attached to the VM | [optional] 
**SizeTb** | Pointer to **float32** | Size of a single disk in TB | [optional] 
**TotalSizeTb** | Pointer to **float32** | Total capacity (all disks combined), in TB | [optional] 
**Type** | Pointer to **string** | Disk type | [optional] 

## Methods

### NewK8sNodeConfigurationLocalDisks

`func NewK8sNodeConfigurationLocalDisks() *K8sNodeConfigurationLocalDisks`

NewK8sNodeConfigurationLocalDisks instantiates a new K8sNodeConfigurationLocalDisks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sNodeConfigurationLocalDisksWithDefaults

`func NewK8sNodeConfigurationLocalDisksWithDefaults() *K8sNodeConfigurationLocalDisks`

NewK8sNodeConfigurationLocalDisksWithDefaults instantiates a new K8sNodeConfigurationLocalDisks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisks

`func (o *K8sNodeConfigurationLocalDisks) GetDisks() int32`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *K8sNodeConfigurationLocalDisks) GetDisksOk() (*int32, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *K8sNodeConfigurationLocalDisks) SetDisks(v int32)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *K8sNodeConfigurationLocalDisks) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetSizeTb

`func (o *K8sNodeConfigurationLocalDisks) GetSizeTb() float32`

GetSizeTb returns the SizeTb field if non-nil, zero value otherwise.

### GetSizeTbOk

`func (o *K8sNodeConfigurationLocalDisks) GetSizeTbOk() (*float32, bool)`

GetSizeTbOk returns a tuple with the SizeTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTb

`func (o *K8sNodeConfigurationLocalDisks) SetSizeTb(v float32)`

SetSizeTb sets SizeTb field to given value.

### HasSizeTb

`func (o *K8sNodeConfigurationLocalDisks) HasSizeTb() bool`

HasSizeTb returns a boolean if a field has been set.

### GetTotalSizeTb

`func (o *K8sNodeConfigurationLocalDisks) GetTotalSizeTb() float32`

GetTotalSizeTb returns the TotalSizeTb field if non-nil, zero value otherwise.

### GetTotalSizeTbOk

`func (o *K8sNodeConfigurationLocalDisks) GetTotalSizeTbOk() (*float32, bool)`

GetTotalSizeTbOk returns a tuple with the TotalSizeTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeTb

`func (o *K8sNodeConfigurationLocalDisks) SetTotalSizeTb(v float32)`

SetTotalSizeTb sets TotalSizeTb field to given value.

### HasTotalSizeTb

`func (o *K8sNodeConfigurationLocalDisks) HasTotalSizeTb() bool`

HasTotalSizeTb returns a boolean if a field has been set.

### GetType

`func (o *K8sNodeConfigurationLocalDisks) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *K8sNodeConfigurationLocalDisks) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *K8sNodeConfigurationLocalDisks) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *K8sNodeConfigurationLocalDisks) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


