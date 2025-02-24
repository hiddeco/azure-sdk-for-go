//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevops

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PipelineTemplateDefinitionsClient contains the methods for the PipelineTemplateDefinitions group.
// Don't use this type directly, use NewPipelineTemplateDefinitionsClient() instead.
type PipelineTemplateDefinitionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPipelineTemplateDefinitionsClient creates a new instance of PipelineTemplateDefinitionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPipelineTemplateDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelineTemplateDefinitionsClient, error) {
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
	client := &PipelineTemplateDefinitionsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Lists all pipeline templates which can be used to configure an Azure Pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-07-01-preview
// options - PipelineTemplateDefinitionsClientListOptions contains the optional parameters for the PipelineTemplateDefinitionsClient.List
// method.
func (client *PipelineTemplateDefinitionsClient) NewListPager(options *PipelineTemplateDefinitionsClientListOptions) *runtime.Pager[PipelineTemplateDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelineTemplateDefinitionsClientListResponse]{
		More: func(page PipelineTemplateDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelineTemplateDefinitionsClientListResponse) (PipelineTemplateDefinitionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PipelineTemplateDefinitionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PipelineTemplateDefinitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PipelineTemplateDefinitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PipelineTemplateDefinitionsClient) listCreateRequest(ctx context.Context, options *PipelineTemplateDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DevOps/pipelineTemplateDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PipelineTemplateDefinitionsClient) listHandleResponse(resp *http.Response) (PipelineTemplateDefinitionsClientListResponse, error) {
	result := PipelineTemplateDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineTemplateDefinitionListResult); err != nil {
		return PipelineTemplateDefinitionsClientListResponse{}, err
	}
	return result, nil
}
