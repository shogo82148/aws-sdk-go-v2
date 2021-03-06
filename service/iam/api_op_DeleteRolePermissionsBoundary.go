// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the permissions boundary for the specified IAM role. Deleting the
// permissions boundary for a role might increase its permissions. For example, it
// might allow anyone who assumes the role to perform all the actions granted in
// its permissions policies.
func (c *Client) DeleteRolePermissionsBoundary(ctx context.Context, params *DeleteRolePermissionsBoundaryInput, optFns ...func(*Options)) (*DeleteRolePermissionsBoundaryOutput, error) {
	stack := middleware.NewStack("DeleteRolePermissionsBoundary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteRolePermissionsBoundaryMiddlewares(stack)
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
	addOpDeleteRolePermissionsBoundaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRolePermissionsBoundary(options.Region), middleware.Before)
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
			OperationName: "DeleteRolePermissionsBoundary",
			Err:           err,
		}
	}
	out := result.(*DeleteRolePermissionsBoundaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRolePermissionsBoundaryInput struct {

	// The name (friendly name, not ARN) of the IAM role from which you want to remove
	// the permissions boundary.
	//
	// This member is required.
	RoleName *string
}

type DeleteRolePermissionsBoundaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteRolePermissionsBoundaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteRolePermissionsBoundary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteRolePermissionsBoundary{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRolePermissionsBoundary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteRolePermissionsBoundary",
	}
}
