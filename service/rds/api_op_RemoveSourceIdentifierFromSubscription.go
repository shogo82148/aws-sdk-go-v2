// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a source identifier from an existing RDS event notification
// subscription.
func (c *Client) RemoveSourceIdentifierFromSubscription(ctx context.Context, params *RemoveSourceIdentifierFromSubscriptionInput, optFns ...func(*Options)) (*RemoveSourceIdentifierFromSubscriptionOutput, error) {
	stack := middleware.NewStack("RemoveSourceIdentifierFromSubscription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRemoveSourceIdentifierFromSubscriptionMiddlewares(stack)
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
	addOpRemoveSourceIdentifierFromSubscriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveSourceIdentifierFromSubscription(options.Region), middleware.Before)
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
			OperationName: "RemoveSourceIdentifierFromSubscription",
			Err:           err,
		}
	}
	out := result.(*RemoveSourceIdentifierFromSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RemoveSourceIdentifierFromSubscriptionInput struct {

	// The source identifier to be removed from the subscription, such as the DB
	// instance identifier for a DB instance or the name of a security group.
	//
	// This member is required.
	SourceIdentifier *string

	// The name of the RDS event notification subscription you want to remove a source
	// identifier from.
	//
	// This member is required.
	SubscriptionName *string
}

type RemoveSourceIdentifierFromSubscriptionOutput struct {

	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRemoveSourceIdentifierFromSubscriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRemoveSourceIdentifierFromSubscription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveSourceIdentifierFromSubscription{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveSourceIdentifierFromSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveSourceIdentifierFromSubscription",
	}
}
