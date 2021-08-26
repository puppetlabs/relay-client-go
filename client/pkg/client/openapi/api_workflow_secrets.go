/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// WorkflowSecretsApiService WorkflowSecretsApi service
type WorkflowSecretsApiService service

type ApiCreateWorkflowSecretRequest struct {
	ctx _context.Context
	ApiService *WorkflowSecretsApiService
	workflowName string
	workflowSecret *WorkflowSecret
}

func (r ApiCreateWorkflowSecretRequest) WorkflowSecret(workflowSecret WorkflowSecret) ApiCreateWorkflowSecretRequest {
	r.workflowSecret = &workflowSecret
	return r
}

func (r ApiCreateWorkflowSecretRequest) Execute() (Entity, *_nethttp.Response, error) {
	return r.ApiService.CreateWorkflowSecretExecute(r)
}

/*
 * CreateWorkflowSecret Add a new secret to the given workflow
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @return ApiCreateWorkflowSecretRequest
 */
func (a *WorkflowSecretsApiService) CreateWorkflowSecret(ctx _context.Context, workflowName string) ApiCreateWorkflowSecretRequest {
	return ApiCreateWorkflowSecretRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
	}
}

/*
 * Execute executes the request
 * @return Entity
 */
func (a *WorkflowSecretsApiService) CreateWorkflowSecretExecute(r ApiCreateWorkflowSecretRequest) (Entity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Entity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowSecretsApiService.CreateWorkflowSecret")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/secrets"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.workflowSecret == nil {
		return localVarReturnValue, nil, reportError("workflowSecret is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.workflowSecret
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteWorkflowSecretRequest struct {
	ctx _context.Context
	ApiService *WorkflowSecretsApiService
	workflowName string
	workflowSecretName string
}


func (r ApiDeleteWorkflowSecretRequest) Execute() (DeletedResource, *_nethttp.Response, error) {
	return r.ApiService.DeleteWorkflowSecretExecute(r)
}

/*
 * DeleteWorkflowSecret Delete the secret associated with the given workflow and secret name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @param workflowSecretName The name of a workflow secret
 * @return ApiDeleteWorkflowSecretRequest
 */
func (a *WorkflowSecretsApiService) DeleteWorkflowSecret(ctx _context.Context, workflowName string, workflowSecretName string) ApiDeleteWorkflowSecretRequest {
	return ApiDeleteWorkflowSecretRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
		workflowSecretName: workflowSecretName,
	}
}

/*
 * Execute executes the request
 * @return DeletedResource
 */
func (a *WorkflowSecretsApiService) DeleteWorkflowSecretExecute(r ApiDeleteWorkflowSecretRequest) (DeletedResource, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeletedResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowSecretsApiService.DeleteWorkflowSecret")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/secrets/{workflowSecretName}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowSecretName"+"}", _neturl.PathEscape(parameterToString(r.workflowSecretName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListWorkflowSecretsRequest struct {
	ctx _context.Context
	ApiService *WorkflowSecretsApiService
	workflowName string
}


func (r ApiListWorkflowSecretsRequest) Execute() (WorkflowSecretsSummary, *_nethttp.Response, error) {
	return r.ApiService.ListWorkflowSecretsExecute(r)
}

/*
 * ListWorkflowSecrets Get all secrets associated with the given workflow
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @return ApiListWorkflowSecretsRequest
 */
func (a *WorkflowSecretsApiService) ListWorkflowSecrets(ctx _context.Context, workflowName string) ApiListWorkflowSecretsRequest {
	return ApiListWorkflowSecretsRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
	}
}

/*
 * Execute executes the request
 * @return WorkflowSecretsSummary
 */
func (a *WorkflowSecretsApiService) ListWorkflowSecretsExecute(r ApiListWorkflowSecretsRequest) (WorkflowSecretsSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkflowSecretsSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowSecretsApiService.ListWorkflowSecrets")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/secrets"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateWorkflowSecretRequest struct {
	ctx _context.Context
	ApiService *WorkflowSecretsApiService
	workflowName string
	workflowSecretName string
	workflowSecretValue *WorkflowSecretValue
}

func (r ApiUpdateWorkflowSecretRequest) WorkflowSecretValue(workflowSecretValue WorkflowSecretValue) ApiUpdateWorkflowSecretRequest {
	r.workflowSecretValue = &workflowSecretValue
	return r
}

func (r ApiUpdateWorkflowSecretRequest) Execute() (Entity, *_nethttp.Response, error) {
	return r.ApiService.UpdateWorkflowSecretExecute(r)
}

/*
 * UpdateWorkflowSecret Update the secret associated with the given workflow and secret name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @param workflowSecretName The name of a workflow secret
 * @return ApiUpdateWorkflowSecretRequest
 */
func (a *WorkflowSecretsApiService) UpdateWorkflowSecret(ctx _context.Context, workflowName string, workflowSecretName string) ApiUpdateWorkflowSecretRequest {
	return ApiUpdateWorkflowSecretRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
		workflowSecretName: workflowSecretName,
	}
}

/*
 * Execute executes the request
 * @return Entity
 */
func (a *WorkflowSecretsApiService) UpdateWorkflowSecretExecute(r ApiUpdateWorkflowSecretRequest) (Entity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Entity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowSecretsApiService.UpdateWorkflowSecret")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/secrets/{workflowSecretName}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowSecretName"+"}", _neturl.PathEscape(parameterToString(r.workflowSecretName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.workflowSecretValue == nil {
		return localVarReturnValue, nil, reportError("workflowSecretValue is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.workflowSecretValue
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
