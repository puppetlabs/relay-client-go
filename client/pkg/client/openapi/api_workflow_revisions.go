/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// WorkflowRevisionsApiService WorkflowRevisionsApi service
type WorkflowRevisionsApiService service

type WorkflowRevisionsApiGetLatestWorkflowRevisionRequest struct {
	ctx          context.Context
	ApiService   *WorkflowRevisionsApiService
	workflowName string
}

func (r WorkflowRevisionsApiGetLatestWorkflowRevisionRequest) Execute() (*Entity, *http.Response, error) {
	return r.ApiService.GetLatestWorkflowRevisionExecute(r)
}

/*
GetLatestWorkflowRevision Retrieve the latest workflow revision

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workflowName Workflow name
 @return WorkflowRevisionsApiGetLatestWorkflowRevisionRequest
*/
func (a *WorkflowRevisionsApiService) GetLatestWorkflowRevision(ctx context.Context, workflowName string) WorkflowRevisionsApiGetLatestWorkflowRevisionRequest {
	return WorkflowRevisionsApiGetLatestWorkflowRevisionRequest{
		ApiService:   a,
		ctx:          ctx,
		workflowName: workflowName,
	}
}

// Execute executes the request
//  @return Entity
func (a *WorkflowRevisionsApiService) GetLatestWorkflowRevisionExecute(r WorkflowRevisionsApiGetLatestWorkflowRevisionRequest) (*Entity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Entity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRevisionsApiService.GetLatestWorkflowRevision")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/revisions/latest"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", url.PathEscape(parameterToString(r.workflowName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WorkflowRevisionsApiGetWorkflowRevisionRequest struct {
	ctx              context.Context
	ApiService       *WorkflowRevisionsApiService
	workflowName     string
	workflowRevision string
}

func (r WorkflowRevisionsApiGetWorkflowRevisionRequest) Execute() (*Entity, *http.Response, error) {
	return r.ApiService.GetWorkflowRevisionExecute(r)
}

/*
GetWorkflowRevision Retrieve workflow revision

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workflowName Workflow name
 @param workflowRevision The workflow revision to reference
 @return WorkflowRevisionsApiGetWorkflowRevisionRequest
*/
func (a *WorkflowRevisionsApiService) GetWorkflowRevision(ctx context.Context, workflowName string, workflowRevision string) WorkflowRevisionsApiGetWorkflowRevisionRequest {
	return WorkflowRevisionsApiGetWorkflowRevisionRequest{
		ApiService:       a,
		ctx:              ctx,
		workflowName:     workflowName,
		workflowRevision: workflowRevision,
	}
}

// Execute executes the request
//  @return Entity
func (a *WorkflowRevisionsApiService) GetWorkflowRevisionExecute(r WorkflowRevisionsApiGetWorkflowRevisionRequest) (*Entity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Entity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRevisionsApiService.GetWorkflowRevision")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/revisions/{workflowRevision}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", url.PathEscape(parameterToString(r.workflowName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowRevision"+"}", url.PathEscape(parameterToString(r.workflowRevision, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WorkflowRevisionsApiPostWorkflowRevisionRequest struct {
	ctx          context.Context
	ApiService   *WorkflowRevisionsApiService
	workflowName string
	body         **os.File
}

// The workflow yaml content
func (r WorkflowRevisionsApiPostWorkflowRevisionRequest) Body(body *os.File) WorkflowRevisionsApiPostWorkflowRevisionRequest {
	r.body = &body
	return r
}

func (r WorkflowRevisionsApiPostWorkflowRevisionRequest) Execute() (*Entity, *http.Response, error) {
	return r.ApiService.PostWorkflowRevisionExecute(r)
}

/*
PostWorkflowRevision Update the workflow revision

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workflowName Workflow name
 @return WorkflowRevisionsApiPostWorkflowRevisionRequest
*/
func (a *WorkflowRevisionsApiService) PostWorkflowRevision(ctx context.Context, workflowName string) WorkflowRevisionsApiPostWorkflowRevisionRequest {
	return WorkflowRevisionsApiPostWorkflowRevisionRequest{
		ApiService:   a,
		ctx:          ctx,
		workflowName: workflowName,
	}
}

// Execute executes the request
//  @return Entity
func (a *WorkflowRevisionsApiService) PostWorkflowRevisionExecute(r WorkflowRevisionsApiPostWorkflowRevisionRequest) (*Entity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Entity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRevisionsApiService.PostWorkflowRevision")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/revisions"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", url.PathEscape(parameterToString(r.workflowName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.puppet.relay.v20200615+yaml"}

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
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
