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

// checks if the KubernetesNodeGroupsInnerNodesInnerDisksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodeGroupsInnerNodesInnerDisksInner{}

// KubernetesNodeGroupsInnerNodesInnerDisksInner struct for KubernetesNodeGroupsInnerNodesInnerDisksInner
type KubernetesNodeGroupsInnerNodesInnerDisksInner struct {
	// Volume ID
	Id *int32 `json:"id,omitempty"`
	// Volume size in gigabytes
	SizeGb *int32 `json:"sizeGb,omitempty"`
	// ID of the volume type
	TypeId *int32 `json:"typeId,omitempty"`
	// Volume type
	Type *string `json:"type,omitempty"`
	// Indicates whether the volume is bootable or not
	IsBootable *bool `json:"isBootable,omitempty"`
}

// NewKubernetesNodeGroupsInnerNodesInnerDisksInner instantiates a new KubernetesNodeGroupsInnerNodesInnerDisksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupsInnerNodesInnerDisksInner() *KubernetesNodeGroupsInnerNodesInnerDisksInner {
	this := KubernetesNodeGroupsInnerNodesInnerDisksInner{}
	return &this
}

// NewKubernetesNodeGroupsInnerNodesInnerDisksInnerWithDefaults instantiates a new KubernetesNodeGroupsInnerNodesInnerDisksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupsInnerNodesInnerDisksInnerWithDefaults() *KubernetesNodeGroupsInnerNodesInnerDisksInner {
	this := KubernetesNodeGroupsInnerNodesInnerDisksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) SetId(v int32) {
	o.Id = &v
}

// GetSizeGb returns the SizeGb field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetSizeGb() int32 {
	if o == nil || IsNil(o.SizeGb) {
		var ret int32
		return ret
	}
	return *o.SizeGb
}

// GetSizeGbOk returns a tuple with the SizeGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetSizeGbOk() (*int32, bool) {
	if o == nil || IsNil(o.SizeGb) {
		return nil, false
	}
	return o.SizeGb, true
}

// HasSizeGb returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) HasSizeGb() bool {
	if o != nil && !IsNil(o.SizeGb) {
		return true
	}

	return false
}

// SetSizeGb gets a reference to the given int32 and assigns it to the SizeGb field.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) SetSizeGb(v int32) {
	o.SizeGb = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetTypeId() int32 {
	if o == nil || IsNil(o.TypeId) {
		var ret int32
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetTypeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) HasTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given int32 and assigns it to the TypeId field.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) SetTypeId(v int32) {
	o.TypeId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) SetType(v string) {
	o.Type = &v
}

// GetIsBootable returns the IsBootable field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetIsBootable() bool {
	if o == nil || IsNil(o.IsBootable) {
		var ret bool
		return ret
	}
	return *o.IsBootable
}

// GetIsBootableOk returns a tuple with the IsBootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) GetIsBootableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBootable) {
		return nil, false
	}
	return o.IsBootable, true
}

// HasIsBootable returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) HasIsBootable() bool {
	if o != nil && !IsNil(o.IsBootable) {
		return true
	}

	return false
}

// SetIsBootable gets a reference to the given bool and assigns it to the IsBootable field.
func (o *KubernetesNodeGroupsInnerNodesInnerDisksInner) SetIsBootable(v bool) {
	o.IsBootable = &v
}

func (o KubernetesNodeGroupsInnerNodesInnerDisksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNodeGroupsInnerNodesInnerDisksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SizeGb) {
		toSerialize["sizeGb"] = o.SizeGb
	}
	if !IsNil(o.TypeId) {
		toSerialize["typeId"] = o.TypeId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsBootable) {
		toSerialize["isBootable"] = o.IsBootable
	}
	return toSerialize, nil
}

type NullableKubernetesNodeGroupsInnerNodesInnerDisksInner struct {
	value *KubernetesNodeGroupsInnerNodesInnerDisksInner
	isSet bool
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerDisksInner) Get() *KubernetesNodeGroupsInnerNodesInnerDisksInner {
	return v.value
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerDisksInner) Set(val *KubernetesNodeGroupsInnerNodesInnerDisksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerDisksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerDisksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupsInnerNodesInnerDisksInner(val *KubernetesNodeGroupsInnerNodesInnerDisksInner) *NullableKubernetesNodeGroupsInnerNodesInnerDisksInner {
	return &NullableKubernetesNodeGroupsInnerNodesInnerDisksInner{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerDisksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerDisksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


