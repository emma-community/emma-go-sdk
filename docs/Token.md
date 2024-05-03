# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtherClaims** | Pointer to **map[string]interface{}** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**IdToken** | Pointer to **int32** |  | [optional] 
**NotBeforePolicy** | Pointer to **int32** |  | [optional] 
**SessionState** | Pointer to **string** |  | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtherClaims

`func (o *Token) GetOtherClaims() map[string]interface{}`

GetOtherClaims returns the OtherClaims field if non-nil, zero value otherwise.

### GetOtherClaimsOk

`func (o *Token) GetOtherClaimsOk() (*map[string]interface{}, bool)`

GetOtherClaimsOk returns a tuple with the OtherClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherClaims

`func (o *Token) SetOtherClaims(v map[string]interface{})`

SetOtherClaims sets OtherClaims field to given value.

### HasOtherClaims

`func (o *Token) HasOtherClaims() bool`

HasOtherClaims returns a boolean if a field has been set.

### GetAccessToken

`func (o *Token) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Token) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Token) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Token) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *Token) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Token) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Token) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Token) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshExpiresIn

`func (o *Token) GetRefreshExpiresIn() int32`

GetRefreshExpiresIn returns the RefreshExpiresIn field if non-nil, zero value otherwise.

### GetRefreshExpiresInOk

`func (o *Token) GetRefreshExpiresInOk() (*int32, bool)`

GetRefreshExpiresInOk returns a tuple with the RefreshExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpiresIn

`func (o *Token) SetRefreshExpiresIn(v int32)`

SetRefreshExpiresIn sets RefreshExpiresIn field to given value.

### HasRefreshExpiresIn

`func (o *Token) HasRefreshExpiresIn() bool`

HasRefreshExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *Token) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Token) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Token) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Token) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenType

`func (o *Token) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Token) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Token) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Token) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetIdToken

`func (o *Token) GetIdToken() int32`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *Token) GetIdTokenOk() (*int32, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *Token) SetIdToken(v int32)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *Token) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetNotBeforePolicy

`func (o *Token) GetNotBeforePolicy() int32`

GetNotBeforePolicy returns the NotBeforePolicy field if non-nil, zero value otherwise.

### GetNotBeforePolicyOk

`func (o *Token) GetNotBeforePolicyOk() (*int32, bool)`

GetNotBeforePolicyOk returns a tuple with the NotBeforePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforePolicy

`func (o *Token) SetNotBeforePolicy(v int32)`

SetNotBeforePolicy sets NotBeforePolicy field to given value.

### HasNotBeforePolicy

`func (o *Token) HasNotBeforePolicy() bool`

HasNotBeforePolicy returns a boolean if a field has been set.

### GetSessionState

`func (o *Token) GetSessionState() string`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *Token) GetSessionStateOk() (*string, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *Token) SetSessionState(v string)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *Token) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


