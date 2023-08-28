//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RabbitMqQueuesClient contains the methods for the RabbitMqQueues group.
// Don't use this type directly, use NewRabbitMqQueuesClient() instead.
type RabbitMqQueuesClient struct {
	internal *arm.Client
	rootScope string
}

// NewRabbitMqQueuesClient creates a new instance of RabbitMqQueuesClient with the specified values.
//   - rootScope - The scope in which the resource is present. UCP Scope is /planes/{planeType}/{planeName}/resourceGroup/{resourcegroupID}
//     and Azure resource scope is
//     /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRabbitMqQueuesClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RabbitMqQueuesClient, error) {
	cl, err := arm.NewClient(moduleName+".RabbitMqQueuesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RabbitMqQueuesClient{
		rootScope: rootScope,
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a RabbitMQQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - rabbitMQQueueName - The name of the RabbitMQQueue portable resource resource
//   - resource - Resource create parameters.
//   - options - RabbitMqQueuesClientBeginCreateOrUpdateOptions contains the optional parameters for the RabbitMqQueuesClient.BeginCreateOrUpdate
//     method.
func (client *RabbitMqQueuesClient) BeginCreateOrUpdate(ctx context.Context, rabbitMQQueueName string, resource RabbitMQQueueResource, options *RabbitMqQueuesClientBeginCreateOrUpdateOptions) (*runtime.Poller[RabbitMqQueuesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, rabbitMQQueueName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RabbitMqQueuesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RabbitMqQueuesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a RabbitMQQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
func (client *RabbitMqQueuesClient) createOrUpdate(ctx context.Context, rabbitMQQueueName string, resource RabbitMQQueueResource, options *RabbitMqQueuesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, rabbitMQQueueName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RabbitMqQueuesClient) createOrUpdateCreateRequest(ctx context.Context, rabbitMQQueueName string, resource RabbitMQQueueResource, options *RabbitMqQueuesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Messaging/rabbitMQQueues/{rabbitMQQueueName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQQueueName == "" {
		return nil, errors.New("parameter rabbitMQQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQQueueName}", url.PathEscape(rabbitMQQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
	return nil, err
}
	return req, nil
}

// BeginDelete - Delete a RabbitMQQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - rabbitMQQueueName - The name of the RabbitMQQueue portable resource resource
//   - options - RabbitMqQueuesClientBeginDeleteOptions contains the optional parameters for the RabbitMqQueuesClient.BeginDelete
//     method.
func (client *RabbitMqQueuesClient) BeginDelete(ctx context.Context, rabbitMQQueueName string, options *RabbitMqQueuesClientBeginDeleteOptions) (*runtime.Poller[RabbitMqQueuesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, rabbitMQQueueName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RabbitMqQueuesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RabbitMqQueuesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a RabbitMQQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
func (client *RabbitMqQueuesClient) deleteOperation(ctx context.Context, rabbitMQQueueName string, options *RabbitMqQueuesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, rabbitMQQueueName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RabbitMqQueuesClient) deleteCreateRequest(ctx context.Context, rabbitMQQueueName string, options *RabbitMqQueuesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Messaging/rabbitMQQueues/{rabbitMQQueueName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQQueueName == "" {
		return nil, errors.New("parameter rabbitMQQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQQueueName}", url.PathEscape(rabbitMQQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a RabbitMQQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - rabbitMQQueueName - The name of the RabbitMQQueue portable resource resource
//   - options - RabbitMqQueuesClientGetOptions contains the optional parameters for the RabbitMqQueuesClient.Get method.
func (client *RabbitMqQueuesClient) Get(ctx context.Context, rabbitMQQueueName string, options *RabbitMqQueuesClientGetOptions) (RabbitMqQueuesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, rabbitMQQueueName, options)
	if err != nil {
		return RabbitMqQueuesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RabbitMqQueuesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RabbitMqQueuesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RabbitMqQueuesClient) getCreateRequest(ctx context.Context, rabbitMQQueueName string, options *RabbitMqQueuesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Messaging/rabbitMQQueues/{rabbitMQQueueName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQQueueName == "" {
		return nil, errors.New("parameter rabbitMQQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQQueueName}", url.PathEscape(rabbitMQQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RabbitMqQueuesClient) getHandleResponse(resp *http.Response) (RabbitMqQueuesClientGetResponse, error) {
	result := RabbitMqQueuesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RabbitMQQueueResource); err != nil {
		return RabbitMqQueuesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByScopePager - List RabbitMQQueueResource resources by Scope
//
// Generated from API version 2022-03-15-privatepreview
//   - options - RabbitMqQueuesClientListByScopeOptions contains the optional parameters for the RabbitMqQueuesClient.NewListByScopePager
//     method.
func (client *RabbitMqQueuesClient) NewListByScopePager(options *RabbitMqQueuesClientListByScopeOptions) (*runtime.Pager[RabbitMqQueuesClientListByScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[RabbitMqQueuesClientListByScopeResponse]{
		More: func(page RabbitMqQueuesClientListByScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RabbitMqQueuesClientListByScopeResponse) (RabbitMqQueuesClientListByScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RabbitMqQueuesClientListByScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RabbitMqQueuesClientListByScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RabbitMqQueuesClientListByScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByScopeHandleResponse(resp)
		},
	})
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *RabbitMqQueuesClient) listByScopeCreateRequest(ctx context.Context, options *RabbitMqQueuesClientListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Messaging/rabbitMQQueues"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByScopeHandleResponse handles the ListByScope response.
func (client *RabbitMqQueuesClient) listByScopeHandleResponse(resp *http.Response) (RabbitMqQueuesClientListByScopeResponse, error) {
	result := RabbitMqQueuesClientListByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RabbitMQQueueResourceListResult); err != nil {
		return RabbitMqQueuesClientListByScopeResponse{}, err
	}
	return result, nil
}

// ListSecrets - Lists secrets values for the specified RabbitMQQueue resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - rabbitMQQueueName - The name of the RabbitMQQueue portable resource resource
//   - body - The content of the action request
//   - options - RabbitMqQueuesClientListSecretsOptions contains the optional parameters for the RabbitMqQueuesClient.ListSecrets
//     method.
func (client *RabbitMqQueuesClient) ListSecrets(ctx context.Context, rabbitMQQueueName string, body map[string]any, options *RabbitMqQueuesClientListSecretsOptions) (RabbitMqQueuesClientListSecretsResponse, error) {
	var err error
	req, err := client.listSecretsCreateRequest(ctx, rabbitMQQueueName, body, options)
	if err != nil {
		return RabbitMqQueuesClientListSecretsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RabbitMqQueuesClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RabbitMqQueuesClientListSecretsResponse{}, err
	}
	resp, err := client.listSecretsHandleResponse(httpResp)
	return resp, err
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *RabbitMqQueuesClient) listSecretsCreateRequest(ctx context.Context, rabbitMQQueueName string, body map[string]any, options *RabbitMqQueuesClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Messaging/rabbitMQQueues/{rabbitMQQueueName}/listSecrets"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQQueueName == "" {
		return nil, errors.New("parameter rabbitMQQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQQueueName}", url.PathEscape(rabbitMQQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *RabbitMqQueuesClient) listSecretsHandleResponse(resp *http.Response) (RabbitMqQueuesClientListSecretsResponse, error) {
	result := RabbitMqQueuesClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RabbitMQListSecretsResult); err != nil {
		return RabbitMqQueuesClientListSecretsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a RabbitMQQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - rabbitMQQueueName - The name of the RabbitMQQueue portable resource resource
//   - properties - The resource properties to be updated.
//   - options - RabbitMqQueuesClientBeginUpdateOptions contains the optional parameters for the RabbitMqQueuesClient.BeginUpdate
//     method.
func (client *RabbitMqQueuesClient) BeginUpdate(ctx context.Context, rabbitMQQueueName string, properties RabbitMQQueueResourceUpdate, options *RabbitMqQueuesClientBeginUpdateOptions) (*runtime.Poller[RabbitMqQueuesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, rabbitMQQueueName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RabbitMqQueuesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RabbitMqQueuesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a RabbitMQQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
func (client *RabbitMqQueuesClient) update(ctx context.Context, rabbitMQQueueName string, properties RabbitMQQueueResourceUpdate, options *RabbitMqQueuesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, rabbitMQQueueName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *RabbitMqQueuesClient) updateCreateRequest(ctx context.Context, rabbitMQQueueName string, properties RabbitMQQueueResourceUpdate, options *RabbitMqQueuesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Messaging/rabbitMQQueues/{rabbitMQQueueName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQQueueName == "" {
		return nil, errors.New("parameter rabbitMQQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQQueueName}", url.PathEscape(rabbitMQQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
	return nil, err
}
	return req, nil
}

