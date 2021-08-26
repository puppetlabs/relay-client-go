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
	"os"
)

// Linger please
var (
	_ _context.Context
)

// WorkflowRunsApiService WorkflowRunsApi service
type WorkflowRunsApiService service

type ApiGetWorkflowRunRequest struct {
	ctx _context.Context
	ApiService *WorkflowRunsApiService
	workflowName string
	workflowRunNumber int32
}


func (r ApiGetWorkflowRunRequest) Execute() (WorkflowRunEntity, *_nethttp.Response, error) {
	return r.ApiService.GetWorkflowRunExecute(r)
}

/*
 * GetWorkflowRun Gets a workflow run accessed with a workflow name and run number
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @param workflowRunNumber Run number of the associated workflow
 * @return ApiGetWorkflowRunRequest
 */
func (a *WorkflowRunsApiService) GetWorkflowRun(ctx _context.Context, workflowName string, workflowRunNumber int32) ApiGetWorkflowRunRequest {
	return ApiGetWorkflowRunRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
		workflowRunNumber: workflowRunNumber,
	}
}

/*
 * Execute executes the request
 * @return WorkflowRunEntity
 */
func (a *WorkflowRunsApiService) GetWorkflowRunExecute(r ApiGetWorkflowRunRequest) (WorkflowRunEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkflowRunEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRunsApiService.GetWorkflowRun")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/runs/{workflowRunNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowRunNumber"+"}", _neturl.PathEscape(parameterToString(r.workflowRunNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.workflowRunNumber < 1 {
		return localVarReturnValue, nil, reportError("workflowRunNumber must be greater than 1")
	}

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

type ApiGetWorkflowRunStepLogRequest struct {
	ctx _context.Context
	ApiService *WorkflowRunsApiService
	workflowName string
	workflowRunNumber int32
	workflowStepName string
	follow *bool
}

func (r ApiGetWorkflowRunStepLogRequest) Follow(follow bool) ApiGetWorkflowRunStepLogRequest {
	r.follow = &follow
	return r
}

func (r ApiGetWorkflowRunStepLogRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.GetWorkflowRunStepLogExecute(r)
}

/*
 * GetWorkflowRunStepLog Returns the log for a workflow step, accessed by workflow name, run number, and step name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @param workflowRunNumber Run number of the associated workflow
 * @param workflowStepName The name of the step in the associated workflow
 * @return ApiGetWorkflowRunStepLogRequest
 */
func (a *WorkflowRunsApiService) GetWorkflowRunStepLog(ctx _context.Context, workflowName string, workflowRunNumber int32, workflowStepName string) ApiGetWorkflowRunStepLogRequest {
	return ApiGetWorkflowRunStepLogRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
		workflowRunNumber: workflowRunNumber,
		workflowStepName: workflowStepName,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *WorkflowRunsApiService) GetWorkflowRunStepLogExecute(r ApiGetWorkflowRunStepLogRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRunsApiService.GetWorkflowRunStepLog")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/runs/{workflowRunNumber}/steps/{workflowStepName}/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowRunNumber"+"}", _neturl.PathEscape(parameterToString(r.workflowRunNumber, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowStepName"+"}", _neturl.PathEscape(parameterToString(r.workflowStepName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.workflowRunNumber < 1 {
		return localVarReturnValue, nil, reportError("workflowRunNumber must be greater than 1")
	}

	if r.follow != nil {
		localVarQueryParams.Add("follow", parameterToString(*r.follow, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/vnd.puppet.relay.v20200615+json"}

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

type ApiGetWorkflowRunsRequest struct {
	ctx _context.Context
	ApiService *WorkflowRunsApiService
	workflowName string
}


func (r ApiGetWorkflowRunsRequest) Execute() (WorkflowRunsSummary, *_nethttp.Response, error) {
	return r.ApiService.GetWorkflowRunsExecute(r)
}

/*
 * GetWorkflowRuns Get all the runs of a workflow
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @return ApiGetWorkflowRunsRequest
 */
func (a *WorkflowRunsApiService) GetWorkflowRuns(ctx _context.Context, workflowName string) ApiGetWorkflowRunsRequest {
	return ApiGetWorkflowRunsRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
	}
}

/*
 * Execute executes the request
 * @return WorkflowRunsSummary
 */
func (a *WorkflowRunsApiService) GetWorkflowRunsExecute(r ApiGetWorkflowRunsRequest) (WorkflowRunsSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkflowRunsSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRunsApiService.GetWorkflowRuns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/runs"
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

type ApiPatchWorkflowRunRequest struct {
	ctx _context.Context
	ApiService *WorkflowRunsApiService
	workflowName string
	workflowRunNumber int32
	updateWorkflowRun *UpdateWorkflowRun
}

func (r ApiPatchWorkflowRunRequest) UpdateWorkflowRun(updateWorkflowRun UpdateWorkflowRun) ApiPatchWorkflowRunRequest {
	r.updateWorkflowRun = &updateWorkflowRun
	return r
}

func (r ApiPatchWorkflowRunRequest) Execute() (WorkflowRunEntity, *_nethttp.Response, error) {
	return r.ApiService.PatchWorkflowRunExecute(r)
}

/*
 * PatchWorkflowRun Update properties of a workflow
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @param workflowRunNumber Run number of the associated workflow
 * @return ApiPatchWorkflowRunRequest
 */
func (a *WorkflowRunsApiService) PatchWorkflowRun(ctx _context.Context, workflowName string, workflowRunNumber int32) ApiPatchWorkflowRunRequest {
	return ApiPatchWorkflowRunRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
		workflowRunNumber: workflowRunNumber,
	}
}

/*
 * Execute executes the request
 * @return WorkflowRunEntity
 */
func (a *WorkflowRunsApiService) PatchWorkflowRunExecute(r ApiPatchWorkflowRunRequest) (WorkflowRunEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkflowRunEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRunsApiService.PatchWorkflowRun")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/runs/{workflowRunNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowRunNumber"+"}", _neturl.PathEscape(parameterToString(r.workflowRunNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.workflowRunNumber < 1 {
		return localVarReturnValue, nil, reportError("workflowRunNumber must be greater than 1")
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
	localVarPostBody = r.updateWorkflowRun
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

type ApiPatchWorkflowRunStepRequest struct {
	ctx _context.Context
	ApiService *WorkflowRunsApiService
	workflowName string
	workflowRunNumber int32
	workflowStepName string
	updateWorkflowRunStep *UpdateWorkflowRunStep
}

func (r ApiPatchWorkflowRunStepRequest) UpdateWorkflowRunStep(updateWorkflowRunStep UpdateWorkflowRunStep) ApiPatchWorkflowRunStepRequest {
	r.updateWorkflowRunStep = &updateWorkflowRunStep
	return r
}

func (r ApiPatchWorkflowRunStepRequest) Execute() (WorkflowRunStep, *_nethttp.Response, error) {
	return r.ApiService.PatchWorkflowRunStepExecute(r)
}

/*
 * PatchWorkflowRunStep Update properties of a workflow run step
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @param workflowRunNumber Run number of the associated workflow
 * @param workflowStepName The name of the step in the associated workflow
 * @return ApiPatchWorkflowRunStepRequest
 */
func (a *WorkflowRunsApiService) PatchWorkflowRunStep(ctx _context.Context, workflowName string, workflowRunNumber int32, workflowStepName string) ApiPatchWorkflowRunStepRequest {
	return ApiPatchWorkflowRunStepRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
		workflowRunNumber: workflowRunNumber,
		workflowStepName: workflowStepName,
	}
}

/*
 * Execute executes the request
 * @return WorkflowRunStep
 */
func (a *WorkflowRunsApiService) PatchWorkflowRunStepExecute(r ApiPatchWorkflowRunStepRequest) (WorkflowRunStep, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkflowRunStep
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRunsApiService.PatchWorkflowRunStep")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/runs/{workflowRunNumber}/steps/{workflowStepName}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowRunNumber"+"}", _neturl.PathEscape(parameterToString(r.workflowRunNumber, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workflowStepName"+"}", _neturl.PathEscape(parameterToString(r.workflowStepName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.workflowRunNumber < 1 {
		return localVarReturnValue, nil, reportError("workflowRunNumber must be greater than 1")
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
	localVarPostBody = r.updateWorkflowRunStep
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

type ApiRunWorkflowRequest struct {
	ctx _context.Context
	ApiService *WorkflowRunsApiService
	workflowName string
	createWorkflowRun *CreateWorkflowRun
}

func (r ApiRunWorkflowRequest) CreateWorkflowRun(createWorkflowRun CreateWorkflowRun) ApiRunWorkflowRequest {
	r.createWorkflowRun = &createWorkflowRun
	return r
}

func (r ApiRunWorkflowRequest) Execute() (WorkflowRunEntity, *_nethttp.Response, error) {
	return r.ApiService.RunWorkflowExecute(r)
}

/*
 * RunWorkflow Runs the given workflow
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowName Workflow name
 * @return ApiRunWorkflowRequest
 */
func (a *WorkflowRunsApiService) RunWorkflow(ctx _context.Context, workflowName string) ApiRunWorkflowRequest {
	return ApiRunWorkflowRequest{
		ApiService: a,
		ctx: ctx,
		workflowName: workflowName,
	}
}

/*
 * Execute executes the request
 * @return WorkflowRunEntity
 */
func (a *WorkflowRunsApiService) RunWorkflowExecute(r ApiRunWorkflowRequest) (WorkflowRunEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkflowRunEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowRunsApiService.RunWorkflow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/workflows/{workflowName}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowName"+"}", _neturl.PathEscape(parameterToString(r.workflowName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.createWorkflowRun
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
