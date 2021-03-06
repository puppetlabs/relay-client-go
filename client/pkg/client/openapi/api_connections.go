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
	"strings"
)

// ConnectionsApiService ConnectionsApi service
type ConnectionsApiService service

type ConnectionsApiCreateConnectionRequest struct {
	ctx                     context.Context
	ApiService              *ConnectionsApiService
	createConnectionRequest *CreateConnectionRequest
}

// Connection to create
func (r ConnectionsApiCreateConnectionRequest) CreateConnectionRequest(createConnectionRequest CreateConnectionRequest) ConnectionsApiCreateConnectionRequest {
	r.createConnectionRequest = &createConnectionRequest
	return r
}

func (r ConnectionsApiCreateConnectionRequest) Execute() (*CreateConnection201Response, *http.Response, error) {
	return r.ApiService.CreateConnectionExecute(r)
}

/*
CreateConnection Create a connection in your account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConnectionsApiCreateConnectionRequest
*/
func (a *ConnectionsApiService) CreateConnection(ctx context.Context) ConnectionsApiCreateConnectionRequest {
	return ConnectionsApiCreateConnectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return CreateConnection201Response
func (a *ConnectionsApiService) CreateConnectionExecute(r ConnectionsApiCreateConnectionRequest) (*CreateConnection201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateConnection201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.CreateConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createConnectionRequest == nil {
		return localVarReturnValue, nil, reportError("createConnectionRequest is required and must be specified")
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
	localVarPostBody = r.createConnectionRequest
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
		var v GetAccessDefaultResponse
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

type ConnectionsApiDeleteConnectionRequest struct {
	ctx          context.Context
	ApiService   *ConnectionsApiService
	connectionId string
}

func (r ConnectionsApiDeleteConnectionRequest) Execute() (*DeletedResource, *http.Response, error) {
	return r.ApiService.DeleteConnectionExecute(r)
}

/*
DeleteConnection Delete a connection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId The connection ID to reference
 @return ConnectionsApiDeleteConnectionRequest
*/
func (a *ConnectionsApiService) DeleteConnection(ctx context.Context, connectionId string) ConnectionsApiDeleteConnectionRequest {
	return ConnectionsApiDeleteConnectionRequest{
		ApiService:   a,
		ctx:          ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return DeletedResource
func (a *ConnectionsApiService) DeleteConnectionExecute(r ConnectionsApiDeleteConnectionRequest) (*DeletedResource, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeletedResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.DeleteConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections/{connectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterToString(r.connectionId, "")), -1)

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
		var v GetAccessDefaultResponse
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

type ConnectionsApiExchangeConnectionOAuth2AuthRequest struct {
	ctx                                 context.Context
	ApiService                          *ConnectionsApiService
	exchangeConnectionOAuth2AuthRequest *ExchangeConnectionOAuth2AuthRequest
}

// Information about the connection to create or update
func (r ConnectionsApiExchangeConnectionOAuth2AuthRequest) ExchangeConnectionOAuth2AuthRequest(exchangeConnectionOAuth2AuthRequest ExchangeConnectionOAuth2AuthRequest) ConnectionsApiExchangeConnectionOAuth2AuthRequest {
	r.exchangeConnectionOAuth2AuthRequest = &exchangeConnectionOAuth2AuthRequest
	return r
}

func (r ConnectionsApiExchangeConnectionOAuth2AuthRequest) Execute() (*ExchangeConnectionOAuth2Auth200Response, *http.Response, error) {
	return r.ApiService.ExchangeConnectionOAuth2AuthExecute(r)
}

/*
ExchangeConnectionOAuth2Auth Complete an OAuth 2.0 exchange and create or update a connection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConnectionsApiExchangeConnectionOAuth2AuthRequest
*/
func (a *ConnectionsApiService) ExchangeConnectionOAuth2Auth(ctx context.Context) ConnectionsApiExchangeConnectionOAuth2AuthRequest {
	return ConnectionsApiExchangeConnectionOAuth2AuthRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ExchangeConnectionOAuth2Auth200Response
func (a *ConnectionsApiService) ExchangeConnectionOAuth2AuthExecute(r ConnectionsApiExchangeConnectionOAuth2AuthRequest) (*ExchangeConnectionOAuth2Auth200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExchangeConnectionOAuth2Auth200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.ExchangeConnectionOAuth2Auth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections/auth/oauth2/exchange"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exchangeConnectionOAuth2AuthRequest == nil {
		return localVarReturnValue, nil, reportError("exchangeConnectionOAuth2AuthRequest is required and must be specified")
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
	localVarPostBody = r.exchangeConnectionOAuth2AuthRequest
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
		var v GetAccessDefaultResponse
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

type ConnectionsApiGetConnectionRequest struct {
	ctx          context.Context
	ApiService   *ConnectionsApiService
	connectionId string
}

func (r ConnectionsApiGetConnectionRequest) Execute() (*GetConnection200Response, *http.Response, error) {
	return r.ApiService.GetConnectionExecute(r)
}

/*
GetConnection Get a connection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId The connection ID to reference
 @return ConnectionsApiGetConnectionRequest
*/
func (a *ConnectionsApiService) GetConnection(ctx context.Context, connectionId string) ConnectionsApiGetConnectionRequest {
	return ConnectionsApiGetConnectionRequest{
		ApiService:   a,
		ctx:          ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return GetConnection200Response
func (a *ConnectionsApiService) GetConnectionExecute(r ConnectionsApiGetConnectionRequest) (*GetConnection200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetConnection200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.GetConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections/{connectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterToString(r.connectionId, "")), -1)

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
		var v GetAccessDefaultResponse
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

type ConnectionsApiGetConnectionsRequest struct {
	ctx        context.Context
	ApiService *ConnectionsApiService
}

func (r ConnectionsApiGetConnectionsRequest) Execute() (*GetConnections200Response, *http.Response, error) {
	return r.ApiService.GetConnectionsExecute(r)
}

/*
GetConnections List all account connections

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConnectionsApiGetConnectionsRequest
*/
func (a *ConnectionsApiService) GetConnections(ctx context.Context) ConnectionsApiGetConnectionsRequest {
	return ConnectionsApiGetConnectionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GetConnections200Response
func (a *ConnectionsApiService) GetConnectionsExecute(r ConnectionsApiGetConnectionsRequest) (*GetConnections200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetConnections200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.GetConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections"

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
		var v GetAccessDefaultResponse
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

type ConnectionsApiReauthorizeConnectionOAuth2AuthRequest struct {
	ctx                                    context.Context
	ApiService                             *ConnectionsApiService
	connectionId                           string
	reauthorizeConnectionOAuth2AuthRequest *ReauthorizeConnectionOAuth2AuthRequest
}

// Reauthorize a connection using an OAuth 2.0 authentication flow
func (r ConnectionsApiReauthorizeConnectionOAuth2AuthRequest) ReauthorizeConnectionOAuth2AuthRequest(reauthorizeConnectionOAuth2AuthRequest ReauthorizeConnectionOAuth2AuthRequest) ConnectionsApiReauthorizeConnectionOAuth2AuthRequest {
	r.reauthorizeConnectionOAuth2AuthRequest = &reauthorizeConnectionOAuth2AuthRequest
	return r
}

func (r ConnectionsApiReauthorizeConnectionOAuth2AuthRequest) Execute() (*ReauthorizeConnectionOAuth2Auth200Response, *http.Response, error) {
	return r.ApiService.ReauthorizeConnectionOAuth2AuthExecute(r)
}

/*
ReauthorizeConnectionOAuth2Auth Reauthorize an existing connection using OAuth 2.0

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId The connection ID to reference
 @return ConnectionsApiReauthorizeConnectionOAuth2AuthRequest
*/
func (a *ConnectionsApiService) ReauthorizeConnectionOAuth2Auth(ctx context.Context, connectionId string) ConnectionsApiReauthorizeConnectionOAuth2AuthRequest {
	return ConnectionsApiReauthorizeConnectionOAuth2AuthRequest{
		ApiService:   a,
		ctx:          ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return ReauthorizeConnectionOAuth2Auth200Response
func (a *ConnectionsApiService) ReauthorizeConnectionOAuth2AuthExecute(r ConnectionsApiReauthorizeConnectionOAuth2AuthRequest) (*ReauthorizeConnectionOAuth2Auth200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReauthorizeConnectionOAuth2Auth200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.ReauthorizeConnectionOAuth2Auth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections/{connectionId}/auth/oauth2"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterToString(r.connectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reauthorizeConnectionOAuth2AuthRequest == nil {
		return localVarReturnValue, nil, reportError("reauthorizeConnectionOAuth2AuthRequest is required and must be specified")
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
	localVarPostBody = r.reauthorizeConnectionOAuth2AuthRequest
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
		var v GetAccessDefaultResponse
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

type ConnectionsApiTestConnectionRequest struct {
	ctx          context.Context
	ApiService   *ConnectionsApiService
	connectionId string
}

func (r ConnectionsApiTestConnectionRequest) Execute() (*TestConnection200Response, *http.Response, error) {
	return r.ApiService.TestConnectionExecute(r)
}

/*
TestConnection Test a connection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId The connection ID to reference
 @return ConnectionsApiTestConnectionRequest
*/
func (a *ConnectionsApiService) TestConnection(ctx context.Context, connectionId string) ConnectionsApiTestConnectionRequest {
	return ConnectionsApiTestConnectionRequest{
		ApiService:   a,
		ctx:          ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return TestConnection200Response
func (a *ConnectionsApiService) TestConnectionExecute(r ConnectionsApiTestConnectionRequest) (*TestConnection200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TestConnection200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.TestConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections/{connectionId}/test"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterToString(r.connectionId, "")), -1)

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
		var v GetAccessDefaultResponse
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

type ConnectionsApiUpdateConnectionRequest struct {
	ctx                     context.Context
	ApiService              *ConnectionsApiService
	connectionId            string
	updateConnectionRequest *UpdateConnectionRequest
}

// Connection info to update
func (r ConnectionsApiUpdateConnectionRequest) UpdateConnectionRequest(updateConnectionRequest UpdateConnectionRequest) ConnectionsApiUpdateConnectionRequest {
	r.updateConnectionRequest = &updateConnectionRequest
	return r
}

func (r ConnectionsApiUpdateConnectionRequest) Execute() (*UpdateConnection200Response, *http.Response, error) {
	return r.ApiService.UpdateConnectionExecute(r)
}

/*
UpdateConnection Update a connection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId The connection ID to reference
 @return ConnectionsApiUpdateConnectionRequest
*/
func (a *ConnectionsApiService) UpdateConnection(ctx context.Context, connectionId string) ConnectionsApiUpdateConnectionRequest {
	return ConnectionsApiUpdateConnectionRequest{
		ApiService:   a,
		ctx:          ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return UpdateConnection200Response
func (a *ConnectionsApiService) UpdateConnectionExecute(r ConnectionsApiUpdateConnectionRequest) (*UpdateConnection200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateConnection200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.UpdateConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections/{connectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterToString(r.connectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateConnectionRequest == nil {
		return localVarReturnValue, nil, reportError("updateConnectionRequest is required and must be specified")
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
	localVarPostBody = r.updateConnectionRequest
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
		var v GetAccessDefaultResponse
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

type ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest struct {
	ctx          context.Context
	ApiService   *ConnectionsApiService
	connectionId string
	url          *string
}

// The URL of a workflow file in a source control repository
func (r ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest) Url(url string) ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest {
	r.url = &url
	return r
}

func (r ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest) Execute() (*ValidateConnectionSourceControlRepositoryURL200Response, *http.Response, error) {
	return r.ApiService.ValidateConnectionSourceControlRepositoryURLExecute(r)
}

/*
ValidateConnectionSourceControlRepositoryURL Validate that the repository workflow URL is accessible

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId The connection ID to reference
 @return ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest
*/
func (a *ConnectionsApiService) ValidateConnectionSourceControlRepositoryURL(ctx context.Context, connectionId string) ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest {
	return ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest{
		ApiService:   a,
		ctx:          ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return ValidateConnectionSourceControlRepositoryURL200Response
func (a *ConnectionsApiService) ValidateConnectionSourceControlRepositoryURLExecute(r ConnectionsApiValidateConnectionSourceControlRepositoryURLRequest) (*ValidateConnectionSourceControlRepositoryURL200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ValidateConnectionSourceControlRepositoryURL200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsApiService.ValidateConnectionSourceControlRepositoryURL")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connections/{connectionId}/capabilities/source-control/validate"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterToString(r.connectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.url == nil {
		return localVarReturnValue, nil, reportError("url is required and must be specified")
	}

	localVarQueryParams.Add("url", parameterToString(*r.url, ""))
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
		var v GetAccessDefaultResponse
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
