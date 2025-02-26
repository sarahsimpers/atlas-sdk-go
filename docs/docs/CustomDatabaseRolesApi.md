# \CustomDatabaseRolesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomDatabaseRole**](CustomDatabaseRolesApi.md#CreateCustomDatabaseRole) | **Post** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Create One Custom Role
[**DeleteCustomDatabaseRole**](CustomDatabaseRolesApi.md#DeleteCustomDatabaseRole) | **Delete** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project
[**GetCustomDatabaseRole**](CustomDatabaseRolesApi.md#GetCustomDatabaseRole) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project
[**ListCustomDatabaseRoles**](CustomDatabaseRolesApi.md#ListCustomDatabaseRoles) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project
[**UpdateCustomDatabaseRole**](CustomDatabaseRolesApi.md#UpdateCustomDatabaseRole) | **Patch** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project



## CreateCustomDatabaseRole

> CustomDBRole CreateCustomDatabaseRole(ctx, groupId, customDBRole CustomDBRole).Execute()

Create One Custom Role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    customDBRole := *openapiclient.NewCustomDBRole("RoleName_example") // CustomDBRole | 

    resp, r, err := sdk.CustomDatabaseRolesApi.CreateCustomDatabaseRole(context.Background(), groupId, &customDBRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.CreateCustomDatabaseRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateCustomDatabaseRole`: CustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.CreateCustomDatabaseRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomDatabaseRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDBRole** | [**CustomDBRole**](CustomDBRole.md) | Creates one custom role in the specified project. | 

### Return type

[**CustomDBRole**](CustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomDatabaseRole

> DeleteCustomDatabaseRole(ctx, groupId, roleName).Execute()

Remove One Custom Role from One Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    roleName := "roleName_example" // string | 

    r, err := sdk.CustomDatabaseRolesApi.DeleteCustomDatabaseRole(context.Background(), groupId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.DeleteCustomDatabaseRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomDatabaseRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDatabaseRole

> CustomDBRole GetCustomDatabaseRole(ctx, groupId, roleName).Execute()

Return One Custom Role in One Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    roleName := "roleName_example" // string | 

    resp, r, err := sdk.CustomDatabaseRolesApi.GetCustomDatabaseRole(context.Background(), groupId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.GetCustomDatabaseRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetCustomDatabaseRole`: CustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.GetCustomDatabaseRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDatabaseRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomDBRole**](CustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomDatabaseRoles

> []CustomDBRole ListCustomDatabaseRoles(ctx, groupId).Execute()

Return All Custom Roles in One Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CustomDatabaseRolesApi.ListCustomDatabaseRoles(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.ListCustomDatabaseRoles``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListCustomDatabaseRoles`: []CustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.ListCustomDatabaseRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomDatabaseRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomDBRole**](CustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomDatabaseRole

> CustomDBRole UpdateCustomDatabaseRole(ctx, groupId, roleName, updateCustomDBRole UpdateCustomDBRole).Execute()

Update One Custom Role in One Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    roleName := "roleName_example" // string | 
    updateCustomDBRole := *openapiclient.NewUpdateCustomDBRole() // UpdateCustomDBRole | 

    resp, r, err := sdk.CustomDatabaseRolesApi.UpdateCustomDatabaseRole(context.Background(), groupId, roleName, &updateCustomDBRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.UpdateCustomDatabaseRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateCustomDatabaseRole`: CustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.UpdateCustomDatabaseRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomDatabaseRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCustomDBRole** | [**UpdateCustomDBRole**](UpdateCustomDBRole.md) | Updates one custom role in the specified project. | 

### Return type

[**CustomDBRole**](CustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

