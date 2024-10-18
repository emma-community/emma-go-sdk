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

// checks if the KubernetesNodeGroupsInnerNodesInnerDataCenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodeGroupsInnerNodesInnerDataCenter{}

// KubernetesNodeGroupsInnerNodesInnerDataCenter Data center where the worker node is deployed
type KubernetesNodeGroupsInnerNodesInnerDataCenter struct {
	// ID of the cloud provider's data center where the worker node is deployed
	Id *string `json:"id,omitempty"`
	// Name of the cloud provider's data center where the worker node is deployed
	Name *string `json:"name,omitempty"`
	// ID of the cloud provider
	ProviderId *int32 `json:"providerId,omitempty"`
	// Name of the cloud provider
	ProviderName *string `json:"providerName,omitempty"`
	// Location ID
	LocationId *int32 `json:"locationId,omitempty"`
	// Location Name (city or state)
	LocationName *string `json:"locationName,omitempty"`
}

// NewKubernetesNodeGroupsInnerNodesInnerDataCenter instantiates a new KubernetesNodeGroupsInnerNodesInnerDataCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupsInnerNodesInnerDataCenter() *KubernetesNodeGroupsInnerNodesInnerDataCenter {
	this := KubernetesNodeGroupsInnerNodesInnerDataCenter{}
	return &this
}

// NewKubernetesNodeGroupsInnerNodesInnerDataCenterWithDefaults instantiates a new KubernetesNodeGroupsInnerNodesInnerDataCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupsInnerNodesInnerDataCenterWithDefaults() *KubernetesNodeGroupsInnerNodesInnerDataCenter {
	this := KubernetesNodeGroupsInnerNodesInnerDataCenter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) SetName(v string) {
	o.Name = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetLocationId() int32 {
	if o == nil || IsNil(o.LocationId) {
		var ret int32
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetLocationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LocationId) {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) HasLocationId() bool {
	if o != nil && !IsNil(o.LocationId) {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given int32 and assigns it to the LocationId field.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) SetLocationId(v int32) {
	o.LocationId = &v
}

// GetLocationName returns the LocationName field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetLocationName() string {
	if o == nil || IsNil(o.LocationName) {
		var ret string
		return ret
	}
	return *o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) GetLocationNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocationName) {
		return nil, false
	}
	return o.LocationName, true
}

// HasLocationName returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) HasLocationName() bool {
	if o != nil && !IsNil(o.LocationName) {
		return true
	}

	return false
}

// SetLocationName gets a reference to the given string and assigns it to the LocationName field.
func (o *KubernetesNodeGroupsInnerNodesInnerDataCenter) SetLocationName(v string) {
	o.LocationName = &v
}

func (o KubernetesNodeGroupsInnerNodesInnerDataCenter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNodeGroupsInnerNodesInnerDataCenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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
	return toSerialize, nil
}

type NullableKubernetesNodeGroupsInnerNodesInnerDataCenter struct {
	value *KubernetesNodeGroupsInnerNodesInnerDataCenter
	isSet bool
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerDataCenter) Get() *KubernetesNodeGroupsInnerNodesInnerDataCenter {
	return v.value
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerDataCenter) Set(val *KubernetesNodeGroupsInnerNodesInnerDataCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerDataCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerDataCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupsInnerNodesInnerDataCenter(val *KubernetesNodeGroupsInnerNodesInnerDataCenter) *NullableKubernetesNodeGroupsInnerNodesInnerDataCenter {
	return &NullableKubernetesNodeGroupsInnerNodesInnerDataCenter{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerDataCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerDataCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


