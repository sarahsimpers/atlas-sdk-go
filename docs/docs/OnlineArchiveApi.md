# \OnlineArchiveApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOnlineArchive**](OnlineArchiveApi.md#CreateOnlineArchive) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive
[**DeleteOnlineArchive**](OnlineArchiveApi.md#DeleteOnlineArchive) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive
[**DownloadOnlineArchiveQueryLogs**](OnlineArchiveApi.md#DownloadOnlineArchiveQueryLogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs
[**GetOnlineArchive**](OnlineArchiveApi.md#GetOnlineArchive) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive
[**ListOnlineArchives**](OnlineArchiveApi.md#ListOnlineArchives) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster
[**UpdateOnlineArchive**](OnlineArchiveApi.md#UpdateOnlineArchive) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive



## CreateOnlineArchive

> OnlineArchive CreateOnlineArchive(ctx, groupId, clusterName, onlineArchive OnlineArchive).Execute()

Create One Online Archive



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
    onlineArchive := *openapiclient.NewOnlineArchive() // OnlineArchive | 

    resp, r, err := sdk.OnlineArchiveApi.CreateOnlineArchive(context.Background(), groupId, clusterName, &onlineArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.CreateOnlineArchive``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateOnlineArchive`: OnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.CreateOnlineArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **onlineArchive** | [**OnlineArchive**](OnlineArchive.md) | Creates one online archive. | 

### Return type

[**OnlineArchive**](OnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOnlineArchive

> map[string]interface{} DeleteOnlineArchive(ctx, groupId, archiveId, clusterName).Execute()

Remove One Online Archive



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
    archiveId := "archiveId_example" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.OnlineArchiveApi.DeleteOnlineArchive(context.Background(), groupId, archiveId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.DeleteOnlineArchive``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteOnlineArchive`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.DeleteOnlineArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to delete. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadOnlineArchiveQueryLogs

> *os.File DownloadOnlineArchiveQueryLogs(ctx, groupId, clusterName).StartDate(startDate).EndDate(endDate).ArchiveOnly(archiveOnly).Execute()

Download Online Archive Query Logs



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
    startDate := int64(1636481348) // int64 |  (optional)
    endDate := int64(1636481348) // int64 |  (optional)
    archiveOnly := true // bool |  (optional) (default to false)

    resp, r, err := sdk.OnlineArchiveApi.DownloadOnlineArchiveQueryLogs(context.Background(), groupId, clusterName).StartDate(startDate).EndDate(endDate).ArchiveOnly(archiveOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.DownloadOnlineArchiveQueryLogs``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DownloadOnlineArchiveQueryLogs`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.DownloadOnlineArchiveQueryLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOnlineArchiveQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **int64** | Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). | 
 **endDate** | **int64** | Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). | 
 **archiveOnly** | **bool** | Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+gzip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnlineArchive

> OnlineArchive GetOnlineArchive(ctx, groupId, archiveId, clusterName).Execute()

Return One Online Archive



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
    archiveId := "archiveId_example" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.OnlineArchiveApi.GetOnlineArchive(context.Background(), groupId, archiveId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.GetOnlineArchive``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetOnlineArchive`: OnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.GetOnlineArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to return. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OnlineArchive**](OnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOnlineArchives

> PaginatedOnlineArchive ListOnlineArchives(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Online Archives for One Cluster



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
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.OnlineArchiveApi.ListOnlineArchives(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.ListOnlineArchives``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListOnlineArchives`: PaginatedOnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.ListOnlineArchives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOnlineArchivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedOnlineArchive**](PaginatedOnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOnlineArchive

> OnlineArchive UpdateOnlineArchive(ctx, groupId, archiveId, clusterName, onlineArchive OnlineArchive).Execute()

Update One Online Archive



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
    archiveId := "archiveId_example" // string | 
    clusterName := "clusterName_example" // string | 
    onlineArchive := *openapiclient.NewOnlineArchive() // OnlineArchive | 

    resp, r, err := sdk.OnlineArchiveApi.UpdateOnlineArchive(context.Background(), groupId, archiveId, clusterName, &onlineArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.UpdateOnlineArchive``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateOnlineArchive`: OnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.UpdateOnlineArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to update. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onlineArchive** | [**OnlineArchive**](OnlineArchive.md) | Updates, pauses, or resumes one online archive. | 

### Return type

[**OnlineArchive**](OnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

