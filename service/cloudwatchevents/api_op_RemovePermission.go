// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Revokes the permission of another AWS account to be able to put events to the
// specified event bus. Specify the account to revoke by the StatementId value that
// you associated with the account when you granted it permission with
// PutPermission. You can find the StatementId by using DescribeEventBus ().
func (c *Client) RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) {
	stack := middleware.NewStack("RemovePermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRemovePermissionMiddlewares(stack)
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
	addOpRemovePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemovePermission(options.Region), middleware.Before)
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
			OperationName: "RemovePermission",
			Err:           err,
		}
	}
	out := result.(*RemovePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemovePermissionInput struct {

	// The statement ID corresponding to the account that is no longer allowed to put
	// events to the default event bus.
	//
	// This member is required.
	StatementId *string

	// The name of the event bus to revoke permissions for. If you omit this, the
	// default event bus is used.
	EventBusName *string
}

type RemovePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRemovePermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRemovePermission{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemovePermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemovePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "RemovePermission",
	}
}
