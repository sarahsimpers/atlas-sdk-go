/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type EncryptionAtRestUsingCustomerKeyManagementApi interface {

	/*
	ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project

	Returns the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Owner role.

**LIMITED TO M10 OR GREATER:** MongoDB Cloud limits this feature to dedicated cluster tiers of M10 and greater.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest
	*/
	ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject(ctx context.Context, groupId string) EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest

	// ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectExecute executes the request
	//  @return EncryptionAtRest
	ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectExecute(r EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest) (*EncryptionAtRest, *http.Response, error)

	/*
	UpdateEncryptionAtRest Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project

	Updates the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Atlas Admin role.

**LIMITED TO M10 OR GREATER:** MongoDB Cloud limits this feature to dedicated cluster tiers of M10 and greater.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest
	*/
	UpdateEncryptionAtRest(ctx context.Context, groupId string) EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest

	// UpdateEncryptionAtRestExecute executes the request
	//  @return EncryptionAtRest
	UpdateEncryptionAtRestExecute(r EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest) (*EncryptionAtRest, *http.Response, error)
}

// EncryptionAtRestUsingCustomerKeyManagementApiService EncryptionAtRestUsingCustomerKeyManagementApi service
type EncryptionAtRestUsingCustomerKeyManagementApiService service

type EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest struct {
	ctx context.Context
	ApiService EncryptionAtRestUsingCustomerKeyManagementApi
	groupId string
}

func (r EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest) Execute() (*EncryptionAtRest, *http.Response, error) {
	return r.ApiService.ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectExecute(r)
}

/*
ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project

Returns the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Owner role.

**LIMITED TO M10 OR GREATER:** MongoDB Cloud limits this feature to dedicated cluster tiers of M10 and greater.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest
*/
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject(ctx context.Context, groupId string) EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest {
	return EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return EncryptionAtRest
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectExecute(r EncryptionAtRestUsingCustomerKeyManagementApiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest) (*EncryptionAtRest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EncryptionAtRest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EncryptionAtRestUsingCustomerKeyManagementApiService.ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/encryptionAtRest"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest struct {
	ctx context.Context
	ApiService EncryptionAtRestUsingCustomerKeyManagementApi
	groupId string
	encryptionAtRest *EncryptionAtRest
}

// Required parameters depend on whether someone has enabled Encryption at Rest using Customer Key Management:  If you have enabled Encryption at Rest using Customer Key Management (CMK), Atlas requires all of the parameters for the desired encryption provider.  - To use AWS Key Management Service (KMS), MongoDB Cloud requires all the fields in the **awsKms** object. - To use Azure Key Vault, MongoDB Cloud requires all the fields in the **azureKeyVault** object. - To use Google Cloud Key Management Service (KMS), MongoDB Cloud requires all the fields in the **googleCloudKms** object.  If you enabled Encryption at Rest using Customer Key  Management, administrators can pass only the changed fields for the **awsKms**, **azureKeyVault**, or **googleCloudKms** object to update the configuration to this endpoint.
func (r EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest) EncryptionAtRest(encryptionAtRest EncryptionAtRest) EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest {
	r.encryptionAtRest = &encryptionAtRest
	return r
}

func (r EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest) Execute() (*EncryptionAtRest, *http.Response, error) {
	return r.ApiService.UpdateEncryptionAtRestExecute(r)
}

/*
UpdateEncryptionAtRest Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project

Updates the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Atlas Admin role.

**LIMITED TO M10 OR GREATER:** MongoDB Cloud limits this feature to dedicated cluster tiers of M10 and greater.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest
*/
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) UpdateEncryptionAtRest(ctx context.Context, groupId string) EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest {
	return EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return EncryptionAtRest
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) UpdateEncryptionAtRestExecute(r EncryptionAtRestUsingCustomerKeyManagementApiUpdateEncryptionAtRestRequest) (*EncryptionAtRest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EncryptionAtRest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EncryptionAtRestUsingCustomerKeyManagementApiService.UpdateEncryptionAtRest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/encryptionAtRest"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if r.encryptionAtRest == nil {
		return localVarReturnValue, nil, reportError("encryptionAtRest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.encryptionAtRest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
