// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Batch operation to revoke permissions from the principal.
func (c *Client) BatchRevokePermissions(ctx context.Context, params *BatchRevokePermissionsInput, optFns ...func(*Options)) (*BatchRevokePermissionsOutput, error) {
	stack := middleware.NewStack("BatchRevokePermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchRevokePermissionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpBatchRevokePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchRevokePermissions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "BatchRevokePermissions",
			Err:           err,
		}
	}
	out := result.(*BatchRevokePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchRevokePermissionsInput struct {

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string

	// A list of up to 20 entries for resource permissions to be revoked by batch
	// operation to the principal.
	//
	// This member is required.
	Entries []*types.BatchPermissionsRequestEntry
}

type BatchRevokePermissionsOutput struct {

	// A list of failures to revoke permissions to the resources.
	Failures []*types.BatchPermissionsFailureEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchRevokePermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchRevokePermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchRevokePermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchRevokePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "BatchRevokePermissions",
	}
}
