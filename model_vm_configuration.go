/*
Public EMMA API

**Base URL:** *<u>https://api.emma.ms/external</u>*  This **Infrastructure API** is for managing the cloud infrastructure within a project.  To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.  After creating the service application, copy the **Client ID** and **Client Secret**. Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.  The Bearer access token is a text string, included in the request header, example:  *-H Authorization: Bearer {token}*  Use this token for API requests. The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
)

// checks if the VmConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmConfiguration{}

// VmConfiguration struct for VmConfiguration
type VmConfiguration struct {
	Id                *int32               `json:"id,omitempty"`
	ProviderId        *int32               `json:"providerId,omitempty"`
	ProviderName      *string              `json:"providerName,omitempty"`
	LocationId        *int32               `json:"locationId,omitempty"`
	LocationName      *string              `json:"locationName,omitempty"`
	DataCenterId      *string              `json:"dataCenterId,omitempty"`
	DataCenterName    *string              `json:"dataCenterName,omitempty"`
	OsId              *int32               `json:"osId,omitempty"`
	OsType            *string              `json:"osType,omitempty"`
	OsVersion         *string              `json:"osVersion,omitempty"`
	CloudNetworkTypes []string             `json:"cloudNetworkTypes,omitempty"`
	VCpuType          *string              `json:"vCpuType,omitempty"`
	VCpu              *int32               `json:"vCpu,omitempty"`
	RamGb             *int32               `json:"ramGb,omitempty"`
	VolumeGb          *int32               `json:"volumeGb,omitempty"`
	VolumeType        *string              `json:"volumeType,omitempty"`
	Cost              *VmConfigurationCost `json:"cost,omitempty"`
}

// NewVmConfiguration instantiates a new VmConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmConfiguration() *VmConfiguration {
	this := VmConfiguration{}
	return &this
}

// NewVmConfigurationWithDefaults instantiates a new VmConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmConfigurationWithDefaults() *VmConfiguration {
	this := VmConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VmConfiguration) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VmConfiguration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VmConfiguration) SetId(v int32) {
	o.Id = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *VmConfiguration) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *VmConfiguration) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *VmConfiguration) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *VmConfiguration) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *VmConfiguration) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *VmConfiguration) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *VmConfiguration) GetLocationId() int32 {
	if o == nil || IsNil(o.LocationId) {
		var ret int32
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetLocationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LocationId) {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *VmConfiguration) HasLocationId() bool {
	if o != nil && !IsNil(o.LocationId) {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given int32 and assigns it to the LocationId field.
func (o *VmConfiguration) SetLocationId(v int32) {
	o.LocationId = &v
}

// GetLocationName returns the LocationName field value if set, zero value otherwise.
func (o *VmConfiguration) GetLocationName() string {
	if o == nil || IsNil(o.LocationName) {
		var ret string
		return ret
	}
	return *o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetLocationNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocationName) {
		return nil, false
	}
	return o.LocationName, true
}

// HasLocationName returns a boolean if a field has been set.
func (o *VmConfiguration) HasLocationName() bool {
	if o != nil && !IsNil(o.LocationName) {
		return true
	}

	return false
}

// SetLocationName gets a reference to the given string and assigns it to the LocationName field.
func (o *VmConfiguration) SetLocationName(v string) {
	o.LocationName = &v
}

// GetDataCenterId returns the DataCenterId field value if set, zero value otherwise.
func (o *VmConfiguration) GetDataCenterId() string {
	if o == nil || IsNil(o.DataCenterId) {
		var ret string
		return ret
	}
	return *o.DataCenterId
}

// GetDataCenterIdOk returns a tuple with the DataCenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetDataCenterIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataCenterId) {
		return nil, false
	}
	return o.DataCenterId, true
}

// HasDataCenterId returns a boolean if a field has been set.
func (o *VmConfiguration) HasDataCenterId() bool {
	if o != nil && !IsNil(o.DataCenterId) {
		return true
	}

	return false
}

// SetDataCenterId gets a reference to the given string and assigns it to the DataCenterId field.
func (o *VmConfiguration) SetDataCenterId(v string) {
	o.DataCenterId = &v
}

// GetDataCenterName returns the DataCenterName field value if set, zero value otherwise.
func (o *VmConfiguration) GetDataCenterName() string {
	if o == nil || IsNil(o.DataCenterName) {
		var ret string
		return ret
	}
	return *o.DataCenterName
}

// GetDataCenterNameOk returns a tuple with the DataCenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetDataCenterNameOk() (*string, bool) {
	if o == nil || IsNil(o.DataCenterName) {
		return nil, false
	}
	return o.DataCenterName, true
}

// HasDataCenterName returns a boolean if a field has been set.
func (o *VmConfiguration) HasDataCenterName() bool {
	if o != nil && !IsNil(o.DataCenterName) {
		return true
	}

	return false
}

// SetDataCenterName gets a reference to the given string and assigns it to the DataCenterName field.
func (o *VmConfiguration) SetDataCenterName(v string) {
	o.DataCenterName = &v
}

// GetOsId returns the OsId field value if set, zero value otherwise.
func (o *VmConfiguration) GetOsId() int32 {
	if o == nil || IsNil(o.OsId) {
		var ret int32
		return ret
	}
	return *o.OsId
}

// GetOsIdOk returns a tuple with the OsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetOsIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OsId) {
		return nil, false
	}
	return o.OsId, true
}

// HasOsId returns a boolean if a field has been set.
func (o *VmConfiguration) HasOsId() bool {
	if o != nil && !IsNil(o.OsId) {
		return true
	}

	return false
}

// SetOsId gets a reference to the given int32 and assigns it to the OsId field.
func (o *VmConfiguration) SetOsId(v int32) {
	o.OsId = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *VmConfiguration) GetOsType() string {
	if o == nil || IsNil(o.OsType) {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetOsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OsType) {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *VmConfiguration) HasOsType() bool {
	if o != nil && !IsNil(o.OsType) {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *VmConfiguration) SetOsType(v string) {
	o.OsType = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *VmConfiguration) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *VmConfiguration) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *VmConfiguration) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetCloudNetworkTypes returns the CloudNetworkTypes field value if set, zero value otherwise.
func (o *VmConfiguration) GetCloudNetworkTypes() []string {
	if o == nil || IsNil(o.CloudNetworkTypes) {
		var ret []string
		return ret
	}
	return o.CloudNetworkTypes
}

// GetCloudNetworkTypesOk returns a tuple with the CloudNetworkTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetCloudNetworkTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.CloudNetworkTypes) {
		return nil, false
	}
	return o.CloudNetworkTypes, true
}

// HasCloudNetworkTypes returns a boolean if a field has been set.
func (o *VmConfiguration) HasCloudNetworkTypes() bool {
	if o != nil && !IsNil(o.CloudNetworkTypes) {
		return true
	}

	return false
}

// SetCloudNetworkTypes gets a reference to the given []string and assigns it to the CloudNetworkTypes field.
func (o *VmConfiguration) SetCloudNetworkTypes(v []string) {
	o.CloudNetworkTypes = v
}

// GetVCpuType returns the VCpuType field value if set, zero value otherwise.
func (o *VmConfiguration) GetVCpuType() string {
	if o == nil || IsNil(o.VCpuType) {
		var ret string
		return ret
	}
	return *o.VCpuType
}

// GetVCpuTypeOk returns a tuple with the VCpuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetVCpuTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VCpuType) {
		return nil, false
	}
	return o.VCpuType, true
}

// HasVCpuType returns a boolean if a field has been set.
func (o *VmConfiguration) HasVCpuType() bool {
	if o != nil && !IsNil(o.VCpuType) {
		return true
	}

	return false
}

// SetVCpuType gets a reference to the given string and assigns it to the VCpuType field.
func (o *VmConfiguration) SetVCpuType(v string) {
	o.VCpuType = &v
}

// GetVCpu returns the VCpu field value if set, zero value otherwise.
func (o *VmConfiguration) GetVCpu() int32 {
	if o == nil || IsNil(o.VCpu) {
		var ret int32
		return ret
	}
	return *o.VCpu
}

// GetVCpuOk returns a tuple with the VCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetVCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.VCpu) {
		return nil, false
	}
	return o.VCpu, true
}

// HasVCpu returns a boolean if a field has been set.
func (o *VmConfiguration) HasVCpu() bool {
	if o != nil && !IsNil(o.VCpu) {
		return true
	}

	return false
}

// SetVCpu gets a reference to the given int32 and assigns it to the VCpu field.
func (o *VmConfiguration) SetVCpu(v int32) {
	o.VCpu = &v
}

// GetRamGb returns the RamGb field value if set, zero value otherwise.
func (o *VmConfiguration) GetRamGb() int32 {
	if o == nil || IsNil(o.RamGb) {
		var ret int32
		return ret
	}
	return *o.RamGb
}

// GetRamGbOk returns a tuple with the RamGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetRamGbOk() (*int32, bool) {
	if o == nil || IsNil(o.RamGb) {
		return nil, false
	}
	return o.RamGb, true
}

// HasRamGb returns a boolean if a field has been set.
func (o *VmConfiguration) HasRamGb() bool {
	if o != nil && !IsNil(o.RamGb) {
		return true
	}

	return false
}

// SetRamGb gets a reference to the given int32 and assigns it to the RamGb field.
func (o *VmConfiguration) SetRamGb(v int32) {
	o.RamGb = &v
}

// GetVolumeGb returns the VolumeGb field value if set, zero value otherwise.
func (o *VmConfiguration) GetVolumeGb() int32 {
	if o == nil || IsNil(o.VolumeGb) {
		var ret int32
		return ret
	}
	return *o.VolumeGb
}

// GetVolumeGbOk returns a tuple with the VolumeGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetVolumeGbOk() (*int32, bool) {
	if o == nil || IsNil(o.VolumeGb) {
		return nil, false
	}
	return o.VolumeGb, true
}

// HasVolumeGb returns a boolean if a field has been set.
func (o *VmConfiguration) HasVolumeGb() bool {
	if o != nil && !IsNil(o.VolumeGb) {
		return true
	}

	return false
}

// SetVolumeGb gets a reference to the given int32 and assigns it to the VolumeGb field.
func (o *VmConfiguration) SetVolumeGb(v int32) {
	o.VolumeGb = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *VmConfiguration) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType) {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeType) {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *VmConfiguration) HasVolumeType() bool {
	if o != nil && !IsNil(o.VolumeType) {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *VmConfiguration) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *VmConfiguration) GetCost() VmConfigurationCost {
	if o == nil || IsNil(o.Cost) {
		var ret VmConfigurationCost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfiguration) GetCostOk() (*VmConfigurationCost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *VmConfiguration) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given VmConfigurationCost and assigns it to the Cost field.
func (o *VmConfiguration) SetCost(v VmConfigurationCost) {
	o.Cost = &v
}

func (o VmConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.LocationId) {
		toSerialize["locationId"] = o.LocationId
	}
	if !IsNil(o.LocationName) {
		toSerialize["locationName"] = o.LocationName
	}
	if !IsNil(o.DataCenterId) {
		toSerialize["dataCenterId"] = o.DataCenterId
	}
	if !IsNil(o.DataCenterName) {
		toSerialize["dataCenterName"] = o.DataCenterName
	}
	if !IsNil(o.OsId) {
		toSerialize["osId"] = o.OsId
	}
	if !IsNil(o.OsType) {
		toSerialize["osType"] = o.OsType
	}
	if !IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !IsNil(o.CloudNetworkTypes) {
		toSerialize["cloudNetworkTypes"] = o.CloudNetworkTypes
	}
	if !IsNil(o.VCpuType) {
		toSerialize["vCpuType"] = o.VCpuType
	}
	if !IsNil(o.VCpu) {
		toSerialize["vCpu"] = o.VCpu
	}
	if !IsNil(o.RamGb) {
		toSerialize["ramGb"] = o.RamGb
	}
	if !IsNil(o.VolumeGb) {
		toSerialize["volumeGb"] = o.VolumeGb
	}
	if !IsNil(o.VolumeType) {
		toSerialize["volumeType"] = o.VolumeType
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	return toSerialize, nil
}

type NullableVmConfiguration struct {
	value *VmConfiguration
	isSet bool
}

func (v NullableVmConfiguration) Get() *VmConfiguration {
	return v.value
}

func (v *NullableVmConfiguration) Set(val *VmConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVmConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVmConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmConfiguration(val *VmConfiguration) *NullableVmConfiguration {
	return &NullableVmConfiguration{value: val, isSet: true}
}

func (v NullableVmConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
