/*
Public EMMA API

### About Infrastructure API  **Base URL:** **<u>https://api.emma.ms/external</u>**   This **Infrastructure API** is for managing the emma cloud infrastructure within a project. The API enables you to view, create, edit, and delete _Virtual machines, Spot instances, Applications, Kubernetes clusters, SSH keys, Security groups, and Backup policies_. For creating the resources you can use the endpoints with the dictionaries: _Data centers, Locations, Providers, Operating systems, Virtual machines configurations, Spot instances configurations, Kubernetes clusters configurations._   ### Authentication   #### 1. Create service application   To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.   - **Read** - only GET methods are allowed in the API.   - **Operate** - some operations are allowed with the resources (e.g. _Start, Reboot,_ and _Shutdown_ of the Virtual machines).   - **Manage** - full creating, updating, and deleting of the resources is allowed.    #### 2. Get access token   - Copy the **Client ID** and **Client Secret** in the service application.  - Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.   _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token \\  -H \"Content-Type: application/json\" \\  -d '{\"clientId\": \"YOUR-CLIENT-ID\", \"clientSecret\": \"YOUR-CLIENT-SECRET\"}'  ```  _For Windows:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token ^  -H \"Content-Type: application/json\" ^  -d \"{\\\"clientId\\\": \\\"YOUR-CLIENT-ID\\\", \\\"clientSecret\\\": \\\"YOUR-CLIENT-SECRET\\\"}\"  ```   #### 3. Use access token in requests  The Bearer access token is a text string, included in the request header, for example:   _For Linux / Mac:_  ```  curl -X GET https://api.emma.ms/external/v1/locations -H \"Authorization: Bearer YOUR-ACCESS-TOKEN-HERE\"  ```   Use this token for the API requests.    #### 4. Refresh token  The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).   To get a new access token send a request to the **_/refresh-token** endpoint:    _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/refresh-token \\  -H \"Content-Type: application/json\" \\  -d '{\"refreshToken\": \"YOUR-REFRESH-TOKEN\"}'  ```       ### Possible response status codes   We use standard HTTP response codes to show the success or failure of requests.   `2xx` - successful responses.   `4xx` - client error responses (the response contains an explanation of the error).   `5xx` - server error responses.   The API uses the following status codes:   | Status Code | Description                  | Notes                                                                  |  |-------------|------------------------------|------------------------------------------------------------------------|  | 200         | OK                           | The request was successful.                                             |  | 201         | Created                      | The object was successfully created. This code is only used with objects that are created immediately.  | 204         | No content                   | A successful request, but there is no additional information to send back in the response body (in a case when the object was deleted).    | 400         | Bad Request                  | The request could not be understood by the server. Incoming parameters might not be valid. |  | 401         | Unauthorized            | The client is unauthenticated. The client must authenticate itself to get the requested response. |  | 403         | Forbidden                   | The client does not have access rights to the content.  | 404         | Not Found                    | The requested resource is not found.                                    |  | 409         | Conflict | This response is sent when a request conflicts with the current state of the object (e.g. deleting the security group with the compute instances in it).|  | 422         | Unprocessable Content   | The request was well-formed but was unable to be followed due to incorrect field values (e.g. creation of a virtual machine in the non-existent data center).  |  | 500         | Internal server Error                 | The server could not return the representation due to an internal server error. | 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// GetStatisticalData200Response - struct for GetStatisticalData200Response
type GetStatisticalData200Response struct {
	ArrayOfKubernetesClusterChangingMetricsResponse *[]KubernetesClusterChangingMetricsResponse
	ArrayOfKubernetesClusterCurrentStateResponse *[]KubernetesClusterCurrentStateResponse
	ArrayOfKubernetesClusterMetricsResponse *[]KubernetesClusterMetricsResponse
	ArrayOfKubernetesClusterObjectStatesResponse *[]KubernetesClusterObjectStatesResponse
	ArrayOfKubernetesClusterObjectsResponse *[]KubernetesClusterObjectsResponse
	ArrayOfProductStatisticsResponse *[]ProductStatisticsResponse
	ArrayOfProjectSummaryResponse *[]ProjectSummaryResponse
	ArrayOfResourceAnalysisResponse *[]ResourceAnalysisResponse
	ArrayOfVmAnalyticsResponse *[]VmAnalyticsResponse
	ArrayOfVmMonitoringResponse *[]VmMonitoringResponse
}

// []KubernetesClusterChangingMetricsResponseAsGetStatisticalData200Response is a convenience function that returns []KubernetesClusterChangingMetricsResponse wrapped in GetStatisticalData200Response
func ArrayOfKubernetesClusterChangingMetricsResponseAsGetStatisticalData200Response(v *[]KubernetesClusterChangingMetricsResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfKubernetesClusterChangingMetricsResponse: v,
	}
}

// []KubernetesClusterCurrentStateResponseAsGetStatisticalData200Response is a convenience function that returns []KubernetesClusterCurrentStateResponse wrapped in GetStatisticalData200Response
func ArrayOfKubernetesClusterCurrentStateResponseAsGetStatisticalData200Response(v *[]KubernetesClusterCurrentStateResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfKubernetesClusterCurrentStateResponse: v,
	}
}

// []KubernetesClusterMetricsResponseAsGetStatisticalData200Response is a convenience function that returns []KubernetesClusterMetricsResponse wrapped in GetStatisticalData200Response
func ArrayOfKubernetesClusterMetricsResponseAsGetStatisticalData200Response(v *[]KubernetesClusterMetricsResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfKubernetesClusterMetricsResponse: v,
	}
}

// []KubernetesClusterObjectStatesResponseAsGetStatisticalData200Response is a convenience function that returns []KubernetesClusterObjectStatesResponse wrapped in GetStatisticalData200Response
func ArrayOfKubernetesClusterObjectStatesResponseAsGetStatisticalData200Response(v *[]KubernetesClusterObjectStatesResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfKubernetesClusterObjectStatesResponse: v,
	}
}

// []KubernetesClusterObjectsResponseAsGetStatisticalData200Response is a convenience function that returns []KubernetesClusterObjectsResponse wrapped in GetStatisticalData200Response
func ArrayOfKubernetesClusterObjectsResponseAsGetStatisticalData200Response(v *[]KubernetesClusterObjectsResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfKubernetesClusterObjectsResponse: v,
	}
}

// []ProductStatisticsResponseAsGetStatisticalData200Response is a convenience function that returns []ProductStatisticsResponse wrapped in GetStatisticalData200Response
func ArrayOfProductStatisticsResponseAsGetStatisticalData200Response(v *[]ProductStatisticsResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfProductStatisticsResponse: v,
	}
}

// []ProjectSummaryResponseAsGetStatisticalData200Response is a convenience function that returns []ProjectSummaryResponse wrapped in GetStatisticalData200Response
func ArrayOfProjectSummaryResponseAsGetStatisticalData200Response(v *[]ProjectSummaryResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfProjectSummaryResponse: v,
	}
}

// []ResourceAnalysisResponseAsGetStatisticalData200Response is a convenience function that returns []ResourceAnalysisResponse wrapped in GetStatisticalData200Response
func ArrayOfResourceAnalysisResponseAsGetStatisticalData200Response(v *[]ResourceAnalysisResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfResourceAnalysisResponse: v,
	}
}

// []VmAnalyticsResponseAsGetStatisticalData200Response is a convenience function that returns []VmAnalyticsResponse wrapped in GetStatisticalData200Response
func ArrayOfVmAnalyticsResponseAsGetStatisticalData200Response(v *[]VmAnalyticsResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfVmAnalyticsResponse: v,
	}
}

// []VmMonitoringResponseAsGetStatisticalData200Response is a convenience function that returns []VmMonitoringResponse wrapped in GetStatisticalData200Response
func ArrayOfVmMonitoringResponseAsGetStatisticalData200Response(v *[]VmMonitoringResponse) GetStatisticalData200Response {
	return GetStatisticalData200Response{
		ArrayOfVmMonitoringResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetStatisticalData200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfKubernetesClusterChangingMetricsResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfKubernetesClusterChangingMetricsResponse)
	if err == nil {
		jsonArrayOfKubernetesClusterChangingMetricsResponse, _ := json.Marshal(dst.ArrayOfKubernetesClusterChangingMetricsResponse)
		if string(jsonArrayOfKubernetesClusterChangingMetricsResponse) == "{}" { // empty struct
			dst.ArrayOfKubernetesClusterChangingMetricsResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfKubernetesClusterChangingMetricsResponse); err != nil {
				dst.ArrayOfKubernetesClusterChangingMetricsResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfKubernetesClusterChangingMetricsResponse = nil
	}

	// try to unmarshal data into ArrayOfKubernetesClusterCurrentStateResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfKubernetesClusterCurrentStateResponse)
	if err == nil {
		jsonArrayOfKubernetesClusterCurrentStateResponse, _ := json.Marshal(dst.ArrayOfKubernetesClusterCurrentStateResponse)
		if string(jsonArrayOfKubernetesClusterCurrentStateResponse) == "{}" { // empty struct
			dst.ArrayOfKubernetesClusterCurrentStateResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfKubernetesClusterCurrentStateResponse); err != nil {
				dst.ArrayOfKubernetesClusterCurrentStateResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfKubernetesClusterCurrentStateResponse = nil
	}

	// try to unmarshal data into ArrayOfKubernetesClusterMetricsResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfKubernetesClusterMetricsResponse)
	if err == nil {
		jsonArrayOfKubernetesClusterMetricsResponse, _ := json.Marshal(dst.ArrayOfKubernetesClusterMetricsResponse)
		if string(jsonArrayOfKubernetesClusterMetricsResponse) == "{}" { // empty struct
			dst.ArrayOfKubernetesClusterMetricsResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfKubernetesClusterMetricsResponse); err != nil {
				dst.ArrayOfKubernetesClusterMetricsResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfKubernetesClusterMetricsResponse = nil
	}

	// try to unmarshal data into ArrayOfKubernetesClusterObjectStatesResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfKubernetesClusterObjectStatesResponse)
	if err == nil {
		jsonArrayOfKubernetesClusterObjectStatesResponse, _ := json.Marshal(dst.ArrayOfKubernetesClusterObjectStatesResponse)
		if string(jsonArrayOfKubernetesClusterObjectStatesResponse) == "{}" { // empty struct
			dst.ArrayOfKubernetesClusterObjectStatesResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfKubernetesClusterObjectStatesResponse); err != nil {
				dst.ArrayOfKubernetesClusterObjectStatesResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfKubernetesClusterObjectStatesResponse = nil
	}

	// try to unmarshal data into ArrayOfKubernetesClusterObjectsResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfKubernetesClusterObjectsResponse)
	if err == nil {
		jsonArrayOfKubernetesClusterObjectsResponse, _ := json.Marshal(dst.ArrayOfKubernetesClusterObjectsResponse)
		if string(jsonArrayOfKubernetesClusterObjectsResponse) == "{}" { // empty struct
			dst.ArrayOfKubernetesClusterObjectsResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfKubernetesClusterObjectsResponse); err != nil {
				dst.ArrayOfKubernetesClusterObjectsResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfKubernetesClusterObjectsResponse = nil
	}

	// try to unmarshal data into ArrayOfProductStatisticsResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfProductStatisticsResponse)
	if err == nil {
		jsonArrayOfProductStatisticsResponse, _ := json.Marshal(dst.ArrayOfProductStatisticsResponse)
		if string(jsonArrayOfProductStatisticsResponse) == "{}" { // empty struct
			dst.ArrayOfProductStatisticsResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfProductStatisticsResponse); err != nil {
				dst.ArrayOfProductStatisticsResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfProductStatisticsResponse = nil
	}

	// try to unmarshal data into ArrayOfProjectSummaryResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfProjectSummaryResponse)
	if err == nil {
		jsonArrayOfProjectSummaryResponse, _ := json.Marshal(dst.ArrayOfProjectSummaryResponse)
		if string(jsonArrayOfProjectSummaryResponse) == "{}" { // empty struct
			dst.ArrayOfProjectSummaryResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfProjectSummaryResponse); err != nil {
				dst.ArrayOfProjectSummaryResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfProjectSummaryResponse = nil
	}

	// try to unmarshal data into ArrayOfResourceAnalysisResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfResourceAnalysisResponse)
	if err == nil {
		jsonArrayOfResourceAnalysisResponse, _ := json.Marshal(dst.ArrayOfResourceAnalysisResponse)
		if string(jsonArrayOfResourceAnalysisResponse) == "{}" { // empty struct
			dst.ArrayOfResourceAnalysisResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfResourceAnalysisResponse); err != nil {
				dst.ArrayOfResourceAnalysisResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfResourceAnalysisResponse = nil
	}

	// try to unmarshal data into ArrayOfVmAnalyticsResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfVmAnalyticsResponse)
	if err == nil {
		jsonArrayOfVmAnalyticsResponse, _ := json.Marshal(dst.ArrayOfVmAnalyticsResponse)
		if string(jsonArrayOfVmAnalyticsResponse) == "{}" { // empty struct
			dst.ArrayOfVmAnalyticsResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfVmAnalyticsResponse); err != nil {
				dst.ArrayOfVmAnalyticsResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfVmAnalyticsResponse = nil
	}

	// try to unmarshal data into ArrayOfVmMonitoringResponse
	err = newStrictDecoder(data).Decode(&dst.ArrayOfVmMonitoringResponse)
	if err == nil {
		jsonArrayOfVmMonitoringResponse, _ := json.Marshal(dst.ArrayOfVmMonitoringResponse)
		if string(jsonArrayOfVmMonitoringResponse) == "{}" { // empty struct
			dst.ArrayOfVmMonitoringResponse = nil
		} else {
			if err = validator.Validate(dst.ArrayOfVmMonitoringResponse); err != nil {
				dst.ArrayOfVmMonitoringResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfVmMonitoringResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfKubernetesClusterChangingMetricsResponse = nil
		dst.ArrayOfKubernetesClusterCurrentStateResponse = nil
		dst.ArrayOfKubernetesClusterMetricsResponse = nil
		dst.ArrayOfKubernetesClusterObjectStatesResponse = nil
		dst.ArrayOfKubernetesClusterObjectsResponse = nil
		dst.ArrayOfProductStatisticsResponse = nil
		dst.ArrayOfProjectSummaryResponse = nil
		dst.ArrayOfResourceAnalysisResponse = nil
		dst.ArrayOfVmAnalyticsResponse = nil
		dst.ArrayOfVmMonitoringResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetStatisticalData200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetStatisticalData200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetStatisticalData200Response) MarshalJSON() ([]byte, error) {
	if src.ArrayOfKubernetesClusterChangingMetricsResponse != nil {
		return json.Marshal(&src.ArrayOfKubernetesClusterChangingMetricsResponse)
	}

	if src.ArrayOfKubernetesClusterCurrentStateResponse != nil {
		return json.Marshal(&src.ArrayOfKubernetesClusterCurrentStateResponse)
	}

	if src.ArrayOfKubernetesClusterMetricsResponse != nil {
		return json.Marshal(&src.ArrayOfKubernetesClusterMetricsResponse)
	}

	if src.ArrayOfKubernetesClusterObjectStatesResponse != nil {
		return json.Marshal(&src.ArrayOfKubernetesClusterObjectStatesResponse)
	}

	if src.ArrayOfKubernetesClusterObjectsResponse != nil {
		return json.Marshal(&src.ArrayOfKubernetesClusterObjectsResponse)
	}

	if src.ArrayOfProductStatisticsResponse != nil {
		return json.Marshal(&src.ArrayOfProductStatisticsResponse)
	}

	if src.ArrayOfProjectSummaryResponse != nil {
		return json.Marshal(&src.ArrayOfProjectSummaryResponse)
	}

	if src.ArrayOfResourceAnalysisResponse != nil {
		return json.Marshal(&src.ArrayOfResourceAnalysisResponse)
	}

	if src.ArrayOfVmAnalyticsResponse != nil {
		return json.Marshal(&src.ArrayOfVmAnalyticsResponse)
	}

	if src.ArrayOfVmMonitoringResponse != nil {
		return json.Marshal(&src.ArrayOfVmMonitoringResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetStatisticalData200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfKubernetesClusterChangingMetricsResponse != nil {
		return obj.ArrayOfKubernetesClusterChangingMetricsResponse
	}

	if obj.ArrayOfKubernetesClusterCurrentStateResponse != nil {
		return obj.ArrayOfKubernetesClusterCurrentStateResponse
	}

	if obj.ArrayOfKubernetesClusterMetricsResponse != nil {
		return obj.ArrayOfKubernetesClusterMetricsResponse
	}

	if obj.ArrayOfKubernetesClusterObjectStatesResponse != nil {
		return obj.ArrayOfKubernetesClusterObjectStatesResponse
	}

	if obj.ArrayOfKubernetesClusterObjectsResponse != nil {
		return obj.ArrayOfKubernetesClusterObjectsResponse
	}

	if obj.ArrayOfProductStatisticsResponse != nil {
		return obj.ArrayOfProductStatisticsResponse
	}

	if obj.ArrayOfProjectSummaryResponse != nil {
		return obj.ArrayOfProjectSummaryResponse
	}

	if obj.ArrayOfResourceAnalysisResponse != nil {
		return obj.ArrayOfResourceAnalysisResponse
	}

	if obj.ArrayOfVmAnalyticsResponse != nil {
		return obj.ArrayOfVmAnalyticsResponse
	}

	if obj.ArrayOfVmMonitoringResponse != nil {
		return obj.ArrayOfVmMonitoringResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetStatisticalData200Response struct {
	value *GetStatisticalData200Response
	isSet bool
}

func (v NullableGetStatisticalData200Response) Get() *GetStatisticalData200Response {
	return v.value
}

func (v *NullableGetStatisticalData200Response) Set(val *GetStatisticalData200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatisticalData200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatisticalData200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatisticalData200Response(val *GetStatisticalData200Response) *NullableGetStatisticalData200Response {
	return &NullableGetStatisticalData200Response{value: val, isSet: true}
}

func (v NullableGetStatisticalData200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatisticalData200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


