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

// checks if the SecurityGroupInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupInstance{}

// SecurityGroupInstance struct for SecurityGroupInstance
type SecurityGroupInstance struct {
	Id             *int32                           `json:"id,omitempty"`
	CreatedAt      *string                          `json:"createdAt,omitempty"`
	CreatedByName  *string                          `json:"createdByName,omitempty"`
	CreatedById    *int32                           `json:"createdById,omitempty"`
	ModifiedAt     *string                          `json:"modifiedAt,omitempty"`
	ModifiedByName *string                          `json:"modifiedByName,omitempty"`
	ModifiedById   *int32                           `json:"modifiedById,omitempty"`
	Name           *string                          `json:"name,omitempty"`
	ProjectId      *int32                           `json:"projectId,omitempty"`
	Status         *string                          `json:"status,omitempty"`
	Provider       *SecurityGroupInstanceProvider   `json:"provider,omitempty"`
	Location       *SecurityGroupInstanceLocation   `json:"location,omitempty"`
	DataCenter     *SecurityGroupInstanceDataCenter `json:"dataCenter,omitempty"`
	Os             *SecurityGroupInstanceOs         `json:"os,omitempty"`
	VCpu           *int32                           `json:"vCpu,omitempty"`
	// vCPU type
	VCpuType *string `json:"vCpuType,omitempty"`
	// Cloud network type
	CloudNetworkType *string                              `json:"cloudNetworkType,omitempty"`
	RamGb            *int32                               `json:"ramGb,omitempty"`
	Disks            []SecurityGroupInstanceDisksInner    `json:"disks,omitempty"`
	Networks         []SecurityGroupInstanceNetworksInner `json:"networks,omitempty"`
	SecurityGroup    *SecurityGroupInstanceSecurityGroup  `json:"securityGroup,omitempty"`
	SshKeyId         *int32                               `json:"sshKeyId,omitempty"`
	UserName         *string                              `json:"userName,omitempty"`
	Cost             *SecurityGroupInstanceCost           `json:"cost,omitempty"`
}

// NewSecurityGroupInstance instantiates a new SecurityGroupInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupInstance() *SecurityGroupInstance {
	this := SecurityGroupInstance{}
	return &this
}

// NewSecurityGroupInstanceWithDefaults instantiates a new SecurityGroupInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupInstanceWithDefaults() *SecurityGroupInstance {
	this := SecurityGroupInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SecurityGroupInstance) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SecurityGroupInstance) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetCreatedByName() string {
	if o == nil || IsNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetCreatedByNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByName) {
		return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasCreatedByName() bool {
	if o != nil && !IsNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *SecurityGroupInstance) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetCreatedById() int32 {
	if o == nil || IsNil(o.CreatedById) {
		var ret int32
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetCreatedByIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given int32 and assigns it to the CreatedById field.
func (o *SecurityGroupInstance) SetCreatedById(v int32) {
	o.CreatedById = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetModifiedAt() string {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetModifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *SecurityGroupInstance) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetModifiedByName returns the ModifiedByName field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetModifiedByName() string {
	if o == nil || IsNil(o.ModifiedByName) {
		var ret string
		return ret
	}
	return *o.ModifiedByName
}

// GetModifiedByNameOk returns a tuple with the ModifiedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetModifiedByNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedByName) {
		return nil, false
	}
	return o.ModifiedByName, true
}

// HasModifiedByName returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasModifiedByName() bool {
	if o != nil && !IsNil(o.ModifiedByName) {
		return true
	}

	return false
}

// SetModifiedByName gets a reference to the given string and assigns it to the ModifiedByName field.
func (o *SecurityGroupInstance) SetModifiedByName(v string) {
	o.ModifiedByName = &v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetModifiedById() int32 {
	if o == nil || IsNil(o.ModifiedById) {
		var ret int32
		return ret
	}
	return *o.ModifiedById
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetModifiedByIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ModifiedById) {
		return nil, false
	}
	return o.ModifiedById, true
}

// HasModifiedById returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasModifiedById() bool {
	if o != nil && !IsNil(o.ModifiedById) {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given int32 and assigns it to the ModifiedById field.
func (o *SecurityGroupInstance) SetModifiedById(v int32) {
	o.ModifiedById = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityGroupInstance) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *SecurityGroupInstance) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecurityGroupInstance) SetStatus(v string) {
	o.Status = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetProvider() SecurityGroupInstanceProvider {
	if o == nil || IsNil(o.Provider) {
		var ret SecurityGroupInstanceProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetProviderOk() (*SecurityGroupInstanceProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given SecurityGroupInstanceProvider and assigns it to the Provider field.
func (o *SecurityGroupInstance) SetProvider(v SecurityGroupInstanceProvider) {
	o.Provider = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetLocation() SecurityGroupInstanceLocation {
	if o == nil || IsNil(o.Location) {
		var ret SecurityGroupInstanceLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetLocationOk() (*SecurityGroupInstanceLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SecurityGroupInstanceLocation and assigns it to the Location field.
func (o *SecurityGroupInstance) SetLocation(v SecurityGroupInstanceLocation) {
	o.Location = &v
}

// GetDataCenter returns the DataCenter field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetDataCenter() SecurityGroupInstanceDataCenter {
	if o == nil || IsNil(o.DataCenter) {
		var ret SecurityGroupInstanceDataCenter
		return ret
	}
	return *o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetDataCenterOk() (*SecurityGroupInstanceDataCenter, bool) {
	if o == nil || IsNil(o.DataCenter) {
		return nil, false
	}
	return o.DataCenter, true
}

// HasDataCenter returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasDataCenter() bool {
	if o != nil && !IsNil(o.DataCenter) {
		return true
	}

	return false
}

// SetDataCenter gets a reference to the given SecurityGroupInstanceDataCenter and assigns it to the DataCenter field.
func (o *SecurityGroupInstance) SetDataCenter(v SecurityGroupInstanceDataCenter) {
	o.DataCenter = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetOs() SecurityGroupInstanceOs {
	if o == nil || IsNil(o.Os) {
		var ret SecurityGroupInstanceOs
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetOsOk() (*SecurityGroupInstanceOs, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given SecurityGroupInstanceOs and assigns it to the Os field.
func (o *SecurityGroupInstance) SetOs(v SecurityGroupInstanceOs) {
	o.Os = &v
}

// GetVCpu returns the VCpu field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetVCpu() int32 {
	if o == nil || IsNil(o.VCpu) {
		var ret int32
		return ret
	}
	return *o.VCpu
}

// GetVCpuOk returns a tuple with the VCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetVCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.VCpu) {
		return nil, false
	}
	return o.VCpu, true
}

// HasVCpu returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasVCpu() bool {
	if o != nil && !IsNil(o.VCpu) {
		return true
	}

	return false
}

// SetVCpu gets a reference to the given int32 and assigns it to the VCpu field.
func (o *SecurityGroupInstance) SetVCpu(v int32) {
	o.VCpu = &v
}

// GetVCpuType returns the VCpuType field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetVCpuType() string {
	if o == nil || IsNil(o.VCpuType) {
		var ret string
		return ret
	}
	return *o.VCpuType
}

// GetVCpuTypeOk returns a tuple with the VCpuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetVCpuTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VCpuType) {
		return nil, false
	}
	return o.VCpuType, true
}

// HasVCpuType returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasVCpuType() bool {
	if o != nil && !IsNil(o.VCpuType) {
		return true
	}

	return false
}

// SetVCpuType gets a reference to the given string and assigns it to the VCpuType field.
func (o *SecurityGroupInstance) SetVCpuType(v string) {
	o.VCpuType = &v
}

// GetCloudNetworkType returns the CloudNetworkType field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetCloudNetworkType() string {
	if o == nil || IsNil(o.CloudNetworkType) {
		var ret string
		return ret
	}
	return *o.CloudNetworkType
}

// GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetCloudNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CloudNetworkType) {
		return nil, false
	}
	return o.CloudNetworkType, true
}

// HasCloudNetworkType returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasCloudNetworkType() bool {
	if o != nil && !IsNil(o.CloudNetworkType) {
		return true
	}

	return false
}

// SetCloudNetworkType gets a reference to the given string and assigns it to the CloudNetworkType field.
func (o *SecurityGroupInstance) SetCloudNetworkType(v string) {
	o.CloudNetworkType = &v
}

// GetRamGb returns the RamGb field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetRamGb() int32 {
	if o == nil || IsNil(o.RamGb) {
		var ret int32
		return ret
	}
	return *o.RamGb
}

// GetRamGbOk returns a tuple with the RamGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetRamGbOk() (*int32, bool) {
	if o == nil || IsNil(o.RamGb) {
		return nil, false
	}
	return o.RamGb, true
}

// HasRamGb returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasRamGb() bool {
	if o != nil && !IsNil(o.RamGb) {
		return true
	}

	return false
}

// SetRamGb gets a reference to the given int32 and assigns it to the RamGb field.
func (o *SecurityGroupInstance) SetRamGb(v int32) {
	o.RamGb = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetDisks() []SecurityGroupInstanceDisksInner {
	if o == nil || IsNil(o.Disks) {
		var ret []SecurityGroupInstanceDisksInner
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetDisksOk() ([]SecurityGroupInstanceDisksInner, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasDisks() bool {
	if o != nil && !IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []SecurityGroupInstanceDisksInner and assigns it to the Disks field.
func (o *SecurityGroupInstance) SetDisks(v []SecurityGroupInstanceDisksInner) {
	o.Disks = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetNetworks() []SecurityGroupInstanceNetworksInner {
	if o == nil || IsNil(o.Networks) {
		var ret []SecurityGroupInstanceNetworksInner
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetNetworksOk() ([]SecurityGroupInstanceNetworksInner, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []SecurityGroupInstanceNetworksInner and assigns it to the Networks field.
func (o *SecurityGroupInstance) SetNetworks(v []SecurityGroupInstanceNetworksInner) {
	o.Networks = v
}

// GetSecurityGroup returns the SecurityGroup field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetSecurityGroup() SecurityGroupInstanceSecurityGroup {
	if o == nil || IsNil(o.SecurityGroup) {
		var ret SecurityGroupInstanceSecurityGroup
		return ret
	}
	return *o.SecurityGroup
}

// GetSecurityGroupOk returns a tuple with the SecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetSecurityGroupOk() (*SecurityGroupInstanceSecurityGroup, bool) {
	if o == nil || IsNil(o.SecurityGroup) {
		return nil, false
	}
	return o.SecurityGroup, true
}

// HasSecurityGroup returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasSecurityGroup() bool {
	if o != nil && !IsNil(o.SecurityGroup) {
		return true
	}

	return false
}

// SetSecurityGroup gets a reference to the given SecurityGroupInstanceSecurityGroup and assigns it to the SecurityGroup field.
func (o *SecurityGroupInstance) SetSecurityGroup(v SecurityGroupInstanceSecurityGroup) {
	o.SecurityGroup = &v
}

// GetSshKeyId returns the SshKeyId field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetSshKeyId() int32 {
	if o == nil || IsNil(o.SshKeyId) {
		var ret int32
		return ret
	}
	return *o.SshKeyId
}

// GetSshKeyIdOk returns a tuple with the SshKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetSshKeyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SshKeyId) {
		return nil, false
	}
	return o.SshKeyId, true
}

// HasSshKeyId returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasSshKeyId() bool {
	if o != nil && !IsNil(o.SshKeyId) {
		return true
	}

	return false
}

// SetSshKeyId gets a reference to the given int32 and assigns it to the SshKeyId field.
func (o *SecurityGroupInstance) SetSshKeyId(v int32) {
	o.SshKeyId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SecurityGroupInstance) SetUserName(v string) {
	o.UserName = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *SecurityGroupInstance) GetCost() SecurityGroupInstanceCost {
	if o == nil || IsNil(o.Cost) {
		var ret SecurityGroupInstanceCost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupInstance) GetCostOk() (*SecurityGroupInstanceCost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *SecurityGroupInstance) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given SecurityGroupInstanceCost and assigns it to the Cost field.
func (o *SecurityGroupInstance) SetCost(v SecurityGroupInstanceCost) {
	o.Cost = &v
}

func (o SecurityGroupInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedByName) {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !IsNil(o.ModifiedByName) {
		toSerialize["modifiedByName"] = o.ModifiedByName
	}
	if !IsNil(o.ModifiedById) {
		toSerialize["modifiedById"] = o.ModifiedById
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.DataCenter) {
		toSerialize["dataCenter"] = o.DataCenter
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.VCpu) {
		toSerialize["vCpu"] = o.VCpu
	}
	if !IsNil(o.VCpuType) {
		toSerialize["vCpuType"] = o.VCpuType
	}
	if !IsNil(o.CloudNetworkType) {
		toSerialize["cloudNetworkType"] = o.CloudNetworkType
	}
	if !IsNil(o.RamGb) {
		toSerialize["ramGb"] = o.RamGb
	}
	if !IsNil(o.Disks) {
		toSerialize["disks"] = o.Disks
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.SecurityGroup) {
		toSerialize["securityGroup"] = o.SecurityGroup
	}
	if !IsNil(o.SshKeyId) {
		toSerialize["sshKeyId"] = o.SshKeyId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	return toSerialize, nil
}

type NullableSecurityGroupInstance struct {
	value *SecurityGroupInstance
	isSet bool
}

func (v NullableSecurityGroupInstance) Get() *SecurityGroupInstance {
	return v.value
}

func (v *NullableSecurityGroupInstance) Set(val *SecurityGroupInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupInstance(val *SecurityGroupInstance) *NullableSecurityGroupInstance {
	return &NullableSecurityGroupInstance{value: val, isSet: true}
}

func (v NullableSecurityGroupInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
