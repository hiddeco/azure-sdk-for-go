//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

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

// PublishedBlueprintsClient contains the methods for the PublishedBlueprints group.
// Don't use this type directly, use NewPublishedBlueprintsClient() instead.
type PublishedBlueprintsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPublishedBlueprintsClient creates a new instance of PublishedBlueprintsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPublishedBlueprintsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PublishedBlueprintsClient, error) {
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
	client := &PublishedBlueprintsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Create - Publish a new version of the blueprint definition with the latest artifacts. Published blueprint definitions are
// immutable.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// versionID - Version of the published blueprint definition.
// options - PublishedBlueprintsClientCreateOptions contains the optional parameters for the PublishedBlueprintsClient.Create
// method.
func (client *PublishedBlueprintsClient) Create(ctx context.Context, resourceScope string, blueprintName string, versionID string, options *PublishedBlueprintsClientCreateOptions) (PublishedBlueprintsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceScope, blueprintName, versionID, options)
	if err != nil {
		return PublishedBlueprintsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PublishedBlueprintsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return PublishedBlueprintsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *PublishedBlueprintsClient) createCreateRequest(ctx context.Context, resourceScope string, blueprintName string, versionID string, options *PublishedBlueprintsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if versionID == "" {
		return nil, errors.New("parameter versionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionId}", url.PathEscape(versionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.PublishedBlueprint != nil {
		return req, runtime.MarshalAsJSON(req, *options.PublishedBlueprint)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *PublishedBlueprintsClient) createHandleResponse(resp *http.Response) (PublishedBlueprintsClientCreateResponse, error) {
	result := PublishedBlueprintsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishedBlueprint); err != nil {
		return PublishedBlueprintsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a published version of a blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// versionID - Version of the published blueprint definition.
// options - PublishedBlueprintsClientDeleteOptions contains the optional parameters for the PublishedBlueprintsClient.Delete
// method.
func (client *PublishedBlueprintsClient) Delete(ctx context.Context, resourceScope string, blueprintName string, versionID string, options *PublishedBlueprintsClientDeleteOptions) (PublishedBlueprintsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceScope, blueprintName, versionID, options)
	if err != nil {
		return PublishedBlueprintsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PublishedBlueprintsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PublishedBlueprintsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *PublishedBlueprintsClient) deleteCreateRequest(ctx context.Context, resourceScope string, blueprintName string, versionID string, options *PublishedBlueprintsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if versionID == "" {
		return nil, errors.New("parameter versionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionId}", url.PathEscape(versionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *PublishedBlueprintsClient) deleteHandleResponse(resp *http.Response) (PublishedBlueprintsClientDeleteResponse, error) {
	result := PublishedBlueprintsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishedBlueprint); err != nil {
		return PublishedBlueprintsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get a published version of a blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// versionID - Version of the published blueprint definition.
// options - PublishedBlueprintsClientGetOptions contains the optional parameters for the PublishedBlueprintsClient.Get method.
func (client *PublishedBlueprintsClient) Get(ctx context.Context, resourceScope string, blueprintName string, versionID string, options *PublishedBlueprintsClientGetOptions) (PublishedBlueprintsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceScope, blueprintName, versionID, options)
	if err != nil {
		return PublishedBlueprintsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PublishedBlueprintsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PublishedBlueprintsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PublishedBlueprintsClient) getCreateRequest(ctx context.Context, resourceScope string, blueprintName string, versionID string, options *PublishedBlueprintsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if versionID == "" {
		return nil, errors.New("parameter versionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionId}", url.PathEscape(versionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PublishedBlueprintsClient) getHandleResponse(resp *http.Response) (PublishedBlueprintsClientGetResponse, error) {
	result := PublishedBlueprintsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishedBlueprint); err != nil {
		return PublishedBlueprintsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List published versions of given blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// options - PublishedBlueprintsClientListOptions contains the optional parameters for the PublishedBlueprintsClient.List
// method.
func (client *PublishedBlueprintsClient) NewListPager(resourceScope string, blueprintName string, options *PublishedBlueprintsClientListOptions) *runtime.Pager[PublishedBlueprintsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PublishedBlueprintsClientListResponse]{
		More: func(page PublishedBlueprintsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PublishedBlueprintsClientListResponse) (PublishedBlueprintsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceScope, blueprintName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PublishedBlueprintsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PublishedBlueprintsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PublishedBlueprintsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PublishedBlueprintsClient) listCreateRequest(ctx context.Context, resourceScope string, blueprintName string, options *PublishedBlueprintsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PublishedBlueprintsClient) listHandleResponse(resp *http.Response) (PublishedBlueprintsClientListResponse, error) {
	result := PublishedBlueprintsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishedBlueprintList); err != nil {
		return PublishedBlueprintsClientListResponse{}, err
	}
	return result, nil
}
