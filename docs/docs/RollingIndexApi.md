# \RollingIndexApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRollingIndex**](RollingIndexApi.md#CreateRollingIndex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/index | Create One Rolling Index



## CreateRollingIndex

> CreateRollingIndex(ctx, groupId, clusterName, indexRequest IndexRequest).Execute()

Create One Rolling Index



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
    clusterName := "clusterName_example" // string | 
    indexRequest := *openapiclient.NewIndexRequest("Collection_example", "Db_example") // IndexRequest | 

    r, err := sdk.RollingIndexApi.CreateRollingIndex(context.Background(), groupId, clusterName, &indexRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RollingIndexApi.CreateRollingIndex``: %v\n", err)
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
**clusterName** | **string** | Human-readable label that identifies the cluster on which MongoDB Cloud creates an index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRollingIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **indexRequest** | [**IndexRequest**](IndexRequest.md) | Rolling index to create on the specified cluster. | 

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

