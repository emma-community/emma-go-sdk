# VmCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Virtual machine name | 
**DataCenterId** | **string** | Provider&#39;s data center ID | 
**OsId** | **int32** | Operating system ID | 
**CloudNetworkType** | **string** | Cloud network type | 
**VCpuType** | **string** | vCPU type | 
**VCpu** | **int32** | Number of virtual Central Processing Units (vCPUs) | 
**RamGb** | **int32** | Capacity of RAM in gigabytes | 
**VolumeType** | **string** | Volume type | 
**VolumeGb** | **int32** | Capacity of volume in gigabytes | 
**SshKeyId** | **int32** | SSH-key ID | 
**SecurityGroupId** | Pointer to **int32** | Security group ID | [optional] 

## Methods

### NewVmCreate

`func NewVmCreate(name string, dataCenterId string, osId int32, cloudNetworkType string, vCpuType string, vCpu int32, ramGb int32, volumeType string, volumeGb int32, sshKeyId int32, ) *VmCreate`

NewVmCreate instantiates a new VmCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmCreateWithDefaults

`func NewVmCreateWithDefaults() *VmCreate`

NewVmCreateWithDefaults instantiates a new VmCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDataCenterId

`func (o *VmCreate) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VmCreate) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VmCreate) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetOsId

`func (o *VmCreate) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *VmCreate) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *VmCreate) SetOsId(v int32)`

SetOsId sets OsId field to given value.


### GetCloudNetworkType

`func (o *VmCreate) GetCloudNetworkType() string`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *VmCreate) GetCloudNetworkTypeOk() (*string, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *VmCreate) SetCloudNetworkType(v string)`

SetCloudNetworkType sets CloudNetworkType field to given value.


### GetVCpuType

`func (o *VmCreate) GetVCpuType() string`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmCreate) GetVCpuTypeOk() (*string, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmCreate) SetVCpuType(v string)`

SetVCpuType sets VCpuType field to given value.


### GetVCpu

`func (o *VmCreate) GetVCpu() int32`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmCreate) GetVCpuOk() (*int32, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmCreate) SetVCpu(v int32)`

SetVCpu sets VCpu field to given value.


### GetRamGb

`func (o *VmCreate) GetRamGb() int32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmCreate) GetRamGbOk() (*int32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmCreate) SetRamGb(v int32)`

SetRamGb sets RamGb field to given value.


### GetVolumeType

`func (o *VmCreate) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VmCreate) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VmCreate) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetVolumeGb

`func (o *VmCreate) GetVolumeGb() int32`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VmCreate) GetVolumeGbOk() (*int32, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VmCreate) SetVolumeGb(v int32)`

SetVolumeGb sets VolumeGb field to given value.


### GetSshKeyId

`func (o *VmCreate) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *VmCreate) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *VmCreate) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.


### GetSecurityGroupId

`func (o *VmCreate) GetSecurityGroupId() int32`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *VmCreate) GetSecurityGroupIdOk() (*int32, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *VmCreate) SetSecurityGroupId(v int32)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *VmCreate) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


