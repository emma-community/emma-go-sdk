/*
Public EMMA API

### About Infrastructure API  **Base URL:** **<u>https://api.emma.ms/external</u>**   This **Infrastructure API** is for managing the emma cloud infrastructure within a project. The API enables you to view, create, edit, and delete _Virtual machines, Spot instances, Applications, Kubernetes clusters, SSH keys, Security groups, and Backup policies_. For creating the resources you can use the endpoints with the dictionaries: _Data centers, Locations, Providers, Operating systems, Virtual machines configurations, Spot instances configurations, Kubernetes clusters configurations._   ### Authentication   #### 1. Create service application   To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.   - **Read** - only GET methods are allowed in the API.   - **Operate** - some operations are allowed with the resources (e.g. _Start, Reboot,_ and _Shutdown_ of the Virtual machines).   - **Manage** - full creating, updating, and deleting of the resources is allowed.    #### 2. Get access token   - Copy the **Client ID** and **Client Secret** in the service application.  - Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.   _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token \\  -H \"Content-Type: application/json\" \\  -d '{\"clientId\": \"YOUR-CLIENT-ID\", \"clientSecret\": \"YOUR-CLIENT-SECRET\"}'  ```  _For Windows:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token ^  -H \"Content-Type: application/json\" ^  -d \"{\\\"clientId\\\": \\\"YOUR-CLIENT-ID\\\", \\\"clientSecret\\\": \\\"YOUR-CLIENT-SECRET\\\"}\"  ```   #### 3. Use access token in requests  The Bearer access token is a text string, included in the request header, for example:   _For Linux / Mac:_  ```  curl -X GET https://api.emma.ms/external/v1/locations -H \"Authorization: Bearer YOUR-ACCESS-TOKEN-HERE\"  ```   Use this token for the API requests.    #### 4. Refresh token  The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).   To get a new access token send a request to the **_/refresh-token** endpoint:    _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/refresh-token \\  -H \"Content-Type: application/json\" \\  -d '{\"refreshToken\": \"YOUR-REFRESH-TOKEN\"}'  ```       ### Possible response status codes   We use standard HTTP response codes to show the success or failure of requests.   `2xx` - successful responses.   `4xx` - client error responses (the response contains an explanation of the error).   `5xx` - server error responses.   The API uses the following status codes:   | Status Code | Description                  | Notes                                                                  |  |-------------|------------------------------|------------------------------------------------------------------------|  | 200         | OK                           | The request was successful.                                             |  | 201         | Created                      | The object was successfully created. This code is only used with objects that are created immediately.  | 204         | No content                   | A successful request, but there is no additional information to send back in the response body (in a case when the object was deleted).    | 400         | Bad Request                  | The request could not be understood by the server. Incoming parameters might not be valid. |  | 401         | Unauthorized            | The client is unauthenticated. The client must authenticate itself to get the requested response. |  | 403         | Forbidden                   | The client does not have access rights to the content.  | 404         | Not Found                    | The requested resource is not found.                                    |  | 409         | Conflict | This response is sent when a request conflicts with the current state of the object (e.g. deleting the security group with the compute instances in it).|  | 422         | Unprocessable Content   | The request was well-formed but was unable to be followed due to incorrect field values (e.g. creation of a virtual machine in the non-existent data center).  |  | 500         | Internal server Error                 | The server could not return the representation due to an internal server error. | 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
)

// checks if the KubernetesNodeGroupsInnerNodesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodeGroupsInnerNodesInner{}

// KubernetesNodeGroupsInnerNodesInner struct for KubernetesNodeGroupsInnerNodesInner
type KubernetesNodeGroupsInnerNodesInner struct {
	// ID of the worker node
	Id *int32 `json:"id,omitempty"`
	// Date and time when the worker node was created
	CreatedAt *string `json:"createdAt,omitempty"`
	// Name of the worker node
	Name *string `json:"name,omitempty"`
	// Status of the worker node
	Status *string `json:"status,omitempty"`
	Provider *KubernetesNodeGroupsInnerNodesInnerProvider `json:"provider,omitempty"`
	Location *KubernetesNodeGroupsInnerNodesInnerLocation `json:"location,omitempty"`
	DataCenter *KubernetesNodeGroupsInnerNodesInnerDataCenter `json:"dataCenter,omitempty"`
	Os *KubernetesNodeGroupsInnerNodesInnerOs `json:"os,omitempty"`
	// Number of virtual Central Processing Units (vCPUs)
	VCpu *int32 `json:"vCpu,omitempty"`
	// Type of the virtual Central Processing Unit
	VCpuType *string `json:"vCpuType,omitempty"`
	// Type of the cloud network (multi-cloud, isolated, or default)
	CloudNetworkType *string `json:"cloudNetworkType,omitempty"`
	// Capacity of the RAM in gigabytes
	RamGb *int32 `json:"ramGb,omitempty"`
	// Volumes attached
	Disks []KubernetesNodeGroupsInnerNodesInnerDisksInner `json:"disks,omitempty"`
	// Public and private networks
	Networks []KubernetesNodeGroupsInnerNodesInnerNetworksInner `json:"networks,omitempty"`
	SecurityGroup *KubernetesNodeGroupsInnerNodesInnerSecurityGroup `json:"securityGroup,omitempty"`
	Cost *KubernetesNodeGroupsInnerNodesInnerCost `json:"cost,omitempty"`
}

// NewKubernetesNodeGroupsInnerNodesInner instantiates a new KubernetesNodeGroupsInnerNodesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupsInnerNodesInner() *KubernetesNodeGroupsInnerNodesInner {
	this := KubernetesNodeGroupsInnerNodesInner{}
	return &this
}

// NewKubernetesNodeGroupsInnerNodesInnerWithDefaults instantiates a new KubernetesNodeGroupsInnerNodesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupsInnerNodesInnerWithDefaults() *KubernetesNodeGroupsInnerNodesInner {
	this := KubernetesNodeGroupsInnerNodesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetStatus(v string) {
	o.Status = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetProvider() KubernetesNodeGroupsInnerNodesInnerProvider {
	if o == nil || IsNil(o.Provider) {
		var ret KubernetesNodeGroupsInnerNodesInnerProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetProviderOk() (*KubernetesNodeGroupsInnerNodesInnerProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given KubernetesNodeGroupsInnerNodesInnerProvider and assigns it to the Provider field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetProvider(v KubernetesNodeGroupsInnerNodesInnerProvider) {
	o.Provider = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetLocation() KubernetesNodeGroupsInnerNodesInnerLocation {
	if o == nil || IsNil(o.Location) {
		var ret KubernetesNodeGroupsInnerNodesInnerLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetLocationOk() (*KubernetesNodeGroupsInnerNodesInnerLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given KubernetesNodeGroupsInnerNodesInnerLocation and assigns it to the Location field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetLocation(v KubernetesNodeGroupsInnerNodesInnerLocation) {
	o.Location = &v
}

// GetDataCenter returns the DataCenter field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetDataCenter() KubernetesNodeGroupsInnerNodesInnerDataCenter {
	if o == nil || IsNil(o.DataCenter) {
		var ret KubernetesNodeGroupsInnerNodesInnerDataCenter
		return ret
	}
	return *o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetDataCenterOk() (*KubernetesNodeGroupsInnerNodesInnerDataCenter, bool) {
	if o == nil || IsNil(o.DataCenter) {
		return nil, false
	}
	return o.DataCenter, true
}

// HasDataCenter returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasDataCenter() bool {
	if o != nil && !IsNil(o.DataCenter) {
		return true
	}

	return false
}

// SetDataCenter gets a reference to the given KubernetesNodeGroupsInnerNodesInnerDataCenter and assigns it to the DataCenter field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetDataCenter(v KubernetesNodeGroupsInnerNodesInnerDataCenter) {
	o.DataCenter = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetOs() KubernetesNodeGroupsInnerNodesInnerOs {
	if o == nil || IsNil(o.Os) {
		var ret KubernetesNodeGroupsInnerNodesInnerOs
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetOsOk() (*KubernetesNodeGroupsInnerNodesInnerOs, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given KubernetesNodeGroupsInnerNodesInnerOs and assigns it to the Os field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetOs(v KubernetesNodeGroupsInnerNodesInnerOs) {
	o.Os = &v
}

// GetVCpu returns the VCpu field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpu() int32 {
	if o == nil || IsNil(o.VCpu) {
		var ret int32
		return ret
	}
	return *o.VCpu
}

// GetVCpuOk returns a tuple with the VCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.VCpu) {
		return nil, false
	}
	return o.VCpu, true
}

// HasVCpu returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasVCpu() bool {
	if o != nil && !IsNil(o.VCpu) {
		return true
	}

	return false
}

// SetVCpu gets a reference to the given int32 and assigns it to the VCpu field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetVCpu(v int32) {
	o.VCpu = &v
}

// GetVCpuType returns the VCpuType field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpuType() string {
	if o == nil || IsNil(o.VCpuType) {
		var ret string
		return ret
	}
	return *o.VCpuType
}

// GetVCpuTypeOk returns a tuple with the VCpuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetVCpuTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VCpuType) {
		return nil, false
	}
	return o.VCpuType, true
}

// HasVCpuType returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasVCpuType() bool {
	if o != nil && !IsNil(o.VCpuType) {
		return true
	}

	return false
}

// SetVCpuType gets a reference to the given string and assigns it to the VCpuType field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetVCpuType(v string) {
	o.VCpuType = &v
}

// GetCloudNetworkType returns the CloudNetworkType field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetCloudNetworkType() string {
	if o == nil || IsNil(o.CloudNetworkType) {
		var ret string
		return ret
	}
	return *o.CloudNetworkType
}

// GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetCloudNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CloudNetworkType) {
		return nil, false
	}
	return o.CloudNetworkType, true
}

// HasCloudNetworkType returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasCloudNetworkType() bool {
	if o != nil && !IsNil(o.CloudNetworkType) {
		return true
	}

	return false
}

// SetCloudNetworkType gets a reference to the given string and assigns it to the CloudNetworkType field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetCloudNetworkType(v string) {
	o.CloudNetworkType = &v
}

// GetRamGb returns the RamGb field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetRamGb() int32 {
	if o == nil || IsNil(o.RamGb) {
		var ret int32
		return ret
	}
	return *o.RamGb
}

// GetRamGbOk returns a tuple with the RamGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetRamGbOk() (*int32, bool) {
	if o == nil || IsNil(o.RamGb) {
		return nil, false
	}
	return o.RamGb, true
}

// HasRamGb returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasRamGb() bool {
	if o != nil && !IsNil(o.RamGb) {
		return true
	}

	return false
}

// SetRamGb gets a reference to the given int32 and assigns it to the RamGb field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetRamGb(v int32) {
	o.RamGb = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetDisks() []KubernetesNodeGroupsInnerNodesInnerDisksInner {
	if o == nil || IsNil(o.Disks) {
		var ret []KubernetesNodeGroupsInnerNodesInnerDisksInner
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetDisksOk() ([]KubernetesNodeGroupsInnerNodesInnerDisksInner, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasDisks() bool {
	if o != nil && !IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []KubernetesNodeGroupsInnerNodesInnerDisksInner and assigns it to the Disks field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetDisks(v []KubernetesNodeGroupsInnerNodesInnerDisksInner) {
	o.Disks = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetNetworks() []KubernetesNodeGroupsInnerNodesInnerNetworksInner {
	if o == nil || IsNil(o.Networks) {
		var ret []KubernetesNodeGroupsInnerNodesInnerNetworksInner
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetNetworksOk() ([]KubernetesNodeGroupsInnerNodesInnerNetworksInner, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []KubernetesNodeGroupsInnerNodesInnerNetworksInner and assigns it to the Networks field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetNetworks(v []KubernetesNodeGroupsInnerNodesInnerNetworksInner) {
	o.Networks = v
}

// GetSecurityGroup returns the SecurityGroup field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetSecurityGroup() KubernetesNodeGroupsInnerNodesInnerSecurityGroup {
	if o == nil || IsNil(o.SecurityGroup) {
		var ret KubernetesNodeGroupsInnerNodesInnerSecurityGroup
		return ret
	}
	return *o.SecurityGroup
}

// GetSecurityGroupOk returns a tuple with the SecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetSecurityGroupOk() (*KubernetesNodeGroupsInnerNodesInnerSecurityGroup, bool) {
	if o == nil || IsNil(o.SecurityGroup) {
		return nil, false
	}
	return o.SecurityGroup, true
}

// HasSecurityGroup returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasSecurityGroup() bool {
	if o != nil && !IsNil(o.SecurityGroup) {
		return true
	}

	return false
}

// SetSecurityGroup gets a reference to the given KubernetesNodeGroupsInnerNodesInnerSecurityGroup and assigns it to the SecurityGroup field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetSecurityGroup(v KubernetesNodeGroupsInnerNodesInnerSecurityGroup) {
	o.SecurityGroup = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInner) GetCost() KubernetesNodeGroupsInnerNodesInnerCost {
	if o == nil || IsNil(o.Cost) {
		var ret KubernetesNodeGroupsInnerNodesInnerCost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) GetCostOk() (*KubernetesNodeGroupsInnerNodesInnerCost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInner) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given KubernetesNodeGroupsInnerNodesInnerCost and assigns it to the Cost field.
func (o *KubernetesNodeGroupsInnerNodesInner) SetCost(v KubernetesNodeGroupsInnerNodesInnerCost) {
	o.Cost = &v
}

func (o KubernetesNodeGroupsInnerNodesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNodeGroupsInnerNodesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	return toSerialize, nil
}

type NullableKubernetesNodeGroupsInnerNodesInner struct {
	value *KubernetesNodeGroupsInnerNodesInner
	isSet bool
}

func (v NullableKubernetesNodeGroupsInnerNodesInner) Get() *KubernetesNodeGroupsInnerNodesInner {
	return v.value
}

func (v *NullableKubernetesNodeGroupsInnerNodesInner) Set(val *KubernetesNodeGroupsInnerNodesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupsInnerNodesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupsInnerNodesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupsInnerNodesInner(val *KubernetesNodeGroupsInnerNodesInner) *NullableKubernetesNodeGroupsInnerNodesInner {
	return &NullableKubernetesNodeGroupsInnerNodesInner{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupsInnerNodesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupsInnerNodesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


