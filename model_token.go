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

// checks if the Token type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Token{}

// Token struct for Token
type Token struct {
	// 
	OtherClaims map[string]interface{} `json:"otherClaims,omitempty"`
	// Access token
	AccessToken *string `json:"accessToken,omitempty"`
	// Seconds until access token expiration
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
	// Seconds until refresh token expiration
	RefreshExpiresIn *int32 `json:"refreshExpiresIn,omitempty"`
	// Refresh token
	RefreshToken *string `json:"refreshToken,omitempty"`
	// Token type
	TokenType *string `json:"tokenType,omitempty"`
	// 
	IdToken *int32 `json:"idToken,omitempty"`
	// Timestamp before which the token is not valid
	NotBeforePolicy *int32 `json:"notBeforePolicy,omitempty"`
	// 
	SessionState *string `json:"sessionState,omitempty"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken() *Token {
	this := Token{}
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetOtherClaims returns the OtherClaims field value if set, zero value otherwise.
func (o *Token) GetOtherClaims() map[string]interface{} {
	if o == nil || IsNil(o.OtherClaims) {
		var ret map[string]interface{}
		return ret
	}
	return o.OtherClaims
}

// GetOtherClaimsOk returns a tuple with the OtherClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetOtherClaimsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OtherClaims) {
		return map[string]interface{}{}, false
	}
	return o.OtherClaims, true
}

// HasOtherClaims returns a boolean if a field has been set.
func (o *Token) HasOtherClaims() bool {
	if o != nil && !IsNil(o.OtherClaims) {
		return true
	}

	return false
}

// SetOtherClaims gets a reference to the given map[string]interface{} and assigns it to the OtherClaims field.
func (o *Token) SetOtherClaims(v map[string]interface{}) {
	o.OtherClaims = v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *Token) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *Token) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *Token) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *Token) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *Token) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *Token) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetRefreshExpiresIn returns the RefreshExpiresIn field value if set, zero value otherwise.
func (o *Token) GetRefreshExpiresIn() int32 {
	if o == nil || IsNil(o.RefreshExpiresIn) {
		var ret int32
		return ret
	}
	return *o.RefreshExpiresIn
}

// GetRefreshExpiresInOk returns a tuple with the RefreshExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetRefreshExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshExpiresIn) {
		return nil, false
	}
	return o.RefreshExpiresIn, true
}

// HasRefreshExpiresIn returns a boolean if a field has been set.
func (o *Token) HasRefreshExpiresIn() bool {
	if o != nil && !IsNil(o.RefreshExpiresIn) {
		return true
	}

	return false
}

// SetRefreshExpiresIn gets a reference to the given int32 and assigns it to the RefreshExpiresIn field.
func (o *Token) SetRefreshExpiresIn(v int32) {
	o.RefreshExpiresIn = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *Token) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *Token) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *Token) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *Token) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *Token) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *Token) SetTokenType(v string) {
	o.TokenType = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *Token) GetIdToken() int32 {
	if o == nil || IsNil(o.IdToken) {
		var ret int32
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetIdTokenOk() (*int32, bool) {
	if o == nil || IsNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *Token) HasIdToken() bool {
	if o != nil && !IsNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given int32 and assigns it to the IdToken field.
func (o *Token) SetIdToken(v int32) {
	o.IdToken = &v
}

// GetNotBeforePolicy returns the NotBeforePolicy field value if set, zero value otherwise.
func (o *Token) GetNotBeforePolicy() int32 {
	if o == nil || IsNil(o.NotBeforePolicy) {
		var ret int32
		return ret
	}
	return *o.NotBeforePolicy
}

// GetNotBeforePolicyOk returns a tuple with the NotBeforePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetNotBeforePolicyOk() (*int32, bool) {
	if o == nil || IsNil(o.NotBeforePolicy) {
		return nil, false
	}
	return o.NotBeforePolicy, true
}

// HasNotBeforePolicy returns a boolean if a field has been set.
func (o *Token) HasNotBeforePolicy() bool {
	if o != nil && !IsNil(o.NotBeforePolicy) {
		return true
	}

	return false
}

// SetNotBeforePolicy gets a reference to the given int32 and assigns it to the NotBeforePolicy field.
func (o *Token) SetNotBeforePolicy(v int32) {
	o.NotBeforePolicy = &v
}

// GetSessionState returns the SessionState field value if set, zero value otherwise.
func (o *Token) GetSessionState() string {
	if o == nil || IsNil(o.SessionState) {
		var ret string
		return ret
	}
	return *o.SessionState
}

// GetSessionStateOk returns a tuple with the SessionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetSessionStateOk() (*string, bool) {
	if o == nil || IsNil(o.SessionState) {
		return nil, false
	}
	return o.SessionState, true
}

// HasSessionState returns a boolean if a field has been set.
func (o *Token) HasSessionState() bool {
	if o != nil && !IsNil(o.SessionState) {
		return true
	}

	return false
}

// SetSessionState gets a reference to the given string and assigns it to the SessionState field.
func (o *Token) SetSessionState(v string) {
	o.SessionState = &v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Token) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OtherClaims) {
		toSerialize["otherClaims"] = o.OtherClaims
	}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !IsNil(o.RefreshExpiresIn) {
		toSerialize["refreshExpiresIn"] = o.RefreshExpiresIn
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !IsNil(o.TokenType) {
		toSerialize["tokenType"] = o.TokenType
	}
	if !IsNil(o.IdToken) {
		toSerialize["idToken"] = o.IdToken
	}
	if !IsNil(o.NotBeforePolicy) {
		toSerialize["notBeforePolicy"] = o.NotBeforePolicy
	}
	if !IsNil(o.SessionState) {
		toSerialize["sessionState"] = o.SessionState
	}
	return toSerialize, nil
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


