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

// checks if the ResourceAnalysisQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAnalysisQuery{}

// ResourceAnalysisQuery struct for ResourceAnalysisQuery
type ResourceAnalysisQuery struct {
	// Query name
	DatasetName string `json:"datasetName"`
	Filters ResourceAnalysisQueryFilters `json:"filters"`
}

type _ResourceAnalysisQuery ResourceAnalysisQuery

// NewResourceAnalysisQuery instantiates a new ResourceAnalysisQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAnalysisQuery(datasetName string, filters ResourceAnalysisQueryFilters) *ResourceAnalysisQuery {
	this := ResourceAnalysisQuery{}
	this.DatasetName = datasetName
	this.Filters = filters
	return &this
}

// NewResourceAnalysisQueryWithDefaults instantiates a new ResourceAnalysisQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAnalysisQueryWithDefaults() *ResourceAnalysisQuery {
	this := ResourceAnalysisQuery{}
	return &this
}

// GetDatasetName returns the DatasetName field value
func (o *ResourceAnalysisQuery) GetDatasetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetName
}

// GetDatasetNameOk returns a tuple with the DatasetName field value
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisQuery) GetDatasetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetName, true
}

// SetDatasetName sets field value
func (o *ResourceAnalysisQuery) SetDatasetName(v string) {
	o.DatasetName = v
}

// GetFilters returns the Filters field value
func (o *ResourceAnalysisQuery) GetFilters() ResourceAnalysisQueryFilters {
	if o == nil {
		var ret ResourceAnalysisQueryFilters
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ResourceAnalysisQuery) GetFiltersOk() (*ResourceAnalysisQueryFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ResourceAnalysisQuery) SetFilters(v ResourceAnalysisQueryFilters) {
	o.Filters = v
}

func (o ResourceAnalysisQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAnalysisQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetName"] = o.DatasetName
	toSerialize["filters"] = o.Filters
	return toSerialize, nil
}

func (o *ResourceAnalysisQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datasetName",
		"filters",
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

	varResourceAnalysisQuery := _ResourceAnalysisQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceAnalysisQuery)

	if err != nil {
		return err
	}

	*o = ResourceAnalysisQuery(varResourceAnalysisQuery)

	return err
}

type NullableResourceAnalysisQuery struct {
	value *ResourceAnalysisQuery
	isSet bool
}

func (v NullableResourceAnalysisQuery) Get() *ResourceAnalysisQuery {
	return v.value
}

func (v *NullableResourceAnalysisQuery) Set(val *ResourceAnalysisQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAnalysisQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAnalysisQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAnalysisQuery(val *ResourceAnalysisQuery) *NullableResourceAnalysisQuery {
	return &NullableResourceAnalysisQuery{value: val, isSet: true}
}

func (v NullableResourceAnalysisQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAnalysisQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

