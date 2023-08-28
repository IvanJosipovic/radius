//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220901privatepreview

// AwsCredentialClientCreateOrUpdateResponse contains the response from method AwsCredentialClient.CreateOrUpdate.
type AwsCredentialClientCreateOrUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	AWSCredentialResource

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// AwsCredentialClientDeleteResponse contains the response from method AwsCredentialClient.Delete.
type AwsCredentialClientDeleteResponse struct {
	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// AwsCredentialClientGetResponse contains the response from method AwsCredentialClient.Get.
type AwsCredentialClientGetResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	AWSCredentialResource
}

// AwsCredentialClientListByRootScopeResponse contains the response from method AwsCredentialClient.NewListByRootScopePager.
type AwsCredentialClientListByRootScopeResponse struct {
	// The response of a AWSCredentialResource list operation.
	AWSCredentialResourceListResult
}

// AzureCredentialClientCreateOrUpdateResponse contains the response from method AzureCredentialClient.CreateOrUpdate.
type AzureCredentialClientCreateOrUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	AzureCredentialResource

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// AzureCredentialClientDeleteResponse contains the response from method AzureCredentialClient.Delete.
type AzureCredentialClientDeleteResponse struct {
	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// AzureCredentialClientGetResponse contains the response from method AzureCredentialClient.Get.
type AzureCredentialClientGetResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	AzureCredentialResource
}

// AzureCredentialClientListByRootScopeResponse contains the response from method AzureCredentialClient.NewListByRootScopePager.
type AzureCredentialClientListByRootScopeResponse struct {
	// The response of a AzureCredentialResource list operation.
	AzureCredentialResourceListResult
}

// PlaneTypesClientGetResponse contains the response from method PlaneTypesClient.Get.
type PlaneTypesClientGetResponse struct {
	// UCP PlaneResource.
	PlaneResource
}

// PlanesClientCreateOrUpdateResponse contains the response from method PlanesClient.CreateOrUpdate.
type PlanesClientCreateOrUpdateResponse struct {
	// UCP PlaneResource.
	PlaneResource

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// PlanesClientDeleteResponse contains the response from method PlanesClient.Delete.
type PlanesClientDeleteResponse struct {
	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// PlanesClientGetResponse contains the response from method PlanesClient.Get.
type PlanesClientGetResponse struct {
	// UCP PlaneResource.
	PlaneResource
}

// PlanesClientListByRootScopeResponse contains the response from method PlanesClient.NewListByRootScopePager.
type PlanesClientListByRootScopeResponse struct {
	// The response of a PlaneResource list operation.
	PlaneResourceListResult
}

// ResourceGroupsClientCreateOrUpdateResponse contains the response from method ResourceGroupsClient.CreateOrUpdate.
type ResourceGroupsClientCreateOrUpdateResponse struct {
	// UCP ResourceGroup.
	ResourceGroupResource

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// ResourceGroupsClientDeleteResponse contains the response from method ResourceGroupsClient.Delete.
type ResourceGroupsClientDeleteResponse struct {
	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// ResourceGroupsClientGetResponse contains the response from method ResourceGroupsClient.Get.
type ResourceGroupsClientGetResponse struct {
	// UCP ResourceGroup.
	ResourceGroupResource
}

// ResourceGroupsClientListByRootScopeResponse contains the response from method ResourceGroupsClient.NewListByRootScopePager.
type ResourceGroupsClientListByRootScopeResponse struct {
	// The response of a ResourceGroupResource list operation.
	ResourceGroupResourceListResult
}

