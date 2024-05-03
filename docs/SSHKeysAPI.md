# \SSHKeysAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSshKey**](SSHKeysAPI.md#GetSshKey) | **Get** /v1/ssh-keys/{sshKeyId} | Get SSH key by id
[**SshKeyDelete**](SSHKeysAPI.md#SshKeyDelete) | **Delete** /v1/ssh-keys/{sshKeyId} | Delete SSH keys
[**SshKeyUpdate**](SSHKeysAPI.md#SshKeyUpdate) | **Put** /v1/ssh-keys/{sshKeyId} | Update SSH keys
[**SshKeys**](SSHKeysAPI.md#SshKeys) | **Get** /v1/ssh-keys | Get list of SSH keys
[**SshKeysCreateImport**](SSHKeysAPI.md#SshKeysCreateImport) | **Post** /v1/ssh-keys | Create or import SSH key



## GetSshKey

> SshKey GetSshKey(ctx, sshKeyId).Execute()

Get SSH key by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emma-community/emma-go-sdk"
)

func main() {
	sshKeyId := int32(56) // int32 | SSH key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHKeysAPI.GetSshKey(context.Background(), sshKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.GetSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshKey`: SshKey
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.GetSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | SSH key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SshKey**](SshKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeyDelete

> SshKeyDelete(ctx, sshKeyId).Execute()

Delete SSH keys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emma-community/emma-go-sdk"
)

func main() {
	sshKeyId := int32(56) // int32 | SSH key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SSHKeysAPI.SshKeyDelete(context.Background(), sshKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.SshKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | SSH key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeyUpdate

> SshKey SshKeyUpdate(ctx, sshKeyId).SshKeyUpdate(sshKeyUpdate).Execute()

Update SSH keys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emma-community/emma-go-sdk"
)

func main() {
	sshKeyId := int32(56) // int32 | SSH key ID
	sshKeyUpdate := *openapiclient.NewSshKeyUpdate("mykey") // SshKeyUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHKeysAPI.SshKeyUpdate(context.Background(), sshKeyId).SshKeyUpdate(sshKeyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.SshKeyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeyUpdate`: SshKey
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.SshKeyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | SSH key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshKeyUpdate** | [**SshKeyUpdate**](SshKeyUpdate.md) |  | 

### Return type

[**SshKey**](SshKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeys

> []SshKey SshKeys(ctx).Execute()

Get list of SSH keys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emma-community/emma-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHKeysAPI.SshKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.SshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeys`: []SshKey
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.SshKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysRequest struct via the builder pattern


### Return type

[**[]SshKey**](SshKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysCreateImport

> SshKeysCreateImport201Response SshKeysCreateImport(ctx).SshKeysCreateImportRequest(sshKeysCreateImportRequest).Execute()

Create or import SSH key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emma-community/emma-go-sdk"
)

func main() {
	sshKeysCreateImportRequest := openapiclient.SshKeysCreateImport_request{SshKeyCreate: openapiclient.NewSshKeyCreate("mykey", "RSA")} // SshKeysCreateImportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHKeysAPI.SshKeysCreateImport(context.Background()).SshKeysCreateImportRequest(sshKeysCreateImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.SshKeysCreateImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeysCreateImport`: SshKeysCreateImport201Response
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.SshKeysCreateImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysCreateImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshKeysCreateImportRequest** | [**SshKeysCreateImportRequest**](SshKeysCreateImportRequest.md) |  | 

### Return type

[**SshKeysCreateImport201Response**](SshKeysCreateImport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

