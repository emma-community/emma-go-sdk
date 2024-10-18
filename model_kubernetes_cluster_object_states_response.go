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

// checks if the KubernetesClusterObjectStatesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesClusterObjectStatesResponse{}

// KubernetesClusterObjectStatesResponse struct for KubernetesClusterObjectStatesResponse
type KubernetesClusterObjectStatesResponse struct {
	// emma's internal metric name
	MetricName *string `json:"metricName,omitempty"`
	// UI group of metrics
	UiMetricGroup *string `json:"uiMetricGroup,omitempty"`
	// UI color for a status
	UiColorGroupId []int32 `json:"uiColorGroupId,omitempty"`
	// UI metric name
	UiMetricName *string `json:"uiMetricName,omitempty"`
	// Name of sub-object
	SubobjectName *string `json:"subobjectName,omitempty"`
	// Sub-object status
	Value []string `json:"value,omitempty"`
}

// NewKubernetesClusterObjectStatesResponse instantiates a new KubernetesClusterObjectStatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterObjectStatesResponse() *KubernetesClusterObjectStatesResponse {
	this := KubernetesClusterObjectStatesResponse{}
	return &this
}

// NewKubernetesClusterObjectStatesResponseWithDefaults instantiates a new KubernetesClusterObjectStatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterObjectStatesResponseWithDefaults() *KubernetesClusterObjectStatesResponse {
	this := KubernetesClusterObjectStatesResponse{}
	return &this
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *KubernetesClusterObjectStatesResponse) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterObjectStatesResponse) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *KubernetesClusterObjectStatesResponse) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *KubernetesClusterObjectStatesResponse) SetMetricName(v string) {
	o.MetricName = &v
}

// GetUiMetricGroup returns the UiMetricGroup field value if set, zero value otherwise.
func (o *KubernetesClusterObjectStatesResponse) GetUiMetricGroup() string {
	if o == nil || IsNil(o.UiMetricGroup) {
		var ret string
		return ret
	}
	return *o.UiMetricGroup
}

// GetUiMetricGroupOk returns a tuple with the UiMetricGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterObjectStatesResponse) GetUiMetricGroupOk() (*string, bool) {
	if o == nil || IsNil(o.UiMetricGroup) {
		return nil, false
	}
	return o.UiMetricGroup, true
}

// HasUiMetricGroup returns a boolean if a field has been set.
func (o *KubernetesClusterObjectStatesResponse) HasUiMetricGroup() bool {
	if o != nil && !IsNil(o.UiMetricGroup) {
		return true
	}

	return false
}

// SetUiMetricGroup gets a reference to the given string and assigns it to the UiMetricGroup field.
func (o *KubernetesClusterObjectStatesResponse) SetUiMetricGroup(v string) {
	o.UiMetricGroup = &v
}

// GetUiColorGroupId returns the UiColorGroupId field value if set, zero value otherwise.
func (o *KubernetesClusterObjectStatesResponse) GetUiColorGroupId() []int32 {
	if o == nil || IsNil(o.UiColorGroupId) {
		var ret []int32
		return ret
	}
	return o.UiColorGroupId
}

// GetUiColorGroupIdOk returns a tuple with the UiColorGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterObjectStatesResponse) GetUiColorGroupIdOk() ([]int32, bool) {
	if o == nil || IsNil(o.UiColorGroupId) {
		return nil, false
	}
	return o.UiColorGroupId, true
}

// HasUiColorGroupId returns a boolean if a field has been set.
func (o *KubernetesClusterObjectStatesResponse) HasUiColorGroupId() bool {
	if o != nil && !IsNil(o.UiColorGroupId) {
		return true
	}

	return false
}

// SetUiColorGroupId gets a reference to the given []int32 and assigns it to the UiColorGroupId field.
func (o *KubernetesClusterObjectStatesResponse) SetUiColorGroupId(v []int32) {
	o.UiColorGroupId = v
}

// GetUiMetricName returns the UiMetricName field value if set, zero value otherwise.
func (o *KubernetesClusterObjectStatesResponse) GetUiMetricName() string {
	if o == nil || IsNil(o.UiMetricName) {
		var ret string
		return ret
	}
	return *o.UiMetricName
}

// GetUiMetricNameOk returns a tuple with the UiMetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterObjectStatesResponse) GetUiMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.UiMetricName) {
		return nil, false
	}
	return o.UiMetricName, true
}

// HasUiMetricName returns a boolean if a field has been set.
func (o *KubernetesClusterObjectStatesResponse) HasUiMetricName() bool {
	if o != nil && !IsNil(o.UiMetricName) {
		return true
	}

	return false
}

// SetUiMetricName gets a reference to the given string and assigns it to the UiMetricName field.
func (o *KubernetesClusterObjectStatesResponse) SetUiMetricName(v string) {
	o.UiMetricName = &v
}

// GetSubobjectName returns the SubobjectName field value if set, zero value otherwise.
func (o *KubernetesClusterObjectStatesResponse) GetSubobjectName() string {
	if o == nil || IsNil(o.SubobjectName) {
		var ret string
		return ret
	}
	return *o.SubobjectName
}

// GetSubobjectNameOk returns a tuple with the SubobjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterObjectStatesResponse) GetSubobjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubobjectName) {
		return nil, false
	}
	return o.SubobjectName, true
}

// HasSubobjectName returns a boolean if a field has been set.
func (o *KubernetesClusterObjectStatesResponse) HasSubobjectName() bool {
	if o != nil && !IsNil(o.SubobjectName) {
		return true
	}

	return false
}

// SetSubobjectName gets a reference to the given string and assigns it to the SubobjectName field.
func (o *KubernetesClusterObjectStatesResponse) SetSubobjectName(v string) {
	o.SubobjectName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KubernetesClusterObjectStatesResponse) GetValue() []string {
	if o == nil || IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterObjectStatesResponse) GetValueOk() ([]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KubernetesClusterObjectStatesResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *KubernetesClusterObjectStatesResponse) SetValue(v []string) {
	o.Value = v
}

func (o KubernetesClusterObjectStatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesClusterObjectStatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.UiMetricGroup) {
		toSerialize["uiMetricGroup"] = o.UiMetricGroup
	}
	if !IsNil(o.UiColorGroupId) {
		toSerialize["uiColorGroupId"] = o.UiColorGroupId
	}
	if !IsNil(o.UiMetricName) {
		toSerialize["uiMetricName"] = o.UiMetricName
	}
	if !IsNil(o.SubobjectName) {
		toSerialize["subobjectName"] = o.SubobjectName
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableKubernetesClusterObjectStatesResponse struct {
	value *KubernetesClusterObjectStatesResponse
	isSet bool
}

func (v NullableKubernetesClusterObjectStatesResponse) Get() *KubernetesClusterObjectStatesResponse {
	return v.value
}

func (v *NullableKubernetesClusterObjectStatesResponse) Set(val *KubernetesClusterObjectStatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterObjectStatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterObjectStatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterObjectStatesResponse(val *KubernetesClusterObjectStatesResponse) *NullableKubernetesClusterObjectStatesResponse {
	return &NullableKubernetesClusterObjectStatesResponse{value: val, isSet: true}
}

func (v NullableKubernetesClusterObjectStatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterObjectStatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


