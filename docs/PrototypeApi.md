# \PrototypeApi

All URIs are relative to *https://quay.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationPrototypePermission**](PrototypeApi.md#CreateOrganizationPrototypePermission) | **Post** /api/v1/organization/{orgname}/prototypes | 
[**DeleteOrganizationPrototypePermission**](PrototypeApi.md#DeleteOrganizationPrototypePermission) | **Delete** /api/v1/organization/{orgname}/prototypes/{prototypeid} | 
[**GetOrganizationPrototypePermissions**](PrototypeApi.md#GetOrganizationPrototypePermissions) | **Get** /api/v1/organization/{orgname}/prototypes | 
[**UpdateOrganizationPrototypePermission**](PrototypeApi.md#UpdateOrganizationPrototypePermission) | **Put** /api/v1/organization/{orgname}/prototypes/{prototypeid} | 



## CreateOrganizationPrototypePermission

> CreateOrganizationPrototypePermission(ctx, orgname).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgname := "orgname_example" // string | The name of the organization
    body := *openapiclient.NewNewPrototype("Role_example", *openapiclient.NewNewPrototypeDelegate("Kind_example", "Name_example")) // NewPrototype | Request body contents.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrototypeApi.CreateOrganizationPrototypePermission(context.Background(), orgname).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrototypeApi.CreateOrganizationPrototypePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPrototypePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewPrototype**](NewPrototype.md) | Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPrototypePermission

> DeleteOrganizationPrototypePermission(ctx, orgname, prototypeid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgname := "orgname_example" // string | The name of the organization
    prototypeid := "prototypeid_example" // string | The ID of the prototype

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrototypeApi.DeleteOrganizationPrototypePermission(context.Background(), orgname, prototypeid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrototypeApi.DeleteOrganizationPrototypePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**prototypeid** | **string** | The ID of the prototype | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPrototypePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPrototypePermissions

> GetOrganizationPrototypePermissions(ctx, orgname).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgname := "orgname_example" // string | The name of the organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrototypeApi.GetOrganizationPrototypePermissions(context.Background(), orgname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrototypeApi.GetOrganizationPrototypePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPrototypePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPrototypePermission

> UpdateOrganizationPrototypePermission(ctx, orgname, prototypeid).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgname := "orgname_example" // string | The name of the organization
    prototypeid := "prototypeid_example" // string | The ID of the prototype
    body := *openapiclient.NewPrototypeUpdate("Role_example") // PrototypeUpdate | Request body contents.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrototypeApi.UpdateOrganizationPrototypePermission(context.Background(), orgname, prototypeid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrototypeApi.UpdateOrganizationPrototypePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**prototypeid** | **string** | The ID of the prototype | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPrototypePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PrototypeUpdate**](PrototypeUpdate.md) | Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

