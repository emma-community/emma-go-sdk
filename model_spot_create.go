/*
Public EMMA API

### About Infrastructure API  **Base URL:** **<u>https://api.emma.ms/external</u>**   This **Infrastructure API** is for managing the emma cloud infrastructure within a project. The API enables you to view, create, edit, and delete _Virtual machines, Spot instances, Applications, Kubernetes clusters, SSH keys, Security groups, and Backup policies_. For creating the resources you can use the endpoints with the dictionaries: _Data centers, Locations, Providers, Operating systems, Virtual machines configurations, Spot instances configurations, Kubernetes clusters configurations._   ### Authentication   #### 1. Create service application   To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.   - **Read** - only GET methods are allowed in the API.   - **Operate** - some operations are allowed with the resources (e.g. _Start, Reboot,_ and _Shutdown_ of the Virtual machines).   - **Manage** - full creating, updating, and deleting of the resources is allowed.    #### 2. Get access token   - Copy the **Client ID** and **Client Secret** in the service application.  - Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.   _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token \\  -H \"Content-Type: application/json\" \\  -d '{\"clientId\": \"YOUR-CLIENT-ID\", \"clientSecret\": \"YOUR-CLIENT-SECRET\"}'  ```  _For Windows:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token ^  -H \"Content-Type: application/json\" ^  -d \"{\\\"clientId\\\": \\\"YOUR-CLIENT-ID\\\", \\\"clientSecret\\\": \\\"YOUR-CLIENT-SECRET\\\"}\"  ```   #### 3. Use access token in requests  The Bearer access token is a text string, included in the request header, for example:   _For Linux / Mac:_  ```  curl -X GET https://api.emma.ms/external/v1/locations -H \"Authorization: Bearer YOUR-ACCESS-TOKEN-HERE\"  ```   Use this token for the API requests.    #### 4. Refresh token  The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).   To get a new access token send a request to the **_/refresh-token** endpoint:    _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/refresh-token \\  -H \"Content-Type: application/json\" \\  -d '{\"refreshToken\": \"YOUR-REFRESH-TOKEN\"}'  ```       ### Possible response status codes   We use standard HTTP response codes to show the success or failure of requests.   `2xx` - successful responses.   `4xx` - client error responses (the response contains an explanation of the error).   `5xx` - server error responses.   The API uses the following status codes:   | Status Code | Description                  | Notes                                                                  |  |-------------|------------------------------|------------------------------------------------------------------------|  | 200         | OK                           | The request was successful.                                             |  | 201         | Created                      | The object was successfully created. This code is only used with objects that are created immediately.  | 204         | No content                   | A successful request, but there is no additional information to send back in the response body (in a case when the object was deleted).    | 400         | Bad Request                  | The request could not be understood by the server. Incoming parameters might not be valid. |  | 401         | Unauthorized            | The client is unauthenticated. The client must authenticate itself to get the requested response. |  | 403         | Forbidden                   | The client does not have access rights to the content.  | 404         | Not Found                    | The requested resource is not found.                                    |  | 409         | Conflict | This response is sent when a request conflicts with the current state of the object (e.g. deleting the security group with the compute instances in it).|  | 422         | Unprocessable Content   | The request was well-formed but was unable to be followed due to incorrect field values (e.g. creation of a virtual machine in the non-existent data center).  |  | 500         | Internal server Error                 | The server could not return the representation due to an internal server error. | 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SpotCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotCreate{}

// SpotCreate struct for SpotCreate
type SpotCreate struct {
	// Name of the spot instance
	Name string `json:"name"`
	// ID of the provider's data center
	DataCenterId string `json:"dataCenterId"`
	// ID of the operating system
	OsId int32 `json:"osId"`
	// Type of the cloud network
	CloudNetworkType string `json:"cloudNetworkType"`
	// Type of the Central Processing Units (vCPUs)
	VCpuType string `json:"vCpuType"`
	// Number of virtual Central Processing Units (vCPUs)
	VCpu int32 `json:"vCpu"`
	// Capacity of the RAM in gigabytes
	RamGb int32 `json:"ramGb"`
	// Volume type
	VolumeType string `json:"volumeType"`
	// Capacity of the volume in gigabytes
	VolumeGb int32 `json:"volumeGb"`
	// SSH-key ID
	SshKeyId int32 `json:"sshKeyId"`
	// User password
	UserPassword *string `json:"userPassword,omitempty"`
	// ID of the security group
	SecurityGroupId *int32 `json:"securityGroupId,omitempty"`
	// Offer price of the spot instance
	Price float32 `json:"price"`
}

type _SpotCreate SpotCreate

// NewSpotCreate instantiates a new SpotCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotCreate(name string, dataCenterId string, osId int32, cloudNetworkType string, vCpuType string, vCpu int32, ramGb int32, volumeType string, volumeGb int32, sshKeyId int32, price float32) *SpotCreate {
	this := SpotCreate{}
	this.Name = name
	this.DataCenterId = dataCenterId
	this.OsId = osId
	this.CloudNetworkType = cloudNetworkType
	this.VCpuType = vCpuType
	this.VCpu = vCpu
	this.RamGb = ramGb
	this.VolumeType = volumeType
	this.VolumeGb = volumeGb
	this.SshKeyId = sshKeyId
	this.Price = price
	return &this
}

// NewSpotCreateWithDefaults instantiates a new SpotCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotCreateWithDefaults() *SpotCreate {
	this := SpotCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SpotCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpotCreate) SetName(v string) {
	o.Name = v
}

// GetDataCenterId returns the DataCenterId field value
func (o *SpotCreate) GetDataCenterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataCenterId
}

// GetDataCenterIdOk returns a tuple with the DataCenterId field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetDataCenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataCenterId, true
}

// SetDataCenterId sets field value
func (o *SpotCreate) SetDataCenterId(v string) {
	o.DataCenterId = v
}

// GetOsId returns the OsId field value
func (o *SpotCreate) GetOsId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OsId
}

// GetOsIdOk returns a tuple with the OsId field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetOsIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsId, true
}

// SetOsId sets field value
func (o *SpotCreate) SetOsId(v int32) {
	o.OsId = v
}

// GetCloudNetworkType returns the CloudNetworkType field value
func (o *SpotCreate) GetCloudNetworkType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudNetworkType
}

// GetCloudNetworkTypeOk returns a tuple with the CloudNetworkType field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetCloudNetworkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudNetworkType, true
}

// SetCloudNetworkType sets field value
func (o *SpotCreate) SetCloudNetworkType(v string) {
	o.CloudNetworkType = v
}

// GetVCpuType returns the VCpuType field value
func (o *SpotCreate) GetVCpuType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VCpuType
}

// GetVCpuTypeOk returns a tuple with the VCpuType field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetVCpuTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VCpuType, true
}

// SetVCpuType sets field value
func (o *SpotCreate) SetVCpuType(v string) {
	o.VCpuType = v
}

// GetVCpu returns the VCpu field value
func (o *SpotCreate) GetVCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VCpu
}

// GetVCpuOk returns a tuple with the VCpu field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetVCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VCpu, true
}

// SetVCpu sets field value
func (o *SpotCreate) SetVCpu(v int32) {
	o.VCpu = v
}

// GetRamGb returns the RamGb field value
func (o *SpotCreate) GetRamGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RamGb
}

// GetRamGbOk returns a tuple with the RamGb field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetRamGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamGb, true
}

// SetRamGb sets field value
func (o *SpotCreate) SetRamGb(v int32) {
	o.RamGb = v
}

// GetVolumeType returns the VolumeType field value
func (o *SpotCreate) GetVolumeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeType, true
}

// SetVolumeType sets field value
func (o *SpotCreate) SetVolumeType(v string) {
	o.VolumeType = v
}

// GetVolumeGb returns the VolumeGb field value
func (o *SpotCreate) GetVolumeGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VolumeGb
}

// GetVolumeGbOk returns a tuple with the VolumeGb field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetVolumeGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeGb, true
}

// SetVolumeGb sets field value
func (o *SpotCreate) SetVolumeGb(v int32) {
	o.VolumeGb = v
}

// GetSshKeyId returns the SshKeyId field value
func (o *SpotCreate) GetSshKeyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SshKeyId
}

// GetSshKeyIdOk returns a tuple with the SshKeyId field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetSshKeyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SshKeyId, true
}

// SetSshKeyId sets field value
func (o *SpotCreate) SetSshKeyId(v int32) {
	o.SshKeyId = v
}

// GetUserPassword returns the UserPassword field value if set, zero value otherwise.
func (o *SpotCreate) GetUserPassword() string {
	if o == nil || IsNil(o.UserPassword) {
		var ret string
		return ret
	}
	return *o.UserPassword
}

// GetUserPasswordOk returns a tuple with the UserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetUserPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.UserPassword) {
		return nil, false
	}
	return o.UserPassword, true
}

// HasUserPassword returns a boolean if a field has been set.
func (o *SpotCreate) HasUserPassword() bool {
	if o != nil && !IsNil(o.UserPassword) {
		return true
	}

	return false
}

// SetUserPassword gets a reference to the given string and assigns it to the UserPassword field.
func (o *SpotCreate) SetUserPassword(v string) {
	o.UserPassword = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *SpotCreate) GetSecurityGroupId() int32 {
	if o == nil || IsNil(o.SecurityGroupId) {
		var ret int32
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetSecurityGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SecurityGroupId) {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *SpotCreate) HasSecurityGroupId() bool {
	if o != nil && !IsNil(o.SecurityGroupId) {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given int32 and assigns it to the SecurityGroupId field.
func (o *SpotCreate) SetSecurityGroupId(v int32) {
	o.SecurityGroupId = &v
}

// GetPrice returns the Price field value
func (o *SpotCreate) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *SpotCreate) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *SpotCreate) SetPrice(v float32) {
	o.Price = v
}

func (o SpotCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["dataCenterId"] = o.DataCenterId
	toSerialize["osId"] = o.OsId
	toSerialize["cloudNetworkType"] = o.CloudNetworkType
	toSerialize["vCpuType"] = o.VCpuType
	toSerialize["vCpu"] = o.VCpu
	toSerialize["ramGb"] = o.RamGb
	toSerialize["volumeType"] = o.VolumeType
	toSerialize["volumeGb"] = o.VolumeGb
	toSerialize["sshKeyId"] = o.SshKeyId
	if !IsNil(o.UserPassword) {
		toSerialize["userPassword"] = o.UserPassword
	}
	if !IsNil(o.SecurityGroupId) {
		toSerialize["securityGroupId"] = o.SecurityGroupId
	}
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

func (o *SpotCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"dataCenterId",
		"osId",
		"cloudNetworkType",
		"vCpuType",
		"vCpu",
		"ramGb",
		"volumeType",
		"volumeGb",
		"sshKeyId",
		"price",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSpotCreate := _SpotCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpotCreate)

	if err != nil {
		return err
	}

	*o = SpotCreate(varSpotCreate)

	return err
}

type NullableSpotCreate struct {
	value *SpotCreate
	isSet bool
}

func (v NullableSpotCreate) Get() *SpotCreate {
	return v.value
}

func (v *NullableSpotCreate) Set(val *SpotCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotCreate(val *SpotCreate) *NullableSpotCreate {
	return &NullableSpotCreate{value: val, isSet: true}
}

func (v NullableSpotCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


