# \AuthenticationAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueToken**](AuthenticationAPI.md#IssueToken) | **Post** /v1/issue-token | Issue token
[**RefreshToken**](AuthenticationAPI.md#RefreshToken) | **Post** /v1/refresh-token | Refresh token



## IssueToken

> Token IssueToken(ctx).Credentials(credentials).Execute()

Issue token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	credentials := *openapiclient.NewCredentials(interface{}(ad6f8c5c-223f-4102-8664-8044b84864e2), interface{}(e0eb8a09-b7fe-4acc-a72b-79ce0dd26068)) // Credentials |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.IssueToken(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.IssueToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueToken`: Token
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.IssueToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshToken

> Token RefreshToken(ctx).RefreshToken(refreshToken).Execute()

Refresh token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	refreshToken := *openapiclient.NewRefreshToken(interface{}(ad6f8c5js6dkj;s43dv3jljhf9fxej4236235tgdgrthreth8044b84864e2)) // RefreshToken |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.RefreshToken(context.Background()).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshToken`: Token
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.RefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshToken** | [**RefreshToken**](RefreshToken.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

