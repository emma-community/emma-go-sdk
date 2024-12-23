# SpotCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the spot instance | 
**DataCenterId** | **string** | ID of the provider&#39;s data center | 
**OsId** | **int32** | ID of the operating system | 
**CloudNetworkType** | **string** | Type of the cloud network | 
**VCpuType** | **string** | Type of the Central Processing Units (vCPUs) | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of the RAM in gigabytes | 
**VolumeType** | **string** | Volume type | 
**VolumeGb** | **int32** | Capacity of the volume in gigabytes | 
**SshKeyId** | Pointer to **int32** | SSH-key ID | [optional] 
**UserPassword** | Pointer to **string** | User password | [optional] 
**SecurityGroupId** | Pointer to **int32** | ID of the security group | [optional] 
**Price** | **float32** | Offer price of the spot instance | 

## Methods

### NewSpotCreate

`func NewSpotCreate(name string, dataCenterId string, osId int32, cloudNetworkType string, vCpuType string, vCpu int32, ramGb int32, volumeType string, volumeGb int32, price float32, ) *SpotCreate`

NewSpotCreate instantiates a new SpotCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotCreateWithDefaults

`func NewSpotCreateWithDefaults() *SpotCreate`

NewSpotCreateWithDefaults instantiates a new SpotCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpotCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpotCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpotCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDataCenterId

`func (o *SpotCreate) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *SpotCreate) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *SpotCreate) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetOsId

`func (o *SpotCreate) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *SpotCreate) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *SpotCreate) SetOsId(v int32)`

SetOsId sets OsId field to given value.


### GetCloudNetworkType

`func (o *SpotCreate) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *SpotCreate) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *SpotCreate) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.


### GetVCpuType

`func (o *SpotCreate) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *SpotCreate) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *SpotCreate) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetVCpu

`func (o *SpotCreate) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *SpotCreate) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *SpotCreate) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetRamGb

`func (o *SpotCreate) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *SpotCreate) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *SpotCreate) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeType

`func (o *SpotCreate) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *SpotCreate) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *SpotCreate) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetVolumeGb

`func (o *SpotCreate) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *SpotCreate) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *SpotCreate) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.


### GetSshKeyId

`func (o *SpotCreate) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *SpotCreate) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *SpotCreate) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *SpotCreate) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetUserPassword

`func (o *SpotCreate) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *SpotCreate) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *SpotCreate) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *SpotCreate) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *SpotCreate) GetSecurityGroupId() int32`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *SpotCreate) GetSecurityGroupIdOk() (*int32, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *SpotCreate) SetSecurityGroupId(v int32)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *SpotCreate) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetPrice

`func (o *SpotCreate) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpotCreate) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SpotCreate) SetPrice(v float32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


