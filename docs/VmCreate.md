# VmCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | Virtual machine name | 
**DataCenterId** | **interface{}** | Provider&#39;s data center ID | 
**OsId** | **interface{}** | Operating system ID | 
**CloudNetworkType** | **interface{}** | Cloud network type | 
**VCpuType** | **interface{}** | vCPU type | 
**VCpu** | **interface{}** | Number of virtual Central Processing Units (vCPUs) | 
**RamGb** | **interface{}** | Capacity of RAM in gigabytes | 
**VolumeType** | **interface{}** | Volume type | 
**VolumeGb** | **interface{}** | Capacity of volume in gigabytes | 
**SshKeyId** | **interface{}** | SSH-key ID | 
**SecurityGroupId** | Pointer to **interface{}** | Security group ID | [optional] 

## Methods

### NewVmCreate

`func NewVmCreate(name interface{}, dataCenterId interface{}, osId interface{}, cloudNetworkType interface{}, vCpuType interface{}, vCpu interface{}, ramGb interface{}, volumeType interface{}, volumeGb interface{}, sshKeyId interface{}, ) *VmCreate`

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

`func (o *VmCreate) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmCreate) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmCreate) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *VmCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDataCenterId

`func (o *VmCreate) GetDataCenterId() interface{}`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VmCreate) GetDataCenterIdOk() (*interface{}, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VmCreate) SetDataCenterId(v interface{})`

SetDataCenterId sets DataCenterId field to given value.


### SetDataCenterIdNil

`func (o *VmCreate) SetDataCenterIdNil(b bool)`

 SetDataCenterIdNil sets the value for DataCenterId to be an explicit nil

### UnsetDataCenterId
`func (o *VmCreate) UnsetDataCenterId()`

UnsetDataCenterId ensures that no value is present for DataCenterId, not even an explicit nil
### GetOsId

`func (o *VmCreate) GetOsId() interface{}`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *VmCreate) GetOsIdOk() (*interface{}, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *VmCreate) SetOsId(v interface{})`

SetOsId sets OsId field to given value.


### SetOsIdNil

`func (o *VmCreate) SetOsIdNil(b bool)`

 SetOsIdNil sets the value for OsId to be an explicit nil

### UnsetOsId
`func (o *VmCreate) UnsetOsId()`

UnsetOsId ensures that no value is present for OsId, not even an explicit nil
### GetCloudNetworkType

`func (o *VmCreate) GetCloudNetworkType() interface{}`

GetCloudNetworkType returns the CloudNetworkType field if non-nil, zero value otherwise.

### GetCloudNetworkTypeOk

`func (o *VmCreate) GetCloudNetworkTypeOk() (*interface{}, bool)`

GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkType

`func (o *VmCreate) SetCloudNetworkType(v interface{})`

SetCloudNetworkType sets CloudNetworkType field to given value.


### SetCloudNetworkTypeNil

`func (o *VmCreate) SetCloudNetworkTypeNil(b bool)`

 SetCloudNetworkTypeNil sets the value for CloudNetworkType to be an explicit nil

### UnsetCloudNetworkType
`func (o *VmCreate) UnsetCloudNetworkType()`

UnsetCloudNetworkType ensures that no value is present for CloudNetworkType, not even an explicit nil
### GetVCpuType

`func (o *VmCreate) GetVCpuType() interface{}`

GetVCpuType returns the VCpuType field if non-nil, zero value otherwise.

### GetVCpuTypeOk

`func (o *VmCreate) GetVCpuTypeOk() (*interface{}, bool)`

GetVCpuTypeOk returns a tuple with the VCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpuType

`func (o *VmCreate) SetVCpuType(v interface{})`

SetVCpuType sets VCpuType field to given value.


### SetVCpuTypeNil

`func (o *VmCreate) SetVCpuTypeNil(b bool)`

 SetVCpuTypeNil sets the value for VCpuType to be an explicit nil

### UnsetVCpuType
`func (o *VmCreate) UnsetVCpuType()`

UnsetVCpuType ensures that no value is present for VCpuType, not even an explicit nil
### GetVCpu

`func (o *VmCreate) GetVCpu() interface{}`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VmCreate) GetVCpuOk() (*interface{}, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VmCreate) SetVCpu(v interface{})`

SetVCpu sets VCpu field to given value.


### SetVCpuNil

`func (o *VmCreate) SetVCpuNil(b bool)`

 SetVCpuNil sets the value for VCpu to be an explicit nil

### UnsetVCpu
`func (o *VmCreate) UnsetVCpu()`

UnsetVCpu ensures that no value is present for VCpu, not even an explicit nil
### GetRamGb

`func (o *VmCreate) GetRamGb() interface{}`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmCreate) GetRamGbOk() (*interface{}, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmCreate) SetRamGb(v interface{})`

SetRamGb sets RamGb field to given value.


### SetRamGbNil

`func (o *VmCreate) SetRamGbNil(b bool)`

 SetRamGbNil sets the value for RamGb to be an explicit nil

### UnsetRamGb
`func (o *VmCreate) UnsetRamGb()`

UnsetRamGb ensures that no value is present for RamGb, not even an explicit nil
### GetVolumeType

`func (o *VmCreate) GetVolumeType() interface{}`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VmCreate) GetVolumeTypeOk() (*interface{}, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VmCreate) SetVolumeType(v interface{})`

SetVolumeType sets VolumeType field to given value.


### SetVolumeTypeNil

`func (o *VmCreate) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *VmCreate) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetVolumeGb

`func (o *VmCreate) GetVolumeGb() interface{}`

GetVolumeGb returns the VolumeGb field if non-nil, zero value otherwise.

### GetVolumeGbOk

`func (o *VmCreate) GetVolumeGbOk() (*interface{}, bool)`

GetVolumeGbOk returns a tuple with the VolumeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGb

`func (o *VmCreate) SetVolumeGb(v interface{})`

SetVolumeGb sets VolumeGb field to given value.


### SetVolumeGbNil

`func (o *VmCreate) SetVolumeGbNil(b bool)`

 SetVolumeGbNil sets the value for VolumeGb to be an explicit nil

### UnsetVolumeGb
`func (o *VmCreate) UnsetVolumeGb()`

UnsetVolumeGb ensures that no value is present for VolumeGb, not even an explicit nil
### GetSshKeyId

`func (o *VmCreate) GetSshKeyId() interface{}`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *VmCreate) GetSshKeyIdOk() (*interface{}, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *VmCreate) SetSshKeyId(v interface{})`

SetSshKeyId sets SshKeyId field to given value.


### SetSshKeyIdNil

`func (o *VmCreate) SetSshKeyIdNil(b bool)`

 SetSshKeyIdNil sets the value for SshKeyId to be an explicit nil

### UnsetSshKeyId
`func (o *VmCreate) UnsetSshKeyId()`

UnsetSshKeyId ensures that no value is present for SshKeyId, not even an explicit nil
### GetSecurityGroupId

`func (o *VmCreate) GetSecurityGroupId() interface{}`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *VmCreate) GetSecurityGroupIdOk() (*interface{}, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *VmCreate) SetSecurityGroupId(v interface{})`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *VmCreate) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *VmCreate) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *VmCreate) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


