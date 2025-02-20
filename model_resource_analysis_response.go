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

// checks if the ResourceAnalysisResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAnalysisResponse{}

// ResourceAnalysisResponse struct for ResourceAnalysisResponse
type ResourceAnalysisResponse struct {
	// emma's internal service type
	Service *string `json:"service,omitempty"`
	Timecode *string `json:"timecode,omitempty"`
	// Total CPU per service, vCPUs
	CpuCoresNumber *float32 `json:"cpuCoresNumber,omitempty"`
	// Total memory per service, GB
	RamTotalAmountGb *float32 `json:"ramTotalAmountGb,omitempty"`
	// Total disk size per service, GB
	DiskSpaceTotalGb *float32 `json:"diskSpaceTotalGb,omitempty"`
	// Dataset type
	Type *string `json:"type,omitempty"`
}

// NewResourceAnalysisResponse instantiates a new ResourceAnalysisResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAnalysisResponse() *ResourceAnalysisResponse {
	this := ResourceAnalysisResponse{}
	return &this
}

// NewResourceAnalysisResponseWithDefaults instantiates a new ResourceAnalysisResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAnalysisResponseWithDefaults() *ResourceAnalysisResponse {
	this := ResourceAnalysisResponse{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ResourceAnalysisResponse) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisResponse) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ResourceAnalysisResponse) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ResourceAnalysisResponse) SetService(v string) {
	o.Service = &v
}

// GetTimecode returns the Timecode field value if set, zero value otherwise.
func (o *ResourceAnalysisResponse) GetTimecode() string {
	if o == nil || IsNil(o.Timecode) {
		var ret string
		return ret
	}
	return *o.Timecode
}

// GetTimecodeOk returns a tuple with the Timecode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisResponse) GetTimecodeOk() (*string, bool) {
	if o == nil || IsNil(o.Timecode) {
		return nil, false
	}
	return o.Timecode, true
}

// HasTimecode returns a boolean if a field has been set.
func (o *ResourceAnalysisResponse) HasTimecode() bool {
	if o != nil && !IsNil(o.Timecode) {
		return true
	}

	return false
}

// SetTimecode gets a reference to the given string and assigns it to the Timecode field.
func (o *ResourceAnalysisResponse) SetTimecode(v string) {
	o.Timecode = &v
}

// GetCpuCoresNumber returns the CpuCoresNumber field value if set, zero value otherwise.
func (o *ResourceAnalysisResponse) GetCpuCoresNumber() float32 {
	if o == nil || IsNil(o.CpuCoresNumber) {
		var ret float32
		return ret
	}
	return *o.CpuCoresNumber
}

// GetCpuCoresNumberOk returns a tuple with the CpuCoresNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisResponse) GetCpuCoresNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.CpuCoresNumber) {
		return nil, false
	}
	return o.CpuCoresNumber, true
}

// HasCpuCoresNumber returns a boolean if a field has been set.
func (o *ResourceAnalysisResponse) HasCpuCoresNumber() bool {
	if o != nil && !IsNil(o.CpuCoresNumber) {
		return true
	}

	return false
}

// SetCpuCoresNumber gets a reference to the given float32 and assigns it to the CpuCoresNumber field.
func (o *ResourceAnalysisResponse) SetCpuCoresNumber(v float32) {
	o.CpuCoresNumber = &v
}

// GetRamTotalAmountGb returns the RamTotalAmountGb field value if set, zero value otherwise.
func (o *ResourceAnalysisResponse) GetRamTotalAmountGb() float32 {
	if o == nil || IsNil(o.RamTotalAmountGb) {
		var ret float32
		return ret
	}
	return *o.RamTotalAmountGb
}

// GetRamTotalAmountGbOk returns a tuple with the RamTotalAmountGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisResponse) GetRamTotalAmountGbOk() (*float32, bool) {
	if o == nil || IsNil(o.RamTotalAmountGb) {
		return nil, false
	}
	return o.RamTotalAmountGb, true
}

// HasRamTotalAmountGb returns a boolean if a field has been set.
func (o *ResourceAnalysisResponse) HasRamTotalAmountGb() bool {
	if o != nil && !IsNil(o.RamTotalAmountGb) {
		return true
	}

	return false
}

// SetRamTotalAmountGb gets a reference to the given float32 and assigns it to the RamTotalAmountGb field.
func (o *ResourceAnalysisResponse) SetRamTotalAmountGb(v float32) {
	o.RamTotalAmountGb = &v
}

// GetDiskSpaceTotalGb returns the DiskSpaceTotalGb field value if set, zero value otherwise.
func (o *ResourceAnalysisResponse) GetDiskSpaceTotalGb() float32 {
	if o == nil || IsNil(o.DiskSpaceTotalGb) {
		var ret float32
		return ret
	}
	return *o.DiskSpaceTotalGb
}

// GetDiskSpaceTotalGbOk returns a tuple with the DiskSpaceTotalGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisResponse) GetDiskSpaceTotalGbOk() (*float32, bool) {
	if o == nil || IsNil(o.DiskSpaceTotalGb) {
		return nil, false
	}
	return o.DiskSpaceTotalGb, true
}

// HasDiskSpaceTotalGb returns a boolean if a field has been set.
func (o *ResourceAnalysisResponse) HasDiskSpaceTotalGb() bool {
	if o != nil && !IsNil(o.DiskSpaceTotalGb) {
		return true
	}

	return false
}

// SetDiskSpaceTotalGb gets a reference to the given float32 and assigns it to the DiskSpaceTotalGb field.
func (o *ResourceAnalysisResponse) SetDiskSpaceTotalGb(v float32) {
	o.DiskSpaceTotalGb = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResourceAnalysisResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResourceAnalysisResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResourceAnalysisResponse) SetType(v string) {
	o.Type = &v
}

func (o ResourceAnalysisResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAnalysisResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Timecode) {
		toSerialize["timecode"] = o.Timecode
	}
	if !IsNil(o.CpuCoresNumber) {
		toSerialize["cpuCoresNumber"] = o.CpuCoresNumber
	}
	if !IsNil(o.RamTotalAmountGb) {
		toSerialize["ramTotalAmountGb"] = o.RamTotalAmountGb
	}
	if !IsNil(o.DiskSpaceTotalGb) {
		toSerialize["diskSpaceTotalGb"] = o.DiskSpaceTotalGb
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableResourceAnalysisResponse struct {
	value *ResourceAnalysisResponse
	isSet bool
}

func (v NullableResourceAnalysisResponse) Get() *ResourceAnalysisResponse {
	return v.value
}

func (v *NullableResourceAnalysisResponse) Set(val *ResourceAnalysisResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAnalysisResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAnalysisResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAnalysisResponse(val *ResourceAnalysisResponse) *NullableResourceAnalysisResponse {
	return &NullableResourceAnalysisResponse{value: val, isSet: true}
}

func (v NullableResourceAnalysisResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAnalysisResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


