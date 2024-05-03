# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtherClaims** | Pointer to **interface{}** |  | [optional] 
**AccessToken** | Pointer to **interface{}** |  | [optional] 
**ExpiresIn** | Pointer to **interface{}** |  | [optional] 
**RefreshExpiresIn** | Pointer to **interface{}** |  | [optional] 
**RefreshToken** | Pointer to **interface{}** |  | [optional] 
**TokenType** | Pointer to **interface{}** |  | [optional] 
**IdToken** | Pointer to **interface{}** |  | [optional] 
**NotBeforePolicy** | Pointer to **interface{}** |  | [optional] 
**SessionState** | Pointer to **interface{}** |  | [optional] 

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

`func (o *Token) GetOtherClaims() interface{}`

GetOtherClaims returns the OtherClaims field if non-nil, zero value otherwise.

### GetOtherClaimsOk

`func (o *Token) GetOtherClaimsOk() (*interface{}, bool)`

GetOtherClaimsOk returns a tuple with the OtherClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherClaims

`func (o *Token) SetOtherClaims(v interface{})`

SetOtherClaims sets OtherClaims field to given value.

### HasOtherClaims

`func (o *Token) HasOtherClaims() bool`

HasOtherClaims returns a boolean if a field has been set.

### GetAccessToken

`func (o *Token) GetAccessToken() interface{}`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Token) GetAccessTokenOk() (*interface{}, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Token) SetAccessToken(v interface{})`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Token) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *Token) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *Token) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetExpiresIn

`func (o *Token) GetExpiresIn() interface{}`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Token) GetExpiresInOk() (*interface{}, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Token) SetExpiresIn(v interface{})`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Token) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### SetExpiresInNil

`func (o *Token) SetExpiresInNil(b bool)`

 SetExpiresInNil sets the value for ExpiresIn to be an explicit nil

### UnsetExpiresIn
`func (o *Token) UnsetExpiresIn()`

UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
### GetRefreshExpiresIn

`func (o *Token) GetRefreshExpiresIn() interface{}`

GetRefreshExpiresIn returns the RefreshExpiresIn field if non-nil, zero value otherwise.

### GetRefreshExpiresInOk

`func (o *Token) GetRefreshExpiresInOk() (*interface{}, bool)`

GetRefreshExpiresInOk returns a tuple with the RefreshExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpiresIn

`func (o *Token) SetRefreshExpiresIn(v interface{})`

SetRefreshExpiresIn sets RefreshExpiresIn field to given value.

### HasRefreshExpiresIn

`func (o *Token) HasRefreshExpiresIn() bool`

HasRefreshExpiresIn returns a boolean if a field has been set.

### SetRefreshExpiresInNil

`func (o *Token) SetRefreshExpiresInNil(b bool)`

 SetRefreshExpiresInNil sets the value for RefreshExpiresIn to be an explicit nil

### UnsetRefreshExpiresIn
`func (o *Token) UnsetRefreshExpiresIn()`

UnsetRefreshExpiresIn ensures that no value is present for RefreshExpiresIn, not even an explicit nil
### GetRefreshToken

`func (o *Token) GetRefreshToken() interface{}`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Token) GetRefreshTokenOk() (*interface{}, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Token) SetRefreshToken(v interface{})`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Token) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *Token) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *Token) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetTokenType

`func (o *Token) GetTokenType() interface{}`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Token) GetTokenTypeOk() (*interface{}, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Token) SetTokenType(v interface{})`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Token) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *Token) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *Token) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
### GetIdToken

`func (o *Token) GetIdToken() interface{}`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *Token) GetIdTokenOk() (*interface{}, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *Token) SetIdToken(v interface{})`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *Token) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### SetIdTokenNil

`func (o *Token) SetIdTokenNil(b bool)`

 SetIdTokenNil sets the value for IdToken to be an explicit nil

### UnsetIdToken
`func (o *Token) UnsetIdToken()`

UnsetIdToken ensures that no value is present for IdToken, not even an explicit nil
### GetNotBeforePolicy

`func (o *Token) GetNotBeforePolicy() interface{}`

GetNotBeforePolicy returns the NotBeforePolicy field if non-nil, zero value otherwise.

### GetNotBeforePolicyOk

`func (o *Token) GetNotBeforePolicyOk() (*interface{}, bool)`

GetNotBeforePolicyOk returns a tuple with the NotBeforePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforePolicy

`func (o *Token) SetNotBeforePolicy(v interface{})`

SetNotBeforePolicy sets NotBeforePolicy field to given value.

### HasNotBeforePolicy

`func (o *Token) HasNotBeforePolicy() bool`

HasNotBeforePolicy returns a boolean if a field has been set.

### SetNotBeforePolicyNil

`func (o *Token) SetNotBeforePolicyNil(b bool)`

 SetNotBeforePolicyNil sets the value for NotBeforePolicy to be an explicit nil

### UnsetNotBeforePolicy
`func (o *Token) UnsetNotBeforePolicy()`

UnsetNotBeforePolicy ensures that no value is present for NotBeforePolicy, not even an explicit nil
### GetSessionState

`func (o *Token) GetSessionState() interface{}`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *Token) GetSessionStateOk() (*interface{}, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *Token) SetSessionState(v interface{})`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *Token) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### SetSessionStateNil

`func (o *Token) SetSessionStateNil(b bool)`

 SetSessionStateNil sets the value for SessionState to be an explicit nil

### UnsetSessionState
`func (o *Token) UnsetSessionState()`

UnsetSessionState ensures that no value is present for SessionState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


