//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureADOnlyAuthenticationsClient contains the methods for the AzureADOnlyAuthentications group.
// Don't use this type directly, use NewAzureADOnlyAuthenticationsClient() instead.
type AzureADOnlyAuthenticationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAzureADOnlyAuthenticationsClient creates a new instance of AzureADOnlyAuthenticationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAzureADOnlyAuthenticationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureADOnlyAuthenticationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AzureADOnlyAuthenticationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create or Update a Azure Active Directory only authentication property for the workspaces
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// azureADOnlyAuthenticationName - name of the property
// azureADOnlyAuthenticationInfo - Azure Active Directory Property
// options - AzureADOnlyAuthenticationsClientBeginCreateOptions contains the optional parameters for the AzureADOnlyAuthenticationsClient.BeginCreate
// method.
func (client *AzureADOnlyAuthenticationsClient) BeginCreate(ctx context.Context, resourceGroupName string, workspaceName string, azureADOnlyAuthenticationName AzureADOnlyAuthenticationName, azureADOnlyAuthenticationInfo AzureADOnlyAuthentication, options *AzureADOnlyAuthenticationsClientBeginCreateOptions) (*runtime.Poller[AzureADOnlyAuthenticationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, workspaceName, azureADOnlyAuthenticationName, azureADOnlyAuthenticationInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AzureADOnlyAuthenticationsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AzureADOnlyAuthenticationsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create or Update a Azure Active Directory only authentication property for the workspaces
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *AzureADOnlyAuthenticationsClient) create(ctx context.Context, resourceGroupName string, workspaceName string, azureADOnlyAuthenticationName AzureADOnlyAuthenticationName, azureADOnlyAuthenticationInfo AzureADOnlyAuthentication, options *AzureADOnlyAuthenticationsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, azureADOnlyAuthenticationName, azureADOnlyAuthenticationInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *AzureADOnlyAuthenticationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, azureADOnlyAuthenticationName AzureADOnlyAuthenticationName, azureADOnlyAuthenticationInfo AzureADOnlyAuthentication, options *AzureADOnlyAuthenticationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/azureADOnlyAuthentications/{azureADOnlyAuthenticationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if azureADOnlyAuthenticationName == "" {
		return nil, errors.New("parameter azureADOnlyAuthenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureADOnlyAuthenticationName}", url.PathEscape(string(azureADOnlyAuthenticationName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, azureADOnlyAuthenticationInfo)
}

// Get - Gets a Azure Active Directory only authentication property
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// azureADOnlyAuthenticationName - name of the property
// options - AzureADOnlyAuthenticationsClientGetOptions contains the optional parameters for the AzureADOnlyAuthenticationsClient.Get
// method.
func (client *AzureADOnlyAuthenticationsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, azureADOnlyAuthenticationName AzureADOnlyAuthenticationName, options *AzureADOnlyAuthenticationsClientGetOptions) (AzureADOnlyAuthenticationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, azureADOnlyAuthenticationName, options)
	if err != nil {
		return AzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureADOnlyAuthenticationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AzureADOnlyAuthenticationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, azureADOnlyAuthenticationName AzureADOnlyAuthenticationName, options *AzureADOnlyAuthenticationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/azureADOnlyAuthentications/{azureADOnlyAuthenticationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if azureADOnlyAuthenticationName == "" {
		return nil, errors.New("parameter azureADOnlyAuthenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureADOnlyAuthenticationName}", url.PathEscape(string(azureADOnlyAuthenticationName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureADOnlyAuthenticationsClient) getHandleResponse(resp *http.Response) (AzureADOnlyAuthenticationsClientGetResponse, error) {
	result := AzureADOnlyAuthenticationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureADOnlyAuthentication); err != nil {
		return AzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of Azure Active Directory only authentication property for a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - AzureADOnlyAuthenticationsClientListOptions contains the optional parameters for the AzureADOnlyAuthenticationsClient.List
// method.
func (client *AzureADOnlyAuthenticationsClient) NewListPager(resourceGroupName string, workspaceName string, options *AzureADOnlyAuthenticationsClientListOptions) *runtime.Pager[AzureADOnlyAuthenticationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureADOnlyAuthenticationsClientListResponse]{
		More: func(page AzureADOnlyAuthenticationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureADOnlyAuthenticationsClientListResponse) (AzureADOnlyAuthenticationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureADOnlyAuthenticationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureADOnlyAuthenticationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureADOnlyAuthenticationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AzureADOnlyAuthenticationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AzureADOnlyAuthenticationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/azureADOnlyAuthentications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureADOnlyAuthenticationsClient) listHandleResponse(resp *http.Response) (AzureADOnlyAuthenticationsClientListResponse, error) {
	result := AzureADOnlyAuthenticationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureADOnlyAuthenticationListResult); err != nil {
		return AzureADOnlyAuthenticationsClientListResponse{}, err
	}
	return result, nil
}
