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

// checks if the KubernetesNodeGroupsInnerNodesInnerSecurityGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodeGroupsInnerNodesInnerSecurityGroup{}

// KubernetesNodeGroupsInnerNodesInnerSecurityGroup Security group of the worker node
type KubernetesNodeGroupsInnerNodesInnerSecurityGroup struct {
	// ID of the security group
	Id *int32 `json:"id,omitempty"`
	// Name of the security group
	Name *string `json:"name,omitempty"`
}

// NewKubernetesNodeGroupsInnerNodesInnerSecurityGroup instantiates a new KubernetesNodeGroupsInnerNodesInnerSecurityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupsInnerNodesInnerSecurityGroup() *KubernetesNodeGroupsInnerNodesInnerSecurityGroup {
	this := KubernetesNodeGroupsInnerNodesInnerSecurityGroup{}
	return &this
}

// NewKubernetesNodeGroupsInnerNodesInnerSecurityGroupWithDefaults instantiates a new KubernetesNodeGroupsInnerNodesInnerSecurityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupsInnerNodesInnerSecurityGroupWithDefaults() *KubernetesNodeGroupsInnerNodesInnerSecurityGroup {
	this := KubernetesNodeGroupsInnerNodesInnerSecurityGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) SetName(v string) {
	o.Name = &v
}

func (o KubernetesNodeGroupsInnerNodesInnerSecurityGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNodeGroupsInnerNodesInnerSecurityGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup struct {
	value *KubernetesNodeGroupsInnerNodesInnerSecurityGroup
	isSet bool
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup) Get() *KubernetesNodeGroupsInnerNodesInnerSecurityGroup {
	return v.value
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup) Set(val *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup(val *KubernetesNodeGroupsInnerNodesInnerSecurityGroup) *NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup {
	return &NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerSecurityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


