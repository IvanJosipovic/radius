// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/project-radius/radius/pkg/armrpc/api/conv"
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	sm "github.com/project-radius/radius/pkg/armrpc/asyncoperation/statusmanager"
	"github.com/project-radius/radius/pkg/armrpc/rest"
	"github.com/project-radius/radius/pkg/ucp/dataprovider"
	"github.com/project-radius/radius/pkg/ucp/resources"
	"github.com/project-radius/radius/pkg/ucp/store"
)

// Operation is the base operation controller.
type Operation[P interface {
	*T
	conv.ResourceDataModel
}, T any] struct {
	options Options

	resourceOptions ResourceOptions[T]
}

// NewOperation creates BaseController instance.
func NewOperation[P interface {
	*T
	conv.ResourceDataModel
}, T any](options Options, resourceOptions ResourceOptions[T]) Operation[P, T] {
	return Operation[P, T]{options, resourceOptions}
}

// Options gets the options for this controller.
func (b *Operation[P, T]) Options() *Options {
	return &b.options
}

// StorageClient gets storage client for this controller.
func (b *Operation[P, T]) StorageClient() store.StorageClient {
	return b.options.StorageClient
}

func (b *Operation[P, T]) DataProvider() dataprovider.DataStorageProvider {
	return b.options.DataProvider
}

// ResourceType gets the resource type for this controller.
func (b *Operation[P, T]) ResourceType() string {
	return b.options.ResourceType
}

// DeploymentProcessor gets the deployment processor for this controller.
func (b *Operation[P, T]) StatusManager() sm.StatusManager {
	return b.options.StatusManager
}

// GetResourceFromRequest extracts and deserializes from HTTP request body to datamodel.
func (c *Operation[P, T]) GetResourceFromRequest(ctx context.Context, req *http.Request) (*T, error) {
	content, err := ReadJSONBody(req)
	if err != nil {
		return nil, err
	}

	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	dm, err := c.resourceOptions.RequestConverter(content, serviceCtx.APIVersion)
	if err != nil {
		return nil, err
	}
	return dm, nil
}

// GetResource is the helper to get the resource via storage client.
func (c *Operation[P, T]) GetResource(ctx context.Context, id resources.ID) (out *T, etag string, err error) {
	etag = ""
	out = new(T)
	var res *store.Object
	if res, err = c.StorageClient().Get(ctx, id.String()); err == nil {
		if err = res.As(out); err == nil {
			etag = res.ETag
			return
		}
	}

	out = nil
	if errors.Is(&store.ErrNotFound{}, err) {
		err = nil
	}
	return
}

// SaveResource is the helper to save the resource via storage client.
func (c *Operation[P, T]) SaveResource(ctx context.Context, id string, in *T, etag string) (string, error) {
	nr := &store.Object{
		Metadata: store.Metadata{
			ID: id,
		},
		Data: in,
	}
	err := c.StorageClient().Save(ctx, nr, store.WithETag(etag))
	if err != nil {
		return "", err
	}
	return nr.ETag, nil
}

// PrepareResource validates incoming request and populate the metadata to new resource.
func (c *Operation[P, T]) PrepareResource(ctx context.Context, req *http.Request, newResource *T, oldResource *T, etag string) (rest.Response, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	if req.Method == http.MethodPatch && oldResource == nil {
		return rest.NewNotFoundResponse(serviceCtx.ResourceID), nil
	}

	if err := ValidateETag(*serviceCtx, etag); err != nil {
		return rest.NewPreconditionFailedResponse(serviceCtx.ResourceID.String(), err.Error()), nil
	}

	if oldResource != nil {
		state := P(oldResource).ProvisioningState()
		if !state.IsTerminal() {
			return rest.NewConflictResponse(fmt.Sprintf(InProgressStateMessageFormat, state)), nil
		}
	}

	if newResource != nil {
		P(newResource).UpdateMetadata(serviceCtx)
		var oldSystemData *v1.SystemData
		if oldResource != nil {
			oldSystemData = P(oldResource).GetSystemData()
		}

		*P(newResource).GetSystemData() = v1.UpdateSystemData(oldSystemData, serviceCtx.SystemData())
	}

	return nil, nil
}

// PrepareAsyncOperation saves the initial state and queue the async operation.
func (c *Operation[P, T]) PrepareAsyncOperation(ctx context.Context, newResource *T, initialState v1.ProvisioningState, asyncTimeout time.Duration, etag *string) (rest.Response, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	P(newResource).SetProvisioningState(initialState)

	var err error
	*etag, err = c.SaveResource(ctx, serviceCtx.ResourceID.String(), newResource, *etag)
	if err != nil {
		return nil, err
	}

	if err := c.StatusManager().QueueAsyncOperation(ctx, serviceCtx, asyncTimeout); err != nil {
		P(newResource).SetProvisioningState(v1.ProvisioningStateFailed)
		_, rbErr := c.SaveResource(ctx, serviceCtx.ResourceID.String(), newResource, *etag)
		if rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	return nil, nil
}

// ConstructSyncResponse constructs synchronous API response.
func (c *Operation[P, T]) ConstructSyncResponse(ctx context.Context, method, etag string, resource *T) (rest.Response, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	versioned, err := c.resourceOptions.ResponseConverter(resource, serviceCtx.APIVersion)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{"ETag": etag}
	return rest.NewOKResponseWithHeaders(versioned, headers), nil
}

// ConstructAsyncResponse constructs asynchronous API response.
func (c *Operation[P, T]) ConstructAsyncResponse(ctx context.Context, method, etag string, resource *T) (rest.Response, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	versioned, err := c.resourceOptions.ResponseConverter(resource, serviceCtx.APIVersion)
	if err != nil {
		return nil, err
	}

	respCode := http.StatusAccepted
	if method == http.MethodPut {
		respCode = http.StatusCreated
	}

	return rest.NewAsyncOperationResponse(versioned, serviceCtx.Location, respCode, serviceCtx.ResourceID, serviceCtx.OperationID, serviceCtx.APIVersion, "", ""), nil
}

// RequestConverter returns the request converter function for this controller.
func (b *Operation[P, T]) RequestConverter() conv.ConvertToDataModel[T] {
	return b.resourceOptions.RequestConverter
}

// ResponseConverter returns the response converter function for this controller.
func (b *Operation[P, T]) ResponseConverter() conv.ConvertToAPIModel[T] {
	return b.resourceOptions.ResponseConverter
}

// DeleteFilters returns the set of filters to execute on delete operations.
func (b *Operation[P, T]) DeleteFilters() []DeleteFilter[T] {
	return b.resourceOptions.DeleteFilters
}

// DeleteFilters returns the set of filters to execute on update (PUT/PATCH) operations.
func (b *Operation[P, T]) UpdateFilters() []UpdateFilter[T] {
	return b.resourceOptions.UpdateFilters
}
