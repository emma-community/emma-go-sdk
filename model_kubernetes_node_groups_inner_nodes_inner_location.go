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

// checks if the KubernetesNodeGroupsInnerNodesInnerLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodeGroupsInnerNodesInnerLocation{}

// KubernetesNodeGroupsInnerNodesInnerLocation Geographic location where the worker node is deployed
type KubernetesNodeGroupsInnerNodesInnerLocation struct {
	// Location ID of the worker node
	Id *int32 `json:"id,omitempty"`
	// Location name (city or state) of the worker node
	Name *string `json:"name,omitempty"`
	// Name of the geographical continent
	Continent *string `json:"continent,omitempty"`
	// Name of the geographical region
	Region *string `json:"region,omitempty"`
	// Approximate latitude of the geographical location
	Latitude *float64 `json:"latitude,omitempty"`
	// Approximate longitude of the geographical location
	Longitude *float64 `json:"longitude,omitempty"`
}

// NewKubernetesNodeGroupsInnerNodesInnerLocation instantiates a new KubernetesNodeGroupsInnerNodesInnerLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupsInnerNodesInnerLocation() *KubernetesNodeGroupsInnerNodesInnerLocation {
	this := KubernetesNodeGroupsInnerNodesInnerLocation{}
	return &this
}

// NewKubernetesNodeGroupsInnerNodesInnerLocationWithDefaults instantiates a new KubernetesNodeGroupsInnerNodesInnerLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupsInnerNodesInnerLocationWithDefaults() *KubernetesNodeGroupsInnerNodesInnerLocation {
	this := KubernetesNodeGroupsInnerNodesInnerLocation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) SetName(v string) {
	o.Name = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetContinent() string {
	if o == nil || IsNil(o.Continent) {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetContinentOk() (*string, bool) {
	if o == nil || IsNil(o.Continent) {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) HasContinent() bool {
	if o != nil && !IsNil(o.Continent) {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) SetContinent(v string) {
	o.Continent = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) SetRegion(v string) {
	o.Region = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *KubernetesNodeGroupsInnerNodesInnerLocation) SetLongitude(v float64) {
	o.Longitude = &v
}

func (o KubernetesNodeGroupsInnerNodesInnerLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNodeGroupsInnerNodesInnerLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Continent) {
		toSerialize["continent"] = o.Continent
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	return toSerialize, nil
}

type NullableKubernetesNodeGroupsInnerNodesInnerLocation struct {
	value *KubernetesNodeGroupsInnerNodesInnerLocation
	isSet bool
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerLocation) Get() *KubernetesNodeGroupsInnerNodesInnerLocation {
	return v.value
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerLocation) Set(val *KubernetesNodeGroupsInnerNodesInnerLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupsInnerNodesInnerLocation(val *KubernetesNodeGroupsInnerNodesInnerLocation) *NullableKubernetesNodeGroupsInnerNodesInnerLocation {
	return &NullableKubernetesNodeGroupsInnerNodesInnerLocation{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupsInnerNodesInnerLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupsInnerNodesInnerLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


