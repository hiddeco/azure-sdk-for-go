package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ResourceGuardProxyClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ResourceGuardProxyClient struct {
	BaseClient
}

// NewResourceGuardProxyClient creates an instance of the ResourceGuardProxyClient client.
func NewResourceGuardProxyClient(subscriptionID string) ResourceGuardProxyClient {
	return NewResourceGuardProxyClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewResourceGuardProxyClientWithBaseURI creates an instance of the ResourceGuardProxyClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewResourceGuardProxyClientWithBaseURI(baseURI string, subscriptionID string) ResourceGuardProxyClient {
	return ResourceGuardProxyClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete ResourceGuardProxy under vault
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
func (client ResourceGuardProxyClient) Delete(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceGuardProxyClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, vaultName, resourceGroupName, resourceGuardProxyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ResourceGuardProxyClient) DeletePreparer(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceGuardProxyName": autorest.Encode("path", resourceGuardProxyName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"vaultName":              autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceGuardProxyClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ResourceGuardProxyClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns ResourceGuardProxy under vault and with the name referenced in request
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
func (client ResourceGuardProxyClient) Get(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string) (result ResourceGuardProxyBaseResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceGuardProxyClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, resourceGuardProxyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ResourceGuardProxyClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceGuardProxyName": autorest.Encode("path", resourceGuardProxyName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"vaultName":              autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceGuardProxyClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ResourceGuardProxyClient) GetResponder(resp *http.Response) (result ResourceGuardProxyBaseResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put add or Update ResourceGuardProxy under vault
// Secures vault critical operations
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
func (client ResourceGuardProxyClient) Put(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string) (result ResourceGuardProxyBaseResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceGuardProxyClient.Put")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutPreparer(ctx, vaultName, resourceGroupName, resourceGuardProxyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "Put", resp, "Failure responding to request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client ResourceGuardProxyClient) PutPreparer(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceGuardProxyName": autorest.Encode("path", resourceGuardProxyName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"vaultName":              autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceGuardProxyClient) PutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client ResourceGuardProxyClient) PutResponder(resp *http.Response) (result ResourceGuardProxyBaseResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UnlockDelete secures delete ResourceGuardProxy operations.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// parameters - request body for operation
func (client ResourceGuardProxyClient) UnlockDelete(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters UnlockDeleteRequest) (result UnlockDeleteResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceGuardProxyClient.UnlockDelete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UnlockDeletePreparer(ctx, vaultName, resourceGroupName, resourceGuardProxyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "UnlockDelete", nil, "Failure preparing request")
		return
	}

	resp, err := client.UnlockDeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "UnlockDelete", resp, "Failure sending request")
		return
	}

	result, err = client.UnlockDeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ResourceGuardProxyClient", "UnlockDelete", resp, "Failure responding to request")
		return
	}

	return
}

// UnlockDeletePreparer prepares the UnlockDelete request.
func (client ResourceGuardProxyClient) UnlockDeletePreparer(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters UnlockDeleteRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceGuardProxyName": autorest.Encode("path", resourceGuardProxyName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"vaultName":              autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}/unlockDelete", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UnlockDeleteSender sends the UnlockDelete request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceGuardProxyClient) UnlockDeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UnlockDeleteResponder handles the response to the UnlockDelete request. The method always
// closes the http.Response Body.
func (client ResourceGuardProxyClient) UnlockDeleteResponder(resp *http.Response) (result UnlockDeleteResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
