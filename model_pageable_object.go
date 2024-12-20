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

// checks if the PageableObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageableObject{}

// PageableObject Pageable object containing pagination information
type PageableObject struct {
	// Offset of the first element on the page
	Offset *int64 `json:"offset,omitempty"`
	Sort *SortObject `json:"sort,omitempty"`
	// Number of elements per page
	PageSize *int32 `json:"pageSize,omitempty"`
	// Current page number (0-based index)
	PageNumber *int32 `json:"pageNumber,omitempty"`
	// Indicates if pagination is applied
	Paged *bool `json:"paged,omitempty"`
	// Indicates if pagination is not applied
	Unpaged *bool `json:"unpaged,omitempty"`
}

// NewPageableObject instantiates a new PageableObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableObject() *PageableObject {
	this := PageableObject{}
	return &this
}

// NewPageableObjectWithDefaults instantiates a new PageableObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableObjectWithDefaults() *PageableObject {
	this := PageableObject{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PageableObject) GetOffset() int64 {
	if o == nil || IsNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableObject) GetOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PageableObject) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *PageableObject) SetOffset(v int64) {
	o.Offset = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PageableObject) GetSort() SortObject {
	if o == nil || IsNil(o.Sort) {
		var ret SortObject
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableObject) GetSortOk() (*SortObject, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PageableObject) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given SortObject and assigns it to the Sort field.
func (o *PageableObject) SetSort(v SortObject) {
	o.Sort = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PageableObject) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableObject) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PageableObject) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *PageableObject) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *PageableObject) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableObject) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *PageableObject) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *PageableObject) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPaged returns the Paged field value if set, zero value otherwise.
func (o *PageableObject) GetPaged() bool {
	if o == nil || IsNil(o.Paged) {
		var ret bool
		return ret
	}
	return *o.Paged
}

// GetPagedOk returns a tuple with the Paged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableObject) GetPagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Paged) {
		return nil, false
	}
	return o.Paged, true
}

// HasPaged returns a boolean if a field has been set.
func (o *PageableObject) HasPaged() bool {
	if o != nil && !IsNil(o.Paged) {
		return true
	}

	return false
}

// SetPaged gets a reference to the given bool and assigns it to the Paged field.
func (o *PageableObject) SetPaged(v bool) {
	o.Paged = &v
}

// GetUnpaged returns the Unpaged field value if set, zero value otherwise.
func (o *PageableObject) GetUnpaged() bool {
	if o == nil || IsNil(o.Unpaged) {
		var ret bool
		return ret
	}
	return *o.Unpaged
}

// GetUnpagedOk returns a tuple with the Unpaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableObject) GetUnpagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unpaged) {
		return nil, false
	}
	return o.Unpaged, true
}

// HasUnpaged returns a boolean if a field has been set.
func (o *PageableObject) HasUnpaged() bool {
	if o != nil && !IsNil(o.Unpaged) {
		return true
	}

	return false
}

// SetUnpaged gets a reference to the given bool and assigns it to the Unpaged field.
func (o *PageableObject) SetUnpaged(v bool) {
	o.Unpaged = &v
}

func (o PageableObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageableObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !IsNil(o.Paged) {
		toSerialize["paged"] = o.Paged
	}
	if !IsNil(o.Unpaged) {
		toSerialize["unpaged"] = o.Unpaged
	}
	return toSerialize, nil
}

type NullablePageableObject struct {
	value *PageableObject
	isSet bool
}

func (v NullablePageableObject) Get() *PageableObject {
	return v.value
}

func (v *NullablePageableObject) Set(val *PageableObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableObject(val *PageableObject) *NullablePageableObject {
	return &NullablePageableObject{value: val, isSet: true}
}

func (v NullablePageableObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


