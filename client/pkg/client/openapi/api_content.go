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
)

// ContentApiService ContentApi service
type ContentApiService service

type ContentApiPostAzureAppRegistrationRequest struct {
	ctx                             context.Context
	ApiService                      *ContentApiService
	postAzureAppRegistrationRequest *PostAzureAppRegistrationRequest
}

// Retrieve an Azure app registration automation payload
func (r ContentApiPostAzureAppRegistrationRequest) PostAzureAppRegistrationRequest(postAzureAppRegistrationRequest PostAzureAppRegistrationRequest) ContentApiPostAzureAppRegistrationRequest {
	r.postAzureAppRegistrationRequest = &postAzureAppRegistrationRequest
	return r
}

func (r ContentApiPostAzureAppRegistrationRequest) Execute() (*PostAzureAppRegistration200Response, *http.Response, error) {
	return r.ApiService.PostAzureAppRegistrationExecute(r)
}

/*
PostAzureAppRegistration Create a new Azure App Registration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentApiPostAzureAppRegistrationRequest
*/
func (a *ContentApiService) PostAzureAppRegistration(ctx context.Context) ContentApiPostAzureAppRegistrationRequest {
	return ContentApiPostAzureAppRegistrationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PostAzureAppRegistration200Response
func (a *ContentApiService) PostAzureAppRegistrationExecute(r ContentApiPostAzureAppRegistrationRequest) (*PostAzureAppRegistration200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PostAzureAppRegistration200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.PostAzureAppRegistration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/templated-content/azure-app-registration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postAzureAppRegistrationRequest == nil {
		return localVarReturnValue, nil, reportError("postAzureAppRegistrationRequest is required and must be specified")
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
	localVarPostBody = r.postAzureAppRegistrationRequest
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